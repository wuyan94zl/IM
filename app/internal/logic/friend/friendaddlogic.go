package friend

import (
	"context"
	"encoding/json"
	"github.com/wuyan94zl/go-zero-blog/app/common/im"
	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"
	"github.com/wuyan94zl/go-zero-blog/app/models/notices"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
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
	log, err := l.svcCtx.NoticeModel.CheckSendAddFriend(id, req.FriendId)
	if log != nil {
		strByte, _ := json.Marshal(types.Notice{
			Id: log.Id, Tp: log.Tp, IsAgree: log.IsAgree, Status: log.Status,
			NickName: "", Content: log.Content, CreateTime: log.CreateTime.Format("2006-01-06 15:04:01"),
		})
		im.SendMessageToUid(uint64(id), uint64(req.FriendId), string(strByte))
		return &types.FriendResponse{
			Status:  false,
			Message: "添加好友请求已经发送",
		}, nil
	}

	link, _ := l.svcCtx.NoticeModel.Insert(l.ctx, &notices.Notices{PubUserId: id, SubUserId: id, Tp: 1, Content: "你请求添加XXX为好友", IsAgree: "未处理", Status: 1, CreateTime: time.Now()})
	linkId, _ := link.LastInsertId()

	addLog := notices.Notices{PubUserId: id, SubUserId: req.FriendId, Tp: 1, Content: "请求添加您为好友！", CreateTime: time.Now(), LinkId: linkId}
	ins, err := l.svcCtx.NoticeModel.Insert(l.ctx, &addLog)

	switch err {
	case nil:
		lastId, _ := ins.LastInsertId()
		strByte, _ := json.Marshal(types.Notice{
			Id: lastId, Tp: 1, IsAgree: "",
			NickName: "", Content: addLog.Content, CreateTime: addLog.CreateTime.Format("2006-01-06 15:04:01"),
		})
		im.SendMessageToUid(uint64(id), uint64(req.FriendId), string(strByte))
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
