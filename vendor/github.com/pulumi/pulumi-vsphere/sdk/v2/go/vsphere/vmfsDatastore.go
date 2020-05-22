// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `.VmfsDatastore` resource can be used to create and manage VMFS
// datastores on an ESXi host or a set of hosts. The resource supports using any
// SCSI device that can generally be used in a datastore, such as local disks, or
// disks presented to a host or multiple hosts over Fibre Channel or iSCSI.
// Devices can be specified manually, or discovered using the
// [`.getVmfsDisks`][data-source-vmfs-disks] data source.
//
// [data-source-vmfs-disks]: /docs/providers/vsphere/d/vmfs_disks.html
//
// ## Auto-Mounting of Datastores Within vCenter
//
// Note that the current behaviour of this resource will auto-mount any created
// datastores to any other host within vCenter that has access to the same disk.
//
// Example: You want to create a datastore with a iSCSI LUN that is visible on 3
// hosts in a single vSphere cluster (`esxi1`, `esxi2` and `esxi3`). When you
// create the datastore on `esxi1`, the datastore will be automatically mounted on
// `esxi2` and `esxi3`, without the need to configure the resource on either of
// those two hosts.
//
// Future versions of this resource may allow you to control the hosts that a
// datastore is mounted to, but currently, this automatic behaviour cannot be
// changed, so keep this in mind when writing your configurations and deploying
// your disks.
//
// ## Increasing Datastore Size
//
// To increase the size of a datastore, you must add additional disks to the
// `disks` attribute. Expanding the size of a datastore by increasing the size of
// an already provisioned disk is currently not supported (but may be in future
// versions of this resource).
//
// > **NOTE:** You cannot decrease the size of a datastore. If the resource
// detects disks removed from the configuration, the provider will give an error.
//
// [cmd-taint]: /docs/commands/taint.html
type VmfsDatastore struct {
	pulumi.CustomResourceState

	// The connectivity status of the datastore. If this is `false`,
	// some other computed attributes may be out of date.
	Accessible pulumi.BoolOutput `pulumi:"accessible"`
	// Maximum capacity of the datastore, in megabytes.
	Capacity pulumi.IntOutput `pulumi:"capacity"`
	// Map of custom attribute ids to attribute
	// value string to set on datastore resource.
	CustomAttributes pulumi.StringMapOutput `pulumi:"customAttributes"`
	// The managed object
	// ID of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId pulumi.StringPtrOutput `pulumi:"datastoreClusterId"`
	// The disks to use with the datastore.
	Disks pulumi.StringArrayOutput `pulumi:"disks"`
	// The relative path to a folder to put this datastore in.
	// This is a path relative to the datacenter you are deploying the datastore to.
	// Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
	// The provider will place a datastore named `test` in a datastore folder
	// located at `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/test`. Conflicts with
	// `datastoreClusterId`.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// Available space of this datastore, in megabytes.
	FreeSpace pulumi.IntOutput `pulumi:"freeSpace"`
	// The managed object ID of
	// the host to set the datastore up on. Note that this is not necessarily the
	// only host that the datastore will be set up on - see
	// here for more info. Forces a
	// new resource if changed.
	HostSystemId pulumi.StringOutput `pulumi:"hostSystemId"`
	// The current maintenance mode state of the datastore.
	MaintenanceMode pulumi.StringOutput `pulumi:"maintenanceMode"`
	// If `true`, more than one host in the datacenter has
	// been configured with access to the datastore.
	MultipleHostAccess pulumi.BoolOutput `pulumi:"multipleHostAccess"`
	// The name of the datastore. Forces a new resource if
	// changed.
	Name pulumi.StringOutput `pulumi:"name"`
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Total additional storage space, in megabytes,
	// potentially used by all virtual machines on this datastore.
	UncommittedSpace pulumi.IntOutput `pulumi:"uncommittedSpace"`
	// The unique locator for the datastore.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewVmfsDatastore registers a new resource with the given unique name, arguments, and options.
func NewVmfsDatastore(ctx *pulumi.Context,
	name string, args *VmfsDatastoreArgs, opts ...pulumi.ResourceOption) (*VmfsDatastore, error) {
	if args == nil || args.Disks == nil {
		return nil, errors.New("missing required argument 'Disks'")
	}
	if args == nil || args.HostSystemId == nil {
		return nil, errors.New("missing required argument 'HostSystemId'")
	}
	if args == nil {
		args = &VmfsDatastoreArgs{}
	}
	var resource VmfsDatastore
	err := ctx.RegisterResource("vsphere:index/vmfsDatastore:VmfsDatastore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVmfsDatastore gets an existing VmfsDatastore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVmfsDatastore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VmfsDatastoreState, opts ...pulumi.ResourceOption) (*VmfsDatastore, error) {
	var resource VmfsDatastore
	err := ctx.ReadResource("vsphere:index/vmfsDatastore:VmfsDatastore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VmfsDatastore resources.
type vmfsDatastoreState struct {
	// The connectivity status of the datastore. If this is `false`,
	// some other computed attributes may be out of date.
	Accessible *bool `pulumi:"accessible"`
	// Maximum capacity of the datastore, in megabytes.
	Capacity *int `pulumi:"capacity"`
	// Map of custom attribute ids to attribute
	// value string to set on datastore resource.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The managed object
	// ID of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId *string `pulumi:"datastoreClusterId"`
	// The disks to use with the datastore.
	Disks []string `pulumi:"disks"`
	// The relative path to a folder to put this datastore in.
	// This is a path relative to the datacenter you are deploying the datastore to.
	// Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
	// The provider will place a datastore named `test` in a datastore folder
	// located at `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/test`. Conflicts with
	// `datastoreClusterId`.
	Folder *string `pulumi:"folder"`
	// Available space of this datastore, in megabytes.
	FreeSpace *int `pulumi:"freeSpace"`
	// The managed object ID of
	// the host to set the datastore up on. Note that this is not necessarily the
	// only host that the datastore will be set up on - see
	// here for more info. Forces a
	// new resource if changed.
	HostSystemId *string `pulumi:"hostSystemId"`
	// The current maintenance mode state of the datastore.
	MaintenanceMode *string `pulumi:"maintenanceMode"`
	// If `true`, more than one host in the datacenter has
	// been configured with access to the datastore.
	MultipleHostAccess *bool `pulumi:"multipleHostAccess"`
	// The name of the datastore. Forces a new resource if
	// changed.
	Name *string `pulumi:"name"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
	// Total additional storage space, in megabytes,
	// potentially used by all virtual machines on this datastore.
	UncommittedSpace *int `pulumi:"uncommittedSpace"`
	// The unique locator for the datastore.
	Url *string `pulumi:"url"`
}

type VmfsDatastoreState struct {
	// The connectivity status of the datastore. If this is `false`,
	// some other computed attributes may be out of date.
	Accessible pulumi.BoolPtrInput
	// Maximum capacity of the datastore, in megabytes.
	Capacity pulumi.IntPtrInput
	// Map of custom attribute ids to attribute
	// value string to set on datastore resource.
	CustomAttributes pulumi.StringMapInput
	// The managed object
	// ID of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId pulumi.StringPtrInput
	// The disks to use with the datastore.
	Disks pulumi.StringArrayInput
	// The relative path to a folder to put this datastore in.
	// This is a path relative to the datacenter you are deploying the datastore to.
	// Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
	// The provider will place a datastore named `test` in a datastore folder
	// located at `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/test`. Conflicts with
	// `datastoreClusterId`.
	Folder pulumi.StringPtrInput
	// Available space of this datastore, in megabytes.
	FreeSpace pulumi.IntPtrInput
	// The managed object ID of
	// the host to set the datastore up on. Note that this is not necessarily the
	// only host that the datastore will be set up on - see
	// here for more info. Forces a
	// new resource if changed.
	HostSystemId pulumi.StringPtrInput
	// The current maintenance mode state of the datastore.
	MaintenanceMode pulumi.StringPtrInput
	// If `true`, more than one host in the datacenter has
	// been configured with access to the datastore.
	MultipleHostAccess pulumi.BoolPtrInput
	// The name of the datastore. Forces a new resource if
	// changed.
	Name pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
	// Total additional storage space, in megabytes,
	// potentially used by all virtual machines on this datastore.
	UncommittedSpace pulumi.IntPtrInput
	// The unique locator for the datastore.
	Url pulumi.StringPtrInput
}

func (VmfsDatastoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmfsDatastoreState)(nil)).Elem()
}

type vmfsDatastoreArgs struct {
	// Map of custom attribute ids to attribute
	// value string to set on datastore resource.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The managed object
	// ID of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId *string `pulumi:"datastoreClusterId"`
	// The disks to use with the datastore.
	Disks []string `pulumi:"disks"`
	// The relative path to a folder to put this datastore in.
	// This is a path relative to the datacenter you are deploying the datastore to.
	// Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
	// The provider will place a datastore named `test` in a datastore folder
	// located at `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/test`. Conflicts with
	// `datastoreClusterId`.
	Folder *string `pulumi:"folder"`
	// The managed object ID of
	// the host to set the datastore up on. Note that this is not necessarily the
	// only host that the datastore will be set up on - see
	// here for more info. Forces a
	// new resource if changed.
	HostSystemId string `pulumi:"hostSystemId"`
	// The name of the datastore. Forces a new resource if
	// changed.
	Name *string `pulumi:"name"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a VmfsDatastore resource.
type VmfsDatastoreArgs struct {
	// Map of custom attribute ids to attribute
	// value string to set on datastore resource.
	CustomAttributes pulumi.StringMapInput
	// The managed object
	// ID of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId pulumi.StringPtrInput
	// The disks to use with the datastore.
	Disks pulumi.StringArrayInput
	// The relative path to a folder to put this datastore in.
	// This is a path relative to the datacenter you are deploying the datastore to.
	// Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
	// The provider will place a datastore named `test` in a datastore folder
	// located at `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/test`. Conflicts with
	// `datastoreClusterId`.
	Folder pulumi.StringPtrInput
	// The managed object ID of
	// the host to set the datastore up on. Note that this is not necessarily the
	// only host that the datastore will be set up on - see
	// here for more info. Forces a
	// new resource if changed.
	HostSystemId pulumi.StringInput
	// The name of the datastore. Forces a new resource if
	// changed.
	Name pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
}

func (VmfsDatastoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmfsDatastoreArgs)(nil)).Elem()
}
