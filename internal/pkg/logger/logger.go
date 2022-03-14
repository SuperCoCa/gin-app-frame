package logger

import (
	"time"

	"github.com/gin-app-frame/internal/pkg/config"
	"github.com/gin-app-frame/internal/pkg/global"
	"github.com/gin-app-frame/internal/pkg/global/variable"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var cfg *config.Config
var options []zap.Option

func InitLogger() *zap.SugaredLogger {
	cfg = global.App.Config

	// 展示调用行
	if cfg.Log.ShowLine {
		options = append(options, zap.AddCaller())
	}

	core := getCore()
	logger := zap.New(core, options...)

	return logger.Sugar()
}

func getCore() zapcore.Core {
	encoder := getEncoder()
	writter := getLogWriter()
	level := getLevel()

	return zapcore.NewCore(encoder, writter, level)
}

// @description 获取编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("[" + "2006-01-02 15:04:05" + "]")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

// @description 获取日志写入位置
func getLogWriter() zapcore.WriteSyncer {
	hook, err := rotatelogs.New(
		variable.LOG_PATH+"/"+cfg.Log.FileName+"-%Y%m%d"+cfg.Log.FileExt,
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(cfg.Log.MaxAge)), // 文件清理前的最长保存时间
		rotatelogs.WithRotationTime(time.Hour*24),                         // 日志分割的时间（每天）
	)
	if err != nil {
		panic(err)
	}

	return zapcore.AddSync(hook)
}

// @description 获取日志写入级别
func getLevel() zapcore.Level {
	var level zapcore.Level
	switch cfg.Log.Level {
	case "debug":
		level = zapcore.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	return level
}
