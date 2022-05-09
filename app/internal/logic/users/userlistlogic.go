package users

import (
	"context"
	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"
	"github.com/wuyan94zl/IM/app/models/user"
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
	id := l.svcCtx.AuthUser.Id
	list, err := l.svcCtx.UserModel.GetListByKeyword(req.Keyword, id)

	if err != nil {
		return nil, err
	}
	isFriend := l.isFriend(list, id)
	var rlt types.UserListResponse
	for _, v := range list {
		item := types.UserList{Id: v.Id, NickName: v.NickName, IsFriend: 0}
		if _, ok := isFriend[v.Id]; ok {
			item.IsFriend = 1
		}
		rlt.List = append(rlt.List, item)
	}
	return &rlt, nil
}

func (l *UserListLogic) isFriend(list []user.Users, userId int64) map[int64]bool {
	var id []interface{}
	isFriend := make(map[int64]bool)
	for _, v := range list {
		id = append(id, v.Id)
	}
	friend, err := l.svcCtx.UserUsersModel.IsFriend(userId, id...)
	if err != nil {
		return isFriend
	}
	for _, v := range friend {
		isFriend[v.HasUserId] = true
	}
	return isFriend
}
