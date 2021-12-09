package command

import (
	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	&AuthCommands,
	&VaultCommands,
	&PasswordCommands,
	&UserCommands,
	&FolderCommands,
}
