package main

import (
	"os"

	"github.com/c0rby/shoppinglist/pkg/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	app := &cli.App{
		Name:  "shoppinglist",
		Usage: "A simple shopping list api and web app",
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "Start the shopping list server",
				Action: func(c *cli.Context) error {
					s := server.New()
					return s.ListenAndServe()
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to run the app")
	}
}
