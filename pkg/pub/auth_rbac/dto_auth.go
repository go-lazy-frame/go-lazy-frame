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
 * long.qian 2021-10-10 22:31 创建
 */

/**
 * @author long.qian
 */

package auth_rbac

// LoginDto 登陆
type LoginDto struct {
	LoginName string `json:"loginName"` // 登录名
	LoginPswd string `json:"loginPassword"` // 登陆密码
}
