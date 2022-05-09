package friend

import (
	"context"
	"encoding/json"
	"github.com/wuyan94zl/IM/app/common/im"
	"github.com/wuyan94zl/IM/app/common/response"
	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"
	"github.com/wuyan94zl/IM/app/models/notices"

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

func (l *FriendDelLogic) FriendDel(req *types.FriendRequest) (resp *types.Response, err error) {
	id := l.svcCtx.AuthUser.Id

	friend, err := l.svcCtx.UserUsersModel.CheckFriend(id, req.FriendId)
	if friend == nil {
		return nil, response.Error(500, "对方不是你的好友")
	}
	_, err = l.svcCtx.UserUsersModel.DeleteFriendByChannelId(im.GenChannelIdByFriend(id, req.FriendId))
	if err != nil {
		return nil, response.Error(500, "删除好友失败"+err.Error())
	}
	notice := notices.Notices{
		PubUserId: id,
		SubUserId: req.FriendId,
		Content:   l.svcCtx.AuthUser.NickName + " 把您从ta的好友列表中移除",
		Status:    1,
		Tp:        notices.FRIEND,
	}
	l.svcCtx.NoticeModel.Insert(l.ctx, &notice)
	strByte, _ := json.Marshal(notice)
	go im.SendMessageToUid(uint64(id), uint64(req.FriendId), string(strByte), 201)
	return &types.Response{
		Message: "删除好友成功",
	}, nil
}
