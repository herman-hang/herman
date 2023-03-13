package role

import "github.com/brianvoe/gofakeit/v6"

// Role 角色填充器
func Role() map[string]interface{} {
	return map[string]interface{}{
		"roles": nil,
		"name":  gofakeit.Name(),
		"role":  gofakeit.Username(),
		"state": gofakeit.RandomInt([]int{1, 2}),
		"rules": []map[string]interface{}{
			{
				"path":   gofakeit.PhoneFormatted(),
				"method": gofakeit.RandomString([]string{"POST", "GET", "PUT", "DELETE"}),
				"name":   gofakeit.Name(),
			},
		},
	}
}
