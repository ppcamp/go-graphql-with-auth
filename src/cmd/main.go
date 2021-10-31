package main

import (
	"os"

	"github.com/ppcamp/go-graphql-with-auth/internal/config"
	postgres "github.com/ppcamp/go-graphql-with-auth/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-graphql-user"
	app.Usage = "Build the file and execute it and then, make some graphql calls"
	app.Flags = config.Flags
	app.Action = run
	app.Run(os.Args)
}

func run(c *cli.Context) error {
	config.Setup()

	// if config.App.Migrate {
	// 	migrate := migrations.SetupMigrations(c.App.Name, config.Database.Url)
	// 	err := migrate.Up()
	// 	if err != nil {
	// 		logrus.WithError(err).Fatal("failed to migrate the data")
	// 	}
	// }

	storage, err := postgres.NewStorage()
	if err != nil {
		logrus.Fatal("couldn't connect to databaseql")
	}

	r := SetupEngine(storage)
	r.Run(config.App.Address)
	return nil
}
