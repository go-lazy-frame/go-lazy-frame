/**
* @author long.qian
*/

package auth_rbac

// =================================================================================
// You can edit this file about RbacPermissions service code.
// eg: func (*rbacPermissionsService) FuncName(d dto.XxxDto) error {}
// =================================================================================

import (
	"errors"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/db"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/query"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/update"
)

type rbacPermissionsService struct {}

var (
	RbacPermissionsService = new(rbacPermissionsService)
)

// CreateRbacPermissions 创建
func (*rbacPermissionsService) CreateRbacPermissions(d RbacPermissionsCreateDto) (*RbacPermissions, error) {
	err := db.DB.Create(&d).Error
	if err != nil {
		return nil, errors.New("create failed：" + err.Error())
	}
	return d.TransformTo(), nil
}

// UpdateRbacPermissions 更新
func (*rbacPermissionsService) UpdateRbacPermissions(d RbacPermissionsUpdateDto) error {
	if d.Id == 0 {
		return errors.New("update failed：" + "id is Empty")
	}
	err := db.DB.Model(&RbacPermissions{}).Where("id = ?", d.Id).Updates(update.GenerateUpdatesMap(d)).Error
	if err != nil {
		return errors.New("update failed：" + err.Error())
	}
	return nil
}

// DeleteSoftById 软删除（要查询被软删除的记录，用Unscoped，例如：db.DB.Unscoped().Where("name = ?", "xxx").Find(&users)）
func (*rbacPermissionsService) DeleteSoftById(id uint) error {
	return db.DB.Where("id = ?", id).Delete(&RbacPermissions{}).Error
}

// DeleteUnscopedById 永久删除
func (*rbacPermissionsService) DeleteUnscopedById(id uint) error {
	return db.DB.Unscoped().Where("id = ?", id).Delete(&RbacPermissions{}).Error
}

// QueryPageRbacPermissions 分页查询
func (*rbacPermissionsService) QueryPageRbacPermissions(d RbacPermissionsPageDto) (*RbacPermissionsPageVo, error) {
	tx := db.DB.Model(&RbacPermissions{})
	whereErr := query.WhereHandler(tx, d)
	if whereErr != nil {
		return nil, whereErr
	}
	orderErr := query.OrderHandler(tx, d)
	if orderErr != nil {
		return nil, orderErr
	}
	var r []RbacPermissions
	var total int64
	tx.Count(&total)
	query.PageHandler(tx, &d.Page)
	tx.Find(&r)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	var voList []RbacPermissionsVo
	for _, m := range r {
		v := RbacPermissionsVo{}
		voList = append(voList, v.Transform(m))
	}

	return &RbacPermissionsPageVo{
		Page: query.Page{
			PageNum: d.PageNum,
			PageSize: d.PageSize,
			Total: total,
		},
		Result: voList,
	}, nil
}

// QueryRbacPermissions 查询
func (*rbacPermissionsService) QueryRbacPermissions(d RbacPermissionsQueryDto) ([]RbacPermissionsVo, error) {
	tx := db.DB.Model(&RbacPermissions{})
	whereErr := query.WhereHandler(tx, d)
	if whereErr != nil {
		return nil, whereErr
	}
	orderErr := query.OrderHandler(tx, d)
	if orderErr != nil {
		return nil, orderErr
	}
	var r []RbacPermissions
	tx.Find(&r)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	var voList []RbacPermissionsVo
	for _, m := range r {
		v := RbacPermissionsVo{}
		voList = append(voList, v.Transform(m))
	}

	return voList, nil
}

// FindById 根据 ID 查询，返回实体
func (*rbacPermissionsService) FindById(id uint) (RbacPermissions, error) {
	m := RbacPermissions{}
	tx := db.DB.Model(&m).Where("id = ?", id).First(&m)
	if m.ID == 0 {
		return m, errors.New("记录不存在")
	}
	return m, tx.Error
}

// FindById4Vo 根据 ID 查询，返回实体 VO 包装结构体
func (*rbacPermissionsService) FindById4Vo(id uint) (RbacPermissionsVo, error) {
	m := RbacPermissions{}
	tx := db.DB.Model(&m).Where("id = ?", id).First(&m)
	if m.ID == 0 {
		return RbacPermissionsVo{}, errors.New("记录不存在")
	}
	v := RbacPermissionsVo{}
	return v.Transform(m), tx.Error
}

// FindByPermission 根据 Permission 字段查询，返回实体
func (*rbacPermissionsService) FindByPermission(permission string) (RbacPermissions, error) {
	m := RbacPermissions{}
	tx := db.DB.Model(&m).Where("permission = ?", permission).First(&m)
	if m.ID == 0 {
		return m, errors.New("记录不存在")
	}
	return m, tx.Error
}

// FindByPermission4Vo 根据 Permission 字段查询，返回 VO 包装结构体
func (*rbacPermissionsService) FindByPermission4Vo(permission string) (RbacPermissionsVo, error) {
	m := RbacPermissions{}
	tx := db.DB.Model(&m).Where("permission = ?", permission).First(&m)
	if m.ID == 0 {
		return RbacPermissionsVo{}, errors.New("记录不存在")
	}
	v := RbacPermissionsVo{}
	return v.Transform(m), tx.Error
}

