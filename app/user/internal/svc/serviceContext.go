package svc

import (
	"bs/app/user/internal/config"
	"bs/app/user/internal/db/user"
	"bs/app/user/internal/middleware"
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
