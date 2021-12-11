package action

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	AUTH_URI_PREFIX = "/auth"
)

type Auth struct {
	APIkey string
	Config GlobalCongig
}

func (a *Auth) AuthLogin() error {
	url := a.Config.APIUrl + AUTH_URI_PREFIX + "/" + "login/" + a.APIkey
	body, err := callAPI(url, "POST", "", nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot get user home dir: %s", err.Error())
	}

	tokenFilepath := filepath.Join(homeDir, a.Config.TokenFile)
	err = os.WriteFile(tokenFilepath, body, 0600)
	if err != nil {
		return fmt.Errorf("cannot create token file %s: %s", tokenFilepath, err.Error())
	}

	return nil
}

func (a *Auth) AuthLogout() error {
	token, err := getTokenFromFile(a.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	url := a.Config.APIUrl + AUTH_URI_PREFIX + "/" + "logout"
	_, err = callAPI(url, "POST", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}
