package friend

import (
	"net/http"

	"github.com/wuyan94zl/go-zero-blog/app/internal/logic/friend"
	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func NoticeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := friend.NewNoticeListLogic(r.Context(), svcCtx)
		resp, err := l.NoticeList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
