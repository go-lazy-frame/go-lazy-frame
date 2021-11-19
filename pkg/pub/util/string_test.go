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
 * long.qian 2021-11-16 14:40 创建
 */

/**
 * @author long.qian
 */

package util

import (
	"fmt"
	"testing"
)

func Test_stringUtil_EditDistanceDP(t *testing.T) {
	a := "赣AV1918"
	//b := ""
	b := "赣AN9178"

	fmt.Println(StringUtil.getMaxEditDistance(a, b, true))
	fmt.Println(StringUtil.EditDistanceDP(a, b, true))
	fmt.Println(StringUtil.MatchingRate(a, b, true))
	fmt.Println(StringUtil.MatchingRate(b, a, true))
}
