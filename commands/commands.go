package commands

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"golang.org/x/sync/errgroup"

	"github.com/CyrivlClth/snowserver/config"
	gsrv "github.com/CyrivlClth/snowserver/grpc/server"
	hsrv "github.com/CyrivlClth/snowserver/http/server"
)

var (
	grpcPort     int
	httpPort     int
	workerID     int
	dataCenterID int
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
	cli.IntFlag{
		Name:        "work",
		Usage:       "指定服务的工作节点ID",
		Value:       0,
		Destination: &workerID,
	},
	cli.IntFlag{
		Name:        "center",
		Usage:       "指定服务的数据中心ID",
		Value:       0,
		Destination: &dataCenterID,
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
	err := config.Init(int64(workerID), int64(dataCenterID))
	if err != nil {
		return err
	}
	g := errgroup.Group{}
	g.Go(func() error {
		return gsrv.Run(fmt.Sprintf(":%d", grpcPort))
	})

	g.Go(func() error {
		gin.SetMode(gin.ReleaseMode)
		return hsrv.Run(fmt.Sprintf(":%d", httpPort))
	})

	return g.Wait()
}

func RunGrpcAction(c *cli.Context) error {
	err := config.Init(int64(workerID), int64(dataCenterID))
	if err != nil {
		return err
	}
	return gsrv.Run(fmt.Sprintf(":%d", grpcPort))
}

func RunHttpAction(c *cli.Context) error {
	gin.SetMode(gin.ReleaseMode)
	err := config.Init(int64(workerID), int64(dataCenterID))
	if err != nil {
		return err
	}
	return hsrv.Run(fmt.Sprintf(":%d", httpPort))
}
