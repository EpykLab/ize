/*
Copyright © 2024 Epyklab contact@epyklab.com
*/

package entries

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

func ListBucketObjects(filter string) {
	endpoint := viper.GetString("bucket.url")
	region := viper.GetString("bucket.region")
	bucket := viper.GetString("bucket.name")

	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(endpoint),
		Region:   aws.String(region),
	})
	if err != nil {
		log.Fatalf("Error creating session: %s", err)
	}

	// Create an S3 service client
	svc := s3.New(sess)

	// List objects
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	if err != nil {
		log.Fatalf("Unable to list items in bucket %q, %v", bucket, err)
	}

	if filter == "all" {
		for _, item := range resp.Contents {
			fmt.Println(*item.Key)
		}
	} else {
		for _, item := range resp.Contents {
			if strings.Contains(item.String(), filter) {
				fmt.Println(item)
			}
		}
	}
}
