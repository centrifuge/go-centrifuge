// +build unit

package coreapi

import (
	"math/big"
	"strings"
	"testing"

	coredocumentpb "github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/jobs"
	documentpb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/document"
	testingdocuments "github.com/centrifuge/go-centrifuge/testingutils/documents"
	testingidentity "github.com/centrifuge/go-centrifuge/testingutils/identity"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTypes_convertAttributes(t *testing.T) {
	attrs := AttributeMap{
		"string_test": {
			Type:  "string",
			Value: "hello, world!",
		},

		"decimal_test": {
			Type:  "decimal",
			Value: "100001.001",
		},
	}

	atts, err := toDocumentAttributes(attrs)
	assert.NoError(t, err)
	assert.Len(t, atts, 2)

	var attrList []documents.Attribute
	for _, v := range atts {
		attrList = append(attrList, v)
	}
	cattrs, err := convertAttributes(attrList)
	assert.NoError(t, err)
	assert.Equal(t, attrs, cattrs)

	attrs["invalid"] = Attribute{Type: "unknown", Value: "some value"}
	_, err = toDocumentAttributes(attrs)
	assert.Error(t, err)

	attrList = append(attrList, documents.Attribute{Value: documents.AttrVal{Type: "invalid"}})
	_, err = convertAttributes(attrList)
	assert.Error(t, err)
}

func TestTypes_deriveResponseHeader(t *testing.T) {
	model := new(testingdocuments.MockModel)
	model.On("GetCollaborators", mock.Anything).Return(documents.CollaboratorsAccess{}, errors.New("error fetching collaborators")).Once()
	_, err := deriveResponseHeader(nil, model, jobs.NewJobID())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error fetching collaborators")
	model.AssertExpectations(t)

	id := utils.RandomSlice(32)
	did1 := testingidentity.GenerateRandomDID()
	did2 := testingidentity.GenerateRandomDID()
	ca := documents.CollaboratorsAccess{
		ReadCollaborators:      []identity.DID{did1},
		ReadWriteCollaborators: []identity.DID{did2},
	}
	model = new(testingdocuments.MockModel)
	model.On("GetCollaborators", mock.Anything).Return(ca, nil).Once()
	model.On("ID").Return(id).Once()
	model.On("CurrentVersion").Return(id).Once()
	model.On("Author").Return(nil, errors.New("somerror"))
	model.On("Timestamp").Return(nil, errors.New("somerror"))
	model.On("NFTs").Return(nil)
	resp, err := deriveResponseHeader(nil, model, jobs.NewJobID())
	assert.NoError(t, err)
	assert.Equal(t, hexutil.Encode(id), resp.DocumentID)
	assert.Equal(t, hexutil.Encode(id), resp.VersionID)
	assert.Len(t, resp.ReadAccess, 1)
	assert.Equal(t, resp.ReadAccess[0].String(), did1.String())
	assert.Len(t, resp.WriteAccess, 1)
	assert.Equal(t, resp.WriteAccess[0].String(), did2.String())
	model.AssertExpectations(t)
}

func invoiceData() map[string]interface{} {
	return map[string]interface{}{
		"number":       "12345",
		"status":       "unpaid",
		"gross_amount": "12.345",
		"recipient":    "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7",
		"date_due":     "2019-05-24T14:48:44.308854Z", // rfc3339nano
		"date_paid":    "2019-05-24T14:48:44Z",        // rfc3339
		"currency":     "EUR",
		"attachments": []map[string]interface{}{
			{
				"name":      "test",
				"file_type": "pdf",
				"size":      1000202,
				"data":      "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7",
				"checksum":  "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF3",
			},
		},
	}
}

func TestTypes_toDocumentCreatePayload(t *testing.T) {
	request := CreateDocumentRequest{Scheme: "invoice"}
	request.Data = invoiceData()

	// success
	payload, err := toDocumentsCreatePayload(request)
	assert.NoError(t, err)
	assert.Equal(t, payload.Scheme, "invoice")
	assert.NotNil(t, payload.Data)

	// failure
	request.Attributes = map[string]Attribute{
		"invalid": {Type: "unknown", Value: "some value"},
	}

	_, err = toDocumentsCreatePayload(request)
	assert.Error(t, err)
}

func TestTypes_convertNFTs(t *testing.T) {
	regIDs := [][]byte{
		utils.RandomSlice(32),
		utils.RandomSlice(32),
	}
	tokIDs := [][]byte{
		utils.RandomSlice(32),
		utils.RandomSlice(32),
	}
	tokIDx := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
	}
	addrs := []common.Address{
		common.BytesToAddress(utils.RandomSlice(20)),
		common.BytesToAddress(utils.RandomSlice(20)),
	}
	tests := []struct {
		name         string
		TR           func() documents.TokenRegistry
		NFTs         func() []*coredocumentpb.NFT
		isErr        bool
		errLen       int
		errMsg       string
		nftLen       int
		expectedNFTs []*documentpb.NFT
	}{
		{
			name: "1 nft, no error",
			TR: func() documents.TokenRegistry {
				m := new(testingdocuments.MockRegistry)
				m.On("OwnerOf", mock.Anything, mock.Anything).Return(addrs[0], nil).Once()
				m.On("CurrentIndexOfToken", mock.Anything, mock.Anything).Return(tokIDx[0], nil).Once()
				return m
			},
			NFTs: func() []*coredocumentpb.NFT {
				return []*coredocumentpb.NFT{
					{
						RegistryId: regIDs[0],
						TokenId:    tokIDs[0],
					},
				}
			},
			isErr:  false,
			nftLen: 1,
			expectedNFTs: []*documentpb.NFT{
				{
					Registry:   hexutil.Encode(regIDs[0][:20]),
					Owner:      addrs[0].Hex(),
					TokenId:    hexutil.Encode(tokIDs[0]),
					TokenIndex: hexutil.Encode(tokIDx[0].Bytes()),
				},
			},
		},
		{
			name: "2 nft, no error",
			TR: func() documents.TokenRegistry {
				m := new(testingdocuments.MockRegistry)
				m.On("OwnerOf", mock.Anything, mock.Anything).Return(addrs[0], nil).Once()
				m.On("OwnerOf", mock.Anything, mock.Anything).Return(addrs[1], nil).Once()
				m.On("CurrentIndexOfToken", mock.Anything, mock.Anything).Return(tokIDx[0], nil).Once()
				m.On("CurrentIndexOfToken", mock.Anything, mock.Anything).Return(tokIDx[1], nil).Once()
				return m
			},
			NFTs: func() []*coredocumentpb.NFT {
				return []*coredocumentpb.NFT{
					{
						RegistryId: regIDs[0],
						TokenId:    tokIDs[0],
					},
					{
						RegistryId: regIDs[1],
						TokenId:    tokIDs[1],
					},
				}
			},
			isErr:  false,
			nftLen: 2,
			expectedNFTs: []*documentpb.NFT{
				{
					Registry:   hexutil.Encode(regIDs[0][:20]),
					Owner:      addrs[0].Hex(),
					TokenId:    hexutil.Encode(tokIDs[0]),
					TokenIndex: hexutil.Encode(tokIDx[0].Bytes()),
				},
				{
					Registry:   hexutil.Encode(regIDs[1][:20]),
					Owner:      addrs[1].Hex(),
					TokenId:    hexutil.Encode(tokIDs[1]),
					TokenIndex: hexutil.Encode(tokIDx[1].Bytes()),
				},
			},
		},
		{
			name: "2 nft, ownerOf error",
			TR: func() documents.TokenRegistry {
				m := new(testingdocuments.MockRegistry)
				m.On("OwnerOf", mock.Anything, mock.Anything).Return(addrs[0], errors.New("owner error")).Once()
				m.On("OwnerOf", mock.Anything, mock.Anything).Return(addrs[1], nil).Once()
				m.On("CurrentIndexOfToken", mock.Anything, mock.Anything).Return(tokIDx[0], nil).Once()
				m.On("CurrentIndexOfToken", mock.Anything, mock.Anything).Return(tokIDx[1], nil).Once()
				return m
			},
			NFTs: func() []*coredocumentpb.NFT {
				return []*coredocumentpb.NFT{
					{
						RegistryId: regIDs[0],
						TokenId:    tokIDs[0],
					},
					{
						RegistryId: regIDs[1],
						TokenId:    tokIDs[1],
					},
				}
			},
			isErr:  true,
			errLen: 1,
			errMsg: "owner",
			nftLen: 1,
			expectedNFTs: []*documentpb.NFT{
				{
					Registry:   hexutil.Encode(regIDs[1][:20]),
					Owner:      addrs[1].Hex(),
					TokenId:    hexutil.Encode(tokIDs[1]),
					TokenIndex: hexutil.Encode(tokIDx[1].Bytes()),
				},
			},
		},
		{
			name: "2 nft, CurrentIndexOfToken error",
			TR: func() documents.TokenRegistry {
				m := new(testingdocuments.MockRegistry)
				m.On("OwnerOf", mock.Anything, mock.Anything).Return(addrs[1], nil).Once()
				m.On("CurrentIndexOfToken", mock.Anything, mock.Anything).Return(tokIDx[0], errors.New("CurrentIndexOfToken error")).Once()
				m.On("CurrentIndexOfToken", mock.Anything, mock.Anything).Return(tokIDx[1], nil).Once()
				return m
			},
			NFTs: func() []*coredocumentpb.NFT {
				return []*coredocumentpb.NFT{
					{
						RegistryId: regIDs[0],
						TokenId:    tokIDs[0],
					},
					{
						RegistryId: regIDs[1],
						TokenId:    tokIDs[1],
					},
				}
			},
			isErr:  true,
			errLen: 1,
			errMsg: "CurrentIndexOfToken",
			nftLen: 1,
			expectedNFTs: []*documentpb.NFT{
				{
					Registry:   hexutil.Encode(regIDs[1][:20]),
					Owner:      addrs[1].Hex(),
					TokenId:    hexutil.Encode(tokIDs[1]),
					TokenIndex: hexutil.Encode(tokIDx[1].Bytes()),
				},
			},
		},
		{
			name: "2 nft, 2 CurrentIndexOfToken error",
			TR: func() documents.TokenRegistry {
				m := new(testingdocuments.MockRegistry)
				m.On("CurrentIndexOfToken", mock.Anything, mock.Anything).Return(tokIDx[0], errors.New("CurrentIndexOfToken error")).Once()
				m.On("CurrentIndexOfToken", mock.Anything, mock.Anything).Return(tokIDx[1], errors.New("CurrentIndexOfToken error")).Once()
				return m
			},
			NFTs: func() []*coredocumentpb.NFT {
				return []*coredocumentpb.NFT{
					{
						RegistryId: regIDs[0],
						TokenId:    tokIDs[0],
					},
					{
						RegistryId: regIDs[1],
						TokenId:    tokIDs[1],
					},
				}
			},
			isErr:  true,
			errLen: 2,
			errMsg: "CurrentIndexOfToken",
			nftLen: 0,
		},
		{
			name: "2 nft, ownerOf and CurrentIndexOfToken error",
			TR: func() documents.TokenRegistry {
				m := new(testingdocuments.MockRegistry)
				m.On("OwnerOf", mock.Anything, mock.Anything).Return(addrs[0], errors.New("owner error")).Once()
				m.On("OwnerOf", mock.Anything, mock.Anything).Return(addrs[1], nil).Once()
				m.On("CurrentIndexOfToken", mock.Anything, mock.Anything).Return(tokIDx[0], nil).Once()
				m.On("CurrentIndexOfToken", mock.Anything, mock.Anything).Return(tokIDx[1], errors.New("CurrentIndexOfToken error")).Once()
				return m
			},
			NFTs: func() []*coredocumentpb.NFT {
				return []*coredocumentpb.NFT{
					{
						RegistryId: regIDs[0],
						TokenId:    tokIDs[0],
					},
					{
						RegistryId: regIDs[1],
						TokenId:    tokIDs[1],
					},
				}
			},
			isErr:  true,
			errLen: 2,
			errMsg: "owner",
			nftLen: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n, err := convertNFTs(test.TR(), test.NFTs())
			if test.isErr {
				assert.Error(t, err)
				assert.Equal(t, errors.Len(err), test.errLen)
				assert.Contains(t, err.Error(), test.errMsg)
			} else {
				assert.NoError(t, err)
			}
			assert.Len(t, n, test.nftLen)
			if test.nftLen > 0 {
				for i, nn := range n {
					assert.Equal(t, strings.ToLower(nn.Registry), strings.ToLower(test.expectedNFTs[i].Registry))
					assert.Equal(t, strings.ToLower(nn.TokenIndex), strings.ToLower(test.expectedNFTs[i].TokenIndex))
					assert.Equal(t, strings.ToLower(nn.TokenID), strings.ToLower(test.expectedNFTs[i].TokenId))
					assert.Equal(t, strings.ToLower(nn.Owner), strings.ToLower(test.expectedNFTs[i].Owner))
				}
			}
		})
	}
}
