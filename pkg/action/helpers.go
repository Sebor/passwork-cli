package action

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type APIToken struct {
	Status string    `json:"status"`
	Data   TokenData `json:"data"`
}

type TokenData struct {
	Token          string `json:"token"`
	TokenExpiredAt int64  `json:"tokenExpiredAt"`
}

type GlobalCongig struct {
	APIUrl    string
	TokenFile string
}

func callAPI(url, method, token string, data io.Reader) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %s", err.Error())
	}

	req.Header.Add("Passwork-Auth", token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot do request: %s", err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read responce body: %s", err.Error())
	}

	fmt.Println(string(body))

	return body, nil
}

func getTokenFromFile(tokenfile string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot get user home directory: %s", err.Error())
	}

	tokenFilepath := filepath.Join(homeDir, tokenfile)
	data, err := os.ReadFile(tokenFilepath)
	if err != nil {
		return "", fmt.Errorf("cannot read tokenfile: %s", err.Error())
	}

	token := APIToken{}
	err = json.Unmarshal(data, &token)
	if err != nil {
		return "", fmt.Errorf("cannot unmarshal token data: %s", err.Error())
	}

	return token.Data.Token, nil
}
