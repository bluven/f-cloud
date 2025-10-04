package middleware

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/bluven/f-cloud/pkg/auth"
	"github.com/bluven/f-cloud/pkg/errorx"
)

type IsAllowedFunc func(ctx context.Context) bool

func newIsAllowedMiddleware(isAllowed IsAllowedFunc) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if isAllowed(ctx) {
				next(w, r)
			} else {
				httpx.Error(w, errorx.ErrForbidden)
			}
		}
	}
}

var (
	AdminOrCurrentUserMiddleware = newIsAllowedMiddleware(func(ctx context.Context) bool {
		return auth.IsAdminOrCurrentUser(ctx, auth.GetUserID(ctx))
	})

	CurrentUserRequiredMiddleware = newIsAllowedMiddleware(func(ctx context.Context) bool {
		return auth.IsCurrentUser(ctx, auth.GetUserID(ctx))
	})

	AdminRequiredMiddleware = newIsAllowedMiddleware(func(ctx context.Context) bool {
		return auth.IsAdmin(ctx)
	})
)
