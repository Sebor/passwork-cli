package main

import (
	"fmt"
	"os"

	"password-cli/pkg/command"

	"github.com/urfave/cli/v2"
)

var (
	AppUsage   = "Passwork Command Line Interface tool (https://passwork.me)"
	AppAuthors = []*cli.Author{
		{
			Name:  "Sebor",
			Email: "sebor@sebor.pro",
		},
	}
	AppVersion      = "v1.0.0"
	AppHelpTemplate = `NAME:
{{.Name}} - {{.Usage}}

USAGE:
{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
{{if len .Authors}}
AUTHOR:
{{range .Authors}}{{ . }}{{end}}
{{end}}
{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
{{range .VisibleFlags}}{{.}}
{{end}}{{end}}
{{if .Copyright }}
COPYRIGHT:
{{.Copyright}}
{{end}}
{{if .Version}}
VERSION:
{{.Version}}
{{end}}
`
)

func main() {
	cli.AppHelpTemplate = AppHelpTemplate

	// Global Flags
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "api-url",
			Aliases:  []string{"au"},
			Usage:    "Passwork API URL",
			EnvVars:  []string{"PASSWORK_API_URL"},
			Required: true,
		},
		&cli.StringFlag{
			Name:    "tokenfile",
			Aliases: []string{"tf"},
			Value:   ".passwork-token",
			Usage:   "Filename inside $HOME/ dir to store API token",
			EnvVars: []string{"PASSWORK_TOKEN_FILENAME"},
		},
	}

	app := &cli.App{
		Usage:                AppUsage,
		Version:              AppVersion,
		Authors:              AppAuthors,
		Flags:                flags,
		Commands:             command.Commands,
		EnableBashCompletion: true,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
