/**
* @author long.qian
*/

package auth_rbac

// =================================================================================
// Code generated by LazyFrame Gen tool. DO NOT EDIT.
// If you want to develop the code, write the file [ gen_vo_rbac_role.go]  in the same directory.
// =================================================================================

import (
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/query"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/util"
)

type RbacRoleVo struct {
	// CreatedAt
    CreatedAt string `json:"createdAt"`
	// DeletedAt
    DeletedAt string `json:"deletedAt"`
	// Id
    Id uint `json:"id"`
	// 角色描述
    RoleDesc string `json:"roleDesc"`
	// 角色名
    RoleName string `json:"roleName"`
	// UpdatedAt
    UpdatedAt string `json:"updatedAt"`
	// 是否可用
    Valid bool `json:"valid"`
}

// Transform 从实体 RbacRole 转换为 Vo
func (RbacRoleVo) Transform(m RbacRole) RbacRoleVo {
	vo := RbacRoleVo{}
	vo.CreatedAt = util.TimeUtil.GetTimeFormatByFormat(m.CreatedAt, util.GolangBirthTime)
	vo.Id = m.ID
	vo.RoleDesc = m.RoleDesc
	vo.RoleName = m.RoleName
	vo.UpdatedAt = util.TimeUtil.GetTimeFormatByFormat(m.UpdatedAt, util.GolangBirthTime)
	vo.Valid = m.Valid
	return vo
}

// TransformTo 从 RbacRoleVo 转换为 实体
func (me RbacRoleVo) TransformTo() RbacRole {
	model := RbacRole{}
	model.CreatedAt = *util.TimeUtil.TimeParseByFormat(me.CreatedAt, util.GolangBirthTime)
	model.ID = me.Id
	model.RoleDesc = me.RoleDesc
	model.RoleName = me.RoleName
	model.UpdatedAt = *util.TimeUtil.TimeParseByFormat(me.UpdatedAt, util.GolangBirthTime)
	model.Valid = me.Valid
	return model
}

type RbacRolePageVo struct {
	query.Page
	Result []RbacRoleVo
}
