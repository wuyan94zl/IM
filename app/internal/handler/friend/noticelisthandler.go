package friend

import (
	"github.com/wuyan94zl/IM/app/common/response"
	"net/http"

	"github.com/wuyan94zl/IM/app/internal/logic/friend"
	"github.com/wuyan94zl/IM/app/internal/svc"
)

func NoticeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := friend.NewNoticeListLogic(r.Context(), svcCtx)
		resp, err := l.NoticeList()
		response.Response(w, r, resp, err)
	}
}
