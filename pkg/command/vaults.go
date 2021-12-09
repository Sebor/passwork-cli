package command

import (
	"fmt"
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
				err := action.VaultList(c.String("api-url"), c.String("tokenfile"))
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
				err := action.VaultTags(c.String("api-url"), c.String("tokenfile"))
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
				err := action.VaultDomain(c.String("api-url"), c.String("tokenfile"))
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:      "folders",
			Aliases:   []string{"f"},
			Usage:     "Get folders by VAULT_ID argument",
			ArgsUsage: "VAULT_ID",
			Action: func(c *cli.Context) error {
				if c.Args().Len() != 1 {
					return fmt.Errorf("ERROR: %s", "You must provide VAULT_ID argument")
				}
				err := action.VaultFolders(c.String("api-url"), c.String("tokenfile"), c.Args().Get(0))
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:      "passwords",
			Aliases:   []string{"p"},
			Usage:     "Get passwords in vault root by VAULT_ID argument",
			ArgsUsage: "VAULT_ID",
			Action: func(c *cli.Context) error {
				if c.Args().Len() != 1 {
					return fmt.Errorf("ERROR: %s", "You must provide VAULT_ID argument")
				}
				err := action.VaultPasswords(c.String("api-url"), c.String("tokenfile"), c.Args().Get(0))
				if err != nil {
					return err
				}
				return nil
			},
		},
	},
}
