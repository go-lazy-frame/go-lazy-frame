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
 * long.qian 2021-10-20 11:04 创建
 */

/**
 * @author long.qian
 */

package util

import (
	"fmt"
	"testing"
	"time"
)

func Test_ipUtil_IsPing(t *testing.T) {
	fmt.Println(IpUtil.IsPing("192.168.0.4", time.Second))
}
