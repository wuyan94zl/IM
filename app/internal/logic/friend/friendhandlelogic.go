package friend

import (
	"context"
	"encoding/json"
	"github.com/wuyan94zl/go-zero-blog/app/common/im"
	"github.com/wuyan94zl/go-zero-blog/app/models/hasusers"

	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendHandleLogic {
	return &FriendHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendHandleLogic) FriendHandle(req *types.FriendHandleRequest) (resp *types.FriendResponse, err error) {
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	log, err := l.svcCtx.ActionLogModel.FindOne(l.ctx, req.ActionLogId)
	if err != nil {
		return nil, err
	}

	if log.SubUserId != id || log.Status == 1 {
		return &types.FriendResponse{
			Status:  false,
			Message: "参数错误或已处理",
		}, nil
	}

	friend, err := l.svcCtx.UserUsersModel.CheckFriend(log.PubUserId, log.SubUserId)
	if len(friend) > 0 || err != nil {
		return &types.FriendResponse{Status: false, Message: "已添加完成"}, nil
	}
	log.Status = 1
	switch req.ActionType {
	case 1:
		l.svcCtx.UserUsersModel.Insert(l.ctx, &hasusers.UserUsers{UserId: id, HasUserId: log.PubUserId})
		l.svcCtx.UserUsersModel.Insert(l.ctx, &hasusers.UserUsers{UserId: log.PubUserId, HasUserId: id})
		log.HandleResult = "同意"
		l.svcCtx.ActionLogModel.Update(l.ctx, log)
		im.SendMessageToUid(uint64(id), uint64(log.PubUserId), "同意了你的好友申请")
		return &types.FriendResponse{Status: true, Message: "添加成功"}, nil
	default:
		log.HandleResult = "拒绝"
		l.svcCtx.ActionLogModel.Update(l.ctx, log)
		im.SendMessageToUid(uint64(id), uint64(log.PubUserId), "拒绝了你的好友申请")
		return &types.FriendResponse{Status: false, Message: "拒绝添加好友"}, nil
	}
}
