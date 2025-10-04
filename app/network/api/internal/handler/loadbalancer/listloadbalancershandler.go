package loadbalancer

import (
	"net/http"

	"github.com/bluven/f-cloud/app/network/api/internal/logic/loadbalancer"
	"github.com/bluven/f-cloud/app/network/api/internal/svc"
	"github.com/bluven/f-cloud/app/network/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListLoadBalancersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListLoadBalancerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := loadbalancer.NewListLoadBalancersLogic(r.Context(), svcCtx)
		resp, err := l.ListLoadBalancers(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
