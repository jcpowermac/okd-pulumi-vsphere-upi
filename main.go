package main

import (
	"github.com/davecgh/go-spew/spew"
	vsphere "github.com/pulumi/pulumi-vsphere/sdk/v2/go/vsphere"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {

		conf := config.New(ctx, "")

		datacenterName := conf.Require("datacenter")
		clusterName := conf.Require("cluster")
		datastoreName := conf.Require("datastore")

		dc, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{Name: &datacenterName})
		if err != nil {
			return err
		}

		if dc != nil && dc.Id != "" {
			args := new(vsphere.LookupComputeClusterArgs)
			args.DatacenterId = &dc.Id
			//args.Name = "/" + datacenterName + "/host/" + clusterName
			args.Name = clusterName

			computeCluster, err := vsphere.LookupComputeCluster(ctx, args, nil)

			//computeCluster, err := vsphere.GetComputeCluster(ctx, clusterName, nil, nil, nil)

			if err != nil {
				return err
			}
			spew.Dump(computeCluster)
		}

		spew.Dump(clusterName)

		datastore, err := vsphere.GetDatastore(ctx, &vsphere.GetDatastoreArgs{
			Name:         datastoreName,
			DatacenterId: &dc.Id,
		})

		if err != nil {
			return err
		}

		spew.Dump(datastore)

		return nil
	})
}
