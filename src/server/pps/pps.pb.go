// Code generated by protoc-gen-go.
// source: server/pps/pps.proto
// DO NOT EDIT!

/*
Package pps is a generated protocol buffer package.

It is generated from these files:
	server/pps/pps.proto

It has these top-level messages:
	StartJobRequest
	StartJobResponse
	FinishJobRequest
*/
package pps

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"
import fuse "github.com/pachyderm/pachyderm/src/server/pfs/fuse"
import _ "github.com/pachyderm/pachyderm/src/client/pfs"
import pachyderm_pps "github.com/pachyderm/pachyderm/src/client/pps"

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

type StartJobRequest struct {
	Job *pachyderm_pps.Job `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
}

func (m *StartJobRequest) Reset()                    { *m = StartJobRequest{} }
func (m *StartJobRequest) String() string            { return proto.CompactTextString(m) }
func (*StartJobRequest) ProtoMessage()               {}
func (*StartJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StartJobRequest) GetJob() *pachyderm_pps.Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type StartJobResponse struct {
	Transform    *pachyderm_pps.Transform `protobuf:"bytes,1,opt,name=transform" json:"transform,omitempty"`
	CommitMounts []*fuse.CommitMount      `protobuf:"bytes,2,rep,name=commit_mounts,json=commitMounts" json:"commit_mounts,omitempty"`
	PodIndex     uint64                   `protobuf:"varint,3,opt,name=pod_index,json=podIndex" json:"pod_index,omitempty"`
}

func (m *StartJobResponse) Reset()                    { *m = StartJobResponse{} }
func (m *StartJobResponse) String() string            { return proto.CompactTextString(m) }
func (*StartJobResponse) ProtoMessage()               {}
func (*StartJobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StartJobResponse) GetTransform() *pachyderm_pps.Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *StartJobResponse) GetCommitMounts() []*fuse.CommitMount {
	if m != nil {
		return m.CommitMounts
	}
	return nil
}

type FinishJobRequest struct {
	Job      *pachyderm_pps.Job `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	Success  bool               `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
	PodIndex uint64             `protobuf:"varint,3,opt,name=pod_index,json=podIndex" json:"pod_index,omitempty"`
}

func (m *FinishJobRequest) Reset()                    { *m = FinishJobRequest{} }
func (m *FinishJobRequest) String() string            { return proto.CompactTextString(m) }
func (*FinishJobRequest) ProtoMessage()               {}
func (*FinishJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FinishJobRequest) GetJob() *pachyderm_pps.Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func init() {
	proto.RegisterType((*StartJobRequest)(nil), "pachyderm.pps.StartJobRequest")
	proto.RegisterType((*StartJobResponse)(nil), "pachyderm.pps.StartJobResponse")
	proto.RegisterType((*FinishJobRequest)(nil), "pachyderm.pps.FinishJobRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for InternalJobAPI service

type InternalJobAPIClient interface {
	StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error)
	FinishJob(ctx context.Context, in *FinishJobRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type internalJobAPIClient struct {
	cc *grpc.ClientConn
}

func NewInternalJobAPIClient(cc *grpc.ClientConn) InternalJobAPIClient {
	return &internalJobAPIClient{cc}
}

func (c *internalJobAPIClient) StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error) {
	out := new(StartJobResponse)
	err := grpc.Invoke(ctx, "/pachyderm.pps.InternalJobAPI/StartJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalJobAPIClient) FinishJob(ctx context.Context, in *FinishJobRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.InternalJobAPI/FinishJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InternalJobAPI service

type InternalJobAPIServer interface {
	StartJob(context.Context, *StartJobRequest) (*StartJobResponse, error)
	FinishJob(context.Context, *FinishJobRequest) (*google_protobuf.Empty, error)
}

func RegisterInternalJobAPIServer(s *grpc.Server, srv InternalJobAPIServer) {
	s.RegisterService(&_InternalJobAPI_serviceDesc, srv)
}

func _InternalJobAPI_StartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalJobAPIServer).StartJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.InternalJobAPI/StartJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalJobAPIServer).StartJob(ctx, req.(*StartJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalJobAPI_FinishJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalJobAPIServer).FinishJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.InternalJobAPI/FinishJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalJobAPIServer).FinishJob(ctx, req.(*FinishJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InternalJobAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.InternalJobAPI",
	HandlerType: (*InternalJobAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartJob",
			Handler:    _InternalJobAPI_StartJob_Handler,
		},
		{
			MethodName: "FinishJob",
			Handler:    _InternalJobAPI_FinishJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("server/pps/pps.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4b, 0xeb, 0x40,
	0x10, 0xc7, 0x9b, 0x97, 0xf7, 0x9e, 0xed, 0xd6, 0x6a, 0x5d, 0x8a, 0x84, 0x14, 0xb4, 0x04, 0x0f,
	0x39, 0x6d, 0xa0, 0x42, 0x3d, 0xab, 0x58, 0x68, 0xa1, 0x20, 0xd1, 0x93, 0x97, 0x92, 0x1f, 0x9b,
	0x36, 0xd2, 0xec, 0xac, 0x3b, 0x1b, 0xb1, 0x7f, 0x8d, 0x47, 0xff, 0x4d, 0x49, 0xda, 0x58, 0x0d,
	0x28, 0x78, 0xd8, 0x85, 0xd9, 0xcf, 0x77, 0x67, 0xe6, 0x3b, 0x0c, 0xe9, 0x21, 0x57, 0xcf, 0x5c,
	0x79, 0x52, 0x62, 0x71, 0x98, 0x54, 0xa0, 0x81, 0x76, 0x64, 0x10, 0x2d, 0xd7, 0x31, 0x57, 0x19,
	0x93, 0x12, 0xed, 0xfe, 0x02, 0x60, 0xb1, 0xe2, 0x5e, 0x09, 0xc3, 0x3c, 0xf1, 0x78, 0x26, 0xf5,
	0x7a, 0xa3, 0xb5, 0xed, 0x2a, 0x43, 0x82, 0x5e, 0x92, 0x23, 0x2f, 0xaf, 0x2d, 0xeb, 0x45, 0xab,
	0x94, 0x0b, 0x5d, 0x32, 0x99, 0x60, 0xfd, 0xf5, 0x73, 0x4d, 0xe7, 0x82, 0x1c, 0xde, 0xe9, 0x40,
	0xe9, 0x29, 0x84, 0x3e, 0x7f, 0xca, 0x39, 0x6a, 0x7a, 0x46, 0xcc, 0x47, 0x08, 0x2d, 0x63, 0x60,
	0xb8, 0xed, 0x21, 0x65, 0x5f, 0x9a, 0x62, 0x85, 0xae, 0xc0, 0xce, 0xab, 0x41, 0xba, 0xbb, 0x9f,
	0x28, 0x41, 0x20, 0xa7, 0x23, 0xd2, 0xd2, 0x2a, 0x10, 0x98, 0x80, 0xca, 0xb6, 0x09, 0xac, 0x5a,
	0x82, 0xfb, 0x8a, 0xfb, 0x3b, 0x29, 0x1d, 0x91, 0x4e, 0x04, 0x59, 0x96, 0xea, 0x79, 0x06, 0xb9,
	0xd0, 0x68, 0xfd, 0x19, 0x98, 0x6e, 0x7b, 0x78, 0xc4, 0x4a, 0x57, 0xd7, 0x25, 0x9a, 0x15, 0xc4,
	0xdf, 0x8f, 0x76, 0x01, 0xd2, 0x3e, 0x69, 0x49, 0x88, 0xe7, 0xa9, 0x88, 0xf9, 0x8b, 0x65, 0x0e,
	0x0c, 0xf7, 0xaf, 0xdf, 0x94, 0x10, 0x4f, 0x8a, 0xd8, 0x01, 0xd2, 0x1d, 0xa7, 0x22, 0xc5, 0xe5,
	0x6f, 0xbd, 0x51, 0x8b, 0xec, 0x61, 0x1e, 0x45, 0x1c, 0x8b, 0x46, 0x0c, 0xb7, 0xe9, 0x57, 0xe1,
	0x8f, 0x05, 0x87, 0x6f, 0x06, 0x39, 0x98, 0x08, 0xcd, 0x95, 0x08, 0x56, 0x53, 0x08, 0x2f, 0x6f,
	0x27, 0x74, 0x46, 0x9a, 0xd5, 0x90, 0xe8, 0x49, 0xad, 0x5c, 0x6d, 0xee, 0xf6, 0xe9, 0xb7, 0x7c,
	0x33, 0x5d, 0xa7, 0x41, 0xc7, 0xa4, 0xf5, 0x61, 0x89, 0xd6, 0xf5, 0x75, 0xb3, 0xf6, 0x31, 0xdb,
	0x6c, 0x10, 0xab, 0x36, 0x88, 0xdd, 0x14, 0x1b, 0xe4, 0x34, 0xae, 0xfe, 0x3d, 0x98, 0x52, 0x62,
	0xf8, 0xbf, 0x04, 0xe7, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x51, 0x7f, 0xf4, 0x86, 0x8f, 0x02,
	0x00, 0x00,
}
