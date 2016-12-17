// Code generated by protoc-gen-gogo.
// source: kafka.proto
// DO NOT EDIT!

/*
Package kafkapb is a generated protocol buffer package.

It is generated from these files:
	kafka.proto

It has these top-level messages:
	ConsumerRequest
	ConsumerResponse
	ProducerRequest
	ProducerResponse
*/
package kafkapb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ConsumerRequest struct {
}

func (m *ConsumerRequest) Reset()                    { *m = ConsumerRequest{} }
func (m *ConsumerRequest) String() string            { return proto.CompactTextString(m) }
func (*ConsumerRequest) ProtoMessage()               {}
func (*ConsumerRequest) Descriptor() ([]byte, []int) { return fileDescriptorKafka, []int{0} }

type ConsumerResponse struct {
	ErrMsg string `protobuf:"bytes,1,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
}

func (m *ConsumerResponse) Reset()                    { *m = ConsumerResponse{} }
func (m *ConsumerResponse) String() string            { return proto.CompactTextString(m) }
func (*ConsumerResponse) ProtoMessage()               {}
func (*ConsumerResponse) Descriptor() ([]byte, []int) { return fileDescriptorKafka, []int{1} }

func (m *ConsumerResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type ProducerRequest struct {
	Topic string               `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Value *google_protobuf.Any `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Key   string               `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *ProducerRequest) Reset()                    { *m = ProducerRequest{} }
func (m *ProducerRequest) String() string            { return proto.CompactTextString(m) }
func (*ProducerRequest) ProtoMessage()               {}
func (*ProducerRequest) Descriptor() ([]byte, []int) { return fileDescriptorKafka, []int{2} }

func (m *ProducerRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *ProducerRequest) GetValue() *google_protobuf.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ProducerRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type ProducerResponse struct {
	Partition int32  `protobuf:"varint,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Offset    int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	ErrMsg    string `protobuf:"bytes,3,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
}

func (m *ProducerResponse) Reset()                    { *m = ProducerResponse{} }
func (m *ProducerResponse) String() string            { return proto.CompactTextString(m) }
func (*ProducerResponse) ProtoMessage()               {}
func (*ProducerResponse) Descriptor() ([]byte, []int) { return fileDescriptorKafka, []int{3} }

func (m *ProducerResponse) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *ProducerResponse) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ProducerResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*ConsumerRequest)(nil), "kafka.ConsumerRequest")
	proto.RegisterType((*ConsumerResponse)(nil), "kafka.ConsumerResponse")
	proto.RegisterType((*ProducerRequest)(nil), "kafka.ProducerRequest")
	proto.RegisterType((*ProducerResponse)(nil), "kafka.ProducerResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KafkaService service

type KafkaServiceClient interface {
	Consumer(ctx context.Context, in *ConsumerRequest, opts ...grpc.CallOption) (*ConsumerResponse, error)
	Producer(ctx context.Context, in *ProducerRequest, opts ...grpc.CallOption) (*ProducerResponse, error)
}

type kafkaServiceClient struct {
	cc *grpc.ClientConn
}

func NewKafkaServiceClient(cc *grpc.ClientConn) KafkaServiceClient {
	return &kafkaServiceClient{cc}
}

func (c *kafkaServiceClient) Consumer(ctx context.Context, in *ConsumerRequest, opts ...grpc.CallOption) (*ConsumerResponse, error) {
	out := new(ConsumerResponse)
	err := grpc.Invoke(ctx, "/kafka.KafkaService/Consumer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaServiceClient) Producer(ctx context.Context, in *ProducerRequest, opts ...grpc.CallOption) (*ProducerResponse, error) {
	out := new(ProducerResponse)
	err := grpc.Invoke(ctx, "/kafka.KafkaService/Producer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KafkaService service

type KafkaServiceServer interface {
	Consumer(context.Context, *ConsumerRequest) (*ConsumerResponse, error)
	Producer(context.Context, *ProducerRequest) (*ProducerResponse, error)
}

func RegisterKafkaServiceServer(s *grpc.Server, srv KafkaServiceServer) {
	s.RegisterService(&_KafkaService_serviceDesc, srv)
}

func _KafkaService_Consumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaServiceServer).Consumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kafka.KafkaService/Consumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaServiceServer).Consumer(ctx, req.(*ConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaService_Producer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProducerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaServiceServer).Producer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kafka.KafkaService/Producer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaServiceServer).Producer(ctx, req.(*ProducerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KafkaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kafka.KafkaService",
	HandlerType: (*KafkaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Consumer",
			Handler:    _KafkaService_Consumer_Handler,
		},
		{
			MethodName: "Producer",
			Handler:    _KafkaService_Producer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kafka.proto",
}

func init() { proto.RegisterFile("kafka.proto", fileDescriptorKafka) }

var fileDescriptorKafka = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x1b, 0x43, 0x5a, 0x3b, 0x15, 0x1a, 0x97, 0xd2, 0xc6, 0xe2, 0xa1, 0xec, 0xa9, 0x28,
	0xa4, 0x50, 0xcf, 0x1e, 0xd4, 0xa3, 0x08, 0x12, 0x6f, 0x5e, 0x64, 0x13, 0x27, 0x21, 0xa4, 0xcd,
	0xc6, 0xdd, 0x4d, 0x21, 0xdf, 0xc1, 0x0f, 0x2d, 0xd9, 0xcd, 0x9f, 0x12, 0x6f, 0x3b, 0x6f, 0xe7,
	0x31, 0xef, 0xfd, 0x60, 0x96, 0xb1, 0x38, 0x63, 0x7e, 0x21, 0xb8, 0xe2, 0xc4, 0xd1, 0xc3, 0xfa,
	0x26, 0xe1, 0x3c, 0x39, 0xe0, 0x4e, 0x8b, 0x61, 0x19, 0xef, 0x58, 0x5e, 0x99, 0x0d, 0x7a, 0x0d,
	0xf3, 0x17, 0x9e, 0xcb, 0xf2, 0x88, 0x22, 0xc0, 0x9f, 0x12, 0xa5, 0xa2, 0xf7, 0xe0, 0xf6, 0x92,
	0x2c, 0x78, 0x2e, 0x91, 0xac, 0x60, 0x82, 0x42, 0x7c, 0x1d, 0x65, 0xe2, 0x59, 0x1b, 0x6b, 0x3b,
	0x0d, 0xc6, 0x28, 0xc4, 0x9b, 0x4c, 0x28, 0xc2, 0xfc, 0x5d, 0xf0, 0xef, 0x32, 0xea, 0xfc, 0x64,
	0x01, 0x8e, 0xe2, 0x45, 0x1a, 0x35, 0x9b, 0x66, 0x20, 0x77, 0xe0, 0x9c, 0xd8, 0xa1, 0x44, 0xef,
	0x62, 0x63, 0x6d, 0x67, 0xfb, 0x85, 0x6f, 0x32, 0xf9, 0x6d, 0x26, 0xff, 0x29, 0xaf, 0x02, 0xb3,
	0x42, 0x5c, 0xb0, 0x33, 0xac, 0x3c, 0x5b, 0xfb, 0xeb, 0x27, 0x65, 0xe0, 0xf6, 0x67, 0x9a, 0x4c,
	0xb7, 0x30, 0x2d, 0x98, 0x50, 0xa9, 0x4a, 0x79, 0xae, 0x6f, 0x39, 0x41, 0x2f, 0x90, 0x25, 0x8c,
	0x79, 0x1c, 0x4b, 0x54, 0xfa, 0xa0, 0x1d, 0x34, 0xd3, 0x79, 0x13, 0xfb, 0xbc, 0xc9, 0xfe, 0xd7,
	0x82, 0xab, 0xd7, 0x1a, 0xd7, 0x07, 0x8a, 0x53, 0x1a, 0x21, 0x79, 0x84, 0xcb, 0x96, 0x03, 0x59,
	0xfa, 0x06, 0xeb, 0x80, 0xd5, 0x7a, 0xf5, 0x4f, 0x37, 0xe1, 0xe8, 0xa8, 0xb6, 0xb7, 0x91, 0x3b,
	0xfb, 0x00, 0x55, 0x67, 0x1f, 0x76, 0xa3, 0xa3, 0xe7, 0xe9, 0xe7, 0x44, 0xff, 0x15, 0x61, 0x38,
	0xd6, 0x8c, 0x1e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x29, 0x2c, 0x3c, 0xdb, 0x01, 0x00,
	0x00,
}
