package dictionary

import "github.com/brianvoe/gofakeit/v6"

// Detail 数据字典明细填充器
// @return map[string]interface{} 数据字典明细信息
func Detail() map[string]interface{} {
	return map[string]interface{}{
		"name":   gofakeit.Name(),
		"code":   gofakeit.Noun(),
		"value":  gofakeit.Noun(),
		"remark": gofakeit.HackerPhrase(),
		"sort":   gofakeit.Number(0, 100),
		"state":  gofakeit.RandomInt([]int{1, 2}),
	}
}
