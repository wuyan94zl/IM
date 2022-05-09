package friend

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wuyan94zl/IM/app/common/im"
	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"
	"github.com/wuyan94zl/IM/app/models/notices"
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
	id := l.svcCtx.AuthUser.Id

	friend, err := l.svcCtx.UserUsersModel.CheckFriend(id, req.FriendId)
	if friend != nil {
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

	noticeAdd := notices.Notices{
		PubUserId:  id,
		SubUserId:  req.FriendId,
		Tp:         notices.FRIEND,
		Content:    fmt.Sprintf("%s 请求添加您为好友", l.svcCtx.AuthUser.NickName),
		Note:       "",
		CreateTime: time.Now(),
	}
	insert, err := l.svcCtx.NoticeModel.Insert(l.ctx, &noticeAdd)
	switch err {
	case nil:
		noticeAdd.Id, _ = insert.LastInsertId()
		strByte, _ := json.Marshal(noticeAdd)
		go im.SendMessageToUid(uint64(id), uint64(req.FriendId), string(strByte), 200)
		return &types.FriendResponse{
			Status:  true,
			Message: "添加好友请求已发送",
		}, nil
	default:
		return nil, err
	}
}
