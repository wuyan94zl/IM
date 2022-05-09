package friend

import (
	"context"
	"github.com/wuyan94zl/IM/app/models/messages"
	"github.com/wuyan94zl/IM/app/models/user"

	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageListLogic {
	return &MessageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageListLogic) MessageList(req *types.MessageListRequest) (resp *types.MessageListResponse, err error) {
	list, err := l.svcCtx.MessageModel.GetListByChannelId(req.ChannelId, req.MinMessageId)
	if err != nil {
		return nil, err
	}

	var rlt types.MessageListResponse
	rlt.List = []types.Message{}
	if len(list) < 1 {
		return &rlt, nil
	}

	users, _ := l.WithUser(list)
	id := l.svcCtx.AuthUser.Id
	for _, item := range list {
		t := 1
		if item.SendUserId != id {
			t = 0
		}
		rlt.List = append(rlt.List, types.Message{
			UserId:   item.SendUserId,
			NickName: users[item.SendUserId].NickName,
			Content:  item.Message,
			SendTime: item.CreateTime.Format("2006-01-02 15:01:05"),
			Tp:       int64(t),
		})
	}
	return &rlt, nil
}

func (l *MessageListLogic) WithUser(msg []messages.Messages) (map[int64]user.Users, error) {
	var uid []interface{}
	for _, v := range msg {
		uid = append(uid, v.SendUserId)
	}
	return l.svcCtx.UserModel.FindByIds(l.ctx, uid...)
}
