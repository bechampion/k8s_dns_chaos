// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: dns.proto

package pb

import (
	grpc "google.golang.org/grpc"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	context "golang.org/x/net/context"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var _DNS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DNS",
	HandlerType: (*DNSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetDNSChaos",
			Handler:    _DNS_SetDNSChaos_Handler,
		},
		{
			MethodName: "CancelDNSChaos",
			Handler:    _DNS_CancelDNSChaos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dns.proto",
}

 type DNSServer interface {                                                                                                                                                            
	SetDNSChaos(context.Context, *SetDNSChaosRequest) (*DNSChaosResponse, error)
	CancelDNSChaos(context.Context, *CancelDNSChaosRequest) (*DNSChaosResponse, error)
}
type SetDNSChaosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pods []*Pod `protobuf:"bytes,2,rep,name=pods,proto3" json:"pods,omitempty"`
	// action means the chaos action, values can be "random" or "error"
	//
	//	"random": return random IP for DNS request
	//	"error":  return error for DNS request
	Action string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	// scope means the chaos scope, values can be "inner", "outer" or "all":
	//
	//	"inner": chaos only works on the inner host in Kubernetes cluster
	//	"outer": chaos only works on the outer host of Kubernetes cluster
	//	"all":   chaos works on all host
	Scope        string   `protobuf:"bytes,4,opt,name=scope,proto3" json:"scope,omitempty"`
	Selector     string   `protobuf:"bytes,5,opt,name=selector,proto3" json:"selector,omitempty"`
	Patterns     []string `protobuf:"bytes,6,rep,name=patterns,proto3" json:"patterns,omitempty"`
	FixedAddress string   `protobuf:"bytes,7,opt,name=fixedaddress,proto3" json:"fixedaddress,omitempty"`
}


type DNSClient interface {
	SetDNSChaos(ctx context.Context, in *SetDNSChaosRequest, opts ...grpc.CallOption) (*DNSChaosResponse, error)
	CancelDNSChaos(ctx context.Context, in *CancelDNSChaosRequest, opts ...grpc.CallOption) (*DNSChaosResponse, error)
}

type dNSClient struct {
	cc *grpc.ClientConn
}

func (c *dNSClient) SetDNSChaos(ctx context.Context, in *SetDNSChaosRequest, opts ...grpc.CallOption) (*DNSChaosResponse, error) {
	out := new(DNSChaosResponse)
	err := c.cc.Invoke(ctx, "/pb.DNS/SetDNSChaos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSClient) CancelDNSChaos(ctx context.Context, in *CancelDNSChaosRequest, opts ...grpc.CallOption) (*DNSChaosResponse, error) {
	out := new(DNSChaosResponse)
	err := c.cc.Invoke(ctx, "/pb.DNS/CancelDNSChaos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func NewDNSClient(cc *grpc.ClientConn) DNSClient {
	return &dNSClient{cc}
}

func RegisterDNSServer(s *grpc.Server, srv DNSServer) {
	s.RegisterService(&_DNS_serviceDesc, srv)
}
func _DNS_SetDNSChaos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDNSChaosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSServer).SetDNSChaos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNS/SetDNSChaos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSServer).SetDNSChaos(ctx, req.(*SetDNSChaosRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _DNS_CancelDNSChaos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelDNSChaosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSServer).CancelDNSChaos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNS/CancelDNSChaos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSServer).CancelDNSChaos(ctx, req.(*CancelDNSChaosRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (x *SetDNSChaosRequest) Reset() {
	*x = SetDNSChaosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDNSChaosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDNSChaosRequest) ProtoMessage() {}

func (x *SetDNSChaosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDNSChaosRequest.ProtoReflect.Descriptor instead.
func (*SetDNSChaosRequest) Descriptor() ([]byte, []int) {
	return file_dns_proto_rawDescGZIP(), []int{0}
}

func (x *SetDNSChaosRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SetDNSChaosRequest) GetPods() []*Pod {
	if x != nil {
		return x.Pods
	}
	return nil
}

func (x *SetDNSChaosRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *SetDNSChaosRequest) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *SetDNSChaosRequest) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

func (x *SetDNSChaosRequest) GetPatterns() []string {
	if x != nil {
		return x.Patterns
	}
	return nil
}

func (x *SetDNSChaosRequest) GetFixedaddress() string {
	if x != nil {
		return x.FixedAddress
	}
	return ""
}

type Pod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Pod) Reset() {
	*x = Pod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pod) ProtoMessage() {}

func (x *Pod) ProtoReflect() protoreflect.Message {
	mi := &file_dns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pod.ProtoReflect.Descriptor instead.
func (*Pod) Descriptor() ([]byte, []int) {
	return file_dns_proto_rawDescGZIP(), []int{1}
}

func (x *Pod) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Pod) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CancelDNSChaosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CancelDNSChaosRequest) Reset() {
	*x = CancelDNSChaosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelDNSChaosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelDNSChaosRequest) ProtoMessage() {}

func (x *CancelDNSChaosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dns_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelDNSChaosRequest.ProtoReflect.Descriptor instead.
func (*CancelDNSChaosRequest) Descriptor() ([]byte, []int) {
	return file_dns_proto_rawDescGZIP(), []int{2}
}

func (x *CancelDNSChaosRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DNSChaosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DNSChaosResponse) Reset() {
	*x = DNSChaosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSChaosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSChaosResponse) ProtoMessage() {}

func (x *DNSChaosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dns_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSChaosResponse.ProtoReflect.Descriptor instead.
func (*DNSChaosResponse) Descriptor() ([]byte, []int) {
	return file_dns_proto_rawDescGZIP(), []int{3}
}

func (x *DNSChaosResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *DNSChaosResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_dns_proto protoreflect.FileDescriptor

var file_dns_proto_rawDesc = []byte{
	0x0a, 0x09, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0xcf, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43, 0x68, 0x61, 0x6f, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x6f,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f,
	0x64, 0x52, 0x04, 0x70, 0x6f, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x12, 0x22, 0x0a,
	0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x37, 0x0a, 0x03, 0x50, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x15, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x44, 0x4e, 0x53, 0x43, 0x68, 0x61, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x10, 0x44, 0x4e, 0x53, 0x43, 0x68,
	0x61, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x89, 0x01, 0x0a, 0x03, 0x44, 0x4e, 0x53, 0x12, 0x3d, 0x0a,
	0x0b, 0x53, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43, 0x68, 0x61, 0x6f, 0x73, 0x12, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43, 0x68, 0x61, 0x6f, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x4e, 0x53, 0x43, 0x68, 0x61,
	0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x44, 0x4e, 0x53, 0x43, 0x68, 0x61, 0x6f, 0x73, 0x12, 0x19,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x44, 0x4e, 0x53, 0x43, 0x68, 0x61,
	0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x44,
	0x4e, 0x53, 0x43, 0x68, 0x61, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dns_proto_rawDescOnce sync.Once
	file_dns_proto_rawDescData = file_dns_proto_rawDesc
)

func file_dns_proto_rawDescGZIP() []byte {
	file_dns_proto_rawDescOnce.Do(func() {
		file_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_dns_proto_rawDescData)
	})
	return file_dns_proto_rawDescData
}

var file_dns_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dns_proto_goTypes = []interface{}{
	(*SetDNSChaosRequest)(nil),    // 0: pb.SetDNSChaosRequest
	(*Pod)(nil),                   // 1: pb.Pod
	(*CancelDNSChaosRequest)(nil), // 2: pb.CancelDNSChaosRequest
	(*DNSChaosResponse)(nil),      // 3: pb.DNSChaosResponse
}
var file_dns_proto_depIdxs = []int32{
	1, // 0: pb.SetDNSChaosRequest.pods:type_name -> pb.Pod
	0, // 1: pb.DNS.SetDNSChaos:input_type -> pb.SetDNSChaosRequest
	2, // 2: pb.DNS.CancelDNSChaos:input_type -> pb.CancelDNSChaosRequest
	3, // 3: pb.DNS.SetDNSChaos:output_type -> pb.DNSChaosResponse
	3, // 4: pb.DNS.CancelDNSChaos:output_type -> pb.DNSChaosResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dns_proto_init() }
func file_dns_proto_init() {
	if File_dns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDNSChaosRequest); i {
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
		file_dns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pod); i {
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
		file_dns_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelDNSChaosRequest); i {
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
		file_dns_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSChaosResponse); i {
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
			RawDescriptor: file_dns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dns_proto_goTypes,
		DependencyIndexes: file_dns_proto_depIdxs,
		MessageInfos:      file_dns_proto_msgTypes,
	}.Build()
	File_dns_proto = out.File
	file_dns_proto_rawDesc = nil
	file_dns_proto_goTypes = nil
	file_dns_proto_depIdxs = nil
}
