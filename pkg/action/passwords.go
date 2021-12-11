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
	Name            string       `json:"name,omitempty"`
	Login           string       `json:"login,omitempty"`
	CryptedPassword string       `json:"cryptedPassword,omitempty"`
	Description     string       `json:"description,omitempty"`
	Url             string       `json:"url,omitempty"`
	MasterHash      string       `json:"masterHash,omitempty"`
	FolderId        string       `json:"folderId,omitempty"`
	VaultId         string       `json:"vaultId,omitempty"`
	Color           int          `json:"color,omitempty"`
	Tags            []string     `json:"tags,omitempty"`
	Config          GlobalCongig `json:"-"`
}

type PasswordSearchQuery struct {
	Query   string       `json:"query"`
	VaultId string       `json:"vaultId,omitempty"`
	Colors  []int        `json:"colors,omitempty"`
	Tags    []string     `json:"tags,omitempty"`
	Config  GlobalCongig `json:"-"`
}

func (p *Password) PasswordAdd() error {
	token, err := getTokenFromFile(p.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	data, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("cannot dump data: %s", err.Error())
	}

	url := p.Config.APIUrl + PASSWORDS_URI_PREFIX
	_, err = callAPI(url, "POST", token, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (p *Password) PasswordGet(id string) error {
	token, err := getTokenFromFile(p.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	url := p.Config.APIUrl + PASSWORDS_URI_PREFIX + "/" + id
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (p *Password) PasswordDelete(id string) error {
	token, err := getTokenFromFile(p.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	url := p.Config.APIUrl + PASSWORDS_URI_PREFIX + "/" + id
	_, err = callAPI(url, "DELETE", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (p *PasswordSearchQuery) PasswordSearch() error {
	token, err := getTokenFromFile(p.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	data, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("ERROR: %s", err.Error())
	}

	url := p.Config.APIUrl + PASSWORDS_URI_PREFIX + "/search"
	_, err = callAPI(url, "POST", token, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (p *Password) PasswordEdit(id string) error {
	token, err := getTokenFromFile(p.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	data, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("ERROR: %s", err.Error())
	}

	url := p.Config.APIUrl + PASSWORDS_URI_PREFIX + "/" + id
	_, err = callAPI(url, "PUT", token, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}
