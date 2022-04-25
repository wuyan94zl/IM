package main

import (
	"flag"
	"fmt"
	"github.com/wuyan94zl/go-zero-blog/app/common/im"
	"github.com/wuyan94zl/go-zero-blog/app/internal/config"
	"github.com/wuyan94zl/go-zero-blog/app/internal/handler"
	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/app-api.yaml", "the config file")

func main() {
	flag.Parse()
	logx.Disable()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	//server.AddRoute(rest.Route{
	//	Method: http.MethodGet,
	//	Path:   "/ws",
	//	Handler: func(w http.ResponseWriter, r *http.Request) {
	//		im.RunWs(w, r, ctx)
	//	},
	//})
	go im.Run(ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
