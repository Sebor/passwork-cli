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
	Config          GlobalConfig `json:"-"`
}

func (v *Vault) VaultList() error {
	url := v.Config.APIUrl + VAULTS_URI_PREFIX + "/list"
	_, err := callAPI(url, "GET", v.Config.ReadToken(), nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (v *Vault) VaultAllTags() error {
	url := v.Config.APIUrl + VAULTS_URI_PREFIX + "/tags"
	_, err := callAPI(url, "GET", v.Config.ReadToken(), nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (v *Vault) VaultDomain() error {
	url := v.Config.APIUrl + VAULTS_URI_PREFIX + "/domain"
	_, err := callAPI(url, "GET", v.Config.ReadToken(), nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (v *Vault) VaultGetFolders(id string) error {
	url := v.Config.APIUrl + VAULTS_URI_PREFIX + "/" + id + "/folders"
	_, err := callAPI(url, "GET", v.Config.ReadToken(), nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (v *Vault) VaultGetPasswords(id string) error {
	url := v.Config.APIUrl + VAULTS_URI_PREFIX + "/" + id + "/passwords"
	_, err := callAPI(url, "GET", v.Config.ReadToken(), nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (v *Vault) VaultGetFullInfo(id string) error {
	url := v.Config.APIUrl + VAULTS_URI_PREFIX + "/" + id + "/fullInfo"
	_, err := callAPI(url, "GET", v.Config.ReadToken(), nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func (v *Vault) VaultGetTags(id string) error {
	url := v.Config.APIUrl + VAULTS_URI_PREFIX + "/" + id + "/tags"
	_, err := callAPI(url, "GET", v.Config.ReadToken(), nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}
