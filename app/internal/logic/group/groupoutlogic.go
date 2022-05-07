package group

import (
	"context"

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

func (l *GroupOutLogic) GroupOut(req *types.GroupOutRequest) (resp *types.GroupResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
