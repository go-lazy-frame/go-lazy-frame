/**
* @author long.qian
*/

package auth_rbac

// =================================================================================
// You can edit this file about RbacLog service code.
// eg: func (*rbacLogService) FuncName(d dto.XxxDto) error {}
// =================================================================================

import (
	"errors"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/db"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/query"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/update"
)

type rbacLogService struct {}

var (
	RbacLogService = new(rbacLogService)
)

// CreateRbacLog 创建
func (*rbacLogService) CreateRbacLog(d RbacLogCreateDto) (*RbacLogCreateDto, error) {
	err := db.DB.Create(&d).Error
	if err != nil {
		return nil, errors.New("create failed：" + err.Error())
	}
	return &d, nil
}

// UpdateRbacLog 更新
func (*rbacLogService) UpdateRbacLog(d RbacLogUpdateDto) error {
	if d.Id == 0 {
		return errors.New("update failed：" + "id is Empty")
	}
	err := db.DB.Model(&RbacLog{}).Where("id = ?", d.Id).Updates(update.GenerateUpdatesMap(d)).Error
	if err != nil {
		return errors.New("update failed：" + err.Error())
	}
	return nil
}

// DeleteSoftById 软删除（要查询被软删除的记录，用Unscoped，例如：db.DB.Unscoped().Where("name = ?", "xxx").Find(&users)）
func (*rbacLogService) DeleteSoftById(id uint) error {
	return db.DB.Where("id = ?", id).Delete(&RbacLog{}).Error
}

// DeleteUnscopedById 永久删除
func (*rbacLogService) DeleteUnscopedById(id uint) error {
	return db.DB.Unscoped().Where("id = ?", id).Delete(&RbacLog{}).Error
}

// QueryPageRbacLog 分页查询
func (*rbacLogService) QueryPageRbacLog(d RbacLogPageDto) (*RbacLogPageVo, error) {
	tx := db.DB.Model(&RbacLog{})
	whereErr := query.WhereHandler(tx, d)
	if whereErr != nil {
		return nil, whereErr
	}
	orderErr := query.OrderHandler(tx, d)
	if orderErr != nil {
		return nil, orderErr
	}
	var r []RbacLog
	var total int64
	tx.Count(&total)
	query.PageHandler(tx, &d.Page)
	tx.Find(&r)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	var voList []RbacLogVo
	for _, m := range r {
		v := RbacLogVo{}
		voList = append(voList, v.Transform(m))
	}

	return &RbacLogPageVo{
		Page: query.Page{
			PageNum: d.PageNum,
			PageSize: d.PageSize,
			Total: total,
		},
		Result: voList,
	}, nil
}

// QueryRbacLog 查询
func (*rbacLogService) QueryRbacLog(d RbacLogQueryDto) ([]RbacLogVo, error) {
	tx := db.DB.Model(&RbacLog{})
	whereErr := query.WhereHandler(tx, d)
	if whereErr != nil {
		return nil, whereErr
	}
	orderErr := query.OrderHandler(tx, d)
	if orderErr != nil {
		return nil, orderErr
	}
	var r []RbacLog
	tx.Find(&r)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	var voList []RbacLogVo
	for _, m := range r {
		v := RbacLogVo{}
		voList = append(voList, v.Transform(m))
	}

	return voList, nil
}

// FindById 根据 ID 查询，返回实体
func (*rbacLogService) FindById(id uint) (RbacLog, error) {
	m := RbacLog{}
	tx := db.DB.Model(&m).Where("id = ?", id).First(&m)
	if m.ID == 0 {
		return m, errors.New("记录不存在")
	}
	return m, tx.Error
}

// FindById4Vo 根据 ID 查询，返回实体 VO 包装结构体
func (*rbacLogService) FindById4Vo(id uint) (RbacLogVo, error) {
	m := RbacLog{}
	tx := db.DB.Model(&m).Where("id = ?", id).First(&m)
	if m.ID == 0 {
		return RbacLogVo{}, errors.New("记录不存在")
	}
	v := RbacLogVo{}
	return v.Transform(m), tx.Error
}

