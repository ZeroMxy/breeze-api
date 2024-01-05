package middleware

import (
	"breeze-api/helper"
	"breeze-api/internal/service"
	"breeze-api/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// 中间件
type Admin struct{}

func (*Admin) Handle(ctx *fiber.Ctx) error {

	id, err := helper.GetTokenPayload(ctx)
	if err != nil {
		return response.Base(ctx, 10401, err.Error(), nil)
	}

	// 获取角色
	admin := (&service.Admin{}).GetDetail(id)
	// 超级管理员直接通过
	if admin.Id == 1 {
		return ctx.Next()
	}
	// 非超级管理员检查角色
	adminRoles := (&service.AdminRoleRelation{}).GetList(admin.Id)
	if len(adminRoles) <= 0 {
		return response.Base(ctx, 10403, "无权限", nil)
	}

	// 获取访问权限
	permission := (&service.Permission{}).GetDetailByPathWithMethod(ctx.Path(), ctx.Method())
	if permission.Id <= 0 || permission.Status != 1 {
		return response.Base(ctx, 10403, "无权限", nil)
	}

	// 绑定权限检查
	for _, adminRole := range adminRoles {
		rolePermission := (&service.RolePermissionRelation{}).GetDetailByRoleIdWithPermissionId(adminRole.RoleId, permission.Id)
		if rolePermission.Id > 0 {
			return ctx.Next()
		}
	}

	return response.Base(ctx, 10403, "无权限", nil)
}
