package group

import (
	"context"

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

func (l *GroupJoinLogic) GroupJoin(req *types.GroupJoinRequest) (resp *types.GroupResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
