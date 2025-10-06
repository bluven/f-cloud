package backup

import (
	"net/http"

	"github.com/bluven/f-cloud/app/storage/api/internal/logic/backup"
	"github.com/bluven/f-cloud/app/storage/api/internal/svc"
	"github.com/bluven/f-cloud/app/storage/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateDiskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateBackupRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := backup.NewCreateDiskLogic(r.Context(), svcCtx)
		resp, err := l.CreateDisk(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
