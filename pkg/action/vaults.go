package action

import (
	"fmt"
)

var (
	VAULTS_URI_PREFIX = "/vaults"
)

type Vault struct {
	Name            string       `json:"name"`
	DomainId        string       `json:"domainId"`
	MPCrypted       string       `json:"mpCrypted"`
	PasswordHash    string       `json:"passwordHash"`
	Salt            string       `json:"salt"`
	PasswordCrypted string       `json:"passwordCrypted"`
	Config          GlobalCongig `json:"-"`
}

func (v *Vault) VaultList() error {
	token, err := getTokenFromFile(v.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	url := v.Config.APIUrl + VAULTS_URI_PREFIX + "/" + "list"
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (v *Vault) VaultTags() error {
	token, err := getTokenFromFile(v.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	url := v.Config.APIUrl + VAULTS_URI_PREFIX + "/" + "tags"
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (v *Vault) VaultDomain() error {
	token, err := getTokenFromFile(v.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	url := v.Config.APIUrl + VAULTS_URI_PREFIX + "/" + "domain"
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (v *Vault) VaultGetFolders(id string) error {
	token, err := getTokenFromFile(v.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	url := v.Config.APIUrl + VAULTS_URI_PREFIX + "/" + id + "/folders"
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (v *Vault) VaultGetPasswords(id string) error {
	token, err := getTokenFromFile(v.Config.TokenFile)
	if err != nil {
		return fmt.Errorf("cannot get token: %s", err.Error())
	}

	url := v.Config.APIUrl + VAULTS_URI_PREFIX + "/" + id + "/passwords"
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}
