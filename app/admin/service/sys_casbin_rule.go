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

type SysCasbinRule struct {
	service.Service
}

// GetPage 获取SysCasbinRule列表
func (e *SysCasbinRule) GetPage(c *dto.SysCasbinRuleGetPageReq, p *actions.DataPermission, list *[]models.SysCasbinRule, count *int64) error {
	var err error
	var data models.SysCasbinRule

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SysCasbinRuleService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SysCasbinRule对象
func (e *SysCasbinRule) Get(d *dto.SysCasbinRuleGetReq, p *actions.DataPermission, model *models.SysCasbinRule) error {
	var data models.SysCasbinRule

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysCasbinRule error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysCasbinRule对象
func (e *SysCasbinRule) Insert(c *dto.SysCasbinRuleInsertReq) error {
	var err error
	var data models.SysCasbinRule
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SysCasbinRuleService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SysCasbinRule对象
func (e *SysCasbinRule) Update(c *dto.SysCasbinRuleUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.SysCasbinRule{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("SysCasbinRuleService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SysCasbinRule
func (e *SysCasbinRule) Remove(d *dto.SysCasbinRuleDeleteReq, p *actions.DataPermission) error {
	var data models.SysCasbinRule

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSysCasbinRule error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
