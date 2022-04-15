package blogs

import (
	"net/http"

	"github.com/wuyan94zl/go-zero-blog/app/internal/logic/blogs"
	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateBlogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateBlogRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := blogs.NewCreateBlogLogic(r.Context(), svcCtx)
		resp, err := l.CreateBlog(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
