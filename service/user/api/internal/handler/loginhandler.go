package handler

import (
	"net/http"

	"ang-miracle.com/go-zero-mall-demo/service/user/api/internal/logic"
	"ang-miracle.com/go-zero-mall-demo/service/user/api/internal/svc"
	"ang-miracle.com/go-zero-mall-demo/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
