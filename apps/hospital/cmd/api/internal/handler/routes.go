// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	hsptGrp "github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/handler/hsptGrp"
	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authentication, serverCtx.Authorization},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/estb",
					Handler: hsptGrp.EstbHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/hospital"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/up",
				Handler: HealthCheckHandler(serverCtx),
			},
		},
		rest.WithPrefix("/health"),
	)
}
