package friend

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wuyan94zl/go-zero-blog/app/models/actionlogs"

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
	fmt.Println(friend, err)
	if len(friend) != 2 || err != nil {
		return &types.FriendResponse{
			Status:  false,
			Message: "对方不是你的好友",
		}, nil
	}

	err = l.svcCtx.UserUsersModel.Delete(l.ctx, friend[0].Id)
	err = l.svcCtx.UserUsersModel.Delete(l.ctx, friend[1].Id)
	l.svcCtx.ActionLogModel.Insert(l.ctx, &actionlogs.ActionLogs{PubUserId: id, SubUserId: req.FriendId, HandleResult: "删除", Status: 1, Tp: 2})
	if err != nil {
		return &types.FriendResponse{
			Status:  false,
			Message: "删除好友失败",
		}, nil
	}

	return &types.FriendResponse{
		Status:  true,
		Message: "删除好友成功",
	}, nil
}
