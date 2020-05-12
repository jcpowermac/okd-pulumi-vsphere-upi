package ignition

import (
	"bytes"
	"text/template"

	"github.com/coreos/ignition/config/v2_1/types"
)

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
}
