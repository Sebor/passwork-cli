package command

import (
	"password-cli/pkg/action"

	"github.com/urfave/cli/v2"
)

var VaultCommands = cli.Command{
	Name:    "vault",
	Aliases: []string{"v"},
	Usage:   "Passwork vault section",
	Subcommands: []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "Get list of user vaults",
			Action: func(c *cli.Context) error {
				a := action.Vault{
					Config: action.GlobalConfig{
						APIUrl:    c.String("api-url"),
						TokenFile: c.String("tokenfile"),
					},
				}
				err := a.VaultList()
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "tags",
			Aliases: []string{"t"},
			Usage:   "Get all tags used in vault",
			Action: func(c *cli.Context) error {
				a := action.Vault{
					Config: action.GlobalConfig{
						APIUrl:    c.String("api-url"),
						TokenFile: c.String("tokenfile"),
					},
				}
				err := a.VaultTags()
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "domain",
			Aliases: []string{"d"},
			Usage:   "Get domain info",
			Action: func(c *cli.Context) error {
				a := action.Vault{
					Config: action.GlobalConfig{
						APIUrl:    c.String("api-url"),
						TokenFile: c.String("tokenfile"),
					},
				}
				err := a.VaultDomain()
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "folders",
			Aliases: []string{"f"},
			Usage:   "Get folders by vault id",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "id",
					Usage:    "Vault ID",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				a := action.Vault{
					Config: action.GlobalConfig{
						APIUrl:    c.String("api-url"),
						TokenFile: c.String("tokenfile"),
					},
				}
				err := a.VaultGetFolders(c.String("id"))
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "passwords",
			Aliases: []string{"p"},
			Usage:   "Get passwords in vault root",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "id",
					Usage:    "Vault ID",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				a := action.Vault{
					Config: action.GlobalConfig{
						APIUrl:    c.String("api-url"),
						TokenFile: c.String("tokenfile"),
					},
				}
				err := a.VaultGetPasswords(c.String("id"))
				if err != nil {
					return err
				}
				return nil
			},
		},
	},
}
