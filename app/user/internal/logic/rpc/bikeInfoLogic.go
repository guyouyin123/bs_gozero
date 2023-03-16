package rpc

import (
	"bs_gozero/app/bike/pb"
	"context"

	"bs_gozero/app/user/internal/svc"
	"bs_gozero/app/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BikeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBikeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BikeInfoLogic {
	return &BikeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BikeInfoLogic) BikeInfo(req *types.BikeInfoRes) (resp *types.BikeInfoReq, err error) {
	// todo: add your logic here and delete this line
	bikeInfo, err := l.svcCtx.RpcBikeServiceClient.RpcGetBikeInfo(l.ctx, &pb.GetBikeInfoRes{Id: req.BikeId})
	if err != nil {
		return nil, err
	}

	return &types.BikeInfoReq{
		BikeId:   bikeInfo.Id,
		BikeName: bikeInfo.BikeName,
	}, nil
}
