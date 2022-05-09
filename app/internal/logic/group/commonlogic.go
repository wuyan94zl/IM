package group

import (
	"context"
	"github.com/wuyan94zl/IM/app/common/response"
	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/models/groups"
)

func checkGroup(ctx context.Context, svcCtx *svc.ServiceContext, groupId int64) (*groups.Groups, error) {
	info, err := svcCtx.GroupModel.FindOne(ctx, groupId)
	if err != nil {
		if err == groups.ErrNotFound {
			return nil, response.Error(404, "群组不存在")
		}
		return nil, err
	}
	return info, nil
}
