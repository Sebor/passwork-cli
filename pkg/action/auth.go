package action

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	AUTH_URI_PREFIX = "/auth"
)

func AuthLogin(url, key, tokenfile string) error {
	url = url + AUTH_URI_PREFIX + "/" + "login/" + key
	body, err := callAPI(url, "POST", "", nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot get user home dir: %s", err.Error())
	}

	tokenFilepath := filepath.Join(homeDir, tokenfile)
	err = os.WriteFile(tokenFilepath, body, 0600)
	if err != nil {
		return fmt.Errorf("cannot create token file %s: %s", tokenFilepath, err.Error())
	}

	return nil
}

func AuthLogout(url, tokenfile string) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	url = url + AUTH_URI_PREFIX + "/" + "logout"
	_, err = callAPI(url, "POST", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}
