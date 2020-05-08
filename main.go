package main

import (
	"github.com/davecgh/go-spew/spew"
	installertypes "github.com/openshift/installer/pkg/types"
	vspheretypes "github.com/openshift/installer/pkg/types/vsphere"
	vsphere "github.com/pulumi/pulumi-vsphere/sdk/v2/go/vsphere"
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
		spew.Dump(ic)

		dc, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
			Name: &ic.Platform.VSphere.Datacenter})
		if err != nil {
			return err
		}

		computeCluster, err := vsphere.LookupComputeCluster(ctx,
			&vsphere.LookupComputeClusterArgs{
				DatacenterId: &dc.Id,
				Name:         ic.Platform.VSphere.Cluster,
			})

		if err != nil {
			return err
		}

		datastore, err := vsphere.GetDatastore(ctx, &vsphere.GetDatastoreArgs{
			Name:         ic.Platform.VSphere.DefaultDatastore,
			DatacenterId: &dc.Id,
		})

		if err != nil {
			return err
		}

		network, err := vsphere.GetNetwork(ctx, &vsphere.GetNetworkArgs{
			Name:         ic.Platform.VSphere.Network,
			DatacenterId: &dc.Id,
		})

		resourcePool, err := vsphere.NewResourcePool(ctx, ic.ObjectMeta.Name, &vsphere.ResourcePoolArgs{
			Name:                 pulumi.StringPtr(ic.ObjectMeta.Name),
			ParentResourcePoolId: pulumi.String(computeCluster.ResourcePoolId),
		})

		if err != nil {
			return nil
		}

		_, err = vsphere.NewFolder(ctx, ic.ObjectMeta.Name, &vsphere.FolderArgs{
			Path:         pulumi.String(ic.ObjectMeta.Name),
			Type:         pulumi.String("vm"),
			DatacenterId: pulumi.StringPtr(dc.Id),
		})
		if err != nil {
			return err
		}

		templateVM, err := vsphere.LookupVirtualMachine(ctx, &vsphere.LookupVirtualMachineArgs{
			Name:         "rhcos-44.81.202003062006-0-vmware.x86_64",
			DatacenterId: &dc.Id,
		})
		if err != nil {
			return err
		}

		vmNetworkInterfaces := vsphere.VirtualMachineNetworkInterfaceArray{
			vsphere.VirtualMachineNetworkInterfaceArgs{
				NetworkId: pulumi.String(network.Id),
			},
		}

		vmDisks := vsphere.VirtualMachineDiskArray{
			vsphere.VirtualMachineDiskArgs{
				Label:           pulumi.StringPtr("disk0"),
				Size:            pulumi.IntPtr(16),
				ThinProvisioned: pulumi.BoolPtr(false),
			},
		}

		vmCloneArgs := &vsphere.VirtualMachineCloneArgs{
			TemplateUuid: pulumi.String(templateVM.Id),
		}

		vmArgs := &vsphere.VirtualMachineArgs{
			Name:           pulumi.StringPtr("jcallen-pulumi-test"),
			ResourcePoolId: resourcePool.ID(),
			DatastoreId:    pulumi.StringPtr(datastore.Id),
			NumCpus:        pulumi.IntPtr(1),
			Memory:         pulumi.IntPtr(1024),
			GuestId:        pulumi.StringPtr("rhel7_64Guest"),
			EnableDiskUuid: pulumi.BoolPtr(true),
			Clone:          vmCloneArgs,

			NetworkInterfaces:       vmNetworkInterfaces,
			Disks:                   vmDisks,
			WaitForGuestNetRoutable: pulumi.BoolPtr(false),
			WaitForGuestNetTimeout:  pulumi.IntPtr(0),

			Folder: pulumi.StringPtr(ic.Platform.VSphere.Folder),
		}

		_, err = vsphere.NewVirtualMachine(ctx, "jcallen-pulumi-test", vmArgs)

		if err != nil {
			return err
		}

		return nil
	})
}
