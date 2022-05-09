package users

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/wuyan94zl/IM/app/common/response"
	utils2 "github.com/wuyan94zl/IM/app/common/utils"
	"github.com/wuyan94zl/IM/app/models/user"
	"time"

	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (*types.JwtTokenResponse, error) {
	info, err := l.svcCtx.UserModel.FindRawByName(l.ctx, req.UserName)
	if err != nil {
		if err == user.ErrNotFound {
			return nil, response.Error(401, "用户名不存在")
		}
		return nil, response.Error(401, err.Error())
	}
	if info.Password != utils2.Md5ByString(req.Password) {
		return nil, response.Error(401, "用户名密码错误")
	}

	// 生成token
	now, accessExpire := time.Now().Unix(), l.svcCtx.Config.JwtAuth.AccessExpire
	data := make(map[string]interface{})
	data["id"] = info.Id
	data["nick_name"] = info.NickName
	token, err := genToken(now, l.svcCtx.Config.JwtAuth.AccessSecret, data, accessExpire)
	return &types.JwtTokenResponse{
		AccessToken:  token,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
		WsToken:      info.Id,
	}, nil
}

func genToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
