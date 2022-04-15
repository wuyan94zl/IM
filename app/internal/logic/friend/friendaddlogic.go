package friend

import (
	"context"
	"encoding/json"
	"github.com/wuyan94zl/go-zero-blog/app/common/im"
	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"
	"github.com/wuyan94zl/go-zero-blog/app/models/actionlogs"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendAddLogic {
	return &FriendAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendAddLogic) FriendAdd(req *types.FriendRequest) (resp *types.FriendResponse, err error) {
	id, _ := l.ctx.Value("id").(json.Number).Int64()

	friend, err := l.svcCtx.UserUsersModel.CheckFriend(id, req.FriendId)
	if len(friend) == 2 || err != nil {
		return &types.FriendResponse{
			Status:  false,
			Message: "对方已经是你好友",
		}, nil
	}
	log, err := l.svcCtx.ActionLogModel.CheckSendAddFriend(id, req.FriendId)
	if log != nil {
		return &types.FriendResponse{
			Status:  false,
			Message: "添加好友请求已经发送",
		}, nil
	}
	_, err = l.svcCtx.ActionLogModel.Insert(l.ctx, &actionlogs.ActionLogs{PubUserId: id, SubUserId: req.FriendId, Tp: 1})
	switch err {
	case nil:
		im.SendMessageToUid(uint64(id), uint64(req.FriendId), "请求添加好友")
		return &types.FriendResponse{
			Status:  true,
			Message: "添加好友请求已发送",
		}, nil
	default:
		return &types.FriendResponse{
			Status:  false,
			Message: "添加好友失败",
		}, nil
	}
}
