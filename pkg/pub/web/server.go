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
 * long.qian 2021-10-02 12:31 创建
 */

/**
 * @author long.qian
 */

package web

import (
	"fmt"
	"github.com/arl/statsviz"
	"github.com/gin-gonic/gin"
	"go-lazy-frame/configs"
	"go-lazy-frame/configs/sand_server"
	"go-lazy-frame/pkg/pub/logger"
	"go-lazy-frame/pkg/pub/util"
	"net/http"
)

// StartWebServer 启动web 服务
func StartWebServer(cross bool, router func(engine *gin.Engine)) {
	defer util.RecoverPanic()

	if configs.GeneralConfig.EnableMon {
		// 开启监控 默认配置下访问：http://localhost:5203/debug/statsviz/
		go func() {
			// Create a serve mux and register statsviz handlers.
			mux := http.NewServeMux()
			_ = statsviz.Register(mux)
			logger.Sugar.Info("监控服务启动：", fmt.Sprintf("http://127.0.0.1%s/debug/statsviz/ , http://%s%s/debug/statsviz/",
				configs.GeneralConfig.MonPort, util.IpUtil.GetLocalIp(), configs.GeneralConfig.MonPort,
			))
			err := http.ListenAndServe(sand_server.Config.MonPort, mux)
			if err != nil {
				fmt.Println("监控服务启动失败", err)
			}
		}()

	}

	r := gin.Default()
	if cross {
		// 设置全局跨域访问
		r.Use(CrossHandler())
	}

	if configs.IsProdEnv() {
		gin.SetMode(gin.ReleaseMode)
	}

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())
	router(r)
	logger.Sugar.Info("Web 服务启动，端口：", configs.GeneralConfig.WebPort)
	err := r.Run("0.0.0.0" + configs.GeneralConfig.WebPort)
	if err != nil {
		panic(err)
	}
}

// CrossHandler 跨域访问
func CrossHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken")
		context.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "OK")
		}

		//处理请求
		context.Next()
	}
}
