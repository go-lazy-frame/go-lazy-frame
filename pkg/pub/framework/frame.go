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
 * long.qian 2021-10-15 13:26 创建
 */

/**
 * @author long.qian
 */

package framework

import (
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/db"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
	"time"
)

// WaitingFrameworkInitialized 等待框架初始化完成，如数据库，日志等。调用此方法，若框架正在初始化中，则会等待，初始化完成后，方法结束
func WaitingFrameworkInitialized() {
	for {
		if db.DBInitialized && logger.LoggerInitialized {
			return
		}
		time.Sleep(time.Millisecond * 200)
	}
}
