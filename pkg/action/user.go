package action

import (
	"fmt"
)

var (
	USER_URI_PREFIX = "/user"
)

func UserGet(url, tokenfile string) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	url = url + USER_URI_PREFIX + "/info"
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}
