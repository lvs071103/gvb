package main

import (
	"fmt"
	"gvb-server/core"
	"gvb-server/global"
)

func main()  {
	// 读取配置文件
	core.InitConf()
	fmt.Println(global.Config)
	// 初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("嘻嘻嘻...")
	global.Log.Errorln("嘻嘻嘻...")
	global.Log.Infoln("嘻嘻嘻...")
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}