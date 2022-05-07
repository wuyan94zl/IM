package group

import (
	"context"

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

func (l *GroupEditLogic) GroupEdit(req *types.GroupEditRequest) (resp *types.GroupResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
