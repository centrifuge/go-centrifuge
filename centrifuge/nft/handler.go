package nft

import (
	"context"

	"github.com/centrifuge/go-centrifuge/centrifuge/centerrors"
	"github.com/centrifuge/go-centrifuge/centrifuge/code"
	"github.com/centrifuge/go-centrifuge/centrifuge/protobufs/gen/go/nft"
	"github.com/ethereum/go-ethereum/common/hexutil"
	logging "github.com/ipfs/go-log"
)

var apiLog = logging.Logger("nft-api")

type PaymentObligationService interface {
	mintNFT(documentID []byte, docType, registryAddress, depositAddress string, proofFields []string) (string, error)
}

type grpcHandler struct {
	service PaymentObligationService
}

// GRPCHandler returns an implementation of invoice.DocumentServiceServer
func GRPCHandler() nftpb.NFTServiceServer {
	return &grpcHandler{service: getPaymentObligationService()}
}

// MintNFT will be called from the client API to mint an NFT
func (g grpcHandler) MintNFT(context context.Context, request *nftpb.NFTMintRequest) (*nftpb.NFTMintResponse, error) {
	apiLog.Infof("Received request to Mint an NFT", request)
	identifier, err := hexutil.Decode(request.Identifier)
	if err != nil {
		return &nftpb.NFTMintResponse{}, centerrors.New(code.Unknown, err.Error())
	}

	tokenID, err := g.service.mintNFT(identifier, request.Type, request.RegistryAddress, request.DepositAddress, request.ProofFields)
	if err != nil {
		return &nftpb.NFTMintResponse{}, centerrors.New(code.Unknown, err.Error())
	}
	return &nftpb.NFTMintResponse{TokenId: tokenID}, nil
}
