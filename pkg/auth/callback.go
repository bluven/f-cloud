package auth

import (
	"net/http"

	"github.com/bluven/f-cloud/pkg/errorx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UnauthorizedCallback(w http.ResponseWriter, r *http.Request, _ error) {
	httpx.WriteJson(w, http.StatusUnauthorized, errorx.ErrNotAuthorized)
}
