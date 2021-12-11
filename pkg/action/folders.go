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
	Config  GlobalConfig `json:"-"`
}

type FolderSearchQuery struct {
	Query   string       `json:"query"`
	VaultId string       `json:"vaultId,omitempty"`
	Config  GlobalConfig `json:"-"`
}

func (f *Folder) FolderAdd() error {
	data, err := json.Marshal(f)
	if err != nil {
		return fmt.Errorf("cannot dump data: %s", err.Error())
	}

	url := f.Config.APIUrl + FOLDERS_URI_PREFIX
	_, err = callAPI(url, "POST", f.Config.ReadToken(), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (f *Folder) FolderGet(id string) error {
	url := f.Config.APIUrl + FOLDERS_URI_PREFIX + "/" + id
	_, err := callAPI(url, "GET", f.Config.ReadToken(), nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (f *Folder) FolderEdit(id string) error {
	data, err := json.Marshal(f)
	if err != nil {
		return fmt.Errorf("cannot dump data: %s", err.Error())
	}

	url := f.Config.APIUrl + FOLDERS_URI_PREFIX + "/" + id
	_, err = callAPI(url, "PUT", f.Config.ReadToken(), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (f *Folder) FolderDelete(id string) error {
	url := f.Config.APIUrl + FOLDERS_URI_PREFIX + "/" + id
	_, err := callAPI(url, "DELETE", f.Config.ReadToken(), nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (f *Folder) FolderGetPasswords(id string) error {
	url := f.Config.APIUrl + FOLDERS_URI_PREFIX + "/" + id + "/passwords"
	_, err := callAPI(url, "GET", f.Config.ReadToken(), nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (f *FolderSearchQuery) FolderSearch() error {
	data, err := json.Marshal(f)
	if err != nil {
		return fmt.Errorf("cannot dump data: %s", err.Error())
	}

	url := f.Config.APIUrl + FOLDERS_URI_PREFIX + "/search"
	_, err = callAPI(url, "POST", f.Config.ReadToken(), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (f *Folder) FolderGetChildren(id string) error {
	url := f.Config.APIUrl + FOLDERS_URI_PREFIX + "/" + id + "/children"
	_, err := callAPI(url, "GET", f.Config.ReadToken(), nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}
