package menu

import "github.com/brianvoe/gofakeit/v6"

// Menu 菜单填充器
// @return map[string]interface{} 菜单信息
func Menu() map[string]interface{} {
	return map[string]interface{}{
		"pid":  0,
		"name": gofakeit.Name(),
		"path": gofakeit.URL(),
		"method": gofakeit.RandomString([]string{
			"GET", "POST", "PUT", "DELETE",
		}),
		"sort": gofakeit.Number(0, 100),
	}
}
