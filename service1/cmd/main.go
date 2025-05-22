package main

import (
	"local/service1/internal/pkg/containers"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "Service",
		Version: "1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
				EnvVars: []string{"ACCOUNT_CONFIG_PATH"},
			},
		},
		Action: runApp,
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func runApp(ctx *cli.Context) error {
	configPath := ctx.String("config")
	container, err := containers.NewContainer(configPath)
	if err != nil {
		return err
	}
	err = container.Start(ctx.Context)
	if err != nil {
		return err
	}
	return nil
}