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
 * long.qian 2021-10-05 13:15 创建
 */

/**
 * @author long.qian
 */

package query

import (
	"fmt"
	"go-lazy-frame/pkg/pub/util"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

// PageHandler 分页处理
func PageHandler(tx *gorm.DB, page *Page) {
	if page.PageNum < 1 {
		page.PageNum = 1
	}
	if page.PageSize <= 0 {
		page.PageSize = 20
	}
	start := page.PageSize * (page.PageNum - 1)
	limit := page.PageSize
	tx.Offset(start)
	tx.Limit(limit)
}

// WhereHandler 查询条件处理
func WhereHandler(tx *gorm.DB, condition interface{}) error {
	// 查询条件组装
	if condition == nil {
		return nil
	}
	fieldType := reflect.TypeOf(condition)
	fieldValue := reflect.ValueOf(condition)

	numField := fieldType.NumField()
	for i := 0; i < numField; i++ {
		field := fieldType.Field(i)
		if field.Name == "Page" {
			continue
		}
		if strings.HasSuffix(fieldType.Name(), "PageDto") && strings.HasSuffix(field.Name, "QueryDto") {
			// struct 内嵌的查询结构体（XxxQueryDto）
			v := fieldValue.Field(i).Interface()
			return WhereHandler(tx, v)
		} else {
			err := validCondition(condition)
			if err != nil {
				return err
			}
			if field.Name == "OrderBy" {
				continue
			}
			queryType := field.Tag.Get("type")
			column := field.Tag.Get("field")
			if queryType != "" {
				if column == "" {
					return fmt.Errorf("%s %s: 缺少 field tag", fieldType.Name(), field.Name)
				}
				v := fieldValue.Field(i)
				if queryType == "between" {
					// between 是数组，判断数组长度
					if !v.IsZero() {
						if vRange, ok := v.Interface().([]interface{}); ok {
							if len(vRange) != 0 {
								if len(vRange) == 2 {
									tx.Where(column + " between ? and ?", vRange[0], vRange[1])
								} else {
									return fmt.Errorf("%s %s: between 类型，数组元素个数必须为 2", fieldType.Name(), field.Name)
								}
							}
						} else {
							return fmt.Errorf("%s %s: between 类型，字段类型必须是数组", fieldType.Name(), field.Name)
						}
					}
				} else {
					if !v.IsNil() {
						switch queryType {
						case "equal":
							tx.Where(column + " = ?", v.Interface())
							break
						case "in":
							if !v.IsZero() {
								if vRange, ok := v.Interface().([]interface{}); ok {
									if len(vRange) > 0 {
										// in 字段值为数组，且长度大于 0
										tx.Where(column + " in ?", v.Interface())
									}
								}
							}
							break
						case "gt":
							tx.Where(column + " > ?", v.Interface())
							break
						case "gte":
							tx.Where(column + " >= ?", v.Interface())
							break
						case "lt":
							tx.Where(column + " < ?", v.Interface())
							break
						case "lte":
							tx.Where(column + " <= ?", v.Interface())
							break
						case "likeRight":
							// likeRight: 左匹配(xxx%)
							tx.Where(column + " like ?", fmt.Sprintf("%v", v.Interface()) + "%")
							break
						case "likeLeft":
							tx.Where(column + " like ?", "%" + fmt.Sprintf("%v", v.Interface()))
							// likeLeft: 右匹配(%xxx)
							break
						case "likeMiddle":
							tx.Where(column + " like ?", "%" + fmt.Sprintf("%v", v.Interface()) + "%")
							// likeMiddle: 模糊匹配(%xxx%)
							break
						default:
							return fmt.Errorf("未知的 type 类型 %s：\n", queryType)
						}
					}
				}
			} else {
				return fmt.Errorf("字段 %s tag 没有定义 type 值\n", field.Name)
			}
		}
	}

	return nil
}

// OrderHandler 查询排序处理
func OrderHandler(tx *gorm.DB, condition interface{}) error {
	fieldValue := reflect.ValueOf(condition)
	field := fieldValue.FieldByName("OrderBy")
	if field.IsValid() {
		var orderBy []string
		v := field.Interface()
		k := reflect.TypeOf(v).Kind()
		if k == reflect.Slice || k == reflect.Array {
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := s.Index(i).String()
				o = strings.TrimSpace(o)
				ss := strings.Split(o, " ")
				if len(ss) > 2 {
					return fmt.Errorf("字段 %s 排序定义非法\n", ss[0])
				}
				if len(ss) == 2 {
					t := strings.ToLower(ss[1])
					if t != "desc" && t != "asc" {
						return fmt.Errorf("字段 %s 排序只能为 desc 或 asc\n", ss[0])
					}
				}
				orderBy = append(orderBy, o)
			}
		} else {
			return fmt.Errorf("类型 %s 的定义字段 OrderBy 只能为 []string 类型\n", fieldValue.Type().Name())
		}
		if len(orderBy) > 0 {
			for _, s := range orderBy {
				tx.Order(s)
			}
		} else {
			// 默认排序：id desc
			tx.Order("id desc")
		}
	} else {
		return fmt.Errorf("类型 %s 没有定义字段 OrderBy []string\n", fieldValue.Type().Name())
	}
	return nil
}

// 条件可用性检查
func validCondition(condition interface{}) error {
	var columns []string
	fieldType := reflect.TypeOf(condition)
	fieldValue := reflect.ValueOf(condition)
	numField := fieldType.NumField()
	for i := 0; i < numField; i++ {
		v := fieldValue.Field(i)
		field := fieldType.Field(i)
		column := field.Tag.Get("field")
		if column != "" {
			queryType := field.Tag.Get("type")
			if queryType == "between" {
				// between 是数组，判断数组长度
				if !v.IsZero() {
					if vRange, ok := v.Interface().([]interface{}); ok {
						if len(vRange) > 0 {
							if util.ArrayUtil.IsExistStringArray(&columns, column) {
								return fmt.Errorf("列 %s 的查询条件重复，一个列只能指定一种查询条件", column)
							}
							columns = append(columns, column)
						}
					}
				}
			} else {
				if !v.IsNil() {
					if util.ArrayUtil.IsExistStringArray(&columns, column) {
						return fmt.Errorf("列 %s 的查询条件重复，一个列只能指定一种查询条件", column)
					}
					columns = append(columns, column)
				}
			}
		}
	}
	return nil
}

