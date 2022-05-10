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

type GroupJoinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupJoinLogic {
	return &GroupJoinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupJoinLogic) GroupJoin(req *types.GroupJoinRequest) (resp *types.Response, err error) {
	info, err := checkGroup(l.ctx, l.svcCtx, req.GroupId)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.GroupUserModel.IsInGroup(l.ctx, info.Id, l.svcCtx.AuthUser.Id)
	if err != groupusers.ErrNotFound {
		if err != nil {
			return nil, err
		}
		return nil, response.Error(500, "你已加入该群组")
	}
	// 向群主发送加群申请消息
	noticeAdd := notices.Notices{
		PubUserId:  l.svcCtx.AuthUser.Id,
		SubUserId:  info.UserId,
		Tp:         notices.GROUP,
		Content:    fmt.Sprintf("%s 申请加入群聊 %s", l.svcCtx.AuthUser.NickName, info.Title),
		Note:       req.Note,
		LinkId:     info.Id,
		CreateTime: time.Now(),
	}
	insert, _ := l.svcCtx.NoticeModel.Insert(l.ctx, &noticeAdd)
	noticeAdd.Id, err = insert.LastInsertId()
	strByte, _ := json.Marshal(noticeAdd)
	go im.SendMessageToUid(uint64(l.svcCtx.AuthUser.Id), uint64(info.UserId), string(strByte), 200)
	return &types.Response{Message: "加群申请已发送"}, nil
}
