package handler

import (
	"net/http"

	"github.com/wuyan94zl/go-zero-blog/app/users/internal/logic"
	"github.com/wuyan94zl/go-zero-blog/app/users/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/users/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUsersLogic(r.Context(), svcCtx)
		resp, err := l.Users(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
