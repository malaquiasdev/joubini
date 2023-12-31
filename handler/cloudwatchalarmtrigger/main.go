package main

import (
	"fmt"
	"joubini/pkg/aws"
	"joubini/pkg/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlerCloudWatchEvent(ev events.CloudWatchAlarmTrigger) error {
	roleARN := utils.GetEnv("AWS_ROLE_ARN", "")
	sessionName := utils.GetEnv("AWS_SESSION_NAME", "test")
  region := utils.GetEnv("AWS_DEFAULT_REGION", "us-east-1")
  
  cred := aws.STSCredentials(roleARN, sessionName, region)

  fmt.Println(*cred.AccessKeyId)

  clientId := utils.GetEnv("LWA_ID", "")
  clientSecret := utils.GetEnv("LWA_SECRET", "")
  refreshToken := utils.GetEnv("LWA_REFRESH_TOKEN", "")

  token := aws.GetToken(clientId, clientSecret, refreshToken)

  fmt.Println(token.AccessToken)
	return nil
}

func main() {
	lambda.Start(handlerCloudWatchEvent)
}