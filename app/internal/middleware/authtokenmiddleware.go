package middleware

import (
	"encoding/json"
	"github.com/wuyan94zl/IM/app/common/auth"
	"github.com/wuyan94zl/IM/app/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
)

type AuthTokenMiddleware struct {
	Config   config.Config
	User     *auth.Info
	RedisCli *redis.Redis
}

func NewAuthTokenMiddleware(c config.Config, user *auth.Info, redisCli *redis.Redis) *AuthTokenMiddleware {
	return &AuthTokenMiddleware{
		Config:   c,
		User:     user,
		RedisCli: redisCli,
	}
}

func (m *AuthTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := r.Context().Value("id").(json.Number).Int64()
		nickName := r.Context().Value("nick_name").(string)
		m.User.Id = id
		m.User.NickName = nickName
		//fmt.Println("用户信息：", l.ctx.Value("id"),l.ctx.Value("nick_name"))
		redis.NewRedisLock(m.RedisCli,"test").Acquire()
		next(w, r)
	}
}
