package group

import (
	"context"
	"fmt"
	"github.com/wuyan94zl/IM/app/common/im"
	"github.com/wuyan94zl/IM/app/common/utils"
	"github.com/wuyan94zl/IM/app/models/groups"
	"github.com/wuyan94zl/IM/app/models/groupusers"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"

	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupAddLogic {
	return &GroupAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupAddLogic) GroupAdd(req *types.GroupAddRequest) (resp *types.Response, err error) {
	groupItem := groups.Groups{
		UserId:      l.svcCtx.AuthUser.Id,
		Title:       req.Title,
		Description: req.Description,
		ChannelId:   utils.Md5ByString(fmt.Sprintf("%s%s%d", req.Title, req.Description, time.Now().UnixNano())),
	}
	err = l.svcCtx.MysqlConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		err := l.svcCtx.GroupModel.TranCreate(ctx, session, &groupItem)
		if err != nil {
			return err
		}
		groupUserItem := groupusers.GroupUsers{
			UserId:    l.svcCtx.AuthUser.Id,
			GroupId:   groupItem.Id,
			ChannelId: groupItem.ChannelId,
			IsManager: 1,
		}
		err = l.svcCtx.GroupUserModel.TranCreate(ctx, session, &groupUserItem)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	go im.JoinChannelIds(uint64(l.svcCtx.AuthUser.Id), groupItem.ChannelId)
	return &types.Response{Message: "创建群组成功"}, nil
}
