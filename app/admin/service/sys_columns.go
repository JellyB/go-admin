package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type SysColumns struct {
	service.Service
}

// GetPage 获取SysColumns列表
func (e *SysColumns) GetPage(c *dto.SysColumnsGetPageReq, p *actions.DataPermission, list *[]models.SysColumns, count *int64) error {
	var err error
	var data models.SysColumns

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SysColumnsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SysColumns对象
func (e *SysColumns) Get(d *dto.SysColumnsGetReq, p *actions.DataPermission, model *models.SysColumns) error {
	var data models.SysColumns

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysColumns error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysColumns对象
func (e *SysColumns) Insert(c *dto.SysColumnsInsertReq) error {
	var err error
	var data models.SysColumns
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SysColumnsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SysColumns对象
func (e *SysColumns) Update(c *dto.SysColumnsUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.SysColumns{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("SysColumnsService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SysColumns
func (e *SysColumns) Remove(d *dto.SysColumnsDeleteReq, p *actions.DataPermission) error {
	var data models.SysColumns

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSysColumns error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
