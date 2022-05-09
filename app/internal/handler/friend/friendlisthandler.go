package friend

import (
	"github.com/wuyan94zl/IM/app/common/response"
	"net/http"

	"github.com/wuyan94zl/IM/app/internal/logic/friend"
	"github.com/wuyan94zl/IM/app/internal/svc"
)

func FriendListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := friend.NewFriendListLogic(r.Context(), svcCtx)
		resp, err := l.FriendList()
		response.Response(w, r, resp, err)
	}
}
