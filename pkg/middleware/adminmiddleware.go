package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/bluven/f-cloud/pkg/auth"
	"github.com/bluven/f-cloud/pkg/errorx"
)

type AdminRequiredMiddleware struct {
}

func NewAdminRequiredMiddleware() AdminRequiredMiddleware {
	return AdminRequiredMiddleware{}
}

func (m AdminRequiredMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if auth.IsAdmin(ctx) {
			next(w, r)
		} else {
			httpx.Error(w, errorx.ErrForbidden)
		}
	}
}
