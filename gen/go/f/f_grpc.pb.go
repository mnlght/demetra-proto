// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: f/f.proto

package fv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	UploadSimple(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadSimpleClient, error)
	UploadMultiple(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadMultipleClient, error)
	DownloadMultiple(ctx context.Context, in *FileDownloadMultipleRequest, opts ...grpc.CallOption) (FileService_DownloadMultipleClient, error)
	DownloadSimple(ctx context.Context, in *FileDownloadSimpleRequest, opts ...grpc.CallOption) (FileService_DownloadSimpleClient, error)
	List(ctx context.Context, in *FilesListRequest, opts ...grpc.CallOption) (*FilesListResponse, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) UploadSimple(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadSimpleClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[0], "/f.FileService/UploadSimple", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceUploadSimpleClient{stream}
	return x, nil
}

type FileService_UploadSimpleClient interface {
	Send(*FileUploadRequest) error
	CloseAndRecv() (*FileUploadResponse, error)
	grpc.ClientStream
}

type fileServiceUploadSimpleClient struct {
	grpc.ClientStream
}

func (x *fileServiceUploadSimpleClient) Send(m *FileUploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileServiceUploadSimpleClient) CloseAndRecv() (*FileUploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileUploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileServiceClient) UploadMultiple(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadMultipleClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[1], "/f.FileService/UploadMultiple", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceUploadMultipleClient{stream}
	return x, nil
}

type FileService_UploadMultipleClient interface {
	Send(*FileUploadRequest) error
	CloseAndRecv() (*FileUploadResponse, error)
	grpc.ClientStream
}

type fileServiceUploadMultipleClient struct {
	grpc.ClientStream
}

func (x *fileServiceUploadMultipleClient) Send(m *FileUploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileServiceUploadMultipleClient) CloseAndRecv() (*FileUploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileUploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileServiceClient) DownloadMultiple(ctx context.Context, in *FileDownloadMultipleRequest, opts ...grpc.CallOption) (FileService_DownloadMultipleClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[2], "/f.FileService/DownloadMultiple", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceDownloadMultipleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileService_DownloadMultipleClient interface {
	Recv() (*FileDownloadMultipleResponse, error)
	grpc.ClientStream
}

type fileServiceDownloadMultipleClient struct {
	grpc.ClientStream
}

func (x *fileServiceDownloadMultipleClient) Recv() (*FileDownloadMultipleResponse, error) {
	m := new(FileDownloadMultipleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileServiceClient) DownloadSimple(ctx context.Context, in *FileDownloadSimpleRequest, opts ...grpc.CallOption) (FileService_DownloadSimpleClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[3], "/f.FileService/DownloadSimple", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceDownloadSimpleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileService_DownloadSimpleClient interface {
	Recv() (*FileDownloadSimpleResponse, error)
	grpc.ClientStream
}

type fileServiceDownloadSimpleClient struct {
	grpc.ClientStream
}

func (x *fileServiceDownloadSimpleClient) Recv() (*FileDownloadSimpleResponse, error) {
	m := new(FileDownloadSimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileServiceClient) List(ctx context.Context, in *FilesListRequest, opts ...grpc.CallOption) (*FilesListResponse, error) {
	out := new(FilesListResponse)
	err := c.cc.Invoke(ctx, "/f.FileService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	UploadSimple(FileService_UploadSimpleServer) error
	UploadMultiple(FileService_UploadMultipleServer) error
	DownloadMultiple(*FileDownloadMultipleRequest, FileService_DownloadMultipleServer) error
	DownloadSimple(*FileDownloadSimpleRequest, FileService_DownloadSimpleServer) error
	List(context.Context, *FilesListRequest) (*FilesListResponse, error)
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) UploadSimple(FileService_UploadSimpleServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadSimple not implemented")
}
func (UnimplementedFileServiceServer) UploadMultiple(FileService_UploadMultipleServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadMultiple not implemented")
}
func (UnimplementedFileServiceServer) DownloadMultiple(*FileDownloadMultipleRequest, FileService_DownloadMultipleServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadMultiple not implemented")
}
func (UnimplementedFileServiceServer) DownloadSimple(*FileDownloadSimpleRequest, FileService_DownloadSimpleServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadSimple not implemented")
}
func (UnimplementedFileServiceServer) List(context.Context, *FilesListRequest) (*FilesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_UploadSimple_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).UploadSimple(&fileServiceUploadSimpleServer{stream})
}

type FileService_UploadSimpleServer interface {
	SendAndClose(*FileUploadResponse) error
	Recv() (*FileUploadRequest, error)
	grpc.ServerStream
}

type fileServiceUploadSimpleServer struct {
	grpc.ServerStream
}

func (x *fileServiceUploadSimpleServer) SendAndClose(m *FileUploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileServiceUploadSimpleServer) Recv() (*FileUploadRequest, error) {
	m := new(FileUploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileService_UploadMultiple_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).UploadMultiple(&fileServiceUploadMultipleServer{stream})
}

type FileService_UploadMultipleServer interface {
	SendAndClose(*FileUploadResponse) error
	Recv() (*FileUploadRequest, error)
	grpc.ServerStream
}

type fileServiceUploadMultipleServer struct {
	grpc.ServerStream
}

func (x *fileServiceUploadMultipleServer) SendAndClose(m *FileUploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileServiceUploadMultipleServer) Recv() (*FileUploadRequest, error) {
	m := new(FileUploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileService_DownloadMultiple_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileDownloadMultipleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServiceServer).DownloadMultiple(m, &fileServiceDownloadMultipleServer{stream})
}

type FileService_DownloadMultipleServer interface {
	Send(*FileDownloadMultipleResponse) error
	grpc.ServerStream
}

type fileServiceDownloadMultipleServer struct {
	grpc.ServerStream
}

func (x *fileServiceDownloadMultipleServer) Send(m *FileDownloadMultipleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FileService_DownloadSimple_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileDownloadSimpleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServiceServer).DownloadSimple(m, &fileServiceDownloadSimpleServer{stream})
}

type FileService_DownloadSimpleServer interface {
	Send(*FileDownloadSimpleResponse) error
	grpc.ServerStream
}

type fileServiceDownloadSimpleServer struct {
	grpc.ServerStream
}

func (x *fileServiceDownloadSimpleServer) Send(m *FileDownloadSimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FileService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/f.FileService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).List(ctx, req.(*FilesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "f.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _FileService_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadSimple",
			Handler:       _FileService_UploadSimple_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UploadMultiple",
			Handler:       _FileService_UploadMultiple_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadMultiple",
			Handler:       _FileService_DownloadMultiple_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DownloadSimple",
			Handler:       _FileService_DownloadSimple_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "f/f.proto",
}
