package friend

import (
	"net/http"

	"github.com/wuyan94zl/go-zero-blog/app/internal/logic/friend"
	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FriendAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := friend.NewFriendAddLogic(r.Context(), svcCtx)
		resp, err := l.FriendAdd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
