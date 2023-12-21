package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type SysCasbinRule struct {
	api.Api
}

// GetPage 获取SysCasbinRule列表
// @Summary 获取SysCasbinRule列表
// @Description 获取SysCasbinRule列表
// @Tags SysCasbinRule
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysCasbinRule}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-casbin-rule [get]
// @Security Bearer
func (e SysCasbinRule) GetPage(c *gin.Context) {
	req := dto.SysCasbinRuleGetPageReq{}
	s := service.SysCasbinRule{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.SysCasbinRule, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取SysCasbinRule失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取SysCasbinRule
// @Summary 获取SysCasbinRule
// @Description 获取SysCasbinRule
// @Tags SysCasbinRule
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SysCasbinRule} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-casbin-rule/{id} [get]
// @Security Bearer
func (e SysCasbinRule) Get(c *gin.Context) {
	req := dto.SysCasbinRuleGetReq{}
	s := service.SysCasbinRule{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SysCasbinRule

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取SysCasbinRule失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建SysCasbinRule
// @Summary 创建SysCasbinRule
// @Description 创建SysCasbinRule
// @Tags SysCasbinRule
// @Accept application/json
// @Product application/json
// @Param data body dto.SysCasbinRuleInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sys-casbin-rule [post]
// @Security Bearer
func (e SysCasbinRule) Insert(c *gin.Context) {
	req := dto.SysCasbinRuleInsertReq{}
	s := service.SysCasbinRule{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建SysCasbinRule失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改SysCasbinRule
// @Summary 修改SysCasbinRule
// @Description 修改SysCasbinRule
// @Tags SysCasbinRule
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysCasbinRuleUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-casbin-rule/{id} [put]
// @Security Bearer
func (e SysCasbinRule) Update(c *gin.Context) {
	req := dto.SysCasbinRuleUpdateReq{}
	s := service.SysCasbinRule{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改SysCasbinRule失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除SysCasbinRule
// @Summary 删除SysCasbinRule
// @Description 删除SysCasbinRule
// @Tags SysCasbinRule
// @Param data body dto.SysCasbinRuleDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-casbin-rule [delete]
// @Security Bearer
func (e SysCasbinRule) Delete(c *gin.Context) {
	s := service.SysCasbinRule{}
	req := dto.SysCasbinRuleDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除SysCasbinRule失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
