package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// VERSION 版本号，可以通过编译的方式指定版本号：go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "0.0.1"

func main() {
	ctx := context.Background()

	app := cli.NewApp()
	app.Name = "zgo"
	app.Version = VERSION
	app.Usage = "zgo cli"
	app.Commands = []*cli.Command{
		runInitCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Println(err.Error())
	}
}

func runInitCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "运行初始化服务",
		Action: func(c *cli.Context) error {
			log.Println("运行初始化服务...")
			return nil
		},
	}
}
