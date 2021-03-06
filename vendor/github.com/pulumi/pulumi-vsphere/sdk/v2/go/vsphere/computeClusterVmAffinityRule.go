// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `.ComputeClusterVmAffinityRule` resource can be used to manage
// VM affinity rules in a cluster, either created by the
// `.ComputeCluster` resource or looked up
// by the `.ComputeCluster` data source.
//
// This rule can be used to tell a set to virtual machines to run together on a
// single host within a cluster. When configured, DRS will make a best effort to
// ensure that the virtual machines run on the same host, or prevent any operation
// that would keep that from happening, depending on the value of the
// `mandatory` flag.
//
// > Keep in mind that this rule can only be used to tell VMs to run together on
// a _non-specific_ host - it can't be used to pin VMs to a host. For that, see
// the
// `.ComputeClusterVmHostRule`
// resource.
//
// > **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
//
// > **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
type ComputeClusterVmAffinityRule struct {
	pulumi.CustomResourceState

	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringOutput `pulumi:"computeClusterId"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory pulumi.BoolPtrOutput `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// The UUIDs of the virtual machines to run
	// on the same host together.
	VirtualMachineIds pulumi.StringArrayOutput `pulumi:"virtualMachineIds"`
}

// NewComputeClusterVmAffinityRule registers a new resource with the given unique name, arguments, and options.
func NewComputeClusterVmAffinityRule(ctx *pulumi.Context,
	name string, args *ComputeClusterVmAffinityRuleArgs, opts ...pulumi.ResourceOption) (*ComputeClusterVmAffinityRule, error) {
	if args == nil || args.ComputeClusterId == nil {
		return nil, errors.New("missing required argument 'ComputeClusterId'")
	}
	if args == nil || args.VirtualMachineIds == nil {
		return nil, errors.New("missing required argument 'VirtualMachineIds'")
	}
	if args == nil {
		args = &ComputeClusterVmAffinityRuleArgs{}
	}
	var resource ComputeClusterVmAffinityRule
	err := ctx.RegisterResource("vsphere:index/computeClusterVmAffinityRule:ComputeClusterVmAffinityRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeClusterVmAffinityRule gets an existing ComputeClusterVmAffinityRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeClusterVmAffinityRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeClusterVmAffinityRuleState, opts ...pulumi.ResourceOption) (*ComputeClusterVmAffinityRule, error) {
	var resource ComputeClusterVmAffinityRule
	err := ctx.ReadResource("vsphere:index/computeClusterVmAffinityRule:ComputeClusterVmAffinityRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeClusterVmAffinityRule resources.
type computeClusterVmAffinityRuleState struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId *string `pulumi:"computeClusterId"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled *bool `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory *bool `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the cluster.
	Name *string `pulumi:"name"`
	// The UUIDs of the virtual machines to run
	// on the same host together.
	VirtualMachineIds []string `pulumi:"virtualMachineIds"`
}

type ComputeClusterVmAffinityRuleState struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringPtrInput
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrInput
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory pulumi.BoolPtrInput
	// The name of the rule. This must be unique in the cluster.
	Name pulumi.StringPtrInput
	// The UUIDs of the virtual machines to run
	// on the same host together.
	VirtualMachineIds pulumi.StringArrayInput
}

func (ComputeClusterVmAffinityRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterVmAffinityRuleState)(nil)).Elem()
}

type computeClusterVmAffinityRuleArgs struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId string `pulumi:"computeClusterId"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled *bool `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory *bool `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the cluster.
	Name *string `pulumi:"name"`
	// The UUIDs of the virtual machines to run
	// on the same host together.
	VirtualMachineIds []string `pulumi:"virtualMachineIds"`
}

// The set of arguments for constructing a ComputeClusterVmAffinityRule resource.
type ComputeClusterVmAffinityRuleArgs struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringInput
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrInput
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory pulumi.BoolPtrInput
	// The name of the rule. This must be unique in the cluster.
	Name pulumi.StringPtrInput
	// The UUIDs of the virtual machines to run
	// on the same host together.
	VirtualMachineIds pulumi.StringArrayInput
}

func (ComputeClusterVmAffinityRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterVmAffinityRuleArgs)(nil)).Elem()
}
