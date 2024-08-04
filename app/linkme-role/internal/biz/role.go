package biz

import (
	"context"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-role/domain"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-role/internal/data"
	"go.uber.org/zap"
)

type PermissionData interface {
	GetPermissions(ctx context.Context) ([]domain.Role, error)                               // 获取权限列表
	AssignPermission(ctx context.Context, userName string, path string, method string) error // 分配权限
	AssignRoleToUser(ctx context.Context, userName, roleName string) error
	RemovePermission(ctx context.Context, userName string, path string, method string) error // 移除权限
	RemoveRoleFromUser(ctx context.Context, userName, roleName string) error
}

// permissionData 是 PermissionRepository 的实现
type permissionData struct {
	l    *zap.Logger     // 日志记录器
	data data.Permission // 数据库访问对象
}

func (p permissionData) GetPermissions(ctx context.Context) ([]domain.Role, error) {
	permission, err := p.data.GetPermissions(ctx)
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func (p permissionData) AssignPermission(ctx context.Context, userName string, path string, method string) error {
	if err := p.data.AssignPermission(ctx, userName, path, method); err != nil {
		return err
	}
	return nil
}

func (p permissionData) AssignRoleToUser(ctx context.Context, userName, roleName string) error {
	if err := p.data.AssignRoleToUser(ctx, userName, roleName); err != nil {
		return err
	}
	return nil
}

func (p permissionData) RemovePermission(ctx context.Context, userName string, path string, method string) error {
	if err := p.data.RemovePermission(ctx, userName, path, method); err != nil {
		return err
	}
	return nil
}

func (p permissionData) RemoveRoleFromUser(ctx context.Context, userName, roleName string) error {
	if err := p.data.RemoveRoleFromUser(ctx, userName, roleName); err != nil {
		return err
	}
	return nil
}

// NewPermissionData 创建一个新的 PermissionRepository
func NewPermissionData(l *zap.Logger, data data.Permission) PermissionData {
	return &permissionData{
		l:    l,
		data: data,
	}
}
