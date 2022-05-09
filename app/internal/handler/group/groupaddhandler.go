package group

import (
	"github.com/wuyan94zl/IM/app/common/response"
	"net/http"

	"github.com/wuyan94zl/IM/app/internal/logic/group"
	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GroupAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupAddRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := group.NewGroupAddLogic(r.Context(), svcCtx)
		resp, err := l.GroupAdd(&req)
		response.Response(w, r, resp, err)
	}
}
