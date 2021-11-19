/**
* @author long.qian
*/

package auth_rbac

// =================================================================================
// You can edit this file about RbacRole controller code.
// =================================================================================

import (
	"github.com/gin-gonic/gin"
	"go-lazy-frame/pkg/pub/util"
	"go-lazy-frame/pkg/pub/web"
)

// RbacRoleController 角色
// Controller 规范：
// 每一个控制器接口方法，都要对应一个 WebXxx 的字段定义（其中的 Xxx 为下面的 struct 方法，既访问接口），用于描述该接口的定义，可参考 controller_rbac_role.go
// 提示：接口变动或接口参数变动，需要执行对应 xxx 脚本进行接口文档的同步更新
type RbacRoleController struct {
	web.CommonController
	WebQueryPage interface{} `url:"/rbac_role/query_page" method:"post"`
	WebQuery interface{} `url:"/rbac_role/query" method:"post"`
	WebCreate interface{} `url:"/rbac_role/create" method:"post"`
	WebUpdate interface{} `url:"/rbac_role/update" method:"post"`
	WebFindById interface{} `url:"/rbac_role/find_by_id" method:"get"`
	WebFindByRoleName interface{} `url:"/rbac_role/find_by_role_name" method:"get"`
}

// Create
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacRoleCreateUsingPOST
// @Tags 角色
// @Summary 角色创建
// @Description 角色创建
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param request body RbacRoleCreateDto{} true "创建"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number(float),boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_role/create [post]
func (me RbacRoleController) Create(c *gin.Context) {
	d := RbacRoleCreateDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	_, err := RbacRoleService.CreateRbacRole(d)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, "OK")
}

// Update
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacRoleUpdateUsingPOST
// @Tags 角色
// @Summary 角色更新
// @Description 角色更新，根据 id 更新
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param request body RbacRoleUpdateDto{} true "更新"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number(float),boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_role/update [post]
func (me RbacRoleController) Update(c *gin.Context) {
	d := RbacRoleUpdateDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	err := RbacRoleService.UpdateRbacRole(d)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, "OK")
}

// QueryPage
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacRoleQueryPageUsingPOST
// @Tags 角色
// @Summary 角色分页查询
// @Description 角色分页查询<br>请求示例：<br>{<br>&nbsp;&nbsp;"start": 0,<br>&nbsp;&nbsp;"limit": 20<br>}<br>其他查询条件，可根据以下条件字段酌情添加查询条件<br>注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param request body RbacRolePageDto{} true "条件查询"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number(float),boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_role/query_page [post]
func (me RbacRoleController) QueryPage(c *gin.Context) {
	d := RbacRolePageDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	r, err := RbacRoleService.QueryPageRbacRole(d)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, r)
}

// Query
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacRoleQueryUsingPOST
// @Tags 角色
// @Summary 角色查询
// @Description 角色查询记录<br>请求示例：<br>{}<br>其他查询条件，可根据以下条件字段酌情添加查询条件<br>注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param request body RbacRoleQueryDto{} true "条件查询"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number(float),boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_role/query [post]
func (me RbacRoleController) Query(c *gin.Context) {
	d := RbacRoleQueryDto{}
	if err := me.BindBodyJson(c, &d); err != nil {
		me.Failed(c, err.Error())
		return
	}
	r, err := RbacRoleService.QueryRbacRole(d)
	if err != nil {
		me.Failed(c, err.Error())
		return
	}
	me.Success(c, r)
}

// FindById
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacRoleFindByIdUsingGET
// @Tags 角色
// @Summary 角色ById
// @Description 角色查询ById
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param id query integer true "查询条件id"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number,boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_role/find_by_id [get]
func (me RbacRoleController) FindById(c *gin.Context) {
	if id,ok := c.GetQuery("id"); ok {
		m, err := RbacRoleService.FindById(util.NumberUtil.StringToUInt(id))
		if err != nil {
			me.Failed(c, err.Error())
			return
		}
		me.Success(c, m)
		return
	}
	me.Failed(c, "参数错误")
}

// FindByRoleName
// 注意：以下的 id 必须设置，且必须全局唯一，否则接口文档页面无法正常显示
// @id RbacRoleFindByRoleNameUsingGET
// @Tags 角色
// @Summary 角色ByRoleName
// @Description 角色查询ByRoleName<br/>字符串类型字段：空串或不传都代表忽略该查询，如果要指定空串为条件，则指定为 STRING__BLANK，例如 { "name": "STRING__BLANK" }
// @Accept json
// @Produce  json
// @Param token header string true "登陆成功后的授权 Token，后续的所有接口header，都要带上 token"
// @Param roleName query string true "查询条件值"
// 		参数名 参数类型(query,path,header,body,formData) 数据类型(string,integer,number,boolean,user defined struct) 是否必传 描述
// @Success 200 {object} vo.ResponseResult{}
// @Failure 500 {object} vo.ResponseResult{}
// @Router /rbac_role/find_by_role_name [get]
func (me RbacRoleController) FindByRoleName(c *gin.Context) {
	if q,ok := c.GetQuery("roleName"); ok {
		m, err := RbacRoleService.FindByRoleName(q)
		if err != nil {
			me.Failed(c, err.Error())
			return
		}
		me.Success(c, m)
		return
	}
	me.Failed(c, "参数错误")
}

