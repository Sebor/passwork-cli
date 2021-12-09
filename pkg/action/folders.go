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
	Name    string `json:"name"`
	VaultId string `json:"vaultId,omitempty"`
}

type FolderSearchQuery struct {
	Query   string `json:"query"`
	VaultId string `json:"vaultId,omitempty"`
}

func FolderAdd(url, tokenfile string, folder Folder) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	data, err := json.Marshal(folder)
	if err != nil {
		return fmt.Errorf("cannot dump data: %s", err.Error())
	}

	url = url + FOLDERS_URI_PREFIX
	_, err = callAPI(url, "POST", token, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func FolderGet(url, tokenfile, id string) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	url = url + FOLDERS_URI_PREFIX + "/" + id
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func FolderEdit(url, tokenfile, id, name string, folder Folder) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	data, err := json.Marshal(folder)
	if err != nil {
		return fmt.Errorf("cannot dump data: %s", err.Error())
	}

	url = url + FOLDERS_URI_PREFIX + "/" + id
	_, err = callAPI(url, "POST", token, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func FolderDelete(url, tokenfile, id string) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	url = url + VAULTS_URI_PREFIX + "/" + id
	_, err = callAPI(url, "DELETE", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func FolderGetPasswords(url, tokenfile, id string) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	url = url + FOLDERS_URI_PREFIX + "/" + id + "/passwords"
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func FolderSearch(url, tokenfile string, query FolderSearchQuery) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	data, err := json.Marshal(query)
	if err != nil {
		return fmt.Errorf("cannot dump data: %s", err.Error())
	}

	url = url + VAULTS_URI_PREFIX + "/search"
	_, err = callAPI(url, "POST", token, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func FolderGetChildren(url, tokenfile, id string) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	url = url + FOLDERS_URI_PREFIX + "/" + id + "/children"
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}
