package models

import (
	"go-admin/common/models"
)

type SysCasbinRule struct {
	models.Model

	models.ModelTime
	models.ControlBy
}

func (SysCasbinRule) TableName() string {
	return "sys_casbin_rule"
}

func (e *SysCasbinRule) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysCasbinRule) GetId() interface{} {
	return e.Id
}
