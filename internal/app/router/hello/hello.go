package hello

import (
	"net/http"

	v1 "github.com/gin-app-frame/internal/app/api/v1"
	"github.com/gin-gonic/gin"
)

func SetHelloGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	router.POST("/register", v1.Register)
}
