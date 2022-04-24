package friend

import (
	"context"
	"encoding/json"
	"github.com/wuyan94zl/go-zero-blog/app/models/user"

	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"

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
	nickName := l.ctx.Value("nick_name").(string)
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	var rlt types.MessageListResponse
	var friend *user.Users
	for i := len(list) - 1; i >= 0; i-- {
		t, n := 1, nickName
		if list[i].SendUserId != id {
			t = 0
			if friend == nil {
				friend, _ = l.svcCtx.UserModel.FindOne(l.ctx, list[i].SendUserId)
			}
			n = friend.NickName
		}
		rlt.List = append(rlt.List, types.Message{
			UserId:   list[i].SendUserId,
			NickName: n,
			Content:  list[i].Message,
			SendTime: list[i].CreateTime.Format("2006-01-02 15:01:05"),
			Tp:       int64(t),
		})
	}
	return &rlt, nil
}
