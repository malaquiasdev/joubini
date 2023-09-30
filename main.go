package main

import (
	"fmt"
	"joubini/pkg/aws"
)

func main() {
	roleARN := ""
	sessionName := "test"
  region := "us-east-1"
	
  cred := aws.STSCredentials(roleARN, sessionName, region)

  fmt.Println(*cred.AccessKeyId)

  clientId := ""
  clientSecret := ""
  refreshToken := ""

  token := aws.GetToken(clientId, clientSecret, refreshToken)

  fmt.Println(token.AccessToken)
}