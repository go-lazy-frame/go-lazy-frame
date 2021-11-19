/**
* @author long.qian
*/

package auth_rbac

// =================================================================================
// Code generated by LazyFrame Gen tool. DO NOT EDIT.
// If you want to develop the code, write the file [ gen_vo_rbac_user.go]  in the same directory.
// =================================================================================

import (
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/query"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/util"
)

type RbacUserVo struct {
	// Id
    Id uint `json:"id"`
	// CreatedAt
    CreatedAt string `json:"createdAt"`
	// UpdatedAt
    UpdatedAt string `json:"updatedAt"`
	// DeletedAt
    DeletedAt string `json:"deletedAt"`
	// 是否超级管理员
    SuperAdmin bool `json:"superAdmin"`
	// 是否管理员
    Admin bool `json:"admin"`
	// 昵称
    Nickname string `json:"nickname"`
	// 手机号
    Phone string `json:"phone"`
	// 登录名
    LoginName string `json:"loginName"`
	// 登陆密码
    LoginPswd string `json:"loginPswd"`
	// 密码加盐
    Salt string `json:"salt"`
	// 状态 1:账号正常 0:账号禁用 -1:账号违规
    Status int64 `json:"status"`
}

// Transform 从实体 RbacUser 转换为 Vo
func (RbacUserVo) Transform(m RbacUser) RbacUserVo {
	vo := RbacUserVo{}
	vo.Id = m.ID
	vo.CreatedAt = util.TimeUtil.GetTimeFormatByFormat(m.CreatedAt, util.GolangBirthTime)
	vo.UpdatedAt = util.TimeUtil.GetTimeFormatByFormat(m.UpdatedAt, util.GolangBirthTime)
	vo.SuperAdmin = m.SuperAdmin
	vo.Admin = m.Admin
	vo.Nickname = m.Nickname
	vo.Phone = m.Phone
	vo.LoginName = m.LoginName
	vo.LoginPswd = m.LoginPswd
	vo.Salt = m.Salt
	vo.Status = m.Status
	return vo
}

// TransformTo 从 RbacUserVo 转换为 实体
func (me RbacUserVo) TransformTo() RbacUser {
	model := RbacUser{}
	model.ID = me.Id
	model.CreatedAt = *util.TimeUtil.TimeParseByFormat(me.CreatedAt, util.GolangBirthTime)
	model.UpdatedAt = *util.TimeUtil.TimeParseByFormat(me.UpdatedAt, util.GolangBirthTime)
	model.SuperAdmin = me.SuperAdmin
	model.Admin = me.Admin
	model.Nickname = me.Nickname
	model.Phone = me.Phone
	model.LoginName = me.LoginName
	model.LoginPswd = me.LoginPswd
	model.Salt = me.Salt
	model.Status = me.Status
	return model
}

type RbacUserPageVo struct {
	query.Page
	Result []RbacUserVo
}
