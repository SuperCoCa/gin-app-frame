package global

import (
	"github.com/gin-app-frame/internal/pkg/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppGlobal struct {
	Logger *zap.SugaredLogger
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
}

var App = new(AppGlobal)
