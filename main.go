package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {

	key := flag.String("key", "", "Access Key ID")
	secret := flag.String("secret", "", "Secret Access Key")
	region := flag.String("region", "us-east-1", "Region")
	noCache := flag.Bool("no-cache", false, "Set no-cache to Cache Control")

	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "NOTE: bucket and file name required\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options] bucket_name filename\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	bucket := args[0]
	filename := args[1]

	file, err := os.Open(filename)
	if err != nil {
		exitErrorf("ERROR: Unable to open file %q, %v", err)
	}

	defer file.Close()

	// Credential

	credential := credentials.NewStaticCredentials(*key, *secret, "")

	sess, err := session.NewSession(&aws.Config{
		CredentialsChainVerboseErrors: aws.Bool(true),
		Credentials:                   credential,
		Region:                        aws.String(*region)},
	)
	if err != nil {
		exitErrorf("ERROR: Failed to create S3 session")
	}

	uploader := s3manager.NewUploader(sess)

	uploadInput := s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	}

	if *noCache {
		fmt.Printf("Set no-cache to %s\n", filename)
		uploadInput.CacheControl = aws.String("no-cache")
	}

	// Upload the file's body to S3 bucket as an object with the key being the
	// same as the filename.
	_, err = uploader.Upload(&uploadInput)
	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
