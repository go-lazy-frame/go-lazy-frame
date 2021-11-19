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
 * long.qian 2021-10-21 13:53 创建
 */

/**
 * @author long.qian
 */

package update

import "reflect"

// GenerateUpdatesMap 根据 UpdateDTO 生成更新 map
func GenerateUpdatesMap(updateDto interface{}) map[string]interface{} {
	var m = make(map[string]interface{})
	fieldType := reflect.TypeOf(updateDto)
	fieldValue := reflect.ValueOf(updateDto)
	numberField := fieldType.NumField()
	for i := 0; i < numberField; i++ {
		if !fieldValue.Field(i).IsNil() {
			m[fieldType.Field(i).Name] = fieldValue.Field(i).Interface()
		}
	}
	return m
}
