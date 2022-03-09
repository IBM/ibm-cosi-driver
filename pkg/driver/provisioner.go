package driver

import (
	"context"
	"fmt"
	cosClient "github.com/IBM/ibm-cosi-driver/pkg/util/cosclient"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

	err := p.cosClient.CreateBucket(bucketName)

	if err != nil {
		return nil, err
	}
	return &cosi.ProvisionerCreateBucketResponse{
		BucketId: bucketName,
	}, nil

}
func (s *ProvisionerServer) ProvisionerDeleteBucket(ctx context.Context,
	req *cosi.ProvisionerDeleteBucketRequest) (*cosi.ProvisionerDeleteBucketResponse, error) {

	return nil, status.Error(codes.Unimplemented, "ProvisionerCreateBucket: not implemented")
}

func (s *ProvisionerServer) ProvisionerGrantBucketAccess(ctx context.Context,
	req *cosi.ProvisionerGrantBucketAccessRequest) (*cosi.ProvisionerGrantBucketAccessResponse, error) {

	return nil, status.Error(codes.Unimplemented, "ProvisionerCreateBucket: not implemented")
}

func (s *ProvisionerServer) ProvisionerRevokeBucketAccess(ctx context.Context,
	req *cosi.ProvisionerRevokeBucketAccessRequest) (*cosi.ProvisionerRevokeBucketAccessResponse, error) {

	return nil, status.Error(codes.Unimplemented, "ProvisionerCreateBucket: not implemented")
}
