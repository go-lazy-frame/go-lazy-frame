/**
* @author long.qian
 */

package auth_rbac

// =================================================================================
// You can edit this file about RbacUser vo code.
// =================================================================================

// AddRbacUserDto RbacUser 创建请求 DTO
type AddRbacUserDto struct {
	// 是否管理员
	Admin bool `json:"admin"`
	// 登录名
	LoginName string `json:"loginName" binding:"required"`
	// 登陆密码
	LoginPswd string `json:"loginPswd" binding:"required"`
	// 重复登陆密码
	RepeatLoginPswd string `json:"repeatLoginPswd" binding:"required"`
	// 昵称
	Nickname string `json:"nickname"`
	// 手机号
	Phone string `json:"phone"`
	// 是否超级管理员
	SuperAdmin bool `json:"superAdmin"`
}
