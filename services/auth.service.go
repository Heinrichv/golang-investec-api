package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	models "../models"
)

// GetAuthenticationToken - Request authentication token from auth service
func GetAuthenticationToken(ClientID string, Secret string) models.AccessTokenResponse {
	client := &http.Client{}

	data := url.Values{}
	data.Set("scope", "accounts")
	data.Add("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", "https://openapi.investec.com/identity/v2/oauth2/token", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.SetBasicAuth(ClientID, Secret)

	if err != nil {
		log.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var response models.AccessTokenResponse
	json.Unmarshal(f, &response)

	return response
}
