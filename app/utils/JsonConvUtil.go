package utils

import (
	"bytes"
	"encoding/json"
	"github.com/fp/fp-gin-framework/app/constants"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Buffer 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

// SnakeJSON 转为下划线JSON
// @param interface{} data 待转数据
// @return converted err 转换完成的数据，错误信息
func SnakeJSON(data interface{}) (converted []byte, err error) {
	// Regexp definitions
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)
	marshalled, err := json.Marshal(data)
	converted = keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)
	return converted, err
}

// CamelJSON 转为驼峰JSON
// @param interface{} data 待转数据
// @return converted err 转换完成的数据，错误信息
func CamelJSON(data interface{}) ([]byte, error) {
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	marshalled, err := json.Marshal(data)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			matchStr := string(match)
			key := matchStr[1 : len(matchStr)-2]
			resKey := LcFirst(CaseToCamel(key))
			return []byte(`"` + resKey + `":`)
		},
	)
	return converted, err
}

// CamelToCase 驼峰式写法转为下划线写法
// @param string data 待转数据
// @return string 返回转换完成的字符串
func CamelToCase(data string) string {
	buffer := &Buffer{Buffer: new(bytes.Buffer)}
	for i, r := range data {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// CaseToCamel 下划线写法转为驼峰写法
// @param string data 待转数据
// @return string 返回转换完成的字符串
func CaseToCamel(data string) string {
	data = strings.Replace(data, "_", " ", -1)
	data = strings.Title(data)
	return strings.Replace(data, " ", "", -1)
}

// UcFirst 首字母大写
// @param string data 待转数据
// @return string 返回转换完成的字符串
func UcFirst(data string) string {
	for i, v := range data {
		return string(unicode.ToUpper(v)) + data[i+1:]
	}
	return ""
}

// LcFirst 首字母小写
// @param string data 待转数据
// @return string 返回转换完成的字符串
func LcFirst(data string) string {
	for i, v := range data {
		return string(unicode.ToLower(v)) + data[i+1:]
	}
	return ""
}

// Append 转换器
// @param i 接收一个任意类型的的数据
// @param *Buffer
func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

// append 将一个字符串追加到缓冲区
// @param string s 接收一个字符串
// @return *Buffer
func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			panic(constants.SystemRAMOut)
		}
	}()
	b.WriteString(s)
	return b
}
