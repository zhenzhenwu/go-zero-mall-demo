package handler

import (
	"net/http"

	"ang-miracle.com/go-zero-mall-demo/service/pay/api/internal/logic"
	"ang-miracle.com/go-zero-mall-demo/service/pay/api/internal/svc"
	"ang-miracle.com/go-zero-mall-demo/service/pay/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CallbackRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCallbackLogic(r.Context(), svcCtx)
		resp, err := l.Callback(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
