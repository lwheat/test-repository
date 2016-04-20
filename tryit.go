package main

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	awssqs "github.com/aws/aws-sdk-go/service/sqs"
)

var (
	AccessKey = "access"
	SecretKey = "secret"
)

func newOne() *SQS {
	creds := credentials.NewChainCredentials(
		[]credentials.Provider{
			&credentials.StaticProvider{Value: credentials.Value{
				AccessKeyID:     AccessKey,
				SecretAccessKey: SecretKey,
				SessionToken:    "",
			}},
			&credentials.EnvProvider{},
			&credentials.SharedCredentialsProvider{},
		})

	_, err := creds.Get()
	if err != nil {
		return nil, err
	}

	awsSession := session.New(&awssdk.Config{
		Credentials: creds,
		Region:      awssdk.String(p.Region),
	})
	awsS3 := awss3.New(awsSession)
	awsSQS := awssqs.New(awsSession)
	return (awsSQS)
}

var x = newOne()
