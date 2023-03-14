package dictionary

import (
	"github.com/brianvoe/gofakeit/v6"
)

// Dictionary 数据字典填充器
// @return map[string]interface{} 数据字典信息
func Dictionary() map[string]interface{} {
	return map[string]interface{}{
		"name":   gofakeit.Name(),
		"code":   gofakeit.LetterN(5),
		"remark": gofakeit.HackerPhrase(),
		"state":  gofakeit.RandomInt([]int{1, 2}),
	}
}
