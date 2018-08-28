// +build unit

package identity

import (
	"testing"

	"context"
	"time"

	"github.com/CentrifugeInc/go-centrifuge/centrifuge/tools"
	"github.com/centrifuge/gocelery"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/stretchr/testify/assert"
)

func TestRegistrationConfirmationTask_ParseKwargsHappyPath(t *testing.T) {
	rct := IdRegistrationConfirmationTask{}
	id := tools.RandomSlice32()
	b32Id := createCentId(id)
	decoded, err := simulateJsonDecode(b32Id)
	err = rct.ParseKwargs(decoded)
	if err != nil {
		t.Errorf("Could not parse %s for [%x]", CentIdParam, id)
	}
	assert.Equal(t, b32Id, rct.CentId, "Resulting id should have the same ID as the input")
}

func TestRegistrationConfirmationTask_ParseKwargsDoesNotExist(t *testing.T) {
	rct := IdRegistrationConfirmationTask{}
	id := tools.RandomSlice32()
	err := rct.ParseKwargs(map[string]interface{}{"notId": id})
	assert.NotNil(t, err, "Should not allow parsing without centId")
}

func TestRegistrationConfirmationTask_ParseKwargsInvalidType(t *testing.T) {
	rct := IdRegistrationConfirmationTask{}
	id := tools.RandomSlice32()
	err := rct.ParseKwargs(map[string]interface{}{CentIdParam: id})
	assert.NotNil(t, err, "Should not parse without the correct type of centId")
}

type MockIdentityCreatedWatcher struct {
	shouldFail bool
	sink       chan<- *EthereumIdentityFactoryContractIdentityCreated
}

func (*MockIdentityCreatedWatcher) WatchIdentityCreated(opts *bind.WatchOpts, sink chan<- *EthereumIdentityFactoryContractIdentityCreated, centrifugeId [][32]byte) (event.Subscription, error) {
	return nil, nil
}

func TestIdRegistrationConfirmationTask_RunTask(t *testing.T) {
	rct := IdRegistrationConfirmationTask{
		CentId: createCentId(tools.RandomSlice32()),
		EthContextInitializer: func() (ctx context.Context, cancelFunc context.CancelFunc) {
			toBeDone := time.Now().Add(time.Duration(1 * time.Millisecond))
			return context.WithDeadline(context.TODO(), toBeDone)
		},
		IdentityCreatedWatcher: &MockIdentityCreatedWatcher{},
	}
	go rct.RunTask()
	time.Sleep(1 * time.Millisecond)
	rct.EthContext.Done()
}

func simulateJsonDecode(b32Id [32]byte) (map[string]interface{}, error) {
	kwargs := map[string]interface{}{CentIdParam: b32Id}
	t1 := gocelery.TaskMessage{Kwargs: kwargs}
	encoded, err := t1.Encode()
	if err != nil {
		return nil, err
	}
	t2, err := gocelery.DecodeTaskMessage(encoded)
	return t2.Kwargs, err
}

func createCentId(id []byte) [32]byte {
	var b32Id [32]byte
	copy(b32Id[:], id[:32])
	return b32Id
}
