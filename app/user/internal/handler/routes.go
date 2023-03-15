// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "bs/app/user/internal/handler/user"
	"bs/app/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.LetMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/getInfo",
					Handler: user.UserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/insertInfo",
					Handler: user.InsertInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/dev/v1"),
	)
}
