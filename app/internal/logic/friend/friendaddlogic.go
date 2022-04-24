package friend

import (
	"context"
	"encoding/json"
	"github.com/wuyan94zl/go-zero-blog/app/common/im"
	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"
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

	isFriend, err := l.svcCtx.UserUsersModel.CheckFriend(id, req.FriendId)
	if len(isFriend) == 2 || err != nil {
		return &types.FriendResponse{
			Status:  false,
			Message: "对方已经是你好友",
		}, nil
	}
	log, err := l.svcCtx.NoticeModel.CheckSendAddFriend(id, req.FriendId)
	if log != nil {
		//strByte, _ := json.Marshal(types.Notice{
		//	Id: log.Id, Tp: log.Tp, IsAgree: log.IsAgree, Status: log.Status,
		//	NickName: "", Content: log.Content, CreateTime: log.CreateTime.Format("2006-01-06 15:04:01"),
		//})
		//im.SendMessageToUid(uint64(id), uint64(req.FriendId), string(strByte))
		return &types.FriendResponse{
			Status:  false,
			Message: "添加好友请求已经发送",
		}, nil
	}

	user, _ := l.svcCtx.UserModel.FindOne(l.ctx, id)
	friend, _ := l.svcCtx.UserModel.FindOne(l.ctx, req.FriendId)
	link, notice, err := l.svcCtx.NoticeModel.AddFriend(user.Id, friend.Id, user.NickName, friend.NickName)
	switch err {
	case nil:
		strByte, _ := json.Marshal(types.Notice{
			Id: notice.Id, Tp: 1, IsAgree: "", LinkId: notice.LinkId,
			NickName: "", Content: notice.Content, CreateTime: notice.CreateTime.Format("2006-01-06 15:04:01"),
		})
		im.SendMessageToUid(uint64(id), uint64(req.FriendId), string(strByte), 200)
		strByte, _ = json.Marshal(types.Notice{
			Id: link.Id, Tp: 1, IsAgree: link.IsAgree, LinkId: link.LinkId, Status: link.Status,
			NickName: "", Content: link.Content, CreateTime: link.CreateTime.Format("2006-01-06 15:04:01"),
		})
		im.SendMessageToUid(uint64(id), uint64(id), string(strByte), 200)
		return &types.FriendResponse{
			Status:  true,
			Message: "添加好友请求已发送",
		}, nil
	default:
		return &types.FriendResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}
}
