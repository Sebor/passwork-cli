package action

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var (
	FOLDERS_URI_PREFIX = "/folders"
)

type Folder struct {
	Name    string       `json:"name"`
	VaultId string       `json:"vaultId,omitempty"`
	Config  GlobalCongig `json:"-"`
}

type FolderSearchQuery struct {
	Query   string       `json:"query"`
	VaultId string       `json:"vaultId,omitempty"`
	Config  GlobalCongig `json:"-"`
}

func (f *Folder) FolderAdd() error {
	token, err := getTokenFromFile(f.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	data, err := json.Marshal(f)
	if err != nil {
		return fmt.Errorf("cannot dump data: %s", err.Error())
	}

	url := f.Config.APIUrl + FOLDERS_URI_PREFIX
	_, err = callAPI(url, "POST", token, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (f *Folder) FolderGet(id string) error {
	token, err := getTokenFromFile(f.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	url := f.Config.APIUrl + FOLDERS_URI_PREFIX + "/" + id
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (f *Folder) FolderEdit(id string) error {
	token, err := getTokenFromFile(f.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}
	data, err := json.Marshal(f)
	if err != nil {
		return fmt.Errorf("cannot dump data: %s", err.Error())
	}

	url := f.Config.APIUrl + FOLDERS_URI_PREFIX + "/" + id
	_, err = callAPI(url, "PUT", token, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (f *Folder) FolderDelete(id string) error {
	token, err := getTokenFromFile(f.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	url := f.Config.APIUrl + FOLDERS_URI_PREFIX + "/" + id
	_, err = callAPI(url, "DELETE", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (f *Folder) FolderGetPasswords(id string) error {
	token, err := getTokenFromFile(f.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	url := f.Config.APIUrl + FOLDERS_URI_PREFIX + "/" + id + "/passwords"
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (f *FolderSearchQuery) FolderSearch() error {
	token, err := getTokenFromFile(f.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	data, err := json.Marshal(f)
	if err != nil {
		return fmt.Errorf("cannot dump data: %s", err.Error())
	}

	url := f.Config.APIUrl + FOLDERS_URI_PREFIX + "/search"
	_, err = callAPI(url, "POST", token, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (f *Folder) FolderGetChildren(id string) error {
	token, err := getTokenFromFile(f.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	url := f.Config.APIUrl + FOLDERS_URI_PREFIX + "/" + id + "/children"
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}
