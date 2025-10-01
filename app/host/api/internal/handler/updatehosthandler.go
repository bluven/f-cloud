package handler

import (
	"net/http"

	"github.com/bluven/f-cloud/app/host/api/internal/logic"
	"github.com/bluven/f-cloud/app/host/api/internal/svc"
	"github.com/bluven/f-cloud/app/host/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateHostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateHostRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateHostLogic(r.Context(), svcCtx)
		resp, err := l.UpdateHost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
