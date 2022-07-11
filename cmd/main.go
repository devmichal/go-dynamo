package main

import (
	"musement.com/commission-service/cmd/migrate"
	"musement.com/commission-service/cmd/server"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "server",
				Aliases:     []string{"s"},
				Description: "Run application to handle endpoints",
				Action:      server.Command,
			},
			{
				Name:        "migrate",
				Aliases:     []string{"m"},
				Description: "Execute database migrations",
				Action:      migrate.Command,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
