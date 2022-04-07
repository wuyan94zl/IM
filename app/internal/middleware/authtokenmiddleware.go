package middleware

import (
	"github.com/wuyan94zl/go-zero-blog/app/common/auth"
	"github.com/wuyan94zl/go-zero-blog/app/internal/config"
	"net/http"
)

type AuthTokenMiddleware struct {
	Config config.Config
	User   *auth.Info
}

func NewAuthTokenMiddleware(c config.Config, user *auth.Info) *AuthTokenMiddleware {
	return &AuthTokenMiddleware{
		Config: c,
		User:   user,
	}
}

func (m *AuthTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("用户信息：", l.ctx.Value("id"),l.ctx.Value("nick_name"))
		next(w, r)
	}
}
