package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/wuyan94zl/go-zero-blog/app/common/auth"
	"github.com/wuyan94zl/go-zero-blog/app/internal/config"
	"net/http"
	"strings"
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
		tokenStr := r.Header.Get("Authorization")
		info, err := m.authToken(tokenStr)
		if err != nil {
			panic("认证失败")
		}
		m.User.Id = int64(info["id"].(float64))
		m.User.NickName = info["nick_name"].(string)
		next(w, r)
	}
}

func (m *AuthTokenMiddleware) authToken(tokenStr string) (jwt.MapClaims, error) {
	kv := strings.Split(tokenStr, " ")
	tokenStr = kv[1]
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("认证失败: %v", token.Header["alg"])
		}
		return []byte(m.Config.JwtAuth.AccessSecret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("认证失败")
	}
}
