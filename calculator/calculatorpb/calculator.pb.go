// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

/*
Package calculatorpb is a generated protocol buffer package.

It is generated from these files:
	calculator/calculatorpb/calculator.proto

It has these top-level messages:
	SumRequest
	SumResponse
	PrimeNumberDecompositionRequest
	PrimeNumberDecompositionResponse
	ComputeAverageRequest
	ComputeAverageResponse
	FindMaximumRequest
	FindMaximumResponse
	SquareRootRequest
	SquareRootResponse
*/
package calculatorpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type SumRequest struct {
	FirstNumber  int32 `protobuf:"varint,1,opt,name=first_number,json=firstNumber" json:"first_number,omitempty"`
	SecondNumber int32 `protobuf:"varint,2,opt,name=second_number,json=secondNumber" json:"second_number,omitempty"`
}

func (m *SumRequest) Reset()                    { *m = SumRequest{} }
func (m *SumRequest) String() string            { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()               {}
func (*SumRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SumRequest) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *SumRequest) GetSecondNumber() int32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type SumResponse struct {
	SumResult int32 `protobuf:"varint,1,opt,name=sum_result,json=sumResult" json:"sum_result,omitempty"`
}

func (m *SumResponse) Reset()                    { *m = SumResponse{} }
func (m *SumResponse) String() string            { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()               {}
func (*SumResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SumResponse) GetSumResult() int32 {
	if m != nil {
		return m.SumResult
	}
	return 0
}

type PrimeNumberDecompositionRequest struct {
	Number int64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *PrimeNumberDecompositionRequest) Reset()                    { *m = PrimeNumberDecompositionRequest{} }
func (m *PrimeNumberDecompositionRequest) String() string            { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionRequest) ProtoMessage()               {}
func (*PrimeNumberDecompositionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PrimeNumberDecompositionRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PrimeNumberDecompositionResponse struct {
	PrimeFactor int64 `protobuf:"varint,1,opt,name=prime_factor,json=primeFactor" json:"prime_factor,omitempty"`
}

func (m *PrimeNumberDecompositionResponse) Reset()         { *m = PrimeNumberDecompositionResponse{} }
func (m *PrimeNumberDecompositionResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionResponse) ProtoMessage()    {}
func (*PrimeNumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3}
}

func (m *PrimeNumberDecompositionResponse) GetPrimeFactor() int64 {
	if m != nil {
		return m.PrimeFactor
	}
	return 0
}

type ComputeAverageRequest struct {
	Number int32 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *ComputeAverageRequest) Reset()                    { *m = ComputeAverageRequest{} }
func (m *ComputeAverageRequest) String() string            { return proto.CompactTextString(m) }
func (*ComputeAverageRequest) ProtoMessage()               {}
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ComputeAverageRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type ComputeAverageResponse struct {
	Average float64 `protobuf:"fixed64,1,opt,name=average" json:"average,omitempty"`
}

func (m *ComputeAverageResponse) Reset()                    { *m = ComputeAverageResponse{} }
func (m *ComputeAverageResponse) String() string            { return proto.CompactTextString(m) }
func (*ComputeAverageResponse) ProtoMessage()               {}
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ComputeAverageResponse) GetAverage() float64 {
	if m != nil {
		return m.Average
	}
	return 0
}

type FindMaximumRequest struct {
	Number int32 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *FindMaximumRequest) Reset()                    { *m = FindMaximumRequest{} }
func (m *FindMaximumRequest) String() string            { return proto.CompactTextString(m) }
func (*FindMaximumRequest) ProtoMessage()               {}
func (*FindMaximumRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FindMaximumRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type FindMaximumResponse struct {
	Maximun int32 `protobuf:"varint,1,opt,name=maximun" json:"maximun,omitempty"`
}

func (m *FindMaximumResponse) Reset()                    { *m = FindMaximumResponse{} }
func (m *FindMaximumResponse) String() string            { return proto.CompactTextString(m) }
func (*FindMaximumResponse) ProtoMessage()               {}
func (*FindMaximumResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FindMaximumResponse) GetMaximun() int32 {
	if m != nil {
		return m.Maximun
	}
	return 0
}

type SquareRootRequest struct {
	Number int32 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *SquareRootRequest) Reset()                    { *m = SquareRootRequest{} }
func (m *SquareRootRequest) String() string            { return proto.CompactTextString(m) }
func (*SquareRootRequest) ProtoMessage()               {}
func (*SquareRootRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SquareRootRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SquareRootResponse struct {
	NumberRoot float64 `protobuf:"fixed64,1,opt,name=number_root,json=numberRoot" json:"number_root,omitempty"`
}

func (m *SquareRootResponse) Reset()                    { *m = SquareRootResponse{} }
func (m *SquareRootResponse) String() string            { return proto.CompactTextString(m) }
func (*SquareRootResponse) ProtoMessage()               {}
func (*SquareRootResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SquareRootResponse) GetNumberRoot() float64 {
	if m != nil {
		return m.NumberRoot
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*PrimeNumberDecompositionRequest)(nil), "calculator.PrimeNumberDecompositionRequest")
	proto.RegisterType((*PrimeNumberDecompositionResponse)(nil), "calculator.PrimeNumberDecompositionResponse")
	proto.RegisterType((*ComputeAverageRequest)(nil), "calculator.ComputeAverageRequest")
	proto.RegisterType((*ComputeAverageResponse)(nil), "calculator.ComputeAverageResponse")
	proto.RegisterType((*FindMaximumRequest)(nil), "calculator.FindMaximumRequest")
	proto.RegisterType((*FindMaximumResponse)(nil), "calculator.FindMaximumResponse")
	proto.RegisterType((*SquareRootRequest)(nil), "calculator.SquareRootRequest")
	proto.RegisterType((*SquareRootResponse)(nil), "calculator.SquareRootResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CalculatorService service

type CalculatorServiceClient interface {
	// unary
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	// Server stream
	PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error)
	// Client stream
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error)
	// BiDi stream
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error)
	// Error handling
	// this RPC will throw and exception if a the sent number is negative
	// the error being sent is of type INVALID_ARGUMENT
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := grpc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CalculatorService_serviceDesc.Streams[0], c.cc, "/calculator.CalculatorService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumberDecompositionResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberDecompositionClient) Recv() (*PrimeNumberDecompositionResponse, error) {
	m := new(PrimeNumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CalculatorService_serviceDesc.Streams[1], c.cc, "/calculator.CalculatorService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceComputeAverageClient{stream}
	return x, nil
}

type CalculatorService_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CalculatorService_serviceDesc.Streams[2], c.cc, "/calculator.CalculatorService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFindMaximumClient{stream}
	return x, nil
}

type CalculatorService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type calculatorServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := grpc.Invoke(ctx, "/calculator.CalculatorService/SquareRoot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CalculatorService service

type CalculatorServiceServer interface {
	// unary
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	// Server stream
	PrimeNumberDecomposition(*PrimeNumberDecompositionRequest, CalculatorService_PrimeNumberDecompositionServer) error
	// Client stream
	ComputeAverage(CalculatorService_ComputeAverageServer) error
	// BiDi stream
	FindMaximum(CalculatorService_FindMaximumServer) error
	// Error handling
	// this RPC will throw and exception if a the sent number is negative
	// the error being sent is of type INVALID_ARGUMENT
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumberDecomposition(m, &calculatorServicePrimeNumberDecompositionServer{stream})
}

type CalculatorService_PrimeNumberDecompositionServer interface {
	Send(*PrimeNumberDecompositionResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberDecompositionServer) Send(m *PrimeNumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAverage(&calculatorServiceComputeAverageServer{stream})
}

type CalculatorService_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).FindMaximum(&calculatorServiceFindMaximumServer{stream})
}

type CalculatorService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type calculatorServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _CalculatorService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _CalculatorService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalculatorService_ComputeAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _CalculatorService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}

func init() { proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5b, 0x6b, 0xd4, 0x40,
	0x14, 0x36, 0x2e, 0xad, 0x78, 0x12, 0x0b, 0x3d, 0xe2, 0xba, 0x04, 0xda, 0xed, 0x8e, 0x2f, 0x81,
	0x2e, 0xdd, 0x52, 0x11, 0xf4, 0x51, 0xab, 0x7d, 0x11, 0x45, 0x12, 0x9f, 0xf4, 0x21, 0x64, 0xd3,
	0xa9, 0x04, 0x32, 0x99, 0x74, 0x2e, 0x8b, 0xfe, 0x41, 0x7f, 0x97, 0xec, 0x4c, 0xb2, 0x99, 0xbd,
	0xc4, 0xed, 0x5b, 0xce, 0x37, 0xdf, 0x65, 0xc8, 0xf9, 0x18, 0x88, 0xf2, 0xac, 0xcc, 0x75, 0x99,
	0x29, 0x2e, 0x66, 0xdd, 0x67, 0x3d, 0x77, 0x86, 0x8b, 0x5a, 0x70, 0xc5, 0x11, 0x3a, 0x84, 0x7c,
	0x07, 0x48, 0x34, 0x8b, 0xe9, 0xbd, 0xa6, 0x52, 0xe1, 0x04, 0x82, 0xbb, 0x42, 0x48, 0x95, 0x56,
	0x9a, 0xcd, 0xa9, 0x18, 0x79, 0x67, 0x5e, 0x74, 0x10, 0xfb, 0x06, 0xfb, 0x6a, 0x20, 0x7c, 0x05,
	0xcf, 0x24, 0xcd, 0x79, 0x75, 0xdb, 0x72, 0x1e, 0x1b, 0x4e, 0x60, 0x41, 0x4b, 0x22, 0x53, 0xf0,
	0x8d, 0xab, 0xac, 0x79, 0x25, 0x29, 0x9e, 0x00, 0x48, 0xcd, 0x52, 0x41, 0xa5, 0x2e, 0x55, 0x63,
	0xfa, 0x54, 0x1a, 0x82, 0x2e, 0x15, 0x79, 0x07, 0xe3, 0x6f, 0xa2, 0x60, 0xd4, 0x8a, 0x3f, 0xd2,
	0x9c, 0xb3, 0x9a, 0xcb, 0x42, 0x15, 0xbc, 0x6a, 0x2f, 0x36, 0x84, 0x43, 0xe7, 0x4a, 0x83, 0xb8,
	0x99, 0xc8, 0x27, 0x38, 0xeb, 0x97, 0x36, 0xe9, 0x13, 0x08, 0xea, 0x25, 0x27, 0xbd, 0xcb, 0x72,
	0xc5, 0x5b, 0x07, 0xdf, 0x60, 0x37, 0x06, 0x22, 0x33, 0x78, 0x71, 0xcd, 0x59, 0xad, 0x15, 0x7d,
	0xbf, 0xa0, 0x22, 0xfb, 0x45, 0x77, 0xe7, 0x1e, 0xac, 0x72, 0xaf, 0x60, 0xb8, 0x29, 0x68, 0xd2,
	0x46, 0xf0, 0x24, 0xb3, 0x90, 0x91, 0x78, 0x71, 0x3b, 0x92, 0x29, 0xe0, 0x4d, 0x51, 0xdd, 0x7e,
	0xc9, 0x7e, 0x17, 0xac, 0xfb, 0xe5, 0x7d, 0x09, 0x33, 0x78, 0xbe, 0xc6, 0xee, 0xec, 0x99, 0x81,
	0xaa, 0x86, 0xdf, 0x8e, 0xe4, 0x1c, 0x8e, 0x93, 0x7b, 0x9d, 0x09, 0x1a, 0x73, 0xae, 0xf6, 0xb9,
	0xbf, 0x01, 0x74, 0xc9, 0x8d, 0xf9, 0x18, 0x7c, 0x7b, 0x9e, 0x0a, 0xce, 0x55, 0x73, 0x7f, 0xb0,
	0xd0, 0x92, 0x78, 0xf5, 0x77, 0x00, 0xc7, 0xd7, 0xab, 0xf2, 0x24, 0x54, 0x2c, 0x8a, 0x9c, 0xe2,
	0x5b, 0x18, 0x24, 0x9a, 0xe1, 0xf0, 0xc2, 0x69, 0x5a, 0x57, 0xaa, 0xf0, 0xe5, 0x16, 0x6e, 0xe3,
	0xc8, 0x23, 0xfc, 0x03, 0xa3, 0xbe, 0xf5, 0xe1, 0xb9, 0x2b, 0xdb, 0xd3, 0x8f, 0x70, 0xfa, 0x30,
	0x72, 0x1b, 0x7c, 0xe9, 0xe1, 0x4f, 0x38, 0x5a, 0xdf, 0x20, 0x4e, 0x5c, 0x8f, 0x9d, 0x75, 0x08,
	0xc9, 0xff, 0x28, 0xad, 0x79, 0xe4, 0x61, 0x0c, 0xbe, 0xb3, 0x3c, 0x3c, 0x75, 0x65, 0xdb, 0x1d,
	0x08, 0xc7, 0xbd, 0xe7, 0xd6, 0x33, 0xf2, 0x2e, 0x3d, 0xfc, 0x0c, 0xd0, 0xad, 0x0c, 0x4f, 0xd6,
	0x7e, 0xea, 0xe6, 0xde, 0xc3, 0xd3, 0xbe, 0x63, 0x6b, 0xf8, 0xe1, 0xe8, 0x47, 0xe0, 0xbe, 0x11,
	0xf3, 0x43, 0xf3, 0x32, 0xbc, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x73, 0x66, 0x37, 0x8f, 0x45,
	0x04, 0x00, 0x00,
}
