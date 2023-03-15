package main

import (
	"bs/pkg/middleware"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"

	"bs/app/user/internal/config"
	"bs/app/user/internal/handler"
	"bs/app/user/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

func main() {
	path, _ := os.Getwd()
	path += "/app/user/etc/user.yaml"
	var configFile = flag.String("f", path, "the config file")
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	if err := svc.InitMysql(c); err != nil {
		logx.Errorf("Fatal initMysql:%s", err.Error())
		return
	} //初始化Mysql

	logx.DisableStat()                                  //禁用每分钟打印的监控日志，建议打开
	server.Use(middleware.NewGlobalMiddleware().Handle) //全局中间件
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
