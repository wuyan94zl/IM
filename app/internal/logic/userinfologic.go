package logic

import (
	"context"
	"github.com/wuyan94zl/go-zero-blog/app/models/user"
	"github.com/wuyan94zl/go-zero-blog/app/utils"

	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"

	"github.com/jinzhu/copier"
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

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (*utils.SuccessTmp, *utils.ErrorTmp) {
	info, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		if err == user.ErrNotFound {
			return nil, utils.Error(401, "用户不存在")
		}
		return nil, utils.Error(401, err.Error())
	}
	resp := types.User{}
	copier.Copy(&resp, info)
	resp.CreateTime = info.CreateTime.Time.Format("2006-01-02 15:04:05")
	resp.UpdateTime = info.UpdateTime.Time.Format("2006-01-02 15:04:05")
	return utils.Success(resp), nil
}
