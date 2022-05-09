package friend

import (
	"github.com/wuyan94zl/IM/app/common/response"
	"net/http"

	"github.com/wuyan94zl/IM/app/internal/logic/friend"
	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MessageListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := friend.NewMessageListLogic(r.Context(), svcCtx)
		resp, err := l.MessageList(&req)
		response.Response(w, r, resp, err)
	}
}
