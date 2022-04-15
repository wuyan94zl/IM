package friend

import (
	"context"
	"encoding/json"

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
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	var rlt types.MessageListResponse
	for _, v := range list {
		t := 0
		if v.SendUserId == id {
			t = 1
		}
		rlt.List = append(rlt.List, types.Message{
			UserId:   v.SendUserId,
			NickName: "test",
			Content:  v.Message,
			SendTime: v.CreateTime.Format("2006-01-02 15:01:05"),
			Tp:       int64(t),
		})
	}
	return &rlt, nil
}
