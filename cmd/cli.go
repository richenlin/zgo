// Copyright 2020 Kratos Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"

	cli "github.com/jawher/mow.cli"
	"github.com/mgutz/ansi"
)

func main() {

	var verbose *bool

	defer func() {
		if err := recover(); err != nil {
			if errs, ok := err.(error); ok {
				fmt.Println()
				if runtime.GOOS == "windows" && errs.Error() == "Incorrect function." {
					fmt.Println(ansi.Color(getWord("Kratos CLI error: CLI has not supported MINGW64 for now, "+
						"please use cmd terminal instead."), "red"))
				} else {
					fmt.Println(ansi.Color("Kratos CLI error: "+errs.Error(), "red"))

					if *verbose {
						fmt.Println(string(debug.Stack()))
					}
				}
				fmt.Println()
			}
		}
	}()

	app := cli.App("kratos", "Kratos CLI tool for developing and generating")

	app.Spec = "[-v]"

	verbose = app.BoolOpt("v verbose", false, "debug info output")
	// quiet

	app.Command("-V version", "display this application version", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			cliInfo()
		}
	})

	app.Command("generate", "generate table model files", func(cmd *cli.Cmd) {

		var (
			config = cmd.StringOpt("c config", "", "config ini path")
			lang   = cmd.StringOpt("l language", "en", "language")
			conn   = cmd.StringOpt("conn connection", "", "connection")
		)

		cmd.Action = func() {
			setDefaultLangSet(*lang)
			generating(*config, *conn)
		}
	})

	app.Command("init", "generate a template project", func(cmd *cli.Cmd) {

		var (
			config = cmd.StringOpt("c config", "", "config ini path")
			lang   = cmd.StringOpt("l language", "en", "language")
		)

		cmd.Action = func() {
			setDefaultLangSet(*lang)
			buildProject(*config)
		}
	})

	app.Command("add", "generate user/permission/roles", func(cmd *cli.Cmd) {

		cmd.Command("user", "generate users", func(cmd *cli.Cmd) {
			var (
				config = cmd.StringOpt("c config", "", "config ini path")
				lang   = cmd.StringOpt("l language", "en", "language")
			)

			cmd.Action = func() {
				setDefaultLangSet(*lang)
				addUser(*config)
			}
		})

		cmd.Command("permission", "generate permissions of table", func(cmd *cli.Cmd) {
			var (
				config = cmd.StringOpt("c config", "", "config ini path")
				lang   = cmd.StringOpt("l language", "en", "language")
			)

			cmd.Action = func() {
				setDefaultLangSet(*lang)
				addPermission(*config)
			}
		})
	})

	_ = app.Run(os.Args)
}
