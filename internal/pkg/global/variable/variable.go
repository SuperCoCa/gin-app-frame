package variable

import (
	"os"
	"path/filepath"
)

var (
	ROOT_PATH   string // 根目录
	CONFIG_PATH string // 配置目录
	LOG_PATH    string // 日志目录
)

func init() {
	ROOT_PATH, _ := os.Getwd()
	CONFIG_PATH = filepath.Join(ROOT_PATH, "configs")
	LOG_PATH = filepath.Join(ROOT_PATH, "runtime/logs")
}
