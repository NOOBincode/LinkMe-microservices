package service

import (
	"context"
	pb "github.com/GoSimplicity/LinkMe-microservices/api/role/v1"
)

type RoleServiceService struct {
	pb.UnimplementedRoleServiceServer
}

func NewRoleServiceService() *RoleServiceService {
	return &RoleServiceService{}
}

func (s *RoleServiceService) GetPermissions(ctx context.Context, req *pb.google_protobuf_Empty) (*pb.GetPermissionsReply, error) {
	return &pb.GetPermissionsReply{}, nil
}
func (s *RoleServiceService) AssignPermission(ctx context.Context, req *pb.AssignPermissionRequest) (*pb.google_protobuf_Empty, error) {
	return &pb.google_protobuf_Empty{}, nil
}
func (s *RoleServiceService) AssignRoleToUser(ctx context.Context, req *pb.AssignRoleToUserRequest) (*pb.google_protobuf_Empty, error) {
	return &pb.google_protobuf_Empty{}, nil
}
func (s *RoleServiceService) RemovePermission(ctx context.Context, req *pb.RemovePermissionRequest) (*pb.google_protobuf_Empty, error) {
	return &pb.google_protobuf_Empty{}, nil
}
func (s *RoleServiceService) RemoveRoleFromUser(ctx context.Context, req *pb.RemoveRoleFromUserRequest) (*pb.google_protobuf_Empty, error) {
	return &pb.google_protobuf_Empty{}, nil
}
