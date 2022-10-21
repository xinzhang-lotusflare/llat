package main

import (
	"fmt"
	"log"
	"os"
	"io"
	"strconv"

	"github.com/urfave/cli/v2"
)

var verbose = false

func createVirtualInterface(time int) error {
	switchInterface(true)
	fmt.Println("Running")
	return nil
}

func run(cCtx *cli.Context) error {
	time := min(cCtx.Int("time"), 12)
	if allowPrintInfo() {
		fmt.Println("It will last for " + strconv.Itoa(time) + " hours")
	}

	return createVirtualInterface(time)
}

func stop(cCtx *cli.Context) error {
	switchInterface(false)
	fmt.Println("Stop")
	return nil
}

func main() {
	if isReleaseMode {
		log.SetOutput(io.Discard)
	}

	app := &cli.App{
		Name:  "bulldozer",
		Usage: "Build you way to LF infra world",
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "Show me the LF VPN",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "time",
						Aliases: []string{"t"},
						Value:   9,
						Usage:   "Hours that the LF VPN will be available for",
					},
					&cli.BoolFlag{
						Name:    "verbose",
						Aliases: []string{"v"},
						Value:   false,
						Usage:   "If enabled, bulldozer will print more logs",
						Action: func(cCtx *cli.Context, value bool) error {
							verbose = value
							return nil
						},
					},
				},
				Action: run,
			},
			{
				Name:    "stop",
				Aliases: []string{"s"},
				Usage:   "No need for LF VPN any more",
				Action:  stop,
			},
			{
				Name:   "install",
				Usage:  "Use homebrew to install required dependencies",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "bash-executable",
						Aliases: []string{ "bash", "b" },
						Value:   "",
						Usage:   "The bash executable for wg-quick, required version: > 4.0.0",
					},
				},
				Action: install,
			},
			{
				Name:   "upgrade",
				Usage:  "Upgrade the bulldozer",
				Action: upgrade,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Println(err.Error())
		fmt.Println("exit abnormally")
		os.Exit(1)
	}
}
