package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lookingforpartner/service/post/api/internal/logic"
	"lookingforpartner/service/post/api/internal/svc"
	"lookingforpartner/service/post/api/internal/types"
)

func GetPostsByAuthorIDHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPostByAuthorIDRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetPostsByAuthorIDLogic(r.Context(), svcCtx)
		resp, err := l.GetPostsByAuthorID(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}