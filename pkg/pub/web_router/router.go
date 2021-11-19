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
 * long.qian 2021-10-02 15:33 创建
 */

/**
 * @author long.qian
 */

package web_router

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-lazy-frame/go-lazy-frame/configs"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/auth_rbac"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/util"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swaggo/swag"
	"github.com/swaggo/swag/gen"
	"net/http"
	"os"
	"path"
	"reflect"
	"strings"
)

const (
	docsCode = `
package router

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var SwaggerDoc = {DOC_JSON}

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "{HOST}",
	BasePath:    "{BASE_PATH}",
	Schemes:     []string{},
	Title:       "LazyFrame",
	Description: "后端系统接口",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(SwaggerDoc)
	if err != nil {
		return SwaggerDoc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return SwaggerDoc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
`
)

var (
	format = "=================================================="
	//go:embed static
	staticFs embed.FS
)

// UpdateDocsGo 更新接口文档
func UpdateDocsGo(dirs []string, docsPath string, config configs.Config, swaggerDoc *string) {
	if _, err := os.Stat(docsPath); os.IsNotExist(err) {
		fmt.Printf("接口文档自动更新失败：docs go 文件: %s 不存在，\n", docsPath)
		return
	}

	if configs.GeneralConfig.EnableRbacAuth {
		projectHome := os.Getenv("projectHome")
		dirs = append(dirs, path.Join(projectHome, "vendor/github.com/go-lazy-frame/go-lazy-frame/pkg/pub/auth_rbac")+"/")
	}

	c := &gen.Config{
		SearchDir:          strings.Join(dirs, ","),
		MainAPIFile:        "./main.go",
		PropNamingStrategy: "",
	}
	searchDirs := strings.Split(c.SearchDir, ",")
	for _, searchDir := range searchDirs {
		if _, err := os.Stat(searchDir); os.IsNotExist(err) {
			logger.Sugar.Errorf("目录: %s 不存在\n", searchDir)
			return
		}
	}

	p := swag.New(swag.SetMarkdownFileDirectory(c.MarkdownFilesDir),
		swag.SetExcludedDirsAndFiles(c.Excludes),
		swag.SetCodeExamplesDirectory(c.CodeExampleFilesDir),
		swag.SetStrict(c.Strict))
	p.PropNamingStrategy = c.PropNamingStrategy
	p.ParseVendor = c.ParseVendor
	p.ParseDependency = c.ParseDependency
	p.ParseInternal = c.ParseInternal

	if err := p.ParseAPIMultiSearchDir(searchDirs, c.MainAPIFile, c.ParseDepth); err != nil {
		logger.Sugar.Error("接口文档自动更新失败", err)
		return
	}
	swagger := p.GetSwagger()

	b, err := jsonIndent(swagger)
	if err != nil {
		logger.Sugar.Error("接口文档自动更新失败", err)
		return
	}

	docJson := string(b)
	content := strings.ReplaceAll(docsCode, "{DOC_JSON}", "`"+docJson+"`")
	content = strings.ReplaceAll(content, "{BASE_PATH}", config.ApiPrefix)
	content = strings.ReplaceAll(content, "{HOST}", "localhost"+config.WebPort)
	// 持久化到文件
	err = os.WriteFile(docsPath, []byte(content), 0755)
	if err != nil {
		logger.Sugar.Error("接口文档自动更新失败", err)
		return
	}
	// 使当前实例生效
	*swaggerDoc = docJson
	logger.Sugar.Info("接口文档自动更新成功")
}

func jsonIndent(data interface{}) ([]byte, error) {
	return json.MarshalIndent(data, "", "    ")
}

// RegisterRouter 注册路由
func RegisterRouter(r *gin.Engine, controllers map[string]interface{}) {
	group := r.Group(configs.GeneralConfig.ApiPrefix)
	if configs.GeneralConfig.EnableRbacAuth {
		group.Use(auth_rbac.RbacHandler())
		auth_rbac.ModelAutoMigrate()
		controllers["AuthController"] = auth_rbac.AuthController{}
		controllers["RbacLogController"] = auth_rbac.RbacLogController{}
		controllers["RbacRoleController"] = auth_rbac.RbacRoleController{}
		controllers["RbacUserController"] = auth_rbac.RbacUserController{}
		controllers["RbacTokenController"] = auth_rbac.RbacTokenController{}
		controllers["RbacPermissionsController"] = auth_rbac.RbacPermissionsController{}
	}
	for controllerName, controller := range controllers {
		fieldType := reflect.TypeOf(controller)
		fieldValue := reflect.ValueOf(controller)

		numField := fieldType.NumField()
		for i := 0; i < numField; i++ {
			field := fieldType.Field(i)
			if strings.HasPrefix(field.Name, "Web") {
				methodName := strings.Replace(field.Name, "Web", "", 1)
				methodValue := fieldValue.MethodByName(methodName)
				if reflect.ValueOf(methodValue).IsZero() {
					logger.Sugar.Errorf("【Controller %s】没有定义方法 %s，停止注册该路由",
						controllerName, methodName,
					)
					continue
				}

				logger.Sugar.Infof("【Controller %s】开始注册路由：%s", controllerName, methodName)
				tag := field.Tag
				url := tag.Get("url")
				switch tag.Get("method") {
				case "*":
					group.Any(url, func(context *gin.Context) {
						methodValue.Call([]reflect.Value{reflect.ValueOf(context)})
					})
					break
				case "post":
					group.POST(url, func(context *gin.Context) {
						methodValue.Call([]reflect.Value{reflect.ValueOf(context)})
					})
					break
				case "get":
					group.GET(url, func(context *gin.Context) {
						methodValue.Call([]reflect.Value{reflect.ValueOf(context)})
					})
					break
				case "put":
					group.PUT(url, func(context *gin.Context) {
						methodValue.Call([]reflect.Value{reflect.ValueOf(context)})
					})
					break
				case "delete":
					group.DELETE(url, func(context *gin.Context) {
						methodValue.Call([]reflect.Value{reflect.ValueOf(context)})
					})
					break
				case "options":
					group.OPTIONS(url, func(context *gin.Context) {
						methodValue.Call([]reflect.Value{reflect.ValueOf(context)})
					})
					break
				case "patch":
					group.PATCH(url, func(context *gin.Context) {
						methodValue.Call([]reflect.Value{reflect.ValueOf(context)})
					})
					break
				case "head":
					group.HEAD(url, func(context *gin.Context) {
						methodValue.Call([]reflect.Value{reflect.ValueOf(context)})
					})
					break
				default:
					logger.Sugar.Errorf("【Controller %s】方法 %s tag[method]定义错误：只能为 *,post,get,put,delete,options,patch,head。请检查字段 %s 的 tag 定义",
						controllerName, methodName, field.Name)
				}

			}
		}
	}

	if configs.GeneralConfig.EnableDoc {
		// 启动接口文档
		r.StaticFS("/doc", http.FS(staticFs))
		logger.Sugar.Infof("接口文档地址：http://127.0.0.1%s/doc/static/doc.html , http://%s%s/doc/static/doc.html",
			configs.GeneralConfig.WebPort, util.IpUtil.GetLocalIp(), configs.GeneralConfig.WebPort,
		)

		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		r.GET("/swagger-resources", func(c *gin.Context) {
			c.JSON(http.StatusOK, []map[string]interface{}{
				{
					"name":           "default",
					"url":            "/swagger/doc.json",
					"swaggerVersion": "2.0",
					"location":       "/swagger/doc.json",
				},
			})
		})
	}

}
