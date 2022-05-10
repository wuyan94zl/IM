package group

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wuyan94zl/IM/app/common/im"
	"github.com/wuyan94zl/IM/app/common/response"
	"github.com/wuyan94zl/IM/app/models/groups"
	"github.com/wuyan94zl/IM/app/models/notices"

	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupOutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupOutLogic {
	return &GroupOutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupOutLogic) GroupOut(req *types.GroupOutRequest) (resp *types.Response, err error) {
	group, err := l.svcCtx.GroupModel.FindOne(l.ctx, req.GroupId)
	if err != nil {
		if err == groups.ErrNotFound {
			return nil, response.Error(404, "群组不存在")
		}
		return nil, err
	}
	if group.UserId == l.svcCtx.AuthUser.Id {
		return nil, response.Error(403, "群主不能退出")
	}
	u, err := l.svcCtx.GroupUserModel.IsInGroup(l.ctx, group.Id, l.svcCtx.AuthUser.Id)
	if err != nil {
		return nil, response.Error(404, "已经退出群组")
	}
	err = l.svcCtx.GroupUserModel.Delete(l.ctx, u.Id)
	if err != nil {
		return nil, err
	}

	notice := notices.Notices{
		PubUserId: l.svcCtx.AuthUser.Id,
		SubUserId: group.UserId,
		Tp:        notices.GROUP,
		Content:   fmt.Sprintf("%s 退出了 %s 群组", l.svcCtx.AuthUser.NickName, group.Title),
		Status:    1,
	}
	l.svcCtx.NoticeModel.Insert(l.ctx, &notice)
	strByte, _ := json.Marshal(notice)
	go im.SendMessageToUid(uint64(l.svcCtx.AuthUser.Id), uint64(group.UserId), string(strByte), 200)

	notice = notices.Notices{
		PubUserId: l.svcCtx.AuthUser.Id,
		SubUserId: l.svcCtx.AuthUser.Id,
		Tp:        notices.GROUP,
		Content:   fmt.Sprintf("您退出了 %s 群组", group.Title),
		Status:    1,
	}
	l.svcCtx.NoticeModel.Insert(l.ctx, &notice)
	strByte, _ = json.Marshal(notice)
	go im.SendMessageToUids(uint64(l.svcCtx.AuthUser.Id), string(strByte), 202, uint64(l.svcCtx.AuthUser.Id))

	return &types.Response{Message: "退出群组成功"}, nil
}
