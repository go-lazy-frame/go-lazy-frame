/**
* @author long.qian
*/

package auth_rbac

// =================================================================================
// Code generated by LazyFrame Gen tool. DO NOT EDIT.
// If you want to develop the code, write the file [ gen_vo_rbac_permissions.go]  in the same directory.
// =================================================================================

import (
	"go-lazy-frame/pkg/pub/query"
	"go-lazy-frame/pkg/pub/util"
)

type RbacPermissionsVo struct {
	// Id
    Id uint `json:"id"`
	// CreatedAt
    CreatedAt string `json:"createdAt"`
	// UpdatedAt
    UpdatedAt string `json:"updatedAt"`
	// DeletedAt
    DeletedAt string `json:"deletedAt"`
	// 权限描述
    Description string `json:"description"`
	// 权限值
    Permission string `json:"permission"`
}

// Transform 从实体 RbacPermissions 转换为 Vo
func (RbacPermissionsVo) Transform(m RbacPermissions) RbacPermissionsVo {
	vo := RbacPermissionsVo{}
	vo.Id = m.ID
	vo.CreatedAt = util.TimeUtil.GetTimeFormatByFormat(m.CreatedAt, util.GolangBirthTime)
	vo.UpdatedAt = util.TimeUtil.GetTimeFormatByFormat(m.UpdatedAt, util.GolangBirthTime)
	vo.Description = m.Description
	vo.Permission = m.Permission
	return vo
}

// TransformTo 从 RbacPermissionsVo 转换为 实体
func (me RbacPermissionsVo) TransformTo() RbacPermissions {
	model := RbacPermissions{}
	model.ID = me.Id
	model.CreatedAt = *util.TimeUtil.TimeParseByFormat(me.CreatedAt, util.GolangBirthTime)
	model.UpdatedAt = *util.TimeUtil.TimeParseByFormat(me.UpdatedAt, util.GolangBirthTime)
	model.Description = me.Description
	model.Permission = me.Permission
	return model
}

type RbacPermissionsPageVo struct {
	query.Page
	Result []RbacPermissionsVo
}
