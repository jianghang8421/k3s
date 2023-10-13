// Code generated for package deploy by go-bindata DO NOT EDIT. (@generated)
// sources:
// manifests/ccm.yaml
// manifests/coredns.yaml
// manifests/local-storage.yaml
// manifests/metrics-server/aggregated-metrics-reader.yaml
// manifests/metrics-server/auth-delegator.yaml
// manifests/metrics-server/auth-reader.yaml
// manifests/metrics-server/metrics-apiservice.yaml
// manifests/metrics-server/metrics-server-deployment.yaml
// manifests/metrics-server/metrics-server-service.yaml
// manifests/metrics-server/resource-reader.yaml
// manifests/rolebindings.yaml
// manifests/traefik.yaml
//go:build !no_stage
// +build !no_stage

package deploy

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _ccmYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x94\x41\x8f\x13\x31\x0c\x85\xef\xf3\x2b\xa2\x1e\x91\xd2\x15\xe2\x82\xe6\x08\x07\xee\x2b\xc1\xdd\x4d\x1e\xdd\xd0\x4c\x1c\xd9\x4e\x61\xf9\xf5\x68\x3a\x5d\x51\x3a\xdb\xd2\x16\x10\x7b\x8b\xac\xf8\xf3\xf3\x73\x62\xaa\xe9\x13\x44\x13\x97\xde\xc9\x8a\xc2\x92\x9a\x3d\xb0\xa4\xef\x64\x89\xcb\x72\xf3\x56\x97\x89\xef\xb6\xaf\xbb\x4d\x2a\xb1\x77\xef\x73\x53\x83\xdc\x73\x46\x37\xc0\x28\x92\x51\xdf\x39\x57\x68\x40\xef\x36\x6f\xd4\x87\xcc\x2d\xfa\xc0\xc5\x84\x73\x86\xf8\x81\x0a\xad\x21\x9d\xb4\x0c\xed\x3b\xef\xa8\xa6\x0f\xc2\xad\xea\x98\xe8\x5d\x60\x96\x98\xca\x61\xbd\xce\x39\x81\x72\x93\x80\xfd\xa5\x0c\x52\x68\xe7\xdc\x16\xb2\xda\xc7\xd6\xb0\x09\x20\x20\xc3\xee\xd8\x6a\x1c\x8f\xb3\x1a\x8b\xc5\x1c\x89\x2d\x8a\x1d\x21\x0f\x50\x95\x2c\x3c\x5c\x0d\x2d\x1c\x8f\x65\x2e\x5e\x2d\xae\xc8\xbd\x53\x23\x6b\xba\x0b\x28\x64\x9b\xc2\x61\xec\x00\x3b\xe9\xbb\x08\xfc\xc4\x99\xf2\x38\x9e\xf0\x31\x27\x9d\x0e\x5f\x6f\x42\xcf\xb4\x5d\xeb\xdd\x9e\x45\x21\x70\x3b\x37\x99\x51\xef\x65\x86\xd2\x00\xad\x34\x93\xf7\x3b\x16\xd5\xaa\x73\x5a\x24\x0c\x5c\x14\xc7\xca\x9e\x9f\x6f\x4c\x1a\x78\x0b\x79\xdc\x3f\xe9\xe7\x1e\x60\x89\x95\x53\x31\xcd\x73\x07\x4f\xcd\xc4\xfb\xee\xf6\x1f\xfb\x2e\x95\x98\xca\xfa\xea\x8f\xcb\x19\xf7\xf8\x3c\xde\x7e\xea\xf2\x4c\xe5\xce\xb9\xf9\xaa\xb8\xa8\x8e\xb6\xd5\x17\x04\xdb\xed\x88\x09\xf1\x51\x21\x97\xe5\xba\x9f\xc3\xee\xdd\xa6\xad\xe0\xf5\x51\x0d\xc3\x7f\x71\xcc\x8f\x7c\x1f\x91\xb1\x26\xe3\xbf\x6a\xe0\xd4\x55\x7f\x54\xe0\xa5\x38\xf7\x87\x96\xa1\x58\x0a\x3b\xb2\x17\x50\x3c\x27\xee\x46\x4b\x7f\xf1\x12\xdf\x0c\x65\xec\xcd\x53\x4d\xe3\xf2\x39\x29\xe3\x9f\xf8\xfb\x23\x00\x00\xff\xff\xde\xc0\x02\x82\x7a\x07\x00\x00")

func ccmYamlBytes() ([]byte, error) {
	return bindataRead(
		_ccmYaml,
		"ccm.yaml",
	)
}

func ccmYaml() (*asset, error) {
	bytes, err := ccmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ccm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _corednsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xdf\x6f\xe3\xbe\x0d\x7f\xcf\x5f\x21\x18\xe8\xcb\x30\xa7\xcd\x8a\xbb\x75\x7a\xeb\x25\xe9\x5d\xb0\x36\x17\x24\xe9\x01\xc5\x30\x14\x8a\xcc\xc4\x5a\x65\x51\x93\xe4\xb4\xd9\xad\xff\xfb\x20\xff\x8a\x95\xb8\xbd\xb6\xbb\x6f\x5e\x62\x9b\x22\x29\x7e\x44\x7e\x48\x31\x2d\x7e\x80\xb1\x02\x15\x25\xdb\x41\xef\x41\xa8\x84\x92\x05\x98\xad\xe0\x70\xc9\x39\xe6\xca\xf5\x32\x70\x2c\x61\x8e\xd1\x1e\x21\x8a\x65\x40\x09\x47\x03\x89\xb2\xd5\xbb\xd5\x8c\x03\x25\x0f\xf9\x0a\x62\xbb\xb3\x0e\xb2\x5e\x1c\xc7\xbd\xb6\x69\xb3\x62\xbc\xcf\x72\x97\xa2\x11\xff\x61\x4e\xa0\xea\x3f\x5c\xd8\xbe\xc0\xd3\xc6\xe9\x50\xe6\xd6\x81\x99\xa3\x84\xc0\xa3\x64\x2b\x90\xd6\x3f\x91\xc2\x85\x51\xe0\xa0\x50\x5d\x21\x3a\xeb\x0c\xd3\x5a\xa8\x4d\xe9\x23\x4e\x60\xcd\x72\xe9\x6c\xb3\xd5\x72\x43\xb4\xde\xb1\xc9\x25\x58\xda\x8b\x09\xd3\xe2\xab\xc1\x5c\x17\x96\x63\x12\x45\x3d\x42\x0c\x58\xcc\x0d\x87\xea\x1b\xa8\x44\xa3\x50\x85\xb1\x98\xd8\x12\x94\xf2\x45\x63\x52\x3e\x34\xf1\xfb\xd7\x2d\x98\x55\xa5\x2b\x85\x75\xc5\xc3\x23\x73\x3c\x3d\xf6\x97\x08\xcb\x71\x0b\x66\x57\xe1\xf0\x8a\x77\x29\x7e\x69\xfd\xff\x42\xfb\x8b\x50\x89\x50\x9b\x00\x74\xa6\x14\xba\x42\xb3\x42\xbe\xcb\x64\x70\x18\x2c\x77\x98\xeb\x84\x39\xa0\x24\x72\x26\x87\xe8\xf7\x9f\x1d\x4a\x98\xc3\xba\xd8\x5f\x85\xe6\x2b\xb1\xf6\x08\x39\x4e\xac\x17\x2c\xdb\x7c\xf5\x2f\xe0\xae\x48\x8c\xce\x12\xf8\x70\xe2\xef\x01\x47\xb5\x16\x9b\x1b\xa6\x3f\x52\x4e\xf5\xf2\x21\x1a\x58\x0b\x09\x94\xfc\xb7\xc0\xb4\x4f\x3f\x9d\x93\x9f\xc5\xa3\xff\x81\x31\x68\x6c\xf3\x9a\x02\x93\x2e\x6d\x5e\x0d\xb0\x64\xd7\xbc\xed\x8f\x83\x9c\xfc\x1c\x5e\xdf\x2e\x96\xe3\xf9\xfd\xe8\xfb\xcd\xe5\x64\xfa\x7c\x42\x84\x8a\x59\x92\x98\x3e\x33\x9a\x11\xa1\x3f\x97\x0f\x7b\x4f\xa4\xa8\x00\x22\x94\x05\x9e\x1b\x68\x7d\x5f\x33\x29\x5d\x6a\x30\xdf\xa4\xdd\x56\x9a\xb5\xcf\xfb\x8d\xa2\x75\x96\x9c\x82\xe3\xa7\x15\x14\xa7\x53\x4c\xe0\x5b\xf1\xb9\xed\xd4\x39\x49\x3e\x9f\xb5\x3e\x18\x90\xc8\x12\x32\xf8\x64\xbb\xb7\xd0\xe1\x4c\x1b\xcc\xc0\xa5\x90\x5b\x42\xff\x36\xf8\x74\xde\x08\xd6\x68\x1e\x99\x49\x48\xbf\xdc\x89\x2f\x47\xb9\xed\x73\x54\xeb\x66\x09\x67\x3c\x05\x72\xbe\xdf\x81\x44\xd4\xbd\x70\x33\x2d\x19\x4b\x56\x4c\x32\xc5\xf7\xf8\x88\x4c\xa3\x71\x61\xa8\x3c\xb7\x0e\xb3\xd3\x3f\xf5\x3d\x1f\x18\x91\x94\xab\xcb\x0d\xbf\xba\xde\x33\x12\x98\xa3\x94\x63\x5a\xdb\x7d\xa1\x8f\x40\x4b\xdc\x65\xf0\x31\x1e\x3f\x28\xe1\x0b\x1b\x33\xad\xab\x25\xa5\xe2\x61\x61\x97\x86\x23\x9f\xa9\xa3\xe9\x22\xea\x59\x0d\x9c\x16\xec\xb6\x15\x7e\x7f\xdf\x84\x75\x68\x76\xd7\x22\x13\x8e\x12\x8f\xa4\xa7\x01\x07\x9b\x5d\xe9\xc3\xed\x34\x50\x32\x47\x29\x85\xda\xdc\x16\x84\x52\x12\x50\xfb\x0b\xad\x00\xcd\xd8\xd3\xad\x62\x5b\x26\x24\x5b\xf9\xaa\x18\x78\x73\x20\x81\x3b\x34\xe5\x9a\xcc\x13\xe4\x75\x2b\x86\xee\x28\x1c\x64\x5a\x36\x86\xdb\x40\x15\x27\x19\xe8\xbf\x84\x43\x1d\x69\x99\x64\x02\x8d\x70\xbb\xa1\x64\xd6\x4e\x4b\x48\x4a\x48\x63\x5e\xd2\x51\xcc\x8d\x70\x82\x33\x19\x55\x2a\x36\x60\x9c\xe9\xc1\xf9\x14\xd0\xa0\x04\xd3\x26\x65\xff\x8b\xc9\x03\xec\x3c\xe0\x95\xb9\xcb\x24\x41\x65\xbf\x2b\xb9\x8b\x5a\x25\x81\xda\x6b\xa2\xa1\x24\x1a\x3f\x09\xeb\x6c\x74\x64\x40\x61\x02\xb1\xa7\xd8\x03\x62\xe7\xa8\x9c\x41\x19\x6b\xc9\x14\xbc\xd1\x26\x21\xb0\x5e\x03\x77\x94\x44\x53\x5c\xf0\x14\x92\x5c\xc2\xdb\x5d\x66\xcc\x23\xf4\x3b\x7c\x79\x0f\x8b\x20\x21\x8e\x33\x16\x2d\x25\x52\xa8\xfc\xa9\x81\x59\xa3\xc4\xcd\x6e\xa1\x3d\x63\x0e\x51\xf9\x04\xf5\x8d\xb8\x0d\x7a\xc6\x9e\x16\x0f\xf0\x58\xa6\x5c\xfd\xab\x35\xff\xee\xa3\x0b\x9d\x78\x8a\xf3\xa5\xd1\x5a\xfd\x98\x82\xba\x55\x96\x39\x61\xd7\xa2\xcc\xdf\x11\x4e\xd1\xd5\x31\xb4\x96\x16\x09\x78\x1c\xc7\x0b\x09\xfe\x7a\x9a\x12\xe2\x4f\x94\x09\x05\xa6\xd1\x88\x8f\xf8\xa0\xfc\x89\x8c\x6d\x80\x92\x93\x9f\x8b\xbb\xc5\x72\x7c\x73\x3f\x1a\x5f\x5d\xde\x5e\x2f\xef\xe7\xe3\xaf\x93\xc5\x72\x7e\xf7\x7c\x62\x98\xe2\x29\x98\xd3\x4c\xf8\xde\x03\x49\x5c\x99\xa8\xff\xe9\xa0\x3f\x38\xeb\x0f\x42\x8b\xb3\x5c\xca\x19\x4a\xc1\x77\x94\x4c\xd6\x53\x74\x33\x03\x16\x8a\x36\x5b\xfe\x82\x51\xa8\x01\xc1\x53\xc6\x41\x90\x19\x64\x68\x76\x94\x0c\xfe\x7a\x76\x23\x82\xbe\xf0\xef\x1c\xec\xe1\x6a\xae\x73\x4a\x06\x67\x67\x59\xa7\x8d\xc0\x04\x33\x1b\x4b\xc9\x3f\x48\x14\xfb\x06\x10\xfd\x99\x44\x01\x07\xd7\x8d\x38\x22\xff\x6c\x54\xb6\x28\xf3\x0c\x6e\x7c\xf5\x06\xa9\x52\x43\xeb\xfb\x7f\x5c\x2e\x6a\xf9\xcf\xfc\xfa\x19\x73\x29\x0d\x58\x3e\x88\x85\x25\xbe\x9e\x29\xf1\x63\xd5\xb1\xe1\xa2\x1d\xc4\xef\xb4\x5f\x75\x91\x5f\xbb\xf1\xfd\x27\x08\xa7\xc9\x9e\x19\x1a\x47\x49\xab\x81\xd6\x5d\x25\xdc\xbe\x36\xe8\x90\xa3\xa4\xe4\x76\x34\x7b\xaf\x9d\xd8\x71\xdd\x69\x6b\x39\x7c\xc5\x56\xd0\xd6\x6b\x6b\x19\x38\x23\x78\xf7\xce\xda\xd6\x8a\x89\xc6\x53\x37\x2a\x07\x4f\xae\x9d\x41\x4c\x4a\x7c\x9c\x19\xb1\x15\x12\x36\x30\xb6\x9c\xc9\x82\x8e\xa9\x1f\x39\x6c\x1b\x75\xce\x34\x5b\x09\x29\x9c\x80\x83\x1c\x64\x49\x12\x7e\x88\xc9\x74\xbc\xbc\xff\x32\x99\x8e\xee\x17\xe3\xf9\x8f\xc9\x70\x1c\x88\x13\x83\xfa\x50\x81\x49\xd9\x71\x70\x73\x44\x77\x25\x24\x54\xb3\x6d\x78\x8c\x52\x6c\x41\x81\xb5\x33\x83\x2b\x68\xdb\x4b\x9d\xd3\x5f\xc1\x85\x2e\x74\x99\x2f\x07\x03\x24\xa9\xd2\x81\x92\x8b\xb3\x8b\xb3\xe0\xb3\xe5\x29\x78\x90\xbf\x2d\x97\xb3\x96\x40\x28\xe1\x04\x93\x23\x90\x6c\xb7\x00\x8e\x2a\xb1\x34\x1c\xe0\x34\x18\x81\x49\x23\x1b\xb4\x65\x4e\x64\x80\xb9\xdb\x0b\x5b\x32\x9b\x73\x0e\xd6\x2e\x53\x03\x36\x45\x99\x84\xd2\x35\x13\x32\x37\xd0\x92\x9e\x07\x63\xb0\x78\x37\x14\xe1\xf0\xdc\x42\x62\x70\x31\xf8\x30\x12\xaf\x00\xf1\x97\x3f\x18\x87\x44\xd9\x9a\x81\x47\xe5\xb5\xab\x12\x94\x04\xf2\x0e\x02\xe3\xf5\xc5\x26\xc4\xad\xbb\xa1\x14\x50\x38\xc8\xec\x61\x4a\x17\x03\x41\xcd\xaa\x07\x7d\xac\x3c\x82\x4e\x61\xa5\xd8\xdc\x16\x3a\x35\x8f\xa5\x6f\xe4\xce\xb7\x84\x16\x1f\x11\xa9\x9f\x56\x3c\x2b\x30\x59\xd5\xe0\x8b\x77\xc2\xea\x92\xd9\x31\x98\xb7\x3a\xf6\x8b\x93\xf9\xd1\x1d\x7d\x7f\xb3\xf1\x13\x47\x99\x9f\x91\xe7\xc2\xa8\x43\x6c\xb9\x61\xfa\xc5\xbb\xfa\x1b\x06\xfd\x7a\x8e\xad\xe6\xd6\x96\xa5\xb7\x5e\x09\xc2\x49\xbd\xcb\x67\xe5\x63\x32\xa3\xed\x4b\xea\x74\xf1\x7c\xd2\x16\xda\x03\xe9\xfd\xf5\x64\xb1\x2c\x96\x34\xcd\x2b\x3e\x68\x4d\xba\xdd\x73\x0e\x3b\x54\xdc\xd1\x7f\x5e\x50\x28\x1b\x47\xdc\xd1\x62\x74\xd8\x89\x0e\x55\x84\xbe\x62\x99\x90\xbb\xba\x08\xc3\x00\x26\xb3\xab\xcb\x9b\xc9\xf5\xdd\xec\xfb\xf5\x64\x78\xf7\x7c\xd2\xfb\x5f\x00\x00\x00\xff\xff\x36\x35\x9b\xa2\xa7\x13\x00\x00")

func corednsYamlBytes() ([]byte, error) {
	return bindataRead(
		_corednsYaml,
		"coredns.yaml",
	)
}

func corednsYaml() (*asset, error) {
	bytes, err := corednsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "coredns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localStorageYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x5f\x6f\xdb\xb6\x16\x7f\xd7\xa7\x38\x57\xb7\x79\xb8\x17\xa5\x9d\x6c\xc3\x32\xb0\xd8\x83\x9b\x38\x69\x80\xc4\x36\x6c\x77\x43\x51\x14\x06\x2d\x1d\xdb\x6c\x28\x92\x20\x29\xb7\x6a\x96\xef\x3e\x90\x94\x1d\xc9\x71\x13\x07\xdb\xde\xa6\x17\x81\x87\xe7\xef\xef\xfc\x23\xd3\xfc\x37\x34\x96\x2b\x49\x61\x7d\x92\xdc\x72\x99\x53\x98\xa0\x59\xf3\x0c\x7b\x59\xa6\x4a\xe9\x92\x02\x1d\xcb\x99\x63\x34\x01\x90\xac\x40\x0a\x42\x65\x4c\x10\xcd\xdc\x8a\x68\xa3\xd6\xdc\xcb\xa3\x21\x36\xca\x11\x56\x0b\x46\x76\xab\x59\x86\x14\x6e\xcb\x39\x12\x5b\x59\x87\x45\x42\x08\x49\x9a\x96\xcd\x9c\x65\x1d\x56\xba\x95\x32\xfc\x1b\x73\x5c\xc9\xce\xed\x2f\xb6\xc3\x55\x77\xeb\xd3\x99\x28\xad\x43\x33\x56\x02\x0f\x77\xc8\x78\x6e\x53\x0a\xb4\x34\x21\xc0\x34\xbf\x34\xaa\xd4\x96\xc2\xc7\x34\xfd\x94\x00\x18\xb4\xaa\x34\x19\x06\x8a\x54\x39\xda\xf4\x35\xa4\xda\xbb\x65\x1d\x4a\xb7\x56\xa2\x2c\x30\x13\x8c\x17\xe1\x26\x53\x72\xc1\x97\x05\xd3\x36\x88\xaf\xd1\xcc\x83\xe8\x12\x9d\xbf\x16\xdc\x86\xff\x17\xe6\xb2\x55\xfa\xe9\x79\x93\x28\x73\xad\xb8\x74\x7b\xcd\x46\xa2\xca\x77\x6c\xfd\xff\x20\xc5\x6b\xf4\x5a\x5b\x82\x99\x41\xe6\x30\x28\xdd\xef\x9f\x75\xca\xb0\x25\xd6\xd0\x3f\x56\x5a\xdf\x67\x82\x59\x8b\x07\x22\xf0\x97\x12\xfd\x96\xcb\x9c\xcb\xe5\xe1\xf9\x9e\x73\x99\x27\x3e\xe9\x63\x5c\x78\xe6\x4d\x78\x4f\x18\x4e\x00\x1e\x17\xd8\x21\x65\x65\xcb\xf9\x67\xcc\x5c\xa8\xac\xbd\x6d\xf3\x4f\x35\x0b\xd3\xda\x3e\xc0\x75\x8e\x5a\xa8\xaa\xc0\x17\xf4\xe9\xf7\x4d\x59\x8d\x19\x0d\x69\x8f\xbc\xef\xb8\xcf\x79\x75\xcd\x0b\xee\x28\x1c\x27\x00\xd6\x19\xe6\x70\x59\x79\x2e\x00\x57\x69\xa4\x30\x56\x42\x70\xb9\x7c\xaf\x73\xe6\x30\xd0\x4d\x93\x12\x59\x01\x0a\xf6\xf5\xbd\x64\x6b\xc6\x05\x9b\x0b\xa4\x70\xe2\xd5\xa1\xc0\xcc\x29\x13\x79\x0a\x5f\x35\xd7\x6c\x8e\xc2\x6e\x84\x98\xd6\x4f\x84\xe1\xb0\xd0\x62\x6b\xa2\x19\xbf\xff\x44\x4b\xd3\x73\xba\x00\x36\xd1\xfb\x4f\x1b\xae\x0c\x77\xd5\x99\x2f\xf6\x41\x00\x33\x8d\x20\x11\x3f\x27\x48\x66\xb8\xe3\x19\x13\x69\xcd\x6f\x5b\xb9\x1f\xbc\x2c\xf1\x01\x4a\x25\xd0\x84\xc2\x6c\x78\x0c\x40\xe0\x16\x2b\x0a\xe9\x59\x6d\xaf\x97\xe7\x4a\xda\xa1\x14\x55\xda\xe0\x02\x50\xda\x4b\x2b\x43\x21\xed\x7f\xe5\xd6\xd9\x74\x8f\x92\xe0\xb9\x2f\xde\x8e\x4f\xba\x91\xe8\x30\xf4\x5e\xa6\xa4\x33\x4a\x10\x2d\x98\xc4\x17\xe8\x05\xc0\xc5\x02\x33\x47\x21\x1d\xa8\x49\xb6\xc2\xbc\x14\xf8\x12\xc3\x05\xf3\x2d\xf7\x77\x59\xf4\x61\x30\x2e\xd1\x6c\x11\x24\xcf\xf5\x41\xfc\x78\xc1\x96\x48\xe1\xe8\x6e\xf2\x61\x32\xed\xdf\xcc\xce\xfb\x17\xbd\xf7\xd7\xd3\xd9\xb8\x7f\x79\x35\x99\x8e\x3f\xdc\x1f\x19\x26\xb3\x15\x9a\xee\x7e\x45\x74\x7d\xdc\x39\xee\xfc\xf0\x53\x5b\xe1\xa8\x14\x62\xa4\x04\xcf\x2a\x0a\x57\x8b\x81\x72\x23\x83\x16\xb7\x09\xf7\xfe\x16\x05\x93\xf9\x43\xba\xc9\x73\x8e\x12\xb0\x8e\x19\xd7\x38\x13\x12\x77\x52\x83\xd4\x45\x97\x75\x23\xb5\xfe\x75\x3e\x5b\x25\xb7\x1c\x71\xbb\xdc\xf8\xda\xb3\x4d\xdb\x11\xaa\x28\x41\x22\x53\x03\xf9\xc2\xf3\x8f\x98\x5b\xd1\x96\x81\x2d\x07\xca\xf5\x63\x65\xa3\xe1\xf9\x6c\xd0\xbb\xe9\x4f\x46\xbd\xb3\x7e\x43\xd9\x9a\x89\x12\x2f\x8c\x2a\x68\x2b\xb7\x0b\x8e\x22\xaf\x47\xf7\x23\x7a\xb4\xbd\xe9\xf1\xce\x76\x82\x25\xcd\xa8\x5e\x10\x50\xa4\xdf\x30\xdd\xb6\xf6\xa8\x60\x6a\x7c\x77\xa7\x70\x7b\x59\x3e\xcc\xe3\x49\xa4\x87\xb9\xf1\xe4\x44\xf6\xeb\x49\x4a\xe5\x9a\x3d\xdf\xdc\xb0\x3b\xad\xc2\x2d\xc9\x71\xc1\x4a\xe1\x48\xb8\xa6\x90\x3a\x53\x62\x9a\x34\xeb\x10\xea\x3a\xf5\x02\x0d\x4b\x31\xf6\x7a\x9b\xde\xa8\x1c\x29\xfc\xce\xb8\xbb\x50\xe6\x82\x1b\xeb\xce\x94\xb4\x65\x81\x26\x31\xf1\xa9\xb3\x29\xda\x73\x14\xe8\x30\x44\x5e\xaf\xc8\x0d\x64\xc9\xce\xb3\xf1\xc9\xcd\xb3\x2d\xd0\xef\x2c\x9d\x8d\x60\xa3\x56\x29\xfc\x41\x02\x20\x77\x75\x6e\xc2\x04\xf1\x15\x70\xc3\x74\x4a\x3f\xd6\xd4\xbb\x6d\xe6\xc2\x7d\x4a\xd3\x4d\xe7\x8e\x7a\xd3\x77\xb3\x8b\xe1\x78\x36\x18\x0e\x66\xd7\x57\x93\x69\xff\x7c\x36\x18\x9e\xf7\x27\xe9\xeb\x07\x19\xef\x9d\x4d\xe9\xc7\xf4\xe8\x6e\x23\x77\x3d\x3c\xeb\x5d\xcf\x26\xd3\xe1\xb8\x77\xd9\x0f\x5a\xee\x8f\xc2\x43\xc7\x7f\xf7\xf5\x3f\x9e\xef\xc3\xfa\x72\xfe\x71\x51\x3b\xfb\xdf\xff\x74\xe7\x5c\x76\xed\x2a\x9c\xbe\xac\xb8\x40\x58\xa2\x53\xda\x59\x48\x0b\x6a\xa9\xa6\x29\x28\x1d\xdb\x37\x57\x0f\x73\x80\x59\x84\x57\x4a\x3b\xe0\xb2\x55\x8b\xfa\x7f\xad\x23\x9b\x5b\x25\x4a\x17\x70\xf8\xf5\xd5\x70\x34\xed\x8d\x2f\x5b\x0c\x6f\xde\xb4\x8e\xb6\x2d\x6e\xf9\x37\xbc\x92\x6f\x2b\x87\xf6\x10\xe9\xa2\x2d\xbd\x56\xc2\x57\xce\x73\x92\x68\x59\x56\xc7\x27\x63\xb7\x15\xb7\x39\x37\x40\x0a\x38\x3e\x3d\x3d\x05\xa2\xe1\xd5\x5d\x33\x90\x08\x6a\xb6\x2a\x54\x0e\xa7\xc7\xc7\xbb\xb7\xdd\x4e\x27\xec\x79\x66\x72\xf5\x45\xfe\x0b\xf5\x93\x50\x9b\x02\x88\x59\xec\x01\x78\x85\x42\xa3\x19\xa9\xbc\x53\xb1\x42\x6c\x51\xdc\xe9\x62\x4f\x8a\x8d\x3e\x52\xf9\xde\x17\x55\xec\xed\xa8\x8d\xe8\x9a\xa9\xf9\x6c\xfa\xfe\x0a\xde\x11\x82\x17\xad\xdd\x82\x1b\xa3\x0c\xe6\x44\xf0\xb9\x61\xa6\x22\xf3\xd2\x56\x73\xf5\x95\x9e\x74\x7e\xfc\xb9\x73\x72\xe0\xde\xfd\x33\x00\x00\xff\xff\xe1\x31\x46\x7e\xec\x0e\x00\x00")

func localStorageYamlBytes() ([]byte, error) {
	return bindataRead(
		_localStorageYaml,
		"local-storage.yaml",
	)
}

func localStorageYaml() (*asset, error) {
	bytes, err := localStorageYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "local-storage.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerAggregatedMetricsReaderYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xcf\x31\x6b\xf4\x30\x0c\xc6\xf1\xdd\x9f\x42\x78\x7e\x93\x97\x6e\xc5\x6b\x87\xee\x1d\xba\x94\x1b\x94\xf8\x21\x27\xce\xb1\x83\x24\xe7\x68\x3f\x7d\xb9\x70\xdc\x58\x68\x27\x0d\x7f\x7e\x0f\xe8\x22\x35\x27\x7a\x29\xdd\x1c\xfa\xd6\x0a\x02\x6f\xf2\x0e\x35\x69\x35\x91\x4e\x3c\x8f\xdc\xfd\xdc\x54\xbe\xd8\xa5\xd5\xf1\xf2\x6c\xa3\xb4\xff\xfb\x53\x58\xe1\x9c\xd9\x39\x05\xa2\xca\x2b\x12\xd9\xa7\x39\xd6\xc4\xcb\xa2\x58\xd8\x91\x87\x15\xae\x32\xdb\xa0\xe0\x0c\x0d\x44\x85\x27\x14\xbb\x11\xfa\x61\xfd\xb1\x30\x78\x1b\x76\xc1\x35\x51\x74\xed\x88\xbf\x71\xc8\xe2\x7f\x71\x9c\x57\xa9\x0f\xa8\xbd\xc0\x52\x18\x88\x37\x79\xd5\xd6\x37\x4b\xf4\x11\xef\x7f\xdd\x7d\x3c\x05\x22\x85\xb5\xae\x33\x8e\xbe\xb5\x6c\xf1\x1f\xc5\xda\x32\xec\xc8\x3b\x74\x3a\xd2\x02\xbf\x95\x22\x76\xdc\x2b\xfb\x7c\x8e\xa7\xf0\x1d\x00\x00\xff\xff\xe5\x1d\x7a\x17\x89\x01\x00\x00")

func metricsServerAggregatedMetricsReaderYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerAggregatedMetricsReaderYaml,
		"metrics-server/aggregated-metrics-reader.yaml",
	)
}

func metricsServerAggregatedMetricsReaderYaml() (*asset, error) {
	bytes, err := metricsServerAggregatedMetricsReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/aggregated-metrics-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerAuthDelegatorYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\x31\x4e\xc4\x40\x0c\x45\xfb\x39\xc5\x5c\x60\x82\xe8\x90\x3b\xa0\xa0\x5f\x24\x7a\x67\xf2\x59\x4c\x92\x71\x64\x7b\x22\xc1\xe9\xd1\x8a\x88\x06\xd8\xf6\x4b\xff\xbd\x57\x4a\x49\xbc\xc9\x0b\xcc\x45\x1b\x65\x1b\xb9\x0e\xdc\xe3\x4d\x4d\x3e\x39\x44\xdb\x30\xdf\xf9\x20\x7a\xb3\xdf\xa6\x59\xda\x44\xf9\x71\xe9\x1e\xb0\x93\x2e\x78\x90\x36\x49\x3b\xa7\x15\xc1\x13\x07\x53\xca\xb9\xf1\x0a\xca\x2b\xc2\xa4\x7a\x71\xd8\x0e\x23\xff\xf0\xc0\x4a\x17\x70\x99\xb0\xe0\xcc\xa1\x96\x4c\x17\x9c\xf0\x7a\x79\xf1\x26\x4f\xa6\x7d\xbb\x52\x90\x72\xfe\x15\xf0\xe3\xfb\x5b\xe0\x7d\x7c\x47\x0d\xa7\x54\x8e\xef\x33\x6c\x97\x8a\xfb\x5a\xb5\xb7\xf8\x27\xf7\x98\x7d\xe3\x0a\xca\x73\x1f\x51\xbe\xf9\xe9\x2b\x00\x00\xff\xff\xa5\xb5\x26\x22\x2f\x01\x00\x00")

func metricsServerAuthDelegatorYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerAuthDelegatorYaml,
		"metrics-server/auth-delegator.yaml",
	)
}

func metricsServerAuthDelegatorYaml() (*asset, error) {
	bytes, err := metricsServerAuthDelegatorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/auth-delegator.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerAuthReaderYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\x3d\x4e\x04\x31\x0c\x46\xfb\x9c\x22\x17\xf0\x22\x3a\x94\x0e\x1a\xfa\x45\xa2\xf7\x64\x3e\xc0\xcc\x8e\x13\xd9\xce\x08\x38\x3d\x1a\xb4\xfc\x34\x4b\x6f\xbf\xef\x3d\x22\x4a\xdc\xe5\x11\xe6\xd2\xb4\x64\x9b\xb8\x1e\x78\xc4\x4b\x33\xf9\xe0\x90\xa6\x87\xe5\xc6\x0f\xd2\xae\xb6\xeb\xb4\x88\xce\x25\x1f\xdb\x09\x77\xa2\xb3\xe8\x73\x5a\x11\x3c\x73\x70\x49\x39\x2b\xaf\x28\x79\x45\x98\x54\x27\x87\x6d\x30\xda\x51\x64\xe0\x19\x76\x3e\xf1\xce\x15\x25\x2f\x63\x02\xf9\xbb\x07\xd6\x64\xed\x84\x23\x9e\x76\x08\x77\xb9\xb7\x36\xfa\x3f\x26\x29\xe7\x5f\x91\x9f\x5d\xbc\x05\x74\x6f\x20\xee\xf2\x67\x1c\x1a\x52\xbf\xde\xbf\x35\x7c\x4c\xaf\xa8\xe1\x25\xd1\x19\xf4\x00\xdb\xa4\xe2\xb6\xd6\x36\x34\x2e\xa4\x5c\xd6\xff\x0c\x00\x00\xff\xff\x2a\x39\xe6\xe4\x44\x01\x00\x00")

func metricsServerAuthReaderYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerAuthReaderYaml,
		"metrics-server/auth-reader.yaml",
	)
}

func metricsServerAuthReaderYaml() (*asset, error) {
	bytes, err := metricsServerAuthReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/auth-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerMetricsApiserviceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8e\x4d\x6a\xc4\x30\x0c\x85\xf7\x3e\x85\x2e\xe0\x74\xb2\x2b\xde\x75\x59\x68\x61\x20\x65\xf6\x1a\x8f\x3a\x88\xe0\x1f\x24\xd9\x90\xdb\x97\x30\x71\xba\x13\x7a\xef\xfb\x24\xef\xbd\xc3\xca\x37\x12\xe5\x92\x03\x60\x65\xa1\x27\xab\x09\x1a\x97\x3c\xad\xef\x3a\x71\x79\xeb\xb3\x5b\x39\x3f\x02\x7c\x5c\x3f\x17\x92\xce\x91\x5c\x22\xc3\x07\x1a\x06\x07\x90\x31\x51\x80\x3e\xdf\xc9\x70\x9e\x12\x99\x70\xd4\x03\x76\x5a\x29\xee\x25\x7d\x81\xfb\x38\x88\xa3\xe9\xf7\x88\xe4\x0c\xb4\x62\xa4\x00\x6b\xbb\x93\xd7\x4d\x8d\x92\x03\x78\x4a\x69\xf5\x44\x86\x1c\xa0\x8f\xdf\x8f\xf3\x0e\x80\xb3\x52\x6c\x42\xcb\xca\xf5\xe7\x6b\xb9\x91\xf0\xef\x16\xc0\xa4\xd1\x10\x5d\x85\x8b\xb0\x6d\xdf\x9c\x39\xb5\x14\x60\xbe\x5c\xfe\x65\x23\x7d\xad\xff\x02\x00\x00\xff\xff\x14\x74\xa9\x1b\x25\x01\x00\x00")

func metricsServerMetricsApiserviceYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerMetricsApiserviceYaml,
		"metrics-server/metrics-apiservice.yaml",
	)
}

func metricsServerMetricsApiserviceYaml() (*asset, error) {
	bytes, err := metricsServerMetricsApiserviceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/metrics-apiservice.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerMetricsServerDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\xc1\x6e\xdb\x46\x10\xbd\xeb\x2b\x06\x2a\x7c\x2b\x6d\x49\xa9\xdd\x80\x80\x0f\x82\xc4\x44\x01\x6c\x47\x10\xe5\x16\x3e\x09\xeb\xe5\xc8\x5a\x78\xb9\xbb\x9d\x19\x2a\x61\x83\xfc\x7b\xb1\xa2\x43\x93\x8e\x1d\xb8\x68\xc3\x03\x0f\x33\x6f\xde\x3c\xbe\x9d\xe5\x24\x49\x32\x50\xc1\xfc\x81\xc4\xc6\xbb\x14\xf6\xe3\xc1\xbd\x71\x45\x0a\x39\xd2\xde\x68\x9c\x6a\xed\x2b\x27\x83\x12\x45\x15\x4a\x54\x3a\x00\x70\xaa\xc4\x14\x4a\x14\x32\x9a\x13\x46\xda\x23\x3d\x84\x39\x28\x8d\x29\xdc\x57\xb7\x98\x70\xcd\x82\xe5\xe0\x69\x07\x15\x02\x9f\xb4\x6d\xe6\x18\xac\xaf\x4b\xfc\x4f\x2d\x00\xac\xba\x45\xcb\xb1\x12\xe0\xfe\x2d\x27\x2a\x84\xef\xca\x39\xa0\x8e\x08\xc2\xbd\x89\x52\x16\x86\xc5\x53\x7d\x61\x4a\x23\x29\x8c\x06\x00\x2c\xa4\x04\xef\xea\x86\x47\xea\x80\x29\xac\xbc\xb5\xc6\xdd\x5d\x87\x42\x09\x1e\xe2\xd4\x8d\x34\x50\x80\x52\x7d\xbe\x76\x6a\xaf\x8c\x55\xb7\x16\x53\x18\x47\x3a\xb4\xa8\xc5\x53\x83\x29\x95\xe8\xdd\x45\x47\xe7\xcb\x4a\x01\x04\xcb\x60\x5b\xfa\xae\x33\xf1\x79\xc1\x9d\xf8\xd8\x5e\x83\x1f\xb5\x00\xf8\x66\x48\x7c\x02\x19\x4f\x46\xea\x99\x55\xcc\x57\x07\xfe\x61\xe3\x6e\xe2\x7c\x81\x89\x26\x23\x46\x2b\x3b\x7c\xc0\x73\x6f\x3c\xae\x5e\x16\x24\xde\x22\x29\x31\xde\x75\x54\x25\x70\x8f\x75\x0a\xc3\xd9\x03\xeb\xb4\x28\xbc\xe3\x8f\xce\xd6\xc3\x16\x03\xe0\x43\xac\xf4\x94\xc2\x30\xfb\x6c\x58\x78\xf8\x1d\xc1\x41\x1b\x79\x8b\xc7\x71\x1e\xc8\xa1\x20\x1f\x1b\x7f\xa2\xbd\x13\xf2\x36\x09\x56\x39\x7c\x25\x27\x00\x6e\xb7\xa8\x25\x85\xe1\x95\xcf\xf5\x0e\x8b\xca\xe2\xeb\x5b\x96\x8a\x05\xe9\xff\xe8\xb5\xf7\xb6\x2a\xb1\xb5\xeb\x17\x28\xa3\xc7\x60\x1c\x48\x19\x80\x3d\x7c\x42\xd0\xca\x01\xab\x2d\xda\x1a\x2a\x46\xd8\x92\x2f\x13\xd6\x14\x67\x0c\x4c\xa9\xee\x90\x41\xb9\xe2\xc4\x13\x10\xaa\x22\xf1\xce\xd6\x10\x4d\x51\xc6\x21\xf1\xe0\xdb\x27\x35\x93\x24\x65\x48\x0a\x43\xad\x3a\x2c\x83\xd4\x73\x43\x29\x7c\xf9\xfa\x10\x7c\xac\x4d\x9f\x14\x3f\x7b\xea\xd0\x88\x48\xe1\xe8\x4b\x7e\x93\xaf\xb3\xcb\xcd\x3c\x7b\x37\xbd\xbe\x58\x6f\x56\xd9\xfb\x0f\xf9\x7a\x75\xf3\xf5\x88\x94\xd3\x3b\xa4\x93\xd2\x10\x79\xc2\x22\xe9\x33\xa5\xfb\xd1\xf1\xd9\xf1\x9b\x96\x50\xd1\x5d\x6f\x82\x92\x44\x23\x49\xd4\x7d\x7e\x22\x65\xe8\x65\x18\x75\x45\x98\x04\x4f\x72\x3e\x1e\x4d\x4e\x47\xbd\x6c\x3c\x37\x8b\x92\x04\xc2\x2d\x52\xec\xac\x8a\x82\x90\x39\x89\x57\x9e\xcf\x8f\xbe\x2c\x57\xd9\xbb\x6c\xb5\xca\xe6\x9b\xe9\x7c\xbe\xca\xf2\x7c\xb3\xbe\x59\x66\xf9\xd7\xa3\x67\x79\x2a\xc6\xe6\x92\xb0\x28\xa9\xf8\xd0\xb6\x07\x6c\x3e\x2c\x21\x64\x6f\xab\x78\x15\xce\xc7\xa7\xdc\x43\x88\xe5\x44\x9b\xb0\x43\x4a\xb8\x32\x82\x7c\xbe\xbe\xc8\x37\xd9\x6c\xbe\xc8\xe2\x3b\x9f\x6e\xfe\xfc\xb0\x5e\x6c\xa6\x59\xbe\x99\x9c\x9e\x6d\xde\xcf\x2e\x37\xf9\x62\xfa\xe6\xed\x6f\xbf\x3e\xe2\x56\xaf\x42\x3d\x61\x1b\x4f\xde\x7e\xc3\x4d\x4e\xcf\x5e\x62\x7b\x11\xd5\x61\x9b\x2d\xa6\xb3\xc5\x74\x32\xda\x2c\x3f\x5e\xdc\x8c\xdf\x8c\x4e\x9f\x23\xfb\x0e\xd4\xba\x10\xcd\xa9\x48\x63\xe7\x8c\x63\xf0\xaf\x0a\x59\x7a\x31\x00\x1d\xaa\x14\xc6\xa3\x51\xd9\x8b\x96\x58\x7a\xaa\x53\xf8\x7d\x74\x69\xda\x44\x3c\x8a\xde\xd4\x34\x33\xbb\x13\x09\xdc\xa9\x6e\xa7\x7b\xe9\x49\x22\x77\x77\x64\xe2\xcf\xd1\x8b\xd7\xde\xa6\xb0\x9e\x2d\x3b\x8a\x55\x61\x1c\x32\x2f\xc9\xdf\x62\x57\x62\xa4\x7f\x8f\xd2\x57\x1d\x94\xec\x52\x38\x89\x55\xf5\xdf\xfd\xcc\xa1\xe9\x53\x4d\x00\xac\x77\x18\xd5\x2e\xd6\xeb\x65\xde\xc9\x18\x67\xc4\x28\x3b\x47\xab\xea\x1c\xb5\x77\x05\x37\xfb\xab\x25\x44\x32\xbe\x68\x53\x93\x4e\x4a\x4c\x89\xbe\x92\x36\x37\xee\xe4\xb8\xd2\x1a\x99\xd7\x3b\x42\xde\x79\x5b\xf4\xb3\x5b\x65\x6c\x45\xd8\xc9\x3e\xde\x4d\x6b\xf6\xf8\xaf\x9d\x88\x45\x3f\xc1\x88\xb3\x1f\x38\x31\x1e\xfd\x74\x2b\x0e\xbf\x9e\xb8\x48\xbd\x13\xfc\x2c\xfd\x69\x56\x45\xdc\x71\x2b\xef\xe5\x9d\xb1\xd8\xec\xd7\x14\x84\x2a\xec\xc2\x2a\x37\xe5\x2b\xef\x22\xec\xf9\xe4\x35\x23\x1d\x6e\x40\xf7\x73\x94\xb5\xfe\xd3\x92\xcc\xde\x58\xbc\xc3\x8c\xb5\xb2\x87\xb5\x9b\xc2\x56\x59\x7e\xe4\x68\xb6\xcb\x65\x5c\x29\xcf\xdc\x8c\xa7\xab\x00\x9a\xe5\xb3\x6c\x8e\x2c\xfe\x67\xff\x09\x00\x00\xff\xff\x70\xb0\x51\x48\x32\x0a\x00\x00")

func metricsServerMetricsServerDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerMetricsServerDeploymentYaml,
		"metrics-server/metrics-server-deployment.yaml",
	)
}

func metricsServerMetricsServerDeploymentYaml() (*asset, error) {
	bytes, err := metricsServerMetricsServerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/metrics-server-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerMetricsServerServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x3f\x4b\x04\x31\x10\xc5\xfb\x7c\x8a\x61\xfb\x28\xe2\x15\x92\x56\xb1\x13\x16\x4e\xec\xe7\x72\x4f\x0d\x9b\x6c\xc2\xcc\xec\xc2\x7e\x7b\xd9\xdc\x59\x1c\x5c\x97\xbc\x79\x7f\x7e\xde\x7b\xc7\x2d\x7d\x41\x34\xd5\x39\xd0\xfa\xe4\xa6\x34\x9f\x03\x1d\x21\x6b\x8a\x70\x05\xc6\x67\x36\x0e\x8e\x68\xe6\x82\x40\x05\x26\x29\xaa\x57\xc8\x0a\xb9\xca\xda\x38\x22\xd0\xb4\x9c\xe0\x75\x53\x43\x71\x44\x99\x4f\xc8\xba\x27\xa9\x5f\x64\x86\x41\x1f\x52\x7d\xbc\x34\x0d\x1f\x37\x55\xc3\x1d\x63\xcc\x8b\x1a\xa4\x3b\xd2\xbe\x30\x98\x2c\x18\x9c\x36\xc4\xbd\x58\x91\x11\xad\xca\x75\xe4\x45\x3d\xb7\x76\x87\xb1\x55\xb1\x4e\xe2\xfb\x33\xd0\xe1\xf0\xdc\x23\x17\x92\x5f\xb3\xa6\xfd\xdf\xa4\x5a\x8d\x35\x07\xfa\x7c\x1d\xbb\x62\x2c\x3f\xb0\xb1\xa7\xfe\x7d\xa9\xbd\x73\x49\x79\x1b\x6b\x4e\x71\x0b\x34\x0a\xbe\x21\x6f\x0b\xe7\xa3\x71\x9c\xdc\x5f\x00\x00\x00\xff\xff\x7b\xf5\x71\x2a\x57\x01\x00\x00")

func metricsServerMetricsServerServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerMetricsServerServiceYaml,
		"metrics-server/metrics-server-service.yaml",
	)
}

func metricsServerMetricsServerServiceYaml() (*asset, error) {
	bytes, err := metricsServerMetricsServerServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/metrics-server-service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerResourceReaderYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\x41\x4b\xc4\x30\x10\x85\xef\xf9\x15\xc3\xde\xd3\xc5\x9b\xe4\xa6\x1e\xbc\xaf\xe0\x3d\x4d\x9f\xbb\x63\xdb\xa4\xcc\x4c\x2a\xfa\xeb\xa5\xb6\x2a\xb8\xb0\x2c\x78\x4a\x98\xe4\x7d\x8f\xf9\xbc\xf7\x2e\x4e\xfc\x0c\x51\x2e\x39\x90\xb4\x31\x35\xb1\xda\xa9\x08\x7f\x44\xe3\x92\x9b\xfe\x56\x1b\x2e\xfb\xf9\xc6\xf5\x9c\xbb\x40\x0f\x43\x55\x83\x1c\xca\x00\x37\xc2\x62\x17\x2d\x06\x47\x94\xe3\x88\x40\xfa\xae\x86\x31\x8c\x30\xe1\xa4\x5e\x21\x33\xc4\x49\x1d\xa0\xc1\x79\x8a\x13\x3f\x4a\xa9\x93\x2e\x09\x4f\xbb\x9d\x23\x12\x68\xa9\x92\xb0\xcd\x72\xe9\xa0\xfb\x0d\xe0\x88\x66\x48\xbb\x3d\x1d\x61\xd7\x31\xa6\xd2\xe9\x2f\xec\x1c\xb2\x9c\x03\xeb\x7a\x79\x8b\x96\x4e\xee\x7f\x26\xee\x39\x77\x9c\x8f\xd7\x0b\x29\x03\x0e\x78\x59\xbe\x7d\xaf\x73\xa1\xd2\x11\x9d\xbb\xbf\x5c\xa0\xb5\x7d\x45\xb2\x2f\xe9\x6b\xf6\x09\x32\x73\xc2\x5d\x4a\xa5\x66\xfb\x89\xff\xc9\xad\x63\x9d\x62\x42\xa0\xbe\xb6\xf0\x2b\xdf\x7d\x06\x00\x00\xff\xff\xdb\x55\x9e\x61\x2a\x02\x00\x00")

func metricsServerResourceReaderYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerResourceReaderYaml,
		"metrics-server/resource-reader.yaml",
	)
}

func metricsServerResourceReaderYaml() (*asset, error) {
	bytes, err := metricsServerResourceReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/resource-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rolebindingsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\x31\x6f\xe3\x30\x0c\x85\x77\xfd\x0a\x21\xbb\x72\x38\xdc\x72\xf0\xd8\x0e\xdd\x03\xb4\xbb\x2c\xb1\x09\x6b\x59\x14\x48\x39\x41\xfb\xeb\x0b\xc7\x4e\xd2\xc4\x76\xe0\xb4\xe9\x66\x0b\xe2\xfb\x48\xbe\x07\xd9\x84\x2f\xc0\x82\x14\x0b\xcd\xa5\x75\x4b\xdb\xe4\x0d\x31\x7e\xd8\x8c\x14\x97\xd5\x7f\x59\x22\xfd\xd9\xfe\x55\x15\x46\x5f\xe8\xc7\xd0\x48\x06\x5e\x51\x80\x07\x8c\x1e\xe3\x5a\xd5\x90\xad\xb7\xd9\x16\x4a\xeb\x68\x6b\x28\x74\xd5\x94\x60\x6c\x42\x01\xde\x02\x9b\xf6\x37\x40\x36\xd6\xd7\x18\x15\x53\x80\x15\xbc\xb6\xb7\x6d\xc2\x27\xa6\x26\x5d\x21\x2b\xad\x07\xe0\x23\x47\xde\x25\x43\x5d\x1c\xf5\x13\xf6\x0c\x69\xca\x37\x70\x59\x0a\x65\x6e\x82\x3c\x0b\xf0\xc4\x14\x4a\x19\x63\xd4\xf7\xb7\x35\xb2\xa6\x43\xfb\xff\xc4\x38\x8a\x99\x29\x04\x60\xc5\x4d\x80\xb3\xc6\xa5\xad\x30\x7a\xb1\x50\x5a\x33\x08\x35\xec\xa0\x3f\x8b\xe4\x41\x94\xd6\x5b\xe0\xb2\x3f\x5a\x43\x9e\x59\x6b\x6b\x90\x64\xdd\xa5\x40\x40\xc9\xfb\x8f\x9d\xcd\x6e\x33\xa2\x15\x21\xef\x88\x2b\x8c\xeb\x7e\xde\x31\xf1\xee\x4e\xa2\x80\x0e\xf7\x04\xa3\x5d\xb7\x0c\x87\x9e\x6f\x45\x8e\x10\x20\xfa\x44\x18\x73\xa7\x9d\xc8\x4f\x69\xb6\x0b\x39\x69\xff\xd0\xc5\xe9\xcc\x4f\x98\x79\xff\xb0\x9f\x03\x4e\x49\x6f\x67\x9c\xc7\xb8\x48\xfb\x75\xc0\xfd\x63\xff\x35\x07\xa6\x4d\xf0\x64\xe4\x07\x49\x1b\xc6\x60\x76\xa8\x7e\xcd\xf8\x91\x71\xee\x67\xfa\x50\xfc\xdc\xf0\xae\x72\x8f\x18\x3a\x79\x78\x1d\xe6\xb5\xf1\x19\x00\x00\xff\xff\x20\xa2\xda\xb0\x09\x06\x00\x00")

func rolebindingsYamlBytes() ([]byte, error) {
	return bindataRead(
		_rolebindingsYaml,
		"rolebindings.yaml",
	)
}

func rolebindingsYaml() (*asset, error) {
	bytes, err := rolebindingsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rolebindings.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _traefikYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x91\x5f\x6b\xdb\x4a\x10\xc5\xdf\xf5\x29\x06\x81\x9f\x2e\x2b\xc5\x86\x0b\x41\x6f\xbe\x8e\x72\x1b\xda\xa6\xc1\x72\x5a\xf2\x64\xc6\xab\xb1\xb5\x78\xb5\xbb\xcc\x8e\x4c\xdd\x34\xdf\xbd\xac\xed\xfc\x83\x40\x4b\x69\xdf\x96\x9d\x99\xdf\x99\x39\x47\x29\x95\x61\x30\x9f\x89\xa3\xf1\xae\x82\x8e\x6c\x5f\x68\x14\xb1\x54\x18\x5f\xee\xc6\xd9\xd6\xb8\xb6\x82\x77\x64\xfb\x59\x87\x2c\x59\x4f\x82\x2d\x0a\x56\x19\x80\xc3\x9e\x2a\x10\x46\x5a\x9b\xad\xd2\xdc\x9e\xfe\x62\x40\x4d\x15\x6c\x87\x15\xa9\xb8\x8f\x42\x7d\x16\x03\xe9\x34\xa2\x13\xa4\x82\x4e\x24\xc4\xaa\x2c\x47\xf7\xef\x6f\xff\xab\xe7\xd7\xf5\xa2\x6e\x96\xd3\x9b\xab\x87\x51\x19\x05\xc5\xe8\xf2\xd0\x18\xcb\x17\x70\x35\x19\x17\x93\x62\xfc\xcf\x10\x0e\x8f\xb3\x42\x36\xdf\xb2\x3f\x78\xc0\xdf\x5b\xfe\xad\xc5\x01\x22\x49\x82\x02\x6c\xac\x5f\xa1\x2d\x8e\x62\x17\xb4\xc6\xc1\xca\x9c\x36\x26\x0a\xef\x2b\xc8\x47\xf7\xcd\x5d\xb3\xa8\x3f\x2e\x2f\xea\xcb\xe9\xed\x87\xc5\x72\x5e\xff\x7f\xd5\x2c\xe6\x77\xcb\xf9\xf4\xcb\xc3\x28\xcf\x00\x76\x68\x07\x8a\x33\xef\x84\x9c\x54\xf0\x5d\x1d\xb8\xc1\xb7\x53\xe7\x7c\x5a\xc9\xbb\x78\xd4\x02\x08\xec\x7b\x92\x8e\x86\x98\x0c\x0a\x3e\x5d\x94\x9f\x9f\x9d\x4f\xf2\x37\x1b\xa2\x66\x0c\x54\x41\x2e\x3c\xd0\xb1\x25\xb0\xdf\x99\x96\xf8\x09\x99\xbc\x62\x47\x42\xf1\xca\x6d\x98\xe2\x53\x01\x20\x0c\x2b\x6b\x62\x47\x6d\x43\xbc\x33\x9a\x9e\x2b\x00\xe4\x70\x65\xa9\x4d\x01\x0c\x74\x22\x1b\xcf\x46\xf6\x33\x8b\x31\x5e\x1f\xc2\xc9\x8f\xb6\x28\x6d\x87\x28\xc4\x4a\xb3\x11\xa3\xd1\x1e\x57\x31\x3d\x6e\x9e\x98\x4c\xc1\x47\x23\xfe\xe0\x1a\xa3\xd3\x1d\x71\xd9\x1b\x66\xcf\xd4\x2a\x6b\x56\x8c\xbc\x57\xa7\x50\x1e\xaf\x15\xdc\x54\x90\x4f\x8a\xf1\x59\xf1\xef\xf1\x4f\xbc\x25\x7e\xe9\x99\x82\x2d\x25\xe4\xec\x24\x3d\x6d\x5b\xef\xe2\x27\x67\xf7\x8f\x10\x1f\xd2\x84\xe7\x0a\xf2\xfa\xab\x89\x12\xf3\x57\x83\xce\xb7\xa4\xd8\x5b\x2a\x9e\x9d\x4a\xde\x6a\xef\x84\xbd\x55\xc1\xa2\xa3\x9f\xb0\x00\x68\xbd\x26\x9d\xc2\xba\xf6\x8d\xee\xa8\x1d\x2c\xfd\x9a\x4c\x8f\xc9\xb9\xdf\xe7\xc7\xd7\xd1\x99\x70\x89\xbd\xb1\xfb\x1b\x6f\x8d\x4e\xba\x37\x4c\x6b\xe2\x8b\x01\x6d\x23\xa8\xb7\x79\xf6\x23\x00\x00\xff\xff\xb9\x01\x23\x5a\x56\x04\x00\x00")

func traefikYamlBytes() ([]byte, error) {
	return bindataRead(
		_traefikYaml,
		"traefik.yaml",
	)
}

func traefikYaml() (*asset, error) {
	bytes, err := traefikYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "traefik.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"ccm.yaml":           ccmYaml,
	"coredns.yaml":       corednsYaml,
	"local-storage.yaml": localStorageYaml,
	"metrics-server/aggregated-metrics-reader.yaml": metricsServerAggregatedMetricsReaderYaml,
	"metrics-server/auth-delegator.yaml":            metricsServerAuthDelegatorYaml,
	"metrics-server/auth-reader.yaml":               metricsServerAuthReaderYaml,
	"metrics-server/metrics-apiservice.yaml":        metricsServerMetricsApiserviceYaml,
	"metrics-server/metrics-server-deployment.yaml": metricsServerMetricsServerDeploymentYaml,
	"metrics-server/metrics-server-service.yaml":    metricsServerMetricsServerServiceYaml,
	"metrics-server/resource-reader.yaml":           metricsServerResourceReaderYaml,
	"rolebindings.yaml":                             rolebindingsYaml,
	"traefik.yaml":                                  traefikYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"ccm.yaml":           &bintree{ccmYaml, map[string]*bintree{}},
	"coredns.yaml":       &bintree{corednsYaml, map[string]*bintree{}},
	"local-storage.yaml": &bintree{localStorageYaml, map[string]*bintree{}},
	"metrics-server": &bintree{nil, map[string]*bintree{
		"aggregated-metrics-reader.yaml": &bintree{metricsServerAggregatedMetricsReaderYaml, map[string]*bintree{}},
		"auth-delegator.yaml":            &bintree{metricsServerAuthDelegatorYaml, map[string]*bintree{}},
		"auth-reader.yaml":               &bintree{metricsServerAuthReaderYaml, map[string]*bintree{}},
		"metrics-apiservice.yaml":        &bintree{metricsServerMetricsApiserviceYaml, map[string]*bintree{}},
		"metrics-server-deployment.yaml": &bintree{metricsServerMetricsServerDeploymentYaml, map[string]*bintree{}},
		"metrics-server-service.yaml":    &bintree{metricsServerMetricsServerServiceYaml, map[string]*bintree{}},
		"resource-reader.yaml":           &bintree{metricsServerResourceReaderYaml, map[string]*bintree{}},
	}},
	"rolebindings.yaml": &bintree{rolebindingsYaml, map[string]*bintree{}},
	"traefik.yaml":      &bintree{traefikYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
