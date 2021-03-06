// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `.ComputeClusterHostGroup` resource can be used to manage groups
// of hosts in a cluster, either created by the
// `.ComputeCluster` resource or looked up
// by the `.ComputeCluster` data source.
//
//
// This resource mainly serves as an input to the
// `.ComputeClusterVmHostRule`
// resource - see the documentation for that resource for further details on how
// to use host groups.
//
// > **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
//
// > **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
type ComputeClusterHostGroup struct {
	pulumi.CustomResourceState

	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringOutput `pulumi:"computeClusterId"`
	// The managed object IDs of
	// the hosts to put in the cluster.
	HostSystemIds pulumi.StringArrayOutput `pulumi:"hostSystemIds"`
	// The name of the host group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewComputeClusterHostGroup registers a new resource with the given unique name, arguments, and options.
func NewComputeClusterHostGroup(ctx *pulumi.Context,
	name string, args *ComputeClusterHostGroupArgs, opts ...pulumi.ResourceOption) (*ComputeClusterHostGroup, error) {
	if args == nil || args.ComputeClusterId == nil {
		return nil, errors.New("missing required argument 'ComputeClusterId'")
	}
	if args == nil {
		args = &ComputeClusterHostGroupArgs{}
	}
	var resource ComputeClusterHostGroup
	err := ctx.RegisterResource("vsphere:index/computeClusterHostGroup:ComputeClusterHostGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeClusterHostGroup gets an existing ComputeClusterHostGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeClusterHostGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeClusterHostGroupState, opts ...pulumi.ResourceOption) (*ComputeClusterHostGroup, error) {
	var resource ComputeClusterHostGroup
	err := ctx.ReadResource("vsphere:index/computeClusterHostGroup:ComputeClusterHostGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeClusterHostGroup resources.
type computeClusterHostGroupState struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId *string `pulumi:"computeClusterId"`
	// The managed object IDs of
	// the hosts to put in the cluster.
	HostSystemIds []string `pulumi:"hostSystemIds"`
	// The name of the host group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name *string `pulumi:"name"`
}

type ComputeClusterHostGroupState struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringPtrInput
	// The managed object IDs of
	// the hosts to put in the cluster.
	HostSystemIds pulumi.StringArrayInput
	// The name of the host group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name pulumi.StringPtrInput
}

func (ComputeClusterHostGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterHostGroupState)(nil)).Elem()
}

type computeClusterHostGroupArgs struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId string `pulumi:"computeClusterId"`
	// The managed object IDs of
	// the hosts to put in the cluster.
	HostSystemIds []string `pulumi:"hostSystemIds"`
	// The name of the host group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ComputeClusterHostGroup resource.
type ComputeClusterHostGroupArgs struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringInput
	// The managed object IDs of
	// the hosts to put in the cluster.
	HostSystemIds pulumi.StringArrayInput
	// The name of the host group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name pulumi.StringPtrInput
}

func (ComputeClusterHostGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterHostGroupArgs)(nil)).Elem()
}
