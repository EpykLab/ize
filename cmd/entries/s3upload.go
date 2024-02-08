/*
Copyright © 2024 Epyklab contact@epyklab.com
*/
package entries

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
)

func UploadToS3(filename string) error {
	endpoint := viper.GetString("bucket.url")
	region := "nyc3"
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: &endpoint,
		Region:   &region,
	}))

	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", filename, err)
	}

	destBucket := viper.GetString("bucket.name")
	sourceString := filename

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(destBucket),
		Key:    aws.String(sourceString),
		Body:   f,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	return nil
}
