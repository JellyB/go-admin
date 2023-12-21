package models

import (
	"go-admin/common/models"
)

type SysColumns struct {
	models.Model

	models.ModelTime
	models.ControlBy
}

func (SysColumns) TableName() string {
	return "sys_columns"
}

func (e *SysColumns) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysColumns) GetId() interface{} {
	return e.ColumnId
}
