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
 * long.qian 2021-10-02 13:14 创建
 */

/**
 * @author long.qian
 */

package web

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-lazy-frame/pkg/pub/vo"
	"net/http"
)

// CommonController 公共 Controller
type CommonController struct {

}

// BindBodyJson 绑定请求体JSON到结构体
func (receiver CommonController) BindBodyJson(c *gin.Context, d interface{}) error {
	return c.ShouldBindBodyWith(d, binding.JSON)
}

// Success 成功响应
func (receiver CommonController) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, vo.ResponseResult{
		Code:    "0",
		Message: "",
		Data:    data,
	})
}

// Failed 失败响应
func (receiver CommonController) Failed(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, vo.ResponseResult{
		Code:    "101",
		Message: msg,
		Data:    nil,
	})
}

// FailedWithCode 失败响应，指定错误码
func (receiver CommonController) FailedWithCode(c *gin.Context, code, msg string) {
	c.JSON(http.StatusOK, vo.ResponseResult{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}

