package svc

import (
	"bs_gozero/app/bike/internal/config"
	"bs_gozero/app/bike/internal/db/bike"
)

type ServiceContext struct {
	Config config.Config
	Option Option
}

type Option struct {
	BikeTmpl bike.Object
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Option: Option{BikeTmpl: bike.New(bike.Option{Db: bsDb.DB(), DbGorm: bsDb})},
	}
}
