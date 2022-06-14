package handler

import (
	"net/http"

	"ang-miracle.com/go-zero-mall-demo/service/product/api/internal/logic"
	"ang-miracle.com/go-zero-mall-demo/service/product/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewTestLogic(r.Context(), svcCtx)
		resp, err := l.Test()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
