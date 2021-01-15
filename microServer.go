package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	rateLimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"microServerTemplate/tools"
)

var (
	Cfg *tools.Config
	Qps = 100
)

func init() {
	fmt.Println("init project config ing...")
	Cfg = tools.ParseConfig("./domain/conf/config.yaml")
	tools.InitOrmEngine(Cfg.Mysql)
	tools.InitRedisEngine(Cfg.Redis)
	fmt.Println("init project end...")
}
func main() {
	etcdUrl := Cfg.Etcd.Host+":"+Cfg.Etcd.Port
	jaegerUrl := Cfg.Jaeger.Host +":"+Cfg.Jaeger.Port
	addr := Cfg.App.Host + ":"+Cfg.App.Port
	//链路追踪
	t,i,e := tools.NewTracer(Cfg.Jaeger.MicroName,jaegerUrl)
	if e !=nil{
		logger.Fatal("jaeger error",e)
	}
	defer i.Close()
	opentracing.SetGlobalTracer(t)
	service := micro.NewService(
		micro.Address(addr),
		micro.Version(Cfg.App.Version),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(rateLimit.NewHandlerWrapper(Qps)),
		micro.Registry(etcd.NewRegistry(registry.Addrs(etcdUrl))),
	)
	service.Init()
	// Register Handler


	// run server
	if err := service.Run(); err != nil {
		logger.Fatal("Run err,",err)
	}
}
