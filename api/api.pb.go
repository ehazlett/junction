// Code generated by protoc-gen-gogo.
// source: api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
*/
package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for JunctionManager service

type JunctionManagerClient interface {
}

type junctionManagerClient struct {
	cc *grpc.ClientConn
}

func NewJunctionManagerClient(cc *grpc.ClientConn) JunctionManagerClient {
	return &junctionManagerClient{cc}
}

// Server API for JunctionManager service

type JunctionManagerServer interface {
}

func RegisterJunctionManagerServer(s *grpc.Server, srv JunctionManagerServer) {
	s.RegisterService(&_JunctionManager_serviceDesc, srv)
}

var _JunctionManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.JunctionManager",
	HandlerType: (*JunctionManagerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptorApi) }

var fileDescriptorApi = []byte{
	// 76 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf,
	0x07, 0xf3, 0xf5, 0x41, 0x2c, 0x88, 0x94, 0x91, 0x20, 0x17, 0xbf, 0x57, 0x69, 0x5e, 0x72, 0x49,
	0x66, 0x7e, 0x9e, 0x6f, 0x62, 0x5e, 0x62, 0x7a, 0x6a, 0x51, 0x12, 0x1b, 0x58, 0xc6, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x92, 0xf6, 0x2d, 0x1e, 0x41, 0x00, 0x00, 0x00,
}
