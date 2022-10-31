package s3_client

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	Client *s3.S3
	config *Config
}

var s3ClientInstance *S3Client

func InstallS3Client(configKeys ...string) *S3Client {
	if s3ClientInstance != nil {
		return s3ClientInstance
	}

	getConfigFromEnv(configKeys...)

	if config == nil || config.AccessKeyID == "" || config.SecretAccessKey == "" {
		err := fmt.Errorf("need config for aws s3 client first")
		panic(err)
	}

	credential := credentials.NewStaticCredentials(config.AccessKeyID, config.SecretAccessKey, "")
	_, err := credential.Get()
	if err != nil {
		if err != nil {
			panic(err)
		}
	}

	awsConfig := aws.NewConfig().WithRegion(config.Region).WithCredentials(credential)
	if config.BaseURL != "" {
		awsConfig = awsConfig.WithEndpoint(config.BaseURL)
	}

	newSession := session.Must(session.NewSession())

	s3ClientInstance = &S3Client{
		Client: s3.New(newSession, awsConfig),
		config: config,
	}

	return s3ClientInstance
}

func GetS3ClientInstance() *S3Client {
	if s3ClientInstance == nil {
		s3ClientInstance = InstallS3Client()
	}

	return s3ClientInstance
}