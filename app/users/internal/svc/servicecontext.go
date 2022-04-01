package svc

import (
	"github.com/wuyan94zl/go-zero-blog/app/users/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
