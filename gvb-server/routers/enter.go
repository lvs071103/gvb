package routers

import (
	"gvb-server/global"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct{
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	// 系统配置API
	routerGroupApp.SettingsRouter()
	return router
}