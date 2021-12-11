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
				a := action.Folder{
					Name:    c.String("name"),
					VaultId: c.String("vault-id"),
					Config: action.GlobalCongig{
						APIUrl:    c.String("api-url"),
						TokenFile: c.String("tokenfile"),
					},
				}
				err := a.FolderAdd()
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
				a := action.Folder{
					Config: action.GlobalCongig{
						APIUrl:    c.String("api-url"),
						TokenFile: c.String("tokenfile"),
					},
				}
				err := a.FolderGet(c.String("id"))
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
					Usage:    "Folder new name",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				a := action.Folder{
					Name: c.String("name"),
					Config: action.GlobalCongig{
						APIUrl:    c.String("api-url"),
						TokenFile: c.String("tokenfile"),
					},
				}
				err := a.FolderEdit(c.String("id"))
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
				a := action.Folder{
					Config: action.GlobalCongig{
						APIUrl:    c.String("api-url"),
						TokenFile: c.String("tokenfile"),
					},
				}
				err := a.FolderDelete(c.String("id"))
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
				a := action.Folder{
					Config: action.GlobalCongig{
						APIUrl:    c.String("api-url"),
						TokenFile: c.String("tokenfile"),
					},
				}
				err := a.FolderGetPasswords(c.String("id"))
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
				a := action.FolderSearchQuery{
					Query:   c.String("query"),
					VaultId: c.String("vault-id"),
					Config: action.GlobalCongig{
						APIUrl:    c.String("api-url"),
						TokenFile: c.String("tokenfile"),
					},
				}
				err := a.FolderSearch()
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
				a := action.Folder{
					Config: action.GlobalCongig{
						APIUrl:    c.String("api-url"),
						TokenFile: c.String("tokenfile"),
					},
				}
				err := a.FolderGetChildren(c.String("id"))
				if err != nil {
					return err
				}
				return nil
			},
		},
	},
}
