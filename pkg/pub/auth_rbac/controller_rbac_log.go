/**
* @author long.qian
*/

package auth_rbac

// =================================================================================
// You can edit this file about RbacLog controller code.
// =================================================================================

import (
	"github.com/gin-gonic/gin"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/util"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/web"
)

// RbacLogController 操作日志
// Controller 规范：
// 每一个控制器接口方法，都要对应一个 WebXxx 的字段定义（其中的 Xxx 为下面的 struct 方法，既访问接口），用于描述该接口的定义，可参考 controller_rbac_log.go
// 提示：接口变动或接口参数变动，需要执行对应 xxx 脚本进行接口文档的同步更新
type RbacLogController struct {
	web.CommonController
	WebQueryPage interface{} `url:"/rbac_log/query_page" method:"post"`
	WebQuery interface{} `url:"/rbac_log/query" method:"post"`
	WebCreate interface{} `url:"/rbac_log/create" method:"post"`
	WebUpdate interface{} `url:"/rbac_log/update" method:"post"`
	WebFindById interface{} `url:"/rbac_log/find_by_id" method:"get"`
}

// Create
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacLogCreateUsingPOST
// @Tags 操作日志
// @Summary 操作日志创建
// @Description 操作日志创建
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param request body RbacLogCreateDto{} true "创建"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number(float),boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_log/create [post]
func (me RbacLogController) Create(c *gin.Context) {
	d := RbacLogCreateDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	_, err := RbacLogService.CreateRbacLog(d)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, "OK")
}

// Update
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacLogUpdateUsingPOST
// @Tags 操作日志
// @Summary 操作日志更新
// @Description 操作日志更新，根据 id 更新
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param request body RbacLogUpdateDto{} true "更新"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number(float),boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_log/update [post]
func (me RbacLogController) Update(c *gin.Context) {
	d := RbacLogUpdateDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	err := RbacLogService.UpdateRbacLog(d)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, "OK")
}

// QueryPage
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacLogQueryPageUsingPOST
// @Tags 操作日志
// @Summary 操作日志分页查询
// @Description 操作日志分页查询<br>请求示例：<br>{<br>&nbsp;&nbsp;"start": 0,<br>&nbsp;&nbsp;"limit": 20<br>}<br>其他查询条件，可根据以下条件字段酌情添加查询条件<br>注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param request body RbacLogPageDto{} true "条件查询"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number(float),boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_log/query_page [post]
func (me RbacLogController) QueryPage(c *gin.Context) {
	d := RbacLogPageDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	r, err := RbacLogService.QueryPageRbacLog(d)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, r)
}

// Query
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacLogQueryUsingPOST
// @Tags 操作日志
// @Summary 操作日志查询
// @Description 操作日志查询记录<br>请求示例：<br>{}<br>其他查询条件，可根据以下条件字段酌情添加查询条件<br>注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param request body RbacLogQueryDto{} true "条件查询"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number(float),boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_log/query [post]
func (me RbacLogController) Query(c *gin.Context) {
	d := RbacLogQueryDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	r, err := RbacLogService.QueryRbacLog(d)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, r)
}

// FindById
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacLogFindByIdUsingGET
// @Tags 操作日志
// @Summary 操作日志ById
// @Description 操作日志查询ById
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param id query integer true "查询条件id"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number,boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_log/find_by_id [get]
func (me RbacLogController) FindById(c *gin.Context) {
	if id,ok := c.GetQuery("id"); ok {
		m, err := RbacLogService.FindById(util.NumberUtil.StringToUInt(id))
		if err != nil {
			me.Failed(c, err.Error())
			return
		}
		me.Success(c, m)
		return
	}
	me.Failed(c, "参数错误")
}

