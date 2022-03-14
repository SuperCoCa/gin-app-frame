package app

import (
	"github.com/gin-app-frame/internal/app/router"
	"github.com/gin-app-frame/internal/pkg/config"
	e "github.com/gin-app-frame/internal/pkg/errors"
	"github.com/gin-app-frame/internal/pkg/global"
	"github.com/gin-app-frame/internal/pkg/logger"
	"github.com/gin-app-frame/internal/pkg/store"
	"github.com/gin-app-frame/internal/pkg/validator"
)

func Run() {
	// 初始化配置
	config.InitConfig()

	// 加载配置
	global.App.Config = config.GetConfig()
	// 加载zap日志
	global.App.Logger = logger.InitLogger()
	// 加载数据库
	global.App.DB = store.InitDatabase()
	defer store.CloseDatabase()
	// 加载redis
	global.App.Redis = store.InitRedis()

	// 注册错误信息
	e.RegisterErrorInfo()

	// 加载自定义验证规则
	validator.InitValidation()

	// 加载路由
	router := router.RegisterRouter()

	router.Run(":" + global.App.Config.Server.HttpPort)
}
