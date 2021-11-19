/**
* @author long.qian
*/

package auth_rbac

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/db"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/query"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/update"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/util"
)

// =================================================================================
// You can edit this file about RbacUser service code.
// eg: func (*rbacUserService) FuncName(d dto.XxxDto) error {}
// =================================================================================

type rbacUserService struct {}

var (
	RbacUserService = new(rbacUserService)
)

// CreateRbacUser 创建
func (me *rbacUserService) CreateRbacUser(d AddRbacUserDto) (*RbacUser, error) {
	if d.LoginName == "" || d.LoginPswd == "" || d.RepeatLoginPswd == "" {
		return nil, errors.New("用户名密码不能为空")
	}
	if d.LoginPswd != d.RepeatLoginPswd {
		return nil, errors.New("密码不匹配")
	}
	salt := uuid.New().String()
	c := &RbacUser{
		SuperAdmin: d.SuperAdmin,
		Admin:      d.Admin,
		Nickname:   d.Nickname,
		Phone:      d.Phone,
		LoginName:  d.LoginName,
		LoginPswd:  me.GeneratePswd(d.LoginPswd, salt),
		Salt:       salt,
		Status:     1,
	}
	err := db.DB.Create(c).Error
	if err != nil {
		return nil, errors.New("create failed：" + err.Error())
	}
	return c, nil
}

// UpdateRbacUser 更新
func (*rbacUserService) UpdateRbacUser(d RbacUserUpdateDto) error {
	if d.Id == 0 {
		return errors.New("update failed：" + "id is Empty")
	}
	err := db.DB.Model(&RbacUser{}).Where("id = ?", d.Id).Updates(update.GenerateUpdatesMap(d)).Error
	if err != nil {
		return errors.New("update failed：" + err.Error())
	}
	return nil
}

// DeleteSoftById 软删除（要查询被软删除的记录，用Unscoped，例如：db.DB.Unscoped().Where("name = ?", "xxx").Find(&users)）
func (*rbacUserService) DeleteSoftById(id uint) error {
	return db.DB.Where("id = ?", id).Delete(&RbacUser{}).Error
}

// DeleteUnscopedById 永久删除
func (*rbacUserService) DeleteUnscopedById(id uint) error {
	return db.DB.Unscoped().Where("id = ?", id).Delete(&RbacUser{}).Error
}

// QueryPageRbacUser 分页查询
func (*rbacUserService) QueryPageRbacUser(d RbacUserPageDto) (*RbacUserPageVo, error) {
	tx := db.DB.Model(&RbacUser{})
	whereErr := query.WhereHandler(tx, d)
	if whereErr != nil {
		return nil, whereErr
	}
	orderErr := query.OrderHandler(tx, d)
	if orderErr != nil {
		return nil, orderErr
	}
	var r []RbacUser
	var total int64
	tx.Count(&total)
	query.PageHandler(tx, &d.Page)
	tx.Find(&r)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	var voList []RbacUserVo
	for _, m := range r {
		vo := RbacUserVo{}
		voList = append(voList, vo.Transform(m))
	}

	return &RbacUserPageVo{
		Page: query.Page{
			PageNum: d.PageNum,
			PageSize: d.PageSize,
			Total: total,
		},
		Result: voList,
	}, nil
}

// QueryRbacUser 查询
func (*rbacUserService) QueryRbacUser(d RbacUserQueryDto) ([]RbacUserVo, error) {
	tx := db.DB.Model(&RbacUser{})
	whereErr := query.WhereHandler(tx, d)
	if whereErr != nil {
		return nil, whereErr
	}
	orderErr := query.OrderHandler(tx, d)
	if orderErr != nil {
		return nil, orderErr
	}
	var r []RbacUser
	tx.Find(&r)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	var voList []RbacUserVo
	for _, m := range r {
		vo := RbacUserVo{}
		voList = append(voList, vo.Transform(m))
	}

	return voList, nil
}

// FindById 根据 ID 查询，返回实体
func (*rbacUserService) FindById(id uint) (RbacUser, error) {
	m := RbacUser{}
	tx := db.DB.Model(&m).Where("id = ?", id).First(&m)
	if m.ID == 0 {
		return m, errors.New("记录不存在")
	}
	return m, tx.Error
}

// FindById4Vo 根据 ID 查询，返回实体 VO 包装结构体
func (*rbacUserService) FindById4Vo(id uint) (RbacUserVo, error) {
	m := RbacUser{}
	tx := db.DB.Model(&m).Where("id = ?", id).First(&m)
	if m.ID == 0 {
		return RbacUserVo{}, errors.New("记录不存在")
	}
	vo := RbacUserVo{}
	return vo.Transform(m), tx.Error
}

// FindByLoginName 根据 LoginName 字段查询，返回实体
func (*rbacUserService) FindByLoginName(loginName string) (RbacUser, error) {
	m := RbacUser{}
	tx := db.DB.Model(&m).Where("login_name = ?", loginName).First(&m)
	if m.ID == 0 {
		return m, errors.New("记录不存在")
	}
	return m, tx.Error
}

// FindByLoginName4Vo 根据 LoginName 字段查询，返回 VO 包装结构体
func (*rbacUserService) FindByLoginName4Vo(loginName string) (RbacUserVo, error) {
	m := RbacUser{}
	tx := db.DB.Model(&m).Where("login_name = ?", loginName).First(&m)
	if m.ID == 0 {
		return RbacUserVo{}, errors.New("记录不存在")
	}
	vo := RbacUserVo{}
	return vo.Transform(m), tx.Error
}

func (me *rbacUserService) FindRbacUserByToken(token string) (RbacUser, error) {
	rbacToken, err := RbacTokenService.FindByToken(token)
	if err != nil {
		return RbacUser{}, fmt.Errorf("Token:%s\n", err.Error())
	}
	user, err := me.FindById(rbacToken.UserId)
	if err != nil {
		return RbacUser{}, fmt.Errorf("User:%s\n", err.Error())
	}
	return user, nil
}

// GeneratePswd 生成密码
func (me *rbacUserService) GeneratePswd(loginPswd string, salt string) string {
	return util.Md5Util.Md5(salt + loginPswd + salt)
}
