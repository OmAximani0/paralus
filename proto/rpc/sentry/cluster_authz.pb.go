// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/rpc/sentry/cluster_authz.proto

package sentry

import (
	_ "github.com/paralus/paralus/proto/types/commonpb/v3"
	controller "github.com/paralus/paralus/proto/types/controller"
	_ "github.com/paralus/paralus/proto/types/sentry"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetUserAuthorizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName                        string                   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	ServiceAccount                  *controller.StepObject   `protobuf:"bytes,2,opt,name=serviceAccount,proto3" json:"serviceAccount,omitempty"`
	ClusterRoles                    []*controller.StepObject `protobuf:"bytes,3,rep,name=clusterRoles,proto3" json:"clusterRoles,omitempty"`
	ClusterRoleBindings             []*controller.StepObject `protobuf:"bytes,4,rep,name=clusterRoleBindings,proto3" json:"clusterRoleBindings,omitempty"`
	Roles                           []*controller.StepObject `protobuf:"bytes,5,rep,name=roles,proto3" json:"roles,omitempty"`
	RoleBindings                    []*controller.StepObject `protobuf:"bytes,6,rep,name=roleBindings,proto3" json:"roleBindings,omitempty"`
	DeleteClusterRoleBindings       []*controller.StepObject `protobuf:"bytes,7,rep,name=deleteClusterRoleBindings,proto3" json:"deleteClusterRoleBindings,omitempty"`
	DeleteRoleBindings              []*controller.StepObject `protobuf:"bytes,8,rep,name=deleteRoleBindings,proto3" json:"deleteRoleBindings,omitempty"`
	Namespaces                      []*controller.StepObject `protobuf:"bytes,9,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	RoleName                        string                   `protobuf:"bytes,10,opt,name=roleName,proto3" json:"roleName,omitempty"`
	IsRead                          bool                     `protobuf:"varint,11,opt,name=isRead,proto3" json:"isRead,omitempty"`
	EnforceOrgAdminOnlySecretAccess bool                     `protobuf:"varint,12,opt,name=enforceOrgAdminOnlySecretAccess,proto3" json:"enforceOrgAdminOnlySecretAccess,omitempty"`
	IsOrgAdmin                      bool                     `protobuf:"varint,13,opt,name=isOrgAdmin,proto3" json:"isOrgAdmin,omitempty"`
}

func (x *GetUserAuthorizationResponse) Reset() {
	*x = GetUserAuthorizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_sentry_cluster_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAuthorizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAuthorizationResponse) ProtoMessage() {}

func (x *GetUserAuthorizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_sentry_cluster_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAuthorizationResponse.ProtoReflect.Descriptor instead.
func (*GetUserAuthorizationResponse) Descriptor() ([]byte, []int) {
	return file_proto_rpc_sentry_cluster_authz_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserAuthorizationResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetUserAuthorizationResponse) GetServiceAccount() *controller.StepObject {
	if x != nil {
		return x.ServiceAccount
	}
	return nil
}

func (x *GetUserAuthorizationResponse) GetClusterRoles() []*controller.StepObject {
	if x != nil {
		return x.ClusterRoles
	}
	return nil
}

func (x *GetUserAuthorizationResponse) GetClusterRoleBindings() []*controller.StepObject {
	if x != nil {
		return x.ClusterRoleBindings
	}
	return nil
}

func (x *GetUserAuthorizationResponse) GetRoles() []*controller.StepObject {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *GetUserAuthorizationResponse) GetRoleBindings() []*controller.StepObject {
	if x != nil {
		return x.RoleBindings
	}
	return nil
}

func (x *GetUserAuthorizationResponse) GetDeleteClusterRoleBindings() []*controller.StepObject {
	if x != nil {
		return x.DeleteClusterRoleBindings
	}
	return nil
}

func (x *GetUserAuthorizationResponse) GetDeleteRoleBindings() []*controller.StepObject {
	if x != nil {
		return x.DeleteRoleBindings
	}
	return nil
}

func (x *GetUserAuthorizationResponse) GetNamespaces() []*controller.StepObject {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *GetUserAuthorizationResponse) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *GetUserAuthorizationResponse) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

func (x *GetUserAuthorizationResponse) GetEnforceOrgAdminOnlySecretAccess() bool {
	if x != nil {
		return x.EnforceOrgAdminOnlySecretAccess
	}
	return false
}

func (x *GetUserAuthorizationResponse) GetIsOrgAdmin() bool {
	if x != nil {
		return x.IsOrgAdmin
	}
	return false
}

type GetUserAuthorizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCN           string `protobuf:"bytes,1,opt,name=userCN,proto3" json:"userCN,omitempty"`
	ClusterID        string `protobuf:"bytes,2,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	CertIssueSeconds int64  `protobuf:"varint,3,opt,name=certIssueSeconds,proto3" json:"certIssueSeconds,omitempty"`
}

func (x *GetUserAuthorizationRequest) Reset() {
	*x = GetUserAuthorizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_sentry_cluster_authz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAuthorizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAuthorizationRequest) ProtoMessage() {}

func (x *GetUserAuthorizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_sentry_cluster_authz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAuthorizationRequest.ProtoReflect.Descriptor instead.
func (*GetUserAuthorizationRequest) Descriptor() ([]byte, []int) {
	return file_proto_rpc_sentry_cluster_authz_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserAuthorizationRequest) GetUserCN() string {
	if x != nil {
		return x.UserCN
	}
	return ""
}

func (x *GetUserAuthorizationRequest) GetClusterID() string {
	if x != nil {
		return x.ClusterID
	}
	return ""
}

func (x *GetUserAuthorizationRequest) GetCertIssueSeconds() int64 {
	if x != nil {
		return x.CertIssueSeconds
	}
	return 0
}

var File_proto_rpc_sentry_cluster_authz_proto protoreflect.FileDescriptor

var file_proto_rpc_sentry_cluster_authz_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62,
	0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x2f, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xde, 0x06, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x4e, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x4a, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0c, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x58, 0x0a, 0x13, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x13, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3c, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x05, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x61, 0x66, 0x61,
	0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x64, 0x0a, 0x19, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x53, 0x74, 0x65, 0x70, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x19, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x56, 0x0a, 0x12, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x65, 0x70, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x12, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x46, 0x0a,
	0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x65, 0x70, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x52, 0x65, 0x61, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x69, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x48, 0x0a, 0x1f, 0x65, 0x6e, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x4f, 0x72, 0x67, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4f, 0x6e, 0x6c, 0x79,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x1f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x4f, 0x72, 0x67, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x4f, 0x6e, 0x6c, 0x79, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x4f, 0x72, 0x67, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x4f, 0x72, 0x67, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x22, 0x7f, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x43, 0x4e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x43, 0x4e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x65, 0x72, 0x74,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x63, 0x65, 0x72, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x32, 0xbd, 0x01, 0x0a, 0x14, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xa4, 0x01,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x72, 0x61, 0x66, 0x61,
	0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x42, 0xc7, 0x04, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x66,
	0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x70,
	0x63, 0x42, 0x11, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x66, 0x61, 0x79, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x72, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0xa2, 0x02, 0x04, 0x52, 0x44, 0x53, 0x52,
	0xaa, 0x02, 0x14, 0x52, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x44, 0x65, 0x76, 0x2e, 0x53, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x52, 0x70, 0x63, 0xca, 0x02, 0x14, 0x52, 0x61, 0x66, 0x61, 0x79, 0x5c,
	0x44, 0x65, 0x76, 0x5c, 0x53, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5c, 0x52, 0x70, 0x63, 0xe2, 0x02,
	0x20, 0x52, 0x61, 0x66, 0x61, 0x79, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x53, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x17, 0x52, 0x61, 0x66, 0x61, 0x79, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x3a, 0x3a,
	0x53, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x92, 0x41, 0xef, 0x02, 0x12,
	0x38, 0x0a, 0x24, 0x53, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x20, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x20, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x0b, 0x0a, 0x09, 0x52, 0x61, 0x66, 0x61, 0x79,
	0x20, 0x44, 0x65, 0x76, 0x32, 0x03, 0x32, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x32, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x79, 0x61, 0x6d, 0x6c,
	0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x79, 0x61, 0x6d, 0x6c, 0x52, 0x50, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x49, 0x0a, 0x47, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68,
	0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74,
	0x6f, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x3b, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x34, 0x0a,
	0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73,
	0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a,
	0x02, 0x01, 0x07, 0x5a, 0x38, 0x0a, 0x25, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x17, 0x08, 0x02, 0x1a, 0x11, 0x58, 0x2d, 0x52, 0x41, 0x46, 0x41, 0x59,
	0x2d, 0x41, 0x50, 0x49, 0x2d, 0x4b, 0x45, 0x59, 0x49, 0x44, 0x20, 0x02, 0x0a, 0x0f, 0x0a, 0x09,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x02, 0x08, 0x01, 0x62, 0x1f, 0x0a,
	0x0e, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x0a,
	0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rpc_sentry_cluster_authz_proto_rawDescOnce sync.Once
	file_proto_rpc_sentry_cluster_authz_proto_rawDescData = file_proto_rpc_sentry_cluster_authz_proto_rawDesc
)

func file_proto_rpc_sentry_cluster_authz_proto_rawDescGZIP() []byte {
	file_proto_rpc_sentry_cluster_authz_proto_rawDescOnce.Do(func() {
		file_proto_rpc_sentry_cluster_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rpc_sentry_cluster_authz_proto_rawDescData)
	})
	return file_proto_rpc_sentry_cluster_authz_proto_rawDescData
}

var file_proto_rpc_sentry_cluster_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_rpc_sentry_cluster_authz_proto_goTypes = []interface{}{
	(*GetUserAuthorizationResponse)(nil), // 0: rafay.dev.sentry.rpc.GetUserAuthorizationResponse
	(*GetUserAuthorizationRequest)(nil),  // 1: rafay.dev.sentry.rpc.GetUserAuthorizationRequest
	(*controller.StepObject)(nil),        // 2: rafay.dev.types.controller.StepObject
}
var file_proto_rpc_sentry_cluster_authz_proto_depIdxs = []int32{
	2, // 0: rafay.dev.sentry.rpc.GetUserAuthorizationResponse.serviceAccount:type_name -> rafay.dev.types.controller.StepObject
	2, // 1: rafay.dev.sentry.rpc.GetUserAuthorizationResponse.clusterRoles:type_name -> rafay.dev.types.controller.StepObject
	2, // 2: rafay.dev.sentry.rpc.GetUserAuthorizationResponse.clusterRoleBindings:type_name -> rafay.dev.types.controller.StepObject
	2, // 3: rafay.dev.sentry.rpc.GetUserAuthorizationResponse.roles:type_name -> rafay.dev.types.controller.StepObject
	2, // 4: rafay.dev.sentry.rpc.GetUserAuthorizationResponse.roleBindings:type_name -> rafay.dev.types.controller.StepObject
	2, // 5: rafay.dev.sentry.rpc.GetUserAuthorizationResponse.deleteClusterRoleBindings:type_name -> rafay.dev.types.controller.StepObject
	2, // 6: rafay.dev.sentry.rpc.GetUserAuthorizationResponse.deleteRoleBindings:type_name -> rafay.dev.types.controller.StepObject
	2, // 7: rafay.dev.sentry.rpc.GetUserAuthorizationResponse.namespaces:type_name -> rafay.dev.types.controller.StepObject
	1, // 8: rafay.dev.sentry.rpc.ClusterAuthorization.GetUserAuthorization:input_type -> rafay.dev.sentry.rpc.GetUserAuthorizationRequest
	0, // 9: rafay.dev.sentry.rpc.ClusterAuthorization.GetUserAuthorization:output_type -> rafay.dev.sentry.rpc.GetUserAuthorizationResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_proto_rpc_sentry_cluster_authz_proto_init() }
func file_proto_rpc_sentry_cluster_authz_proto_init() {
	if File_proto_rpc_sentry_cluster_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rpc_sentry_cluster_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAuthorizationResponse); i {
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
		file_proto_rpc_sentry_cluster_authz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAuthorizationRequest); i {
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
			RawDescriptor: file_proto_rpc_sentry_cluster_authz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rpc_sentry_cluster_authz_proto_goTypes,
		DependencyIndexes: file_proto_rpc_sentry_cluster_authz_proto_depIdxs,
		MessageInfos:      file_proto_rpc_sentry_cluster_authz_proto_msgTypes,
	}.Build()
	File_proto_rpc_sentry_cluster_authz_proto = out.File
	file_proto_rpc_sentry_cluster_authz_proto_rawDesc = nil
	file_proto_rpc_sentry_cluster_authz_proto_goTypes = nil
	file_proto_rpc_sentry_cluster_authz_proto_depIdxs = nil
}
