package users

import (
	"context"
	"fmt"
	"github.com/wuyan94zl/go-zero-blog/app/common/utils"
	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"

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

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (*utils.SuccessTmp, *utils.ErrorTmp) {
	fmt.Println("用户信息：", l.svcCtx.AuthUser)
	info, _ := l.svcCtx.UserModel.FindOne(l.ctx, l.svcCtx.AuthUser.Id)
	return utils.Success(info), nil
}
