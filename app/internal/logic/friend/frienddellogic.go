package friend

import (
	"context"
	"encoding/json"
	"github.com/wuyan94zl/go-zero-blog/app/common/im"
	"github.com/wuyan94zl/go-zero-blog/app/models/notices"
	"strconv"

	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendDelLogic {
	return &FriendDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendDelLogic) FriendDel(req *types.FriendRequest) (resp *types.FriendResponse, err error) {
	id, _ := l.ctx.Value("id").(json.Number).Int64()

	friend, err := l.svcCtx.UserUsersModel.CheckFriend(id, req.FriendId)
	if len(friend) != 2 || err != nil {
		return &types.FriendResponse{
			Status:  false,
			Message: "对方不是你的好友",
		}, nil
	}

	err = l.svcCtx.UserUsersModel.Delete(l.ctx, friend[0].Id)
	err = l.svcCtx.UserUsersModel.Delete(l.ctx, friend[1].Id)
	l.svcCtx.NoticeModel.Insert(l.ctx, &notices.Notices{PubUserId: id, SubUserId: req.FriendId, Content: "把删除从ta的好友移除", Status: 1, Tp: 2})
	if err != nil {
		return &types.FriendResponse{
			Status:  false,
			Message: "删除好友失败",
		}, nil
	}

	go im.SendMessageToChannelIds(uint64(id), strconv.FormatInt(req.FriendId, 10), 202, im.GenChannelIdByFriend(id, req.FriendId))
	return &types.FriendResponse{
		Status:  true,
		Message: "删除好友成功",
	}, nil
}
