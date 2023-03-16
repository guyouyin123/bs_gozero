package logic

import (
	"context"

	"bs_gozero/app/bike/internal/svc"
	"bs_gozero/app/bike/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RpcGetBikeInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRpcGetBikeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RpcGetBikeInfoLogic {
	return &RpcGetBikeInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RpcGetBikeInfoLogic) RpcGetBikeInfo(in *pb.GetBikeInfoRes) (*pb.GetBikeInfoReq, error) {
	// todo: add your logic here and delete this line

	return &pb.GetBikeInfoReq{}, nil
}
