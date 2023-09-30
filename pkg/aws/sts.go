package aws

import (
	"context"
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
)

// STSAssumeRoleAPI defines the interface for the AssumeRole function.
// We use this interface to test the function using a mocked service.
type stsAssumeRoleAPI interface {
	AssumeRole(ctx context.Context,
		params *sts.AssumeRoleInput,
		optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error)
}

// TakeRole gets temporary security credentials to access resources.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If successful, an AssumeRoleOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to AssumeRole.
func assumeRole(c context.Context, api stsAssumeRoleAPI, input *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
	return api.AssumeRole(c, input)
}

func STSCredentials(awsRoleARN string, awsSessionName string, awsRegion string) *types.Credentials {
	roleARN := flag.String("r", awsRoleARN, "The Amazon Resource Name (ARN) of the role to assume")
	sessionName := flag.String("s", awsSessionName, "The name of the session")
	flag.Parse()

	if *roleARN == "" || *sessionName == "" {
		fmt.Println("You must supply a role ARN and session name")
		fmt.Println("-r ROLE-ARN -s SESSION-NAME")
		return nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
	if err != nil {
		panic("sts configuration error, " + err.Error())
	}

		client := sts.NewFromConfig(cfg)

	input := &sts.AssumeRoleInput{
		RoleArn:         roleARN,
		RoleSessionName: sessionName,
	}

	result, err := assumeRole(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error assuming the role:")
		fmt.Println(err)
		return nil
	}

	return result.Credentials
}