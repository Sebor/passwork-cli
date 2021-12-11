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
				a := action.User{
					Config: action.GlobalCongig{
						APIUrl:    c.String("api-url"),
						TokenFile: c.String("tokenfile"),
					},
				}
				err := a.UserGet()
				if err != nil {
					return err
				}
				return nil
			},
		},
	},
}
