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
 * long.qian 2021-09-30 15:23 创建
 */

/**
 * @author long.qian
 */

package main

import (
	"github.com/go-lazy-frame/go-lazy-frame/configs"
	"github.com/go-lazy-frame/go-lazy-frame/internal/go-lazy-frame-example/router"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/db"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/web"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

// @title Go-Lazy_frame-Example
// @version 1.0
// @description Go-Lazy_frame-Example
// @host localhost:8890
// @contact.name qianlong
// @contact.email 642321251@qq.com
// @BasePath /api/v1
func main() {
	logger.Init(configs.GeneralConfig.AppName)
	db.Init()

	// 启动 Web 服务
	go func() {
		web.StartWebServer(true, router.Router)
	}()

	killListener()

	select {}
}

func killListener() {
	if runtime.GOOS == "darwin" || runtime.GOOS == "freebsd" || runtime.GOOS == "linux" {
		//创建监听退出chan
		c := make(chan os.Signal)
		//监听指定信号 ctrl+c kill
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		go func() {
			for s := range c {
				switch s {
				case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
					logger.Sugar.Info("成功关闭系统 ", s)
					os.Exit(0)
				//case syscall.SIGUSR1:
				//case syscall.SIGUSR2:
				default:
					logger.Sugar.Info("系统信号other ", s)
				}
			}
		}()
	}
}
