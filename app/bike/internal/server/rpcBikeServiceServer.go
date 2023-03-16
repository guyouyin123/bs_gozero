// Code generated by goctl. DO NOT EDIT!
// Source: bike.proto

package server

import (
	"context"

	"bs_gozero/app/bike/internal/logic"
	"bs_gozero/app/bike/internal/svc"
	"bs_gozero/app/bike/pb"
)

type RpcBikeServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedRpcBikeServiceServer
}

func NewRpcBikeServiceServer(svcCtx *svc.ServiceContext) *RpcBikeServiceServer {
	return &RpcBikeServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *RpcBikeServiceServer) RpcGetBikeInfo(ctx context.Context, in *pb.GetBikeInfoRes) (*pb.GetBikeInfoReq, error) {
	l := logic.NewRpcGetBikeInfoLogic(ctx, s.svcCtx)
	return l.RpcGetBikeInfo(in)
}
