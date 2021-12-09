package command

import (
	"encoding/base64"
	"password-cli/pkg/action"

	"github.com/urfave/cli/v2"
)

var PasswordCommands = cli.Command{
	Name:    "password",
	Aliases: []string{"p"},
	Usage:   "Passwork password section",
	Subcommands: []*cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Create new password",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "name",
					Usage:    "Password name",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "login",
					Usage:    "Password login",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "description",
					Usage: "Password description",
				},
				&cli.StringFlag{
					Name:     "password",
					Usage:    "Password",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "vault-id",
					Usage:    "Password Vault ID",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "folder-id",
					Usage:    "Password Folder ID",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "url",
					Usage: "Password URL",
				},
				&cli.StringSliceFlag{
					Name:  "tags",
					Usage: "Password Tags",
				},
				&cli.StringFlag{
					Name:  "master-hash",
					Usage: "Password Master Hash",
				},
				&cli.IntFlag{
					Name:  "color",
					Usage: "Password Color",
				},
			},
			Action: func(c *cli.Context) error {
				password := action.Password{
					Name:            c.String("name"),
					Login:           c.String("login"),
					CryptedPassword: base64.StdEncoding.EncodeToString([]byte(c.String("password"))),
					Description:     c.String("description"),
					Url:             c.String("url"),
					FolderId:        c.String("folder-id"),
					VaultId:         c.String("vault-id"),
					MasterHash:      c.String("master-hash"),
					Color:           c.Int("color"),
					Tags:            c.StringSlice("tags"),
				}
				err := action.PasswordAdd(c.String("api-url"), c.String("tokenfile"), password)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "Search passwords by filter parameters",
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
				&cli.IntSliceFlag{
					Name:  "colors",
					Usage: "Colors int slice",
				},
				&cli.StringSliceFlag{
					Name:  "tags",
					Usage: "Tags string slice",
				},
			},
			Action: func(c *cli.Context) error {
				query := action.PasswordSearchQuery{
					Query:   c.String("query"),
					VaultId: c.String("vault-id"),
					Colors:  c.IntSlice("colors"),
					Tags:    c.StringSlice("tags"),
				}
				err := action.PasswordSearch(c.String("api-url"), c.String("tokenfile"), query)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Get password by ID",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "id",
					Usage:    "Password ID",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				err := action.PasswordGet(c.String("api-url"), c.String("tokenfile"), c.String("id"))
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Delete password by ID",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "id",
					Usage:    "Password ID",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				err := action.PasswordDelete(c.String("api-url"), c.String("tokenfile"), c.String("id"))
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "edit",
			Aliases: []string{"e"},
			Usage:   "Edit password by ID",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "id",
					Usage:    "Password ID",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "name",
					Usage: "Password name",
				},
				&cli.StringFlag{
					Name:  "login",
					Usage: "Password login",
				},
				&cli.StringFlag{
					Name:  "description",
					Usage: "Password description",
				},
				&cli.StringFlag{
					Name:  "password",
					Usage: "Password",
				},
				&cli.StringFlag{
					Name:  "vault-id",
					Usage: "Password Vault ID",
				},
				&cli.StringFlag{
					Name:  "folder-id",
					Usage: "Password Folder ID",
				},
				&cli.StringFlag{
					Name:  "url",
					Usage: "Password URL",
				},
				&cli.StringSliceFlag{
					Name:  "tags",
					Usage: "Password Tags",
				},
				&cli.StringFlag{
					Name:  "master-hash",
					Usage: "Password Master Hash",
				},
				&cli.IntFlag{
					Name:  "color",
					Usage: "Password Color",
				},
			},
			Action: func(c *cli.Context) error {
				password := action.Password{
					Name:            c.String("name"),
					Login:           c.String("login"),
					CryptedPassword: base64.StdEncoding.EncodeToString([]byte(c.String("password"))),
					Description:     c.String("description"),
					Url:             c.String("url"),
					FolderId:        c.String("folder-id"),
					VaultId:         c.String("vault-id"),
					MasterHash:      c.String("master-hash"),
					Color:           c.Int("color"),
					Tags:            c.StringSlice("tags"),
				}
				err := action.PasswordEdit(c.String("api-url"), c.String("tokenfile"), c.String("id"), password)
				if err != nil {
					return err
				}
				return nil
			},
		},
	},
}
