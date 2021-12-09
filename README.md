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

## Roadmap

* Add remaining methods
* Add auto login feature when access token is expired
