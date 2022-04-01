package logic

import (
	"context"
	"github.com/wuyan94zl/go-zero-blog/app/models/user"
	"strconv"

	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"

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

func (l *UserAddLogic) UserAdd(req *types.UserAddRequest) (resp *types.Response, err error) {
	u := user.Users{
		UserName: req.UserName,
		NickName: req.NickName,
		Password: req.Password,
		Mobile:   req.Mobile,
	}
	insert, err := l.svcCtx.UserModel.Insert(l.ctx, &u)
	if err != nil {
		return nil, err
	}
	id, _ := insert.LastInsertId()
	resp.Message = strconv.Itoa(int(id))
	return resp, nil
}
