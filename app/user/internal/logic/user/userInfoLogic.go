package user

import (
	"bs_gozero/app/user/internal/svc"
	"bs_gozero/app/user/internal/types"
	"bs_gozero/pkg/qhttpx"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRes) (resp *types.UserInfoReq, errMsg *qhttpx.CommonErrorMsg) {
	// todo: add your logic here and delete this line
	userInfo := l.svcCtx.Option.UserTmpl.QueryByBid(req.UserId)
	if userInfo == nil {
		err := qhttpx.CommonErrorMsg{Code: "100102", Msg: "数据库没有查到哦"}
		l.Logger.Info(fmt.Sprintf("errCode=%s,errMsg=%s", err.Code, err.Msg))
		//logx.Errorf("errCode=%s,errMsg=%s", err.Code, err.Msg)
		return nil, &err
	}
	return &types.UserInfoReq{
		UserId: userInfo.Id,
		Name:   userInfo.Name,
	}, nil
}
