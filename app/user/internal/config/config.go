package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		BsMaster struct {
			Url     string
			Maxopen int
			Maxidle int
		}
	}
}
