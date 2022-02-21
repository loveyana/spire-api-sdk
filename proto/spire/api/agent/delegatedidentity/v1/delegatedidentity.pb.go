// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: spire/api/agent/delegatedidentity/v1/delegatedidentity.proto

package delegatedidentityv1

import (
	types "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// X.509 SPIFFE Verifiable Identity Document with the private key.
type X509SVIDWithKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The workload X509-SVID.
	X509Svid *types.X509SVID `protobuf:"bytes,1,opt,name=x509_svid,json=x509Svid,proto3" json:"x509_svid,omitempty"`
	// Private key (encoding DER PKCS#8).
	X509SvidKey []byte `protobuf:"bytes,2,opt,name=x509_svid_key,json=x509SvidKey,proto3" json:"x509_svid_key,omitempty"`
}

func (x *X509SVIDWithKey) Reset() {
	*x = X509SVIDWithKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *X509SVIDWithKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*X509SVIDWithKey) ProtoMessage() {}

func (x *X509SVIDWithKey) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use X509SVIDWithKey.ProtoReflect.Descriptor instead.
func (*X509SVIDWithKey) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDescGZIP(), []int{0}
}

func (x *X509SVIDWithKey) GetX509Svid() *types.X509SVID {
	if x != nil {
		return x.X509Svid
	}
	return nil
}

func (x *X509SVIDWithKey) GetX509SvidKey() []byte {
	if x != nil {
		return x.X509SvidKey
	}
	return nil
}

// SubscribeToX509SVIDsRequest is used by clients to subscribe the set of SVIDs that
// any given workload is entitled to. Clients subscribe to a workload's SVIDs by providing
// a set of selectors describing the workload.
type SubscribeToX509SVIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Selectors describing the workload to subscribe to.
	Selectors []*types.Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
}

func (x *SubscribeToX509SVIDsRequest) Reset() {
	*x = SubscribeToX509SVIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToX509SVIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToX509SVIDsRequest) ProtoMessage() {}

func (x *SubscribeToX509SVIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeToX509SVIDsRequest.ProtoReflect.Descriptor instead.
func (*SubscribeToX509SVIDsRequest) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDescGZIP(), []int{1}
}

func (x *SubscribeToX509SVIDsRequest) GetSelectors() []*types.Selector {
	if x != nil {
		return x.Selectors
	}
	return nil
}

type SubscribeToX509SVIDsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X509Svids []*X509SVIDWithKey `protobuf:"bytes,1,rep,name=x509_svids,json=x509Svids,proto3" json:"x509_svids,omitempty"`
	// Names of the trust domains that this workload should federates with.
	FederatesWith []string `protobuf:"bytes,2,rep,name=federates_with,json=federatesWith,proto3" json:"federates_with,omitempty"`
}

func (x *SubscribeToX509SVIDsResponse) Reset() {
	*x = SubscribeToX509SVIDsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToX509SVIDsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToX509SVIDsResponse) ProtoMessage() {}

func (x *SubscribeToX509SVIDsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeToX509SVIDsResponse.ProtoReflect.Descriptor instead.
func (*SubscribeToX509SVIDsResponse) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeToX509SVIDsResponse) GetX509Svids() []*X509SVIDWithKey {
	if x != nil {
		return x.X509Svids
	}
	return nil
}

func (x *SubscribeToX509SVIDsResponse) GetFederatesWith() []string {
	if x != nil {
		return x.FederatesWith
	}
	return nil
}

type SubscribeToX509BundlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeToX509BundlesRequest) Reset() {
	*x = SubscribeToX509BundlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToX509BundlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToX509BundlesRequest) ProtoMessage() {}

func (x *SubscribeToX509BundlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeToX509BundlesRequest.ProtoReflect.Descriptor instead.
func (*SubscribeToX509BundlesRequest) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDescGZIP(), []int{3}
}

// SubscribeToX509BundlesResponse contains all bundles that the agent is tracking,
// including the local bundle. When an update occurs, or bundles are added or removed,
// a new response with the full set of bundles is sent.
type SubscribeToX509BundlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A map keyed by trust domain name, with ASN.1 DER-encoded
	// X.509 CA certificates as the values
	CaCertificates map[string][]byte `protobuf:"bytes,1,rep,name=ca_certificates,json=caCertificates,proto3" json:"ca_certificates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SubscribeToX509BundlesResponse) Reset() {
	*x = SubscribeToX509BundlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToX509BundlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToX509BundlesResponse) ProtoMessage() {}

func (x *SubscribeToX509BundlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeToX509BundlesResponse.ProtoReflect.Descriptor instead.
func (*SubscribeToX509BundlesResponse) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDescGZIP(), []int{4}
}

func (x *SubscribeToX509BundlesResponse) GetCaCertificates() map[string][]byte {
	if x != nil {
		return x.CaCertificates
	}
	return nil
}

type FetchJWTSVIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The audience(s) the workload intends to authenticate against.
	Audience []string `protobuf:"bytes,1,rep,name=audience,proto3" json:"audience,omitempty"`
	// Required. Selectors describing the workload to fetch.
	Selectors []*types.Selector `protobuf:"bytes,2,rep,name=selectors,proto3" json:"selectors,omitempty"`
}

func (x *FetchJWTSVIDsRequest) Reset() {
	*x = FetchJWTSVIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchJWTSVIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchJWTSVIDsRequest) ProtoMessage() {}

func (x *FetchJWTSVIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchJWTSVIDsRequest.ProtoReflect.Descriptor instead.
func (*FetchJWTSVIDsRequest) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDescGZIP(), []int{5}
}

func (x *FetchJWTSVIDsRequest) GetAudience() []string {
	if x != nil {
		return x.Audience
	}
	return nil
}

func (x *FetchJWTSVIDsRequest) GetSelectors() []*types.Selector {
	if x != nil {
		return x.Selectors
	}
	return nil
}

// The FetchJWTSVIDsResponse message conveys JWT-SVIDs.
type FetchJWTSVIDsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The list of returned JWT-SVIDs.
	Svids []*types.JWTSVID `protobuf:"bytes,1,rep,name=svids,proto3" json:"svids,omitempty"`
}

func (x *FetchJWTSVIDsResponse) Reset() {
	*x = FetchJWTSVIDsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchJWTSVIDsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchJWTSVIDsResponse) ProtoMessage() {}

func (x *FetchJWTSVIDsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchJWTSVIDsResponse.ProtoReflect.Descriptor instead.
func (*FetchJWTSVIDsResponse) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDescGZIP(), []int{6}
}

func (x *FetchJWTSVIDsResponse) GetSvids() []*types.JWTSVID {
	if x != nil {
		return x.Svids
	}
	return nil
}

// The SubscribeToJWTBundlesRequest message conveys parameters for requesting JWKS bundles.
// There are currently no such parameters.
type SubscribeToJWTBundlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeToJWTBundlesRequest) Reset() {
	*x = SubscribeToJWTBundlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToJWTBundlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToJWTBundlesRequest) ProtoMessage() {}

func (x *SubscribeToJWTBundlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeToJWTBundlesRequest.ProtoReflect.Descriptor instead.
func (*SubscribeToJWTBundlesRequest) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDescGZIP(), []int{7}
}

// The SubscribeToJWTBundlesReponse conveys JWKS bundles.
type SubscribeToJWTBundlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. JWK encoded JWT bundles, keyed by the SPIFFE ID of the trust
	// domain.
	Bundles map[string][]byte `protobuf:"bytes,1,rep,name=bundles,proto3" json:"bundles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SubscribeToJWTBundlesResponse) Reset() {
	*x = SubscribeToJWTBundlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToJWTBundlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToJWTBundlesResponse) ProtoMessage() {}

func (x *SubscribeToJWTBundlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeToJWTBundlesResponse.ProtoReflect.Descriptor instead.
func (*SubscribeToJWTBundlesResponse) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDescGZIP(), []int{8}
}

func (x *SubscribeToJWTBundlesResponse) GetBundles() map[string][]byte {
	if x != nil {
		return x.Bundles
	}
	return nil
}

var File_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto protoreflect.FileDescriptor

var file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x78, 0x35, 0x30, 0x39, 0x73, 0x76, 0x69, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6a, 0x77, 0x74, 0x73, 0x76, 0x69, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x0f, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x57,
	0x69, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x09, 0x78, 0x35, 0x30, 0x39, 0x5f, 0x73,
	0x76, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x58, 0x35, 0x30, 0x39,
	0x53, 0x56, 0x49, 0x44, 0x52, 0x08, 0x78, 0x35, 0x30, 0x39, 0x53, 0x76, 0x69, 0x64, 0x12, 0x22,
	0x0a, 0x0d, 0x78, 0x35, 0x30, 0x39, 0x5f, 0x73, 0x76, 0x69, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x78, 0x35, 0x30, 0x39, 0x53, 0x76, 0x69, 0x64, 0x4b,
	0x65, 0x79, 0x22, 0x56, 0x0a, 0x1b, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54,
	0x6f, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x1c, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56,
	0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0a, 0x78,
	0x35, 0x30, 0x39, 0x5f, 0x73, 0x76, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x57,
	0x69, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x78, 0x35, 0x30, 0x39, 0x53, 0x76, 0x69, 0x64,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x77,
	0x69, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x22, 0x1f, 0x0a, 0x1d, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x58, 0x35, 0x30, 0x39, 0x42, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xe7, 0x01, 0x0a, 0x1e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x58, 0x35, 0x30, 0x39, 0x42, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x81, 0x01, 0x0a,
	0x0f, 0x63, 0x61, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x58, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x58, 0x35, 0x30, 0x39, 0x42, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0e, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73,
	0x1a, 0x41, 0x0a, 0x13, 0x43, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x6b, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4a, 0x57, 0x54, 0x53,
	0x56, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x22, 0x47, 0x0a, 0x15, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49, 0x44,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x73, 0x76, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x53, 0x56,
	0x49, 0x44, 0x52, 0x05, 0x73, 0x76, 0x69, 0x64, 0x73, 0x22, 0x1e, 0x0a, 0x1c, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x4a, 0x57, 0x54, 0x42, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xc7, 0x01, 0x0a, 0x1d, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x4a, 0x57, 0x54, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x07, 0x62,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x4a,
	0x57, 0x54, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x42, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x32, 0x8d, 0x05, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x9f, 0x01, 0x0a, 0x14, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49,
	0x44, 0x73, 0x12, 0x41, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x54, 0x6f, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0xa5, 0x01, 0x0a, 0x16,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x58, 0x35, 0x30, 0x39, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x12, 0x43, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x58, 0x35, 0x30, 0x39, 0x42, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x64, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x58, 0x35,
	0x30, 0x39, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x12, 0x88, 0x01, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4a, 0x57, 0x54,
	0x53, 0x56, 0x49, 0x44, 0x73, 0x12, 0x3a, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3b, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4a, 0x57,
	0x54, 0x53, 0x56, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa2,
	0x01, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x4a, 0x57,
	0x54, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x12, 0x42, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x4a, 0x57, 0x54, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x4a,
	0x57, 0x54, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x42, 0x60, 0x5a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x31, 0x3b, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDescOnce sync.Once
	file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDescData = file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDesc
)

func file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDescGZIP() []byte {
	file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDescOnce.Do(func() {
		file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDescData)
	})
	return file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDescData
}

var file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_goTypes = []interface{}{
	(*X509SVIDWithKey)(nil),                // 0: spire.api.agent.delegatedidentity.v1.X509SVIDWithKey
	(*SubscribeToX509SVIDsRequest)(nil),    // 1: spire.api.agent.delegatedidentity.v1.SubscribeToX509SVIDsRequest
	(*SubscribeToX509SVIDsResponse)(nil),   // 2: spire.api.agent.delegatedidentity.v1.SubscribeToX509SVIDsResponse
	(*SubscribeToX509BundlesRequest)(nil),  // 3: spire.api.agent.delegatedidentity.v1.SubscribeToX509BundlesRequest
	(*SubscribeToX509BundlesResponse)(nil), // 4: spire.api.agent.delegatedidentity.v1.SubscribeToX509BundlesResponse
	(*FetchJWTSVIDsRequest)(nil),           // 5: spire.api.agent.delegatedidentity.v1.FetchJWTSVIDsRequest
	(*FetchJWTSVIDsResponse)(nil),          // 6: spire.api.agent.delegatedidentity.v1.FetchJWTSVIDsResponse
	(*SubscribeToJWTBundlesRequest)(nil),   // 7: spire.api.agent.delegatedidentity.v1.SubscribeToJWTBundlesRequest
	(*SubscribeToJWTBundlesResponse)(nil),  // 8: spire.api.agent.delegatedidentity.v1.SubscribeToJWTBundlesResponse
	nil,                                    // 9: spire.api.agent.delegatedidentity.v1.SubscribeToX509BundlesResponse.CaCertificatesEntry
	nil,                                    // 10: spire.api.agent.delegatedidentity.v1.SubscribeToJWTBundlesResponse.BundlesEntry
	(*types.X509SVID)(nil),                 // 11: spire.api.types.X509SVID
	(*types.Selector)(nil),                 // 12: spire.api.types.Selector
	(*types.JWTSVID)(nil),                  // 13: spire.api.types.JWTSVID
}
var file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_depIdxs = []int32{
	11, // 0: spire.api.agent.delegatedidentity.v1.X509SVIDWithKey.x509_svid:type_name -> spire.api.types.X509SVID
	12, // 1: spire.api.agent.delegatedidentity.v1.SubscribeToX509SVIDsRequest.selectors:type_name -> spire.api.types.Selector
	0,  // 2: spire.api.agent.delegatedidentity.v1.SubscribeToX509SVIDsResponse.x509_svids:type_name -> spire.api.agent.delegatedidentity.v1.X509SVIDWithKey
	9,  // 3: spire.api.agent.delegatedidentity.v1.SubscribeToX509BundlesResponse.ca_certificates:type_name -> spire.api.agent.delegatedidentity.v1.SubscribeToX509BundlesResponse.CaCertificatesEntry
	12, // 4: spire.api.agent.delegatedidentity.v1.FetchJWTSVIDsRequest.selectors:type_name -> spire.api.types.Selector
	13, // 5: spire.api.agent.delegatedidentity.v1.FetchJWTSVIDsResponse.svids:type_name -> spire.api.types.JWTSVID
	10, // 6: spire.api.agent.delegatedidentity.v1.SubscribeToJWTBundlesResponse.bundles:type_name -> spire.api.agent.delegatedidentity.v1.SubscribeToJWTBundlesResponse.BundlesEntry
	1,  // 7: spire.api.agent.delegatedidentity.v1.DelegatedIdentity.SubscribeToX509SVIDs:input_type -> spire.api.agent.delegatedidentity.v1.SubscribeToX509SVIDsRequest
	3,  // 8: spire.api.agent.delegatedidentity.v1.DelegatedIdentity.SubscribeToX509Bundles:input_type -> spire.api.agent.delegatedidentity.v1.SubscribeToX509BundlesRequest
	5,  // 9: spire.api.agent.delegatedidentity.v1.DelegatedIdentity.FetchJWTSVIDs:input_type -> spire.api.agent.delegatedidentity.v1.FetchJWTSVIDsRequest
	7,  // 10: spire.api.agent.delegatedidentity.v1.DelegatedIdentity.SubscribeToJWTBundles:input_type -> spire.api.agent.delegatedidentity.v1.SubscribeToJWTBundlesRequest
	2,  // 11: spire.api.agent.delegatedidentity.v1.DelegatedIdentity.SubscribeToX509SVIDs:output_type -> spire.api.agent.delegatedidentity.v1.SubscribeToX509SVIDsResponse
	4,  // 12: spire.api.agent.delegatedidentity.v1.DelegatedIdentity.SubscribeToX509Bundles:output_type -> spire.api.agent.delegatedidentity.v1.SubscribeToX509BundlesResponse
	6,  // 13: spire.api.agent.delegatedidentity.v1.DelegatedIdentity.FetchJWTSVIDs:output_type -> spire.api.agent.delegatedidentity.v1.FetchJWTSVIDsResponse
	8,  // 14: spire.api.agent.delegatedidentity.v1.DelegatedIdentity.SubscribeToJWTBundles:output_type -> spire.api.agent.delegatedidentity.v1.SubscribeToJWTBundlesResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_init() }
func file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_init() {
	if File_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*X509SVIDWithKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeToX509SVIDsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeToX509SVIDsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeToX509BundlesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeToX509BundlesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchJWTSVIDsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchJWTSVIDsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeToJWTBundlesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeToJWTBundlesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_goTypes,
		DependencyIndexes: file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_depIdxs,
		MessageInfos:      file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_msgTypes,
	}.Build()
	File_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto = out.File
	file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_rawDesc = nil
	file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_goTypes = nil
	file_spire_api_agent_delegatedidentity_v1_delegatedidentity_proto_depIdxs = nil
}
