# Passwork CLI

**passwork-cli** - is a CLI tool to interact with [passwork](<https://passwork.me>) API.

There is no API public documentation but official js [client](<https://github.com/passwork-me/js-connector/tree/main/src/rest-modules>) is presented.

## Currently supported methods

* Auth
  * login
  * logout
* Folders
  * add
  * get
  * edit
  * delete
  * passwords
  * search
  * children
* Passwords
  * add
  * get
  * search
  * delete
  * edit
* Users
  * get
* Vaults
  * list
  * tags
  * domain
  * folders
  * passwords

## Build

```bash
go build cmd/passwork/passwork.go 
```

## Help

```bash
./passwork -h
```

## Auto-completion

### Bash

#### Enabling

To enable auto-completion for the current shell session, a bash script,
`autocomplete/bash_autocomplete` is included in this repo.

To use `autocomplete/bash_autocomplete` set an environment variable named `PROG` to
the name of your program and then `source` the `autocomplete/bash_autocomplete` file.

For example, if your cli program is called `passwork`:

`PROG=passwork source path/to/cli/autocomplete/bash_autocomplete`

Auto-completion is now enabled for the current shell, but will not persist into a new shell.

#### Distribution and Persistent Autocompletion

Copy `autocomplete/bash_autocomplete` into `/etc/bash_completion.d/` and rename
it to the name of the program you wish to add autocomplete support for (or
automatically install it there if you are distributing a package). Don't forget
to source the file or restart your shell to activate the auto-completion.

```bash
sudo cp path/to/autocomplete/bash_autocomplete /etc/bash_completion.d/passwork
source /etc/bash_completion.d/passwork
```

Alternatively, you can just document that users should `source` the generic
`autocomplete/bash_autocomplete` and set `$PROG` within their bash configuration
file, adding these lines:

```bash
PROG=passwork
source path/to/cli/autocomplete/bash_autocomplete
```

Keep in mind that if they are enabling auto-completion for more than one program,
they will need to set `PROG` and source `autocomplete/bash_autocomplete` for each
program, like so:

```bash
PROG=<program1>
source path/to/cli/autocomplete/bash_autocomplete
PROG=<program2>
source path/to/cli/autocomplete/bash_autocomplete
```

### ZSH

Auto-completion for ZSH is also supported using the `autocomplete/zsh_autocomplete`
file included in this repo. Two environment variables are used, `PROG` and `_CLI_ZSH_AUTOCOMPLETE_HACK`.
Set `PROG` to the program name as before, set `_CLI_ZSH_AUTOCOMPLETE_HACK` to `1`, and
then `source path/to/autocomplete/zsh_autocomplete`. Adding the following lines to your ZSH
configuration file (usually `.zshrc`) will allow the auto-completion to persist across new shells:

```zsh
PROG=passwork
_CLI_ZSH_AUTOCOMPLETE_HACK=1
source  path/to/autocomplete/zsh_autocomplete
```

### PowerShell

Auto-completion for PowerShell is also supported using the `autocomplete/powershell_autocomplete.ps1`
file included in this repo.

Rename the script to `<my program>.ps1` and move it anywhere in your file system.
The location of script does not matter, only the file name of the script has to match
the your program's binary name.

To activate it, enter `& path/to/autocomplete/<my program>.ps1`

To persist across new shells, open the PowerShell profile (with `code $profile` or `notepad $profile`)
and add the line:

```ps
& path/to/autocomplete/<my program>.ps1
```

## Roadmap

* Add remaining methods
* Add auto login feature when access token is expired
* Add github actions
* Add handling of http response code
