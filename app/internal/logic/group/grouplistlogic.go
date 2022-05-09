package group

import (
	"context"
	"fmt"
	"github.com/wuyan94zl/IM/app/models/groupusers"

	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupListLogic) GroupList(req *types.GroupRequest) (resp *types.GroupListResponse, err error) {
	groups, err := l.svcCtx.GroupUserModel.InGroups(l.ctx, l.svcCtx.AuthUser.Id)
	fmt.Println(groups,err)
	if err != nil {
		return nil, err
	}
	var rlt types.GroupListResponse
	rlt.List = []types.Group{}
	if len(groups) == 0 {
		return &rlt, nil
	}
	err = WithGroup(l.svcCtx, groups, &rlt)
	if err != nil {
		return nil, err
	}
	return &rlt, nil
}

func WithGroup(svcCtx *svc.ServiceContext, list []groupusers.GroupUsers, resp *types.GroupListResponse) error {
	var groupIds []interface{}
	for _, v := range list {
		groupIds = append(groupIds, v.GroupId)
	}
	groups, err := svcCtx.GroupModel.FindByIds(groupIds...)
	if err != nil {
		return err
	}
	for _, v := range list {
		resp.List = append(resp.List, types.Group{
			Id:          v.GroupId,
			Title:       groups[v.GroupId].Title,
			Description: groups[v.GroupId].Description,
			ChannelId: groups[v.GroupId].ChannelId,
			IsManager:   v.IsManager,
		})
	}
	return nil
}
