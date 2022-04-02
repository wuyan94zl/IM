package users

import (
	"context"
	"github.com/wuyan94zl/go-zero-blog/app/models/user"
	"github.com/wuyan94zl/go-zero-blog/app/utils"
	"time"

	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.RegisterRequest) (*utils.SuccessTmp, *utils.ErrorTmp) {
	// 验证用户是否存在
	_, err := l.svcCtx.UserModel.FindRaw(l.ctx, "user_name", req.UserName)
	if err == nil {
		return nil, utils.Error(401, "用户已存在")
	}
	// 注册用户
	register := user.Users{
		UserName: req.UserName,
		NickName: req.NickName,
		Password: utils.Md5ByString(req.Password),
		Mobile:   req.Mobile,
	}
	insert, err := l.svcCtx.UserModel.Insert(l.ctx, &register)
	if err != nil {
		return nil, utils.Error(401, "用户注册失败")
	}
	id, _ := insert.LastInsertId()

	// 生成token
	now, accessExpire := time.Now().Unix(), l.svcCtx.Config.JwtAuth.AccessExpire
	info := make(map[string]interface{})
	info["id"] = id
	info["nick_name"] = req.NickName
	token, err := genToken(now, l.svcCtx.Config.JwtAuth.AccessSecret, info, accessExpire)
	return utils.Success(types.JwtTokenResponse{
		AccessToken:  token,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}), nil

}
