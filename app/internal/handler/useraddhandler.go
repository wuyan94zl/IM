package handler

import (
	"net/http"

	"github.com/wuyan94zl/go-zero-blog/app/internal/logic"
	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserAddRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserAddLogic(r.Context(), svcCtx)
		resp, err := l.UserAdd(&req)
		if err != nil {
			httpx.OkJson(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
