package logic

import (
	"context"

	"github.com/wuyan94zl/go-zero-blog/app/users/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/users/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersLogic {
	return &UsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UsersLogic) Users(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
