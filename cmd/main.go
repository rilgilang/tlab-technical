package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/urfave/cli/v3"
	"log"
	"os"
	"tlab/bootstrap/config"
	"tlab/bootstrap/container"
	"tlab/src/infrastructure/http"
)

func main() {
	ctn, err := container.NewContainer()
	if err != nil {
		panic("cannot initialize container: " + err.Error())
	}

	cmd := cli.Command{}

	cmd.Commands = append(cmd.Commands,
		&cli.Command{
			Name:  "run",
			Usage: "Run GUS http",
			Action: func(context.Context, *cli.Command) error {
				// Initialize container

				// Get echo instance
				e := ctn.Get(container.EchoDefName).(*echo.Echo)

				http.SetupRoutes(e, &ctn)

				// Get config and start server
				cfg := ctn.Get(container.ConfigDefName).(config.Config)
				port := fmt.Sprintf(":%s", cfg.ServerPort)
				if err := e.Start(port); err != nil {
					log.Fatal("Cannot start server:", err)
				}
				return nil
			},
		},
	)

	cmd.Commands = append(cmd.Commands, Migration(&ctn)...)

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
