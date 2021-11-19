/**
* @author long.qian
*/

package auth_rbac

// =================================================================================
// You can edit this file about RbacToken controller code.
// =================================================================================

import (
	"github.com/gin-gonic/gin"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/util"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/web"
)

// RbacTokenController Token
// Controller 规范：
// 每一个控制器接口方法，都要对应一个 WebXxx 的字段定义（其中的 Xxx 为下面的 struct 方法，既访问接口），用于描述该接口的定义，可参考 controller_rbac_token.go
// 提示：接口变动或接口参数变动，需要执行对应 xxx 脚本进行接口文档的同步更新
type RbacTokenController struct {
	web.CommonController
	WebQueryPage interface{} `url:"/rbac_token/query_page" method:"post"`
	WebQuery interface{} `url:"/rbac_token/query" method:"post"`
	WebCreate interface{} `url:"/rbac_token/create" method:"post"`
	WebUpdate interface{} `url:"/rbac_token/update" method:"post"`
	WebFindById interface{} `url:"/rbac_token/find_by_id" method:"get"`
	WebFindByToken interface{} `url:"/rbac_token/find_by_token" method:"get"`
}

// Create
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacTokenCreateUsingPOST
// @Tags Token
// @Summary Token创建
// @Description Token创建
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param request body RbacTokenCreateDto{} true "创建"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number(float),boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_token/create [post]
func (me RbacTokenController) Create(c *gin.Context) {
	d := RbacTokenCreateDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	_, err := RbacTokenService.CreateRbacToken(d)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, "OK")
}

// Update
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacTokenUpdateUsingPOST
// @Tags Token
// @Summary Token更新
// @Description Token更新，根据 id 更新
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param request body RbacTokenUpdateDto{} true "更新"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number(float),boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_token/update [post]
func (me RbacTokenController) Update(c *gin.Context) {
	d := RbacTokenUpdateDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	err := RbacTokenService.UpdateRbacToken(d)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, "OK")
}

// QueryPage
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacTokenQueryPageUsingPOST
// @Tags Token
// @Summary Token分页查询
// @Description Token分页查询<br>请求示例：<br>{<br>&nbsp;&nbsp;"start": 0,<br>&nbsp;&nbsp;"limit": 20<br>}<br>其他查询条件，可根据以下条件字段酌情添加查询条件<br>注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param request body RbacTokenPageDto{} true "条件查询"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number(float),boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_token/query_page [post]
func (me RbacTokenController) QueryPage(c *gin.Context) {
	d := RbacTokenPageDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	r, err := RbacTokenService.QueryPageRbacToken(d)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, r)
}

// Query
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacTokenQueryUsingPOST
// @Tags Token
// @Summary Token查询
// @Description Token查询记录<br>请求示例：<br>{}<br>其他查询条件，可根据以下条件字段酌情添加查询条件<br>注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param request body RbacTokenQueryDto{} true "条件查询"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number(float),boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_token/query [post]
func (me RbacTokenController) Query(c *gin.Context) {
	d := RbacTokenQueryDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	r, err := RbacTokenService.QueryRbacToken(d)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, r)
}

// FindById
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacTokenFindByIdUsingGET
// @Tags Token
// @Summary TokenById
// @Description Token查询ById
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param id query integer true "查询条件id"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number,boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_token/find_by_id [get]
func (me RbacTokenController) FindById(c *gin.Context) {
	if id,ok := c.GetQuery("id"); ok {
		m, err := RbacTokenService.FindById(util.NumberUtil.StringToUInt(id))
		if err != nil {
			me.Failed(c, err.Error())
			return
		}
		me.Success(c, m)
		return
	}
	me.Failed(c, "参数错误")
}

// FindByToken
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacTokenFindByTokenUsingGET
// @Tags Token
// @Summary TokenByToken
// @Description Token查询ByToken<br/>字符串类型字段：空串或不传都代表忽略该查询，如果要指定空串为条件，则指定为 STRING__BLANK，例如 { "name": "STRING__BLANK" }
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param token query string true "查询条件值"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number,boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_token/find_by_token [get]
func (me RbacTokenController) FindByToken(c *gin.Context) {
	if q,ok := c.GetQuery("token"); ok {
		m, err := RbacTokenService.FindByToken(q)
		if err != nil {
			me.Failed(c, err.Error())
			return
		}
		me.Success(c, m)
		return
	}
	me.Failed(c, "参数错误")
}

