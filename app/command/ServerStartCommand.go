package command

import (
	"github.com/herman-hang/herman/servers"
	"github.com/spf13/cobra"
)

// StartServerCmd 服务启动
var (
	host           string
	port           uint
	IsMigrate      bool
	StartServerCmd = &cobra.Command{
		Use:          "server",
		Short:        "This is a herman service",
		SilenceUsage: true,
		Example:      "herman server --host=0.0.0.0 --port=8000",
		Run: func(cmd *cobra.Command, args []string) {
			servers.NewServer(host, port)
		},
	}
)

// init 命令参数绑定
func init() {
	StartServerCmd.Flags().StringVarP(&host, "host", "H", "127.0.0.1", "HTTP server host")
	StartServerCmd.Flags().UintVarP(&port, "port", "p", 8000, "HTTP server port")
	StartServerCmd.Flags().BoolVarP(&IsMigrate, "migrate", "m", false, "Whether to migrate the database")
}
