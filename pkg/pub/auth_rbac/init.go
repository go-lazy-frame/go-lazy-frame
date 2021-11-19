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
 * long.qian 2021-10-15 13:23 创建
 */

/**
 * @author long.qian
 */

package auth_rbac

import (
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/framework"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
)

func init() {
	go func() {
		// 等待系统资源初始化完毕
		framework.WaitingFrameworkInitialized()
		// 初始化 admin 用户
		_ ,err := RbacUserService.FindByLoginName("admin")
		if err != nil {
			_, err := RbacUserService.CreateRbacUser(AddRbacUserDto{
				Admin:      true,
				LoginName:  "admin",
				LoginPswd:  "123456",
				RepeatLoginPswd: "123456",
				SuperAdmin: true,
			})
			if err != nil {
				logger.Sugar.Error("用户初始化失败", err)
			} else {
				logger.Sugar.Info("超级管理员用户初始化成功，账号：admin，密码：123456")
			}
		}
	}()
}
