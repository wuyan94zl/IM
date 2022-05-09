package friend

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wuyan94zl/IM/app/common/im"
	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"
	"github.com/wuyan94zl/IM/app/models/hasusers"
	"github.com/wuyan94zl/IM/app/models/notices"
	"time"

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
	id := l.svcCtx.AuthUser.Id
	log, err := l.svcCtx.NoticeModel.FindOne(l.ctx, req.ActionLogId)
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
	if friend != nil {
		return &types.FriendResponse{Status: false, Message: "已添加完成"}, nil
	}
	log.Status = 1
	return l.setFriend(log, req.ActionType, id)
}

func (l *FriendHandleLogic) setFriend(log *notices.Notices, tp, id int64) (resp *types.FriendResponse, err error) {
	rlt := new(types.FriendResponse)
	if tp == 1 {
		l.svcCtx.UserUsersModel.Insert(l.ctx, &hasusers.UserUsers{UserId: id, HasUserId: log.PubUserId, ChannelId: im.GenChannelIdByFriend(id, log.PubUserId)})
		l.svcCtx.UserUsersModel.Insert(l.ctx, &hasusers.UserUsers{UserId: log.PubUserId, HasUserId: id, ChannelId: im.GenChannelIdByFriend(id, log.PubUserId)})
		log.IsAgree = "已同意"
		rlt.Status = true
		rlt.Message = "添加成功"
	} else {
		tp = 0
		log.IsAgree = "已拒绝"
		rlt.Status = false
		rlt.Message = "拒绝添加好友"
	}
	l.svcCtx.NoticeModel.Update(l.ctx, log)

	handleNotice := notices.Notices{
		PubUserId:  id,
		SubUserId:  log.PubUserId,
		Tp:         notices.FRIEND,
		Content:    fmt.Sprintf("%v %s 你的添加好友请求", l.svcCtx.AuthUser.NickName, log.IsAgree),
		IsAgree:    log.IsAgree,
		Status:     1,
		CreateTime: time.Now(),
	}
	noticeIns, _ := l.svcCtx.NoticeModel.Insert(l.ctx, &handleNotice)
	insertId, _ := noticeIns.LastInsertId()
	handleNotice.Id = insertId

	if tp == 1 {
		channelId := im.GenChannelIdByFriend(id, log.PubUserId)
		go im.JoinChannelIds(uint64(id), channelId)
		go im.JoinChannelIds(uint64(log.PubUserId), channelId)
	}
	strByte, _ := json.Marshal(handleNotice)
	go im.SendMessageToUid(uint64(id), uint64(handleNotice.SubUserId), string(strByte), uint8(200+tp))
	return rlt, nil
}
