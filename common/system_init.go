package common

import (
	"context"
	"qqchat/common/qqlog"
)

// 全局上下文 生命周期控制
var Ctx = context.Background()

func init() {
	// 初始化配置
	InitConfig()
	// 初始化日志Logrus日志库
	qqlog.InitLogger()
	// 初始化gorm-mysql数据库
	InitGorm()
	//初始化redis
	InitRedis(Ctx)
}
