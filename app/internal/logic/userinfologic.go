package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wuyan94zl/go-zero-blog/app/models/user"

	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.Response, err error) {
	info, err := l.svcCtx.UserModel.FindOne(l.ctx, int64(req.Id))
	if err != nil {
		if err == user.ErrNotFound {
			return nil, fmt.Errorf("用户不存在")
		}
		return nil, err
	}
	strByte, _ := json.Marshal(info)
	resp.Message = string(strByte)
	return resp, nil
}
