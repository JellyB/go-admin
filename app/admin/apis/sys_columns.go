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

type SysColumns struct {
	api.Api
}

// GetPage 获取SysColumns列表
// @Summary 获取SysColumns列表
// @Description 获取SysColumns列表
// @Tags SysColumns
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysColumns}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-columns [get]
// @Security Bearer
func (e SysColumns) GetPage(c *gin.Context) {
	req := dto.SysColumnsGetPageReq{}
	s := service.SysColumns{}
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
	list := make([]models.SysColumns, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取SysColumns失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取SysColumns
// @Summary 获取SysColumns
// @Description 获取SysColumns
// @Tags SysColumns
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SysColumns} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-columns/{id} [get]
// @Security Bearer
func (e SysColumns) Get(c *gin.Context) {
	req := dto.SysColumnsGetReq{}
	s := service.SysColumns{}
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
	var object models.SysColumns

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取SysColumns失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建SysColumns
// @Summary 创建SysColumns
// @Description 创建SysColumns
// @Tags SysColumns
// @Accept application/json
// @Product application/json
// @Param data body dto.SysColumnsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sys-columns [post]
// @Security Bearer
func (e SysColumns) Insert(c *gin.Context) {
	req := dto.SysColumnsInsertReq{}
	s := service.SysColumns{}
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
		e.Error(500, err, fmt.Sprintf("创建SysColumns失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改SysColumns
// @Summary 修改SysColumns
// @Description 修改SysColumns
// @Tags SysColumns
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysColumnsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-columns/{id} [put]
// @Security Bearer
func (e SysColumns) Update(c *gin.Context) {
	req := dto.SysColumnsUpdateReq{}
	s := service.SysColumns{}
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
		e.Error(500, err, fmt.Sprintf("修改SysColumns失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除SysColumns
// @Summary 删除SysColumns
// @Description 删除SysColumns
// @Tags SysColumns
// @Param data body dto.SysColumnsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-columns [delete]
// @Security Bearer
func (e SysColumns) Delete(c *gin.Context) {
	s := service.SysColumns{}
	req := dto.SysColumnsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除SysColumns失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
