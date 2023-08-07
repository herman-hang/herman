package command

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/herman-hang/herman/kernel/app"
	"github.com/spf13/cobra"
)

// HermanVersionCmd 获取herman版本号
var (
	HermanVersionCmd = &cobra.Command{
		Use:          "version",
		Short:        "Get herman version",
		Example:      "herman version",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf(`Herman version: %v`, color.GreenString(app.Version))
			return nil
		},
	}
)
