package main

import (
	"context"
	"fmt"
	"tlab/bootstrap/container"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sarulabs/di/v2"
	"github.com/urfave/cli/v3"
)

func Migration(ctn *di.Container) []*cli.Command {
	cmd := []*cli.Command{}
	cmd = append(cmd, &cli.Command{
		Name:  "migration",
		Usage: "Run current migration files",
		Commands: []*cli.Command{
			{
				Name:  "up",
				Usage: "run migration",
				Action: func(context.Context, *cli.Command) error {
					migrations := &migrate.FileMigrationSource{
						Dir: "./bootstrap/database/migration",
					}

					_, err := migrate.Exec(ctn.Get(container.DBDefName).(*sqlx.DB).DB, "postgres", migrations, migrate.Up)
					if err != nil {
						panic(err)
					}
					return nil
				},
			},
			{
				Name:  "down",
				Usage: "down known migration",
				Action: func(context.Context, *cli.Command) error {
					fmt.Println("down")
					migrations := &migrate.FileMigrationSource{
						Dir: "./bootstrap/database/migration",
					}

					_, err := migrate.Exec(ctn.Get(container.DBDefName).(*sqlx.DB).DB, "postgres", migrations, migrate.Down)
					if err != nil {
						panic(err)
					}
					return nil
				},
			},
			{
				Name:  "rollback",
				Usage: "rollback migration to specific migration",
				Commands: []*cli.Command{
					{
						Name:  "one-step",
						Usage: "run migration",
						Action: func(context.Context, *cli.Command) error {
							migrations := &migrate.FileMigrationSource{
								Dir: "./bootstrap/database/migration",
							}

							_, err := migrate.ExecMax(ctn.Get(container.DBDefName).(*sqlx.DB).DB, "postgres", migrations, migrate.Down, 1)
							if err != nil {
								panic(err)
							}
							return nil
						},
					},
				},
			},
		},
	})

	return cmd
}
