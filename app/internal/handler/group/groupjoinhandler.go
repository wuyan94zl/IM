package group

import (
	"net/http"

	"github.com/wuyan94zl/IM/app/common/response"
	"github.com/wuyan94zl/IM/app/internal/logic/group"
	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GroupJoinHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupJoinRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := group.NewGroupJoinLogic(r.Context(), svcCtx)
		resp, err := l.GroupJoin(&req)
		response.Response(w, r, resp, err)
	}
}
