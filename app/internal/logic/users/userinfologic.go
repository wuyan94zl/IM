package users

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/wuyan94zl/IM/app/common/utils"
	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/internal/types"

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
	// id 为啥是json.Number
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	// nick_name string
	//nickName := l.ctx.Value("nick_name").(string)
	// l.ctx.Value("key") key为token解析的map key
	fmt.Println(id)
	info, _ := l.svcCtx.UserModel.FindOne(l.ctx, id)
	return utils.Success(info), nil
}
