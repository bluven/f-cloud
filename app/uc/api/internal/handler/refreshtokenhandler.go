package handler

import (
	"net/http"

	"github.com/bluven/f-cloud/app/uc/api/internal/logic"
	"github.com/bluven/f-cloud/app/uc/api/internal/svc"
	"github.com/bluven/f-cloud/app/uc/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RefreshTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmptyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRefreshTokenLogic(r.Context(), svcCtx)
		resp, err := l.RefreshToken(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
