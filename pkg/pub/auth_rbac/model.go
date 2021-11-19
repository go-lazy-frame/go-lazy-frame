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
 * long.qian 2021-10-10 20:57 创建
 */

/**
 * @author long.qian
 */

package auth_rbac

import (
	"github.com/go-lazy-frame/go-lazy-frame/configs"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/db"
)

func ModelAutoMigrate() {
	// 权限
	if configs.GeneralConfig.EnableRbacAuth {
		err := db.DB.AutoMigrate(
			RbacUser{},
			RbacLog{},
			RbacPermissions{},
			RbacToken{},
			RbacRole{},
		)
		if err != nil {
			panic(err)
		}
	}
}
