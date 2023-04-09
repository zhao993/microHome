package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"house/dao/mysql"
	"house/dao/redis"
	"house/handler"
	pb "house/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "house"
	version = "latest"
)

func main() {
	//连接redis
	redis.Init()
	defer redis.Close()
	//初始化mysql
	err := mysql.Init()
	if err != nil {
		fmt.Println("mysql init error: ", err)
		return
	}
	defer mysql.Close()
	//初始化consul
	consulReg := consul.NewRegistry()
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Address("127.0.0.1:9094"), //防止随机生成port
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg), //注册到consul中
	)

	// Register handler
	if err := pb.RegisterHouseHandler(srv.Server(), new(handler.House)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
