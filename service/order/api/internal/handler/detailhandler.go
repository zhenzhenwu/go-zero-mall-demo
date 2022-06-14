package handler

import (
	"net/http"

	"ang-miracle.com/go-zero-mall-demo/service/order/api/internal/logic"
	"ang-miracle.com/go-zero-mall-demo/service/order/api/internal/svc"
	"ang-miracle.com/go-zero-mall-demo/service/order/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
