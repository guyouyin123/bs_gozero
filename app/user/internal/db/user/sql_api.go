package user

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

func (o Option) QueryByBid2(id int) (info *Info) {
	sqlStr := `select id,name from user where 1=1`
	qstr := fmt.Sprintf(" and id = %d", id)
	sqlStr += qstr
	r := new(Info)
	err := o.Db.QueryRow(sqlStr).Scan(&r.Id, &r.Name)
	if err != nil {
		logx.Errorf("QueryByBidErr:%s", err.Error())
		return nil
	}
	return r
}
