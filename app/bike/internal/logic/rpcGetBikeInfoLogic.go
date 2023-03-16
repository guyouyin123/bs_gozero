package logic

import (
	"bs_gozero/app/bike/internal/svc"
	"bs_gozero/app/bike/pb"
	"context"
	"errors"
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
	//if md, ok := metadata.FromIncomingContext(l.ctx); ok {
	//	tmp := md.Get("user") //类似于网管，接收业务传过来端值
	//	fmt.Println("rpc客户端传过来端值：", tmp)
	//}
	bikeInfo := l.svcCtx.Option.BikeTmpl.QueryByBid(in.Id)
	if bikeInfo == nil {
		return nil, errors.New("数据库没有查到哦")
	}
	return &pb.GetBikeInfoReq{
		Id:       bikeInfo.Id,
		BikeName: bikeInfo.BikeName,
	}, nil
}
