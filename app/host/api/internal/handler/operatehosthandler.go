package handler

import (
	"net/http"

	"github.com/bluven/f-cloud/app/host/api/internal/logic"
	"github.com/bluven/f-cloud/app/host/api/internal/svc"
	"github.com/bluven/f-cloud/app/host/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OperateHostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OperateHostRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOperateHostLogic(r.Context(), svcCtx)
		resp, err := l.OperateHost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
