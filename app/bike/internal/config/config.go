package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		BsMaster struct {
			Url     string
			Maxopen int
			Maxidle int
		}
	}
}
