package admin

import (
	"github.com/brianvoe/gofakeit/v6"
)

// Admin 管理员填充器
func Admin() map[string]interface{} {
	return map[string]interface{}{
		"user":         gofakeit.Username(),
		"password":     gofakeit.Password(false, false, true, false, false, 10),
		"photo":        gofakeit.ImageURL(100, 100),
		"name":         gofakeit.Name(),
		"card":         "450981200008272525",
		"sex":          gofakeit.RandomInt([]int{1, 2, 3}),
		"age":          gofakeit.Number(18, 60),
		"region":       gofakeit.Country(),
		"phone":        "18888888888",
		"email":        gofakeit.Email(),
		"introduction": gofakeit.Sentence(10),
		"state":        gofakeit.RandomInt([]int{1, 2}),
		"sort":         gofakeit.Number(1, 100),
	}
}
