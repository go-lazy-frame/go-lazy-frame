/**
* @author long.qian
*/

package auth_rbac

// =================================================================================
// Code generated by LazyFrame Gen tool. DO NOT EDIT.
// If you want to develop the code, write the file [ gen_dto_rbac_log.go]  in the same directory.
// =================================================================================

import (
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/query"
)

// RbacLogCreateDto RbacLog 创建请求 DTO
type RbacLogCreateDto struct {
	// 请求Body体参数
    Body *string `json:"body"`
	// Id
    Id *uint `json:"id"`
	// IP地址
    Ip *string `json:"ip" binding:"required"`
	// 登陆用户
    LoginName *string `json:"loginName" binding:"required"`
	// 状态 0：正常 1：鉴权失败
    Status *int64 `json:"status"`
	// 访问地址
    Url *string `json:"url" binding:"required"`
	// 地址参数
    UrlParams *string `json:"urlParams" binding:"required"`
}

// TableName 指定创建体表名
func (RbacLogCreateDto) TableName() string {
	return "rbac_log"
}

// RbacLogQueryDto 查询请求 DTO
type RbacLogQueryDto struct {
	// 【body】请求Body体参数 全匹配
	Body interface{} `json:"body" field:"body"  type:"equal"`
	// 【body】请求Body体参数 in 查询
	BodyIn []interface{} `json:"bodyIn" field:"body"  type:"in"`
	// 【body】请求Body体参数 左匹配(xxx%)
	BodyLeft interface{} `json:"bodyLeft" field:"body"  type:"likeRight"`
	// 【body】请求Body体参数 右匹配(%xxx，查询有性能影响)
	BodyRight interface{} `json:"bodyRight" field:"body"  type:"likeLeft"`
	// 【body】请求Body体参数 模糊匹配(%xxx%，查询有性能影响)
	BodyMiddle interface{} `json:"bodyMiddle" field:"body"  type:"likeMiddle"`
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
	// 【ip】IP地址 全匹配
	Ip interface{} `json:"ip" field:"ip"  type:"equal"`
	// 【ip】IP地址 in 查询
	IpIn []interface{} `json:"ipIn" field:"ip"  type:"in"`
	// 【ip】IP地址 左匹配(xxx%)
	IpLeft interface{} `json:"ipLeft" field:"ip"  type:"likeRight"`
	// 【ip】IP地址 右匹配(%xxx，查询有性能影响)
	IpRight interface{} `json:"ipRight" field:"ip"  type:"likeLeft"`
	// 【ip】IP地址 模糊匹配(%xxx%，查询有性能影响)
	IpMiddle interface{} `json:"ipMiddle" field:"ip"  type:"likeMiddle"`
	// 【login_name】登陆用户 全匹配
	LoginName interface{} `json:"loginName" field:"login_name"  type:"equal"`
	// 【login_name】登陆用户 in 查询
	LoginNameIn []interface{} `json:"loginNameIn" field:"login_name"  type:"in"`
	// 【login_name】登陆用户 左匹配(xxx%)
	LoginNameLeft interface{} `json:"loginNameLeft" field:"login_name"  type:"likeRight"`
	// 【login_name】登陆用户 右匹配(%xxx，查询有性能影响)
	LoginNameRight interface{} `json:"loginNameRight" field:"login_name"  type:"likeLeft"`
	// 【login_name】登陆用户 模糊匹配(%xxx%，查询有性能影响)
	LoginNameMiddle interface{} `json:"loginNameMiddle" field:"login_name"  type:"likeMiddle"`
	// 【status】状态 0：正常 1：鉴权失败 全匹配
	Status interface{} `json:"status" field:"status"  type:"equal"`
	// 【status】状态 0：正常 1：鉴权失败 in 查询
	StatusIn []interface{} `json:"statusIn" field:"status"  type:"in"`
	// 【status】状态 0：正常 1：鉴权失败 范围（包含边界）
	StatusBetween []interface{} `json:"statusBetween" field:"status"  type:"between"`
	// 【status】状态 0：正常 1：鉴权失败 大于
	StatusGt interface{} `json:"statusGt" field:"status"  type:"gt"`
	// 【status】状态 0：正常 1：鉴权失败 大于等于
	StatusGte interface{} `json:"statusGte" field:"status"  type:"gte"`
	// 【status】状态 0：正常 1：鉴权失败 小于
	StatusLt interface{} `json:"statusLt" field:"status"  type:"lt"`
	// 【status】状态 0：正常 1：鉴权失败 小于等于
	StatusLte interface{} `json:"statusLte" field:"status"  type:"lte"`
	// 【updated_at】UpdatedAt 时间范围（包含边界）
	UpdatedAtBetween []interface{} `json:"updatedAtBetween" field:"updated_at" type:"between"`
	// 【url】访问地址 全匹配
	Url interface{} `json:"url" field:"url"  type:"equal"`
	// 【url】访问地址 in 查询
	UrlIn []interface{} `json:"urlIn" field:"url"  type:"in"`
	// 【url】访问地址 左匹配(xxx%)
	UrlLeft interface{} `json:"urlLeft" field:"url"  type:"likeRight"`
	// 【url】访问地址 右匹配(%xxx，查询有性能影响)
	UrlRight interface{} `json:"urlRight" field:"url"  type:"likeLeft"`
	// 【url】访问地址 模糊匹配(%xxx%，查询有性能影响)
	UrlMiddle interface{} `json:"urlMiddle" field:"url"  type:"likeMiddle"`
	// 【url_params】地址参数 全匹配
	UrlParams interface{} `json:"urlParams" field:"url_params"  type:"equal"`
	// 【url_params】地址参数 in 查询
	UrlParamsIn []interface{} `json:"urlParamsIn" field:"url_params"  type:"in"`
	// 【url_params】地址参数 左匹配(xxx%)
	UrlParamsLeft interface{} `json:"urlParamsLeft" field:"url_params"  type:"likeRight"`
	// 【url_params】地址参数 右匹配(%xxx，查询有性能影响)
	UrlParamsRight interface{} `json:"urlParamsRight" field:"url_params"  type:"likeLeft"`
	// 【url_params】地址参数 模糊匹配(%xxx%，查询有性能影响)
	UrlParamsMiddle interface{} `json:"urlParamsMiddle" field:"url_params"  type:"likeMiddle"`
	// 排序，例如：["id desc", "name asc"] 字段名为每个字段说明后的开头中【】内的内容，方式只能为 desc 或 asc。支持指定多个字段排序
	OrderBy []string `json:"orderBy"`
}

// RbacLogUpdateDto 更新请求 DTO
type RbacLogUpdateDto struct {
	// 请求Body体参数
    Body interface{} `json:"body"`
	// Id
    Id interface{} `json:"id"`
	// IP地址
    Ip interface{} `json:"ip"`
	// 登陆用户
    LoginName interface{} `json:"loginName"`
	// 状态 0：正常 1：鉴权失败
    Status interface{} `json:"status"`
	// 访问地址
    Url interface{} `json:"url"`
	// 地址参数
    UrlParams interface{} `json:"urlParams"`
}

// RbacLogPageDto 查询请求（分页） DTO
type RbacLogPageDto struct {
	query.Page
	RbacLogQueryDto
}

