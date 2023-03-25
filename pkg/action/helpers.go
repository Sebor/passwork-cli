package action

import (
	"encoding/json"
	"fmt"
	"io"
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

type GlobalConfig struct {
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body: %s", err.Error())
	}

	fmt.Println(string(body))

	return body, nil
}

func (g *GlobalConfig) ReadToken() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tokenFilepath := filepath.Join(homeDir, g.TokenFile)
	data, err := os.ReadFile(tokenFilepath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	token := APIToken{}
	err = json.Unmarshal(data, &token)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return token.Data.Token
}

func (g *GlobalConfig) WriteToken(body []byte) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tokenFilepath := filepath.Join(homeDir, g.TokenFile)
	err = os.WriteFile(tokenFilepath, body, 0600)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
