package backup

import (
	"net/http"

	"github.com/bluven/f-cloud/app/storage/api/internal/logic/backup"
	"github.com/bluven/f-cloud/app/storage/api/internal/svc"
	"github.com/bluven/f-cloud/app/storage/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteBackupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := backup.NewDeleteBackupLogic(r.Context(), svcCtx)
		resp, err := l.DeleteBackup(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
