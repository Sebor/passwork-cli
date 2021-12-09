package action

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var (
	PASSWORDS_URI_PREFIX = "/passwords"
)

type Password struct {
	Name            string   `json:"name,omitempty"`
	Login           string   `json:"login,omitempty"`
	CryptedPassword string   `json:"cryptedPassword,omitempty"`
	Description     string   `json:"description,omitempty"`
	Url             string   `json:"url,omitempty"`
	MasterHash      string   `json:"masterHash,omitempty"`
	FolderId        string   `json:"folderId,omitempty"`
	VaultId         string   `json:"vaultId,omitempty"`
	Color           int      `json:"color,omitempty"`
	Tags            []string `json:"tags,omitempty"`
}

type PasswordSearchQuery struct {
	Query   string   `json:"query"`
	VaultId string   `json:"vaultId,omitempty"`
	Colors  []int    `json:"colors,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}

func PasswordAdd(url, tokenfile string, password Password) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	data, err := json.Marshal(password)
	if err != nil {
		return fmt.Errorf("cannot dump data: %s", err.Error())
	}

	url = url + PASSWORDS_URI_PREFIX
	_, err = callAPI(url, "POST", token, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func PasswordGet(url, tokenfile, id string) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	url = url + PASSWORDS_URI_PREFIX + "/" + id
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func PasswordDelete(url, tokenfile, id string) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	url = url + PASSWORDS_URI_PREFIX + "/" + id
	_, err = callAPI(url, "DELETE", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func PasswordSearch(url, tokenfile string, query PasswordSearchQuery) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	data, err := json.Marshal(query)
	if err != nil {
		return fmt.Errorf("ERROR: %s", err.Error())
	}

	url = url + PASSWORDS_URI_PREFIX + "/search"
	_, err = callAPI(url, "POST", token, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func PasswordEdit(url, tokenfile, id string, password Password) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	data, err := json.Marshal(password)
	if err != nil {
		return fmt.Errorf("ERROR: %s", err.Error())
	}

	url = url + PASSWORDS_URI_PREFIX + "/" + id
	_, err = callAPI(url, "PUT", token, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}
