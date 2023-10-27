package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-northeast-2"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           "http://localhost.localstack.cloud:4566",
				SigningRegion: "ap-northeast-2",
			}, nil
		})),
	)
	if err != nil {
		log.Fatalf("Unable to load SDK config: %v", err)
	}

	client := s3.NewFromConfig(cfg)

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String("my-bucket"),
		Prefix: aws.String("my-prefix/"),
	}

	output, err := client.ListObjectsV2(context.TODO(), input)
	if err != nil {
		log.Fatalf("Failed to list objects: %v", err)
	}

	for _, object := range output.Contents {
		fmt.Println("Object name:", *object.Key)
	}
	fmt.Println("Total objects: ", len(output.Contents))
	// Total objects:  0
}
