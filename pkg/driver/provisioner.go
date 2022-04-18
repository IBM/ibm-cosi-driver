package driver

import (
	"context"
	"fmt"
	cosClient "github.com/IBM/ibm-cosi-driver/pkg/util/cosclient"
	namewriter "github.com/IBM/ibm-cosi-driver/pkg/util/namewriter"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"

	cosi "sigs.k8s.io/container-object-storage-interface-spec"
)

type ProvisionerServer struct {
	cosClient *cosClient.ObjectStorageClient

	// iamClient   *iam.Client
	provisioner string
}

func (p *ProvisionerServer) ProvisionerCreateBucket(ctx context.Context,
	req *cosi.ProvisionerCreateBucketRequest) (*cosi.ProvisionerCreateBucketResponse, error) {
	fmt.Println()

	bucketName := req.GetName()

	klog.InfoS("PROVISIONER", "CREATE BUCKET", bucketName)

	namewriter.UpdateBucketrData("test-user", bucketName)

	err := p.cosClient.CreateBucket(bucketName)

	if err != nil {
		return nil, err
	}

	return &cosi.ProvisionerCreateBucketResponse{
		BucketId: bucketName,
	}, nil

}
func (p *ProvisionerServer) ProvisionerDeleteBucket(ctx context.Context,
	req *cosi.ProvisionerDeleteBucketRequest) (*cosi.ProvisionerDeleteBucketResponse, error) {

	return nil, status.Error(codes.Unimplemented, "ProvisionerCreateBucket: not implemented")
}

func (p *ProvisionerServer) ProvisionerGrantBucketAccess(ctx context.Context,
	req *cosi.ProvisionerGrantBucketAccessRequest) (*cosi.ProvisionerGrantBucketAccessResponse, error) {

	// bucketNmae := req.GetBucketId()
	accountId := req.GetAccountName()
	bucketName := req.GetBucketId()
	accessPolicy := req.GetAccessPolicy()
	klog.InfoS("Granting user accessPolicy to bucket", "userName", accountId, "bucketName", bucketName, "accessPolicy", accessPolicy)

	creds, err := p.cosClient.GetCreds()
	if err != nil {
		klog.Error("Could not retrive creds", err)

		return nil, err

	}

	return &cosi.ProvisionerGrantBucketAccessResponse{
		AccountId:   accountId,
		Credentials: creds,
	}, nil

}

func (p *ProvisionerServer) ProvisionerRevokeBucketAccess(ctx context.Context,
	req *cosi.ProvisionerRevokeBucketAccessRequest) (*cosi.ProvisionerRevokeBucketAccessResponse, error) {

	return nil, status.Error(codes.Unimplemented, "ProvisionerCreateBucket: not implemented")
}
