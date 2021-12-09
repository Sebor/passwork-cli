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

func callAPI(url, method, token string, data io.Reader) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, fmt.Errorf("ERROR: %s", err.Error())
	}

	req.Header.Add("Passwork-Auth", token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ERROR: %s", err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ERROR: %s", err.Error())
	}

	fmt.Println(string(body))

	return body, nil
}

func getTokenFromFile(tokenfile string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("ERROR: %s", err.Error())
	}

	tokenFilepath := filepath.Join(homeDir, tokenfile)
	data, err := os.ReadFile(tokenFilepath)
	if err != nil {
		return "", fmt.Errorf("ERROR: %s", err.Error())
	}

	token := APIToken{}
	err = json.Unmarshal(data, &token)
	if err != nil {
		return "", fmt.Errorf("ERROR: %s", err.Error())
	}

	return token.Data.Token, nil
}
