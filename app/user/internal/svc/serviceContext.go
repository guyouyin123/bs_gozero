package svc

import (
	"bs_gozero/app/bike/rpcbikeservice"
	"bs_gozero/app/user/internal/config"
	"bs_gozero/app/user/internal/db/user"
	"bs_gozero/app/user/internal/middleware"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ServiceContext struct {
	Config               config.Config
	Option               Option
	LetMiddleware        rest.Middleware               //局部中间件
	RpcBikeServiceClient rpcbikeservice.RpcBikeService //Bike服务Rpc注册
}

type Option struct {
	UserTmpl user.Object
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:               c,
		LetMiddleware:        middleware.NewLetMiddleware().Handle, //局部中间件初始化
		Option:               Option{UserTmpl: user.New(user.Option{Db: bsDb.DB(), DbGorm: bsDb})},
		RpcBikeServiceClient: rpcbikeservice.NewRpcBikeService(zrpc.MustNewClient(c.RpcBikeServiceConf, zrpc.WithUnaryClientInterceptor(SetHeader))), //rpc客户端拦截器
	}
}

func SetHeader(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Println("====rpc客户端拦截器之前====")
	fmt.Println("====req:", req)
	fmt.Println("====method:", method)

	md := metadata.New(map[string]string{"userName": "Jeff"}) //传值,类似于网管传业务需要端数据
	ctx = metadata.NewOutgoingContext(ctx, md)

	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}
	fmt.Println("====rpc客户端拦截器之后====")
	return nil
}
