/**
* @author long.qian
*/

package auth_rbac

// =================================================================================
// Code generated by LazyFrame Gen tool. DO NOT EDIT.
// If you want to develop the code, write the file [ gen_dto_rbac_role.go]  in the same directory.
// =================================================================================

import (
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/query"
	"gorm.io/gorm"
)

// RbacRoleCreateDto RbacRole 创建请求 DTO
// 注意：为解决 Golang 的零值问题，该 DTO 字段都为指针类型。设值时，可使用 util.ValueUtil 的 XxxPointer 系列方法进行赋值
type RbacRoleCreateDto struct {
	gorm.Model
	// 角色描述
    RoleDesc *string `json:"roleDesc" gorm:"default:"`
	// 角色名
    RoleName *string `json:"roleName" binding:"required"`
	// 是否可用
    Valid *bool `json:"valid" gorm:"default:1"`
}

// TableName 指定创建体表名
func (RbacRoleCreateDto) TableName() string {
	return "rbac_role"
}

// TransformTo 从 RbacRoleCreateDto 转换为 实体
func (me RbacRoleCreateDto) TransformTo() *RbacRole {
	model := &RbacRole{}
	model.Model = me.Model
	if me.RoleDesc != nil {
		model.RoleDesc = *me.RoleDesc
	}
	if me.RoleName != nil {
		model.RoleName = *me.RoleName
	}
	if me.Valid != nil {
		model.Valid = *me.Valid
	}
	return model
}

// RbacRoleQueryDto 查询请求 DTO
type RbacRoleQueryDto struct {
	// 【created_at】CreatedAt 时间范围（包含边界）
	CreatedAtBetween []interface{} `json:"createdAtBetween" field:"created_at" type:"between"`
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
	// 【role_desc】角色描述 全匹配
	RoleDesc interface{} `json:"roleDesc" field:"role_desc"  type:"equal"`
	// 【role_desc】角色描述 in 查询
	RoleDescIn []interface{} `json:"roleDescIn" field:"role_desc"  type:"in"`
	// 【role_desc】角色描述 左匹配(xxx%)
	RoleDescLeft interface{} `json:"roleDescLeft" field:"role_desc"  type:"likeRight"`
	// 【role_desc】角色描述 右匹配(%xxx，查询有性能影响)
	RoleDescRight interface{} `json:"roleDescRight" field:"role_desc"  type:"likeLeft"`
	// 【role_desc】角色描述 模糊匹配(%xxx%，查询有性能影响)
	RoleDescMiddle interface{} `json:"roleDescMiddle" field:"role_desc"  type:"likeMiddle"`
	// 【role_name】角色名 全匹配
	RoleName interface{} `json:"roleName" field:"role_name"  type:"equal"`
	// 【role_name】角色名 in 查询
	RoleNameIn []interface{} `json:"roleNameIn" field:"role_name"  type:"in"`
	// 【role_name】角色名 左匹配(xxx%)
	RoleNameLeft interface{} `json:"roleNameLeft" field:"role_name"  type:"likeRight"`
	// 【role_name】角色名 右匹配(%xxx，查询有性能影响)
	RoleNameRight interface{} `json:"roleNameRight" field:"role_name"  type:"likeLeft"`
	// 【role_name】角色名 模糊匹配(%xxx%，查询有性能影响)
	RoleNameMiddle interface{} `json:"roleNameMiddle" field:"role_name"  type:"likeMiddle"`
	// 【updated_at】UpdatedAt 时间范围（包含边界）
	UpdatedAtBetween []interface{} `json:"updatedAtBetween" field:"updated_at" type:"between"`
	// 【valid】是否可用 全匹配
	Valid interface{} `json:"valid" field:"valid"  type:"equal"`
	// 【valid】是否可用 in 查询
	ValidIn []interface{} `json:"validIn" field:"valid"  type:"in"`
	// 排序，例如：["id desc", "name asc"] 字段名为每个字段说明后的开头中【】内的内容，方式只能为 desc 或 asc。支持指定多个字段排序
	OrderBy []string `json:"orderBy"`
}

// RbacRoleUpdateDto 更新请求 DTO
type RbacRoleUpdateDto struct {
	// Id
    Id interface{} `json:"id"`
	// 角色描述
    RoleDesc interface{} `json:"roleDesc"`
	// 角色名
    RoleName interface{} `json:"roleName"`
	// 是否可用
    Valid interface{} `json:"valid"`
}

// RbacRolePageDto 查询请求（分页） DTO
type RbacRolePageDto struct {
	query.Page
	RbacRoleQueryDto
}

