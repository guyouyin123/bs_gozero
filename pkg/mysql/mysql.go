package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zeromicro/go-zero/core/logx"
)

func InitMysqlDb(url string, maxopen, maxidle int) (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", url)
	if err != nil {
		logx.Errorf("mysql open faild,err=%v", err.Error())
		return
	}
	sqlDb := db.DB()
	sqlDb.SetMaxOpenConns(maxopen / 2)
	sqlDb.SetMaxIdleConns(maxidle / 2)
	err = sqlDb.Ping()
	if err != nil {
		logx.Errorf("mysqlDb ping faild,err=%v", err.Error())
		return
	}
	db.LogMode(true)
	return
}
