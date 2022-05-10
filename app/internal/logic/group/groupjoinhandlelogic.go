package group

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wuyan94zl/IM/app/common/im"
	"github.com/wuyan94zl/IM/app/common/response"
	"github.com/wuyan94zl/IM/app/models/groupusers"
	"github.com/wuyan94zl/IM/app/models/notices"
	"time"

	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupJoinHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupJoinHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupJoinHandleLogic {
	return &GroupJoinHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupJoinHandleLogic) GroupJoinHandle(req *types.GroupJoinHandleRequest) (resp *types.Response, err error) {
	log, err := l.svcCtx.NoticeModel.FindOne(l.ctx, req.JoinId)
	if err != nil {
		if err == notices.ErrNotFound {
			return nil, response.Error(404, "操作数据不存在")
		}
		return nil, err
	}
	id := l.svcCtx.AuthUser.Id
	if log.SubUserId != id || log.Status == 1 {
		return nil, response.Error(404, "参数错误或已处理")
	}

	_, err = l.svcCtx.GroupUserModel.IsInGroup(l.ctx, log.LinkId, log.PubUserId)
	if err == nil {
		return nil, response.Error(400, "已经加入群组")
	}
	log.Status = 1
	return l.setGroupJoin(log, req.ActionType, id)
}

func (l *GroupJoinHandleLogic) setGroupJoin(log *notices.Notices, tp, id int64) (resp *types.Response, err error) {
	rlt := new(types.Response)
	sendTp := 200
	group, _ := l.svcCtx.GroupModel.FindOne(l.ctx, log.LinkId)
	if tp == 1 {
		groupUser := groupusers.GroupUsers{
			UserId:    log.PubUserId,
			GroupId:   log.LinkId,
			ChannelId: group.ChannelId,
			IsManager: 0,
		}
		l.svcCtx.GroupUserModel.Insert(l.ctx, &groupUser)
		log.IsAgree = "已同意"
		rlt.Message = "同意加入群组"
		go im.JoinChannelIds(uint64(log.PubUserId), group.ChannelId)
		sendTp = 202
	} else {
		log.IsAgree = "已拒绝"
		rlt.Message = "拒绝加入群组"
	}
	l.svcCtx.NoticeModel.Update(l.ctx, log)

	handleNotice := notices.Notices{
		PubUserId:  id,
		SubUserId:  log.PubUserId,
		Tp:         notices.GROUP,
		Content:    fmt.Sprintf("%v %s 你加入 `%s` 群组", l.svcCtx.AuthUser.NickName, log.IsAgree, group.Title),
		IsAgree:    log.IsAgree,
		Status:     1,
		CreateTime: time.Now(),
	}
	_, _ = l.svcCtx.NoticeModel.Insert(l.ctx, &handleNotice)
	strByte, _ := json.Marshal(handleNotice)
	go im.SendMessageToUid(uint64(id), uint64(handleNotice.SubUserId), string(strByte), uint8(sendTp))
	return rlt, nil
}
