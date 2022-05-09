package users

import (
	"context"
	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (*types.UserInfoResponse, error) {
	id := l.svcCtx.AuthUser.Id
	info, err := l.svcCtx.UserModel.FindOne(l.ctx, id)
	return &types.UserInfoResponse{
		Id:       info.Id,
		UserName: info.UserName,
		NickName: info.NickName,
		Mobile:   info.Mobile,
	}, err
}
