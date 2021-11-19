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

// LoginVo 登陆
type LoginVo struct {
	LoginName   string              `json:"loginName"`
	Nickname    string              `json:"nickname"`
	Phone       string              `json:"phone"`
	Token       string              `json:"token"`
	Permissions []RbacPermissionsVo `json:"permissions"`
}
