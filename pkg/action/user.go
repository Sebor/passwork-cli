package action

import (
	"fmt"
)

var (
	USER_URI_PREFIX = "/user"
)

type User struct {
	Config GlobalConfig
}

func (u *User) UserGet() error {
	url := u.Config.APIUrl + USER_URI_PREFIX + "/info"
	_, err := callAPI(url, "GET", u.Config.ReadToken(), nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}
