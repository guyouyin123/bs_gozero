package svc

import (
	"bs_gozero/app/user/internal/config"
	"bs_gozero/app/user/internal/db/user"
	"bs_gozero/app/user/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config        config.Config
	Option        Option
	LetMiddleware rest.Middleware //局部中间件
}

type Option struct {
	UserTmpl user.Object
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		LetMiddleware: middleware.NewLetMiddleware().Handle, //局部中间件初始化
		Option:        Option{UserTmpl: user.New(user.Option{Db: bsDb.DB(), DbGorm: bsDb})},
	}
}
