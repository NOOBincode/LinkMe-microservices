package data

import (
	"context"
	"errors"
	"fmt"
	userpb "github.com/GoSimplicity/LinkMe-microservices/api/user/v1"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-role/domain"
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/transport"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// Permission 定义了权限仓库接口
type Permission interface {
	GetPermissions(ctx context.Context) ([]domain.Role, error)                               // 获取权限列表
	AssignPermission(ctx context.Context, userName string, path string, method string) error // 分配权限
	AssignRoleToUser(ctx context.Context, userName, roleName string) error
	RemovePermission(ctx context.Context, userName string, path string, method string) error // 移除权限
	RemoveRoleFromUser(ctx context.Context, userName, roleName string) error
}

type RoleUseCase struct {
	db         *gorm.DB
	l          *zap.Logger
	ce         *casbin.Enforcer
	p          Permission
	userClient userpb.UserClient
}

func NewRoleUseCase(db *gorm.DB, l *zap.Logger, ce *casbin.Enforcer, p Permission) *RoleUseCase {
	return &RoleUseCase{db: db, l: l, ce: ce, p: p}
}

func (s *RoleUseCase) GetPermissions(ctx context.Context) ([]domain.Role, error) {
	var policies []domain.Role
	err := s.db.WithContext(ctx).Table("casbin_rule").Find(&policies).Error
	if err != nil {
		return nil, err
	}
	return policies, nil
}
func (s *RoleUseCase) AssignPermission(ctx context.Context, userName string, path string, method string) error {
	userID, err := s.getUserId(ctx)
	if err != nil {
		return err
	}
	userIDStr := strconv.FormatInt(userID, 10)
	ok, err := s.ce.AddPolicy(userIDStr, path, method)
	if err != nil {
		s.l.Error("Failed to add policy", zap.Error(err))
		return err
	}
	if !ok {
		s.l.Error("Failed to add policy", zap.Error(err))
		return fmt.Errorf("policy already exists for user %d,path %s,method %s", userID, path, method)
	}
	return nil
}

func (s *RoleUseCase) AssignRoleToUser(ctx context.Context, userName, roleName string) error {
	userID, err := s.getUserId(ctx)
	if err != nil {
		s.l.Error("Failed to get user id", zap.Error(err))
		return err
	}
	userIDStr := strconv.FormatInt(userID, 10)
	roleIDStr := strconv.FormatInt(userID, 10)
	ok, err := s.ce.AddGroupingPolicy(userIDStr, roleIDStr)
	if err != nil {
		s.l.Error("Failed to add role to user", zap.Error(err))
		return err
	}
	if !ok {
		s.l.Error("role already assigned to user", zap.Error(err))
		return fmt.Errorf("role %s already assigned to user %s", roleName, userName)
	}
	return nil
}

func (s *RoleUseCase) RemovePermission(ctx context.Context, userName string, path string, method string) error {
	userID, err := s.getUserId(ctx)
	if err != nil {
		s.l.Error("RemovePermission:Failed to get user id", zap.Error(err))
		return err
	}
	userIDStr := strconv.FormatInt(userID, 10)
	ok, err := s.ce.RemovePolicy(userIDStr, path, method)
	if err != nil {
		s.l.Error("Failed to remove policy", zap.Error(err))
		return err
	}
	if !ok {
		s.l.Error("policy does not exist", zap.Error(err))
		return fmt.Errorf("policy does not exist for user %d,path %s,method %s", userID, path, method)
	}
	return nil
}

func (s *RoleUseCase) RemoveRoleFromUser(ctx context.Context, userName, roleName string) error {
	userID, err := s.getUserId(ctx)
	if err != nil {
		s.l.Error("RemoveRoleFromUser:Failed to get user id", zap.Error(err))
		return err
	}
	roleID, err := s.getUserId(ctx)
	if err != nil {
		s.l.Error("RemoveRoleFromUser:Failed to get user id", zap.Error(err))
		return err
	}

	userIDStr := strconv.FormatInt(userID, 10)
	roleIDStr := strconv.FormatInt(roleID, 10)
	ok, err := s.ce.RemoveGroupingPolicy(userIDStr, roleIDStr)
	if err != nil {
		s.l.Error("RemoveRoleFromUser:Failed to remove role from user", zap.Error(err))
		return err
	}
	if !ok {
		s.l.Error("RemoveRoleFromUser:role not assigned to user", zap.Error(err))
		return fmt.Errorf("role %s not assigned to user %s", roleName, userName)
	}
	return nil
}

func (s *RoleUseCase) getUserId(ctx context.Context) (int64, error) {
	// 从 kratos 上下文中获取传输信息
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return -1, errors.New("failed to get transport from context")
	}
	// 获取 Authorization 头
	token := tr.RequestHeader().Get("Authorization")
	if token == "" {
		return -1, errors.New("authorization token not provided")
	}
	// 移除 "Bearer " 前缀
	tokenStr := strings.TrimPrefix(token, "Bearer ")
	if tokenStr == "" {
		return -1, errors.New("authorization token is empty after trim")
	}
	// 调用 userClient 获取用户信息
	info, err := s.userClient.GetUserInfo(ctx, &userpb.GetUserInfoRequest{Token: tokenStr})
	if err != nil {
		return -1, err
	}
	return info.UserId, nil
}
