package commands

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"golang.org/x/sync/errgroup"

	gsrv "github.com/CyrivlClth/snowserver/grpc/server"
	hsrv "github.com/CyrivlClth/snowserver/http/server"
)

var (
	grpcPort int
	httpPort int
)

var ServerFlags = []cli.Flag{
	cli.IntFlag{
		Name:        "gp",
		Usage:       "指定grpc服务的端口号，默认50051",
		Value:       50051,
		Destination: &grpcPort,
	},
	cli.IntFlag{
		Name:        "hp",
		Usage:       "指定http服务的端口号，默认50051",
		Value:       8080,
		Destination: &httpPort,
	},
	cli.BoolFlag{
		Name:  "d",
		Usage: "指定服务后台运行",
	},
}

var ServerCommand = cli.Command{
	Name:   "server",
	Usage:  "运行grpc和http服务",
	Action: RunAllAction,
	Subcommands: []cli.Command{
		{
			Name:   "grpc",
			Usage:  "仅运行grpc服务",
			Action: RunGrpcAction,
		},
		{
			Name:   "http",
			Usage:  "仅运行http服务",
			Action: RunHttpAction,
		},
	},
}

func RunAllAction(c *cli.Context) error {
	g := errgroup.Group{}
	g.Go(func() error {
		return RunGrpcAction(c)
	})

	g.Go(func() error {
		return RunHttpAction(c)
	})

	return g.Wait()
}

func RunGrpcAction(c *cli.Context) error {
	return gsrv.Run(fmt.Sprintf(":%d", grpcPort))
}

func RunHttpAction(c *cli.Context) error {
	gin.SetMode(gin.ReleaseMode)
	return hsrv.Run(fmt.Sprintf(":%d", httpPort))
}
