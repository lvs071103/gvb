package routers

import (
	"gvb-server/global"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("", func(ctx *gin.Context) {
		ctx.String(200, "hello world")
	})
	return router
}