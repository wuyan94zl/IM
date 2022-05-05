package svc

import (
	"github.com/wuyan94zl/IM/app/common/auth"
	"github.com/wuyan94zl/IM/app/internal/config"
	"github.com/wuyan94zl/IM/app/internal/middleware"
	"github.com/wuyan94zl/IM/app/models/hasusers"
	"github.com/wuyan94zl/IM/app/models/messages"
	"github.com/wuyan94zl/IM/app/models/notices"
	"github.com/wuyan94zl/IM/app/models/sendqueue"
	"github.com/wuyan94zl/IM/app/models/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	AuthToken      rest.Middleware
	UserModel      user.UsersModel
	UserUsersModel hasusers.UserUsersModel
	NoticeModel    notices.NoticesModel
	MessageModel   messages.MessagesModel
	SendQueueModel sendqueue.SendQueuesModel
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
		NoticeModel:    notices.NewNoticesModel(sqlx.NewMysql(c.DB.DataSource)),
		MessageModel:   messages.NewMessagesModel(sqlx.NewMysql(c.DB.DataSource)),
		SendQueueModel: sendqueue.NewSendQueuesModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
