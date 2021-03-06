// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `.getDatastore` data source can be used to discover the ID of a
// datastore in vSphere. This is useful to fetch the ID of a datastore that you
// want to use to create virtual machines in using the
// `.VirtualMachine` resource.
func GetDatastore(ctx *pulumi.Context, args *GetDatastoreArgs, opts ...pulumi.InvokeOption) (*GetDatastoreResult, error) {
	var rv GetDatastoreResult
	err := ctx.Invoke("vsphere:index/getDatastore:getDatastore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatastore.
type GetDatastoreArgs struct {
	// The managed object reference
	// ID of the datacenter the datastore is located in. This
	// can be omitted if the search path used in `name` is an absolute path. For
	// default datacenters, use the id attribute from an empty `.Datacenter`
	// data source.
	DatacenterId *string `pulumi:"datacenterId"`
	// The name of the datastore. This can be a name or path.
	Name string `pulumi:"name"`
}

// A collection of values returned by getDatastore.
type GetDatastoreResult struct {
	DatacenterId *string `pulumi:"datacenterId"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}
