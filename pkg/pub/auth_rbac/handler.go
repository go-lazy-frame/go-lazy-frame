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
 * long.qian 2021-10-10 16:49 创建
 */

/**
 * @author long.qian
 */

package auth_rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-lazy-frame/go-lazy-frame/configs"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/util"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/vo"
	"net/http"
	"strings"
)

var (
	// NotAuthUrl 无需鉴权的访问地址前缀
	NotAuthUrl = []string{
		configs.GeneralConfig.ApiPrefix + "/login",
		"/doc",
		"/swagger",
		"/swagger-resources",
	}
)

// RbacHandler 拦截器
func RbacHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		num := int64(0)
		rbacLog := RbacLogCreateDto{
			Ip:        &c.Request.RemoteAddr,
			Url:       &c.Request.URL.Path,
			UrlParams: &c.Request.URL.RawQuery,
			Status:    &num,
		}
		body := make(map[string]interface{})
		err := c.ShouldBindBodyWith(&body, binding.JSON)
		if err == nil {
			s := util.JsonUtil.ParseObjToJsonString(body)
			rbacLog.Body = &s
		}
		defer func() {
			// 保存操作记录
			_, err = RbacLogService.CreateRbacLog(rbacLog)
			if err != nil {
				logger.Sugar.Error("操作记录保存错误：", err)
			}
		}()

		// 获取请求的URI
		requestURI := c.Request.URL.RequestURI()
		// 判断是否需要鉴权
		for _, s := range NotAuthUrl {
			if strings.HasPrefix(requestURI, s) {
				c.Next()
				return
			}
		}
		for _, s := range configs.GeneralConfig.NotAuthUrl {
			if strings.HasPrefix(requestURI, s) {
				c.Next()
				return
			}
		}

		token := c.GetHeader("token")
		if token == "" {
			c.JSON(http.StatusOK, vo.ResponseResult{
				Code:    "403",
				Message: "token not exist",
				Data:    nil,
			})
			c.Abort()
			status := int64(1)
			rbacLog.Status = &status
			return
		}
		user, err := RbacUserService.FindRbacUserByToken(token)
		if err != nil {
			c.JSON(http.StatusOK, vo.ResponseResult{
				Code:    "403",
				Message: err.Error(),
				Data:    nil,
			})
			c.Abort()
			status := int64(1)
			rbacLog.Status = &status
			return
		}
		rbacLog.LoginName = &user.LoginName
		if user.Status != 1 {
			c.JSON(http.StatusOK, vo.ResponseResult{
				Code:    "403",
				Message: "账号已不可用",
				Data:    nil,
			})
			c.Abort()
			status := int64(1)
			rbacLog.Status = &status
			return
		}

		// 判断用户是否拥有权限，若没有地址访问的权限控制，只是菜单权限，可先不实现该功能
	}
}
