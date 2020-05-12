package main

import (
	pulumivm "okd-pulumi-vsphere-upi/pkg/vm"

	"github.com/davecgh/go-spew/spew"

	installertypes "github.com/openshift/installer/pkg/types"
	vspheretypes "github.com/openshift/installer/pkg/types/vsphere"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

func newInstallConfig(conf *config.Config) *installertypes.InstallConfig {
	ic := &installertypes.InstallConfig{}

	ic.ObjectMeta.Name = conf.Require("clusterid")
	ic.Platform.VSphere = &vspheretypes.Platform{
		Datacenter:       conf.Require("datacenter"),
		DefaultDatastore: conf.Require("datastore"),
		Folder:           conf.Require("folder"),
		Cluster:          conf.Require("cluster"),
		Network:          conf.Require("network"),
	}

	return ic
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		conf := config.New(ctx, "")
		ic := newInstallConfig(conf)

		// Write ic out to disk to be used with openshift-install
		spew.Dump(ic)

		vm, err := pulumivm.NewVirtualMachine(ctx, ic)
		if err != nil {
			return err
		}

		// todo:
		// find template?
		err = vm.CreateCoreOSVirtualMachine(pulumivm.Bootstrap, 1, "foo")
		if err != nil {
			return err
		}

		return nil
	})
}
