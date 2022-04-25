package friend

import (
	"context"
	"encoding/json"
	"fmt"
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

	isFriend, err := l.svcCtx.UserUsersModel.CheckFriend(id, req.FriendId)
	if len(isFriend) == 2 || err != nil {
		return &types.FriendResponse{
			Status:  false,
			Message: "对方已经是你好友",
		}, nil
	}
	log, err := l.svcCtx.NoticeModel.CheckSendAddFriend(id, req.FriendId)
	if log != nil {
		return &types.FriendResponse{
			Status:  false,
			Message: "添加好友请求已经发送",
		}, nil
	}

	user, _ := l.svcCtx.UserModel.FindOne(l.ctx, id)
	noticeAdd := notices.Notices{PubUserId: id, SubUserId: req.FriendId, Tp: 1, Content: fmt.Sprintf("%s 请求添加您为好友", user.NickName), CreateTime: time.Now()}
	insert, _ := l.svcCtx.NoticeModel.Insert(l.ctx, &noticeAdd)
	noticeAdd.Id, err = insert.LastInsertId()
	switch err {
	case nil:
		strByte, _ := json.Marshal(noticeAdd)
		im.SendMessageToUid(uint64(id), uint64(req.FriendId), string(strByte), 200)
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
