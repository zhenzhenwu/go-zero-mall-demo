package handler

import (
	"net/http"

	"ang-miracle.com/go-zero-mall-demo/service/product/api/internal/logic"
	"ang-miracle.com/go-zero-mall-demo/service/product/api/internal/svc"
	"ang-miracle.com/go-zero-mall-demo/service/product/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateLogic(r.Context(), svcCtx)
		resp, err := l.Update(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
