package group

import (
	"net/http"

	"github.com/wuyan94zl/IM/app/common/response"
	"github.com/wuyan94zl/IM/app/internal/logic/group"
	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GroupListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := group.NewGroupListLogic(r.Context(), svcCtx)
		resp, err := l.GroupList(&req)
		response.Response(w, r, resp, err)
	}
}
