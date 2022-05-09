package group

import (
	"context"
	"github.com/wuyan94zl/IM/app/common/response"
	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupEditLogic {
	return &GroupEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupEditLogic) GroupEdit(req *types.GroupEditRequest) (resp *types.Response, err error) {
	info, err := checkGroup(l.ctx,l.svcCtx,req.GroupId)
	if err != nil {
		return nil, err
	}
	if info.UserId != l.svcCtx.AuthUser.Id {
		return nil, response.Error(403, "无权限更新")
	}
	info.Title = req.Title
	info.Description = req.Description
	err = l.svcCtx.GroupModel.Update(l.ctx, info)
	if err != nil {
		return nil, err
	}
	return &types.Response{Message: "更新成功"}, nil
}
