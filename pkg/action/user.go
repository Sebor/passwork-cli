package action

import (
	"fmt"
)

var (
	USER_URI_PREFIX = "/user"
)

type User struct {
	Config GlobalCongig
}

func (u *User) UserGet() error {
	token, err := getTokenFromFile(u.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	url := u.Config.APIUrl + USER_URI_PREFIX + "/info"
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}
