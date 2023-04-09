package main

import (
	"fmt"
	"go.uber.org/zap"
	"microHome/web/dao/mysql"
	"microHome/web/dao/redis"
	"microHome/web/logger"
	"microHome/web/router"
	"microHome/web/setting"
)

func main() {
	//1. 加载配置
	if err := setting.Init(); err != nil {
		fmt.Println("加载配置失败：", err)
		return
	}
	//2.初始化日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf); err != nil {
		fmt.Println("初始化日志失败：", err)
		return
	}
	//3.初始化redis连接池
	redis.Init(setting.Conf.RedisConfig)
	defer redis.Close()
	//4.初始化mysql数据库
	if err := mysql.Init(setting.Conf.MysqlConfig); err != nil {
		zap.L().Error("初始化mysql数据库失败", zap.Error(err))
		return
	}
	defer mysql.Close()
	//5.注册路由
	r := router.SetUp(setting.Conf.AppMode)
	//6.启动服务
	r.Run(setting.Conf.AppPort)
}
