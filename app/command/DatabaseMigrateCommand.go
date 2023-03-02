package command

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/herman-hang/herman/servers/settings"
	"github.com/spf13/cobra"
)

var MigrationCmd = &cobra.Command{
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
		m, err := migrate.NewWithDatabaseInstance(
			"file://database/migrations",
			"mysql",
			driver,
		)
		if err != nil {
			return err
		}

		if err := up(m); err != nil {
			return err
		}
		return nil
	},
}

// up 执行迁移
// @param *migrate.Migrate migrate 迁移对象
// @return error 错误信息
func up(migrate *migrate.Migrate) error {
	if err := migrate.Up(); err != nil {
		return err
	}
	return nil
}

// down 执行回滚
// @param *migrate.Migrate migrate 迁移对象
// @return error 错误信息
func down(migrate *migrate.Migrate) error {
	if err := migrate.Down(); err != nil {
		return err
	}
	return nil
}
