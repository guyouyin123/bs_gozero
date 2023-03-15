package user

import (
	"bs_gozero/pkg/mysql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

func instance() *Option {
	url := "root:123456@tcp(127.0.0.1:3307)/bs?charset=utf8"
	db, err := mysql.InitMysqlDb(url, 10, 5)
	if err != nil {
		fmt.Printf("InitGorm error=%v", err)
		return nil
	}
	return New(Option{
		Db:     db.DB(),
		DbGorm: db,
	})
}

func TestQueryByBid(t *testing.T) {
	r := instance()
	userInfo := r.QueryByBid2(2)
	fmt.Println(userInfo)
}

func TestInsert(t *testing.T) {
	r := instance()
	i := Info{
		Name: "aa",
	}
	id := r.Insert(&i)
	fmt.Println(id)
}
