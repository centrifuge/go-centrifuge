package ethereum

import (
	"context"
	"math/big"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/jobs"
	"github.com/centrifuge/go-centrifuge/queue"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("geth-client")
var gc Client
var gcMu sync.RWMutex

// Config defines functions to get ethereum details
type Config interface {
	GetEthereumMaxGasPrice() *big.Int
	GetEthereumNodeURL() string
	GetEthereumAccount(accountName string) (account *config.AccountConfig, err error)
	GetEthereumIntervalRetry() time.Duration
	GetEthereumMaxRetries() int
	GetEthereumContextReadWaitTimeout() time.Duration
}

// DefaultWaitForTransactionMiningContext returns context with timeout for write operations
func DefaultWaitForTransactionMiningContext(d time.Duration) (ctx context.Context, cancelFunc context.CancelFunc) {
	toBeDone := time.Now().Add(d)
	return context.WithDeadline(context.Background(), toBeDone)
}

// EthClient abstracts the implementation of eth client
type EthClient interface {
	ethereum.ChainReader
	ethereum.ChainSyncReader
	ethereum.TransactionReader
	bind.ContractBackend

	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
}

// Client can be implemented by any chain client
type Client interface {

	// GetEthClient returns the ethereum client
	GetEthClient() EthClient

	// GetNodeURL returns the node url
	GetNodeURL() *url.URL

	// GetBlockByNumber returns the block by number
	GetBlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)

	// GetTxOpts returns a cached options if available else creates and returns new options
	GetTxOpts(ctx context.Context, accountName string) (*bind.TransactOpts, error)

	// SubmitTransactionWithRetries submits transaction to the ethereum chain
	// Blocking Function that sends transaction using reflection wrapped in a retrial block. It is based on the ErrTransactionUnderpriced error,
	// meaning that a transaction is being attempted to run twice, and the logic is to override the existing one. As we have constant
	// gas prices that means that a concurrent transaction race condition event has happened.
	// - contractMethod: Contract Method that implements GenericEthereumAsset (usually autogenerated binding from abi)
	// - params: Arbitrary number of parameters that are passed to the function fname call
	// Note: contractMethod must always return "*types.Job, error"
	SubmitTransactionWithRetries(contractMethod interface{}, opts *bind.TransactOpts, params ...interface{}) (tx *types.Transaction, err error)

	// GetGethCallOpts returns the Call options with default
	GetGethCallOpts(pending bool) (*bind.CallOpts, context.CancelFunc)

	// TransactionByHash returns a Ethereum transaction
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)

	// TransactionReceipt return receipt of a transaction
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

// gethClient implements Client for Ethereum
type gethClient struct {
	client    EthClient
	rpcClient *rpc.Client
	host      *url.URL
	accounts  map[string]*bind.TransactOpts
	accMu     sync.Mutex // accMu to protect accounts
	config    Config

	// txMu to ensure one transaction at a time per client
	txMu sync.Mutex
}

// NewGethClient returns an gethClient which implements Client
func NewGethClient(config Config) (Client, error) {
	// This might be removed as soon as we support multiple ethereum keys per account, the error might not be thrown at startup
	acc, err := config.GetEthereumAccount("main")
	if err != nil {
		return nil, err
	}
	if acc.Key == "" {
		return nil, ErrEthKeyNotProvided
	}
	if acc.Password == "" {
		log.Warningf("Main Ethereum Password not provided")
	}

	log.Info("Opening connection to Ethereum:", config.GetEthereumNodeURL())
	u, err := url.Parse(config.GetEthereumNodeURL())
	if err != nil {
		return nil, errors.NewTypedError(ErrEthURL, err)
	}

	c, err := rpc.Dial(u.String())
	if err != nil {
		return nil, errors.NewTypedError(ErrEthURL, err)
	}

	return &gethClient{
		client:    ethclient.NewClient(c),
		rpcClient: c,
		host:      u,
		accounts:  make(map[string]*bind.TransactOpts),
		txMu:      sync.Mutex{},
		config:    config,
		accMu:     sync.Mutex{},
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
// TODO change upstream context in NFT or tx manager instead of passing background context internally
func (gc *gethClient) GetTxOpts(ctx context.Context, accountName string) (opts *bind.TransactOpts, err error) {
	gc.accMu.Lock()
	defer gc.accMu.Unlock()

	if opts, ok := gc.accounts[accountName]; ok {
		return gc.copyOpts(context.Background(), opts)
	}

	txOpts, err := gc.getGethTxOpts(accountName)
	if err != nil {
		return nil, err
	}

	gc.accounts[accountName] = txOpts
	return gc.copyOpts(context.Background(), txOpts)
}

// copyOpts copies tx opts each time to avoid any race conditions when modifying, for example gasLimit
func (gc *gethClient) copyOpts(ctx context.Context, original *bind.TransactOpts) (nOpts *bind.TransactOpts, err error) {
	nOpts = &bind.TransactOpts{
		From:   original.From,
		Signer: original.Signer,
	}
	nOpts.GasPrice, err = gc.getOptimalGasPrice(ctx) // use oracle
	if err != nil {
		return nil, errors.NewTypedError(ErrEthTransaction, errors.New("failed to create new transaction opts: %v", err))
	}
	// Note that gasLimit is must be set at a later point based on the operation performed
	nOpts.Context = ctx
	return nOpts, nil
}

// GetEthClient returns the ethereum client
func (gc *gethClient) GetEthClient() EthClient {
	return gc.client
}

// GetBlockByNumber returns the block by number
func (gc *gethClient) GetBlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	var current int
	var blk *types.Block
	var err error
	maxTries := 10

	for {
		current++
		if current == maxTries {
			return nil, errors.New("Error retrying getting block number %d: %v", number, err)
		}
		blk, err = gc.client.BlockByNumber(ctx, number)
		if err != nil || blk == nil {
			log.Warningf("[%d/%d] Error looking block number[%d][%v]: %v", current, maxTries, number, blk, err)
			time.Sleep(2 * time.Second)
			continue
		}

		break
	}

	return blk, nil
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
		return nil, errors.NewTypedError(ErrEthTransaction, errors.New("failed to get ethereum account: %v", err))
	}

	opts, err := bind.NewTransactor(strings.NewReader(account.Key), account.Password)
	if err != nil {
		return nil, errors.NewTypedError(ErrEthTransaction, errors.New("failed to create new transaction opts: %v", err))
	}
	return opts, nil
}

// getOptimalGasPrice get the optimal current gas price from eth client
func (gc *gethClient) getOptimalGasPrice(ctx context.Context) (*big.Int, error) {
	// we don't acquire the exclusive lock since this method must only be called by other public methods that has the lock
	suggested, err := gc.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	if gc.config.GetEthereumMaxGasPrice().Cmp(suggested) == -1 {
		log.Warningf("suggested gas price %s is greater than max allowed %s", suggested.String(), gc.config.GetEthereumMaxGasPrice().String())
		return gc.config.GetEthereumMaxGasPrice(), nil
	}

	return suggested, nil
}

// QueueEthTXStatusTask starts a new queuing transaction check task.
func QueueEthTXStatusTask(
	accountID identity.DID,
	jobID jobs.JobID,
	txHash common.Hash,
	queuer queue.TaskQueuer) (res queue.TaskResult, err error) {
	return QueueEthTXStatusTaskWithValue(accountID, jobID, txHash, queuer, nil)
}

// QueueEthTXStatusTaskWithValue starts a new queuing transaction check task with a filtered value.
func QueueEthTXStatusTaskWithValue(
	accountID identity.DID,
	jobID jobs.JobID,
	txHash common.Hash,
	queuer queue.TaskQueuer,
	txValue *jobs.JobValue) (res queue.TaskResult, err error) {
	params := map[string]interface{}{
		jobs.JobIDParam:         jobID.String(),
		TransactionAccountParam: accountID.String(),
		TransactionTxHashParam:  txHash.String(),
	}
	if txValue != nil {
		params[TransactionEventName] = txValue.Key
		params[TransactionEventValueIdx] = txValue.KeyIdx
	}

	return queuer.EnqueueJob(EthTXStatusTaskName, params)
}

/**
SubmitTransactionWithRetries submits transaction to the ethereum chain
Blocking Function that sends transaction using reflection wrapped in a retrial block. It is based on the ErrTransactionUnderpriced error,
meaning that a transaction is being attempted to run twice, and the logic is to override the existing one. As we have constant
gas prices that means that a concurrent transaction race condition event has happened.
- contractMethod: Contract Method that implements GenericEthereumAsset (usually autogenerated binding from abi)
- params: Arbitrary number of parameters that are passed to the function fname call
Note: contractMethod must always return "*types.Job, error"
*/
func (gc *gethClient) SubmitTransactionWithRetries(contractMethod interface{}, opts *bind.TransactOpts, params ...interface{}) (*types.Transaction, error) {
	gc.txMu.Lock()
	defer gc.txMu.Unlock()

	var current int
	f := reflect.ValueOf(contractMethod)
	maxTries := gc.config.GetEthereumMaxRetries()

	var err error
	for {

		if current >= maxTries {
			return nil, errors.NewTypedError(ErrEthTransaction, errors.New("max concurrent transaction tries reached: %v", err))
		}

		current++
		err = gc.setNonce(opts)
		if err != nil {
			return nil, errors.NewTypedError(ErrEthTransaction, errors.New("failed to increment nonce: %v", err))
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

		if (err.Error() == ErrTransactionUnderpriced.Error()) || (err.Error() == ErrNonceTooLow.Error()) {
			log.Warningf("Concurrent transaction identified, trying again [%d/%d]\n", current, maxTries)
			time.Sleep(gc.config.GetEthereumIntervalRetry())
			continue
		}

		return nil, err
	}

}

// GetGethCallOpts returns the Call options with default
func (gc *gethClient) GetGethCallOpts(pending bool) (*bind.CallOpts, context.CancelFunc) {
	// Assuring that pending transactions are taken into account by go-ethereum when asking for things like
	// specific transactions and client's nonce
	// with timeout context, in case eth node is not in sync
	ctx, cancel := gc.defaultReadContext()
	return &bind.CallOpts{Pending: pending, Context: ctx}, cancel
}

// setNonce updates the opts.Nonce to next valid nonce
func (gc *gethClient) setNonce(opts *bind.TransactOpts) error {
	ctx, cancel := gc.defaultReadContext()
	defer cancel()

	// get current nonce
	n, err := gc.client.PendingNonceAt(ctx, opts.From)
	if err != nil {
		return errors.NewTypedError(ErrEthTransaction, errors.New("failed to get chain nonce for %s: %v", opts.From.String(), err))
	}

	// set the nonce
	opts.Nonce = new(big.Int).SetUint64(n)
	return nil
}

// BindContract returns a bind contract at the address with corresponding ABI
func BindContract(address common.Address, abi abi.ABI, client Client) *bind.BoundContract {
	return bind.NewBoundContract(address, abi, client.GetEthClient(), client.GetEthClient(), client.GetEthClient())
}
