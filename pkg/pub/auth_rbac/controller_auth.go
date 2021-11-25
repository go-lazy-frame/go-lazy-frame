//             ,%%%%%%%%,
//           ,%%/\%%%%/\%%
//          ,%%%\c "" J/%%%
// %.       %%%%/ o  o \%%%
// `%%.     %%%%    _  |%%%
//  `%%     `%%%%(__Y__)%%'
//  //       ;%%%%`\-/%%%'
// ((       /  `%%%%%%%'
//  \\    .'          |
//   \\  /       \  | |
//    \\/攻城狮保佑) | |
//     \         /_ | |__
//     (___________)))))))                   `\/'
/*
 * 修订记录:
 * long.qian 2021-10-10 22:28 创建
 */

/**
 * @author long.qian
 */

package auth_rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/web"
)

// AuthController 鉴权
// Controller 规范：
// 每一个控制器接口方法，都要对应一个 WebXxx 的字段定义（其中的 Xxx 为下面的 struct 方法，既访问接口），用于描述该接口的定义，可参考 controller_rbac_log.go
// 提示：接口变动或接口参数变动，需要执行对应 xxx 脚本进行接口文档的同步更新
type AuthController struct {
	web.CommonController
	WebLogin  interface{} `url:"/login" method:"post"`
	WebLogout interface{} `url:"/logout" method:"post"`
}

// Login
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id AuthLoginUsingPOST
// @Tags 登陆注销
// @Summary 登陆
// @Description 登陆
// @Accept json
// @Produce  json
// @Param request body LoginDto{} true "创建"
// 		参数名 参数类型 参数对象类型 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /login [post]
func (me AuthController) Login(c *gin.Context) {
	d := LoginDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	user, err := RbacUserService.FindByLoginName(d.LoginName)
	if err != nil {
		me.Failed(c, "账号或密码错误")
		return
	}
	if user.Status != 1 {
		me.Failed(c, "账号已不可用")
		return
	}
	if RbacUserService.GeneratePswd(d.LoginPswd, user.Salt) != user.LoginPswd {
		me.Failed(c, "账号或密码错误")
		return
	}
	token, err := RbacTokenService.GetToken(&user)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}

	permissions, err := RbacUserService.GetPermissions(user)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}

	me.Success(c, LoginVo{
		LoginName:   user.LoginName,
		Nickname:    user.Nickname,
		Phone:       user.Phone,
		Token:       token,
		Permissions: permissions,
	})
}

// Logout
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id AuthLogoutUsingPOST
// @Tags 登陆注销
// @Summary 注销
// @Description 注销
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// 		参数名 参数类型 参数对象类型 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /logout [post]
func (me AuthController) Logout(c *gin.Context) {
	token := c.GetHeader("token")
	err := RbacTokenService.Logout(token)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, "Success")
}
