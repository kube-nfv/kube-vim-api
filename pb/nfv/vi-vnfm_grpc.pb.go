// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: vi-vnfm.proto

package nfv

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

const (
	ViVnfm_QueryImages_FullMethodName                                                 = "/kubenvf.kubevim.api.pb.vi_vnfm/QueryImages"
	ViVnfm_QueryImage_FullMethodName                                                  = "/kubenvf.kubevim.api.pb.vi_vnfm/QueryImage"
	ViVnfm_AllocateVirtualisedComputeResource_FullMethodName                          = "/kubenvf.kubevim.api.pb.vi_vnfm/AllocateVirtualisedComputeResource"
	ViVnfm_QueryVirtualisedComputeResource_FullMethodName                             = "/kubenvf.kubevim.api.pb.vi_vnfm/QueryVirtualisedComputeResource"
	ViVnfm_CreateComputeResourceAffinityOrAntiAffinityConstraintsGroup_FullMethodName = "/kubenvf.kubevim.api.pb.vi_vnfm/CreateComputeResourceAffinityOrAntiAffinityConstraintsGroup"
	ViVnfm_CreateComputeFlavour_FullMethodName                                        = "/kubenvf.kubevim.api.pb.vi_vnfm/CreateComputeFlavour"
	ViVnfm_QueryComputeFlavour_FullMethodName                                         = "/kubenvf.kubevim.api.pb.vi_vnfm/QueryComputeFlavour"
	ViVnfm_DeleteComputeFlavour_FullMethodName                                        = "/kubenvf.kubevim.api.pb.vi_vnfm/DeleteComputeFlavour"
	ViVnfm_AllocateVirtualisedNetworkResource_FullMethodName                          = "/kubenvf.kubevim.api.pb.vi_vnfm/AllocateVirtualisedNetworkResource"
	ViVnfm_QueryVirtualisedNetworkResource_FullMethodName                             = "/kubenvf.kubevim.api.pb.vi_vnfm/QueryVirtualisedNetworkResource"
	ViVnfm_TerminateVirtualisedNetworkResource_FullMethodName                         = "/kubenvf.kubevim.api.pb.vi_vnfm/TerminateVirtualisedNetworkResource"
)

// ViVnfmClient is the client API for ViVnfm service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ViVnfmClient interface {
	// Software Image Management Interface
	// Requirements: The Software Image Management interface produced by the VIM on the reference point Vi-Vnfm
	// shall support querying information of software image(s) from the VIM.
	// Result: As a result of this operation, the producer (VIM) shall indicate to the consumer (VNFM) whether or not it was possible to process the query
	QueryImages(ctx context.Context, in *QueryImagesRequest, opts ...grpc.CallOption) (*QueryImagesResponse, error)
	// This operation allows querying the information about a specific software image in the image repository managed by the VIM.
	// Result: As a result of this operation, the producer (VIM) shall indicate to the consumer (VNFM) whether or not it was possible to process the query.
	QueryImage(ctx context.Context, in *QueryImageRequest, opts ...grpc.CallOption) (*QueryImageResponse, error)
	// This operation allows requesting the allocation of virtualised compute resources as indicated by the consumer functional block.
	// Result: After successful operation, the VIM has created the internal management objects for the virtualised compute resource and allocated this
	// resource according to the input requirements and constraints. In addition, the VIM shall return to the VNFM information on the newly instantiated
	// virtualised compute resource plus any additional information about the allocate request operation.
	// The VIM may also return intermediate status reports during the allocation process. If the operation was not successful,
	// the VIM shall return to the VNFM appropriate error information.
	AllocateVirtualisedComputeResource(ctx context.Context, in *AllocateComputeRequest, opts ...grpc.CallOption) (*AllocateComputeResponse, error)
	// This operation allows querying information about instantiated virtualised compute resources.
	QueryVirtualisedComputeResource(ctx context.Context, in *QueryComputeRequest, opts ...grpc.CallOption) (*QueryComputeResponse, error)
	// This operation allows an authorized consumer functional block to request the creation of a resource affinity or
	// anti-affinity constraints group. An anti-affinity group contains resources that are not placed in proximity, e.g. that do not
	// share the same physical NFVI node. An affinity group contains resources that are placed in proximity, e.g. that do share
	// the same physical NFVI node.
	// This operation shall be supported by the VIM. It shall be supported by the VNFM, if the VNFM supports named
	// resource groups for affinity/anti-affinity.
	CreateComputeResourceAffinityOrAntiAffinityConstraintsGroup(ctx context.Context, in *CreateComputeResourceAffinityOrAntiAffinityConstraintsGroupRequest, opts ...grpc.CallOption) (*CreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse, error)
	// This operation allows requesting the creation of a flavour as indicated by the consumer functional block.
	// Result: After successful operation, the VIM has created the Compute Flavour.
	// In addition, the VIM shall return to the VNFM information on the newly created Compute Flavour.
	// If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
	CreateComputeFlavour(ctx context.Context, in *CreateComputeFlavourRequest, opts ...grpc.CallOption) (*CreateComputeFlavourResponse, error)
	// This operation allows querying information about created Compute Flavours.
	// Result: After successful operation, the VIM has queried the internal management objects for the Compute Flavours. The result of the query shall indicate with
	// a standard success/error result if the query has been processed correctly. For a particular query, information about the
	// Compute Flavours that the VNFM has access to and that are matching the filter shall be returned.
	QueryComputeFlavour(ctx context.Context, in *QueryComputeFlavourRequest, opts ...grpc.CallOption) (*QueryComputeFlavourResponse, error)
	// This operation allows deleting a Compute Flavour.
	// Result: After successful operation, the VIM has deleted the Compute Flavour, so no new Virtualised Compute Resource can be allocated based on it.
	// The already allocated Virtualised Compute Resources are not affected. If the operation was not successful,
	// the VIM shall return to the VNFM appropriate error information.
	DeleteComputeFlavour(ctx context.Context, in *DeleteComputeFlavourRequest, opts ...grpc.CallOption) (*DeleteComputeFlavourResponse, error)
	// This operation allows requesting the allocation of virtualised network resources as indicated by the consumer functional block.
	// Result: After successful operation, the VIM has created the internal management objects for the virtualised network resource and
	// allocated this resource. In addition, the VIM shall return to the VNFM information on the newly instantiated virtualised network resource
	// plus any additional information about the allocate request operation. The VIM may also return intermediate status reports during the allocation process.
	// If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
	AllocateVirtualisedNetworkResource(ctx context.Context, in *AllocateNetworkRequest, opts ...grpc.CallOption) (*AllocateNetworkResponse, error)
	// This operation allows querying information about instantiated virtualised network resources.
	// Result: After successful operation, the VIM has queried the internal management objects for the virtualised network resources.
	// The result of the query shall indicate with a standard success/error result if the query has been processed correctly. For a
	// particular query, information about the network resources that the VNFM has access to and that are matching the filter
	// shall be returned.
	QueryVirtualisedNetworkResource(ctx context.Context, in *QueryNetworkRequest, opts ...grpc.CallOption) (*QueryNetworkResponse, error)
	// This operation allows de-allocating and terminating one or more an instantiated virtualised network resource(s).
	// When the operation is done on multiple ids, it is assumed to be best-effort, i.e. it can succeed for a subset of the ids, and
	// fail for the remaining ones.
	// Result: After successful operation, the VIM has terminated the virtualised network resources and removed the internal
	// management objects for those resources. In addition, the VIM shall return to the VNFM information on the terminated
	// virtualised network resource plus any additional information about the terminate request operation.
	// If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
	//
	// Note(dmalovan): ETSI GS NFV-IFA 006 (7.4.1.5.4) Operation result attached above is not coresponds to the reality, since
	// Ouput parameters defined in the (7.4.1.5.3) Output parameters block are not contains any (C) "additional information about
	// the terminated request operation" and (C) "appropriate error information"
	TerminateVirtualisedNetworkResource(ctx context.Context, in *TerminateNetworkRequest, opts ...grpc.CallOption) (*TerminateNetworkResponse, error)
}

type viVnfmClient struct {
	cc grpc.ClientConnInterface
}

func NewViVnfmClient(cc grpc.ClientConnInterface) ViVnfmClient {
	return &viVnfmClient{cc}
}

func (c *viVnfmClient) QueryImages(ctx context.Context, in *QueryImagesRequest, opts ...grpc.CallOption) (*QueryImagesResponse, error) {
	out := new(QueryImagesResponse)
	err := c.cc.Invoke(ctx, ViVnfm_QueryImages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viVnfmClient) QueryImage(ctx context.Context, in *QueryImageRequest, opts ...grpc.CallOption) (*QueryImageResponse, error) {
	out := new(QueryImageResponse)
	err := c.cc.Invoke(ctx, ViVnfm_QueryImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viVnfmClient) AllocateVirtualisedComputeResource(ctx context.Context, in *AllocateComputeRequest, opts ...grpc.CallOption) (*AllocateComputeResponse, error) {
	out := new(AllocateComputeResponse)
	err := c.cc.Invoke(ctx, ViVnfm_AllocateVirtualisedComputeResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viVnfmClient) QueryVirtualisedComputeResource(ctx context.Context, in *QueryComputeRequest, opts ...grpc.CallOption) (*QueryComputeResponse, error) {
	out := new(QueryComputeResponse)
	err := c.cc.Invoke(ctx, ViVnfm_QueryVirtualisedComputeResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viVnfmClient) CreateComputeResourceAffinityOrAntiAffinityConstraintsGroup(ctx context.Context, in *CreateComputeResourceAffinityOrAntiAffinityConstraintsGroupRequest, opts ...grpc.CallOption) (*CreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse, error) {
	out := new(CreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse)
	err := c.cc.Invoke(ctx, ViVnfm_CreateComputeResourceAffinityOrAntiAffinityConstraintsGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viVnfmClient) CreateComputeFlavour(ctx context.Context, in *CreateComputeFlavourRequest, opts ...grpc.CallOption) (*CreateComputeFlavourResponse, error) {
	out := new(CreateComputeFlavourResponse)
	err := c.cc.Invoke(ctx, ViVnfm_CreateComputeFlavour_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viVnfmClient) QueryComputeFlavour(ctx context.Context, in *QueryComputeFlavourRequest, opts ...grpc.CallOption) (*QueryComputeFlavourResponse, error) {
	out := new(QueryComputeFlavourResponse)
	err := c.cc.Invoke(ctx, ViVnfm_QueryComputeFlavour_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viVnfmClient) DeleteComputeFlavour(ctx context.Context, in *DeleteComputeFlavourRequest, opts ...grpc.CallOption) (*DeleteComputeFlavourResponse, error) {
	out := new(DeleteComputeFlavourResponse)
	err := c.cc.Invoke(ctx, ViVnfm_DeleteComputeFlavour_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viVnfmClient) AllocateVirtualisedNetworkResource(ctx context.Context, in *AllocateNetworkRequest, opts ...grpc.CallOption) (*AllocateNetworkResponse, error) {
	out := new(AllocateNetworkResponse)
	err := c.cc.Invoke(ctx, ViVnfm_AllocateVirtualisedNetworkResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viVnfmClient) QueryVirtualisedNetworkResource(ctx context.Context, in *QueryNetworkRequest, opts ...grpc.CallOption) (*QueryNetworkResponse, error) {
	out := new(QueryNetworkResponse)
	err := c.cc.Invoke(ctx, ViVnfm_QueryVirtualisedNetworkResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viVnfmClient) TerminateVirtualisedNetworkResource(ctx context.Context, in *TerminateNetworkRequest, opts ...grpc.CallOption) (*TerminateNetworkResponse, error) {
	out := new(TerminateNetworkResponse)
	err := c.cc.Invoke(ctx, ViVnfm_TerminateVirtualisedNetworkResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ViVnfmServer is the server API for ViVnfm service.
// All implementations must embed UnimplementedViVnfmServer
// for forward compatibility
type ViVnfmServer interface {
	// Software Image Management Interface
	// Requirements: The Software Image Management interface produced by the VIM on the reference point Vi-Vnfm
	// shall support querying information of software image(s) from the VIM.
	// Result: As a result of this operation, the producer (VIM) shall indicate to the consumer (VNFM) whether or not it was possible to process the query
	QueryImages(context.Context, *QueryImagesRequest) (*QueryImagesResponse, error)
	// This operation allows querying the information about a specific software image in the image repository managed by the VIM.
	// Result: As a result of this operation, the producer (VIM) shall indicate to the consumer (VNFM) whether or not it was possible to process the query.
	QueryImage(context.Context, *QueryImageRequest) (*QueryImageResponse, error)
	// This operation allows requesting the allocation of virtualised compute resources as indicated by the consumer functional block.
	// Result: After successful operation, the VIM has created the internal management objects for the virtualised compute resource and allocated this
	// resource according to the input requirements and constraints. In addition, the VIM shall return to the VNFM information on the newly instantiated
	// virtualised compute resource plus any additional information about the allocate request operation.
	// The VIM may also return intermediate status reports during the allocation process. If the operation was not successful,
	// the VIM shall return to the VNFM appropriate error information.
	AllocateVirtualisedComputeResource(context.Context, *AllocateComputeRequest) (*AllocateComputeResponse, error)
	// This operation allows querying information about instantiated virtualised compute resources.
	QueryVirtualisedComputeResource(context.Context, *QueryComputeRequest) (*QueryComputeResponse, error)
	// This operation allows an authorized consumer functional block to request the creation of a resource affinity or
	// anti-affinity constraints group. An anti-affinity group contains resources that are not placed in proximity, e.g. that do not
	// share the same physical NFVI node. An affinity group contains resources that are placed in proximity, e.g. that do share
	// the same physical NFVI node.
	// This operation shall be supported by the VIM. It shall be supported by the VNFM, if the VNFM supports named
	// resource groups for affinity/anti-affinity.
	CreateComputeResourceAffinityOrAntiAffinityConstraintsGroup(context.Context, *CreateComputeResourceAffinityOrAntiAffinityConstraintsGroupRequest) (*CreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse, error)
	// This operation allows requesting the creation of a flavour as indicated by the consumer functional block.
	// Result: After successful operation, the VIM has created the Compute Flavour.
	// In addition, the VIM shall return to the VNFM information on the newly created Compute Flavour.
	// If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
	CreateComputeFlavour(context.Context, *CreateComputeFlavourRequest) (*CreateComputeFlavourResponse, error)
	// This operation allows querying information about created Compute Flavours.
	// Result: After successful operation, the VIM has queried the internal management objects for the Compute Flavours. The result of the query shall indicate with
	// a standard success/error result if the query has been processed correctly. For a particular query, information about the
	// Compute Flavours that the VNFM has access to and that are matching the filter shall be returned.
	QueryComputeFlavour(context.Context, *QueryComputeFlavourRequest) (*QueryComputeFlavourResponse, error)
	// This operation allows deleting a Compute Flavour.
	// Result: After successful operation, the VIM has deleted the Compute Flavour, so no new Virtualised Compute Resource can be allocated based on it.
	// The already allocated Virtualised Compute Resources are not affected. If the operation was not successful,
	// the VIM shall return to the VNFM appropriate error information.
	DeleteComputeFlavour(context.Context, *DeleteComputeFlavourRequest) (*DeleteComputeFlavourResponse, error)
	// This operation allows requesting the allocation of virtualised network resources as indicated by the consumer functional block.
	// Result: After successful operation, the VIM has created the internal management objects for the virtualised network resource and
	// allocated this resource. In addition, the VIM shall return to the VNFM information on the newly instantiated virtualised network resource
	// plus any additional information about the allocate request operation. The VIM may also return intermediate status reports during the allocation process.
	// If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
	AllocateVirtualisedNetworkResource(context.Context, *AllocateNetworkRequest) (*AllocateNetworkResponse, error)
	// This operation allows querying information about instantiated virtualised network resources.
	// Result: After successful operation, the VIM has queried the internal management objects for the virtualised network resources.
	// The result of the query shall indicate with a standard success/error result if the query has been processed correctly. For a
	// particular query, information about the network resources that the VNFM has access to and that are matching the filter
	// shall be returned.
	QueryVirtualisedNetworkResource(context.Context, *QueryNetworkRequest) (*QueryNetworkResponse, error)
	// This operation allows de-allocating and terminating one or more an instantiated virtualised network resource(s).
	// When the operation is done on multiple ids, it is assumed to be best-effort, i.e. it can succeed for a subset of the ids, and
	// fail for the remaining ones.
	// Result: After successful operation, the VIM has terminated the virtualised network resources and removed the internal
	// management objects for those resources. In addition, the VIM shall return to the VNFM information on the terminated
	// virtualised network resource plus any additional information about the terminate request operation.
	// If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
	//
	// Note(dmalovan): ETSI GS NFV-IFA 006 (7.4.1.5.4) Operation result attached above is not coresponds to the reality, since
	// Ouput parameters defined in the (7.4.1.5.3) Output parameters block are not contains any (C) "additional information about
	// the terminated request operation" and (C) "appropriate error information"
	TerminateVirtualisedNetworkResource(context.Context, *TerminateNetworkRequest) (*TerminateNetworkResponse, error)
	mustEmbedUnimplementedViVnfmServer()
}

// UnimplementedViVnfmServer must be embedded to have forward compatible implementations.
type UnimplementedViVnfmServer struct {
}

func (UnimplementedViVnfmServer) QueryImages(context.Context, *QueryImagesRequest) (*QueryImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryImages not implemented")
}
func (UnimplementedViVnfmServer) QueryImage(context.Context, *QueryImageRequest) (*QueryImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryImage not implemented")
}
func (UnimplementedViVnfmServer) AllocateVirtualisedComputeResource(context.Context, *AllocateComputeRequest) (*AllocateComputeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocateVirtualisedComputeResource not implemented")
}
func (UnimplementedViVnfmServer) QueryVirtualisedComputeResource(context.Context, *QueryComputeRequest) (*QueryComputeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryVirtualisedComputeResource not implemented")
}
func (UnimplementedViVnfmServer) CreateComputeResourceAffinityOrAntiAffinityConstraintsGroup(context.Context, *CreateComputeResourceAffinityOrAntiAffinityConstraintsGroupRequest) (*CreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComputeResourceAffinityOrAntiAffinityConstraintsGroup not implemented")
}
func (UnimplementedViVnfmServer) CreateComputeFlavour(context.Context, *CreateComputeFlavourRequest) (*CreateComputeFlavourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComputeFlavour not implemented")
}
func (UnimplementedViVnfmServer) QueryComputeFlavour(context.Context, *QueryComputeFlavourRequest) (*QueryComputeFlavourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryComputeFlavour not implemented")
}
func (UnimplementedViVnfmServer) DeleteComputeFlavour(context.Context, *DeleteComputeFlavourRequest) (*DeleteComputeFlavourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComputeFlavour not implemented")
}
func (UnimplementedViVnfmServer) AllocateVirtualisedNetworkResource(context.Context, *AllocateNetworkRequest) (*AllocateNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocateVirtualisedNetworkResource not implemented")
}
func (UnimplementedViVnfmServer) QueryVirtualisedNetworkResource(context.Context, *QueryNetworkRequest) (*QueryNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryVirtualisedNetworkResource not implemented")
}
func (UnimplementedViVnfmServer) TerminateVirtualisedNetworkResource(context.Context, *TerminateNetworkRequest) (*TerminateNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TerminateVirtualisedNetworkResource not implemented")
}
func (UnimplementedViVnfmServer) mustEmbedUnimplementedViVnfmServer() {}

// UnsafeViVnfmServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ViVnfmServer will
// result in compilation errors.
type UnsafeViVnfmServer interface {
	mustEmbedUnimplementedViVnfmServer()
}

func RegisterViVnfmServer(s grpc.ServiceRegistrar, srv ViVnfmServer) {
	s.RegisterService(&ViVnfm_ServiceDesc, srv)
}

func _ViVnfm_QueryImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViVnfmServer).QueryImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViVnfm_QueryImages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViVnfmServer).QueryImages(ctx, req.(*QueryImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViVnfm_QueryImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViVnfmServer).QueryImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViVnfm_QueryImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViVnfmServer).QueryImage(ctx, req.(*QueryImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViVnfm_AllocateVirtualisedComputeResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocateComputeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViVnfmServer).AllocateVirtualisedComputeResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViVnfm_AllocateVirtualisedComputeResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViVnfmServer).AllocateVirtualisedComputeResource(ctx, req.(*AllocateComputeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViVnfm_QueryVirtualisedComputeResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryComputeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViVnfmServer).QueryVirtualisedComputeResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViVnfm_QueryVirtualisedComputeResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViVnfmServer).QueryVirtualisedComputeResource(ctx, req.(*QueryComputeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViVnfm_CreateComputeResourceAffinityOrAntiAffinityConstraintsGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateComputeResourceAffinityOrAntiAffinityConstraintsGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViVnfmServer).CreateComputeResourceAffinityOrAntiAffinityConstraintsGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViVnfm_CreateComputeResourceAffinityOrAntiAffinityConstraintsGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViVnfmServer).CreateComputeResourceAffinityOrAntiAffinityConstraintsGroup(ctx, req.(*CreateComputeResourceAffinityOrAntiAffinityConstraintsGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViVnfm_CreateComputeFlavour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateComputeFlavourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViVnfmServer).CreateComputeFlavour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViVnfm_CreateComputeFlavour_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViVnfmServer).CreateComputeFlavour(ctx, req.(*CreateComputeFlavourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViVnfm_QueryComputeFlavour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryComputeFlavourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViVnfmServer).QueryComputeFlavour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViVnfm_QueryComputeFlavour_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViVnfmServer).QueryComputeFlavour(ctx, req.(*QueryComputeFlavourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViVnfm_DeleteComputeFlavour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteComputeFlavourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViVnfmServer).DeleteComputeFlavour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViVnfm_DeleteComputeFlavour_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViVnfmServer).DeleteComputeFlavour(ctx, req.(*DeleteComputeFlavourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViVnfm_AllocateVirtualisedNetworkResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocateNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViVnfmServer).AllocateVirtualisedNetworkResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViVnfm_AllocateVirtualisedNetworkResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViVnfmServer).AllocateVirtualisedNetworkResource(ctx, req.(*AllocateNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViVnfm_QueryVirtualisedNetworkResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViVnfmServer).QueryVirtualisedNetworkResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViVnfm_QueryVirtualisedNetworkResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViVnfmServer).QueryVirtualisedNetworkResource(ctx, req.(*QueryNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViVnfm_TerminateVirtualisedNetworkResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminateNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViVnfmServer).TerminateVirtualisedNetworkResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViVnfm_TerminateVirtualisedNetworkResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViVnfmServer).TerminateVirtualisedNetworkResource(ctx, req.(*TerminateNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ViVnfm_ServiceDesc is the grpc.ServiceDesc for ViVnfm service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ViVnfm_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kubenvf.kubevim.api.pb.vi_vnfm",
	HandlerType: (*ViVnfmServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryImages",
			Handler:    _ViVnfm_QueryImages_Handler,
		},
		{
			MethodName: "QueryImage",
			Handler:    _ViVnfm_QueryImage_Handler,
		},
		{
			MethodName: "AllocateVirtualisedComputeResource",
			Handler:    _ViVnfm_AllocateVirtualisedComputeResource_Handler,
		},
		{
			MethodName: "QueryVirtualisedComputeResource",
			Handler:    _ViVnfm_QueryVirtualisedComputeResource_Handler,
		},
		{
			MethodName: "CreateComputeResourceAffinityOrAntiAffinityConstraintsGroup",
			Handler:    _ViVnfm_CreateComputeResourceAffinityOrAntiAffinityConstraintsGroup_Handler,
		},
		{
			MethodName: "CreateComputeFlavour",
			Handler:    _ViVnfm_CreateComputeFlavour_Handler,
		},
		{
			MethodName: "QueryComputeFlavour",
			Handler:    _ViVnfm_QueryComputeFlavour_Handler,
		},
		{
			MethodName: "DeleteComputeFlavour",
			Handler:    _ViVnfm_DeleteComputeFlavour_Handler,
		},
		{
			MethodName: "AllocateVirtualisedNetworkResource",
			Handler:    _ViVnfm_AllocateVirtualisedNetworkResource_Handler,
		},
		{
			MethodName: "QueryVirtualisedNetworkResource",
			Handler:    _ViVnfm_QueryVirtualisedNetworkResource_Handler,
		},
		{
			MethodName: "TerminateVirtualisedNetworkResource",
			Handler:    _ViVnfm_TerminateVirtualisedNetworkResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vi-vnfm.proto",
}
