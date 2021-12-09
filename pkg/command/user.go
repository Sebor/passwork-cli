package command

import (
	"password-cli/pkg/action"

	"github.com/urfave/cli/v2"
)

var UserCommands = cli.Command{
	Name:    "user",
	Aliases: []string{"u"},
	Usage:   "Passwork user section",
	Subcommands: []*cli.Command{
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Get user info",
			Action: func(c *cli.Context) error {
				err := action.UserGet(c.String("api-url"), c.String("tokenfile"))
				if err != nil {
					return err
				}
				return nil
			},
		},
	},
}
