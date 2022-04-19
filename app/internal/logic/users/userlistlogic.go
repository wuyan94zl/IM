package users

import (
	"context"
	"fmt"

	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListRequest) (resp *types.UserListResponse, err error) {
	list, err := l.svcCtx.UserModel.GetListByKeyword(req.Keyword)
	fmt.Println(list,err)
	if err != nil {
		return nil, err
	}
	var rlt types.UserListResponse
	for _, v := range list {
		rlt.List = append(rlt.List, types.UserList{Id: v.Id, NickName: v.NickName})
	}
	return &rlt, nil
}
