package main

import (
	"okd-pulumi-vsphere-upi/pkg/rhcos"
	_ "okd-pulumi-vsphere-upi/pkg/vm"

	"github.com/davecgh/go-spew/spew"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {

	/* Order of execution (existing terraform)
	 ********************************
	 * 1. IPAM
	 * 2. Create vSphere objects RP, Folder
	 * 3. DNS
	 * 4. Load Balance
	 * 5. Bootstrap
	 * 6. Compute, Control Plane
	 */

	pulumi.Run(func(ctx *pulumi.Context) error {
		//conf := config.New(ctx, "")
		//ic := installer.NewInstallConfig(conf)

		// Write ic out to disk to be used with openshift-install
		//spew.Dump(ic)

		/*
			vm, err := pulumivm.NewVirtualMachine(ctx, ic)
			if err != nil {
				return err
			}
		*/

		meta, err := rhcos.GetRHCOSPerBranch("release-4.4")
		if err != nil {
			return err
		}
		filePath, err := rhcos.DownloadRHCOSImage(meta)
		if err != nil {
			return err
		}

		err = rhcos.ExtractImage(filePath)
		if err != nil {
			return err
		}
		spew.Dump(filePath)
		/*

			vm.CreateCoreOSTemplate(filePath, "4.4")

			err = vm.CreateCoreOSVirtualMachine(pulumivm.Bootstrap, 1, "foo")
			if err != nil {
				return err
			}
		*/

		return nil
	})
}
