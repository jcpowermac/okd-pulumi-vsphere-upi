// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a VMware vSphere datacenter resource. This can be used as the primary
// container of inventory objects such as hosts and virtual machines.
type Datacenter struct {
	pulumi.CustomResourceState

	// Map of custom attribute ids to value
	// strings to set for datacenter resource. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes pulumi.StringMapOutput `pulumi:"customAttributes"`
	// The folder where the datacenter should be created.
	// Forces a new resource if changed.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// Managed object ID of this datacenter.
	Moid pulumi.StringOutput `pulumi:"moid"`
	// The name of the datacenter. This name needs to be unique
	// within the folder. Forces a new resource if changed.
	Name pulumi.StringOutput `pulumi:"name"`
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewDatacenter registers a new resource with the given unique name, arguments, and options.
func NewDatacenter(ctx *pulumi.Context,
	name string, args *DatacenterArgs, opts ...pulumi.ResourceOption) (*Datacenter, error) {
	if args == nil {
		args = &DatacenterArgs{}
	}
	var resource Datacenter
	err := ctx.RegisterResource("vsphere:index/datacenter:Datacenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatacenter gets an existing Datacenter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatacenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatacenterState, opts ...pulumi.ResourceOption) (*Datacenter, error) {
	var resource Datacenter
	err := ctx.ReadResource("vsphere:index/datacenter:Datacenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Datacenter resources.
type datacenterState struct {
	// Map of custom attribute ids to value
	// strings to set for datacenter resource. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The folder where the datacenter should be created.
	// Forces a new resource if changed.
	Folder *string `pulumi:"folder"`
	// Managed object ID of this datacenter.
	Moid *string `pulumi:"moid"`
	// The name of the datacenter. This name needs to be unique
	// within the folder. Forces a new resource if changed.
	Name *string `pulumi:"name"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
}

type DatacenterState struct {
	// Map of custom attribute ids to value
	// strings to set for datacenter resource. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes pulumi.StringMapInput
	// The folder where the datacenter should be created.
	// Forces a new resource if changed.
	Folder pulumi.StringPtrInput
	// Managed object ID of this datacenter.
	Moid pulumi.StringPtrInput
	// The name of the datacenter. This name needs to be unique
	// within the folder. Forces a new resource if changed.
	Name pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
}

func (DatacenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*datacenterState)(nil)).Elem()
}

type datacenterArgs struct {
	// Map of custom attribute ids to value
	// strings to set for datacenter resource. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The folder where the datacenter should be created.
	// Forces a new resource if changed.
	Folder *string `pulumi:"folder"`
	// The name of the datacenter. This name needs to be unique
	// within the folder. Forces a new resource if changed.
	Name *string `pulumi:"name"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a Datacenter resource.
type DatacenterArgs struct {
	// Map of custom attribute ids to value
	// strings to set for datacenter resource. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes pulumi.StringMapInput
	// The folder where the datacenter should be created.
	// Forces a new resource if changed.
	Folder pulumi.StringPtrInput
	// The name of the datacenter. This name needs to be unique
	// within the folder. Forces a new resource if changed.
	Name pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
}

func (DatacenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datacenterArgs)(nil)).Elem()
}
