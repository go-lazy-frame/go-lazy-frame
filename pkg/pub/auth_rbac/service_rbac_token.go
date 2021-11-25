/**
* @author long.qian
 */

package auth_rbac

import (
	"errors"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/db"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/query"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/update"
	"github.com/google/uuid"
)

// =================================================================================
// You can edit this file about RbacToken service code.
// eg: func (*rbacTokenService) FuncName(d dto.XxxDto) error {}
// =================================================================================

type rbacTokenService struct{}

var (
	RbacTokenService = new(rbacTokenService)
)

// CreateRbacToken 创建
func (*rbacTokenService) CreateRbacToken(d RbacTokenCreateDto) (*RbacTokenCreateDto, error) {
	err := db.DB.Create(&d).Error
	if err != nil {
		return nil, errors.New("create failed：" + err.Error())
	}
	return &d, nil
}

// UpdateRbacToken 更新
func (*rbacTokenService) UpdateRbacToken(d RbacTokenUpdateDto) error {
	if d.Id == 0 {
		return errors.New("update failed：" + "id is Empty")
	}
	err := db.DB.Model(&RbacToken{}).Where("id = ?", d.Id).Updates(update.GenerateUpdatesMap(d)).Error
	if err != nil {
		return errors.New("update failed：" + err.Error())
	}
	return nil
}

// DeleteSoftById 软删除（要查询被软删除的记录，用Unscoped，例如：db.DB.Unscoped().Where("name = ?", "xxx").Find(&users)）
func (*rbacTokenService) DeleteSoftById(id uint) error {
	return db.DB.Where("id = ?", id).Delete(&RbacToken{}).Error
}

// DeleteUnscopedById 永久删除
func (*rbacTokenService) DeleteUnscopedById(id uint) error {
	return db.DB.Unscoped().Where("id = ?", id).Delete(&RbacToken{}).Error
}

// QueryPageRbacToken 分页查询
func (*rbacTokenService) QueryPageRbacToken(d RbacTokenPageDto) (*RbacTokenPageVo, error) {
	tx := db.DB.Model(&RbacToken{})
	whereErr := query.WhereHandler(tx, d)
	if whereErr != nil {
		return nil, whereErr
	}
	orderErr := query.OrderHandler(tx, d)
	if orderErr != nil {
		return nil, orderErr
	}
	var r []RbacToken
	var total int64
	tx.Count(&total)
	query.PageHandler(tx, &d.Page)
	tx.Find(&r)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	var voList []RbacTokenVo
	for _, m := range r {
		vo := RbacTokenVo{}
		voList = append(voList, vo.Transform(m))
	}

	return &RbacTokenPageVo{
		Page: query.Page{
			PageNum: d.PageNum,
			PageSize: d.PageSize,
			Total: total,
		},
		Result: voList,
	}, nil
}

// QueryRbacToken 查询
func (*rbacTokenService) QueryRbacToken(d RbacTokenQueryDto) ([]RbacTokenVo, error) {
	tx := db.DB.Model(&RbacToken{})
	whereErr := query.WhereHandler(tx, d)
	if whereErr != nil {
		return nil, whereErr
	}
	orderErr := query.OrderHandler(tx, d)
	if orderErr != nil {
		return nil, orderErr
	}
	var r []RbacToken
	tx.Find(&r)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	var voList []RbacTokenVo
	for _, m := range r {
		vo := RbacTokenVo{}
		voList = append(voList, vo.Transform(m))
	}

	return voList, nil
}

// FindById 根据 ID 查询，返回实体
func (*rbacTokenService) FindById(id uint) (RbacToken, error) {
	m := RbacToken{}
	tx := db.DB.Model(&m).Where("id = ?", id).First(&m)
	if m.ID == 0 {
		return m, errors.New("记录不存在")
	}
	return m, tx.Error
}

// FindById4Vo 根据 ID 查询，返回实体 VO 包装结构体
func (*rbacTokenService) FindById4Vo(id uint) (RbacTokenVo, error) {
	m := RbacToken{}
	tx := db.DB.Model(&m).Where("id = ?", id).First(&m)
	if m.ID == 0 {
		return RbacTokenVo{}, errors.New("记录不存在")
	}
	vo := RbacTokenVo{}
	return vo.Transform(m), tx.Error
}

// 根据 Token 字段查询，返回实体
func (*rbacTokenService) FindByToken(token string) (RbacToken, error) {
	m := RbacToken{}
	tx := db.DB.Model(&m).Where("token = ?", token).First(&m)
	if m.ID == 0 {
		return m, errors.New("记录不存在")
	}
	return m, tx.Error
}

// FindByToken4Vo 根据 Token 字段查询，返回 VO 包装结构体
func (*rbacTokenService) FindByToken4Vo(token string) (RbacTokenVo, error) {
	m := RbacToken{}
	tx := db.DB.Model(&m).Where("token = ?", token).First(&m)
	if m.ID == 0 {
		return RbacTokenVo{}, errors.New("记录不存在")
	}
	vo := RbacTokenVo{}
	return vo.Transform(m), tx.Error
}

func (*rbacTokenService) QueryByUserId(userId uint) ([]RbacToken, error) {
	tx := db.DB.Model(&RbacToken{})
	tx.Where("user_id = ?", userId)
	var r []RbacToken
	tx.Find(&r)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetToken 获取并维护用户 Token
func (s *rbacTokenService) GetToken(user *RbacUser) (string, error) {
	if user == nil {
		return "", errors.New("用户为空")
	}
	exists, err := s.QueryByUserId(user.ID)
	if err != nil {
		return "", err
	}
	// 删除已存在的 token，保证 token 仅最后一次登陆有效
	for _, exist := range exists {
		err := s.DeleteUnscopedById(exist.ID)
		if err != nil {
			logger.Sugar.Error("删除错误", err)
			return "", err
		}
		logger.Sugar.Infof("用户 %s 已清除旧 Token\n", user.LoginName)
	}
	t := uuid.New().String()
	token, err := s.CreateRbacToken(RbacTokenCreateDto{
		Token:  &t,
		UserId: &user.ID,
	})
	if err != nil {
		return "", err
	}

	return *token.Token ,nil
}

// IsValid 判断 token 是否可用
func (s *rbacTokenService) IsValid(token string) bool {
	return true
}

// Logout 注销
func (s *rbacTokenService) Logout(token string) error {
	t, err := s.FindByToken(token)
	if err != nil {
		return err
	}
	return s.DeleteUnscopedById(t.ID)
}
