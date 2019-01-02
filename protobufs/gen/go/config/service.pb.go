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
	return fileDescriptor_service_48f740854e8728dc, []int{0}
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
	return fileDescriptor_service_48f740854e8728dc, []int{1}
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
	return fileDescriptor_service_48f740854e8728dc, []int{2}
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
	return fileDescriptor_service_48f740854e8728dc, []int{3}
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
	return fileDescriptor_service_48f740854e8728dc, []int{4}
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
	return fileDescriptor_service_48f740854e8728dc, []int{5}
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
	return fileDescriptor_service_48f740854e8728dc, []int{6}
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

func init() { proto.RegisterFile("config/service.proto", fileDescriptor_service_48f740854e8728dc) }

var fileDescriptor_service_48f740854e8728dc = []byte{
	// 1243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xeb, 0x6e, 0x13, 0x47,
	0x14, 0x96, 0x09, 0xe4, 0x32, 0xb6, 0x73, 0x99, 0x24, 0xb0, 0x2c, 0x21, 0x2c, 0x5b, 0x95, 0x5a,
	0x88, 0xc4, 0x92, 0x5b, 0x89, 0xc2, 0x0f, 0x24, 0x93, 0x44, 0x51, 0xd4, 0x42, 0xad, 0x2d, 0x08,
	0xb5, 0x55, 0xb5, 0x9a, 0x78, 0x4f, 0xec, 0x15, 0xeb, 0x99, 0x61, 0xf6, 0x6c, 0x12, 0xab, 0xaa,
	0x54, 0xf1, 0x08, 0xee, 0x2b, 0xf4, 0x57, 0x5f, 0xa7, 0xaf, 0xd0, 0xbe, 0x47, 0x35, 0x97, 0x4d,
	0x9c, 0xd8, 0x21, 0xe2, 0x97, 0x3d, 0xe7, 0xf2, 0x7d, 0xe7, 0x3b, 0x73, 0x66, 0x66, 0xc9, 0x5a,
	0x57, 0xf0, 0xa3, 0xb4, 0xd7, 0xcc, 0x41, 0x1d, 0xa7, 0x5d, 0xd8, 0x96, 0x4a, 0xa0, 0xa0, 0xb3,
	0xd6, 0xea, 0x6f, 0xf4, 0x84, 0xe8, 0x65, 0xd0, 0x64, 0x32, 0x6d, 0x32, 0xce, 0x05, 0x32, 0x4c,
	0x05, 0xcf, 0x6d, 0x94, 0xbf, 0xe9, 0xbc, 0x66, 0x75, 0x58, 0x1c, 0x35, 0x93, 0x42, 0x99, 0x00,
	0xe7, 0xbf, 0x77, 0xd9, 0x0f, 0x03, 0x89, 0x43, 0xe7, 0x7c, 0x62, 0x7e, 0xba, 0x5b, 0x3d, 0xe0,
	0x5b, 0xf9, 0x09, 0xeb, 0xf5, 0x40, 0x35, 0x85, 0x34, 0xf0, 0x93, 0x54, 0x61, 0x8b, 0x2c, 0xef,
	0x03, 0xbe, 0x01, 0xce, 0x38, 0x46, 0xf0, 0xa1, 0x80, 0x1c, 0xe9, 0x26, 0x21, 0x69, 0x02, 0x1c,
	0xd3, 0xa3, 0x14, 0x94, 0x57, 0x09, 0x2a, 0x8d, 0x85, 0x68, 0xcc, 0x12, 0xbe, 0x20, 0x6b, 0xfb,
	0x80, 0xed, 0x2c, 0x2b, 0xd3, 0x72, 0x29, 0x78, 0x0e, 0xf4, 0x11, 0xb9, 0x99, 0x30, 0x64, 0x5e,
	0x25, 0x98, 0x69, 0x54, 0x5b, 0x74, 0xdb, 0x6a, 0xdd, 0xb6, 0x51, 0xbb, 0x0c, 0x59, 0x64, 0xfc,
	0xe1, 0xaf, 0x64, 0xf5, 0xad, 0x4c, 0x18, 0xc2, 0x67, 0xd1, 0x9e, 0xc1, 0xdf, 0x08, 0x2a, 0x9f,
	0x84, 0xff, 0x89, 0x2c, 0xed, 0x61, 0x1f, 0x14, 0x14, 0x83, 0x76, 0xb7, 0x2b, 0x0a, 0x8e, 0xd4,
	0x23, 0x73, 0x2c, 0x49, 0x14, 0xe4, 0xb9, 0xc3, 0x2d, 0x97, 0x74, 0x99, 0xcc, 0xbc, 0x87, 0xa1,
	0xc1, 0x5c, 0x88, 0xf4, 0x5f, 0xea, 0x93, 0x79, 0xc9, 0xf2, 0xfc, 0x44, 0xa8, 0xc4, 0x9b, 0x31,
	0xe6, 0xb3, 0x75, 0xb8, 0x45, 0xe6, 0xbe, 0x83, 0x61, 0x87, 0xa5, 0x4a, 0x27, 0xca, 0xe2, 0xd0,
	0xc1, 0xe9, 0xbf, 0xc6, 0x72, 0x8c, 0x25, 0x94, 0x3c, 0xc6, 0xf0, 0xbf, 0x1b, 0x84, 0x9c, 0x97,
	0x47, 0xbf, 0x25, 0x55, 0xc0, 0x7e, 0xcc, 0x6c, 0x51, 0x26, 0xb5, 0xda, 0xba, 0x53, 0xea, 0xb8,
	0x54, 0x73, 0x44, 0x00, 0xfb, 0x65, 0xfd, 0x4f, 0x89, 0xa7, 0x33, 0x13, 0x38, 0x62, 0x45, 0x86,
	0x25, 0x42, 0xcc, 0xd9, 0x00, 0x1c, 0xdf, 0x3a, 0x60, 0x7f, 0xd7, 0xba, 0x5d, 0xd2, 0x6b, 0x36,
	0x00, 0xfa, 0x8a, 0x7c, 0xa1, 0xa0, 0x0b, 0xe9, 0x31, 0xc4, 0x70, 0x0c, 0x3a, 0x45, 0xe8, 0x6e,
	0x76, 0xcd, 0x0c, 0xc4, 0xc0, 0x13, 0x29, 0x52, 0x8e, 0x4e, 0x67, 0xe0, 0x42, 0xf7, 0x74, 0xe4,
	0xeb, 0xb1, 0xc0, 0x3d, 0x17, 0x47, 0x1f, 0x90, 0xaa, 0xdd, 0x10, 0x1c, 0xc6, 0x69, 0xe2, 0xdd,
	0x1c, 0xdf, 0x23, 0x1c, 0x1e, 0x24, 0xf4, 0x19, 0x59, 0xce, 0xd3, 0x1e, 0x4f, 0x79, 0x2f, 0x7e,
	0x0f, 0xc3, 0x58, 0xb2, 0x54, 0x79, 0xb7, 0x8c, 0xce, 0xa5, 0x52, 0xa7, 0x6b, 0x60, 0xb4, 0xe8,
	0x02, 0xcb, 0x86, 0x3e, 0x23, 0xcb, 0x80, 0x7d, 0x56, 0x60, 0xff, 0x3c, 0x75, 0xf6, 0x8a, 0x54,
	0x17, 0xe8, 0xd6, 0xe1, 0xdf, 0x73, 0x84, 0xec, 0x98, 0x10, 0xd3, 0xe7, 0x87, 0xa4, 0x96, 0xa3,
	0x50, 0xac, 0x07, 0xb1, 0x64, 0xd8, 0x77, 0x7b, 0x54, 0x75, 0xb6, 0x0e, 0xc3, 0x3e, 0xbd, 0x4b,
	0xe6, 0x65, 0x4b, 0xc6, 0x52, 0x28, 0xbb, 0x61, 0xb7, 0xa2, 0x39, 0xd9, 0x92, 0x1d, 0xa1, 0x90,
	0x3e, 0x22, 0x4b, 0xda, 0x05, 0xa7, 0x08, 0x8a, 0xb3, 0x2c, 0x4e, 0xa5, 0x6b, 0x4f, 0x5d, 0xb6,
	0xe4, 0x9e, 0xb3, 0x1e, 0x48, 0xfa, 0x03, 0xb9, 0xad, 0xe3, 0xba, 0x82, 0x73, 0xe8, 0x9a, 0x76,
	0x62, 0x3a, 0x00, 0x51, 0xa0, 0x69, 0x4b, 0xb5, 0x75, 0x77, 0xdb, 0x9e, 0xd2, 0xed, 0xf2, 0x94,
	0x6e, 0xef, 0xba, 0x53, 0x1c, 0xad, 0xc9, 0x96, 0xdc, 0x39, 0xcb, 0x7b, 0x63, 0xd3, 0x74, 0x73,
	0xf5, 0x65, 0x01, 0xca, 0x96, 0x75, 0xcb, 0x94, 0x45, 0xac, 0xc9, 0x54, 0xf6, 0x25, 0x59, 0x74,
	0x01, 0xe5, 0x30, 0xcf, 0xda, 0xc2, 0xac, 0xb5, 0xed, 0x46, 0xfa, 0x01, 0xa9, 0xf2, 0x62, 0x10,
	0x9f, 0x08, 0xf5, 0x1e, 0x54, 0xee, 0xcd, 0x59, 0x1c, 0x5e, 0x0c, 0xde, 0x59, 0x0b, 0xdd, 0x22,
	0xab, 0xd6, 0x19, 0x9f, 0xb0, 0x14, 0x4d, 0xd9, 0xf1, 0x20, 0xf7, 0xe6, 0x4d, 0xe0, 0xb2, 0x75,
	0xbd, 0x63, 0x29, 0xea, 0xc2, 0x5e, 0xe5, 0x34, 0x20, 0x35, 0x3d, 0x7c, 0x5c, 0x24, 0x10, 0x17,
	0x2a, 0xf3, 0x16, 0xec, 0xae, 0x03, 0xf6, 0x5f, 0x8b, 0x04, 0xde, 0xaa, 0x8c, 0xfe, 0x42, 0xee,
	0xeb, 0x88, 0xae, 0xe0, 0x08, 0xa7, 0x18, 0x2b, 0x60, 0xc9, 0x39, 0xb4, 0xee, 0x08, 0xb9, 0xae,
	0x23, 0x77, 0x01, 0xfb, 0x3b, 0x36, 0x3d, 0x02, 0x96, 0x94, 0xec, 0xba, 0x2d, 0x91, 0x9d, 0xfd,
	0x12, 0xfc, 0x02, 0x6e, 0xf5, 0x3a, 0xdc, 0xf5, 0x73, 0xdc, 0x71, 0xcc, 0x7d, 0x42, 0x35, 0x66,
	0xca, 0x11, 0xd4, 0x31, 0xcb, 0x62, 0x05, 0xa8, 0x86, 0x5e, 0xed, 0x3a, 0x34, 0x3d, 0xa0, 0x07,
	0x2e, 0x27, 0xd2, 0x29, 0x7a, 0x58, 0x34, 0xd0, 0x80, 0x9d, 0x1a, 0x8c, 0x14, 0x72, 0xaf, 0x1e,
	0x54, 0x1a, 0xf5, 0xa8, 0x0e, 0xd8, 0x7f, 0xc5, 0x4e, 0x23, 0x6b, 0xa4, 0x21, 0xd1, 0x86, 0xb8,
	0xc7, 0xf2, 0x58, 0xaa, 0xb4, 0x0b, 0xde, 0x62, 0x50, 0x69, 0xdc, 0x8c, 0xf4, 0x7d, 0xb0, 0xcf,
	0xf2, 0x8e, 0x36, 0x8d, 0xc7, 0x64, 0xe9, 0x20, 0x45, 0x6f, 0x69, 0x3c, 0xe6, 0x7b, 0x6d, 0xd2,
	0x7c, 0x78, 0x1a, 0x4b, 0x21, 0xb2, 0x18, 0x38, 0x3b, 0xcc, 0x20, 0xf1, 0x96, 0x83, 0x4a, 0x63,
	0x3e, 0xaa, 0xe3, 0x69, 0x47, 0x88, 0x6c, 0xcf, 0x1a, 0xf5, 0x85, 0xc7, 0x01, 0xf5, 0x56, 0x7a,
	0x2b, 0xf6, 0xc2, 0x73, 0x4b, 0xfa, 0x15, 0x59, 0x3a, 0x14, 0x02, 0x73, 0x54, 0x4c, 0xc6, 0x12,
	0xf4, 0x84, 0xd0, 0x60, 0xa6, 0xb1, 0x10, 0x2d, 0x9e, 0x99, 0x3b, 0xda, 0x4a, 0xef, 0x13, 0xe2,
	0x72, 0xf4, 0x51, 0x5f, 0x35, 0xaa, 0x16, 0x9c, 0xe5, 0x20, 0xa1, 0x4f, 0x49, 0x7d, 0xc0, 0x52,
	0x1e, 0x97, 0x87, 0xdf, 0x5b, 0xbb, 0xf2, 0x5a, 0xae, 0xe9, 0xc0, 0x03, 0x17, 0xd7, 0xfa, 0x6b,
	0x9e, 0xd4, 0xed, 0x61, 0xfd, 0xd1, 0x3e, 0x8d, 0x94, 0x91, 0x85, 0x7d, 0x40, 0x6b, 0xa3, 0xb7,
	0x27, 0xda, 0xbf, 0xa7, 0x1f, 0x37, 0xff, 0x0c, 0xf8, 0xfc, 0xa0, 0x87, 0x8d, 0x51, 0x7b, 0xc5,
	0x5f, 0xda, 0x07, 0x0c, 0xf4, 0x24, 0x06, 0xd6, 0xf3, 0xf1, 0x9f, 0x7f, 0xff, 0xbc, 0xb1, 0x48,
	0x6b, 0x4d, 0xf7, 0x00, 0xeb, 0xb9, 0xa5, 0x85, 0xa1, 0xb0, 0x35, 0x51, 0xaf, 0x84, 0xba, 0xfc,
	0xf2, 0xf9, 0x53, 0xaa, 0x0f, 0x9f, 0x8f, 0xda, 0xab, 0xfe, 0x8a, 0x26, 0xb1, 0xc6, 0x71, 0x9a,
	0x4d, 0xba, 0x51, 0xd2, 0xa0, 0x71, 0xe6, 0xcd, 0xdf, 0xce, 0x5f, 0xac, 0xdf, 0xe9, 0x90, 0xd4,
	0xc7, 0x5f, 0xca, 0xfc, 0x4a, 0x75, 0x1b, 0x63, 0x25, 0x4d, 0x3c, 0xac, 0x61, 0x6b, 0xd4, 0xf6,
	0xfc, 0xdb, 0xba, 0x84, 0x76, 0x96, 0x5d, 0x2c, 0x23, 0x37, 0x75, 0xac, 0xd0, 0xa5, 0x4b, 0x75,
	0xd0, 0x8c, 0xd4, 0x76, 0x14, 0x30, 0x04, 0xd7, 0xd7, 0x29, 0xfd, 0x9b, 0xda, 0xd3, 0x6f, 0x46,
	0x6d, 0xdf, 0xf7, 0x6c, 0x6a, 0x1e, 0xe8, 0xe6, 0x05, 0x36, 0x28, 0xd0, 0x8f, 0xab, 0x65, 0x0b,
	0x2f, 0x34, 0xf7, 0x79, 0xe5, 0x31, 0xfd, 0x50, 0xb2, 0xb9, 0x16, 0x4f, 0x69, 0xe4, 0xd4, 0xe6,
	0x3e, 0x1b, 0xb5, 0x37, 0x7c, 0xbf, 0x64, 0xb3, 0xb5, 0x4f, 0xf0, 0xad, 0x85, 0x97, 0xd5, 0x69,
	0xca, 0x1e, 0xa9, 0xd9, 0xaf, 0x88, 0xcf, 0x14, 0xd8, 0x1c, 0xb5, 0xd7, 0x7d, 0xf7, 0x01, 0x72,
	0x41, 0xa0, 0xd5, 0xe6, 0x4f, 0x68, 0xfb, 0x58, 0x29, 0x99, 0x9c, 0xb8, 0x7b, 0x25, 0xea, 0x94,
	0xaf, 0x98, 0xa9, 0x2a, 0xdb, 0xa3, 0xf6, 0x1d, 0x7f, 0xbd, 0xa4, 0xbc, 0xa0, 0xd2, 0x90, 0x3e,
	0xf4, 0x3f, 0x39, 0x46, 0xba, 0x88, 0x8c, 0xd4, 0x76, 0x21, 0x83, 0x33, 0xb5, 0x57, 0x0d, 0xd2,
	0x15, 0xf6, 0xf0, 0x89, 0x51, 0x6d, 0x21, 0x26, 0x55, 0x2f, 0x3e, 0xbe, 0x78, 0x5c, 0xfe, 0xa8,
	0x94, 0x74, 0xd7, 0x1e, 0x99, 0xab, 0x08, 0x5f, 0x18, 0xcd, 0x25, 0xe1, 0xa4, 0xe6, 0xcd, 0xc7,
	0x9f, 0xd4, 0xfc, 0xf2, 0x11, 0x21, 0x5d, 0x31, 0x70, 0xb4, 0x2f, 0x6b, 0xee, 0xae, 0xe8, 0x68,
	0x92, 0x4e, 0xe5, 0xe7, 0x79, 0x6b, 0x97, 0x87, 0x87, 0xb3, 0x86, 0xf7, 0xeb, 0xff, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xd9, 0x4e, 0xef, 0x2d, 0x70, 0x0b, 0x00, 0x00,
}
