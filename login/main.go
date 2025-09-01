package main

import (
	"duedemo/service"

	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/registry/etcd/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/mode"
)

func main() {
	mode.SetMode(mode.DebugMode)

	container := due.NewContainer()

	locator := redis.NewLocator()

	registry := etcd.NewRegistry()

	component := node.NewNode(
		node.WithLocator(locator),
		node.WithRegistry(registry),
	)

	container.Add(component)

	service.NewLogin(component.Proxy()).Init()

	container.Serve()
}
