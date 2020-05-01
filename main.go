package main

import (
	vsphere "github.com/pulumi/pulumi-vsphere/sdk/v2/go/vsphere"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {

		conf := config.New(ctx, "")

		datacenterName := conf.Require("datacenter")
		clusterName := conf.Require("cluster")
		datastoreName := conf.Require("datastore")
		networkName := conf.Require("network")

		clusterID := conf.Require("clusterID")

		dc, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{Name: &datacenterName})
		if err != nil {
			return err
		}

		computeCluster, err := vsphere.LookupComputeCluster(ctx,
			&vsphere.LookupComputeClusterArgs{
				DatacenterId: &dc.Id,
				Name:         clusterName,
			})

		if err != nil {
			return err
		}

		datastore, err := vsphere.GetDatastore(ctx, &vsphere.GetDatastoreArgs{
			Name:         datastoreName,
			DatacenterId: &dc.Id,
		})

		if err != nil {
			return err
		}

		network, err := vsphere.GetNetwork(ctx, &vsphere.GetNetworkArgs{
			Name:         networkName,
			DatacenterId: &dc.Id,
		})

		resourcePool, err := vsphere.NewResourcePool(ctx, clusterID, &vsphere.ResourcePoolArgs{
			Name:                 pulumi.StringPtr(clusterID),
			ParentResourcePoolId: pulumi.String(computeCluster.ResourcePoolId),
		})

		if err != nil {
			return nil
		}

		_, err = vsphere.NewFolder(ctx, clusterID, &vsphere.FolderArgs{
			Path:         pulumi.String(clusterID),
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

			Folder: pulumi.StringPtr(clusterID),
		}

		_, err = vsphere.NewVirtualMachine(ctx, "jcallen-pulumi-test", vmArgs)

		if err != nil {
			return err
		}

		//spew.Dump(vm)

		return nil
	})
}
