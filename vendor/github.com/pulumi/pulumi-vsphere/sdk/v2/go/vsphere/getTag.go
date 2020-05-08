// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupTag(ctx *pulumi.Context, args *LookupTagArgs, opts ...pulumi.InvokeOption) (*LookupTagResult, error) {
	var rv LookupTagResult
	err := ctx.Invoke("vsphere:index/getTag:getTag", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTag.
type LookupTagArgs struct {
	// The ID of the tag category the tag is located in.
	CategoryId string `pulumi:"categoryId"`
	// The name of the tag.
	Name string `pulumi:"name"`
}

// A collection of values returned by getTag.
type LookupTagResult struct {
	CategoryId  string `pulumi:"categoryId"`
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}
