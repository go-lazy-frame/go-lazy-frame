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
 * long.qian 2021-10-03 14:48 创建
 */

/**
 * @author long.qian
 */

package vo

// ResponseResult 响应结果
type ResponseResult struct {
	// 响应码 0为成功，其他都为错误码
	Code string `json:"code"`
	// 错误信息
	Message string `json:"message"`
	// 响应的结果数据
	Data interface{} `json:"data"`
}
