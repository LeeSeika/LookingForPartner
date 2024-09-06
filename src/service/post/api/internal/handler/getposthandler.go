package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lookingforpartner/service/post/api/internal/logic"
	"lookingforpartner/service/post/api/internal/svc"
	"lookingforpartner/service/post/api/internal/types"
)

func GetPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPostRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetPostLogic(r.Context(), svcCtx)
		resp, err := l.GetPost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
