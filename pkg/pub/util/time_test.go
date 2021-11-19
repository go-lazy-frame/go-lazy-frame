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
 * long.qian 2021-10-25 16:21 创建
 */

/**
 * @author long.qian
 */

package util

import (
	"fmt"
	"testing"
)

func Test_timeUtil_IsWeekend(t *testing.T) {
	fmt.Println(TimeUtil.IsWeekend(*TimeUtil.TimeParse("2021-10-25 04:23:58")))
	fmt.Println(TimeUtil.IsWeekend(*TimeUtil.TimeParse("2021-10-26 04:23:58")))
	fmt.Println(TimeUtil.IsWeekend(*TimeUtil.TimeParse("2021-10-27 04:23:58")))
	fmt.Println(TimeUtil.IsWeekend(*TimeUtil.TimeParse("2021-10-28 04:23:58")))
	fmt.Println(TimeUtil.IsWeekend(*TimeUtil.TimeParse("2021-10-29 04:23:58")))
	fmt.Println(TimeUtil.IsWeekend(*TimeUtil.TimeParse("2021-10-30 04:23:58")))
	fmt.Println(TimeUtil.IsWeekend(*TimeUtil.TimeParse("2021-10-31 04:23:58")))
}

func Test_timeUtil_GetCurrentMonthStartEnd(t *testing.T) {
	fmt.Println(TimeUtil.GetCurrentMonthStart())
	fmt.Println(TimeUtil.GetCurrentMonthEnd())
}
