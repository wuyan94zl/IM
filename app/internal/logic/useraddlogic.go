package logic

import (
	"context"
	"fmt"
	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"
	"github.com/wuyan94zl/go-zero-blog/app/models/user"
	"github.com/wuyan94zl/go-zero-blog/app/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req *types.UserAddRequest) (*utils.SuccessTmp, *utils.ErrorTmp) {
	u := user.Users{
		UserName: req.UserName,
		NickName: req.NickName,
		Password: req.Password,
		Mobile:   req.Mobile,
	}
	insert, err := l.svcCtx.UserModel.Insert(l.ctx, &u)
	if err != nil {
		return nil, utils.Error(500, err.Error())
	}
	id, _ := insert.LastInsertId()

	return utils.Success(fmt.Sprintf("用户ID：%d", id)), nil
}
