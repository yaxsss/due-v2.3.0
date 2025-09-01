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
		group.AddRouteHandler(int32(pb.Route_QuickStart), false, g.quickStart)
	})
}

func (g *Game) quickStart(ctx node.Context) {
	resp := &pb.QuickStartRes{}
	defer func() {
		if err := ctx.Response(resp); err != nil {
			log.Error("quickStart response error", "error", err)
		}
	}()

	resp.Code = pb.Code_Success
}
