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
 * long.qian 2021-10-02 12:45 创建
 */

/**
 * @author long.qian
 */

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-lazy-frame/go-lazy-frame/configs"
	go_lazy_frame_example "github.com/go-lazy-frame/go-lazy-frame/configs/go-lazy-frame-example"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/web_router"
	"os"
	"path"
	"strings"
)

var (
	// 注册所有的 Controller Struts
	controllers = map[string]interface{}{
		//以下代码的格式请勿改动，否则代码生成工具无法注册路由
		// === 路由注册开始 ===
		// === 路由注册结束 ===
	}
)

// Router 注册路由
func Router(r *gin.Engine) {
	initDoc()
	web_router.RegisterRouter(r, controllers)
}

// 更新接口文档
func initDoc() {
	if configs.IsLocalEnv() {
		projectHome := os.Getenv("projectHome")
		docsPath := os.Getenv("docsPath")
		if projectHome != "" && docsPath != "" {
			docsPath = strings.ReplaceAll(docsPath, "{ProjectHome}", projectHome)
			logger.Sugar.Info("自动更新接口文档，project.home：", projectHome)
			logger.Sugar.Info("自动更新接口文档，docs.path：", docsPath)
			dirs := []string{
				path.Join(projectHome, "cmd/go-lazy-frame-example") + "/",
				path.Join(projectHome, "pkg/pub/query") + "/",
				path.Join(projectHome, "pkg/pub/web") + "/",
				path.Join(projectHome, "pkg/pub/vo") + "/",
			}
			web_router.UpdateDocsGo(dirs, docsPath, go_lazy_frame_example.Config.Config, &SwaggerDoc)
		}
	}
}
