// Copyright (c) 2023 Arista Networks, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.11.3
// source: ovgs.proto

package ovgs

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
	OwnershipVoucherService_CreateGroup_FullMethodName         = "/ovgs.v1.OwnershipVoucherService/CreateGroup"
	OwnershipVoucherService_DeleteGroup_FullMethodName         = "/ovgs.v1.OwnershipVoucherService/DeleteGroup"
	OwnershipVoucherService_GetGroup_FullMethodName            = "/ovgs.v1.OwnershipVoucherService/GetGroup"
	OwnershipVoucherService_AddUserRole_FullMethodName         = "/ovgs.v1.OwnershipVoucherService/AddUserRole"
	OwnershipVoucherService_RemoveUserRole_FullMethodName      = "/ovgs.v1.OwnershipVoucherService/RemoveUserRole"
	OwnershipVoucherService_GetUserRole_FullMethodName         = "/ovgs.v1.OwnershipVoucherService/GetUserRole"
	OwnershipVoucherService_AddSerial_FullMethodName           = "/ovgs.v1.OwnershipVoucherService/AddSerial"
	OwnershipVoucherService_RemoveSerial_FullMethodName        = "/ovgs.v1.OwnershipVoucherService/RemoveSerial"
	OwnershipVoucherService_GetSerial_FullMethodName           = "/ovgs.v1.OwnershipVoucherService/GetSerial"
	OwnershipVoucherService_CreateDomainCert_FullMethodName    = "/ovgs.v1.OwnershipVoucherService/CreateDomainCert"
	OwnershipVoucherService_DeleteDomainCert_FullMethodName    = "/ovgs.v1.OwnershipVoucherService/DeleteDomainCert"
	OwnershipVoucherService_GetDomainCert_FullMethodName       = "/ovgs.v1.OwnershipVoucherService/GetDomainCert"
	OwnershipVoucherService_GetOwnershipVoucher_FullMethodName = "/ovgs.v1.OwnershipVoucherService/GetOwnershipVoucher"
)

// OwnershipVoucherServiceClient is the client API for OwnershipVoucherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OwnershipVoucherServiceClient interface {
	// CreateGroup creates a group as a child of an existing group.
	// Errors will be returned:
	// INVALID_ARGUMENT if either parent or description is empty
	// NOT_FOUND if the parent group doesn't exist, as specified in request
	// ALREADY_EXISTS if a group already exists with the same parent group
	// and description.
	// PERMISSION_DENIED if the user doesn't have access to the parent group
	// Roles with permission to invoke this = ADMIN
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	// DeleteGroup deletes a named group. All associated cert_ids and child
	// groups must have been deleted and all associated components must have
	// been removed before the group can be deleted.
	// Errors will be returned:
	// INVALID_ARGUMENT if group_id = root group (the precreated root group
	// cannot be deleted) or if cert_ids, users, or child_group_ids is
	// non-empty.
	// NOT_FOUND if the group doesn't exist
	// PERMISSION_DENIED if user doesn't have access to parent group
	// Roles with permission to invoke this = ADMIN
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error)
	// GetGroup returns the domain-certs (keyed by id), components,
	// user/role mappings for that group, and the child_group_ids.
	// Errors will be returned:
	// NOT_FOUND if the group doesn't exist
	// PERMISSION_DENIED if the user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN, ASSIGNER, REQUESTOR
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error)
	// AddUserRole will assign a role to a user in a named group.
	// Username is unique to an username, org_id, user_type tuple.
	// Errors will be returned:
	// INVALID_ARGUMENT if any field is empty
	// NOT_FOUND if the group doesn't exist, as specified in the request
	// FAILED_PRECONDITION if any of user tuple (username, org_id, user_type)
	// or group_id do not exist.
	// ALREADY_EXISTS if user already exists in the group
	// PERMISSION_DENIED if the user doesn't have access to the group.
	// Roles with permission to invoke this = ADMIN
	AddUserRole(ctx context.Context, in *AddUserRoleRequest, opts ...grpc.CallOption) (*AddUserRoleResponse, error)
	// RemoveUserRole removes a role from a user in a named group.
	// Username is unique to an username, org_id, user_type tuple.
	// Errors will be returned:
	// INVALID_ARGUMENT if any field is empty
	// NOT_FOUND if the group doesn't exist or if the user tuple is not a
	// member of the group.
	// PERMISSION_DENIED if user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN
	RemoveUserRole(ctx context.Context, in *RemoveUserRoleRequest, opts ...grpc.CallOption) (*RemoveUserRoleResponse, error)
	// GetUserRole returns the roles the user is assigned in the group.
	// Username is unique to an username, org_id, user_type tuple.
	// A user can only view roles of another user in the groups that
	// it has a role assigned to.
	// Errors will be returned:
	// INVALID_ARGUMENT if any field is empty
	// NOT_FOUND if the group doesn't exist or the user tuple is not a member.
	// Roles with permission to invoke this = ADMIN, ASSIGNER, REQUESTOR
	GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*GetUserRoleResponse, error)
	// AddSerial assigns the component to the group.
	// Errors will be returned:
	// INVALID_ARGUMENT if any field of component or group_id is empty or the
	// IEN isn't applicable for the voucher issuer.
	// NOT_FOUND if the component or group_id doesn't exist
	// ALREADY_EXISTS if component is already a member of the group.
	// PERMISSION_DENIED if the user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN, ASSIGNER
	AddSerial(ctx context.Context, in *AddSerialRequest, opts ...grpc.CallOption) (*AddSerialResponse, error)
	// RemoveSerial removes the component from the group.
	// Errors will be returned:
	// INVALID_ARGUMENT if any field of component or group_id is empty or the
	// IEN isn't applicable for the voucher issuer.
	// NOT_FOUND if the component or group_id doesn't exist
	// PERMISSION_DENIED if user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN, ASSIGNER
	RemoveSerial(ctx context.Context, in *RemoveSerialRequest, opts ...grpc.CallOption) (*RemoveSerialResponse, error)
	// GetSerial returns component, groups the component belongs to.
	// Errors will be returned:
	// INVALID_ARGUMENT if any field of component or group_id is empty or the
	// IEN isn't applicable for the voucher issuer.
	// NOT_FOUND if the component doesn't exist.
	// PERMISSION_DENIED if the user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN, ASSIGNER, REQUESTOR
	GetSerial(ctx context.Context, in *GetSerialRequest, opts ...grpc.CallOption) (*GetSerialResponse, error)
	// CreateDomainCert creates the certificate in the group.
	// Errors will be returned:
	// INVALID_ARGUMENT if expiry_time is empty or in the past, the
	// supplied cert is invalid (such expired or malformed), or any of the
	// fields are empty.
	// NOT_FOUND if the group_id doesn't exist.
	// ALREADY_EXISTS if the certificate_der,revocation_checks,expiry_time
	// tuple already exists in the group.
	// PERMISSION_DENIED if the user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN, ASSIGNER
	CreateDomainCert(ctx context.Context, in *CreateDomainCertRequest, opts ...grpc.CallOption) (*CreateDomainCertResponse, error)
	// DeleteDomainCert deletes the cert_id.
	// Errors will be returned:
	// NOT_FOUND if the cert_id doesn't exist
	// PERMISSION_DENIED if user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN, ASSIGNER
	DeleteDomainCert(ctx context.Context, in *DeleteDomainCertRequest, opts ...grpc.CallOption) (*DeleteDomainCertResponse, error)
	// GetDomainCert returns the details of the cert_id.
	// NOT_FOUND if the cert_id doesn't exist.
	// PERMISSION_DENIED if user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN, ASSIGNER, REQUESTOR
	GetDomainCert(ctx context.Context, in *GetDomainCertRequest, opts ...grpc.CallOption) (*GetDomainCertResponse, error)
	// GetOwnershipVoucher issues an ownership voucher for the component (if it
	// exists/if applicable)
	// Errors will be returned:
	// INVALID_ARGUMENT if any field of the request is empty, lifetime is in
	// the past, or the IEN supplied isn't applicable for the voucher issuer
	// FAILED_PRECONDITION if the component or cert_id do not exist.
	// PERMISSION_DENIED if user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN, ASSIGNER, REQUESTOR
	GetOwnershipVoucher(ctx context.Context, in *GetOwnershipVoucherRequest, opts ...grpc.CallOption) (*GetOwnershipVoucherResponse, error)
}

type ownershipVoucherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOwnershipVoucherServiceClient(cc grpc.ClientConnInterface) OwnershipVoucherServiceClient {
	return &ownershipVoucherServiceClient{cc}
}

func (c *ownershipVoucherServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, OwnershipVoucherService_CreateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownershipVoucherServiceClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error) {
	out := new(DeleteGroupResponse)
	err := c.cc.Invoke(ctx, OwnershipVoucherService_DeleteGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownershipVoucherServiceClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error) {
	out := new(GetGroupResponse)
	err := c.cc.Invoke(ctx, OwnershipVoucherService_GetGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownershipVoucherServiceClient) AddUserRole(ctx context.Context, in *AddUserRoleRequest, opts ...grpc.CallOption) (*AddUserRoleResponse, error) {
	out := new(AddUserRoleResponse)
	err := c.cc.Invoke(ctx, OwnershipVoucherService_AddUserRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownershipVoucherServiceClient) RemoveUserRole(ctx context.Context, in *RemoveUserRoleRequest, opts ...grpc.CallOption) (*RemoveUserRoleResponse, error) {
	out := new(RemoveUserRoleResponse)
	err := c.cc.Invoke(ctx, OwnershipVoucherService_RemoveUserRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownershipVoucherServiceClient) GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*GetUserRoleResponse, error) {
	out := new(GetUserRoleResponse)
	err := c.cc.Invoke(ctx, OwnershipVoucherService_GetUserRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownershipVoucherServiceClient) AddSerial(ctx context.Context, in *AddSerialRequest, opts ...grpc.CallOption) (*AddSerialResponse, error) {
	out := new(AddSerialResponse)
	err := c.cc.Invoke(ctx, OwnershipVoucherService_AddSerial_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownershipVoucherServiceClient) RemoveSerial(ctx context.Context, in *RemoveSerialRequest, opts ...grpc.CallOption) (*RemoveSerialResponse, error) {
	out := new(RemoveSerialResponse)
	err := c.cc.Invoke(ctx, OwnershipVoucherService_RemoveSerial_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownershipVoucherServiceClient) GetSerial(ctx context.Context, in *GetSerialRequest, opts ...grpc.CallOption) (*GetSerialResponse, error) {
	out := new(GetSerialResponse)
	err := c.cc.Invoke(ctx, OwnershipVoucherService_GetSerial_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownershipVoucherServiceClient) CreateDomainCert(ctx context.Context, in *CreateDomainCertRequest, opts ...grpc.CallOption) (*CreateDomainCertResponse, error) {
	out := new(CreateDomainCertResponse)
	err := c.cc.Invoke(ctx, OwnershipVoucherService_CreateDomainCert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownershipVoucherServiceClient) DeleteDomainCert(ctx context.Context, in *DeleteDomainCertRequest, opts ...grpc.CallOption) (*DeleteDomainCertResponse, error) {
	out := new(DeleteDomainCertResponse)
	err := c.cc.Invoke(ctx, OwnershipVoucherService_DeleteDomainCert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownershipVoucherServiceClient) GetDomainCert(ctx context.Context, in *GetDomainCertRequest, opts ...grpc.CallOption) (*GetDomainCertResponse, error) {
	out := new(GetDomainCertResponse)
	err := c.cc.Invoke(ctx, OwnershipVoucherService_GetDomainCert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownershipVoucherServiceClient) GetOwnershipVoucher(ctx context.Context, in *GetOwnershipVoucherRequest, opts ...grpc.CallOption) (*GetOwnershipVoucherResponse, error) {
	out := new(GetOwnershipVoucherResponse)
	err := c.cc.Invoke(ctx, OwnershipVoucherService_GetOwnershipVoucher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OwnershipVoucherServiceServer is the server API for OwnershipVoucherService service.
// All implementations must embed UnimplementedOwnershipVoucherServiceServer
// for forward compatibility
type OwnershipVoucherServiceServer interface {
	// CreateGroup creates a group as a child of an existing group.
	// Errors will be returned:
	// INVALID_ARGUMENT if either parent or description is empty
	// NOT_FOUND if the parent group doesn't exist, as specified in request
	// ALREADY_EXISTS if a group already exists with the same parent group
	// and description.
	// PERMISSION_DENIED if the user doesn't have access to the parent group
	// Roles with permission to invoke this = ADMIN
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	// DeleteGroup deletes a named group. All associated cert_ids and child
	// groups must have been deleted and all associated components must have
	// been removed before the group can be deleted.
	// Errors will be returned:
	// INVALID_ARGUMENT if group_id = root group (the precreated root group
	// cannot be deleted) or if cert_ids, users, or child_group_ids is
	// non-empty.
	// NOT_FOUND if the group doesn't exist
	// PERMISSION_DENIED if user doesn't have access to parent group
	// Roles with permission to invoke this = ADMIN
	DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error)
	// GetGroup returns the domain-certs (keyed by id), components,
	// user/role mappings for that group, and the child_group_ids.
	// Errors will be returned:
	// NOT_FOUND if the group doesn't exist
	// PERMISSION_DENIED if the user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN, ASSIGNER, REQUESTOR
	GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error)
	// AddUserRole will assign a role to a user in a named group.
	// Username is unique to an username, org_id, user_type tuple.
	// Errors will be returned:
	// INVALID_ARGUMENT if any field is empty
	// NOT_FOUND if the group doesn't exist, as specified in the request
	// FAILED_PRECONDITION if any of user tuple (username, org_id, user_type)
	// or group_id do not exist.
	// ALREADY_EXISTS if user already exists in the group
	// PERMISSION_DENIED if the user doesn't have access to the group.
	// Roles with permission to invoke this = ADMIN
	AddUserRole(context.Context, *AddUserRoleRequest) (*AddUserRoleResponse, error)
	// RemoveUserRole removes a role from a user in a named group.
	// Username is unique to an username, org_id, user_type tuple.
	// Errors will be returned:
	// INVALID_ARGUMENT if any field is empty
	// NOT_FOUND if the group doesn't exist or if the user tuple is not a
	// member of the group.
	// PERMISSION_DENIED if user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN
	RemoveUserRole(context.Context, *RemoveUserRoleRequest) (*RemoveUserRoleResponse, error)
	// GetUserRole returns the roles the user is assigned in the group.
	// Username is unique to an username, org_id, user_type tuple.
	// A user can only view roles of another user in the groups that
	// it has a role assigned to.
	// Errors will be returned:
	// INVALID_ARGUMENT if any field is empty
	// NOT_FOUND if the group doesn't exist or the user tuple is not a member.
	// Roles with permission to invoke this = ADMIN, ASSIGNER, REQUESTOR
	GetUserRole(context.Context, *GetUserRoleRequest) (*GetUserRoleResponse, error)
	// AddSerial assigns the component to the group.
	// Errors will be returned:
	// INVALID_ARGUMENT if any field of component or group_id is empty or the
	// IEN isn't applicable for the voucher issuer.
	// NOT_FOUND if the component or group_id doesn't exist
	// ALREADY_EXISTS if component is already a member of the group.
	// PERMISSION_DENIED if the user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN, ASSIGNER
	AddSerial(context.Context, *AddSerialRequest) (*AddSerialResponse, error)
	// RemoveSerial removes the component from the group.
	// Errors will be returned:
	// INVALID_ARGUMENT if any field of component or group_id is empty or the
	// IEN isn't applicable for the voucher issuer.
	// NOT_FOUND if the component or group_id doesn't exist
	// PERMISSION_DENIED if user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN, ASSIGNER
	RemoveSerial(context.Context, *RemoveSerialRequest) (*RemoveSerialResponse, error)
	// GetSerial returns component, groups the component belongs to.
	// Errors will be returned:
	// INVALID_ARGUMENT if any field of component or group_id is empty or the
	// IEN isn't applicable for the voucher issuer.
	// NOT_FOUND if the component doesn't exist.
	// PERMISSION_DENIED if the user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN, ASSIGNER, REQUESTOR
	GetSerial(context.Context, *GetSerialRequest) (*GetSerialResponse, error)
	// CreateDomainCert creates the certificate in the group.
	// Errors will be returned:
	// INVALID_ARGUMENT if expiry_time is empty or in the past, the
	// supplied cert is invalid (such expired or malformed), or any of the
	// fields are empty.
	// NOT_FOUND if the group_id doesn't exist.
	// ALREADY_EXISTS if the certificate_der,revocation_checks,expiry_time
	// tuple already exists in the group.
	// PERMISSION_DENIED if the user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN, ASSIGNER
	CreateDomainCert(context.Context, *CreateDomainCertRequest) (*CreateDomainCertResponse, error)
	// DeleteDomainCert deletes the cert_id.
	// Errors will be returned:
	// NOT_FOUND if the cert_id doesn't exist
	// PERMISSION_DENIED if user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN, ASSIGNER
	DeleteDomainCert(context.Context, *DeleteDomainCertRequest) (*DeleteDomainCertResponse, error)
	// GetDomainCert returns the details of the cert_id.
	// NOT_FOUND if the cert_id doesn't exist.
	// PERMISSION_DENIED if user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN, ASSIGNER, REQUESTOR
	GetDomainCert(context.Context, *GetDomainCertRequest) (*GetDomainCertResponse, error)
	// GetOwnershipVoucher issues an ownership voucher for the component (if it
	// exists/if applicable)
	// Errors will be returned:
	// INVALID_ARGUMENT if any field of the request is empty, lifetime is in
	// the past, or the IEN supplied isn't applicable for the voucher issuer
	// FAILED_PRECONDITION if the component or cert_id do not exist.
	// PERMISSION_DENIED if user doesn't have access to the group
	// Roles with permission to invoke this = ADMIN, ASSIGNER, REQUESTOR
	GetOwnershipVoucher(context.Context, *GetOwnershipVoucherRequest) (*GetOwnershipVoucherResponse, error)
	mustEmbedUnimplementedOwnershipVoucherServiceServer()
}

// UnimplementedOwnershipVoucherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOwnershipVoucherServiceServer struct {
}

func (UnimplementedOwnershipVoucherServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedOwnershipVoucherServiceServer) DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedOwnershipVoucherServiceServer) GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedOwnershipVoucherServiceServer) AddUserRole(context.Context, *AddUserRoleRequest) (*AddUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserRole not implemented")
}
func (UnimplementedOwnershipVoucherServiceServer) RemoveUserRole(context.Context, *RemoveUserRoleRequest) (*RemoveUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserRole not implemented")
}
func (UnimplementedOwnershipVoucherServiceServer) GetUserRole(context.Context, *GetUserRoleRequest) (*GetUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRole not implemented")
}
func (UnimplementedOwnershipVoucherServiceServer) AddSerial(context.Context, *AddSerialRequest) (*AddSerialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSerial not implemented")
}
func (UnimplementedOwnershipVoucherServiceServer) RemoveSerial(context.Context, *RemoveSerialRequest) (*RemoveSerialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSerial not implemented")
}
func (UnimplementedOwnershipVoucherServiceServer) GetSerial(context.Context, *GetSerialRequest) (*GetSerialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSerial not implemented")
}
func (UnimplementedOwnershipVoucherServiceServer) CreateDomainCert(context.Context, *CreateDomainCertRequest) (*CreateDomainCertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDomainCert not implemented")
}
func (UnimplementedOwnershipVoucherServiceServer) DeleteDomainCert(context.Context, *DeleteDomainCertRequest) (*DeleteDomainCertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDomainCert not implemented")
}
func (UnimplementedOwnershipVoucherServiceServer) GetDomainCert(context.Context, *GetDomainCertRequest) (*GetDomainCertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomainCert not implemented")
}
func (UnimplementedOwnershipVoucherServiceServer) GetOwnershipVoucher(context.Context, *GetOwnershipVoucherRequest) (*GetOwnershipVoucherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOwnershipVoucher not implemented")
}
func (UnimplementedOwnershipVoucherServiceServer) mustEmbedUnimplementedOwnershipVoucherServiceServer() {
}

// UnsafeOwnershipVoucherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OwnershipVoucherServiceServer will
// result in compilation errors.
type UnsafeOwnershipVoucherServiceServer interface {
	mustEmbedUnimplementedOwnershipVoucherServiceServer()
}

func RegisterOwnershipVoucherServiceServer(s grpc.ServiceRegistrar, srv OwnershipVoucherServiceServer) {
	s.RegisterService(&OwnershipVoucherService_ServiceDesc, srv)
}

func _OwnershipVoucherService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnershipVoucherServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnershipVoucherService_CreateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnershipVoucherServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnershipVoucherService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnershipVoucherServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnershipVoucherService_DeleteGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnershipVoucherServiceServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnershipVoucherService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnershipVoucherServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnershipVoucherService_GetGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnershipVoucherServiceServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnershipVoucherService_AddUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnershipVoucherServiceServer).AddUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnershipVoucherService_AddUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnershipVoucherServiceServer).AddUserRole(ctx, req.(*AddUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnershipVoucherService_RemoveUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnershipVoucherServiceServer).RemoveUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnershipVoucherService_RemoveUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnershipVoucherServiceServer).RemoveUserRole(ctx, req.(*RemoveUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnershipVoucherService_GetUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnershipVoucherServiceServer).GetUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnershipVoucherService_GetUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnershipVoucherServiceServer).GetUserRole(ctx, req.(*GetUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnershipVoucherService_AddSerial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSerialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnershipVoucherServiceServer).AddSerial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnershipVoucherService_AddSerial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnershipVoucherServiceServer).AddSerial(ctx, req.(*AddSerialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnershipVoucherService_RemoveSerial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSerialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnershipVoucherServiceServer).RemoveSerial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnershipVoucherService_RemoveSerial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnershipVoucherServiceServer).RemoveSerial(ctx, req.(*RemoveSerialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnershipVoucherService_GetSerial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSerialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnershipVoucherServiceServer).GetSerial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnershipVoucherService_GetSerial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnershipVoucherServiceServer).GetSerial(ctx, req.(*GetSerialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnershipVoucherService_CreateDomainCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDomainCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnershipVoucherServiceServer).CreateDomainCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnershipVoucherService_CreateDomainCert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnershipVoucherServiceServer).CreateDomainCert(ctx, req.(*CreateDomainCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnershipVoucherService_DeleteDomainCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDomainCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnershipVoucherServiceServer).DeleteDomainCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnershipVoucherService_DeleteDomainCert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnershipVoucherServiceServer).DeleteDomainCert(ctx, req.(*DeleteDomainCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnershipVoucherService_GetDomainCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnershipVoucherServiceServer).GetDomainCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnershipVoucherService_GetDomainCert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnershipVoucherServiceServer).GetDomainCert(ctx, req.(*GetDomainCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnershipVoucherService_GetOwnershipVoucher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOwnershipVoucherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnershipVoucherServiceServer).GetOwnershipVoucher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnershipVoucherService_GetOwnershipVoucher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnershipVoucherServiceServer).GetOwnershipVoucher(ctx, req.(*GetOwnershipVoucherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OwnershipVoucherService_ServiceDesc is the grpc.ServiceDesc for OwnershipVoucherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OwnershipVoucherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ovgs.v1.OwnershipVoucherService",
	HandlerType: (*OwnershipVoucherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _OwnershipVoucherService_CreateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _OwnershipVoucherService_DeleteGroup_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _OwnershipVoucherService_GetGroup_Handler,
		},
		{
			MethodName: "AddUserRole",
			Handler:    _OwnershipVoucherService_AddUserRole_Handler,
		},
		{
			MethodName: "RemoveUserRole",
			Handler:    _OwnershipVoucherService_RemoveUserRole_Handler,
		},
		{
			MethodName: "GetUserRole",
			Handler:    _OwnershipVoucherService_GetUserRole_Handler,
		},
		{
			MethodName: "AddSerial",
			Handler:    _OwnershipVoucherService_AddSerial_Handler,
		},
		{
			MethodName: "RemoveSerial",
			Handler:    _OwnershipVoucherService_RemoveSerial_Handler,
		},
		{
			MethodName: "GetSerial",
			Handler:    _OwnershipVoucherService_GetSerial_Handler,
		},
		{
			MethodName: "CreateDomainCert",
			Handler:    _OwnershipVoucherService_CreateDomainCert_Handler,
		},
		{
			MethodName: "DeleteDomainCert",
			Handler:    _OwnershipVoucherService_DeleteDomainCert_Handler,
		},
		{
			MethodName: "GetDomainCert",
			Handler:    _OwnershipVoucherService_GetDomainCert_Handler,
		},
		{
			MethodName: "GetOwnershipVoucher",
			Handler:    _OwnershipVoucherService_GetOwnershipVoucher_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ovgs.proto",
}
