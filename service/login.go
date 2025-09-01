package service

import (
	"context"
	"duedemo/pb"
	"duedemo/shared/jwt"

	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/log"
)

type Login struct {
	ctx   context.Context
	proxy *node.Proxy
	jwt   *jwt.JWT
}

func NewLogin(proxy *node.Proxy) *Login {
	return &Login{
		ctx:   context.Background(),
		proxy: proxy,
		jwt:   jwt.Instance(),
	}
}

func (l *Login) Init() {
	l.proxy.Router().Group(func(group *node.RouterGroup) {
		group.AddRouteHandler(int32(pb.Route_Register), false, l.register)
		group.AddRouteHandler(int32(pb.Route_Login), false, l.login)
	})
}

// 用户注册
func (l *Login) register(ctx node.Context) {
	req := &pb.RegisterReq{}
	resp := &pb.RegisterResp{}
	defer func() {
		if err := ctx.Response(resp); err != nil {
			log.Error("register response error", "error", err)
		}
	}()
	var err error
	if err = ctx.Parse(req); err != nil {
		log.Error("register parse error", "error", err)
		resp.Code = pb.Code_Failed
		return
	}

	if req.Name != "admin" || req.Password != "admin123" {
		resp.Code = pb.Code_Failed
		return
	}

	resp.Code = pb.Code_Success
}

// 用户登录
func (l *Login) login(ctx node.Context) {
	req := &pb.LoginReq{}
	resp := &pb.LoginResp{}
	defer func() {
		if err := ctx.Response(resp); err != nil {
			log.Error("login response error", "error", err)
		}
	}()
	var err error
	if err = ctx.Parse(req); err != nil {
		log.Error("login parse error", "error", err)
		resp.Code = pb.Code_Failed
		return
	}

	if req.Name != "admin" {
		resp.Code = pb.Code_Failed
		return
	}

	resp.Code = pb.Code_Success
}
