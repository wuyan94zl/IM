package users

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/wuyan94zl/go-zero-blog/app/models/user"
	"github.com/wuyan94zl/go-zero-blog/app/utils"
	"time"

	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"

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

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (*utils.SuccessTmp, *utils.ErrorTmp) {
	info, err := l.svcCtx.UserModel.FindRaw(l.ctx, "user_name", req.UserName)
	if err != nil {
		if err == user.ErrNotFound {
			return nil, utils.Error(401, "用户名不存在")
		}
		return nil, utils.Error(401, err.Error())
	}
	if info.Password != utils.Md5ByString(req.Password) {
		return nil, utils.Error(401, "用户名密码错误")
	}

	// 生成token
	now, accessExpire := time.Now().Unix(), l.svcCtx.Config.JwtAuth.AccessExpire
	data := make(map[string]interface{})
	data["id"] = info.Id
	data["nick_name"] = info.NickName
	token, err := genToken(now, l.svcCtx.Config.JwtAuth.AccessSecret, data, accessExpire)

	return utils.Success(types.JwtTokenResponse{
		AccessToken:  token,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}), nil
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
