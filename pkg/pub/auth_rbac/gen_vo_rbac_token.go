/**
* @author long.qian
*/

package auth_rbac

// =================================================================================
// Code generated by LazyFrame Gen tool. DO NOT EDIT.
// If you want to develop the code, write the file [ gen_vo_rbac_token.go]  in the same directory.
// =================================================================================

import (
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/query"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/util"
)

type RbacTokenVo struct {
	// Id
    Id uint `json:"id"`
	// CreatedAt
    CreatedAt string `json:"createdAt"`
	// UpdatedAt
    UpdatedAt string `json:"updatedAt"`
	// DeletedAt
    DeletedAt string `json:"deletedAt"`
	// Token
    Token string `json:"token"`
	// 用户ID
    UserId uint `json:"userId"`
}

// Transform 从实体 RbacToken 转换为 Vo
func (RbacTokenVo) Transform(m RbacToken) RbacTokenVo {
	vo := RbacTokenVo{}
	vo.Id = m.ID
	vo.CreatedAt = util.TimeUtil.GetTimeFormatByFormat(m.CreatedAt, util.GolangBirthTime)
	vo.UpdatedAt = util.TimeUtil.GetTimeFormatByFormat(m.UpdatedAt, util.GolangBirthTime)
	vo.Token = m.Token
	vo.UserId = m.UserId
	return vo
}

// TransformTo 从 RbacTokenVo 转换为 实体
func (me RbacTokenVo) TransformTo() RbacToken {
	model := RbacToken{}
	model.ID = me.Id
	model.CreatedAt = *util.TimeUtil.TimeParseByFormat(me.CreatedAt, util.GolangBirthTime)
	model.UpdatedAt = *util.TimeUtil.TimeParseByFormat(me.UpdatedAt, util.GolangBirthTime)
	model.Token = me.Token
	model.UserId = me.UserId
	return model
}

type RbacTokenPageVo struct {
	query.Page
	Result []RbacTokenVo
}
