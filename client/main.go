package main

import (
	"github.com/dobyte/due/network/ws/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/client"
	"github.com/dobyte/due/v2/mode"
)

func main() {
	mode.SetMode(mode.DebugMode)

	container := due.NewContainer()

	component := client.NewClient(
		client.WithClient(ws.NewClient()),
	)

	NewLogin(component.Proxy()).Init()

	container.Add(component)

	container.Serve()
}
