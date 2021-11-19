/**
* @author long.qian
*/

package auth_rbac

// =================================================================================
// Code generated by LazyFrame Gen tool. DO NOT EDIT.
// If you want to develop the code, write the file [ gen_dto_rbac_token.go]  in the same directory.
// =================================================================================

import (
	"go-lazy-frame/pkg/pub/query"
)

// RbacTokenCreateDto RbacToken 创建请求 DTO
type RbacTokenCreateDto struct {
	// Token
    Token string `json:"token" binding:"required"`
	// 用户ID
    UserId uint `json:"userId" binding:"required"`
}

// TransformTo 从 RbacTokenCreateDto 转换为 实体
func (me RbacTokenCreateDto) TransformTo() *RbacToken {
	model := &RbacToken{
		Token: me.Token,
		UserId: me.UserId,
	}
	return model
}

// RbacTokenQueryDto 查询请求 DTO
type RbacTokenQueryDto struct {
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
	// 【token】Token 全匹配
	Token interface{} `json:"token" field:"token"  type:"equal"`
	// 【token】Token in 查询
	TokenIn []interface{} `json:"tokenIn" field:"token"  type:"in"`
	// 【token】Token 左匹配(xxx%)
	TokenLeft interface{} `json:"tokenLeft" field:"token"  type:"likeRight"`
	// 【token】Token 右匹配(%xxx，查询有性能影响)
	TokenRight interface{} `json:"tokenRight" field:"token"  type:"likeLeft"`
	// 【token】Token 模糊匹配(%xxx%，查询有性能影响)
	TokenMiddle interface{} `json:"tokenMiddle" field:"token"  type:"likeMiddle"`
	// 【user_id】用户ID 全匹配
	UserId interface{} `json:"userId" field:"user_id"  type:"equal"`
	// 【user_id】用户ID in 查询
	UserIdIn []interface{} `json:"userIdIn" field:"user_id"  type:"in"`
	// 排序，例如：["id desc", "name asc"] 字段名为每个字段说明后的开头中【】内的内容，方式只能为 desc 或 asc。支持指定多个字段排序
	OrderBy []string `json:"orderBy"`
}

// RbacTokenUpdateDto 更新请求 DTO
type RbacTokenUpdateDto struct {
	// Id
    Id interface{} `json:"id"`
	// Token
    Token interface{} `json:"token"`
	// 用户ID
    UserId interface{} `json:"userId"`
}

// RbacTokenPageDto 查询请求（分页） DTO
type RbacTokenPageDto struct {
	query.Page
	RbacTokenQueryDto
}
