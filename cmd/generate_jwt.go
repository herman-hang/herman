package command

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// GenerateJwtCmd 随机生成JWT令牌
var (
	GenerateJwtCmd = &cobra.Command{
		Use:          "jwt:secret",
		Short:        "Generate a secret for JWT",
		Example:      "jwt:secret",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			// 生成32字节的随机字节数组
			key := make([]byte, 32)
			if _, err := rand.Read(key); err != nil {
				return err
			}
			// 对随机字节数组进行Base64编码
			encodedKey := base64.StdEncoding.EncodeToString(key)

			fmt.Println("Jwt secret:", color.GreenString(encodedKey))
			return nil
		},
	}
)
