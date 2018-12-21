// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/service.proto

package configpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetTenantRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTenantRequest) Reset()         { *m = GetTenantRequest{} }
func (m *GetTenantRequest) String() string { return proto.CompactTextString(m) }
func (*GetTenantRequest) ProtoMessage()    {}
func (*GetTenantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5680e2fa0c56bd70, []int{0}
}
func (m *GetTenantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTenantRequest.Unmarshal(m, b)
}
func (m *GetTenantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTenantRequest.Marshal(b, m, deterministic)
}
func (dst *GetTenantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTenantRequest.Merge(dst, src)
}
func (m *GetTenantRequest) XXX_Size() int {
	return xxx_messageInfo_GetTenantRequest.Size(m)
}
func (m *GetTenantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTenantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTenantRequest proto.InternalMessageInfo

func (m *GetTenantRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type GetAllTenantResponse struct {
	Data                 []*TenantData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetAllTenantResponse) Reset()         { *m = GetAllTenantResponse{} }
func (m *GetAllTenantResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllTenantResponse) ProtoMessage()    {}
func (*GetAllTenantResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5680e2fa0c56bd70, []int{1}
}
func (m *GetAllTenantResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllTenantResponse.Unmarshal(m, b)
}
func (m *GetAllTenantResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllTenantResponse.Marshal(b, m, deterministic)
}
func (dst *GetAllTenantResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllTenantResponse.Merge(dst, src)
}
func (m *GetAllTenantResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllTenantResponse.Size(m)
}
func (m *GetAllTenantResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllTenantResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllTenantResponse proto.InternalMessageInfo

func (m *GetAllTenantResponse) GetData() []*TenantData {
	if m != nil {
		return m.Data
	}
	return nil
}

type UpdateTenantRequest struct {
	Identifier           string      `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Data                 *TenantData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateTenantRequest) Reset()         { *m = UpdateTenantRequest{} }
func (m *UpdateTenantRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTenantRequest) ProtoMessage()    {}
func (*UpdateTenantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5680e2fa0c56bd70, []int{2}
}
func (m *UpdateTenantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTenantRequest.Unmarshal(m, b)
}
func (m *UpdateTenantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTenantRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateTenantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTenantRequest.Merge(dst, src)
}
func (m *UpdateTenantRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTenantRequest.Size(m)
}
func (m *UpdateTenantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTenantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTenantRequest proto.InternalMessageInfo

func (m *UpdateTenantRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *UpdateTenantRequest) GetData() *TenantData {
	if m != nil {
		return m.Data
	}
	return nil
}

type EthereumAccount struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EthereumAccount) Reset()         { *m = EthereumAccount{} }
func (m *EthereumAccount) String() string { return proto.CompactTextString(m) }
func (*EthereumAccount) ProtoMessage()    {}
func (*EthereumAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5680e2fa0c56bd70, []int{3}
}
func (m *EthereumAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EthereumAccount.Unmarshal(m, b)
}
func (m *EthereumAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EthereumAccount.Marshal(b, m, deterministic)
}
func (dst *EthereumAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthereumAccount.Merge(dst, src)
}
func (m *EthereumAccount) XXX_Size() int {
	return xxx_messageInfo_EthereumAccount.Size(m)
}
func (m *EthereumAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_EthereumAccount.DiscardUnknown(m)
}

var xxx_messageInfo_EthereumAccount proto.InternalMessageInfo

func (m *EthereumAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EthereumAccount) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *EthereumAccount) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type KeyPair struct {
	Pub                  string   `protobuf:"bytes,1,opt,name=pub,proto3" json:"pub,omitempty"`
	Pvt                  string   `protobuf:"bytes,2,opt,name=pvt,proto3" json:"pvt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyPair) Reset()         { *m = KeyPair{} }
func (m *KeyPair) String() string { return proto.CompactTextString(m) }
func (*KeyPair) ProtoMessage()    {}
func (*KeyPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5680e2fa0c56bd70, []int{4}
}
func (m *KeyPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyPair.Unmarshal(m, b)
}
func (m *KeyPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyPair.Marshal(b, m, deterministic)
}
func (dst *KeyPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyPair.Merge(dst, src)
}
func (m *KeyPair) XXX_Size() int {
	return xxx_messageInfo_KeyPair.Size(m)
}
func (m *KeyPair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyPair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyPair proto.InternalMessageInfo

func (m *KeyPair) GetPub() string {
	if m != nil {
		return m.Pub
	}
	return ""
}

func (m *KeyPair) GetPvt() string {
	if m != nil {
		return m.Pvt
	}
	return ""
}

type TenantData struct {
	EthAccount                       *EthereumAccount `protobuf:"bytes,1,opt,name=eth_account,json=ethAccount,proto3" json:"eth_account,omitempty"`
	EthDefaultAccountName            string           `protobuf:"bytes,2,opt,name=eth_default_account_name,json=ethDefaultAccountName,proto3" json:"eth_default_account_name,omitempty"`
	ReceiveEventNotificationEndpoint string           `protobuf:"bytes,3,opt,name=receive_event_notification_endpoint,json=receiveEventNotificationEndpoint,proto3" json:"receive_event_notification_endpoint,omitempty"`
	IdentityId                       string           `protobuf:"bytes,4,opt,name=identity_id,json=identityId,proto3" json:"identity_id,omitempty"`
	SigningKeyPair                   *KeyPair         `protobuf:"bytes,5,opt,name=signing_key_pair,json=signingKeyPair,proto3" json:"signing_key_pair,omitempty"`
	EthauthKeyPair                   *KeyPair         `protobuf:"bytes,6,opt,name=ethauth_key_pair,json=ethauthKeyPair,proto3" json:"ethauth_key_pair,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}         `json:"-"`
	XXX_unrecognized                 []byte           `json:"-"`
	XXX_sizecache                    int32            `json:"-"`
}

func (m *TenantData) Reset()         { *m = TenantData{} }
func (m *TenantData) String() string { return proto.CompactTextString(m) }
func (*TenantData) ProtoMessage()    {}
func (*TenantData) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5680e2fa0c56bd70, []int{5}
}
func (m *TenantData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TenantData.Unmarshal(m, b)
}
func (m *TenantData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TenantData.Marshal(b, m, deterministic)
}
func (dst *TenantData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TenantData.Merge(dst, src)
}
func (m *TenantData) XXX_Size() int {
	return xxx_messageInfo_TenantData.Size(m)
}
func (m *TenantData) XXX_DiscardUnknown() {
	xxx_messageInfo_TenantData.DiscardUnknown(m)
}

var xxx_messageInfo_TenantData proto.InternalMessageInfo

func (m *TenantData) GetEthAccount() *EthereumAccount {
	if m != nil {
		return m.EthAccount
	}
	return nil
}

func (m *TenantData) GetEthDefaultAccountName() string {
	if m != nil {
		return m.EthDefaultAccountName
	}
	return ""
}

func (m *TenantData) GetReceiveEventNotificationEndpoint() string {
	if m != nil {
		return m.ReceiveEventNotificationEndpoint
	}
	return ""
}

func (m *TenantData) GetIdentityId() string {
	if m != nil {
		return m.IdentityId
	}
	return ""
}

func (m *TenantData) GetSigningKeyPair() *KeyPair {
	if m != nil {
		return m.SigningKeyPair
	}
	return nil
}

func (m *TenantData) GetEthauthKeyPair() *KeyPair {
	if m != nil {
		return m.EthauthKeyPair
	}
	return nil
}

type ConfigData struct {
	StoragePath               string             `protobuf:"bytes,1,opt,name=storage_path,json=storagePath,proto3" json:"storage_path,omitempty"`
	P2PPort                   int32              `protobuf:"varint,2,opt,name=p2p_port,json=p2pPort,proto3" json:"p2p_port,omitempty"`
	P2PExternalIp             string             `protobuf:"bytes,3,opt,name=p2p_external_ip,json=p2pExternalIp,proto3" json:"p2p_external_ip,omitempty"`
	P2PConnectionTimeout      *duration.Duration `protobuf:"bytes,4,opt,name=p2p_connection_timeout,json=p2pConnectionTimeout,proto3" json:"p2p_connection_timeout,omitempty"`
	ServerPort                int32              `protobuf:"varint,5,opt,name=server_port,json=serverPort,proto3" json:"server_port,omitempty"`
	ServerAddress             string             `protobuf:"bytes,6,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	NumWorkers                int32              `protobuf:"varint,7,opt,name=num_workers,json=numWorkers,proto3" json:"num_workers,omitempty"`
	WorkerWaitTimeMs          int32              `protobuf:"varint,8,opt,name=worker_wait_time_ms,json=workerWaitTimeMs,proto3" json:"worker_wait_time_ms,omitempty"`
	EthNodeUrl                string             `protobuf:"bytes,9,opt,name=eth_node_url,json=ethNodeUrl,proto3" json:"eth_node_url,omitempty"`
	EthContextReadWaitTimeout *duration.Duration `protobuf:"bytes,10,opt,name=eth_context_read_wait_timeout,json=ethContextReadWaitTimeout,proto3" json:"eth_context_read_wait_timeout,omitempty"`
	EthContextWaitTimeout     *duration.Duration `protobuf:"bytes,11,opt,name=eth_context_wait_timeout,json=ethContextWaitTimeout,proto3" json:"eth_context_wait_timeout,omitempty"`
	EthIntervalRetry          *duration.Duration `protobuf:"bytes,12,opt,name=eth_interval_retry,json=ethIntervalRetry,proto3" json:"eth_interval_retry,omitempty"`
	EthMaxRetries             uint32             `protobuf:"varint,13,opt,name=eth_max_retries,json=ethMaxRetries,proto3" json:"eth_max_retries,omitempty"`
	EthGasPrice               uint64             `protobuf:"varint,14,opt,name=eth_gas_price,json=ethGasPrice,proto3" json:"eth_gas_price,omitempty"`
	EthGasLimit               uint64             `protobuf:"varint,15,opt,name=eth_gas_limit,json=ethGasLimit,proto3" json:"eth_gas_limit,omitempty"`
	TxPoolEnabled             bool               `protobuf:"varint,16,opt,name=tx_pool_enabled,json=txPoolEnabled,proto3" json:"tx_pool_enabled,omitempty"`
	Network                   string             `protobuf:"bytes,17,opt,name=network,proto3" json:"network,omitempty"`
	BootstrapPeers            []string           `protobuf:"bytes,18,rep,name=bootstrap_peers,json=bootstrapPeers,proto3" json:"bootstrap_peers,omitempty"`
	NetworkId                 uint32             `protobuf:"varint,19,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	MainIdentity              *TenantData        `protobuf:"bytes,20,opt,name=main_identity,json=mainIdentity,proto3" json:"main_identity,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}           `json:"-"`
	XXX_unrecognized          []byte             `json:"-"`
	XXX_sizecache             int32              `json:"-"`
}

func (m *ConfigData) Reset()         { *m = ConfigData{} }
func (m *ConfigData) String() string { return proto.CompactTextString(m) }
func (*ConfigData) ProtoMessage()    {}
func (*ConfigData) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5680e2fa0c56bd70, []int{6}
}
func (m *ConfigData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigData.Unmarshal(m, b)
}
func (m *ConfigData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigData.Marshal(b, m, deterministic)
}
func (dst *ConfigData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigData.Merge(dst, src)
}
func (m *ConfigData) XXX_Size() int {
	return xxx_messageInfo_ConfigData.Size(m)
}
func (m *ConfigData) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigData.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigData proto.InternalMessageInfo

func (m *ConfigData) GetStoragePath() string {
	if m != nil {
		return m.StoragePath
	}
	return ""
}

func (m *ConfigData) GetP2PPort() int32 {
	if m != nil {
		return m.P2PPort
	}
	return 0
}

func (m *ConfigData) GetP2PExternalIp() string {
	if m != nil {
		return m.P2PExternalIp
	}
	return ""
}

func (m *ConfigData) GetP2PConnectionTimeout() *duration.Duration {
	if m != nil {
		return m.P2PConnectionTimeout
	}
	return nil
}

func (m *ConfigData) GetServerPort() int32 {
	if m != nil {
		return m.ServerPort
	}
	return 0
}

func (m *ConfigData) GetServerAddress() string {
	if m != nil {
		return m.ServerAddress
	}
	return ""
}

func (m *ConfigData) GetNumWorkers() int32 {
	if m != nil {
		return m.NumWorkers
	}
	return 0
}

func (m *ConfigData) GetWorkerWaitTimeMs() int32 {
	if m != nil {
		return m.WorkerWaitTimeMs
	}
	return 0
}

func (m *ConfigData) GetEthNodeUrl() string {
	if m != nil {
		return m.EthNodeUrl
	}
	return ""
}

func (m *ConfigData) GetEthContextReadWaitTimeout() *duration.Duration {
	if m != nil {
		return m.EthContextReadWaitTimeout
	}
	return nil
}

func (m *ConfigData) GetEthContextWaitTimeout() *duration.Duration {
	if m != nil {
		return m.EthContextWaitTimeout
	}
	return nil
}

func (m *ConfigData) GetEthIntervalRetry() *duration.Duration {
	if m != nil {
		return m.EthIntervalRetry
	}
	return nil
}

func (m *ConfigData) GetEthMaxRetries() uint32 {
	if m != nil {
		return m.EthMaxRetries
	}
	return 0
}

func (m *ConfigData) GetEthGasPrice() uint64 {
	if m != nil {
		return m.EthGasPrice
	}
	return 0
}

func (m *ConfigData) GetEthGasLimit() uint64 {
	if m != nil {
		return m.EthGasLimit
	}
	return 0
}

func (m *ConfigData) GetTxPoolEnabled() bool {
	if m != nil {
		return m.TxPoolEnabled
	}
	return false
}

func (m *ConfigData) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *ConfigData) GetBootstrapPeers() []string {
	if m != nil {
		return m.BootstrapPeers
	}
	return nil
}

func (m *ConfigData) GetNetworkId() uint32 {
	if m != nil {
		return m.NetworkId
	}
	return 0
}

func (m *ConfigData) GetMainIdentity() *TenantData {
	if m != nil {
		return m.MainIdentity
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTenantRequest)(nil), "config.GetTenantRequest")
	proto.RegisterType((*GetAllTenantResponse)(nil), "config.GetAllTenantResponse")
	proto.RegisterType((*UpdateTenantRequest)(nil), "config.UpdateTenantRequest")
	proto.RegisterType((*EthereumAccount)(nil), "config.EthereumAccount")
	proto.RegisterType((*KeyPair)(nil), "config.KeyPair")
	proto.RegisterType((*TenantData)(nil), "config.TenantData")
	proto.RegisterType((*ConfigData)(nil), "config.ConfigData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigServiceClient is the client API for ConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigServiceClient interface {
	GetConfig(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ConfigData, error)
	GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*TenantData, error)
	GetAllTenants(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAllTenantResponse, error)
	CreateConfig(ctx context.Context, in *ConfigData, opts ...grpc.CallOption) (*ConfigData, error)
	CreateTenant(ctx context.Context, in *TenantData, opts ...grpc.CallOption) (*TenantData, error)
	UpdateConfig(ctx context.Context, in *ConfigData, opts ...grpc.CallOption) (*ConfigData, error)
	UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*TenantData, error)
	DeleteConfig(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type configServiceClient struct {
	cc *grpc.ClientConn
}

func NewConfigServiceClient(cc *grpc.ClientConn) ConfigServiceClient {
	return &configServiceClient{cc}
}

func (c *configServiceClient) GetConfig(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ConfigData, error) {
	out := new(ConfigData)
	err := c.cc.Invoke(ctx, "/config.ConfigService/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*TenantData, error) {
	out := new(TenantData)
	err := c.cc.Invoke(ctx, "/config.ConfigService/GetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetAllTenants(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAllTenantResponse, error) {
	out := new(GetAllTenantResponse)
	err := c.cc.Invoke(ctx, "/config.ConfigService/GetAllTenants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) CreateConfig(ctx context.Context, in *ConfigData, opts ...grpc.CallOption) (*ConfigData, error) {
	out := new(ConfigData)
	err := c.cc.Invoke(ctx, "/config.ConfigService/CreateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) CreateTenant(ctx context.Context, in *TenantData, opts ...grpc.CallOption) (*TenantData, error) {
	out := new(TenantData)
	err := c.cc.Invoke(ctx, "/config.ConfigService/CreateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) UpdateConfig(ctx context.Context, in *ConfigData, opts ...grpc.CallOption) (*ConfigData, error) {
	out := new(ConfigData)
	err := c.cc.Invoke(ctx, "/config.ConfigService/UpdateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*TenantData, error) {
	out := new(TenantData)
	err := c.cc.Invoke(ctx, "/config.ConfigService/UpdateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) DeleteConfig(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/config.ConfigService/DeleteConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) DeleteTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/config.ConfigService/DeleteTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServiceServer is the server API for ConfigService service.
type ConfigServiceServer interface {
	GetConfig(context.Context, *empty.Empty) (*ConfigData, error)
	GetTenant(context.Context, *GetTenantRequest) (*TenantData, error)
	GetAllTenants(context.Context, *empty.Empty) (*GetAllTenantResponse, error)
	CreateConfig(context.Context, *ConfigData) (*ConfigData, error)
	CreateTenant(context.Context, *TenantData) (*TenantData, error)
	UpdateConfig(context.Context, *ConfigData) (*ConfigData, error)
	UpdateTenant(context.Context, *UpdateTenantRequest) (*TenantData, error)
	DeleteConfig(context.Context, *empty.Empty) (*empty.Empty, error)
	DeleteTenant(context.Context, *GetTenantRequest) (*empty.Empty, error)
}

func RegisterConfigServiceServer(s *grpc.Server, srv ConfigServiceServer) {
	s.RegisterService(&_ConfigService_serviceDesc, srv)
}

func _ConfigService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetConfig(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/GetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetAllTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetAllTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/GetAllTenants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetAllTenants(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_CreateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).CreateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/CreateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).CreateConfig(ctx, req.(*ConfigData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/CreateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).CreateTenant(ctx, req.(*TenantData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).UpdateConfig(ctx, req.(*ConfigData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_UpdateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).UpdateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/UpdateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).UpdateTenant(ctx, req.(*UpdateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_DeleteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).DeleteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/DeleteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).DeleteConfig(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_DeleteTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).DeleteTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/DeleteTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).DeleteTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "config.ConfigService",
	HandlerType: (*ConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _ConfigService_GetConfig_Handler,
		},
		{
			MethodName: "GetTenant",
			Handler:    _ConfigService_GetTenant_Handler,
		},
		{
			MethodName: "GetAllTenants",
			Handler:    _ConfigService_GetAllTenants_Handler,
		},
		{
			MethodName: "CreateConfig",
			Handler:    _ConfigService_CreateConfig_Handler,
		},
		{
			MethodName: "CreateTenant",
			Handler:    _ConfigService_CreateTenant_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _ConfigService_UpdateConfig_Handler,
		},
		{
			MethodName: "UpdateTenant",
			Handler:    _ConfigService_UpdateTenant_Handler,
		},
		{
			MethodName: "DeleteConfig",
			Handler:    _ConfigService_DeleteConfig_Handler,
		},
		{
			MethodName: "DeleteTenant",
			Handler:    _ConfigService_DeleteTenant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config/service.proto",
}

func init() { proto.RegisterFile("config/service.proto", fileDescriptor_service_5680e2fa0c56bd70) }

var fileDescriptor_service_5680e2fa0c56bd70 = []byte{
	// 1248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xfb, 0x6e, 0x1b, 0xc5,
	0x17, 0x96, 0x9b, 0x36, 0x97, 0xf1, 0x25, 0xc9, 0x24, 0x69, 0xb7, 0xdb, 0xa6, 0xdd, 0xdf, 0xfe,
	0x44, 0xb1, 0xaa, 0x26, 0x96, 0x0c, 0xa2, 0x2d, 0x12, 0x08, 0x37, 0x89, 0xa2, 0x08, 0x5a, 0xac,
	0xa5, 0x55, 0x05, 0x08, 0xad, 0x26, 0xde, 0x13, 0x7b, 0xd5, 0xf5, 0xcc, 0x30, 0x7b, 0x9c, 0xc4,
	0x42, 0x15, 0x82, 0x47, 0x30, 0x8f, 0xc0, 0x7f, 0xbc, 0x0e, 0xaf, 0x00, 0xef, 0x81, 0xe6, 0xb2,
	0x89, 0x1d, 0x3b, 0x89, 0xfa, 0x97, 0x3d, 0xe7, 0xf2, 0x7d, 0xe7, 0x3b, 0x73, 0x66, 0x66, 0xc9,
	0x7a, 0x47, 0xf0, 0xa3, 0xb4, 0xdb, 0xc8, 0x41, 0x1d, 0xa7, 0x1d, 0xd8, 0x96, 0x4a, 0xa0, 0xa0,
	0xf3, 0xd6, 0xea, 0xdf, 0xef, 0x0a, 0xd1, 0xcd, 0xa0, 0xc1, 0x64, 0xda, 0x60, 0x9c, 0x0b, 0x64,
	0x98, 0x0a, 0x9e, 0xdb, 0x28, 0xff, 0x81, 0xf3, 0x9a, 0xd5, 0xe1, 0xe0, 0xa8, 0x91, 0x0c, 0x94,
	0x09, 0x70, 0xfe, 0x7b, 0x17, 0xfd, 0xd0, 0x97, 0x38, 0x74, 0xce, 0x27, 0xe6, 0xa7, 0xb3, 0xd5,
	0x05, 0xbe, 0x95, 0x9f, 0xb0, 0x6e, 0x17, 0x54, 0x43, 0x48, 0x03, 0x3f, 0x4d, 0x15, 0x36, 0xc9,
	0xca, 0x3e, 0xe0, 0x6b, 0xe0, 0x8c, 0x63, 0x04, 0x3f, 0x0f, 0x20, 0x47, 0xfa, 0x80, 0x90, 0x34,
	0x01, 0x8e, 0xe9, 0x51, 0x0a, 0xca, 0x2b, 0x05, 0xa5, 0xfa, 0x52, 0x34, 0x66, 0x09, 0xbf, 0x24,
	0xeb, 0xfb, 0x80, 0xad, 0x2c, 0x2b, 0xd2, 0x72, 0x29, 0x78, 0x0e, 0xf4, 0x11, 0xb9, 0x99, 0x30,
	0x64, 0x5e, 0x29, 0x98, 0xab, 0x97, 0x9b, 0x74, 0xdb, 0x6a, 0xdd, 0xb6, 0x51, 0xbb, 0x0c, 0x59,
	0x64, 0xfc, 0xe1, 0x4f, 0x64, 0xed, 0x8d, 0x4c, 0x18, 0xc2, 0x07, 0xd1, 0x9e, 0xc1, 0xdf, 0x08,
	0x4a, 0x57, 0xc2, 0x7f, 0x4f, 0x96, 0xf7, 0xb0, 0x07, 0x0a, 0x06, 0xfd, 0x56, 0xa7, 0x23, 0x06,
	0x1c, 0xa9, 0x47, 0x16, 0x58, 0x92, 0x28, 0xc8, 0x73, 0x87, 0x5b, 0x2c, 0xe9, 0x0a, 0x99, 0x7b,
	0x07, 0x43, 0x83, 0xb9, 0x14, 0xe9, 0xbf, 0xd4, 0x27, 0x8b, 0x92, 0xe5, 0xf9, 0x89, 0x50, 0x89,
	0x37, 0x67, 0xcc, 0x67, 0xeb, 0x70, 0x8b, 0x2c, 0x7c, 0x0d, 0xc3, 0x36, 0x4b, 0x95, 0x4e, 0x94,
	0x83, 0x43, 0x07, 0xa7, 0xff, 0x1a, 0xcb, 0x31, 0x16, 0x50, 0xf2, 0x18, 0xc3, 0x7f, 0x6f, 0x10,
	0x72, 0x5e, 0x1e, 0x7d, 0x46, 0xca, 0x80, 0xbd, 0x98, 0xd9, 0xa2, 0x4c, 0x6a, 0xb9, 0x79, 0xa7,
	0xd0, 0x71, 0xa1, 0xe6, 0x88, 0x00, 0xf6, 0x8a, 0xfa, 0x9f, 0x12, 0x4f, 0x67, 0x26, 0x70, 0xc4,
	0x06, 0x19, 0x16, 0x08, 0x31, 0x67, 0x7d, 0x70, 0x7c, 0x1b, 0x80, 0xbd, 0x5d, 0xeb, 0x76, 0x49,
	0xaf, 0x58, 0x1f, 0xe8, 0x4b, 0xf2, 0x7f, 0x05, 0x1d, 0x48, 0x8f, 0x21, 0x86, 0x63, 0xd0, 0x29,
	0x42, 0x77, 0xb3, 0x63, 0x66, 0x20, 0x06, 0x9e, 0x48, 0x91, 0x72, 0x74, 0x3a, 0x03, 0x17, 0xba,
	0xa7, 0x23, 0x5f, 0x8d, 0x05, 0xee, 0xb9, 0x38, 0xfa, 0x90, 0x94, 0xed, 0x86, 0xe0, 0x30, 0x4e,
	0x13, 0xef, 0xe6, 0xf8, 0x1e, 0xe1, 0xf0, 0x20, 0xa1, 0xcf, 0xc9, 0x4a, 0x9e, 0x76, 0x79, 0xca,
	0xbb, 0xf1, 0x3b, 0x18, 0xc6, 0x92, 0xa5, 0xca, 0xbb, 0x65, 0x74, 0x2e, 0x17, 0x3a, 0x5d, 0x03,
	0xa3, 0x9a, 0x0b, 0x2c, 0x1a, 0xfa, 0x9c, 0xac, 0x00, 0xf6, 0xd8, 0x00, 0x7b, 0xe7, 0xa9, 0xf3,
	0x97, 0xa4, 0xba, 0x40, 0xb7, 0x0e, 0xff, 0x5a, 0x20, 0x64, 0xc7, 0x84, 0x98, 0x3e, 0xff, 0x8f,
	0x54, 0x72, 0x14, 0x8a, 0x75, 0x21, 0x96, 0x0c, 0x7b, 0x6e, 0x8f, 0xca, 0xce, 0xd6, 0x66, 0xd8,
	0xa3, 0x77, 0xc9, 0xa2, 0x6c, 0xca, 0x58, 0x0a, 0x65, 0x37, 0xec, 0x56, 0xb4, 0x20, 0x9b, 0xb2,
	0x2d, 0x14, 0xd2, 0x47, 0x64, 0x59, 0xbb, 0xe0, 0x14, 0x41, 0x71, 0x96, 0xc5, 0xa9, 0x74, 0xed,
	0xa9, 0xca, 0xa6, 0xdc, 0x73, 0xd6, 0x03, 0x49, 0xbf, 0x25, 0xb7, 0x75, 0x5c, 0x47, 0x70, 0x0e,
	0x1d, 0xd3, 0x4e, 0x4c, 0xfb, 0x20, 0x06, 0x68, 0xda, 0x52, 0x6e, 0xde, 0xdd, 0xb6, 0xa7, 0x74,
	0xbb, 0x38, 0xa5, 0xdb, 0xbb, 0xee, 0x14, 0x47, 0xeb, 0xb2, 0x29, 0x77, 0xce, 0xf2, 0x5e, 0xdb,
	0x34, 0xdd, 0x5c, 0x7d, 0x59, 0x80, 0xb2, 0x65, 0xdd, 0x32, 0x65, 0x11, 0x6b, 0x32, 0x95, 0x7d,
	0x44, 0x6a, 0x2e, 0xa0, 0x18, 0xe6, 0x79, 0x5b, 0x98, 0xb5, 0xb6, 0xdc, 0x48, 0x3f, 0x24, 0x65,
	0x3e, 0xe8, 0xc7, 0x27, 0x42, 0xbd, 0x03, 0x95, 0x7b, 0x0b, 0x16, 0x87, 0x0f, 0xfa, 0x6f, 0xad,
	0x85, 0x6e, 0x91, 0x35, 0xeb, 0x8c, 0x4f, 0x58, 0x8a, 0xa6, 0xec, 0xb8, 0x9f, 0x7b, 0x8b, 0x26,
	0x70, 0xc5, 0xba, 0xde, 0xb2, 0x14, 0x75, 0x61, 0x2f, 0x73, 0x1a, 0x90, 0x8a, 0x1e, 0x3e, 0x2e,
	0x12, 0x88, 0x07, 0x2a, 0xf3, 0x96, 0xec, 0xae, 0x03, 0xf6, 0x5e, 0x89, 0x04, 0xde, 0xa8, 0x8c,
	0xfe, 0x48, 0x36, 0x75, 0x44, 0x47, 0x70, 0x84, 0x53, 0x8c, 0x15, 0xb0, 0xe4, 0x1c, 0x5a, 0x77,
	0x84, 0x5c, 0xd7, 0x91, 0xbb, 0x80, 0xbd, 0x1d, 0x9b, 0x1e, 0x01, 0x4b, 0x0a, 0x76, 0xdd, 0x96,
	0xc8, 0xce, 0x7e, 0x01, 0x3e, 0x81, 0x5b, 0xbe, 0x0e, 0x77, 0xe3, 0x1c, 0x77, 0x1c, 0x73, 0x9f,
	0x50, 0x8d, 0x99, 0x72, 0x04, 0x75, 0xcc, 0xb2, 0x58, 0x01, 0xaa, 0xa1, 0x57, 0xb9, 0x0e, 0x4d,
	0x0f, 0xe8, 0x81, 0xcb, 0x89, 0x74, 0x8a, 0x1e, 0x16, 0x0d, 0xd4, 0x67, 0xa7, 0x06, 0x23, 0x85,
	0xdc, 0xab, 0x06, 0xa5, 0x7a, 0x35, 0xaa, 0x02, 0xf6, 0x5e, 0xb2, 0xd3, 0xc8, 0x1a, 0x69, 0x48,
	0xb4, 0x21, 0xee, 0xb2, 0x3c, 0x96, 0x2a, 0xed, 0x80, 0x57, 0x0b, 0x4a, 0xf5, 0x9b, 0x91, 0xbe,
	0x0f, 0xf6, 0x59, 0xde, 0xd6, 0xa6, 0xf1, 0x98, 0x2c, 0xed, 0xa7, 0xe8, 0x2d, 0x8f, 0xc7, 0x7c,
	0xa3, 0x4d, 0x9a, 0x0f, 0x4f, 0x63, 0x29, 0x44, 0x16, 0x03, 0x67, 0x87, 0x19, 0x24, 0xde, 0x4a,
	0x50, 0xaa, 0x2f, 0x46, 0x55, 0x3c, 0x6d, 0x0b, 0x91, 0xed, 0x59, 0xa3, 0xbe, 0xf0, 0x38, 0xa0,
	0xde, 0x4a, 0x6f, 0xd5, 0x5e, 0x78, 0x6e, 0x49, 0x3f, 0x26, 0xcb, 0x87, 0x42, 0x60, 0x8e, 0x8a,
	0xc9, 0x58, 0x82, 0x9e, 0x10, 0x1a, 0xcc, 0xd5, 0x97, 0xa2, 0xda, 0x99, 0xb9, 0xad, 0xad, 0x74,
	0x93, 0x10, 0x97, 0xa3, 0x8f, 0xfa, 0x9a, 0x51, 0xb5, 0xe4, 0x2c, 0x07, 0x09, 0x7d, 0x4a, 0xaa,
	0x7d, 0x96, 0xf2, 0xb8, 0x38, 0xfc, 0xde, 0xfa, 0xa5, 0xd7, 0x72, 0x45, 0x07, 0x1e, 0xb8, 0xb8,
	0xe6, 0x9f, 0x8b, 0xa4, 0x6a, 0x0f, 0xeb, 0x77, 0xf6, 0x69, 0xa4, 0x8c, 0x2c, 0xed, 0x03, 0x5a,
	0x1b, 0xbd, 0x3d, 0xd5, 0xfe, 0x3d, 0xfd, 0xb8, 0xf9, 0x67, 0xc0, 0xe7, 0x07, 0x3d, 0xac, 0x8f,
	0x5a, 0xab, 0xfe, 0xf2, 0x3e, 0x60, 0xa0, 0x27, 0x31, 0xb0, 0x9e, 0xdf, 0xff, 0xfe, 0xe7, 0x8f,
	0x1b, 0x35, 0x5a, 0x69, 0xb8, 0x07, 0x58, 0xcf, 0x2d, 0x45, 0x43, 0x61, 0x6b, 0xa2, 0x5e, 0x01,
	0x75, 0xf1, 0xe5, 0xf3, 0x67, 0x54, 0x1f, 0x3e, 0x1f, 0xb5, 0xd6, 0xfc, 0x55, 0x4d, 0x62, 0x8d,
	0xe3, 0x34, 0x9b, 0xf4, 0x5e, 0x41, 0x83, 0xc6, 0xd9, 0xf8, 0xe5, 0xfc, 0xc1, 0x7a, 0x4f, 0xdf,
	0x93, 0xea, 0xf8, 0x43, 0x99, 0x5f, 0x2a, 0xee, 0xfe, 0x58, 0x45, 0x53, 0xef, 0x6a, 0xf8, 0xd9,
	0xa8, 0xe5, 0xf9, 0xb7, 0x75, 0x05, 0xad, 0x2c, 0x9b, 0xac, 0x22, 0x37, 0x65, 0x6c, 0xd0, 0xb5,
	0x0b, 0x65, 0x64, 0x69, 0x8e, 0x34, 0x23, 0x95, 0x1d, 0x05, 0x0c, 0xc1, 0xb5, 0x76, 0x46, 0x0b,
	0x67, 0xb6, 0xf5, 0xd3, 0x51, 0xcb, 0xf7, 0x3d, 0x9b, 0x9a, 0x07, 0xba, 0x7f, 0x81, 0x0d, 0x0a,
	0xf4, 0xfb, 0x6a, 0x18, 0x57, 0xc3, 0x89, 0xfe, 0x7e, 0x5e, 0x7a, 0x4c, 0x65, 0xc1, 0xe6, 0xba,
	0x3c, 0xa3, 0x97, 0x33, 0xfb, 0xfb, 0x6c, 0xd4, 0xba, 0xef, 0xfb, 0x05, 0x9b, 0xad, 0x7f, 0x8a,
	0x6f, 0x2d, 0xac, 0x4d, 0x2a, 0xd4, 0x8c, 0x5d, 0x52, 0xb1, 0xdf, 0x11, 0x1f, 0xa8, 0xaf, 0x31,
	0x6a, 0x6d, 0xf8, 0xee, 0x13, 0x64, 0x42, 0x9f, 0x95, 0xe6, 0x4f, 0x49, 0xfb, 0xad, 0x54, 0x30,
	0x39, 0x6d, 0xf7, 0x0a, 0xd4, 0x19, 0xdf, 0x31, 0x33, 0x45, 0x7e, 0x35, 0x6a, 0xdd, 0xf1, 0x37,
	0x0a, 0xca, 0x09, 0x91, 0x86, 0x34, 0xf0, 0xaf, 0x1a, 0x24, 0x5d, 0x43, 0x46, 0x2a, 0xbb, 0x90,
	0xc1, 0x99, 0xd8, 0xcb, 0x46, 0xe9, 0x12, 0x7b, 0xf8, 0xc4, 0x88, 0xb6, 0x10, 0xd3, 0xa2, 0x6b,
	0x8f, 0x27, 0xcf, 0xcb, 0xaf, 0x05, 0xdb, 0xb5, 0x47, 0xe6, 0x32, 0xbe, 0x2f, 0x8c, 0xe2, 0x82,
	0x6f, 0x5a, 0xf1, 0xe6, 0xe3, 0xab, 0x14, 0xbf, 0x78, 0x44, 0x48, 0x47, 0xf4, 0x1d, 0xeb, 0x8b,
	0x8a, 0xbb, 0x2a, 0xda, 0x9a, 0xa3, 0x5d, 0xfa, 0x61, 0xd1, 0xda, 0xe5, 0xe1, 0xe1, 0xbc, 0xa1,
	0xfd, 0xe4, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x06, 0x5c, 0x73, 0x03, 0x6f, 0x0b, 0x00, 0x00,
}
