package core

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// 初始化项目目录
var (
	RootPath     string
	AppPath      string
	ConfigPath   string
	StoragePath  string
	RoutePath    string
	ResourcePath string
)

// init 初始化项目目录
func init() {
	// 获取当前文件所在目录
	dir := absPathing("./")
	// 向上遍历文件目录，直到找到项目根目录
	for {
		if _, err := os.Stat(filepath.Join(dir, "main.go")); err == nil || !os.IsNotExist(err) {
			RootPath = dir
			AppPath = filepath.Join(dir, "app")
			ConfigPath = filepath.Join(dir, "config")
			StoragePath = filepath.Join(dir, "storages")
			RoutePath = filepath.Join(dir, "routes")
			ResourcePath = filepath.Join(dir, "resources")
			break
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			fmt.Println("Can not find project root path")
			os.Exit(1)
		}
		dir = parentDir
	}
}

// absPathing 获取绝对路径
// @param inPath string 当前路径
// @return string 绝对路径
func absPathing(inPath string) string {
	if inPath == "$HOME" || strings.HasPrefix(inPath, "$HOME"+string(os.PathSeparator)) {
		inPath = userHomeDir() + inPath[5:]
	}

	inPath = os.ExpandEnv(inPath)

	if filepath.IsAbs(inPath) {
		return filepath.Clean(inPath)
	}

	p, err := filepath.Abs(inPath)
	if err == nil {
		return filepath.Clean(p)
	}
	return ""
}

// userHomeDir 返回当前用户的主目录
// @return string 主目录
func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
