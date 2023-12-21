package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysCasbinRuleGetPageReq struct {
	dto.Pagination `search:"-"`
	SysCasbinRuleOrder
}

type SysCasbinRuleOrder struct {
}

func (m *SysCasbinRuleGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysCasbinRuleInsertReq struct {
	common.ControlBy
}

func (s *SysCasbinRuleInsertReq) Generate(model *models.SysCasbinRule) {
}

func (s *SysCasbinRuleInsertReq) GetId() interface{} {
	return s.Id
}

type SysCasbinRuleUpdateReq struct {
	common.ControlBy
}

func (s *SysCasbinRuleUpdateReq) Generate(model *models.SysCasbinRule) {
}

func (s *SysCasbinRuleUpdateReq) GetId() interface{} {
	return s.Id
}

// SysCasbinRuleGetReq 功能获取请求参数
type SysCasbinRuleGetReq struct {
}

func (s *SysCasbinRuleGetReq) GetId() interface{} {
	return s.Id
}

// SysCasbinRuleDeleteReq 功能删除请求参数
type SysCasbinRuleDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysCasbinRuleDeleteReq) GetId() interface{} {
	return s.Ids
}
