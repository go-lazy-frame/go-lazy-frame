/**
* @author long.qian
*/

package auth_rbac

// =================================================================================
// Code generated by LazyFrame Gen tool. DO NOT EDIT.
// If you want to develop the code, write the file [ gen_dto_rbac_user.go]  in the same directory.
// =================================================================================

import (
	"go-lazy-frame/pkg/pub/query"
)

// RbacUserCreateDto RbacUser 创建请求 DTO
type RbacUserCreateDto struct {
	// 是否超级管理员
    SuperAdmin bool `json:"superAdmin"`
	// 是否管理员
    Admin bool `json:"admin"`
	// 昵称
    Nickname string `json:"nickname"`
	// 手机号
    Phone string `json:"phone"`
	// 登录名
    LoginName string `json:"loginName" binding:"required"`
	// 登陆密码
    LoginPswd string `json:"loginPswd" binding:"required"`
	// 密码加盐
    Salt string `json:"salt" binding:"required"`
	// 状态 1:账号正常 0:账号禁用 -1:账号违规
    Status int64 `json:"status"`
}

// TransformTo 从 RbacUserCreateDto 转换为 实体
func (me RbacUserCreateDto) TransformTo() *RbacUser {
	model := &RbacUser{
		SuperAdmin: me.SuperAdmin,
		Admin: me.Admin,
		Nickname: me.Nickname,
		Phone: me.Phone,
		LoginName: me.LoginName,
		LoginPswd: me.LoginPswd,
		Salt: me.Salt,
		Status: me.Status,
	}
	return model
}

// RbacUserQueryDto 查询请求 DTO
type RbacUserQueryDto struct {
	// 【id】Id 全匹配
	Id interface{} `json:"id" field:"id"  type:"equal"`
	// 【id】Id in 查询
	IdIn []interface{} `json:"idIn" field:"id"  type:"in"`
	// 【id】Id 范围（包含边界）
	IdBetween []interface{} `json:"idBetween" field:"id"  type:"between"`
	// 【id】Id 大于
	IdGt interface{} `json:"idGt" field:"id"  type:"gt"`
	// 【id】Id 大于等于
	IdGte interface{} `json:"idGte" field:"id"  type:"gte"`
	// 【id】Id 小于
	IdLt interface{} `json:"idLt" field:"id"  type:"lt"`
	// 【id】Id 小于等于
	IdLte interface{} `json:"idLte" field:"id"  type:"lte"`
	// 【created_at】CreatedAt 时间范围（包含边界）
	CreatedAtBetween []interface{} `json:"createdAtBetween" field:"created_at" type:"between"`
	// 【updated_at】UpdatedAt 时间范围（包含边界）
	UpdatedAtBetween []interface{} `json:"updatedAtBetween" field:"updated_at" type:"between"`
	// 【super_admin】是否超级管理员 全匹配
	SuperAdmin interface{} `json:"superAdmin" field:"super_admin"  type:"equal"`
	// 【super_admin】是否超级管理员 in 查询
	SuperAdminIn []interface{} `json:"superAdminIn" field:"super_admin"  type:"in"`
	// 【admin】是否管理员 全匹配
	Admin interface{} `json:"admin" field:"admin"  type:"equal"`
	// 【admin】是否管理员 in 查询
	AdminIn []interface{} `json:"adminIn" field:"admin"  type:"in"`
	// 【nickname】昵称 全匹配
	Nickname interface{} `json:"nickname" field:"nickname"  type:"equal"`
	// 【nickname】昵称 in 查询
	NicknameIn []interface{} `json:"nicknameIn" field:"nickname"  type:"in"`
	// 【nickname】昵称 左匹配(xxx%)
	NicknameLeft interface{} `json:"nicknameLeft" field:"nickname"  type:"likeRight"`
	// 【nickname】昵称 右匹配(%xxx，查询有性能影响)
	NicknameRight interface{} `json:"nicknameRight" field:"nickname"  type:"likeLeft"`
	// 【nickname】昵称 模糊匹配(%xxx%，查询有性能影响)
	NicknameMiddle interface{} `json:"nicknameMiddle" field:"nickname"  type:"likeMiddle"`
	// 【phone】手机号 全匹配
	Phone interface{} `json:"phone" field:"phone"  type:"equal"`
	// 【phone】手机号 in 查询
	PhoneIn []interface{} `json:"phoneIn" field:"phone"  type:"in"`
	// 【phone】手机号 左匹配(xxx%)
	PhoneLeft interface{} `json:"phoneLeft" field:"phone"  type:"likeRight"`
	// 【phone】手机号 右匹配(%xxx，查询有性能影响)
	PhoneRight interface{} `json:"phoneRight" field:"phone"  type:"likeLeft"`
	// 【phone】手机号 模糊匹配(%xxx%，查询有性能影响)
	PhoneMiddle interface{} `json:"phoneMiddle" field:"phone"  type:"likeMiddle"`
	// 【login_name】登录名 全匹配
	LoginName interface{} `json:"loginName" field:"login_name"  type:"equal"`
	// 【login_name】登录名 in 查询
	LoginNameIn []interface{} `json:"loginNameIn" field:"login_name"  type:"in"`
	// 【login_name】登录名 左匹配(xxx%)
	LoginNameLeft interface{} `json:"loginNameLeft" field:"login_name"  type:"likeRight"`
	// 【login_name】登录名 右匹配(%xxx，查询有性能影响)
	LoginNameRight interface{} `json:"loginNameRight" field:"login_name"  type:"likeLeft"`
	// 【login_name】登录名 模糊匹配(%xxx%，查询有性能影响)
	LoginNameMiddle interface{} `json:"loginNameMiddle" field:"login_name"  type:"likeMiddle"`
	// 【login_pswd】登陆密码 全匹配
	LoginPswd interface{} `json:"loginPswd" field:"login_pswd"  type:"equal"`
	// 【login_pswd】登陆密码 in 查询
	LoginPswdIn []interface{} `json:"loginPswdIn" field:"login_pswd"  type:"in"`
	// 【login_pswd】登陆密码 左匹配(xxx%)
	LoginPswdLeft interface{} `json:"loginPswdLeft" field:"login_pswd"  type:"likeRight"`
	// 【login_pswd】登陆密码 右匹配(%xxx，查询有性能影响)
	LoginPswdRight interface{} `json:"loginPswdRight" field:"login_pswd"  type:"likeLeft"`
	// 【login_pswd】登陆密码 模糊匹配(%xxx%，查询有性能影响)
	LoginPswdMiddle interface{} `json:"loginPswdMiddle" field:"login_pswd"  type:"likeMiddle"`
	// 【salt】密码加盐 全匹配
	Salt interface{} `json:"salt" field:"salt"  type:"equal"`
	// 【salt】密码加盐 in 查询
	SaltIn []interface{} `json:"saltIn" field:"salt"  type:"in"`
	// 【salt】密码加盐 左匹配(xxx%)
	SaltLeft interface{} `json:"saltLeft" field:"salt"  type:"likeRight"`
	// 【salt】密码加盐 右匹配(%xxx，查询有性能影响)
	SaltRight interface{} `json:"saltRight" field:"salt"  type:"likeLeft"`
	// 【salt】密码加盐 模糊匹配(%xxx%，查询有性能影响)
	SaltMiddle interface{} `json:"saltMiddle" field:"salt"  type:"likeMiddle"`
	// 【status】状态 1:账号正常 0:账号禁用 -1:账号违规 全匹配
	Status interface{} `json:"status" field:"status"  type:"equal"`
	// 【status】状态 1:账号正常 0:账号禁用 -1:账号违规 in 查询
	StatusIn []interface{} `json:"statusIn" field:"status"  type:"in"`
	// 【status】状态 1:账号正常 0:账号禁用 -1:账号违规 范围（包含边界）
	StatusBetween []interface{} `json:"statusBetween" field:"status"  type:"between"`
	// 【status】状态 1:账号正常 0:账号禁用 -1:账号违规 大于
	StatusGt interface{} `json:"statusGt" field:"status"  type:"gt"`
	// 【status】状态 1:账号正常 0:账号禁用 -1:账号违规 大于等于
	StatusGte interface{} `json:"statusGte" field:"status"  type:"gte"`
	// 【status】状态 1:账号正常 0:账号禁用 -1:账号违规 小于
	StatusLt interface{} `json:"statusLt" field:"status"  type:"lt"`
	// 【status】状态 1:账号正常 0:账号禁用 -1:账号违规 小于等于
	StatusLte interface{} `json:"statusLte" field:"status"  type:"lte"`
	// 排序，例如：["id desc", "name asc"] 字段名为每个字段说明后的开头中【】内的内容，方式只能为 desc 或 asc。支持指定多个字段排序
	OrderBy []string `json:"orderBy"`
}

// RbacUserUpdateDto 更新请求 DTO
type RbacUserUpdateDto struct {
	// Id
    Id interface{} `json:"id"`
	// 是否超级管理员
    SuperAdmin interface{} `json:"superAdmin"`
	// 是否管理员
    Admin interface{} `json:"admin"`
	// 昵称
    Nickname interface{} `json:"nickname"`
	// 手机号
    Phone interface{} `json:"phone"`
	// 登录名
    LoginName interface{} `json:"loginName"`
	// 登陆密码
    LoginPswd interface{} `json:"loginPswd"`
	// 密码加盐
    Salt interface{} `json:"salt"`
	// 状态 1:账号正常 0:账号禁用 -1:账号违规
    Status interface{} `json:"status"`
}

// RbacUserPageDto 查询请求（分页） DTO
type RbacUserPageDto struct {
	query.Page
	RbacUserQueryDto
}
