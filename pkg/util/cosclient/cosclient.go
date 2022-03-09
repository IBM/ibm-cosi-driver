package cosclient

import (
	"fmt"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/awserr"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
	"github.com/golang/glog"
)

type ObjectStorageClient struct {
	svc *s3.S3
}

type ObjectStorageCredentials struct {
	AuthType          string
	AccessKey         string
	SecretKey         string
	APIKey            string
	ServiceInstanceID string
	IAMEndpoint       string
}

func NewCOSClient(endpoint, locationConstraint string, creds *ObjectStorageCredentials) (*ObjectStorageClient, error) {
	var sdkCreds *credentials.Credentials
	if creds.AuthType == "iam" {
		sdkCreds = ibmiam.NewStaticCredentials(aws.NewConfig(), creds.APIKey, creds.ServiceInstanceID, creds.IAMEndpoint)
	} else {
		sdkCreds = credentials.NewStaticCredentials(creds.AccessKey, creds.SecretKey, "")
	}
	sess := session.Must(session.NewSession(&aws.Config{
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String(endpoint),
		Credentials:      sdkCreds,
		Region:           aws.String(locationConstraint),
	}))

	service := s3.New(sess)

	return &ObjectStorageClient{
		svc: service,
	}, nil
}

func (s *ObjectStorageClient) CreateBucket(name string) error {
	return s.createBucket(name)
}

func (s *ObjectStorageClient) createBucket(name string) error {
	_, err := s.svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(name),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "BucketAlreadyOwnedByYou" {
			glog.Warning(fmt.Sprintf("bucket '%s' already exists", name))
			return nil
		}
		return err
	}

	return nil

}
