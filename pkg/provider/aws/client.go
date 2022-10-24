package awsClient

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/spf13/viper"
	"log"
	"os"
)

func NewClient() (aws.Config, error) {
	if viper.GetString("aws.auth.credentials") == "true" {
		cfg, err := config.LoadDefaultConfig(
			context.TODO(),
			config.WithRegion(viper.GetString("aws.region")),
			config.WithRetryMaxAttempts(3),
		)
		if err != nil {
			fmt.Printf("Error while trying to authenticate to AWS using credentials: %s\n", err)
		}

		return cfg, err

	} else if viper.GetString("aws.auth.profile") == "true" {
		cfg, err := config.LoadDefaultConfig(
			context.TODO(),
			config.WithSharedConfigProfile(viper.GetString("aws.auth.profileName")),
			config.WithRegion(viper.GetString("aws.region")),
			config.WithRetryMaxAttempts(3),
		)
		if err != nil {
			fmt.Printf("Error while trying to authenticate to AWS using SSO credentials: %s\n", err)
		}

		return cfg, err

	} else {
		fmt.Println("Couldn't find a specified auth method.")
		os.Exit(1)
	}

	return aws.Config{}, nil
}

func NewEksClient() *eks.Client {
	cfg, err := NewClient()
	if err != nil {
		log.Fatalln(err)
	}

	clt := eks.NewFromConfig(cfg)

	return clt
}
