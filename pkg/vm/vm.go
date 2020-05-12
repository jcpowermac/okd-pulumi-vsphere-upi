package vm

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	installertypes "github.com/openshift/installer/pkg/types"
	vsphere "github.com/pulumi/pulumi-vsphere/sdk/v2/go/vsphere"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// VirtualMachine - all the objects we need...
type VirtualMachine struct {
	Ctx *pulumi.Context

	Datacenter     *vsphere.Datacenter
	ResourcePool   *vsphere.ResourcePool
	Folder         *vsphere.Folder
	ComputeCluster *vsphere.ComputeCluster

	LookupDatacenter     *vsphere.LookupDatacenterResult
	LookupComputeCluster *vsphere.LookupComputeClusterResult
	LookupResourcePool   *vsphere.LookupResourcePoolResult
	LookupFolder         *vsphere.LookupFolderResult

	Datastore *vsphere.GetDatastoreResult
	Network   *vsphere.GetNetworkResult
}

// NewResourcePool - create a new resourcepool, set resourcepool in VirtualMahchine
func (vm *VirtualMachine) NewResourcePool(name string) error {
	resourcePool, err := vsphere.NewResourcePool(vm.Ctx, name, &vsphere.ResourcePoolArgs{
		Name:                 pulumi.StringPtr(name),
		ParentResourcePoolId: pulumi.String(vm.LookupComputeCluster.ResourcePoolId),
	})

	if err != nil {
		return err
	}
	vm.ResourcePool = resourcePool
	return nil
}

// NewFolder - Create a new folder, set folder in VirtualMachine
func (vm *VirtualMachine) NewFolder(name string) error {
	folder, err := vsphere.NewFolder(vm.Ctx, name, &vsphere.FolderArgs{
		Path:         pulumi.String(name),
		Type:         pulumi.String("vm"),
		DatacenterId: pulumi.StringPtr(*vm.LookupComputeCluster.DatacenterId),
	})

	if err != nil {
		return err
	}
	vm.Folder = folder
	return nil
}

// NewVirtualMachine - new object
func NewVirtualMachine(ctx *pulumi.Context, ic *installertypes.InstallConfig) (*VirtualMachine, error) {
	lookupDatacenter, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{Name: &ic.Platform.VSphere.Datacenter})
	if err != nil {
		return nil, err
	}

	datacenter, err := vsphere.GetDatacenter(ctx, *lookupDatacenter.Name, pulumi.ID(lookupDatacenter.Id), &vsphere.DatacenterState{
		Name: pulumi.StringPtr(*lookupDatacenter.Name),
		Moid: pulumi.StringPtr(lookupDatacenter.Id),
	})
	if err != nil {
		return nil, err
	}

	spew.Dump(datacenter)

	lookupComputeCluster, err := vsphere.LookupComputeCluster(ctx,
		&vsphere.LookupComputeClusterArgs{
			DatacenterId: &lookupDatacenter.Id,
			Name:         ic.Platform.VSphere.Cluster,
		})

	if err != nil {
		return nil, err
	}

	computeCluster, err := vsphere.GetComputeCluster(ctx, lookupComputeCluster.Name, pulumi.ID(lookupComputeCluster.Id), &vsphere.ComputeClusterState{
		DatacenterId: pulumi.StringPtr(lookupDatacenter.Id),
		Name:         pulumi.StringPtr(lookupComputeCluster.Name),
	})

	defaultRPName := "*"
	lookupRP, err := vsphere.LookupResourcePool(ctx, &vsphere.LookupResourcePoolArgs{
		DatacenterId: &lookupComputeCluster.Id,
		Name:         &defaultRPName,
	})

	if err != nil {
		return nil, err
	}

	resourcePool, err := vsphere.GetResourcePool(ctx, *lookupRP.Name, nil, &vsphere.ResourcePoolState{
		ParentResourcePoolId: pulumi.StringPtr(lookupRP.Id),
		Name:                 pulumi.StringPtr(*lookupRP.Name),
	})

	if err != nil {
		return nil, err
	}

	defaultFolderPath := fmt.Sprintf("/%s/vm", *lookupDatacenter.Name)
	lookupFolder, err := vsphere.LookupFolder(ctx, &vsphere.LookupFolderArgs{
		Path: defaultFolderPath,
	})

	if err != nil {
		return nil, err
	}

	folder, err := vsphere.GetFolder(ctx, defaultFolderPath, pulumi.ID(lookupFolder.Id), &vsphere.FolderState{
		Path: pulumi.StringPtr(lookupFolder.Path),
	})

	datastore, err := vsphere.GetDatastore(ctx, &vsphere.GetDatastoreArgs{
		Name:         ic.Platform.VSphere.DefaultDatastore,
		DatacenterId: &lookupDatacenter.Id,
	})

	if err != nil {
		return nil, err
	}

	network, err := vsphere.GetNetwork(ctx, &vsphere.GetNetworkArgs{
		Name:         ic.Platform.VSphere.Network,
		DatacenterId: &lookupDatacenter.Id,
	})
	if err != nil {
		return nil, err
	}

	return &VirtualMachine{
		LookupDatacenter: lookupDatacenter,

		Datacenter:           datacenter,
		ResourcePool:         resourcePool,
		Network:              network,
		Datastore:            datastore,
		ComputeCluster:       computeCluster,
		LookupFolder:         lookupFolder,
		LookupComputeCluster: lookupComputeCluster,
		Folder:               folder,
		Ctx:                  ctx,
	}, nil
}

// InstanceType - Types of instances
type InstanceType string

//InstanceType enum
const (
	Bootstrap InstanceType = "bootstrap"
	Master    InstanceType = "control-plane"
	Worker    InstanceType = "compute"
)

/* TODO:
 * - ignition
 * - ip addressing
 */

// CreateCoreOSVirtualMachine - creates a virtual machine
func (vm VirtualMachine) CreateCoreOSVirtualMachine(instanceType InstanceType, quantity int, template string) error {

	templateVM, err := vsphere.LookupVirtualMachine(vm.Ctx, &vsphere.LookupVirtualMachineArgs{
		Name:         template,
		DatacenterId: &vm.LookupDatacenter.Id,
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
