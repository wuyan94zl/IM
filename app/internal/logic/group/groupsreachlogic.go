package group

import (
	"context"
	"github.com/wuyan94zl/IM/app/common/response"
	"github.com/wuyan94zl/IM/app/models/groups"

	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupSreachLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupSreachLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupSreachLogic {
	return &GroupSreachLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupSreachLogic) GroupSreach(req *types.GroupSreachRequest) (resp *types.GroupSreachResponse, err error) {
	group, err := l.svcCtx.GroupModel.FindByTitle(l.ctx, req.Title)
	if err != nil {
		if err == groups.ErrNotFound {
			return nil, response.Error(404, "群组名称不存在")
		}
		return nil, err
	}
	manager, _ := l.svcCtx.UserModel.FindOne(l.ctx, group.UserId)
	_, err = l.svcCtx.GroupUserModel.IsInGroup(l.ctx, group.Id, l.svcCtx.AuthUser.Id)
	isJoin := 1
	if err != nil {
		isJoin = 0
	}
	return &types.GroupSreachResponse{
		Id:          group.Id,
		Title:       group.Title,
		Description: group.Description,
		Manager:     manager.NickName,
		IsJoin:      int64(isJoin),
	}, nil
}
