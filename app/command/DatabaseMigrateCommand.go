package command

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	commandConstant "github.com/herman-hang/herman/app/constants/command"
	"github.com/herman-hang/herman/servers/settings"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// MigrationCmd 数据库迁移
var (
	direction    string
	version      uint
	number       uint
	MigrationCmd = &cobra.Command{
		Use:          "migrate",
		Short:        "Run database migrations",
		Example:      "herman migrate --direction=up --number=1",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := Migrate(direction); err != nil {
				return err
			}
			return nil
		},
	}
)

// init 命令参数绑定
// @return void
func init() {
	// 迁移方式，up和down
	MigrationCmd.Flags().StringVarP(&direction, "direction", "d", "up", "Database migration")
	// 执行指定数据库版本，主要在出现Error: Dirty database version XX.使用
	MigrationCmd.Flags().UintVarP(&version, "version", "v", 0, "Database version")
	// 执行迁移的版本次数，比如回滚1个版本，可以执行herman -d down -n 1，不指定则全部迁移
	MigrationCmd.Flags().UintVarP(&number, "number", "n", 0, "Database migration steps")
}

// Migrate 数据库迁移
// @return error 错误信息
func Migrate(direction string) error {
	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true",
		settings.Config.MysqlConfig.User,
		settings.Config.MysqlConfig.Password,
		settings.Config.MysqlConfig.Host,
		settings.Config.MysqlConfig.Port,
		settings.Config.MysqlConfig.Dbname,
	))
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}
	// 该目录下的文件格式严格要求：版本号_描述.up(down).sql
	dirPath := "database/migrations"
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", dirPath), "mysql", driver)
	if err != nil {
		return err
	}

	// 获取当前最新的迁移版本
	currentVersion, _, _ := m.Version()
	err, fileNames := getFileNames(dirPath, direction)
	if err != nil {
		return err
	}
	if err := selectExec(m, fileNames, currentVersion, direction); err != nil {
		return err
	}
	return nil
}

// selectExec 根据命令执行
// @param *migrate.Migrate m 迁移实例
// @param []string fileNames 文件名切片
// @param uint currentVersion 当前数据库最新版本
// @param string direction 迁移方向
// @return error 错误信息
func selectExec(m *migrate.Migrate, fileNames []string, currentVersion uint, direction string) error {
	switch direction {
	case "up": // 执行迁移
		fmt.Println(commandConstant.DataBaseMigrateUp)
		err, isUp := up(m, fileNames, currentVersion)
		if err != nil {
			return err
		}
		if err == nil && !isUp {
			fmt.Println(color.GreenString(commandConstant.DataBaseMigrateUpSuccess))
		}
	case "down": // 执行回滚
		fmt.Println(commandConstant.DataBaseMigrateDown)
		err, isDown := down(m, currentVersion)
		if err != nil {
			return err
		}
		if err == nil && !isDown {
			fmt.Println(color.GreenString(commandConstant.DataBaseMigrateDownSuccess))
		}
	case "force": // 如果出现Error: Dirty database version XX. Fix and force version. 可以执行herman -d force -v (XX-1)即可回退版本
		fmt.Println(commandConstant.ForceChange)
		if err := m.Force(int(version)); err != nil {
			return err
		}
		fmt.Println(color.GreenString(commandConstant.ForceChangeSuccess))
	case "drop":
		fmt.Println(commandConstant.DropDatabase)
		if err := m.Drop(); err != nil {
			return err
		}
		fmt.Println(color.GreenString(commandConstant.DropDatabaseSuccess))
	default:
		return errors.New("invalid migration direction")
	}
	return nil
}

// up 数据库迁移
// @param *migrate.Migrate migrate 迁移实例
// @param []string fileNames 文件名切片
// @param uint currentVersion 当前数据库最新版本
// @return error 错误信息，是否为数据库已迁移到最新版本
func up(migrate *migrate.Migrate, fileNames []string, currentVersion uint) (error, bool) {
	version := getNewVersion(fileNames[len(fileNames)-1])
	if currentVersion == version {
		fmt.Println(color.GreenString(commandConstant.DataBaseNews))
		return nil, true
	}
	if number > 0 {
		if err := migrate.Steps(int(number)); err != nil {
			return err, false
		}
	}
	if number == 0 {
		if err := migrate.Up(); err != nil {
			return err, false
		}
	}
	return nil, false
}

// down 数据库回滚
// @param *migrate.Migrate migrate 迁移实例
// @param uint currentVersion 当前数据库最新版本
// @return error bool 错误信息，是否数据库当前已全部回滚
func down(migrate *migrate.Migrate, currentVersion uint) (error, bool) {
	if currentVersion == 0 {
		fmt.Println(color.GreenString(commandConstant.DataBaseEmpty))
		return nil, true
	}
	if number > 0 {
		if err := migrate.Steps(int(-number)); err != nil {
			return err, false
		}
	}
	if number == 0 {
		if err := migrate.Down(); err != nil {
			return err, false
		}
	}
	return nil, false
}

// getNewVersion 获取最新版本号
// @param string fileName 文件名
// @param uint 返回版本号
func getNewVersion(fileName string) uint {
	// 根据文件名提取版本号
	versionSlice := strings.Split(fileName, "_")
	newVersion, _ := strconv.ParseUint(versionSlice[0], 10, 32)
	return uint(newVersion)
}

// getFileNames 获取指定目录下所有文件名称并升序排序
// @param string dirPath 目录
// @param string direction 只包含direction文件名称进行排序
// @return err fileNames 错误信息，文件名字符串切片
func getFileNames(dirPath string, direction string) (err error, upFileNames []string) {
	var allFileNames []string
	// 读取指定目录
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err, nil
	}
	for _, file := range files {
		// 只获取文件名，忽略目录名
		if !file.IsDir() {
			if file.Size() == 0 {
				return errors.New(fmt.Sprintf("%s%s", file.Name(), commandConstant.ContentIsEmpty)), nil
			}
			allFileNames = append(allFileNames, file.Name())
			if strings.Contains(file.Name(), direction) {
				upFileNames = append(upFileNames, file.Name())
			}
		}
	}
	// 对文件名进行升序排序
	sort.Strings(upFileNames)

	if err := isFileExist(upFileNames, allFileNames); err != nil {
		return err, nil
	}

	return nil, upFileNames
}

// isFileExist 判断回滚文件是否存在
// @param []string upFileNames 迁移文件名切片
// @param []string allFileNames 全部文件名切片
// @return error 错误信息
func isFileExist(upFileNames []string, allFileNames []string) error {
	// 判断回滚文件是否存在，并且内容是否为空
	for _, upFileName := range upFileNames {
		split := strings.Split(upFileName, ".")
		downFileName := split[0] + ".down.sql"
		if !strings.Contains(strings.Join(allFileNames, ","), downFileName) {
			return errors.New(fmt.Sprintf(commandConstant.DownFileNotExist, upFileName, downFileName))
		}
	}
	return nil
}
