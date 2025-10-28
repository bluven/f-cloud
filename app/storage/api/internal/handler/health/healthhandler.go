package health

import (
	"net/http"

	"github.com/bluven/f-cloud/app/storage/api/internal/logic/health"
	"github.com/bluven/f-cloud/app/storage/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HealthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := health.NewHealthLogic(r.Context(), svcCtx)
		err := l.Health()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
