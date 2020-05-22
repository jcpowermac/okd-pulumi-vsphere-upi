package rhcos

import (
	"archive/tar"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// Metadata - from the installer.
type Metadata struct {
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

// GetRHCOSPerBranch - get the current version of the rhcos.json
func GetRHCOSPerBranch(branch string) (*Metadata, error) {
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
		var meta *Metadata
		if err := json.Unmarshal(bodyBytes, &meta); err != nil {
			return meta, errors.Wrap(err, "failed to parse RHCOS build metadata")
		}

		return meta, nil
	}

	return nil, errors.Errorf("failed to download RHCOS build metadata")
}

// ExtractImage - extracts the OVA to access the ovf file
// https://medium.com/@skdomino/taring-untaring-files-in-go-6b07cf56bc07
func ExtractImage(filePath string) error {

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	tr := tar.NewReader(file)
	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		target := filepath.Join(wd, header.Name)

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}
}

// DownloadRHCOSImage - download the image and place in current working directory
func DownloadRHCOSImage(meta *Metadata) (string, error) {
	base, err := url.Parse(meta.BaseURI)
	if err != nil {
		return "", err
	}
	image, err := url.Parse(meta.Images.VMware.Path)
	if err != nil {
		return "", err
	}

	fullURL := fmt.Sprintf("%s?sha256=%s",
		base.ResolveReference(image).String(),
		meta.Images.VMware.SHA256)

	// Check that we have generated a valid URL
	_, err = url.ParseRequestURI(fullURL)
	if err != nil {
		return "", err
	}
	fileName := fmt.Sprintf("%x.ova", md5.Sum([]byte(fullURL)))

	wd, err := os.Getwd()
	if err != nil {
		return fileName, err
	}

	filePath := filepath.Join(wd, fileName)

	_, err = os.Stat(filePath)
	if err == nil {
		return filePath, nil
	}

	resp, err := http.Get(fullURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return "", errors.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(filePath)
	if err != nil {
		return fileName, err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
