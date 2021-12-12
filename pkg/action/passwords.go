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
	VaultTo         string       `json:"vaultTo,omitempty"`
	FolderTo        string       `json:"folderTo,omitempty"`
	Config          GlobalConfig `json:"-"`
}

type PasswordSearchQuery struct {
	Query   string       `json:"query"`
	VaultId string       `json:"vaultId,omitempty"`
	Colors  []int        `json:"colors,omitempty"`
	Tags    []string     `json:"tags,omitempty"`
	Config  GlobalConfig `json:"-"`
}

func (p *Password) PasswordAdd() error {
	data, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("cannot dump data: %s", err.Error())
	}

	url := p.Config.APIUrl + PASSWORDS_URI_PREFIX
	_, err = callAPI(url, "POST", p.Config.ReadToken(), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (p *Password) PasswordGet(id string) error {
	url := p.Config.APIUrl + PASSWORDS_URI_PREFIX + "/" + id
	_, err := callAPI(url, "GET", p.Config.ReadToken(), nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (p *Password) PasswordDelete(id string) error {
	url := p.Config.APIUrl + PASSWORDS_URI_PREFIX + "/" + id
	_, err := callAPI(url, "DELETE", p.Config.ReadToken(), nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (p *PasswordSearchQuery) PasswordSearch() error {
	data, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("cannot dump data: %s", err.Error())
	}

	url := p.Config.APIUrl + PASSWORDS_URI_PREFIX + "/search"
	_, err = callAPI(url, "POST", p.Config.ReadToken(), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (p *Password) PasswordEdit(id string) error {
	data, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("cannot dump data: %s", err.Error())
	}

	url := p.Config.APIUrl + PASSWORDS_URI_PREFIX + "/" + id
	_, err = callAPI(url, "PUT", p.Config.ReadToken(), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (p *Password) PasswordCopy(id string) error {
	data, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("cannot dump data: %s", err.Error())
	}

	url := p.Config.APIUrl + PASSWORDS_URI_PREFIX + "/" + id + "/copy"
	_, err = callAPI(url, "POST", p.Config.ReadToken(), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (p *Password) PasswordMove(id string) error {
	data, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("cannot dump data: %s", err.Error())
	}

	url := p.Config.APIUrl + PASSWORDS_URI_PREFIX + "/" + id + "/move"
	_, err = callAPI(url, "POST", p.Config.ReadToken(), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (p *Password) PasswordGetRecent() error {
	url := p.Config.APIUrl + PASSWORDS_URI_PREFIX + "/recent"
	_, err := callAPI(url, "GET", p.Config.ReadToken(), nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}
