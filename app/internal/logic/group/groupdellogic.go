package group

import (
	"context"
	"fmt"
	"github.com/wuyan94zl/IM/app/common/response"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

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

func (l *GroupDelLogic) GroupDel(req *types.GroupDelRequest) (resp *types.Response, err error) {
	info, err := checkGroup(l.ctx,l.svcCtx,req.GroupId)
	if err != nil {
		return nil, err
	}

	if info.UserId != l.svcCtx.AuthUser.Id {
		return nil, response.Error(403, "无删除权限")
	}

	err = l.svcCtx.MysqlConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		err := l.svcCtx.GroupUserModel.TranDeleteByGroupId(ctx, session, req.GroupId)
		if err != nil {
			return err
		}
		err = l.svcCtx.GroupModel.TranDelete(ctx, session, req.GroupId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, response.Error(500, fmt.Sprintf("删除失败：%v", err))
	}

	return &types.Response{Message: "删除成功"}, nil
}
