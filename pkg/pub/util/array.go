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
 * long.qian 2021-10-10 17:26 创建
 */

/**
 * @author long.qian
 */

package util

var (
	ArrayUtil = new(arrayUtil)
)

type arrayUtil struct {

}

func (me *arrayUtil) IsExistIntArray(array *[]int64, obj int64) bool {
	for _, i := range *array {
		if i == obj {
			return true
		}
	}
	return false
}

func (me *arrayUtil) IsExistStringArray(array *[]string, obj string) bool {
	for _, i := range *array {
		if i == obj {
			return true
		}
	}
	return false
}

// SubtractIntArray 整型数组求差集 arr1 - arr2
func (me *arrayUtil) SubtractIntArray(arr1 *[]int64, arr2 *[]int64) []int64 {
	var arr []int64
	for _, item1 := range *arr1 {
		if !me.IsExistIntArray(arr2, item1) {
			arr = append(arr, item1)
		}
	}
	return arr
}
