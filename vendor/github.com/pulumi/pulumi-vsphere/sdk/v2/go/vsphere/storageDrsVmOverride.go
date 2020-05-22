// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `.StorageDrsVmOverride` resource can be used to add a Storage DRS
// override to a datastore cluster for a specific virtual machine. With this
// resource, one can enable or disable Storage DRS, and control the automation
// level and disk affinity for a single virtual machine without affecting the rest
// of the datastore cluster.
//
// For more information on vSphere datastore clusters and Storage DRS, see [this
// page][ref-vsphere-datastore-clusters].
//
// [ref-vsphere-datastore-clusters]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-598DF695-107E-406B-9C95-0AF961FC227A.html
type StorageDrsVmOverride struct {
	pulumi.CustomResourceState

	// The managed object reference
	// ID of the datastore cluster to put the override in.
	// Forces a new resource if changed.
	DatastoreClusterId pulumi.StringOutput `pulumi:"datastoreClusterId"`
	// Overrides any Storage DRS automation
	// levels for this virtual machine. Can be one of `automated` or `manual`. When
	// not specified, the datastore cluster's settings are used according to the
	// specific SDRS subsystem.
	SdrsAutomationLevel pulumi.StringPtrOutput `pulumi:"sdrsAutomationLevel"`
	// Overrides the default Storage DRS setting for
	// this virtual machine. When not specified, the datastore cluster setting is
	// used.
	SdrsEnabled pulumi.StringPtrOutput `pulumi:"sdrsEnabled"`
	// Overrides the intra-VM affinity setting
	// for this virtual machine. When `true`, all disks for this virtual machine
	// will be kept on the same datastore. When `false`, Storage DRS may locate
	// individual disks on different datastores if it helps satisfy cluster
	// requirements. When not specified, the datastore cluster's settings are used.
	SdrsIntraVmAffinity pulumi.StringPtrOutput `pulumi:"sdrsIntraVmAffinity"`
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId pulumi.StringOutput `pulumi:"virtualMachineId"`
}

// NewStorageDrsVmOverride registers a new resource with the given unique name, arguments, and options.
func NewStorageDrsVmOverride(ctx *pulumi.Context,
	name string, args *StorageDrsVmOverrideArgs, opts ...pulumi.ResourceOption) (*StorageDrsVmOverride, error) {
	if args == nil || args.DatastoreClusterId == nil {
		return nil, errors.New("missing required argument 'DatastoreClusterId'")
	}
	if args == nil || args.VirtualMachineId == nil {
		return nil, errors.New("missing required argument 'VirtualMachineId'")
	}
	if args == nil {
		args = &StorageDrsVmOverrideArgs{}
	}
	var resource StorageDrsVmOverride
	err := ctx.RegisterResource("vsphere:index/storageDrsVmOverride:StorageDrsVmOverride", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageDrsVmOverride gets an existing StorageDrsVmOverride resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageDrsVmOverride(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageDrsVmOverrideState, opts ...pulumi.ResourceOption) (*StorageDrsVmOverride, error) {
	var resource StorageDrsVmOverride
	err := ctx.ReadResource("vsphere:index/storageDrsVmOverride:StorageDrsVmOverride", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageDrsVmOverride resources.
type storageDrsVmOverrideState struct {
	// The managed object reference
	// ID of the datastore cluster to put the override in.
	// Forces a new resource if changed.
	DatastoreClusterId *string `pulumi:"datastoreClusterId"`
	// Overrides any Storage DRS automation
	// levels for this virtual machine. Can be one of `automated` or `manual`. When
	// not specified, the datastore cluster's settings are used according to the
	// specific SDRS subsystem.
	SdrsAutomationLevel *string `pulumi:"sdrsAutomationLevel"`
	// Overrides the default Storage DRS setting for
	// this virtual machine. When not specified, the datastore cluster setting is
	// used.
	SdrsEnabled *string `pulumi:"sdrsEnabled"`
	// Overrides the intra-VM affinity setting
	// for this virtual machine. When `true`, all disks for this virtual machine
	// will be kept on the same datastore. When `false`, Storage DRS may locate
	// individual disks on different datastores if it helps satisfy cluster
	// requirements. When not specified, the datastore cluster's settings are used.
	SdrsIntraVmAffinity *string `pulumi:"sdrsIntraVmAffinity"`
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId *string `pulumi:"virtualMachineId"`
}

type StorageDrsVmOverrideState struct {
	// The managed object reference
	// ID of the datastore cluster to put the override in.
	// Forces a new resource if changed.
	DatastoreClusterId pulumi.StringPtrInput
	// Overrides any Storage DRS automation
	// levels for this virtual machine. Can be one of `automated` or `manual`. When
	// not specified, the datastore cluster's settings are used according to the
	// specific SDRS subsystem.
	SdrsAutomationLevel pulumi.StringPtrInput
	// Overrides the default Storage DRS setting for
	// this virtual machine. When not specified, the datastore cluster setting is
	// used.
	SdrsEnabled pulumi.StringPtrInput
	// Overrides the intra-VM affinity setting
	// for this virtual machine. When `true`, all disks for this virtual machine
	// will be kept on the same datastore. When `false`, Storage DRS may locate
	// individual disks on different datastores if it helps satisfy cluster
	// requirements. When not specified, the datastore cluster's settings are used.
	SdrsIntraVmAffinity pulumi.StringPtrInput
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId pulumi.StringPtrInput
}

func (StorageDrsVmOverrideState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageDrsVmOverrideState)(nil)).Elem()
}

type storageDrsVmOverrideArgs struct {
	// The managed object reference
	// ID of the datastore cluster to put the override in.
	// Forces a new resource if changed.
	DatastoreClusterId string `pulumi:"datastoreClusterId"`
	// Overrides any Storage DRS automation
	// levels for this virtual machine. Can be one of `automated` or `manual`. When
	// not specified, the datastore cluster's settings are used according to the
	// specific SDRS subsystem.
	SdrsAutomationLevel *string `pulumi:"sdrsAutomationLevel"`
	// Overrides the default Storage DRS setting for
	// this virtual machine. When not specified, the datastore cluster setting is
	// used.
	SdrsEnabled *string `pulumi:"sdrsEnabled"`
	// Overrides the intra-VM affinity setting
	// for this virtual machine. When `true`, all disks for this virtual machine
	// will be kept on the same datastore. When `false`, Storage DRS may locate
	// individual disks on different datastores if it helps satisfy cluster
	// requirements. When not specified, the datastore cluster's settings are used.
	SdrsIntraVmAffinity *string `pulumi:"sdrsIntraVmAffinity"`
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId string `pulumi:"virtualMachineId"`
}

// The set of arguments for constructing a StorageDrsVmOverride resource.
type StorageDrsVmOverrideArgs struct {
	// The managed object reference
	// ID of the datastore cluster to put the override in.
	// Forces a new resource if changed.
	DatastoreClusterId pulumi.StringInput
	// Overrides any Storage DRS automation
	// levels for this virtual machine. Can be one of `automated` or `manual`. When
	// not specified, the datastore cluster's settings are used according to the
	// specific SDRS subsystem.
	SdrsAutomationLevel pulumi.StringPtrInput
	// Overrides the default Storage DRS setting for
	// this virtual machine. When not specified, the datastore cluster setting is
	// used.
	SdrsEnabled pulumi.StringPtrInput
	// Overrides the intra-VM affinity setting
	// for this virtual machine. When `true`, all disks for this virtual machine
	// will be kept on the same datastore. When `false`, Storage DRS may locate
	// individual disks on different datastores if it helps satisfy cluster
	// requirements. When not specified, the datastore cluster's settings are used.
	SdrsIntraVmAffinity pulumi.StringPtrInput
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId pulumi.StringInput
}

func (StorageDrsVmOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageDrsVmOverrideArgs)(nil)).Elem()
}
