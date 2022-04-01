package svc

import (
	"github.com/wuyan94zl/go-zero-blog/app/internal/config"
	"github.com/wuyan94zl/go-zero-blog/app/models/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel user.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: user.NewUsersModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
