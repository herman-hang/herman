package system

import "github.com/brianvoe/gofakeit/v6"

// System 系统设置信息填充器
func System() map[string]interface{} {
	return map[string]interface{}{
		"address":     gofakeit.Sentence(10),
		"copyright":   gofakeit.Sentence(10),
		"description": "Herman基于Gin，Casbin，Kafka，Mysql，Redis，Zap，Cobra，Grom开发，专注于后端快速上手的一款开源，简单，轻量框架。 ",
		"email":       gofakeit.Email(),
		"icoFileId":   gofakeit.Number(1, 100),
		"isWebsite":   gofakeit.RandomInt([]int{1, 2}),
		"keywords":    gofakeit.Sentence(5),
		"logoFileId":  gofakeit.Number(1, 100),
		"name":        gofakeit.Name(),
		"record":      gofakeit.Sentence(1),
		"telephone":   gofakeit.PhoneFormatted(),
		"title":       gofakeit.Sentence(5),
	}
}
