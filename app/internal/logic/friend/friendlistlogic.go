package friend

import (
	"context"
	"github.com/wuyan94zl/IM/app/common/im"
	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList() (resp *types.FriendList, err error) {
	id := l.svcCtx.AuthUser.Id

	friends, err := l.svcCtx.UserModel.Friends(l.svcCtx.UserUsersModel, id)
	list := new(types.FriendList)
	for _, friend := range friends {
		list.List = append(list.List, types.Friend{UserId: friend.Id, NickName: friend.NickName, IsFriend: 1, ChannelId: im.GenChannelIdByFriend(id, friend.Id)})
	}
	return list, nil
}
