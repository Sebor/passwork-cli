package command

import (
	"password-cli/pkg/action"

	"github.com/urfave/cli/v2"
)

var FolderCommands = cli.Command{
	Name:    "folders",
	Aliases: []string{"f"},
	Usage:   "Passwork folder section",
	Subcommands: []*cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Create new folder",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "name",
					Usage:    "Folder name",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "vault-id",
					Usage:    "Vault ID",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				folder := action.Folder{
					Name:    c.String("name"),
					VaultId: c.String("vault-id"),
				}
				err := action.FolderAdd(c.String("api-url"), c.String("tokenfile"), folder)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Get folder by ID",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "id",
					Usage:    "Folder ID",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				err := action.FolderGet(c.String("api-url"), c.String("tokenfile"), c.String("id"))
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "edit",
			Aliases: []string{"e"},
			Usage:   "Rename folder by ID",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "id",
					Usage:    "Folder ID",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "name",
					Value:    "name",
					Usage:    "Folder new name",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				folder := action.Folder{
					Name: c.String("name"),
				}
				err := action.FolderEdit(c.String("api-url"), c.String("tokenfile"), c.String("id"), c.String("name"), folder)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Delete folder by ID",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "id",
					Usage:    "Folder ID",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				err := action.FolderDelete(c.String("api-url"), c.String("tokenfile"), c.String("id"))
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "passwords",
			Aliases: []string{"p"},
			Usage:   "Get list of passwords in folder",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "id",
					Usage:    "Folder ID",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				err := action.FolderGetPasswords(c.String("api-url"), c.String("tokenfile"), c.String("id"))
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "Search folders by filter parameters",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "query",
					Usage:    "Query string",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "vault-id",
					Usage: "Vault ID",
				},
			},
			Action: func(c *cli.Context) error {
				query := action.FolderSearchQuery{
					Query:   c.String("query"),
					VaultId: c.String("vault-id"),
				}
				err := action.FolderSearch(c.String("api-url"), c.String("tokenfile"), query)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "children",
			Aliases: []string{"c"},
			Usage:   "Get list of children folders",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "id",
					Usage:    "Folder ID",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				err := action.FolderGetChildren(c.String("api-url"), c.String("tokenfile"), c.String("id"))
				if err != nil {
					return err
				}
				return nil
			},
		},
	},
}
