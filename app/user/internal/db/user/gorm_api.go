package user

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type Object interface {
	QueryByBid(id int) (info *Info)  //gorm
	QueryByBid2(id int) (info *Info) //sql
	Insert(info *Info) (id int)      //gorm
}

type Option struct {
	//Log    *qlog.Entry
	Db     *sql.DB
	DbGorm *gorm.DB
}

func New(in Option) *Option {
	return &Option{
		//Log:    in.Log,
		Db:     in.Db,
		DbGorm: in.DbGorm,
	}
}

func (o Option) QueryByBid(id int) *Info {
	res := new(Info)
	err := o.DbGorm.Model(Info{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		logx.Errorf("QueryByBidErr:%s", err.Error())
		return nil
	}
	return res
}

func (o Option) Insert(i *Info) (id int) {
	i.Created = time.Now().Unix()
	i.Updated = time.Now().Unix()
	//err := o.DbGorm.Table("user").Create(&i).Error
	err := o.DbGorm.Model(Info{}).Create(&i).Error
	if err != nil {
		logx.Errorf("UserInsertErr:%s", err.Error())
		return 0
	}
	return i.Id
}
