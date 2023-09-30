package aws

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Token struct {
  AccessToken string `json:"access_token"`
  RefreshToken string `json:"refresh_token"`
  TokenType string `json:"token_type"`
  ExpiresIn int16 `json:"expires_in"`
}

func GetToken(clientId string, clientSecret string, refreshToken string) *Token {
  newToken := &Token{}

  urlStr := "https://api.amazon.com/auth/o2/token"
  payload := url.Values{}
  payload.Set("grant_type", "refresh_token")
  payload.Set("client_id", clientId)
  payload.Set("client_secret", clientSecret)
  payload.Set("refresh_token", refreshToken)

  client := &http.Client {
  }

  req, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(payload.Encode()))
  if err != nil {
    fmt.Printf("error could not create request client: %s\n", err)
    return newToken
  }
  
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  res, err := client.Do(req)
  if err != nil {
    fmt.Printf("error client http request: %s\n", err)
    return newToken
  }
  
  defer res.Body.Close()
  body, err := io.ReadAll(res.Body)
  if err != nil {
    fmt.Printf("error could not read response body: %s\n", err)
    return newToken
  }

  json.Unmarshal(body, newToken)

  return newToken
}