package p2pcontroller

import (
	"context"
	"github.com/CentrifugeInc/centrifuge-protobufs/gen/go/p2p"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/p2p/grpcservice"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/notification"
)

type P2PService struct{}

func (srv *P2PService) Post(ctx context.Context, req *p2ppb.P2PMessage) (rep *p2ppb.P2PReply, err error) {
	// We can define at runtime which notification backend to use pub/sub, queue, webhook ...
	var svc = grpcservice.P2PService{&notification.WebhookSender{}}
	return svc.HandleP2PPost(ctx, req)
}
