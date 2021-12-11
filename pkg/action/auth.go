package action

import (
	"fmt"
)

var (
	AUTH_URI_PREFIX = "/auth"
)

type Auth struct {
	APIkey string
	Config GlobalConfig
}

func (a *Auth) AuthLogin() error {
	url := a.Config.APIUrl + AUTH_URI_PREFIX + "/" + "login/" + a.APIkey
	body, err := callAPI(url, "POST", "", nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	a.Config.WriteToken(body)

	return nil
}

func (a *Auth) AuthLogout() error {
	url := a.Config.APIUrl + AUTH_URI_PREFIX + "/" + "logout"
	_, err := callAPI(url, "POST", a.Config.ReadToken(), nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}
