package main

import (
	"fmt"
	"gvb-server/core"
	"gvb-server/global"
	"gvb-server/routers"
)

func main()  {
	// 读取配置文件
	core.InitConf()
	fmt.Println(global.Config)
	// 初始化日志
	global.Log = core.InitLogger()
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
	// 路由初始化
	router := routers.InitRouter()
	addr := global.Config.System.ServerAddr()
	global.Log.Infof("gvb-server正在运行在：%s", addr)
	router.Run(addr)
}