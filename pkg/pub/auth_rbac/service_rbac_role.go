/**
* @author long.qian
*/

package auth_rbac

// =================================================================================
// You can edit this file about RbacRole service code.
// eg: func (*rbacRoleService) FuncName(d dto.XxxDto) error {}
// =================================================================================

import (
	"errors"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/db"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/query"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/update"
)

type rbacRoleService struct {}

var (
	RbacRoleService = new(rbacRoleService)
)

// CreateRbacRole 创建
func (*rbacRoleService) CreateRbacRole(d RbacRoleCreateDto) (*RbacRole, error) {
	err := db.DB.Create(&d).Error
	if err != nil {
		return nil, errors.New("create failed：" + err.Error())
	}
	return d.TransformTo(), nil
}

// UpdateRbacRole 更新
func (*rbacRoleService) UpdateRbacRole(d RbacRoleUpdateDto) error {
	if d.Id == 0 {
		return errors.New("update failed：" + "id is Empty")
	}
	err := db.DB.Model(&RbacRole{}).Where("id = ?", d.Id).Updates(update.GenerateUpdatesMap(d)).Error
	if err != nil {
		return errors.New("update failed：" + err.Error())
	}
	return nil
}

// DeleteSoftById 软删除（要查询被软删除的记录，用Unscoped，例如：db.DB.Unscoped().Where("name = ?", "xxx").Find(&users)）
func (*rbacRoleService) DeleteSoftById(id uint) error {
	return db.DB.Where("id = ?", id).Delete(&RbacRole{}).Error
}

// DeleteUnscopedById 永久删除
func (*rbacRoleService) DeleteUnscopedById(id uint) error {
	return db.DB.Unscoped().Where("id = ?", id).Delete(&RbacRole{}).Error
}

// QueryPageRbacRole 分页查询
func (*rbacRoleService) QueryPageRbacRole(d RbacRolePageDto) (*RbacRolePageVo, error) {
	tx := db.DB.Model(&RbacRole{})
	whereErr := query.WhereHandler(tx, d)
	if whereErr != nil {
		return nil, whereErr
	}
	orderErr := query.OrderHandler(tx, d)
	if orderErr != nil {
		return nil, orderErr
	}
	var r []RbacRole
	var total int64
	tx.Count(&total)
	query.PageHandler(tx, &d.Page)
	tx.Find(&r)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	var voList []RbacRoleVo
	for _, m := range r {
		v := RbacRoleVo{}
		voList = append(voList, v.Transform(m))
	}

	return &RbacRolePageVo{
		Page: query.Page{
			PageNum: d.PageNum,
			PageSize: d.PageSize,
			Total: total,
		},
		Result: voList,
	}, nil
}

// QueryRbacRole 查询
func (*rbacRoleService) QueryRbacRole(d RbacRoleQueryDto) ([]RbacRoleVo, error) {
	tx := db.DB.Model(&RbacRole{})
	whereErr := query.WhereHandler(tx, d)
	if whereErr != nil {
		return nil, whereErr
	}
	orderErr := query.OrderHandler(tx, d)
	if orderErr != nil {
		return nil, orderErr
	}
	var r []RbacRole
	tx.Find(&r)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	var voList []RbacRoleVo
	for _, m := range r {
		v := RbacRoleVo{}
		voList = append(voList, v.Transform(m))
	}

	return voList, nil
}

// FindById 根据 ID 查询，返回实体
func (*rbacRoleService) FindById(id uint) (RbacRole, error) {
	m := RbacRole{}
	tx := db.DB.Model(&m).Where("id = ?", id).First(&m)
	if m.ID == 0 {
		return m, errors.New("记录不存在")
	}
	return m, tx.Error
}

// FindById4Vo 根据 ID 查询，返回实体 VO 包装结构体
func (*rbacRoleService) FindById4Vo(id uint) (RbacRoleVo, error) {
	m := RbacRole{}
	tx := db.DB.Model(&m).Where("id = ?", id).First(&m)
	if m.ID == 0 {
		return RbacRoleVo{}, errors.New("记录不存在")
	}
	v := RbacRoleVo{}
	return v.Transform(m), tx.Error
}

// FindByRoleName 根据 RoleName 字段查询，返回实体
func (*rbacRoleService) FindByRoleName(roleName string) (RbacRole, error) {
	m := RbacRole{}
	tx := db.DB.Model(&m).Where("role_name = ?", roleName).First(&m)
	if m.ID == 0 {
		return m, errors.New("记录不存在")
	}
	return m, tx.Error
}

// FindByRoleName4Vo 根据 RoleName 字段查询，返回 VO 包装结构体
func (*rbacRoleService) FindByRoleName4Vo(roleName string) (RbacRoleVo, error) {
	m := RbacRole{}
	tx := db.DB.Model(&m).Where("role_name = ?", roleName).First(&m)
	if m.ID == 0 {
		return RbacRoleVo{}, errors.New("记录不存在")
	}
	v := RbacRoleVo{}
	return v.Transform(m), tx.Error
}

