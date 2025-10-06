package instance

import (
	"net/http"

	"github.com/bluven/f-cloud/app/instance/api/internal/logic/instance"
	"github.com/bluven/f-cloud/app/instance/api/internal/svc"
	"github.com/bluven/f-cloud/app/instance/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpgradeInstanceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeInstanceRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := instance.NewUpgradeInstanceLogic(r.Context(), svcCtx)
		resp, err := l.UpgradeInstance(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
