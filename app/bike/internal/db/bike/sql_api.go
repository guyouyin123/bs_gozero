package bike

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

func (o Option) QueryByBid2(id int64) (info *Info) {
	sqlStr := `select id,name from bike where 1=1`
	qstr := fmt.Sprintf(" and id = %d", id)
	sqlStr += qstr
	r := new(Info)
	err := o.Db.QueryRow(sqlStr).Scan(&r.Id, &r.BikeName, &r.Created, &r.Updated)
	if err != nil {
		logx.Errorf("QueryByBidErr:%s", err.Error())
		return nil
	}
	return r
}
