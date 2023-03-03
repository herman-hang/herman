package command

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/herman-hang/herman/servers/settings"
	"github.com/spf13/cobra"
)

// MigrationCmd 数据库迁移
var (
	direction    string
	MigrationCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Run database migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			mysqlConfig := settings.Config.MysqlConfig
			db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true",
				mysqlConfig.User,
				mysqlConfig.Password,
				mysqlConfig.Host,
				mysqlConfig.Port,
				mysqlConfig.Dbname,
			))
			driver, err := mysql.WithInstance(db, &mysql.Config{})
			if err != nil {
				return err
			}
			m, err := migrate.NewWithDatabaseInstance("file://database/migrations", "mysql", driver)
			if err != nil {
				return err
			}

			if err = migrateFunc(m, direction); err != nil {
				return err
			}
			return nil
		},
	}
)

// init 命令参数绑定
// @return void
func init() {
	StartServerCmd.Flags().StringVarP(&direction, "direction", "d", "up", "Database migration")
}

// migrateFunc 数据库迁移
// @param *migrate.Migrate migrate 迁移实例
// @param string direction 迁移方向
// @return error
func migrateFunc(migrate *migrate.Migrate, direction string) error {
	switch direction {
	case "up": // 执行迁移
		if err := migrate.Up(); err != nil {
			return err
		}
	case "down": // 执行回滚
		if err := migrate.Down(); err != nil {
			return err
		}
	default:
		return errors.New("invalid migration direction")
	}
	return nil
}
