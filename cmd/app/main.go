package main

import (
	"context"
	"os"
	"github.com/suisrc/zgo/app"
	"github.com/suisrc/zgo/modules/logger"

	"github.com/urfave/cli/v2"
)

// VERSION 版本号，可以通过编译的方式指定版本号：go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "0.0.1"

func main() {
	ctx := context.Background()

	app := cli.NewApp()
	app.Name = "zgo"
	app.Version = VERSION
	app.Usage = "GIN + ENT+ CASBIN + WIRE"
	app.Commands = []*cli.Command{
		runWebCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.Errorf(ctx, err.Error())
	}
}

func runWebCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "web",
		Usage: "运行web服务",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "conf",
				Aliases:     []string{"c"},
				Usage:       "配置文件(.json,.yaml,.toml)",
				DefaultText: "zgo.toml",
				//Required:   true,
			},
		},
		Action: func(c *cli.Context) error {
			return app.Run(ctx,
				app.SetConfigFile(c.String("conf")),
				app.SetVersion(VERSION))
		},
	}
}
