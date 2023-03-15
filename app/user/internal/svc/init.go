package svc

import (
	"bs_gozero/app/user/internal/config"
	"bs_gozero/pkg/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	bsDb *gorm.DB
)

func InitMysql(c config.Config) error {
	url := c.Mysql.BsMaster.Url
	maxopen := c.Mysql.BsMaster.Maxopen
	maxidle := c.Mysql.BsMaster.Maxidle
	db, err := mysql.InitMysqlDb(url, maxopen, maxidle)
	if err != nil {
		logx.Errorf("Fatal initMysql:%s", err.Error())
		return err
	}
	bsDb = db
	return nil
}
