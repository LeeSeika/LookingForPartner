package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lookingforpartner/service/comment/api/internal/logic"
	"lookingforpartner/service/comment/api/internal/svc"
	"lookingforpartner/service/comment/api/internal/types"
)

func GetCommentsByPostIDHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCommentsByPostIDRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetCommentsByPostIDLogic(r.Context(), svcCtx)
		resp, err := l.GetCommentsByPostID(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}