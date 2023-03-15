package user

import (
	"bs_gozero/pkg/qhttpx"
	"net/http"

	"bs_gozero/app/user/internal/logic/user"
	"bs_gozero/app/user/internal/svc"
	"bs_gozero/app/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func InsertInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInsertInfoRes
		if err := httpx.Parse(r, &req); err != nil {
			qhttpx.SendParamError(w, err)
			return
		}

		l := user.NewInsertInfoLogic(r.Context(), svcCtx)
		err := l.InsertInfo(&req)
		if err != nil {
			qhttpx.SendFailed(w, err)
		} else {
			qhttpx.SendOk(w)
		}
	}
}
