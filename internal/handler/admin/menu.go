package admin

import (
	"breeze-api/internal/model"
	"breeze-api/internal/service"
	"breeze-api/pkg/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// 菜单请求
type Menu struct{}

// 创建菜单
func (*Menu) Create(ctx *fiber.Ctx) error {

	type request struct {
		ParentId  int    `json:"parentId"`
		Name      string `json:"name"`
		Type      int    `json:"type"`
		Sort      int    `json:"sort"`
		Path      string `json:"path"`
		Component string `json:"component"`
		Icon      string `json:"icon"`
		Redirect  string `json:"redirect"`
		Status    string `json:"status"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Name == "" || req.Type <= 0 || req.Path == "" {
		return response.Error(ctx, "参数错误")
	}

	err := (&service.Menu{}).Create(&model.Menu{
		ParentId:  req.ParentId,
		Name:      req.Name,
		Type:      req.Type,
		Sort:      req.Sort,
		Path:      req.Path,
		Component: req.Component,
		Icon:      req.Icon,
		Redirect:  req.Redirect,
		Status:    req.Status,
	})

	if err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 更新菜单
func (*Menu) Update(ctx *fiber.Ctx) error {

	type request struct {
		Id        int    `json:"id"`
		ParentId  int    `json:"parentId"`
		Name      string `json:"name"`
		Type      int    `json:"type"`
		Sort      int    `json:"sort"`
		Path      string `json:"path"`
		Component string `json:"component"`
		Icon      string `json:"icon"`
		Redirect  string `json:"redirect"`
		Status    string `json:"status"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	err := (&service.Menu{}).Update(&model.Menu{
		Id:        req.Id,
		ParentId:  req.ParentId,
		Name:      req.Name,
		Type:      req.Type,
		Sort:      req.Sort,
		Path:      req.Path,
		Component: req.Component,
		Icon:      req.Icon,
		Redirect:  req.Redirect,
		Status:    req.Status,
	})

	if err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 删除菜单
func (*Menu) Delete(ctx *fiber.Ctx) error {

	type request struct {
		Id int `json:"id"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	list := (&service.Menu{}).GetChildrenList(req.Id, 0)
	if len(list) > 0 {
		return response.Error(ctx, "存在下级菜单")
	}

	err := (&service.Menu{}).Delete(req.Id)
	if err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 菜单列表
func (*Menu) Tree(ctx *fiber.Ctx) error {

	tree := (&service.Menu{}).GetChildrenList(0, 0)

	return response.Success(ctx, "成功", map[string]interface{}{
		"tree": tree,
	})
}

// 菜单详情
func (*Menu) Detail(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Query("id"))
	if id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	menu := (&service.Menu{}).GetDetail(id)

	return response.Success(ctx, "成功", map[string]interface{}{
		"menu": menu,
	})
}
