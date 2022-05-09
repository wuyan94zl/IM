package svc

import (
	"github.com/wuyan94zl/IM/app/common/auth"
	"github.com/wuyan94zl/IM/app/internal/config"
	"github.com/wuyan94zl/IM/app/internal/middleware"
	"github.com/wuyan94zl/IM/app/models/groups"
	"github.com/wuyan94zl/IM/app/models/groupusers"
	"github.com/wuyan94zl/IM/app/models/hasusers"
	"github.com/wuyan94zl/IM/app/models/messages"
	"github.com/wuyan94zl/IM/app/models/notices"
	"github.com/wuyan94zl/IM/app/models/sendqueue"
	"github.com/wuyan94zl/IM/app/models/user"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	RedisCli       *redis.Redis
	AuthToken      rest.Middleware
	UserModel      user.UsersModel
	UserUsersModel hasusers.UserUsersModel
	NoticeModel    notices.NoticesModel
	MessageModel   messages.MessagesModel
	SendQueueModel sendqueue.SendQueuesModel
	GroupModel     groups.GroupsModel
	GroupUserModel groupusers.GroupUsersModel
	MysqlConn      sqlx.SqlConn
	AuthUser       *auth.Info
}

func NewServiceContext(c config.Config) *ServiceContext {
	authUser := new(auth.Info)
	redisCli := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})
	return &ServiceContext{
		Config:         c,
		RedisCli:       redisCli,
		MysqlConn:      sqlx.NewMysql(c.DB.DataSource),
		AuthToken:      middleware.NewAuthTokenMiddleware(c, authUser, redisCli).Handle,
		AuthUser:       authUser,
		UserModel:      user.NewUsersModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserUsersModel: hasusers.NewUserUsersModel(sqlx.NewMysql(c.DB.DataSource)),
		NoticeModel:    notices.NewNoticesModel(sqlx.NewMysql(c.DB.DataSource)),
		MessageModel:   messages.NewMessagesModel(sqlx.NewMysql(c.DB.DataSource)),
		SendQueueModel: sendqueue.NewSendQueuesModel(sqlx.NewMysql(c.DB.DataSource)),
		GroupModel:     groups.NewGroupsModel(sqlx.NewMysql(c.DB.DataSource)),
		GroupUserModel: groupusers.NewGroupUsersModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
