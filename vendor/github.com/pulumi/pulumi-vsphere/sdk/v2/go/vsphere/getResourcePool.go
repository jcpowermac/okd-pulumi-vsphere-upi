// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `.ResourcePool` data source can be used to discover the ID of a
// resource pool in vSphere. This is useful to fetch the ID of a resource pool
// that you want to use to create virtual machines in using the
// `.VirtualMachine` resource.
func LookupResourcePool(ctx *pulumi.Context, args *LookupResourcePoolArgs, opts ...pulumi.InvokeOption) (*LookupResourcePoolResult, error) {
	var rv LookupResourcePoolResult
	err := ctx.Invoke("vsphere:index/getResourcePool:getResourcePool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResourcePool.
type LookupResourcePoolArgs struct {
	// The managed object reference
	// ID of the datacenter the resource pool is located in.
	// This can be omitted if the search path used in `name` is an absolute path.
	// For default datacenters, use the id attribute from an empty
	// `.Datacenter` data source.
	DatacenterId *string `pulumi:"datacenterId"`
	// The name of the resource pool. This can be a name or
	// path. This is required when using vCenter.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getResourcePool.
type LookupResourcePoolResult struct {
	DatacenterId *string `pulumi:"datacenterId"`
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
}
