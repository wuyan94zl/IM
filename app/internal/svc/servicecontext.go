package svc

import (
	"github.com/wuyan94zl/go-zero-blog/app/common/auth"
	"github.com/wuyan94zl/go-zero-blog/app/internal/config"
	"github.com/wuyan94zl/go-zero-blog/app/internal/middleware"
	"github.com/wuyan94zl/go-zero-blog/app/models/actionlogs"
	"github.com/wuyan94zl/go-zero-blog/app/models/hasusers"
	"github.com/wuyan94zl/go-zero-blog/app/models/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	AuthToken      rest.Middleware
	UserModel      user.UsersModel
	UserUsersModel hasusers.UserUsersModel
	ActionLogModel actionlogs.ActionLogsModel
	AuthUser       *auth.Info
}

func NewServiceContext(c config.Config) *ServiceContext {
	authUser := new(auth.Info)
	return &ServiceContext{
		Config:         c,
		AuthToken:      middleware.NewAuthTokenMiddleware(c, authUser).Handle,
		AuthUser:       authUser,
		UserModel:      user.NewUsersModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserUsersModel: hasusers.NewUserUsersModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		ActionLogModel: actionlogs.NewActionLogsModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
