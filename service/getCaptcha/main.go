package main

import (
	"getCaptcha/dao/redis"
	"getCaptcha/handler"
	pb "getCaptcha/proto"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "getCaptcha"
	version = "latest"
)

func main() {
	//初始化redis连接池
	redis.Init()
	defer redis.Close()
	//初始化consul
	consulReg := consul.NewRegistry()
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Address("127.0.0.1:9090"), //防止随机生成port
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg), //注册到consul中
	)

	// Register handler
	if err := pb.RegisterGetCaptchaHandler(srv.Server(), new(handler.GetCaptcha)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
