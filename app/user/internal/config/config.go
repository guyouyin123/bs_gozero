package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
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
	RpcBikeServiceConf zrpc.RpcClientConf
}
