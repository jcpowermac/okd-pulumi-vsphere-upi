package ignition

import (
	"bytes"
	"text/template"

	"github.com/coreos/ignition/config/v2_1/types"
)

//todo
//data "ignition_file" "hostname" {

const ifcfgPath = "./ifcfg.tmpl"

//Ifcfg - interface config
type Ifcfg struct {
	DNS1      string
	DNS2      string
	IPAddress string
	Device    string
	Prefix    string
	Gateway   string
	Domain    string
}

//https://github.com/terraform-providers/terraform-provider-ignition/blob/b247e4a11c394f3d6a9b8f58fb139b6b2f149c42/ignition/resource_ignition_config.go#L173

func createIgnitionConfig() {
	i := types.Ignition{}
	i.Version = types.MaxVersion.String()
}

func createIgnitionFile(user, group, contents, path string, mode int) *types.File {
	return &types.File{
		FileEmbedded1: types.FileEmbedded1{
			Contents: types.FileContents{
				Source: contents,
			},
			Mode: mode,
		},
		Node: types.Node{
			Filesystem: "root",
			Path:       path,
			User: types.NodeUser{
				Name: user,
			},
			Group: types.NodeGroup{
				Name: group,
			},
		},
	}
}

// CreateIfcfg - create ifcfg
func CreateIfcfg(ifcfg *Ifcfg) *types.File {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("ifcfg").ParseFiles(ifcfgPath)

	if err != nil {
		return nil
	}

	err = tmpl.Execute(buf, ifcfg)
	if err != nil {
		return nil
	}

	return createIgnitionFile("root",
		"root",
		buf.String(),
		"/etc/sysconfig/network-scripts/ifcfg-ens192", 644)

	/*
		return &types.File{
			FileEmbedded1: types.FileEmbedded1{
				Contents: types.FileContents{
					Source: buf.String(),
				},
				Mode: 420,
			},
			Node: types.Node{
				Filesystem: "root",
				Path:       "/etc/sysconfig/network-scripts/ifcfg-ens192",
				User: types.NodeUser{
					Name: "root",
				},
				Group: types.NodeGroup{
					Name: "root",
				},
			},
		}
	*/
}
