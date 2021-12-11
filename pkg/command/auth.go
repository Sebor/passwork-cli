package command

import (
	"password-cli/pkg/action"

	"github.com/urfave/cli/v2"
)

var AuthCommands = cli.Command{
	Name:    "auth",
	Aliases: []string{"a"},
	Usage:   "Passwork auth section",
	Subcommands: []*cli.Command{
		{
			Name:    "login",
			Aliases: []string{"li"},
			Usage:   "Fetch session token by API key",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "api-key",
					Aliases:     []string{"ak"},
					Usage:       "Passwork API key",
					DefaultText: "$PASSWORK_API_KEY",
					EnvVars:     []string{"PASSWORK_API_KEY"},
					Required:    true,
				},
			},
			Action: func(c *cli.Context) error {
				a := action.Auth{
					APIkey: c.String("api-key"),
					Config: action.GlobalCongig{
						APIUrl:    c.String("api-url"),
						TokenFile: c.String("tokenfile"),
					},
				}
				err := a.AuthLogin()
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "logout",
			Aliases: []string{"lo"},
			Usage:   "Close session",
			Action: func(c *cli.Context) error {
				a := action.Auth{
					Config: action.GlobalCongig{
						APIUrl:    c.String("api-url"),
						TokenFile: c.String("tokenfile"),
					},
				}
				err := a.AuthLogout()
				if err != nil {
					return err
				}
				return nil
			},
		},
	},
}
