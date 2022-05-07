// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	friend "github.com/wuyan94zl/IM/app/internal/handler/friend"
	group "github.com/wuyan94zl/IM/app/internal/handler/group"
	users "github.com/wuyan94zl/IM/app/internal/handler/users"
	"github.com/wuyan94zl/IM/app/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: users.UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: users.UserLoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthToken},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/info",
					Handler: users.UserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/list",
					Handler: users.UserListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthToken},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/friend/add",
					Handler: friend.FriendAddHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/friend/handle",
					Handler: friend.FriendHandleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/friend/del",
					Handler: friend.FriendDelHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/friend/list",
					Handler: friend.FriendListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthToken},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/message/list",
					Handler: friend.MessageListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/notice/list",
					Handler: friend.NoticeListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthToken},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/group/add",
					Handler: group.GroupAddHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/group/edit",
					Handler: group.GroupEditHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/group/Del",
					Handler: group.GroupDelHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/group/join",
					Handler: group.GroupJoinHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/group/out",
					Handler: group.GroupOutHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/group/join/handle",
					Handler: group.GroupJoinHandleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/group/remove",
					Handler: group.GroupRemoveHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/group/list",
					Handler: group.GroupListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
