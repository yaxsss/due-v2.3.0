package main

import (
	"duedemo/pb"

	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/client"
	"github.com/dobyte/due/v2/log"
)

type Login struct {
	proxy *client.Proxy
}

func NewLogin(proxy *client.Proxy) *Login {
	return &Login{proxy: proxy}
}

func (l *Login) Init() {
	l.proxy.AddRouteHandler(int32(pb.Route_Register), l.register)
	l.proxy.AddRouteHandler(int32(pb.Route_Login), l.login)
	l.proxy.AddHookListener(cluster.Start, l.startHandler)
	l.proxy.AddEventListener(cluster.Connect, l.connectHandler)
}

func (l *Login) register(ctx *client.Context) {
	// req := &pb.RegisterReq{}
	resp := &pb.RegisterResp{}
	err := ctx.Parse(resp)
	if err != nil {
		log.Error("register parse error", "error", err)
		return
	}

	if resp.Code != pb.Code_Success {
		log.Error("register failed", "error", resp)
		return
	}

	log.Info("register success")
}

func (l *Login) login(ctx *client.Context) {
	resp := &pb.LoginResp{}
	err := ctx.Parse(resp)
	if err != nil {
		log.Error("login parse error", "error", err)
		return
	}

	if resp.Code != pb.Code_Success {
		log.Error("login failed", "error", resp)
		return
	}

	msg := &cluster.Message{
		Route: int32(pb.Route_QuickStart),
		Data:  []byte{},
	}

	err = ctx.Conn().Push(msg)
	if err != nil {
		log.Error("login push error", "error", err)
	}
}

func (l *Login) startHandler(proxy *client.Proxy) {
	if _, err := proxy.Dial(); err != nil {
		log.Error("start handler dial error", "error", err)
		return
	}
}

func (l *Login) connectHandler(conn *client.Conn) {
	log.Info("connect handler")
	msg := &cluster.Message{
		Route: int32(pb.Route_Login),
		Data: &pb.LoginReq{
			Name:     "123456",
			Password: "123456",
		},
	}
	err := conn.Push(msg)
	if err != nil {
		log.Error("connect handler push error", "error", err)
	}
}
