// Code generated by protoc-gen-go. DO NOT EDIT.
// source: apigrpc/apigrpc.proto

package apigrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
	api "spaceship/api"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("apigrpc/apigrpc.proto", fileDescriptor_84e2d31978c605c7) }

var fileDescriptor_84e2d31978c605c7 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xbf, 0x4a, 0x03, 0x41,
	0x10, 0x87, 0xb9, 0x08, 0x8a, 0x87, 0x36, 0x07, 0x1a, 0x3c, 0xad, 0x82, 0x08, 0x8a, 0xd9, 0x45,
	0xed, 0xec, 0x92, 0x22, 0x75, 0x60, 0x3b, 0xbb, 0xc9, 0xb2, 0xd9, 0x1b, 0x30, 0x3b, 0xc3, 0xee,
	0x9c, 0x62, 0x6b, 0x69, 0x6b, 0x63, 0x6d, 0xe9, 0xeb, 0xf8, 0x0a, 0x3e, 0x88, 0xe4, 0xee, 0x48,
	0x8c, 0x68, 0x35, 0x2c, 0xbf, 0x6f, 0xbf, 0xf9, 0x93, 0x1f, 0x00, 0xa3, 0x8f, 0x6c, 0x75, 0x57,
	0x15, 0x47, 0x12, 0x2a, 0xf6, 0x13, 0x83, 0x75, 0xa9, 0x42, 0x56, 0xc0, 0x58, 0x9e, 0x78, 0x22,
	0x7f, 0xef, 0x96, 0x90, 0x86, 0x10, 0x48, 0x40, 0x90, 0x42, 0x6a, 0xe1, 0xf2, 0xb8, 0x4b, 0x9b,
	0xd7, 0xac, 0x9e, 0x6b, 0xb7, 0x60, 0x79, 0xea, 0xc2, 0xcb, 0xa6, 0xd8, 0xa1, 0x77, 0x61, 0x98,
	0x1e, 0xc1, 0x7b, 0x17, 0x35, 0x71, 0xf3, 0xfd, 0x0f, 0x55, 0x7f, 0xd5, 0xb7, 0xed, 0xc5, 0xd8,
	0x06, 0xd7, 0x6f, 0x59, 0xbe, 0x6b, 0x96, 0x99, 0xa9, 0x90, 0x8b, 0x97, 0x2c, 0xef, 0x8f, 0x6a,
	0xa9, 0x5c, 0x10, 0xb4, 0x20, 0x6e, 0x82, 0xc1, 0xbb, 0xc8, 0x11, 0x83, 0x14, 0x67, 0x6a, 0x63,
	0x76, 0xf5, 0x0f, 0x57, 0x1e, 0xfe, 0xe2, 0x8c, 0x4b, 0x09, 0x29, 0x0c, 0xf4, 0xf3, 0xe7, 0xd7,
	0x6b, 0xef, 0x7c, 0x70, 0xaa, 0x1f, 0xae, 0x34, 0x58, 0x4b, 0x75, 0x10, 0x0d, 0x3f, 0x24, 0x7a,
	0xbe, 0xb6, 0xdc, 0x66, 0x17, 0xe3, 0x49, 0x7e, 0x24, 0x51, 0x59, 0x5a, 0x28, 0x60, 0x4e, 0x9b,
	0xd6, 0xf1, 0xde, 0x6a, 0xe8, 0x11, 0xe3, 0x34, 0xbb, 0xdb, 0xe9, 0xee, 0xfc, 0xde, 0xdb, 0x32,
	0x66, 0xfa, 0xd1, 0x5b, 0x2f, 0x35, 0xdb, 0x6e, 0x36, 0xbd, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0x49, 0xdb, 0x7e, 0x4d, 0x93, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SpaceShipClient is the client API for SpaceShip service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpaceShipClient interface {
	AuthenticateFingerprint(ctx context.Context, in *api.AuthenticateFingerprint, opts ...grpc.CallOption) (*api.Session, error)
}

type spaceShipClient struct {
	cc *grpc.ClientConn
}

func NewSpaceShipClient(cc *grpc.ClientConn) SpaceShipClient {
	return &spaceShipClient{cc}
}

func (c *spaceShipClient) AuthenticateFingerprint(ctx context.Context, in *api.AuthenticateFingerprint, opts ...grpc.CallOption) (*api.Session, error) {
	out := new(api.Session)
	err := c.cc.Invoke(ctx, "/spaceship.api.SpaceShip/AuthenticateFingerprint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpaceShipServer is the server API for SpaceShip service.
type SpaceShipServer interface {
	AuthenticateFingerprint(context.Context, *api.AuthenticateFingerprint) (*api.Session, error)
}

func RegisterSpaceShipServer(s *grpc.Server, srv SpaceShipServer) {
	s.RegisterService(&_SpaceShip_serviceDesc, srv)
}

func _SpaceShip_AuthenticateFingerprint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.AuthenticateFingerprint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceShipServer).AuthenticateFingerprint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spaceship.api.SpaceShip/AuthenticateFingerprint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceShipServer).AuthenticateFingerprint(ctx, req.(*api.AuthenticateFingerprint))
	}
	return interceptor(ctx, in, info, handler)
}

var _SpaceShip_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spaceship.api.SpaceShip",
	HandlerType: (*SpaceShipServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthenticateFingerprint",
			Handler:    _SpaceShip_AuthenticateFingerprint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apigrpc/apigrpc.proto",
}
