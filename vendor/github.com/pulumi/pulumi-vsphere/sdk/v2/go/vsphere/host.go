// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a VMware vSphere host resource. This represents an ESXi host that
// can be used either as part of a Compute Cluster or Standalone.
type Host struct {
	pulumi.CustomResourceState

	// The ID of the Compute Cluster this host should
	// be added to. This should not be set if `datacenter` is set.
	Cluster pulumi.StringPtrOutput `pulumi:"cluster"`
	// If set to false then the host will be disconected.
	// Default is `false`.
	Connected pulumi.BoolPtrOutput `pulumi:"connected"`
	// The ID of the datacenter this host should
	// be added to. This should not be set if `cluster` is set.
	Datacenter pulumi.StringPtrOutput `pulumi:"datacenter"`
	// If set to true then it will force the host to be added, even
	// if the host is already connected to a different vSphere instance. Default is `false`
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// FQDN or IP address of the host to be added.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The license key that will be applied to the host.
	// The license key is expected to be present in vSphere.
	License pulumi.StringPtrOutput `pulumi:"license"`
	// Set the lockdown state of the host. Valid options are
	// `disabled`, `normal`, and `strict`. Default is `disabled`.
	Lockdown pulumi.StringPtrOutput `pulumi:"lockdown"`
	// Set the management state of the host. Default is `false`.
	Maintenance pulumi.BoolPtrOutput `pulumi:"maintenance"`
	// Password that will be used by vSphere to authenticate
	// to the host.
	Password pulumi.StringOutput `pulumi:"password"`
	// Host's certificate SHA-1 thumbprint. If not set the the
	// CA that signed the host's certificate should be trusted. If the CA is not trusted
	// and no thumbprint is set then the operation will fail.
	Thumbprint pulumi.StringPtrOutput `pulumi:"thumbprint"`
	// Username that will be used by vSphere to authenticate
	// to the host.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewHost registers a new resource with the given unique name, arguments, and options.
func NewHost(ctx *pulumi.Context,
	name string, args *HostArgs, opts ...pulumi.ResourceOption) (*Host, error) {
	if args == nil || args.Hostname == nil {
		return nil, errors.New("missing required argument 'Hostname'")
	}
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	if args == nil || args.Username == nil {
		return nil, errors.New("missing required argument 'Username'")
	}
	if args == nil {
		args = &HostArgs{}
	}
	var resource Host
	err := ctx.RegisterResource("vsphere:index/host:Host", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHost gets an existing Host resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostState, opts ...pulumi.ResourceOption) (*Host, error) {
	var resource Host
	err := ctx.ReadResource("vsphere:index/host:Host", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Host resources.
type hostState struct {
	// The ID of the Compute Cluster this host should
	// be added to. This should not be set if `datacenter` is set.
	Cluster *string `pulumi:"cluster"`
	// If set to false then the host will be disconected.
	// Default is `false`.
	Connected *bool `pulumi:"connected"`
	// The ID of the datacenter this host should
	// be added to. This should not be set if `cluster` is set.
	Datacenter *string `pulumi:"datacenter"`
	// If set to true then it will force the host to be added, even
	// if the host is already connected to a different vSphere instance. Default is `false`
	Force *bool `pulumi:"force"`
	// FQDN or IP address of the host to be added.
	Hostname *string `pulumi:"hostname"`
	// The license key that will be applied to the host.
	// The license key is expected to be present in vSphere.
	License *string `pulumi:"license"`
	// Set the lockdown state of the host. Valid options are
	// `disabled`, `normal`, and `strict`. Default is `disabled`.
	Lockdown *string `pulumi:"lockdown"`
	// Set the management state of the host. Default is `false`.
	Maintenance *bool `pulumi:"maintenance"`
	// Password that will be used by vSphere to authenticate
	// to the host.
	Password *string `pulumi:"password"`
	// Host's certificate SHA-1 thumbprint. If not set the the
	// CA that signed the host's certificate should be trusted. If the CA is not trusted
	// and no thumbprint is set then the operation will fail.
	Thumbprint *string `pulumi:"thumbprint"`
	// Username that will be used by vSphere to authenticate
	// to the host.
	Username *string `pulumi:"username"`
}

type HostState struct {
	// The ID of the Compute Cluster this host should
	// be added to. This should not be set if `datacenter` is set.
	Cluster pulumi.StringPtrInput
	// If set to false then the host will be disconected.
	// Default is `false`.
	Connected pulumi.BoolPtrInput
	// The ID of the datacenter this host should
	// be added to. This should not be set if `cluster` is set.
	Datacenter pulumi.StringPtrInput
	// If set to true then it will force the host to be added, even
	// if the host is already connected to a different vSphere instance. Default is `false`
	Force pulumi.BoolPtrInput
	// FQDN or IP address of the host to be added.
	Hostname pulumi.StringPtrInput
	// The license key that will be applied to the host.
	// The license key is expected to be present in vSphere.
	License pulumi.StringPtrInput
	// Set the lockdown state of the host. Valid options are
	// `disabled`, `normal`, and `strict`. Default is `disabled`.
	Lockdown pulumi.StringPtrInput
	// Set the management state of the host. Default is `false`.
	Maintenance pulumi.BoolPtrInput
	// Password that will be used by vSphere to authenticate
	// to the host.
	Password pulumi.StringPtrInput
	// Host's certificate SHA-1 thumbprint. If not set the the
	// CA that signed the host's certificate should be trusted. If the CA is not trusted
	// and no thumbprint is set then the operation will fail.
	Thumbprint pulumi.StringPtrInput
	// Username that will be used by vSphere to authenticate
	// to the host.
	Username pulumi.StringPtrInput
}

func (HostState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostState)(nil)).Elem()
}

type hostArgs struct {
	// The ID of the Compute Cluster this host should
	// be added to. This should not be set if `datacenter` is set.
	Cluster *string `pulumi:"cluster"`
	// If set to false then the host will be disconected.
	// Default is `false`.
	Connected *bool `pulumi:"connected"`
	// The ID of the datacenter this host should
	// be added to. This should not be set if `cluster` is set.
	Datacenter *string `pulumi:"datacenter"`
	// If set to true then it will force the host to be added, even
	// if the host is already connected to a different vSphere instance. Default is `false`
	Force *bool `pulumi:"force"`
	// FQDN or IP address of the host to be added.
	Hostname string `pulumi:"hostname"`
	// The license key that will be applied to the host.
	// The license key is expected to be present in vSphere.
	License *string `pulumi:"license"`
	// Set the lockdown state of the host. Valid options are
	// `disabled`, `normal`, and `strict`. Default is `disabled`.
	Lockdown *string `pulumi:"lockdown"`
	// Set the management state of the host. Default is `false`.
	Maintenance *bool `pulumi:"maintenance"`
	// Password that will be used by vSphere to authenticate
	// to the host.
	Password string `pulumi:"password"`
	// Host's certificate SHA-1 thumbprint. If not set the the
	// CA that signed the host's certificate should be trusted. If the CA is not trusted
	// and no thumbprint is set then the operation will fail.
	Thumbprint *string `pulumi:"thumbprint"`
	// Username that will be used by vSphere to authenticate
	// to the host.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a Host resource.
type HostArgs struct {
	// The ID of the Compute Cluster this host should
	// be added to. This should not be set if `datacenter` is set.
	Cluster pulumi.StringPtrInput
	// If set to false then the host will be disconected.
	// Default is `false`.
	Connected pulumi.BoolPtrInput
	// The ID of the datacenter this host should
	// be added to. This should not be set if `cluster` is set.
	Datacenter pulumi.StringPtrInput
	// If set to true then it will force the host to be added, even
	// if the host is already connected to a different vSphere instance. Default is `false`
	Force pulumi.BoolPtrInput
	// FQDN or IP address of the host to be added.
	Hostname pulumi.StringInput
	// The license key that will be applied to the host.
	// The license key is expected to be present in vSphere.
	License pulumi.StringPtrInput
	// Set the lockdown state of the host. Valid options are
	// `disabled`, `normal`, and `strict`. Default is `disabled`.
	Lockdown pulumi.StringPtrInput
	// Set the management state of the host. Default is `false`.
	Maintenance pulumi.BoolPtrInput
	// Password that will be used by vSphere to authenticate
	// to the host.
	Password pulumi.StringInput
	// Host's certificate SHA-1 thumbprint. If not set the the
	// CA that signed the host's certificate should be trusted. If the CA is not trusted
	// and no thumbprint is set then the operation will fail.
	Thumbprint pulumi.StringPtrInput
	// Username that will be used by vSphere to authenticate
	// to the host.
	Username pulumi.StringInput
}

func (HostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostArgs)(nil)).Elem()
}
