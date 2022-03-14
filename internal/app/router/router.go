package router

import (
	"github.com/gin-app-frame/internal/app/router/hello"
	"github.com/gin-gonic/gin"
)

// 注册路由
func RegisterRouter() *gin.Engine {
	router := gin.Default()

	// 分组路由
	apiV1 := router.Group("/api/v1")
	// 载入分组路由
	hello.SetHelloGroupRoutes(apiV1)

	return router
}
