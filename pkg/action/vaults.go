package action

import (
	"fmt"
)

var (
	VAULTS_URI_PREFIX = "/vaults"
)

type Vault struct {
	Name            string `json:"name"`
	DomainId        string `json:"domainId"`
	MPCrypted       string `json:"mpCrypted"`
	PasswordHash    string `json:"passwordHash"`
	Salt            string `json:"salt"`
	PasswordCrypted string `json:"passwordCrypted"`
}

func VaultList(url, tokenfile string) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	url = url + VAULTS_URI_PREFIX + "/" + "list"
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func VaultTags(url, tokenfile string) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	url = url + VAULTS_URI_PREFIX + "/" + "tags"
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func VaultDomain(url, tokenfile string) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	url = url + VAULTS_URI_PREFIX + "/" + "domain"
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func VaultFolders(url, tokenfile, vaultID string) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	url = url + VAULTS_URI_PREFIX + "/" + vaultID + "/folders"
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}

func VaultPasswords(url, tokenfile, vaultID string) error {
	token, err := getTokenFromFile(tokenfile)
	if err != nil {
		return fmt.Errorf("cannot get get token: %s", err.Error())
	}

	url = url + VAULTS_URI_PREFIX + "/" + vaultID + "/passwords"
	_, err = callAPI(url, "GET", token, nil)
	if err != nil {
		return fmt.Errorf("cannot call API: %s", err.Error())
	}

	return nil
}
