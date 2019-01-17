package ethereum

import (
	"context"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/queue"
	"github.com/centrifuge/go-centrifuge/transactions"
	"math/big"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	logging "github.com/ipfs/go-log"
	"github.com/satori/go.uuid"
)

const (
	transactionUnderpriced = errors.Error("replacement transaction underpriced")
	nonceTooLow            = errors.Error("nonce too low")
)

var log = logging.Logger("geth-client")
var gc Client
var gcMu sync.RWMutex

// Config defines functions to get ethereum details
type Config interface {
	GetEthereumGasPrice() *big.Int
	GetEthereumGasLimit() uint64
	GetEthereumNodeURL() string
	GetEthereumAccount(accountName string) (account *config.AccountConfig, err error)
	GetEthereumIntervalRetry() time.Duration
	GetEthereumMaxRetries() int
	GetTxPoolAccessEnabled() bool
	GetEthereumContextReadWaitTimeout() time.Duration
}

// DefaultWaitForTransactionMiningContext returns context with timeout for write operations
func DefaultWaitForTransactionMiningContext(d time.Duration) (ctx context.Context, cancelFunc context.CancelFunc) {
	toBeDone := time.Now().Add(d)
	return context.WithDeadline(context.Background(), toBeDone)
}

// Client can be implemented by any chain client
type Client interface {

	// GetEthClient returns the ethereum client
	GetEthClient() *ethclient.Client

	// GetNodeURL returns the node url
	GetNodeURL() *url.URL

	// GetTxOpts returns a cached options if available else creates and returns new options
	GetTxOpts(accountName string) (*bind.TransactOpts, error)

	// SubmitTransactionWithRetries submits transaction to the ethereum chain
	// Blocking Function that sends transaction using reflection wrapped in a retrial block. It is based on the transactionUnderpriced error,
	// meaning that a transaction is being attempted to run twice, and the logic is to override the existing one. As we have constant
	// gas prices that means that a concurrent transaction race condition event has happened.
	// - contractMethod: Contract Method that implements GenericEthereumAsset (usually autogenerated binding from abi)
	// - params: Arbitrary number of parameters that are passed to the function fname call
	// Note: contractMethod must always return "*types.Transaction, error"
	SubmitTransactionWithRetries(contractMethod interface{}, opts *bind.TransactOpts, params ...interface{}) (tx *types.Transaction, err error)

	// GetGethCallOpts returns the Call options with default
	GetGethCallOpts() (*bind.CallOpts, context.CancelFunc)

	// TransactionByHash returns a Ethereum transaction
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)

	// TransactionReceipt return receipt of a transaction
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

// gethClient implements Client for Ethereum
type gethClient struct {
	client    *ethclient.Client
	rpcClient *rpc.Client
	host      *url.URL
	accounts  map[string]*bind.TransactOpts
	accMu     sync.Mutex // accMu to protect accounts
	config    Config
	txService transactions.Service
	queue *queue.Server

	// txMu to ensure one transaction at a time per client
	txMu sync.Mutex
}

// NewGethClient returns an gethClient which implements Client
func NewGethClient(config Config,transService transactions.Service,queue *queue.Server) (Client, error) {
	log.Info("Opening connection to Ethereum:", config.GetEthereumNodeURL())
	u, err := url.Parse(config.GetEthereumNodeURL())
	if err != nil {
		return nil, errors.New("failed to parse ethereum node URL: %v", err)
	}

	c, err := rpc.Dial(u.String())
	if err != nil {
		return nil, errors.New("failed to connect to ethereum node: %v", err)
	}

	return &gethClient{
		client:    ethclient.NewClient(c),
		rpcClient: c,
		host:      u,
		accounts:  make(map[string]*bind.TransactOpts),
		txMu:      sync.Mutex{},
		accMu:     sync.Mutex{},
		config:    config,
		txService: transService,
		queue: queue,
	}, nil
}

// SetClient sets the Client
// Note that this is a singleton and is the same connection for the whole application.
func SetClient(client Client) {
	gcMu.Lock()
	defer gcMu.Unlock()
	gc = client
}

// GetClient returns the current Client
func GetClient() Client {
	gcMu.RLock()
	defer gcMu.RUnlock()
	return gc
}

// defaultReadContext returns context with timeout for read operations
func (gc *gethClient) defaultReadContext() (ctx context.Context, cancelFunc context.CancelFunc) {
	toBeDone := time.Now().Add(gc.config.GetEthereumContextReadWaitTimeout())
	return context.WithDeadline(context.Background(), toBeDone)
}

// GetTxOpts returns a cached options if available else creates and returns new options
func (gc *gethClient) GetTxOpts(accountName string) (*bind.TransactOpts, error) {
	gc.accMu.Lock()
	defer gc.accMu.Unlock()

	if opts, ok := gc.accounts[accountName]; ok {
		return opts, nil
	}

	txOpts, err := gc.getGethTxOpts(accountName)
	if err != nil {
		return nil, err
	}

	gc.accounts[accountName] = txOpts
	return txOpts, nil
}

// GetEthClient returns the ethereum client
func (gc *gethClient) GetEthClient() *ethclient.Client {
	return gc.client
}

// GetNodeURL returns the node url
func (gc *gethClient) GetNodeURL() *url.URL {
	return gc.host
}

// TransactionByHash returns a Ethereum transaction
func (gc *gethClient) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return gc.client.TransactionByHash(ctx, hash)
}

// TransactionReceipt return receipt of a transaction
func (gc *gethClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return gc.client.TransactionReceipt(ctx, txHash)
}

// getGethTxOpts retrieves the geth transaction options for the given account name. The account name influences which configuration
// is used.
func (gc *gethClient) getGethTxOpts(accountName string) (*bind.TransactOpts, error) {
	account, err := gc.config.GetEthereumAccount(accountName)
	if err != nil {
		return nil, errors.New("failed to get ethereum account: %v", err)
	}

	opts, err := bind.NewTransactor(strings.NewReader(account.Key), account.Password)
	if err != nil {
		return nil, errors.New("failed to create new transaction opts: %v", err)
	}

	opts.GasPrice = gc.config.GetEthereumGasPrice()
	opts.GasLimit = gc.config.GetEthereumGasLimit()
	opts.Context = context.Background()
	return opts, nil
}


func (gc *gethClient) queueTaskTransactionStatus(tenantID identity.CentID, txHash string) (txID uuid.UUID, err error) {
	tx, err := gc.txService.CreateTransaction(tenantID, "polling Ethereum transaction status")
	if err != nil {
		return txID, err
	}
	_, err = gc.queue.EnqueueJob(TransactionStatusTaskName, map[string]interface{}{
		transactions.TxIDParam:           tx.ID.String(),
		TransactionAccountParam: tenantID.String(),
		TransactionTxHashParam:  txHash,
	})

	return tx.ID, err
}

// SubmitTransaction creates an Ethereum transactions with retries
func (gc *gethClient) SubmitTransaction(context context.Context, tendantID identity.CentID, contractMethod interface{}, opts *bind.TransactOpts, params ...interface{}) (*uuid.UUID,*types.Transaction, error) {

	tx, err := gc.SubmitTransactionWithRetries(contractMethod,opts, params)
	if err != nil {
		return nil,nil, errors.New("Submit Ethereum transaction failed", err)
	}

	txHash := tx.Hash()
	txID, err := gc.queueTaskTransactionStatus(tendantID, txHash.String())

	if err != nil {
		return nil,nil, errors.New("Failed to generated a queue task to poll the Ethereum transaction status", err)
	}
	return &txID, tx, nil

}


/**
SubmitTransactionWithRetries submits transaction to the ethereum chain
Blocking Function that sends transaction using reflection wrapped in a retrial block. It is based on the transactionUnderpriced error,
meaning that a transaction is being attempted to run twice, and the logic is to override the existing one. As we have constant
gas prices that means that a concurrent transaction race condition event has happened.
- contractMethod: Contract Method that implements GenericEthereumAsset (usually autogenerated binding from abi)
- params: Arbitrary number of parameters that are passed to the function fname call
Note: contractMethod must always return "*types.Transaction, error"
*/
func (gc *gethClient) SubmitTransactionWithRetries(contractMethod interface{}, opts *bind.TransactOpts, params ...interface{}) (*types.Transaction, error) {
	var current int
	f := reflect.ValueOf(contractMethod)
	maxTries := gc.config.GetEthereumMaxRetries()

	gc.txMu.Lock()
	defer gc.txMu.Unlock()

	var err error
	for {

		if current >= maxTries {
			return nil, errors.New("max concurrent transaction tries reached: %v", err)
		}

		current++
		err = gc.incrementNonce(opts, gc.config.GetTxPoolAccessEnabled(), gc.client, gc.rpcClient)
		if err != nil {
			return nil, errors.New("failed to increment nonce: %v", err)
		}

		if opts.Nonce != nil {
			log.Infof("Incrementing Nonce to [%v]\n", opts.Nonce.String())
		}

		var in []reflect.Value
		in = append(in, reflect.ValueOf(opts))
		for _, p := range params {
			in = append(in, reflect.ValueOf(p))
		}

		result := f.Call(in)
		var tx *types.Transaction
		if !result[0].IsNil() {
			tx = result[0].Interface().(*types.Transaction)
		}

		if !result[1].IsNil() {
			err = result[1].Interface().(error)
		}

		if err == nil {
			return tx, nil
		}

		if (err.Error() == transactionUnderpriced.Error()) || (err.Error() == nonceTooLow.Error()) {
			log.Warningf("Concurrent transaction identified, trying again [%d/%d]\n", current, maxTries)
			time.Sleep(gc.config.GetEthereumIntervalRetry())
			continue
		}

		return nil, err
	}

}

// GetGethCallOpts returns the Call options with default
func (gc *gethClient) GetGethCallOpts() (*bind.CallOpts, context.CancelFunc) {
	// Assuring that pending transactions are taken into account by go-ethereum when asking for things like
	// specific transactions and client's nonce
	// with timeout context, in case eth node is not in sync
	ctx, cancel := gc.defaultReadContext()
	return &bind.CallOpts{Pending: true, Context: ctx}, cancel
}

// noncer defines functions to get the next nonce
type noncer interface {
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
}

// callContexter defines functions to get CallContext
type callContexter interface {
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
}

// incrementNonce updates the opts.Nonce to next valid nonce
// We pick the current nonce by getting latest transactions included in the blocks including pending blocks
// then we check the txpool to see if there any new transactions from the address that are not included in any block
// If there are no pending transactions in txpool, we use the current nonce + 1
// else we increment the greater of current nonce or the nonce derived from txpool
func (gc *gethClient) incrementNonce(opts *bind.TransactOpts, txpoolAccessEnabled bool, noncer noncer, cc callContexter) error {
	// check if the txpool access is enabled
	if !txpoolAccessEnabled {
		log.Warningf("Ethereum Client doesn't support txpool API, may cause transactions failures")
		opts.Nonce = nil
		return nil
	}

	ctx, cancel := gc.defaultReadContext()
	defer cancel()

	// get current nonce
	n, err := noncer.PendingNonceAt(ctx, opts.From)
	if err != nil {
		return errors.New("failed to get chain nonce for %s: %v", opts.From.String(), err)
	}

	// set the nonce
	opts.Nonce = new(big.Int).SetUint64(n)

	// check for any transactions in txpool
	res := make(map[string]map[string]map[string]string)
	err = cc.CallContext(ctx, &res, "txpool_inspect")
	if err != nil {
		return errors.New("failed to get txpool data: %v", err)
	}

	// no pending transaction from this account in tx pool
	if len(res["pending"][opts.From.Hex()]) < 1 {
		return nil
	}

	var keys []int
	for k := range res["pending"][opts.From.Hex()] {
		ki, err := strconv.Atoi(k)
		if err != nil {
			return errors.New("failed to convert nonce: %v", err)
		}

		keys = append(keys, ki)
	}

	// there are some pending transactions in txpool, check their nonce
	// pick the largest one and increment it
	sort.Ints(keys)
	lastPoolNonce := keys[len(keys)-1]
	if uint64(lastPoolNonce) >= n {
		opts.Nonce = new(big.Int).Add(big.NewInt(int64(lastPoolNonce)), big.NewInt(1))
	}

	return nil
}
