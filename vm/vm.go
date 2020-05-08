package vm

import (
	installertypes "github.com/openshift/installer/pkg/types"
	vsphere "github.com/pulumi/pulumi-vsphere/sdk/v2/go/vsphere"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Create

type VirtualMachine struct {
	Datacenter     *vsphere.LookupDatacenterResult
	ComputeCluster *vsphere.LookupComputeClusterResult
	Datastore      *vsphere.GetDatastoreResult
	Network        *vsphere.GetNetworkResult
	ResourcePool   *vsphere.ResourcePool
	Folder         *vsphere.Folder
	Ctx            *pulumi.Context
}

// NewVirtualMachine - new object
func NewVirtualMachine(ctx *pulumi.Context, ic *installertypes.InstallConfig) (*VirtualMachine, error) {
	dc, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{Name: &ic.Platform.VSphere.Datacenter})
	if err != nil {
		return nil, err
	}

	computeCluster, err := vsphere.LookupComputeCluster(ctx,
		&vsphere.LookupComputeClusterArgs{
			DatacenterId: &dc.Id,
			Name:         ic.Platform.VSphere.Cluster,
		})

	if err != nil {
		return nil, err
	}

	datastore, err := vsphere.GetDatastore(ctx, &vsphere.GetDatastoreArgs{
		Name:         ic.Platform.VSphere.DefaultDatastore,
		DatacenterId: &dc.Id,
	})

	if err != nil {
		return nil, err
	}

	network, err := vsphere.GetNetwork(ctx, &vsphere.GetNetworkArgs{
		Name:         ic.Platform.VSphere.Network,
		DatacenterId: &dc.Id,
	})
	if err != nil {
		return nil, err
	}

	resourcePool, err := vsphere.NewResourcePool(ctx, ic.ObjectMeta.Name, &vsphere.ResourcePoolArgs{
		Name:                 pulumi.StringPtr(ic.ObjectMeta.Name),
		ParentResourcePoolId: pulumi.String(computeCluster.ResourcePoolId),
	})

	if err != nil {
		return nil, err
	}

	folder, err := vsphere.NewFolder(ctx, ic.ObjectMeta.Name, &vsphere.FolderArgs{
		Path:         pulumi.String(ic.ObjectMeta.Name),
		Type:         pulumi.String("vm"),
		DatacenterId: pulumi.StringPtr(dc.Id),
	})

	if err != nil {
		return nil, err
	}

	return &VirtualMachine{
		Datacenter:     dc,
		ResourcePool:   resourcePool,
		Network:        network,
		Datastore:      datastore,
		ComputeCluster: computeCluster,
		Folder:         folder,
		Ctx:            ctx,
	}, nil
}

// CreateCoreOSVirtualMachine - creates a virtual machine
func (vm VirtualMachine) CreateCoreOSVirtualMachine() error {

	templateVM, err := vsphere.LookupVirtualMachine(vm.Ctx, &vsphere.LookupVirtualMachineArgs{
		Name:         "rhcos-44.81.202003062006-0-vmware.x86_64",
		DatacenterId: &vm.Datacenter.Id,
	})
	if err != nil {
		return err
	}

	vmNetworkInterfaces := vsphere.VirtualMachineNetworkInterfaceArray{
		vsphere.VirtualMachineNetworkInterfaceArgs{
			NetworkId: pulumi.String(vm.Network.Id),
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
		ResourcePoolId: vm.ResourcePool.ID(),
		DatastoreId:    pulumi.StringPtr(vm.Datastore.Id),
		NumCpus:        pulumi.IntPtr(1),
		Memory:         pulumi.IntPtr(1024),
		GuestId:        pulumi.StringPtr("rhel7_64Guest"),
		EnableDiskUuid: pulumi.BoolPtr(true),
		Clone:          vmCloneArgs,

		NetworkInterfaces:       vmNetworkInterfaces,
		Disks:                   vmDisks,
		WaitForGuestNetRoutable: pulumi.BoolPtr(false),
		WaitForGuestNetTimeout:  pulumi.IntPtr(0),

		Folder: vm.Folder.Path,
	}

	_, err = vsphere.NewVirtualMachine(vm.Ctx, "jcallen-pulumi-test", vmArgs)

	if err != nil {
		return err
	}

	return nil
}
