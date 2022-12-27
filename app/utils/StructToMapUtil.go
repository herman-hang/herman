package utils

import (
	"fmt"
	"reflect"
)

// ToMap 结构体转为Map[string]interface{}
// @param interface in 待转结构体
// @param string tagName 根据指定结构体标签作为key
// @return map[string]interface{} error 返回一个结构体转换好的map值和错误信息
func ToMap(in interface{}, tagName string) (out map[string]interface{}, err error) {
	out = make(map[string]interface{})
	v := reflect.ValueOf(in)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// 非结构体返回错误提示
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段,指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}

	return out, nil
}
