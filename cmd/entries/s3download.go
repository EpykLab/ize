package entries

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

func DownloadFromS3(filename string) {
	endpoint := viper.GetString("bucket.url")
	region := viper.GetString("bucket.region")
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: &endpoint,
		Region:   &region,
	}))

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Unable to open file %q, %v", filename, err)
	}
	defer file.Close()

	// sess, _ := session.NewSession(&aws.Config{
	//     Region: aws.String("us-west-2")}, // Specify the region
	// )

	downloader := s3.New(sess)

	_, err = downloader.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(endpoint),
		Key:    aws.String(filename),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				log.Fatalf("Error: %s, %s", s3.ErrCodeNoSuchKey, aerr.Error())
			default:
				log.Fatalf("Error occurred: %s", aerr.Error())
			}
		} else {
			log.Fatalf("Error occurred: %s", err.Error())
		}
	}

	// Write the content to a file
	if _, err := downloader.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(endpoint),
		Key:    aws.String(filename),
	}); err != nil {
		log.Fatalf("Failed to download file, %v", err)
	}

	log.Printf("File downloaded, %s", filename)
}
