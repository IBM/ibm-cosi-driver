package main

import (
	"context"
	"flag"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"k8s.io/klog/v2"

	"github.com/IBM/ibm-cosi-driver/pkg/driver"
	"github.com/IBM/ibm-cosi-driver/pkg/util/cosclient"
	"sigs.k8s.io/container-object-storage-interface-provisioner-sidecar/pkg/provisioner"
)

const provisionerName = "ibm.objectstorage.k8s.io"

var (
	driverAddress      = "unix:///var/lib/cosi/cosi.sock"
	endpoint           = ""
	locationConstraint = ""
	accessKey          = ""
	secretKey          = ""
)

var cmd = &cobra.Command{
	Use:           "ibm-cosi-driver",
	Short:         "IBM COSI driver implementation",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context(), args)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	flag.Set("alsologtostderr", "true")
	kflags := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(kflags)

	persistentFlags := cmd.PersistentFlags()
	persistentFlags.AddGoFlagSet(kflags)

	stringFlag := persistentFlags.StringVarP

	stringFlag(&driverAddress,
		"driver-addr",
		"d",
		driverAddress,
		"path to unix domain socket where driver should listen")

	stringFlag(&endpoint,
		"endpoint",
		"e",
		endpoint,
		"object storage endpoint")

	stringFlag(&locationConstraint,
		"location",
		"l",
		locationConstraint,
		"object storage region")

	stringFlag(&accessKey,
		"accesskey",
		"a",
		accessKey,
		"access key for object storage")

	stringFlag(&secretKey,
		"secretkey",
		"s",
		secretKey,
		"secret key for object storage")

	viper.BindPFlags(cmd.PersistentFlags())
	cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		if viper.IsSet(f.Name) && viper.GetString(f.Name) != "" {
			cmd.PersistentFlags().Set(f.Name, viper.GetString(f.Name))
		}
	})
}

func run(ctx context.Context, args []string) error {
	creds := &cosclient.ObjectStorageCredentials{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
	identityServer, bucketProvisioner, err := driver.NewDriver(ctx, creds, endpoint, locationConstraint, provisionerName)
	if err != nil {
		return err
	}

	server, err := provisioner.NewDefaultCOSIProvisionerServer(driverAddress,
		identityServer,
		bucketProvisioner)
	if err != nil {
		return err
	}
	return server.Run(ctx)
}
