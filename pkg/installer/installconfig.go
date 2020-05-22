package installer

import (
	installertypes "github.com/openshift/installer/pkg/types"
	vspheretypes "github.com/openshift/installer/pkg/types/vsphere"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

// NewInstallConfig - create a install-config since the struct already
// contains all the information we need
func NewInstallConfig(conf *config.Config) *installertypes.InstallConfig {
	ic := &installertypes.InstallConfig{}

	ic.APIVersion = "v1"
	ic.BaseDomain = conf.Require("basedomain")
	ic.SSHKey = conf.Require("sshkey")
	ic.PullSecret = conf.Require("pullsecret")

	ic.ObjectMeta.Name = conf.Require("clustername")
	ic.Platform.VSphere = &vspheretypes.Platform{
		Datacenter:       conf.Require("datacenter"),
		DefaultDatastore: conf.Require("datastore"),
		Folder:           conf.Require("folder"),
		Cluster:          conf.Require("cluster"),
		Network:          conf.Require("network"),
	}

	return ic
}
