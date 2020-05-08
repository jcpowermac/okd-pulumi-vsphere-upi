module okd-pulumi-vsphere-upi

go 1.13

require (
	github.com/ajeddeloh/go-json v0.0.0-20200220154158-5ae607161559 // indirect
	github.com/coreos/ignition v0.35.0
	github.com/davecgh/go-spew v1.1.1
	github.com/openshift/installer v0.9.0-master.0.20200507004117-189b16eba4e0
	github.com/pulumi/pulumi-vsphere/sdk/v2 v2.2.1
	github.com/pulumi/pulumi/sdk/v2 v2.0.0
	go4.org v0.0.0-20200411211856-f5505b9728dd // indirect
	golang.org/x/tools v0.0.0-20200507172948-ff647943c706 // indirect
)

replace (
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0 // Pin non-versioned import to v22.0.0
	github.com/metal3-io/baremetal-operator => github.com/openshift/baremetal-operator v0.0.0-20200206190020-71b826cc0f0a // Use OpenShift fork
	github.com/metal3-io/cluster-api-provider-baremetal => github.com/openshift/cluster-api-provider-baremetal v0.0.0-20190821174549-a2a477909c1d // Pin OpenShift fork
	github.com/openshift/api => github.com/openshift/api v0.0.0-20200413201024-c6e8c9b6eb9a // Pin API
	github.com/openshift/client-go => github.com/openshift/client-go v0.0.0-20200116152001-92a2713fa240 // Pin client-go
	github.com/terraform-providers/terraform-provider-azurerm => github.com/openshift/terraform-provider-azurerm v1.41.1-openshift-3 // Pin to openshift fork with IPv6 fixes
	go.etcd.io/etcd => go.etcd.io/etcd v0.0.0-20191023171146-3cf2f69b5738 // Pin to version used by k8s.io/apiserver
	k8s.io/client-go => k8s.io/client-go v0.18.2 // Pinned to keep from using an older v12.0.0 version that go mod thinks is newer
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200121204235-bf4fb3bd569c // Pin to version used by k8s.io/apiserver
	sigs.k8s.io/cluster-api-provider-aws => github.com/openshift/cluster-api-provider-aws v0.2.1-0.20200316201703-923caeb1d0d8 // Pin OpenShift fork
	sigs.k8s.io/cluster-api-provider-azure => github.com/openshift/cluster-api-provider-azure v0.1.0-alpha.3.0.20200120114645-8a9592f1f87b // Pin OpenShift fork
	sigs.k8s.io/cluster-api-provider-openstack => github.com/openshift/cluster-api-provider-openstack v0.0.0-20200323110431-3311de91e078 // Pin OpenShift fork
)
