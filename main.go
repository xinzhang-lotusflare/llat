package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

const VERSION = "0.1"

var verbose = false

func run(cCtx *cli.Context) error {
	time := min(cCtx.Int("time"), 12)
	if allowPrintInfo() {
		fmt.Println("It will last for " + strconv.Itoa(time) + " hours")
	}

	if err := switchInterface(true); err != nil {
		return err
	}
	fmt.Println("Running")
	return nil
}

func stop(cCtx *cli.Context) error {
	if err := switchInterface(false); err != nil {
		return err
	}
	fmt.Println("Stop")
	return nil
}

func get_version(cCtx *cli.Context) error {
	fmt.Println("llat version", VERSION)
	return nil
}

func main() {
	if isReleaseMode {
		log.SetOutput(io.Discard)
	}

	app := &cli.App{
		Name:  "llat",
		Usage: "Guide you a way for LF VPN",
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "Start to use LF VPN",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "verbose",
						Aliases: []string{"v"},
						Value:   false,
						Usage:   "If enabled, llat will print more logs",
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
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "verbose",
						Aliases: []string{"v"},
						Value:   false,
						Usage:   "If enabled, llat will print more logs",
						Action: func(cCtx *cli.Context, value bool) error {
							verbose = value
							return nil
						},
					},
				},
				Action: stop,
			},
			{
				Name:  "install",
				Usage: "Prepare workspace: ~/.llat",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "bash-executable",
						Aliases: []string{"bash", "b"},
						Value:   "",
						Usage:   "The bash executable for wg-quick, required version: > 4.0.0",
					},
					&cli.StringFlag{
						Name:    "llat-config",
						Aliases: []string{"config", "c"},
						Value:   "",
						Usage:   "The config file provided by LLat admin",
					},
				},
				Action: install,
			},
			{
				Name:   "version",
				Usage:  "Check current llat version",
				Action: get_version,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Println(err.Error())
		fmt.Println("exit abnormally")
		os.Exit(1)
	}
}
