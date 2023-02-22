// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	ping "github.com/v3nooonn/trytry/apps/bff/api/internal/handler/ping"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.allowed},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: ping.PingHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/ping"),
	)
}
