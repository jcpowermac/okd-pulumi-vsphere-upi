// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `.DistributedPortGroup` resource can be used to manage vSphere
// distributed virtual port groups. These port groups are connected to distributed
// virtual switches, which can be managed by the
// `.DistributedVirtualSwitch` resource.
//
// Distributed port groups can be used as networks for virtual machines, allowing
// VMs to use the networking supplied by a distributed virtual switch (DVS), with
// a set of policies that apply to that individual newtork, if desired.
//
// For an overview on vSphere networking concepts, see [this
// page][ref-vsphere-net-concepts]. For more information on vSphere DVS
// portgroups, see [this page][ref-vsphere-dvportgroup].
//
// [ref-vsphere-net-concepts]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.networking.doc/GUID-2B11DBB8-CB3C-4AFF-8885-EFEA0FC562F4.html
// [ref-vsphere-dvportgroup]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.networking.doc/GUID-69933F6E-2442-46CF-AA17-1196CB9A0A09.html
//
// > **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
type DistributedPortGroup struct {
	pulumi.CustomResourceState

	// List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
	ActiveUplinks pulumi.StringArrayOutput `pulumi:"activeUplinks"`
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
	// that of its own.
	AllowForgedTransmits pulumi.BoolOutput `pulumi:"allowForgedTransmits"`
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges pulumi.BoolOutput `pulumi:"allowMacChanges"`
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous pulumi.BoolOutput `pulumi:"allowPromiscuous"`
	// Allows the port group to create additional ports
	// past the limit specified in `numberOfPorts` if necessary. Default: `true`.
	AutoExpand pulumi.BoolPtrOutput `pulumi:"autoExpand"`
	// Indicates whether to block all ports by default.
	BlockAllPorts pulumi.BoolOutput `pulumi:"blockAllPorts"`
	// Allow the port shutdown
	// policy to be overridden on an individual port.
	BlockOverrideAllowed pulumi.BoolPtrOutput `pulumi:"blockOverrideAllowed"`
	// Enable beacon probing on the ports this policy applies to.
	CheckBeacon pulumi.BoolOutput `pulumi:"checkBeacon"`
	// Version string of the configuration that this spec is trying to change.
	ConfigVersion pulumi.StringOutput `pulumi:"configVersion"`
	// Map of custom attribute ids to attribute
	// value string to set for port group.
	CustomAttributes pulumi.StringMapOutput `pulumi:"customAttributes"`
	// An optional description for the port group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Allow VMDirectPath Gen2 on the ports this policy applies to.
	DirectpathGen2Allowed pulumi.BoolOutput `pulumi:"directpathGen2Allowed"`
	// The ID of the DVS to add the
	// port group to. Forces a new resource if changed.
	DistributedVirtualSwitchUuid pulumi.StringOutput `pulumi:"distributedVirtualSwitchUuid"`
	// The average egress bandwidth in bits per second if egress shaping is enabled on the port.
	EgressShapingAverageBandwidth pulumi.IntOutput `pulumi:"egressShapingAverageBandwidth"`
	// The maximum egress burst size allowed in bytes if egress shaping is enabled on the port.
	EgressShapingBurstSize pulumi.IntOutput `pulumi:"egressShapingBurstSize"`
	// True if the traffic shaper is enabled for egress traffic on the port.
	EgressShapingEnabled pulumi.BoolOutput `pulumi:"egressShapingEnabled"`
	// The peak egress bandwidth during bursts in bits per second if egress traffic shaping is enabled on the port.
	EgressShapingPeakBandwidth pulumi.IntOutput `pulumi:"egressShapingPeakBandwidth"`
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback pulumi.BoolOutput `pulumi:"failback"`
	// The average ingress bandwidth in bits per second if ingress shaping is enabled on the port.
	IngressShapingAverageBandwidth pulumi.IntOutput `pulumi:"ingressShapingAverageBandwidth"`
	// The maximum ingress burst size allowed in bytes if ingress shaping is enabled on the port.
	IngressShapingBurstSize pulumi.IntOutput `pulumi:"ingressShapingBurstSize"`
	// True if the traffic shaper is enabled for ingress traffic on the port.
	IngressShapingEnabled pulumi.BoolOutput `pulumi:"ingressShapingEnabled"`
	// The peak ingress bandwidth during bursts in bits per second if ingress traffic shaping is enabled on the port.
	IngressShapingPeakBandwidth pulumi.IntOutput `pulumi:"ingressShapingPeakBandwidth"`
	// The generated UUID of the portgroup.
	Key pulumi.StringOutput `pulumi:"key"`
	// Whether or not to enable LACP on all uplink ports.
	LacpEnabled pulumi.BoolOutput `pulumi:"lacpEnabled"`
	// The uplink LACP mode to use. Can be one of active or passive.
	LacpMode pulumi.StringOutput `pulumi:"lacpMode"`
	// Allow a port in this port group to be
	// moved to another port group while it is connected.
	LivePortMovingAllowed pulumi.BoolPtrOutput `pulumi:"livePortMovingAllowed"`
	// The name of the port group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates whether to enable netflow on all ports.
	NetflowEnabled pulumi.BoolOutput `pulumi:"netflowEnabled"`
	// Allow the Netflow
	// policy on this port group to be overridden on an individual
	// port.
	NetflowOverrideAllowed pulumi.BoolPtrOutput `pulumi:"netflowOverrideAllowed"`
	// The key of a network resource pool
	// to associate with this port group. The default is `-1`, which implies no
	// association.
	NetworkResourcePoolKey pulumi.StringPtrOutput `pulumi:"networkResourcePoolKey"`
	// Allow the network
	// resource pool set on this port group to be overridden on an individual port.
	NetworkResourcePoolOverrideAllowed pulumi.BoolPtrOutput `pulumi:"networkResourcePoolOverrideAllowed"`
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches pulumi.BoolOutput `pulumi:"notifySwitches"`
	// The number of ports available on this port
	// group. Cannot be decreased below the amount of used ports on the port group.
	NumberOfPorts pulumi.IntOutput `pulumi:"numberOfPorts"`
	// Reset a port's settings to the
	// settings defined on this port group policy when the port disconnects.
	PortConfigResetAtDisconnect pulumi.BoolPtrOutput `pulumi:"portConfigResetAtDisconnect"`
	// An optional formatting policy for naming of
	// the ports in this port group. See the `portNameFormat` attribute listed
	// [here][ext-vsphere-portname-format] for details on the format syntax.
	PortNameFormat pulumi.StringPtrOutput `pulumi:"portNameFormat"`
	// The secondary VLAN ID for this port.
	PortPrivateSecondaryVlanId pulumi.IntOutput `pulumi:"portPrivateSecondaryVlanId"`
	// Allow the security policy
	// settings defined in this port group policy to be
	// overridden on an individual port.
	SecurityPolicyOverrideAllowed pulumi.BoolPtrOutput `pulumi:"securityPolicyOverrideAllowed"`
	// Allow the traffic shaping
	// options on this port group policy to be overridden
	// on an individual port.
	ShapingOverrideAllowed pulumi.BoolPtrOutput `pulumi:"shapingOverrideAllowed"`
	// List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
	StandbyUplinks pulumi.StringArrayOutput `pulumi:"standbyUplinks"`
	// A list of tag IDs to apply to this object.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid,
	// failover_explicit, or loadbalance_loadbased.
	TeamingPolicy pulumi.StringOutput `pulumi:"teamingPolicy"`
	// Allow any traffic filters on
	// this port group to be overridden on an individual port.
	TrafficFilterOverrideAllowed pulumi.BoolPtrOutput `pulumi:"trafficFilterOverrideAllowed"`
	// If true, a copy of packets sent to the switch will always be forwarded to an uplink in addition to the regular packet
	// forwarded done by the switch.
	TxUplink pulumi.BoolOutput `pulumi:"txUplink"`
	// The port group type. Can be one of `earlyBinding` (static
	// binding) or `ephemeral`. Default: `earlyBinding`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// Allow the uplink teaming
	// options on this port group to be overridden on an
	// individual port.
	UplinkTeamingOverrideAllowed pulumi.BoolPtrOutput `pulumi:"uplinkTeamingOverrideAllowed"`
	// The VLAN ID for single VLAN mode. 0 denotes no VLAN.
	VlanId pulumi.IntOutput `pulumi:"vlanId"`
	// Allow the VLAN settings
	// on this port group to be overridden on an individual port.
	VlanOverrideAllowed pulumi.BoolPtrOutput `pulumi:"vlanOverrideAllowed"`
	// The VLAN ID for single VLAN mode. 0 denotes no VLAN.
	VlanRanges DistributedPortGroupVlanRangeArrayOutput `pulumi:"vlanRanges"`
}

// NewDistributedPortGroup registers a new resource with the given unique name, arguments, and options.
func NewDistributedPortGroup(ctx *pulumi.Context,
	name string, args *DistributedPortGroupArgs, opts ...pulumi.ResourceOption) (*DistributedPortGroup, error) {
	if args == nil || args.DistributedVirtualSwitchUuid == nil {
		return nil, errors.New("missing required argument 'DistributedVirtualSwitchUuid'")
	}
	if args == nil {
		args = &DistributedPortGroupArgs{}
	}
	var resource DistributedPortGroup
	err := ctx.RegisterResource("vsphere:index/distributedPortGroup:DistributedPortGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDistributedPortGroup gets an existing DistributedPortGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDistributedPortGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DistributedPortGroupState, opts ...pulumi.ResourceOption) (*DistributedPortGroup, error) {
	var resource DistributedPortGroup
	err := ctx.ReadResource("vsphere:index/distributedPortGroup:DistributedPortGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DistributedPortGroup resources.
type distributedPortGroupState struct {
	// List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
	ActiveUplinks []string `pulumi:"activeUplinks"`
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
	// that of its own.
	AllowForgedTransmits *bool `pulumi:"allowForgedTransmits"`
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges *bool `pulumi:"allowMacChanges"`
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous *bool `pulumi:"allowPromiscuous"`
	// Allows the port group to create additional ports
	// past the limit specified in `numberOfPorts` if necessary. Default: `true`.
	AutoExpand *bool `pulumi:"autoExpand"`
	// Indicates whether to block all ports by default.
	BlockAllPorts *bool `pulumi:"blockAllPorts"`
	// Allow the port shutdown
	// policy to be overridden on an individual port.
	BlockOverrideAllowed *bool `pulumi:"blockOverrideAllowed"`
	// Enable beacon probing on the ports this policy applies to.
	CheckBeacon *bool `pulumi:"checkBeacon"`
	// Version string of the configuration that this spec is trying to change.
	ConfigVersion *string `pulumi:"configVersion"`
	// Map of custom attribute ids to attribute
	// value string to set for port group.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// An optional description for the port group.
	Description *string `pulumi:"description"`
	// Allow VMDirectPath Gen2 on the ports this policy applies to.
	DirectpathGen2Allowed *bool `pulumi:"directpathGen2Allowed"`
	// The ID of the DVS to add the
	// port group to. Forces a new resource if changed.
	DistributedVirtualSwitchUuid *string `pulumi:"distributedVirtualSwitchUuid"`
	// The average egress bandwidth in bits per second if egress shaping is enabled on the port.
	EgressShapingAverageBandwidth *int `pulumi:"egressShapingAverageBandwidth"`
	// The maximum egress burst size allowed in bytes if egress shaping is enabled on the port.
	EgressShapingBurstSize *int `pulumi:"egressShapingBurstSize"`
	// True if the traffic shaper is enabled for egress traffic on the port.
	EgressShapingEnabled *bool `pulumi:"egressShapingEnabled"`
	// The peak egress bandwidth during bursts in bits per second if egress traffic shaping is enabled on the port.
	EgressShapingPeakBandwidth *int `pulumi:"egressShapingPeakBandwidth"`
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback *bool `pulumi:"failback"`
	// The average ingress bandwidth in bits per second if ingress shaping is enabled on the port.
	IngressShapingAverageBandwidth *int `pulumi:"ingressShapingAverageBandwidth"`
	// The maximum ingress burst size allowed in bytes if ingress shaping is enabled on the port.
	IngressShapingBurstSize *int `pulumi:"ingressShapingBurstSize"`
	// True if the traffic shaper is enabled for ingress traffic on the port.
	IngressShapingEnabled *bool `pulumi:"ingressShapingEnabled"`
	// The peak ingress bandwidth during bursts in bits per second if ingress traffic shaping is enabled on the port.
	IngressShapingPeakBandwidth *int `pulumi:"ingressShapingPeakBandwidth"`
	// The generated UUID of the portgroup.
	Key *string `pulumi:"key"`
	// Whether or not to enable LACP on all uplink ports.
	LacpEnabled *bool `pulumi:"lacpEnabled"`
	// The uplink LACP mode to use. Can be one of active or passive.
	LacpMode *string `pulumi:"lacpMode"`
	// Allow a port in this port group to be
	// moved to another port group while it is connected.
	LivePortMovingAllowed *bool `pulumi:"livePortMovingAllowed"`
	// The name of the port group.
	Name *string `pulumi:"name"`
	// Indicates whether to enable netflow on all ports.
	NetflowEnabled *bool `pulumi:"netflowEnabled"`
	// Allow the Netflow
	// policy on this port group to be overridden on an individual
	// port.
	NetflowOverrideAllowed *bool `pulumi:"netflowOverrideAllowed"`
	// The key of a network resource pool
	// to associate with this port group. The default is `-1`, which implies no
	// association.
	NetworkResourcePoolKey *string `pulumi:"networkResourcePoolKey"`
	// Allow the network
	// resource pool set on this port group to be overridden on an individual port.
	NetworkResourcePoolOverrideAllowed *bool `pulumi:"networkResourcePoolOverrideAllowed"`
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches *bool `pulumi:"notifySwitches"`
	// The number of ports available on this port
	// group. Cannot be decreased below the amount of used ports on the port group.
	NumberOfPorts *int `pulumi:"numberOfPorts"`
	// Reset a port's settings to the
	// settings defined on this port group policy when the port disconnects.
	PortConfigResetAtDisconnect *bool `pulumi:"portConfigResetAtDisconnect"`
	// An optional formatting policy for naming of
	// the ports in this port group. See the `portNameFormat` attribute listed
	// [here][ext-vsphere-portname-format] for details on the format syntax.
	PortNameFormat *string `pulumi:"portNameFormat"`
	// The secondary VLAN ID for this port.
	PortPrivateSecondaryVlanId *int `pulumi:"portPrivateSecondaryVlanId"`
	// Allow the security policy
	// settings defined in this port group policy to be
	// overridden on an individual port.
	SecurityPolicyOverrideAllowed *bool `pulumi:"securityPolicyOverrideAllowed"`
	// Allow the traffic shaping
	// options on this port group policy to be overridden
	// on an individual port.
	ShapingOverrideAllowed *bool `pulumi:"shapingOverrideAllowed"`
	// List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
	StandbyUplinks []string `pulumi:"standbyUplinks"`
	// A list of tag IDs to apply to this object.
	Tags []string `pulumi:"tags"`
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid,
	// failover_explicit, or loadbalance_loadbased.
	TeamingPolicy *string `pulumi:"teamingPolicy"`
	// Allow any traffic filters on
	// this port group to be overridden on an individual port.
	TrafficFilterOverrideAllowed *bool `pulumi:"trafficFilterOverrideAllowed"`
	// If true, a copy of packets sent to the switch will always be forwarded to an uplink in addition to the regular packet
	// forwarded done by the switch.
	TxUplink *bool `pulumi:"txUplink"`
	// The port group type. Can be one of `earlyBinding` (static
	// binding) or `ephemeral`. Default: `earlyBinding`.
	Type *string `pulumi:"type"`
	// Allow the uplink teaming
	// options on this port group to be overridden on an
	// individual port.
	UplinkTeamingOverrideAllowed *bool `pulumi:"uplinkTeamingOverrideAllowed"`
	// The VLAN ID for single VLAN mode. 0 denotes no VLAN.
	VlanId *int `pulumi:"vlanId"`
	// Allow the VLAN settings
	// on this port group to be overridden on an individual port.
	VlanOverrideAllowed *bool `pulumi:"vlanOverrideAllowed"`
	// The VLAN ID for single VLAN mode. 0 denotes no VLAN.
	VlanRanges []DistributedPortGroupVlanRange `pulumi:"vlanRanges"`
}

type DistributedPortGroupState struct {
	// List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
	ActiveUplinks pulumi.StringArrayInput
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
	// that of its own.
	AllowForgedTransmits pulumi.BoolPtrInput
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges pulumi.BoolPtrInput
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous pulumi.BoolPtrInput
	// Allows the port group to create additional ports
	// past the limit specified in `numberOfPorts` if necessary. Default: `true`.
	AutoExpand pulumi.BoolPtrInput
	// Indicates whether to block all ports by default.
	BlockAllPorts pulumi.BoolPtrInput
	// Allow the port shutdown
	// policy to be overridden on an individual port.
	BlockOverrideAllowed pulumi.BoolPtrInput
	// Enable beacon probing on the ports this policy applies to.
	CheckBeacon pulumi.BoolPtrInput
	// Version string of the configuration that this spec is trying to change.
	ConfigVersion pulumi.StringPtrInput
	// Map of custom attribute ids to attribute
	// value string to set for port group.
	CustomAttributes pulumi.StringMapInput
	// An optional description for the port group.
	Description pulumi.StringPtrInput
	// Allow VMDirectPath Gen2 on the ports this policy applies to.
	DirectpathGen2Allowed pulumi.BoolPtrInput
	// The ID of the DVS to add the
	// port group to. Forces a new resource if changed.
	DistributedVirtualSwitchUuid pulumi.StringPtrInput
	// The average egress bandwidth in bits per second if egress shaping is enabled on the port.
	EgressShapingAverageBandwidth pulumi.IntPtrInput
	// The maximum egress burst size allowed in bytes if egress shaping is enabled on the port.
	EgressShapingBurstSize pulumi.IntPtrInput
	// True if the traffic shaper is enabled for egress traffic on the port.
	EgressShapingEnabled pulumi.BoolPtrInput
	// The peak egress bandwidth during bursts in bits per second if egress traffic shaping is enabled on the port.
	EgressShapingPeakBandwidth pulumi.IntPtrInput
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback pulumi.BoolPtrInput
	// The average ingress bandwidth in bits per second if ingress shaping is enabled on the port.
	IngressShapingAverageBandwidth pulumi.IntPtrInput
	// The maximum ingress burst size allowed in bytes if ingress shaping is enabled on the port.
	IngressShapingBurstSize pulumi.IntPtrInput
	// True if the traffic shaper is enabled for ingress traffic on the port.
	IngressShapingEnabled pulumi.BoolPtrInput
	// The peak ingress bandwidth during bursts in bits per second if ingress traffic shaping is enabled on the port.
	IngressShapingPeakBandwidth pulumi.IntPtrInput
	// The generated UUID of the portgroup.
	Key pulumi.StringPtrInput
	// Whether or not to enable LACP on all uplink ports.
	LacpEnabled pulumi.BoolPtrInput
	// The uplink LACP mode to use. Can be one of active or passive.
	LacpMode pulumi.StringPtrInput
	// Allow a port in this port group to be
	// moved to another port group while it is connected.
	LivePortMovingAllowed pulumi.BoolPtrInput
	// The name of the port group.
	Name pulumi.StringPtrInput
	// Indicates whether to enable netflow on all ports.
	NetflowEnabled pulumi.BoolPtrInput
	// Allow the Netflow
	// policy on this port group to be overridden on an individual
	// port.
	NetflowOverrideAllowed pulumi.BoolPtrInput
	// The key of a network resource pool
	// to associate with this port group. The default is `-1`, which implies no
	// association.
	NetworkResourcePoolKey pulumi.StringPtrInput
	// Allow the network
	// resource pool set on this port group to be overridden on an individual port.
	NetworkResourcePoolOverrideAllowed pulumi.BoolPtrInput
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches pulumi.BoolPtrInput
	// The number of ports available on this port
	// group. Cannot be decreased below the amount of used ports on the port group.
	NumberOfPorts pulumi.IntPtrInput
	// Reset a port's settings to the
	// settings defined on this port group policy when the port disconnects.
	PortConfigResetAtDisconnect pulumi.BoolPtrInput
	// An optional formatting policy for naming of
	// the ports in this port group. See the `portNameFormat` attribute listed
	// [here][ext-vsphere-portname-format] for details on the format syntax.
	PortNameFormat pulumi.StringPtrInput
	// The secondary VLAN ID for this port.
	PortPrivateSecondaryVlanId pulumi.IntPtrInput
	// Allow the security policy
	// settings defined in this port group policy to be
	// overridden on an individual port.
	SecurityPolicyOverrideAllowed pulumi.BoolPtrInput
	// Allow the traffic shaping
	// options on this port group policy to be overridden
	// on an individual port.
	ShapingOverrideAllowed pulumi.BoolPtrInput
	// List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
	StandbyUplinks pulumi.StringArrayInput
	// A list of tag IDs to apply to this object.
	Tags pulumi.StringArrayInput
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid,
	// failover_explicit, or loadbalance_loadbased.
	TeamingPolicy pulumi.StringPtrInput
	// Allow any traffic filters on
	// this port group to be overridden on an individual port.
	TrafficFilterOverrideAllowed pulumi.BoolPtrInput
	// If true, a copy of packets sent to the switch will always be forwarded to an uplink in addition to the regular packet
	// forwarded done by the switch.
	TxUplink pulumi.BoolPtrInput
	// The port group type. Can be one of `earlyBinding` (static
	// binding) or `ephemeral`. Default: `earlyBinding`.
	Type pulumi.StringPtrInput
	// Allow the uplink teaming
	// options on this port group to be overridden on an
	// individual port.
	UplinkTeamingOverrideAllowed pulumi.BoolPtrInput
	// The VLAN ID for single VLAN mode. 0 denotes no VLAN.
	VlanId pulumi.IntPtrInput
	// Allow the VLAN settings
	// on this port group to be overridden on an individual port.
	VlanOverrideAllowed pulumi.BoolPtrInput
	// The VLAN ID for single VLAN mode. 0 denotes no VLAN.
	VlanRanges DistributedPortGroupVlanRangeArrayInput
}

func (DistributedPortGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*distributedPortGroupState)(nil)).Elem()
}

type distributedPortGroupArgs struct {
	// List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
	ActiveUplinks []string `pulumi:"activeUplinks"`
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
	// that of its own.
	AllowForgedTransmits *bool `pulumi:"allowForgedTransmits"`
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges *bool `pulumi:"allowMacChanges"`
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous *bool `pulumi:"allowPromiscuous"`
	// Allows the port group to create additional ports
	// past the limit specified in `numberOfPorts` if necessary. Default: `true`.
	AutoExpand *bool `pulumi:"autoExpand"`
	// Indicates whether to block all ports by default.
	BlockAllPorts *bool `pulumi:"blockAllPorts"`
	// Allow the port shutdown
	// policy to be overridden on an individual port.
	BlockOverrideAllowed *bool `pulumi:"blockOverrideAllowed"`
	// Enable beacon probing on the ports this policy applies to.
	CheckBeacon *bool `pulumi:"checkBeacon"`
	// Map of custom attribute ids to attribute
	// value string to set for port group.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// An optional description for the port group.
	Description *string `pulumi:"description"`
	// Allow VMDirectPath Gen2 on the ports this policy applies to.
	DirectpathGen2Allowed *bool `pulumi:"directpathGen2Allowed"`
	// The ID of the DVS to add the
	// port group to. Forces a new resource if changed.
	DistributedVirtualSwitchUuid string `pulumi:"distributedVirtualSwitchUuid"`
	// The average egress bandwidth in bits per second if egress shaping is enabled on the port.
	EgressShapingAverageBandwidth *int `pulumi:"egressShapingAverageBandwidth"`
	// The maximum egress burst size allowed in bytes if egress shaping is enabled on the port.
	EgressShapingBurstSize *int `pulumi:"egressShapingBurstSize"`
	// True if the traffic shaper is enabled for egress traffic on the port.
	EgressShapingEnabled *bool `pulumi:"egressShapingEnabled"`
	// The peak egress bandwidth during bursts in bits per second if egress traffic shaping is enabled on the port.
	EgressShapingPeakBandwidth *int `pulumi:"egressShapingPeakBandwidth"`
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback *bool `pulumi:"failback"`
	// The average ingress bandwidth in bits per second if ingress shaping is enabled on the port.
	IngressShapingAverageBandwidth *int `pulumi:"ingressShapingAverageBandwidth"`
	// The maximum ingress burst size allowed in bytes if ingress shaping is enabled on the port.
	IngressShapingBurstSize *int `pulumi:"ingressShapingBurstSize"`
	// True if the traffic shaper is enabled for ingress traffic on the port.
	IngressShapingEnabled *bool `pulumi:"ingressShapingEnabled"`
	// The peak ingress bandwidth during bursts in bits per second if ingress traffic shaping is enabled on the port.
	IngressShapingPeakBandwidth *int `pulumi:"ingressShapingPeakBandwidth"`
	// Whether or not to enable LACP on all uplink ports.
	LacpEnabled *bool `pulumi:"lacpEnabled"`
	// The uplink LACP mode to use. Can be one of active or passive.
	LacpMode *string `pulumi:"lacpMode"`
	// Allow a port in this port group to be
	// moved to another port group while it is connected.
	LivePortMovingAllowed *bool `pulumi:"livePortMovingAllowed"`
	// The name of the port group.
	Name *string `pulumi:"name"`
	// Indicates whether to enable netflow on all ports.
	NetflowEnabled *bool `pulumi:"netflowEnabled"`
	// Allow the Netflow
	// policy on this port group to be overridden on an individual
	// port.
	NetflowOverrideAllowed *bool `pulumi:"netflowOverrideAllowed"`
	// The key of a network resource pool
	// to associate with this port group. The default is `-1`, which implies no
	// association.
	NetworkResourcePoolKey *string `pulumi:"networkResourcePoolKey"`
	// Allow the network
	// resource pool set on this port group to be overridden on an individual port.
	NetworkResourcePoolOverrideAllowed *bool `pulumi:"networkResourcePoolOverrideAllowed"`
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches *bool `pulumi:"notifySwitches"`
	// The number of ports available on this port
	// group. Cannot be decreased below the amount of used ports on the port group.
	NumberOfPorts *int `pulumi:"numberOfPorts"`
	// Reset a port's settings to the
	// settings defined on this port group policy when the port disconnects.
	PortConfigResetAtDisconnect *bool `pulumi:"portConfigResetAtDisconnect"`
	// An optional formatting policy for naming of
	// the ports in this port group. See the `portNameFormat` attribute listed
	// [here][ext-vsphere-portname-format] for details on the format syntax.
	PortNameFormat *string `pulumi:"portNameFormat"`
	// The secondary VLAN ID for this port.
	PortPrivateSecondaryVlanId *int `pulumi:"portPrivateSecondaryVlanId"`
	// Allow the security policy
	// settings defined in this port group policy to be
	// overridden on an individual port.
	SecurityPolicyOverrideAllowed *bool `pulumi:"securityPolicyOverrideAllowed"`
	// Allow the traffic shaping
	// options on this port group policy to be overridden
	// on an individual port.
	ShapingOverrideAllowed *bool `pulumi:"shapingOverrideAllowed"`
	// List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
	StandbyUplinks []string `pulumi:"standbyUplinks"`
	// A list of tag IDs to apply to this object.
	Tags []string `pulumi:"tags"`
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid,
	// failover_explicit, or loadbalance_loadbased.
	TeamingPolicy *string `pulumi:"teamingPolicy"`
	// Allow any traffic filters on
	// this port group to be overridden on an individual port.
	TrafficFilterOverrideAllowed *bool `pulumi:"trafficFilterOverrideAllowed"`
	// If true, a copy of packets sent to the switch will always be forwarded to an uplink in addition to the regular packet
	// forwarded done by the switch.
	TxUplink *bool `pulumi:"txUplink"`
	// The port group type. Can be one of `earlyBinding` (static
	// binding) or `ephemeral`. Default: `earlyBinding`.
	Type *string `pulumi:"type"`
	// Allow the uplink teaming
	// options on this port group to be overridden on an
	// individual port.
	UplinkTeamingOverrideAllowed *bool `pulumi:"uplinkTeamingOverrideAllowed"`
	// The VLAN ID for single VLAN mode. 0 denotes no VLAN.
	VlanId *int `pulumi:"vlanId"`
	// Allow the VLAN settings
	// on this port group to be overridden on an individual port.
	VlanOverrideAllowed *bool `pulumi:"vlanOverrideAllowed"`
	// The VLAN ID for single VLAN mode. 0 denotes no VLAN.
	VlanRanges []DistributedPortGroupVlanRange `pulumi:"vlanRanges"`
}

// The set of arguments for constructing a DistributedPortGroup resource.
type DistributedPortGroupArgs struct {
	// List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
	ActiveUplinks pulumi.StringArrayInput
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
	// that of its own.
	AllowForgedTransmits pulumi.BoolPtrInput
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges pulumi.BoolPtrInput
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous pulumi.BoolPtrInput
	// Allows the port group to create additional ports
	// past the limit specified in `numberOfPorts` if necessary. Default: `true`.
	AutoExpand pulumi.BoolPtrInput
	// Indicates whether to block all ports by default.
	BlockAllPorts pulumi.BoolPtrInput
	// Allow the port shutdown
	// policy to be overridden on an individual port.
	BlockOverrideAllowed pulumi.BoolPtrInput
	// Enable beacon probing on the ports this policy applies to.
	CheckBeacon pulumi.BoolPtrInput
	// Map of custom attribute ids to attribute
	// value string to set for port group.
	CustomAttributes pulumi.StringMapInput
	// An optional description for the port group.
	Description pulumi.StringPtrInput
	// Allow VMDirectPath Gen2 on the ports this policy applies to.
	DirectpathGen2Allowed pulumi.BoolPtrInput
	// The ID of the DVS to add the
	// port group to. Forces a new resource if changed.
	DistributedVirtualSwitchUuid pulumi.StringInput
	// The average egress bandwidth in bits per second if egress shaping is enabled on the port.
	EgressShapingAverageBandwidth pulumi.IntPtrInput
	// The maximum egress burst size allowed in bytes if egress shaping is enabled on the port.
	EgressShapingBurstSize pulumi.IntPtrInput
	// True if the traffic shaper is enabled for egress traffic on the port.
	EgressShapingEnabled pulumi.BoolPtrInput
	// The peak egress bandwidth during bursts in bits per second if egress traffic shaping is enabled on the port.
	EgressShapingPeakBandwidth pulumi.IntPtrInput
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback pulumi.BoolPtrInput
	// The average ingress bandwidth in bits per second if ingress shaping is enabled on the port.
	IngressShapingAverageBandwidth pulumi.IntPtrInput
	// The maximum ingress burst size allowed in bytes if ingress shaping is enabled on the port.
	IngressShapingBurstSize pulumi.IntPtrInput
	// True if the traffic shaper is enabled for ingress traffic on the port.
	IngressShapingEnabled pulumi.BoolPtrInput
	// The peak ingress bandwidth during bursts in bits per second if ingress traffic shaping is enabled on the port.
	IngressShapingPeakBandwidth pulumi.IntPtrInput
	// Whether or not to enable LACP on all uplink ports.
	LacpEnabled pulumi.BoolPtrInput
	// The uplink LACP mode to use. Can be one of active or passive.
	LacpMode pulumi.StringPtrInput
	// Allow a port in this port group to be
	// moved to another port group while it is connected.
	LivePortMovingAllowed pulumi.BoolPtrInput
	// The name of the port group.
	Name pulumi.StringPtrInput
	// Indicates whether to enable netflow on all ports.
	NetflowEnabled pulumi.BoolPtrInput
	// Allow the Netflow
	// policy on this port group to be overridden on an individual
	// port.
	NetflowOverrideAllowed pulumi.BoolPtrInput
	// The key of a network resource pool
	// to associate with this port group. The default is `-1`, which implies no
	// association.
	NetworkResourcePoolKey pulumi.StringPtrInput
	// Allow the network
	// resource pool set on this port group to be overridden on an individual port.
	NetworkResourcePoolOverrideAllowed pulumi.BoolPtrInput
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches pulumi.BoolPtrInput
	// The number of ports available on this port
	// group. Cannot be decreased below the amount of used ports on the port group.
	NumberOfPorts pulumi.IntPtrInput
	// Reset a port's settings to the
	// settings defined on this port group policy when the port disconnects.
	PortConfigResetAtDisconnect pulumi.BoolPtrInput
	// An optional formatting policy for naming of
	// the ports in this port group. See the `portNameFormat` attribute listed
	// [here][ext-vsphere-portname-format] for details on the format syntax.
	PortNameFormat pulumi.StringPtrInput
	// The secondary VLAN ID for this port.
	PortPrivateSecondaryVlanId pulumi.IntPtrInput
	// Allow the security policy
	// settings defined in this port group policy to be
	// overridden on an individual port.
	SecurityPolicyOverrideAllowed pulumi.BoolPtrInput
	// Allow the traffic shaping
	// options on this port group policy to be overridden
	// on an individual port.
	ShapingOverrideAllowed pulumi.BoolPtrInput
	// List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
	StandbyUplinks pulumi.StringArrayInput
	// A list of tag IDs to apply to this object.
	Tags pulumi.StringArrayInput
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid,
	// failover_explicit, or loadbalance_loadbased.
	TeamingPolicy pulumi.StringPtrInput
	// Allow any traffic filters on
	// this port group to be overridden on an individual port.
	TrafficFilterOverrideAllowed pulumi.BoolPtrInput
	// If true, a copy of packets sent to the switch will always be forwarded to an uplink in addition to the regular packet
	// forwarded done by the switch.
	TxUplink pulumi.BoolPtrInput
	// The port group type. Can be one of `earlyBinding` (static
	// binding) or `ephemeral`. Default: `earlyBinding`.
	Type pulumi.StringPtrInput
	// Allow the uplink teaming
	// options on this port group to be overridden on an
	// individual port.
	UplinkTeamingOverrideAllowed pulumi.BoolPtrInput
	// The VLAN ID for single VLAN mode. 0 denotes no VLAN.
	VlanId pulumi.IntPtrInput
	// Allow the VLAN settings
	// on this port group to be overridden on an individual port.
	VlanOverrideAllowed pulumi.BoolPtrInput
	// The VLAN ID for single VLAN mode. 0 denotes no VLAN.
	VlanRanges DistributedPortGroupVlanRangeArrayInput
}

func (DistributedPortGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*distributedPortGroupArgs)(nil)).Elem()
}
