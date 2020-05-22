// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `.ResourcePool` resource can be used to create and manage
// resource pools in standalone hosts or on compute clusters.
//
// For more information on vSphere resource pools, see [this
// page][ref-vsphere-resource_pools].
//
// [ref-vsphere-resource_pools]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-60077B40-66FF-4625-934A-641703ED7601.html
type ResourcePool struct {
	pulumi.CustomResourceState

	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable pulumi.BoolPtrOutput `pulumi:"cpuExpandable"`
	// The CPU utilization of a resource pool will not exceed
	// this limit, even if there are available resources. Set to `-1` for unlimited.
	// Default: `-1`
	CpuLimit pulumi.IntPtrOutput `pulumi:"cpuLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
	CpuReservation pulumi.IntPtrOutput `pulumi:"cpuReservation"`
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `cpuShares` will be
	// ignored.  Default: `normal`
	CpuShareLevel pulumi.StringPtrOutput `pulumi:"cpuShareLevel"`
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `cpuShareLevel` must be `custom`.
	CpuShares pulumi.IntOutput `pulumi:"cpuShares"`
	// A list of custom attributes to set on this resource.
	CustomAttributes pulumi.StringMapOutput `pulumi:"customAttributes"`
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable pulumi.BoolPtrOutput `pulumi:"memoryExpandable"`
	// The CPU utilization of a resource pool will not exceed
	// this limit, even if there are available resources. Set to `-1` for unlimited.
	// Default: `-1`
	MemoryLimit pulumi.IntPtrOutput `pulumi:"memoryLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
	MemoryReservation pulumi.IntPtrOutput `pulumi:"memoryReservation"`
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `memoryShares` will be
	// ignored.  Default: `normal`
	MemoryShareLevel pulumi.StringPtrOutput `pulumi:"memoryShareLevel"`
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `memoryShareLevel` must be `custom`.
	MemoryShares pulumi.IntOutput `pulumi:"memoryShares"`
	// The name of the resource pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// The managed object ID
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a resource pool
	// from one parent resource pool to another, both must share a common root
	// resource pool or the move will fail.
	ParentResourcePoolId pulumi.StringOutput `pulumi:"parentResourcePoolId"`
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewResourcePool registers a new resource with the given unique name, arguments, and options.
func NewResourcePool(ctx *pulumi.Context,
	name string, args *ResourcePoolArgs, opts ...pulumi.ResourceOption) (*ResourcePool, error) {
	if args == nil || args.ParentResourcePoolId == nil {
		return nil, errors.New("missing required argument 'ParentResourcePoolId'")
	}
	if args == nil {
		args = &ResourcePoolArgs{}
	}
	var resource ResourcePool
	err := ctx.RegisterResource("vsphere:index/resourcePool:ResourcePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcePool gets an existing ResourcePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcePoolState, opts ...pulumi.ResourceOption) (*ResourcePool, error) {
	var resource ResourcePool
	err := ctx.ReadResource("vsphere:index/resourcePool:ResourcePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcePool resources.
type resourcePoolState struct {
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable *bool `pulumi:"cpuExpandable"`
	// The CPU utilization of a resource pool will not exceed
	// this limit, even if there are available resources. Set to `-1` for unlimited.
	// Default: `-1`
	CpuLimit *int `pulumi:"cpuLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
	CpuReservation *int `pulumi:"cpuReservation"`
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `cpuShares` will be
	// ignored.  Default: `normal`
	CpuShareLevel *string `pulumi:"cpuShareLevel"`
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `cpuShareLevel` must be `custom`.
	CpuShares *int `pulumi:"cpuShares"`
	// A list of custom attributes to set on this resource.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable *bool `pulumi:"memoryExpandable"`
	// The CPU utilization of a resource pool will not exceed
	// this limit, even if there are available resources. Set to `-1` for unlimited.
	// Default: `-1`
	MemoryLimit *int `pulumi:"memoryLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
	MemoryReservation *int `pulumi:"memoryReservation"`
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `memoryShares` will be
	// ignored.  Default: `normal`
	MemoryShareLevel *string `pulumi:"memoryShareLevel"`
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `memoryShareLevel` must be `custom`.
	MemoryShares *int `pulumi:"memoryShares"`
	// The name of the resource pool.
	Name *string `pulumi:"name"`
	// The managed object ID
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a resource pool
	// from one parent resource pool to another, both must share a common root
	// resource pool or the move will fail.
	ParentResourcePoolId *string `pulumi:"parentResourcePoolId"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
}

type ResourcePoolState struct {
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable pulumi.BoolPtrInput
	// The CPU utilization of a resource pool will not exceed
	// this limit, even if there are available resources. Set to `-1` for unlimited.
	// Default: `-1`
	CpuLimit pulumi.IntPtrInput
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
	CpuReservation pulumi.IntPtrInput
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `cpuShares` will be
	// ignored.  Default: `normal`
	CpuShareLevel pulumi.StringPtrInput
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `cpuShareLevel` must be `custom`.
	CpuShares pulumi.IntPtrInput
	// A list of custom attributes to set on this resource.
	CustomAttributes pulumi.StringMapInput
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable pulumi.BoolPtrInput
	// The CPU utilization of a resource pool will not exceed
	// this limit, even if there are available resources. Set to `-1` for unlimited.
	// Default: `-1`
	MemoryLimit pulumi.IntPtrInput
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
	MemoryReservation pulumi.IntPtrInput
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `memoryShares` will be
	// ignored.  Default: `normal`
	MemoryShareLevel pulumi.StringPtrInput
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `memoryShareLevel` must be `custom`.
	MemoryShares pulumi.IntPtrInput
	// The name of the resource pool.
	Name pulumi.StringPtrInput
	// The managed object ID
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a resource pool
	// from one parent resource pool to another, both must share a common root
	// resource pool or the move will fail.
	ParentResourcePoolId pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
}

func (ResourcePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePoolState)(nil)).Elem()
}

type resourcePoolArgs struct {
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable *bool `pulumi:"cpuExpandable"`
	// The CPU utilization of a resource pool will not exceed
	// this limit, even if there are available resources. Set to `-1` for unlimited.
	// Default: `-1`
	CpuLimit *int `pulumi:"cpuLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
	CpuReservation *int `pulumi:"cpuReservation"`
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `cpuShares` will be
	// ignored.  Default: `normal`
	CpuShareLevel *string `pulumi:"cpuShareLevel"`
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `cpuShareLevel` must be `custom`.
	CpuShares *int `pulumi:"cpuShares"`
	// A list of custom attributes to set on this resource.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable *bool `pulumi:"memoryExpandable"`
	// The CPU utilization of a resource pool will not exceed
	// this limit, even if there are available resources. Set to `-1` for unlimited.
	// Default: `-1`
	MemoryLimit *int `pulumi:"memoryLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
	MemoryReservation *int `pulumi:"memoryReservation"`
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `memoryShares` will be
	// ignored.  Default: `normal`
	MemoryShareLevel *string `pulumi:"memoryShareLevel"`
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `memoryShareLevel` must be `custom`.
	MemoryShares *int `pulumi:"memoryShares"`
	// The name of the resource pool.
	Name *string `pulumi:"name"`
	// The managed object ID
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a resource pool
	// from one parent resource pool to another, both must share a common root
	// resource pool or the move will fail.
	ParentResourcePoolId string `pulumi:"parentResourcePoolId"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a ResourcePool resource.
type ResourcePoolArgs struct {
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable pulumi.BoolPtrInput
	// The CPU utilization of a resource pool will not exceed
	// this limit, even if there are available resources. Set to `-1` for unlimited.
	// Default: `-1`
	CpuLimit pulumi.IntPtrInput
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
	CpuReservation pulumi.IntPtrInput
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `cpuShares` will be
	// ignored.  Default: `normal`
	CpuShareLevel pulumi.StringPtrInput
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `cpuShareLevel` must be `custom`.
	CpuShares pulumi.IntPtrInput
	// A list of custom attributes to set on this resource.
	CustomAttributes pulumi.StringMapInput
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable pulumi.BoolPtrInput
	// The CPU utilization of a resource pool will not exceed
	// this limit, even if there are available resources. Set to `-1` for unlimited.
	// Default: `-1`
	MemoryLimit pulumi.IntPtrInput
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
	MemoryReservation pulumi.IntPtrInput
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `memoryShares` will be
	// ignored.  Default: `normal`
	MemoryShareLevel pulumi.StringPtrInput
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `memoryShareLevel` must be `custom`.
	MemoryShares pulumi.IntPtrInput
	// The name of the resource pool.
	Name pulumi.StringPtrInput
	// The managed object ID
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a resource pool
	// from one parent resource pool to another, both must share a common root
	// resource pool or the move will fail.
	ParentResourcePoolId pulumi.StringInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
}

func (ResourcePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePoolArgs)(nil)).Elem()
}
