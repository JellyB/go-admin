package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysColumnsGetPageReq struct {
	dto.Pagination `search:"-"`
	SysColumnsOrder
}

type SysColumnsOrder struct {
}

func (m *SysColumnsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysColumnsInsertReq struct {
	common.ControlBy
}

func (s *SysColumnsInsertReq) Generate(model *models.SysColumns) {
}

func (s *SysColumnsInsertReq) GetId() interface{} {
	return s.ColumnId
}

type SysColumnsUpdateReq struct {
	common.ControlBy
}

func (s *SysColumnsUpdateReq) Generate(model *models.SysColumns) {
}

func (s *SysColumnsUpdateReq) GetId() interface{} {
	return s.ColumnId
}

// SysColumnsGetReq 功能获取请求参数
type SysColumnsGetReq struct {
}

func (s *SysColumnsGetReq) GetId() interface{} {
	return s.ColumnId
}

// SysColumnsDeleteReq 功能删除请求参数
type SysColumnsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysColumnsDeleteReq) GetId() interface{} {
	return s.Ids
}
