/**
* @author long.qian
*/

package auth_rbac

// =================================================================================
// You can edit this file about RbacUser controller code.
// =================================================================================

import (
	"github.com/gin-gonic/gin"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/util"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/web"
)

// RbacUserController 用户
// Controller 规范：
// 每一个控制器接口方法，都要对应一个 WebXxx 的字段定义（其中的 Xxx 为下面的 struct 方法，既访问接口），用于描述该接口的定义，可参考 controller_rbac_user.go
// 提示：接口变动或接口参数变动，需要执行对应 xxx 脚本进行接口文档的同步更新
type RbacUserController struct {
	web.CommonController
	WebQueryPage interface{} `url:"/rbac_user/query_page" method:"post"`
	WebQuery interface{} `url:"/rbac_user/query" method:"post"`
	WebCreate interface{} `url:"/rbac_user/create" method:"post"`
	WebUpdate interface{} `url:"/rbac_user/update" method:"post"`
	WebFindById interface{} `url:"/rbac_user/find_by_id" method:"get"`
	WebDeleteById interface{} `url:"/rbac_user/delete_by_id" method:"get"`
	WebResetPasswordById interface{} `url:"/rbac_user/reset_password_by_id" method:"get"`
	WebFindByLoginName interface{} `url:"/rbac_user/find_by_login_name" method:"get"`
}

// Create
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacUserCreateUsingPOST
// @Tags 用户
// @Summary 用户创建
// @Description 用户创建
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param request body AddRbacUserDto{} true "创建"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number(float),boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_user/create [post]
func (me RbacUserController) Create(c *gin.Context) {
	d := AddRbacUserDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	_, err := RbacUserService.CreateRbacUser(d)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, "OK")
}

// Update
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacUserUpdateUsingPOST
// @Tags 用户
// @Summary 用户更新
// @Description 用户更新，根据 id 更新
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param request body RbacUserUpdateDto{} true "更新"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number(float),boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_user/update [post]
func (me RbacUserController) Update(c *gin.Context) {
	d := RbacUserUpdateDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	err := RbacUserService.UpdateRbacUser(d)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, "OK")
}

// QueryPage
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacUserQueryPageUsingPOST
// @Tags 用户
// @Summary 用户分页查询
// @Description 用户分页查询<br>请求示例：<br>{<br>&nbsp;&nbsp;"start": 0,<br>&nbsp;&nbsp;"limit": 20<br>}<br>其他查询条件，可根据以下条件字段酌情添加查询条件<br>注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param request body RbacUserPageDto{} true "条件查询"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number(float),boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_user/query_page [post]
func (me RbacUserController) QueryPage(c *gin.Context) {
	d := RbacUserPageDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	r, err := RbacUserService.QueryPageRbacUser(d)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, r)
}

// Query
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacUserQueryUsingPOST
// @Tags 用户
// @Summary 用户查询
// @Description 用户查询记录<br>请求示例：<br>{}<br>其他查询条件，可根据以下条件字段酌情添加查询条件<br>注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param request body RbacUserQueryDto{} true "条件查询"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number(float),boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_user/query [post]
func (me RbacUserController) Query(c *gin.Context) {
	d := RbacUserQueryDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	r, err := RbacUserService.QueryRbacUser(d)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, r)
}

// FindById
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacUserFindByIdUsingGET
// @Tags 用户
// @Summary 用户ById
// @Description 用户查询ById
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param id query integer true "查询条件id"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number,boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_user/find_by_id [get]
func (me RbacUserController) FindById(c *gin.Context) {
	if id,ok := c.GetQuery("id"); ok {
		m, err := RbacUserService.FindById(util.NumberUtil.StringToUInt(id))
		if err != nil {
			me.Failed(c, err.Error())
			return
		}
		me.Success(c, m)
		return
	}
	me.Failed(c, "参数错误")
}

// FindByLoginName
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacUserFindByLoginNameUsingGET
// @Tags 用户
// @Summary 用户ByLoginName
// @Description 用户查询ByLoginName<br/>字符串类型字段：空串或不传都代表忽略该查询，如果要指定空串为条件，则指定为 STRING__BLANK，例如 { "name": "STRING__BLANK" }
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param loginName query string true "查询条件值"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number,boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_user/find_by_login_name [get]
func (me RbacUserController) FindByLoginName(c *gin.Context) {
	if q,ok := c.GetQuery("loginName"); ok {
		m, err := RbacUserService.FindByLoginName(q)
		if err != nil {
			me.Failed(c, err.Error())
			return
		}
		me.Success(c, m)
		return
	}
	me.Failed(c, "参数错误")
}

// DeleteById
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacUserDeleteByIdUsingGET
// @Tags 用户
// @Summary 删除ById
// @Description 用户删除ById
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param id query integer true "id"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number,boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_user/delete_by_id [get]
func (me RbacUserController) DeleteById(c *gin.Context) {
	token := c.GetHeader("token")
	if id,ok := c.GetQuery("id"); ok {
		err := RbacUserService.DeleteById(util.NumberUtil.StringToUInt(id), token)
		if err != nil {
			me.Failed(c, err.Error())
			return
		}
		me.Success(c, "删除成功")
		return
	}
	me.Failed(c, "参数错误")
}

// ResetPasswordById
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacUserResetPasswordByIdUsingGET
// @Tags 用户
// @Summary 重置密码
// @Description 重置密码
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param id query integer true "id"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number,boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_user/reset_password_by_id [get]
func (me RbacUserController) ResetPasswordById(c *gin.Context) {
	token := c.GetHeader("token")
	if id,ok := c.GetQuery("id"); ok {
		err := RbacUserService.ResetPasswordById(util.NumberUtil.StringToUInt(id), token)
		if err != nil {
			me.Failed(c, err.Error())
			return
		}
		me.Success(c, "该账户密码已重置为 123456")
		return
	}
	me.Failed(c, "参数错误")
}

