package users

import (
	"net/http"

	"github.com/wuyan94zl/IM/app/internal/logic/users"
	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := users.NewUserRegisterLogic(r.Context(), svcCtx)
		resp, err := l.UserRegister(&req)
		if err != nil {
			httpx.OkJson(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
