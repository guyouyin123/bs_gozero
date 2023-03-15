package user

import (
	"bs_gozero/app/user/internal/db/user"
	"bs_gozero/app/user/internal/svc"
	"bs_gozero/app/user/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertInfoLogic {
	return &InsertInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertInfoLogic) InsertInfo(req *types.UserInsertInfoRes) error {
	// todo: add your logic here and delete this line

	userInfo := user.Info{
		Name: req.Name,
	}
	id := l.svcCtx.Option.UserTmpl.Insert(&userInfo)
	if id <= 0 {
		logx.Infof("UserInsertInfo:%s", userInfo)
		return nil
	}
	return nil
}
