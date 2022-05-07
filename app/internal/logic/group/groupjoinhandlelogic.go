package group

import (
	"context"

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

func (l *GroupJoinHandleLogic) GroupJoinHandle(req *types.GroupJoinHandleRequest) (resp *types.GroupResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
