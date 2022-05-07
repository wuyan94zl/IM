package group

import (
	"context"

	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupDelLogic {
	return &GroupDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupDelLogic) GroupDel(req *types.GroupDelRequest) (resp *types.GroupResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
