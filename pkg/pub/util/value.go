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
 * long.qian 2021-11-28 15:33 创建
 */

/**
 * @author long.qian
 */

package util

var(
	ValueUtil = new(valueUtil)
)

type valueUtil struct {

}

func (receiver *valueUtil) IntPointer(v int) *int {
	return &v
}

func (receiver *valueUtil) UintPointer(v uint) *uint {
	return &v
}

func (receiver *valueUtil) Int64Pointer(v int64) *int64 {
	return &v
}

func (receiver *valueUtil) Int32Pointer(v int32) *int32 {
	return &v
}

func (receiver *valueUtil) Float32Pointer(v float32) *float32 {
	return &v
}

func (receiver *valueUtil) Float64Pointer(v float64) *float64 {
	return &v
}

func (receiver *valueUtil) BoolPointer(v bool) *bool {
	return &v
}

func (receiver *valueUtil) StringPointer(v string) *string {
	return &v
}

