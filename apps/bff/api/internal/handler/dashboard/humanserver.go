package dashboard

import (
	"net/http"

	e "github.com/levvel-health/rpms-service/pkg/error"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/logic/dashboard"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HumanServerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := dashboard.NewHumanServerLogic(r.Context(), svcCtx)
		resp, err := l.HumanServer()

		e.RespHandler(r, w, resp, err)
	}
}
