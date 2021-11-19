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
 * long.qian 2021-10-12 14:05 创建
 */

/**
 * @author long.qian
 */

package util

import (
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
	"strconv"
)

var (
	NumberUtil = new(numberUtil)
)

type numberUtil struct {

}

// StringToInt64 字符串转int64类型，转换错误将会输出日志并返回0
func (me *numberUtil) StringToInt64(number string) int64 {
	if number == "" {
		return 0
	}
	i, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		logger.Sugar.Error(err)
		return 0
	}
	return i
}

// StringToInt 字符串转int类型，转换错误将会输出日志并返回0
func (me *numberUtil) StringToInt(number string) int {
	if number == "" {
		return 0
	}
	i, err := strconv.Atoi(number)
	if err != nil {
		logger.Sugar.Error(err)
		return 0
	}
	return i
}

// StringToUInt 字符串转uint类型，转换错误将会输出日志并返回0
func (me *numberUtil) StringToUInt(number string) uint {
	if number == "" {
		return 0
	}
	i, err := strconv.Atoi(number)
	if err != nil {
		logger.Sugar.Error(err)
		return 0
	}
	return uint(i)
}

// StringToFloat64 字符串转float64类型，转换错误将会输出日志并返回0
func (me *numberUtil) StringToFloat64(number string) float64 {
	if number == "" {
		return 0
	}
	i, err := strconv.ParseFloat(number, 64)
	if err != nil {
		logger.Sugar.Error(err)
		return 0
	}
	return i
}
