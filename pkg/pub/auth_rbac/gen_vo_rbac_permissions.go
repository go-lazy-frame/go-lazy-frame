/**
* @author long.qian
*/

package auth_rbac

// =================================================================================
// Code generated by LazyFrame Gen tool. DO NOT EDIT.
// If you want to develop the code, write the file [ gen_vo_rbac_permissions.go]  in the same directory.
// =================================================================================

import (
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/query"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/util"
)

type RbacPermissionsVo struct {
	// CreatedAt
    CreatedAt string `json:"createdAt"`
	// DeletedAt
    DeletedAt string `json:"deletedAt"`
	// 权限描述
    Description string `json:"description"`
	// Id
    Id uint `json:"id"`
	// 权限值
    Permission string `json:"permission"`
	// UpdatedAt
    UpdatedAt string `json:"updatedAt"`
}

// Transform 从实体 RbacPermissions 转换为 Vo
func (RbacPermissionsVo) Transform(m RbacPermissions) RbacPermissionsVo {
	vo := RbacPermissionsVo{}
	vo.CreatedAt = util.TimeUtil.GetTimeFormatByFormat(m.CreatedAt, util.GolangBirthTime)
	vo.Description = m.Description
	vo.Id = m.ID
	vo.Permission = m.Permission
	vo.UpdatedAt = util.TimeUtil.GetTimeFormatByFormat(m.UpdatedAt, util.GolangBirthTime)
	return vo
}

// TransformTo 从 RbacPermissionsVo 转换为 实体
func (me RbacPermissionsVo) TransformTo() RbacPermissions {
	model := RbacPermissions{}
	model.CreatedAt = *util.TimeUtil.TimeParseByFormat(me.CreatedAt, util.GolangBirthTime)
	model.Description = me.Description
	model.ID = me.Id
	model.Permission = me.Permission
	model.UpdatedAt = *util.TimeUtil.TimeParseByFormat(me.UpdatedAt, util.GolangBirthTime)
	return model
}

type RbacPermissionsPageVo struct {
	query.Page
	Result []RbacPermissionsVo
}
