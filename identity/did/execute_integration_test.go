// +build integration

/*

The identity contract has a method called execute which forwards
a call to another contract. In Ethereum it is a useful pattern especially
related to identity smart contracts.


The correct behaviour currently is tested by calling the anchorRepository
commit method. It could be any other contract or method as well to simple test the
correct implementation.

*/

package did

import (
	"testing"

	"github.com/centrifuge/go-centrifuge/anchors"
	"github.com/centrifuge/go-centrifuge/ethereum"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestExecute_successful(t *testing.T) {
	did := deployIdentityContract(t)
	aCtx := getTestDIDContext(t, *did)
	idSrv := initIdentity(cfg)
	anchorAddress := getAnchorAddress()

	// init params
	testAnchorId, _ := anchors.ToAnchorID(utils.RandomSlice(32))
	rootHash := utils.RandomSlice(32)
	testRootHash, _ := anchors.ToDocumentRoot(rootHash)
	proofs := [][anchors.DocumentProofLength]byte{utils.RandomByte32()}

	err := idSrv.Execute(aCtx, anchorAddress, anchors.AnchorContractABI, "commit", testAnchorId.BigInt(), testRootHash, proofs)
	assert.Nil(t, err, "Execute method calls should be successful")

	checkAnchor(t, testAnchorId, rootHash)

}

func TestExecute_fail_falseMethodName(t *testing.T) {
	did := deployIdentityContract(t)
	aCtx := getTestDIDContext(t, *did)

	idSrv := initIdentity(cfg)
	anchorAddress := getAnchorAddress()

	testAnchorId, _ := anchors.ToAnchorID(utils.RandomSlice(32))
	rootHash := utils.RandomSlice(32)
	testRootHash, _ := anchors.ToDocumentRoot(rootHash)

	proofs := [][anchors.DocumentProofLength]byte{utils.RandomByte32()}

	err := idSrv.Execute(aCtx, anchorAddress, anchors.AnchorContractABI, "fakeMethod", testAnchorId.BigInt(), testRootHash, proofs)
	assert.Error(t, err, "should throw an error because method is not existing in abi")

}

func TestExecute_fail_MissingParam(t *testing.T) {
	did := deployIdentityContract(t)
	aCtx := getTestDIDContext(t, *did)
	idSrv := initIdentity(cfg)
	anchorAddress := getAnchorAddress()

	testAnchorId, _ := anchors.ToAnchorID(utils.RandomSlice(32))
	rootHash := utils.RandomSlice(32)
	testRootHash, _ := anchors.ToDocumentRoot(rootHash)

	err := idSrv.Execute(aCtx, anchorAddress, anchors.AnchorContractABI, "commit", testAnchorId.BigInt(), testRootHash)
	assert.Error(t, err, "should throw an error because method is not existing in abi")
}

func checkAnchor(t *testing.T, anchorId anchors.AnchorID, expectedRootHash []byte) {
	anchorAddress := getAnchorAddress()
	client := ctx[ethereum.BootstrappedEthereumClient].(ethereum.Client)

	// check if anchor has been added
	opts, _ := client.GetGethCallOpts(false)
	anchorContract := bindAnchorContract(t, anchorAddress)
	result, err := anchorContract.GetAnchorById(opts, anchorId.BigInt())
	assert.Nil(t, err, "get anchor should be successful")
	assert.Equal(t, expectedRootHash[:], result.DocumentRoot[:], "committed anchor should be the same")
}

// Checks the standard behaviour of the anchor contract
func TestAnchorWithoutExecute_successful(t *testing.T) {
	client := ctx[ethereum.BootstrappedEthereumClient].(ethereum.Client)
	anchorAddress := getAnchorAddress()
	anchorContract := bindAnchorContract(t, anchorAddress)

	testAnchorId, _ := anchors.ToAnchorID(utils.RandomSlice(32))
	rootHash := utils.RandomSlice(32)
	testRootHash, _ := anchors.ToDocumentRoot(rootHash)

	//commit without execute method
	txStatus := commitAnchorWithoutExecute(t, anchorContract, testAnchorId, testRootHash)
	assert.Equal(t, ethereum.TransactionStatusSuccess, txStatus.Status, "transaction should be successful")

	opts, _ := client.GetGethCallOpts(false)
	result, err := anchorContract.GetAnchorById(opts, testAnchorId.BigInt())
	assert.Nil(t, err, "get anchor should be successful")
	assert.Equal(t, rootHash[:], result.DocumentRoot[:], "committed anchor should be the same")

}

// Checks the standard behaviour of the anchor contract
func commitAnchorWithoutExecute(t *testing.T, anchorContract *anchors.AnchorContract, anchorId anchors.AnchorID, rootHash anchors.DocumentRoot) *ethereum.WatchTransaction {
	client := ctx[ethereum.BootstrappedEthereumClient].(ethereum.Client)

	proofs := [][anchors.DocumentProofLength]byte{utils.RandomByte32()}

	opts, err := client.GetTxOpts(cfg.GetEthereumDefaultAccountName())
	tx, err := client.SubmitTransactionWithRetries(anchorContract.Commit, opts, anchorId.BigInt(), rootHash, proofs)

	assert.Nil(t, err, "submit transaction should be successful")

	watchTrans := make(chan *ethereum.WatchTransaction)
	go waitForTransaction(client, tx.Hash(), watchTrans)
	return <-watchTrans
}

func bindAnchorContract(t *testing.T, address common.Address) *anchors.AnchorContract {
	client := ctx[ethereum.BootstrappedEthereumClient].(ethereum.Client)
	anchorContract, err := anchors.NewAnchorContract(address, client.GetEthClient())
	assert.Nil(t, err, "bind anchor contract should not throw an error")

	return anchorContract
}
