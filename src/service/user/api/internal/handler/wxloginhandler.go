package handler

import (
	"lookingforpartner/service/user/api/internal/logic"
	"lookingforpartner/service/user/api/internal/svc"
	"lookingforpartner/service/user/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func WxLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WxLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewWxLoginLogic(r.Context(), svcCtx)
		resp, err := l.WxLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
