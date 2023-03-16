package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"

	"bs_gozero/app/bike/internal/config"
	"bs_gozero/app/bike/internal/server"
	"bs_gozero/app/bike/internal/svc"
	"bs_gozero/app/bike/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	path, _ := os.Getwd()
	path += "/app/bike/etc/bike.yaml"
	var configFile = flag.String("f", path, "the config file")
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	
	if err := svc.InitMysql(c); err != nil {
		logx.Errorf("Fatal initMysql:%s", err.Error())
		return
	} //初始化Mysql

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterRpcBikeServiceServer(grpcServer, server.NewRpcBikeServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
