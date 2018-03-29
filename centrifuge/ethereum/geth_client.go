package ethereum

import (
	"log"
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	"github.com/go-errors/errors"
	"math/big"
	"time"
	"context"
	"sync"
	"net/url"
	"github.com/ethereum/go-ethereum/core/types"
	"reflect"
)

const (
	mainAccountName = "main"
	defaultMaxRetries = 200
	defaultIntervalRetry = (100 * time.Millisecond)
	TransactionUnderpriced = "replacement transaction underpriced"
)

var gc *GethClient
var gcInit sync.Once

// getDefaultContextTimeout retrieves the default duration before an Ethereum call context should time out
func getDefaultContextTimeout() (time.Duration) {
	return viper.GetDuration("ethereum.contextWaitTimeout")
}

func getMaxRetries() (x int) {
	x = viper.GetInt("ethereum.maxRetries")
	if x == 0 {
		x = defaultMaxRetries
	}
	return
}

func getIntervalRetry() (x time.Duration) {
	x = viper.GetDuration("ethereum.intervalRetry")
	if x == (0 * time.Second) {
		x = defaultIntervalRetry
	}
	return
}

func DefaultWaitForTransactionMiningContext() (ctx context.Context) {
	toBeDone := time.Now().Add(getDefaultContextTimeout())
	ctx, _ = context.WithDeadline(context.TODO(), toBeDone)
	return
}

// Abstract the "ethereum client" out so we can more easily support other clients
// besides Geth (e.g. quorum)
// Also make it easier to mock tests
type EthereumClient interface {
	GetClient() (*ethclient.Client)
}

type GethClient struct {
	Client *ethclient.Client
	Host *url.URL
}

func (gethClient GethClient) GetClient() (*ethclient.Client) {
	return gethClient.Client
}

// GetConnection returns the connection to the configured `ethereum.gethSocket`.
// Note that this is a singleton and is the same connection for the whole application.
func GetConnection() (EthereumClient) {
	gcInit.Do(func() {
		u, err := url.Parse(viper.GetString("ethereum.gethSocket"))
		if err != nil {
			log.Fatalf("Failed to connect to parse ethereum.gethSocket URL: %v", err)
		}
		client, err := ethclient.Dial(u.String())
		if err != nil {
			log.Fatalf("Failed to connect to the Ethereum client [%s]: %v", u.String(), err)
		} else {
			gc = &GethClient{client, u}
		}
	})
	return gc
}

// GetGethTxOpts retrieves the geth transaction options for the given account name. The account name influences which configuration
// is used. If no account name is provided the account as defined by `mainAccountName` constant is used
// It is not supported to call with more than one account name.
func GetGethTxOpts(optionalAccountName ...string) (*bind.TransactOpts, error) {
	var accountName string
	accsLen := len(optionalAccountName)
	if accsLen > 1 {
		err := errors.Errorf("error in use of method. can deal with maximum of one account name for ethereum transaction options. please check your code.")
		log.Fatalf(err.Error())
		return nil, err
	} else {
		switch accsLen {
		case 1:
			accountName = optionalAccountName[0]
		default:
			accountName = mainAccountName
		}
	}

	key := viper.GetString("ethereum.accounts." + accountName + ".key")

	// TODO: this could be done more elegantly if support for additional ways to configure keys should be added later on
	// e.g. if key files would be supported instead of inline keys
	if key == "" {
		err := errors.Errorf("could not find configured ethereum key for account [%v]. please check your configuration.\n", accountName)
		log.Printf(err.Error())
		return nil, err
	}

	password := viper.GetString("ethereum.accounts." + accountName + ".password")

	authedTransactionOpts, err := bind.NewTransactor(strings.NewReader(key), password)
	if err != nil {
		err = errors.Errorf("Failed to load key with error: %v", err);
		log.Println(err.Error())
		return nil, err
	} else {
		authedTransactionOpts.GasPrice = big.NewInt(viper.GetInt64("ethereum.gasPrice"))
		authedTransactionOpts.GasLimit = uint64(viper.GetInt64("ethereum.gasLimit"))
		return authedTransactionOpts, nil
	}
}

/**
 Function that sends transaction using reflection wrapped in a retrial block. It is based on the TransactionUnderpriced error,
 meaning that a transaction is being attempted to run twice, and the logic is to override the existing one. As we have constant
 gas prices that means that a concurrent transaction race condition event has happened.
 - contractMethod: Contract Method that implements GenericEthereumAsset (usually autogenerated binding from abi)
 - params: Arbitrary number of parameters that are passed to the function fname call
 */
func SubmitTransactionWithRetries(contractMethod interface{}, params ... interface{}) (tx *types.Transaction, err error) {
	done := false
	maxTries := getMaxRetries()
	current := 0
	var f reflect.Value
	var in []reflect.Value
	var result []reflect.Value
	for !done {
		if current >= maxTries {
			log.Println("Max Concurrent transaction tries reached")
			break
		}
		current += 1
		f = reflect.ValueOf(contractMethod)
		in = make([]reflect.Value, len(params))
		for k, param := range params {
			in[k] = reflect.ValueOf(param)
		}
		result = f.Call(in)
		tx = result[0].Interface().(*types.Transaction)
		err = nil
		if result[1].Interface() != nil {
			err = result[1].Interface().(error)
		}

		if err != nil {
			if err.Error() == TransactionUnderpriced {
				log.Printf("Concurrent transaction identified, trying again [%d/%d]\n", current, maxTries )
				time.Sleep(getIntervalRetry())
			} else {
				done = true
			}
		} else {
			done = true
		}
	}

	return
}

func GetGethCallOpts() (auth *bind.CallOpts) {
	// Assuring that pending transactions are taken into account by go-ethereum when asking for things like
	// specific transactions and client's nonce
	return &bind.CallOpts{Pending: true}
}
