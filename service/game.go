package service

import (
	"context"
	"duedemo/pb"

	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/log"
)

type Game struct {
	proxy *node.Proxy
	ctx   context.Context
}

func NewGame(proxy *node.Proxy) *Game {
	return &Game{
		proxy: proxy,
		ctx:   context.Background(),
	}
}

func (g *Game) Init() {
	g.proxy.Router().Group(func(group *node.RouterGroup) {
		group.AddRouteHandler(int32(pb.Route_QuickStart), true, g.quickStart)
	})
}

func (g *Game) quickStart(ctx node.Context) {
	ctx.UnbindGate(1)
	resp := &pb.QuickStartRes{}
	defer func() {
		if err := ctx.Response(resp); err != nil {
			log.Error("quickStart response error", "error", err)
		}
	}()
	log.Info("quickStart ok")
	ctx.BindGate(1)
	resp.Code = pb.Code_Success
}
