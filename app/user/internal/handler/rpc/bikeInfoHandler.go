package rpc

import (
	"bs_gozero/pkg/qhttpx"
	"net/http"

	"bs_gozero/app/user/internal/logic/rpc"
	"bs_gozero/app/user/internal/svc"
	"bs_gozero/app/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BikeInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BikeInfoRes
		if err := httpx.Parse(r, &req); err != nil {
			qhttpx.SendParamError(w, err)
			return
		}

		l := rpc.NewBikeInfoLogic(r.Context(), svcCtx)
		resp, err := l.BikeInfo(&req)
		if err != nil {
			qhttpx.SendFailed(w, err)
		} else {
			qhttpx.SendData(w, resp)
		}
	}
}
