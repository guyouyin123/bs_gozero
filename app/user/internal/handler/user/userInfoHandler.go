package user

import (
	"bs_gozero/pkg/qhttpx"
	"net/http"

	"bs_gozero/app/user/internal/logic/user"
	"bs_gozero/app/user/internal/svc"
	"bs_gozero/app/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoRes
		if err := httpx.Parse(r, &req); err != nil {
			qhttpx.SendParamError(w, err)
			return
		}

		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, errMsg := l.UserInfo(&req)
		if errMsg != nil {
			qhttpx.SendCode(w, errMsg.Code, errMsg.Msg, "")
		} else {
			qhttpx.SendData(w, resp)
		}
	}
}
