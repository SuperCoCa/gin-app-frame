package store

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-app-frame/internal/pkg/config"
	"github.com/gin-app-frame/internal/pkg/global"
	"github.com/gin-app-frame/internal/pkg/global/variable"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var cfg *config.Config
var zLogger *zap.SugaredLogger
var db *gorm.DB

// @description 初始化数据库
func InitDatabase() *gorm.DB {
	cfg = global.App.Config
	zLogger = global.App.Logger

	// 根据数据库引擎类型初始化
	switch cfg.Database.DbDriver {
	case "mysql":
		return initMysqlGorm()
	default:
		return initMysqlGorm()
	}
}

// @description 关闭数据库连接
func CloseDatabase() {
	if db != nil {
		sqlDb, _ := db.DB()
		sqlDb.Close()
	}
}

// @description 初始化mysql
func initMysqlGorm() *gorm.DB {
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.DbName,
		cfg.Database.Charset,
		strconv.FormatBool(cfg.Database.ParseTime),
	)

	if db, err = gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{Logger: getGormLogger()}); err != nil {
		zLogger.Panic("mysql connect failed: %s", err)
		return nil
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(cfg.Database.MaxIdleConnections)
	sqlDb.SetMaxOpenConns(cfg.Database.MaxOpenConnections)

	return db
}

// @description 自定义gorm的writer
func getGormLogWriter() logger.Writer {
	hook, err := rotatelogs.New(
		variable.LOG_PATH+"/"+cfg.Database.LogFileName+"-%Y%m%d"+cfg.Database.LogFileExt,
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(cfg.Database.LogMaxAge)), // 文件清理前的最长保存时间
		rotatelogs.WithRotationTime(time.Hour*24),                                 // 日志分割的时间（每天）
	)
	if err != nil {
		panic(err)
	}

	return log.New(hook, "["+time.Now().Format("2006-01-02 15:04:05")+"] ", log.Lmsgprefix)
}

// @description 获取logger
func getGormLogger() logger.Interface {
	var logMode logger.LogLevel
	switch cfg.Database.LogLevel {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	writer := getGormLogWriter()
	config := logger.Config{
		LogLevel:                  logMode,                // 日志记录等级
		SlowThreshold:             200 * time.Millisecond, // 慢SQL时间定义
		IgnoreRecordNotFoundError: false,                  // 是否忽略记录未找到的错误
		Colorful:                  false,                  // 关闭彩色打印
	}

	return logger.New(writer, config)
}
