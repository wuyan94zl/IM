package friend

import (
	"net/http"

	"github.com/wuyan94zl/go-zero-blog/app/internal/logic/friend"
	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FriendHandleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendHandleRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := friend.NewFriendHandleLogic(r.Context(), svcCtx)
		resp, err := l.FriendHandle(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
