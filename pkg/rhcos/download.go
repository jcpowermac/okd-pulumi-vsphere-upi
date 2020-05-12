package rhcos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type metadata struct {
	AMIs map[string]struct {
		HVM string `json:"hvm"`
	} `json:"amis"`
	Azure struct {
		Image string `json:"image"`
		URL   string `json:"url"`
	}
	GCP struct {
		Image string `json:"image"`
		URL   string `json:"url"`
	}
	BaseURI string `json:"baseURI"`
	Images  struct {
		QEMU struct {
			Path               string `json:"path"`
			SHA256             string `json:"sha256"`
			UncompressedSHA256 string `json:"uncompressed-sha256"`
		} `json:"qemu"`
		OpenStack struct {
			Path               string `json:"path"`
			SHA256             string `json:"sha256"`
			UncompressedSHA256 string `json:"uncompressed-sha256"`
		} `json:"openstack"`
		VMware struct {
			Path   string `json:"path"`
			SHA256 string `json:"sha256"`
		} `json:"vmware"`
	} `json:"images"`
	OSTreeVersion string `json:"ostree-version"`
}

func getRHCOSPerBranch(branch string) (*metadata, error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/openshift/installer/%s/data/data/rhcos.json", branch)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var meta *metadata
		if err := json.Unmarshal(bodyBytes, &meta); err != nil {
			return meta, errors.Wrap(err, "failed to parse RHCOS build metadata")
		}

		return meta, nil
	}

	return nil, errors.Errorf("failed to download RHCOS build metadata")
}
