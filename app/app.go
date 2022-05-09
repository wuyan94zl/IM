package main

import (
	"flag"
	"fmt"
	"github.com/wuyan94zl/IM/app/common/im"
	"github.com/wuyan94zl/IM/app/internal/config"
	"github.com/wuyan94zl/IM/app/internal/handler"
	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/app-api.yaml", "the config file")

func main() {
	flag.Parse()
	//logx.Disable()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	//server.Start()

	group := service.NewServiceGroup()
	defer group.Stop()
	group.Add(server)
	group.Add(im.Server{Ctx: ctx})
	group.Start()
}
