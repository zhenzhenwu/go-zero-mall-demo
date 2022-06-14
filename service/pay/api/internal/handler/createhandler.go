package handler

import (
	"net/http"

	"ang-miracle.com/go-zero-mall-demo/service/pay/api/internal/logic"
	"ang-miracle.com/go-zero-mall-demo/service/pay/api/internal/svc"
	"ang-miracle.com/go-zero-mall-demo/service/pay/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
