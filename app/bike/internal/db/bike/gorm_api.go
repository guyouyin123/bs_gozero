package bike

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type Object interface {
	QueryByBid(id int64) (info *Info)  //gorm
	QueryByBid2(id int64) (info *Info) //sql
	Insert(info *Info) (id int64)      //gorm
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

func (o Option) QueryByBid(id int64) *Info {
	res := new(Info)
	err := o.DbGorm.Model(Info{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		logx.Errorf("QueryByBidErr:%s", err.Error())
		return nil
	}
	return res
}

func (o Option) Insert(i *Info) (id int64) {
	i.Created = time.Now().Unix()
	i.Updated = time.Now().Unix()
	err := o.DbGorm.Model(Info{}).Create(&i).Error
	if err != nil {
		logx.Errorf("InserttErr:%s", err.Error())
		return 0
	}
	return i.Id
}
