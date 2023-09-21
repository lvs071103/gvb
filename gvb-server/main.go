package main

import (
	"fmt"
	"gvb-server/core"
	"gvb-server/flag"
	"gvb-server/global"
	"gvb-server/routers"
)

func main() {
	// 读取配置文件
	core.InitConf()
	fmt.Println(global.Config)
	// 初始化日志
	global.Log = core.InitLogger()
	// 初始化数据库连接
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
	// 命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		return
	} else {
		flag.SwitchOption(option)
	}

	// 路由初始化
	router := routers.InitRouter()
	addr := global.Config.System.ServerAddr()
	global.Log.Infof("gvb-server正在运行在：%s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
