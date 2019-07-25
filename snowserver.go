package main

import (
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/CyrivlClth/snowserver/commands"
)

const (
	appName = "Snowflake server"
	detail  = "Snowflake 分布式ID工作节点"
	version = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = detail
	app.Version = version
	app.Flags = commands.ServerFlags
	app.Commands = []cli.Command{
		commands.ServerCommand,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
