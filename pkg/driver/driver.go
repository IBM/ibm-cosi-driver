package driver

import (
	"context"

	"github.com/IBM/ibm-cosi-driver/pkg/util/cosclient"
	"k8s.io/klog/v2"
)

func NewDriver(ctx context.Context, creds *cosclient.ObjectStorageCredentials, endpoint, locationConstraint, provisioner string) (*IdentityServer, *ProvisionerServer, error) {

	cosClient, err := cosclient.NewCOSClient(endpoint, locationConstraint, creds)
	if err != nil {
		klog.Fatalln(err)
	}
	return &IdentityServer{
			provisioner: provisioner,
		}, &ProvisionerServer{
			provisioner: provisioner,
			cosClient:   cosClient,
		}, nil
}
