// Code generated by go-bindata.
// sources:
// ../../resources/.testing_config.yaml.swp
// ../../resources/default_config.yaml
// ../../resources/testing_config.yaml
// DO NOT EDIT!

package resources

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _resourcesTesting_configYamlSwp = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\xda\xbd\x8e\xeb\xc4\x17\x00\xf0\x73\x6f\xfb\x5f\x6b\xff\x02\x1e\xc0\x0c\xc5\x36\xf9\x98\x19\x8f\xed\xf1\x48\x34\xd0\xdc\x2b\x74\xa1\xe0\xe3\x16\x14\xe8\xcc\xcc\x71\xd6\xca\xc6\x0e\xf6\x78\x77\xa3\xcb\x02\xcf\x01\x0f\x42\x41\x47\x43\x8d\x90\x68\xe9\x28\x78\x00\x6a\xe4\x24\xb0\x2b\x9a\xbd\x20\xc4\x15\x62\x7e\x89\x64\xcd\x89\xcf\x99\x73\x26\x4a\x11\x25\x96\x7f\xf0\xf8\x49\xaa\x17\x1c\x00\xe0\xff\x00\xf3\xef\xbe\xfe\xf0\xd1\xb7\x29\x7c\xfe\x10\xe0\x12\x57\x3d\xb6\x01\xee\x73\xd5\xf5\xeb\xb6\xb9\xbe\xf7\xbe\x4f\x8f\x05\x97\x53\xc2\xd2\x51\x1b\xe6\xab\x6e\x39\xf4\x6e\xb9\x6a\xc2\xf9\x68\x17\xae\xdb\xec\xc3\x7d\x53\x8f\x2b\x5a\xae\xba\xf9\x9d\x55\x4f\x43\x37\xf6\x8e\x86\x65\xa0\x21\x34\xed\xea\x23\xd7\xb5\x75\xb3\x5a\xec\x70\x73\x71\xef\xde\x51\xf4\x9f\x36\x86\x7a\xae\x13\xc8\xa4\xd8\x7f\xd4\x5f\x63\xaf\xa6\x2f\xbf\xf4\xfe\x8b\xee\x2a\x8a\xa2\x28\x8a\xa2\x28\x8a\xa2\x7f\x50\xd8\x3e\x80\xcf\x00\xe0\xe1\x71\xfd\xca\xf1\xfa\xe0\x0f\xd7\x28\x8a\xa2\x28\x8a\xa2\x28\x8a\xa2\x28\x8a\xfe\xbd\xd0\x03\xfc\xf0\x3f\x80\x37\x4e\x0e\xbf\xff\xff\xf6\xfd\xff\x97\x53\x80\x9f\x4f\x01\x7e\x3c\x05\xf8\x29\x01\xf8\x3e\x01\xf8\x32\x01\xf8\x22\x01\xb8\x49\x00\x3e\x49\x00\xc6\x04\x80\x12\x80\x77\x12\x80\xb7\x13\x80\x47\x09\xc0\xeb\x09\x00\x24\x00\xdf\x9c\x00\x7c\x75\x02\x70\x73\x72\xa8\xfd\xe6\xc9\x0b\x1e\x34\x8a\xa2\x28\x8a\xa2\x28\x8a\xa2\xbf\x41\x9a\xa6\xe9\xb6\x6f\x2e\x31\xd0\x5b\xb4\x33\xe9\x62\xb1\x3c\x3c\xe9\x1a\x37\xdb\x8b\xbb\x7f\xdc\xa5\x70\x8e\x63\x38\x5f\xac\x69\xb7\xd8\xd2\xe6\x90\x3a\xda\x8b\xc6\x3d\x6f\xe6\x76\xb4\xc7\xcc\x63\xc4\xfc\x99\xfd\x87\x66\xd5\x62\x18\x7b\x12\x7f\xa1\x85\x3b\xc9\xb7\x5d\x4c\xc1\xa6\x5d\x19\x58\xd3\x6e\x30\xd3\x61\x5c\x75\xfd\x9a\xfa\xa7\xd8\x84\xf7\x9a\x0d\x3d\x79\xd7\xa4\x02\xd2\xb4\x1d\x37\x4f\xf7\x2f\x0c\x26\x95\xf0\xf1\x48\x23\x19\x80\xc6\x53\x1b\x9a\xb0\x7b\xec\x4d\x7a\xc6\xaf\xb9\xb8\x7d\x9c\x1d\x0e\x36\x4d\xd1\xfb\x9e\x86\xc1\xa4\xba\xb2\x1c\x75\x91\xeb\xcc\x29\xa5\x14\xba\xda\x97\xc2\xaa\x22\x23\xee\x33\x97\xe7\x48\x42\x09\x89\xf9\x31\x6f\x8b\xc3\x70\xd5\xf5\x53\xe5\xb3\x63\x68\x3d\x4d\x78\xf6\x8c\x1d\x4b\x32\xc3\x9e\xb7\x26\x9b\x31\xd7\xef\xb6\xa1\x63\xe6\x19\x73\xcd\xf6\x9c\x7a\x66\x18\xd2\x30\x17\x52\xcf\x5d\xe8\xa7\x1b\xf6\xe1\x40\xd7\x81\x19\xe6\xca\xb2\xaa\x75\x56\x56\xbe\x2c\xb9\xaf\xa4\xab\x9d\xf0\xde\x2b\xd4\x75\x26\x7c\x8e\x1c\xbd\xd3\xb5\x44\x6e\x25\x0a\xc5\x45\x56\x72\x9f\x15\x19\xaf\x33\xed\xb8\xd3\xf8\x7b\xbd\x2d\xf6\xb8\x19\xa6\x6d\x9b\x4b\x66\x58\x56\x38\x51\x68\x2a\x33\x5b\x57\x9a\xd7\x54\xe6\x96\x97\xb2\xac\x75\xc5\xb1\x14\xe8\xd9\xcd\x8c\xad\x7d\xcd\x0c\x1b\xf6\x0d\xb3\xfd\xf2\xb6\x88\x5f\x5f\x50\xcb\x4c\x26\x67\xac\x65\x46\x16\x52\x28\x35\x63\x5b\x66\xc4\x8c\xf5\xcc\xe8\x19\x1b\xf0\x62\x1a\xc0\x93\xb0\x24\x0a\xca\x5c\xa5\x45\xa5\x94\x17\xe4\x50\x5a\x6d\x65\x49\x8a\x0a\xe2\x36\xb7\xb5\x55\x99\x25\x9e\x95\x05\xe6\x5e\x6b\x5d\xd5\x58\x94\x15\x4a\x2d\xa4\x9c\x1a\xd9\xa0\x9b\x8e\xc2\x09\xa9\xad\x16\x79\x9e\xe7\x16\x05\xa1\x2f\x1d\x52\xc5\x0b\x4e\x5a\x2b\x89\xb5\x43\x9d\xe5\x85\xe7\x85\xca\x73\xeb\x2b\xcc\xcb\x5c\x5a\x2c\x6a\xe7\x78\x25\xa9\x9e\x2a\x35\x9e\x19\xa6\x72\xe2\x05\xc7\x62\xee\x25\xd2\x5c\x65\x56\xcf\x2b\x29\xeb\xb9\x52\x5a\x56\xaa\xaa\x7c\x56\x7a\x36\x63\x97\xd4\x0f\x4d\x37\x0d\x79\x73\x78\xe3\x37\xd8\xb4\x26\x85\x34\x45\xe7\xba\xb1\x0d\x83\x01\x0a\xe7\xd4\xd3\xb8\x31\xf0\x6b\x00\x00\x00\xff\xff\xbe\x77\x2b\xbb\x00\x30\x00\x00")

func resourcesTesting_configYamlSwpBytes() ([]byte, error) {
	return bindataRead(
		_resourcesTesting_configYamlSwp,
		"resources/.testing_config.yaml.swp",
	)
}

func resourcesTesting_configYamlSwp() (*asset, error) {
	bytes, err := resourcesTesting_configYamlSwpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/.testing_config.yaml.swp", size: 12288, mode: os.FileMode(420), modTime: time.Unix(1539035446, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesDefault_configYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x95\xcb\x76\xe2\x38\x10\x86\xf7\x7e\x0a\x1d\x66\x3d\xa0\x8b\xad\x8b\x77\x24\x4d\xfa\x12\x92\x03\x84\x0e\x1d\x76\xba\x94\x41\x0d\x58\xc4\x96\xb9\xf4\xd3\xcf\x71\x08\x93\x99\x9c\x09\x33\x0f\x30\xac\xb0\x4a\xf5\x57\xb9\xbe\xdf\x92\x85\x32\x56\xbe\x68\x16\x70\x0f\x71\x1f\xaa\x55\x8e\xde\x96\x26\x4d\x5d\x7b\x5d\x2e\xfd\x7a\x3d\x88\xcb\x89\x2f\x57\x60\x8e\x49\x79\xda\x58\xe7\x09\x42\xbf\xa1\x61\xb0\x7a\x8d\x22\xd4\xd1\x97\x0b\x64\x43\x19\x2b\x6d\x23\xd2\xce\x55\x50\xd7\x50\xa3\x12\xc0\xa1\x18\x90\x01\x54\x43\x44\x7b\x1f\x97\x08\xca\x1d\xda\xe9\xca\x6b\xb3\x86\xba\x9b\xa0\x73\x7e\x2b\x89\x90\x77\x39\x62\x8c\xbd\xfc\x87\xb8\x84\x0a\x9a\xcd\x6b\x77\x5f\x5d\x8e\x24\x93\xa7\x98\x09\x21\xd6\xb1\xd2\xdb\x11\x40\x55\x9f\x72\x11\xfa\x1d\x75\x7a\x7e\x9b\xf6\x08\x15\x5d\xdc\xc5\x5d\xd2\x8b\x76\xdb\x63\x92\x62\xda\xf3\xdb\xa2\xee\x8d\x37\xd3\xf1\xc1\xec\x57\xcd\xfc\xe9\xe9\x53\xd1\xfc\x9a\x9a\xc3\xa0\x3f\x81\xe9\xfd\xf5\x30\xfc\x3a\x1e\xb3\x4c\xee\xc6\xe5\xe2\x71\x37\xba\xfb\x39\x7c\x5a\x75\xfe\x55\x96\x9d\x65\x1f\x0b\x3e\xb8\xe7\x9b\xd5\xf3\x0c\x7e\xce\x6e\x67\xf4\x79\xd4\x10\xfe\x63\xeb\x3e\xb3\xd5\xb7\x40\xa6\x6c\xb3\xd4\xcb\xd1\x55\xf6\x00\x59\x49\x4e\xb2\xe7\x71\xf5\xcf\xd3\x3a\xbf\x84\x77\x50\x46\x1f\x8f\x37\xda\xc6\x50\x1d\x73\xd4\xe9\xbc\x8b\x4c\x60\xe1\xeb\xf8\xb7\x90\x2e\xed\x32\x54\x17\x02\xdb\x50\xfb\x3f\xe5\x12\xf4\x21\xe9\xaf\x65\x84\x45\xa5\xa3\x0f\xe5\x1b\x11\x82\xc9\xff\x44\xfe\x4a\x04\x1f\x04\xa3\x8c\x98\x54\x31\xcc\xb5\x35\x5c\x1b\x2c\x52\x8d\x05\x13\xb2\x60\x82\x17\x56\x19\x67\xa8\x20\xfc\x02\x3b\x7c\x70\x5c\x5a\x8e\x85\x31\x96\x62\x89\x31\xe6\x52\x31\x92\x62\x45\xb8\xe6\x92\x92\xc2\x09\xc2\x0c\x55\xf4\x43\xca\xf8\x60\x32\x41\x0a\x48\xb1\x03\xea\x38\x01\x95\xa5\xae\x50\x56\xa7\xa0\x38\x29\xb8\x4a\xa5\x2c\x08\xcb\x34\xbe\x08\xfd\xf5\xf3\x7e\x03\x9e\x91\xff\xc0\x94\x65\x5d\x4a\xb3\x2e\xc5\xb8\x9b\xd2\xf7\x5c\x09\xfd\xc4\x6e\x43\x98\x0d\xbd\xb7\xe3\xc7\xfd\x74\x39\xbd\x7a\xe2\x87\x5b\x3b\x0a\xc3\x82\x4f\xc6\x4f\xdf\x6e\xb6\xfb\x82\x54\x22\xdb\x0f\x0f\x74\x3e\x61\xdb\x6b\x47\xde\xd3\x7d\x2d\x20\x79\x97\x12\xfc\x51\x81\xf1\xfc\xae\x2f\x3f\x8f\xbe\x54\xbb\xc1\xfc\x4a\xed\xdd\x2a\x7c\xb7\xfd\xfe\xe6\x7a\xfe\x65\xab\xe0\x78\x9c\xa7\x0f\x03\xb9\xb8\xa9\xd8\x72\x7a\xff\xe3\x65\x08\xff\x68\xe3\xf4\x03\x07\xa0\x0b\x16\x50\xd8\x51\x95\x66\x82\x80\x60\x32\xa5\x5c\x09\xcd\xb9\x11\x5a\x29\x8d\x95\x73\xdc\x0a\xe6\x58\xc6\xdd\x45\x0b\x28\xce\xb1\xc5\x4c\x39\x46\x48\x9a\x31\x5d\x60\x97\x49\x9b\x71\xce\x05\x65\x4e\x59\x5a\x68\xe1\x38\xd8\x0b\x16\x60\x29\xd5\x22\x35\x4c\x52\x47\x94\xd3\x3c\x55\x52\x1a\x26\xb8\xc3\x90\x6a\x9e\x71\x23\x4c\xa1\x33\xe7\x2e\x1c\x09\xf8\x20\x0a\xd9\x5a\x47\x2b\x89\x09\x75\xa2\xd0\x59\x66\x25\x66\xc6\x68\x4a\x39\x36\xd6\x01\xa4\x26\x03\xd7\x49\x92\xe7\x06\x1a\x68\xcd\x50\x36\x9b\x59\xa8\x56\xad\x35\x10\x4d\x10\xda\xbf\x3c\xcc\xb4\x8f\x53\xbf\x81\xbb\x87\x1c\x91\x24\x39\x0f\xbb\x4d\x70\x50\xe8\x66\x1d\xfb\xd6\x86\xa6\x8c\xf7\x7a\x03\x39\xea\x6c\xb4\x2f\xdb\xd6\xca\xe0\xe0\xfb\x64\x98\xa3\x7d\x9d\xf7\x7a\xeb\xf6\x6e\x59\x86\x3a\xe6\x2a\x4b\x79\x82\xd0\x42\xd7\xa3\xca\x5b\x68\x4f\xa3\xf3\xef\xb4\x3c\xf4\x1b\x1f\x73\x94\x0a\x42\x99\x94\xc9\x09\x23\x1c\xe2\xb9\x91\xd0\xc4\x1c\x75\x38\xc6\x75\x5b\x66\xa3\x0f\x13\x88\x95\x6f\xe1\xd2\x17\x09\x5f\x46\xa8\x76\x7a\xdd\x2e\xb7\xc3\xa0\x2f\xfb\xe2\x61\x14\xc2\xba\x6f\x2d\xd4\xf5\xa0\x6c\x6f\x2b\x97\xa3\x58\x35\x90\x24\x7f\x04\x00\x00\xff\xff\xa1\xfa\x63\x77\x39\x07\x00\x00")

func resourcesDefault_configYamlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesDefault_configYaml,
		"resources/default_config.yaml",
	)
}

func resourcesDefault_configYaml() (*asset, error) {
	bytes, err := resourcesDefault_configYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/default_config.yaml", size: 1849, mode: os.FileMode(420), modTime: time.Unix(1538517588, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesTesting_configYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x52\xcb\x6e\xe3\x3a\x0c\xdd\xfb\x2b\x0c\x6e\xb2\x71\x52\xbd\x5f\x7f\x70\x71\x71\x57\x77\x80\xae\x29\x89\x6a\x8c\xc4\x8e\x47\xb6\xdb\x06\x45\xff\x7d\xe0\x34\x9d\x6e\xa7\x03\x69\x41\x12\x3a\x87\x87\xe2\xa1\xe5\x48\x95\xd6\x21\x34\x6d\x8b\x29\x5d\xd6\x71\x99\xb7\xb8\x6d\x07\xec\xc7\xd0\xde\xc2\xb6\x3d\xd1\x35\xb4\xbb\x37\xc0\x9c\x2b\xcd\x33\x04\x70\x3e\x32\x74\x46\x3b\x99\x94\x52\x0a\x53\xc9\x96\x47\x65\x24\xb1\x2c\x93\xd6\x48\x5c\x71\x81\x1a\x3a\x48\xf5\x3a\x2d\x17\x08\x6f\x90\xfa\xe9\x48\x15\x02\x20\xcd\x7b\x2e\xdc\x3e\x2d\x75\x7b\x70\x2b\x2f\xf4\xba\x40\x80\x64\xad\x2f\x4e\x5a\x9f\xad\x65\xd9\x8b\x54\x12\xcf\x39\x2b\x74\x45\xf2\xac\x91\x61\x4e\xae\x08\x64\x51\x20\x57\x8c\x4b\xcb\xb2\x34\x92\x15\xe9\x12\x4b\x0e\x7f\xf3\x4d\x58\x71\x98\xb7\xb6\xfd\x33\x04\x90\x26\x71\xe3\xc8\xca\x58\xbc\x63\x85\xac\x8e\xcc\x0a\x5b\x9c\x67\x68\x39\x66\x78\xef\xe0\x94\x0b\x04\x98\x6f\x82\xe1\x96\x7e\x91\xe4\xd3\x99\x46\x08\x52\x74\x30\x42\x10\x46\x70\xa5\x3a\x98\x20\xf0\x0e\x2a\x04\xd7\xc1\x8c\xe7\x6d\x80\x4c\x3c\x12\x37\x24\x93\x77\xdc\x2b\x95\x39\x25\x14\xd1\x45\x61\x49\x91\x21\x16\x75\x2c\x51\xc9\x48\x4c\x5a\x83\x3a\x3b\xe7\x7c\x41\x63\x3d\x0a\xc7\x85\xd8\x84\x0c\x98\xb6\xaf\x48\x5c\xb8\xe8\xb8\xd6\x5a\x47\xe4\x84\xd9\x26\x24\xcf\x0c\x23\xe7\x94\xc0\x92\xd0\x49\x6d\x32\x33\x4a\xeb\x98\x3d\x6a\xab\x45\x44\x53\x52\x62\x5e\x50\xd9\x98\xfa\x0c\x01\x94\x26\x66\x18\x9a\x7d\x16\x48\x7b\x25\xa3\xdb\x7b\x21\xca\x5e\x29\x27\xbc\xf2\x3e\x4b\x9b\xa1\x83\x67\xaa\x73\x7f\xd9\x86\x7c\xdf\xdd\x17\x3f\xe1\x3c\xbf\x5c\x6a\x0e\xed\xee\xb3\x74\xf7\x40\x68\xff\xd4\x02\x4d\xd3\x67\x1a\x97\x7e\xb9\xfe\xb3\xf1\xb0\x57\xc6\xbf\xce\xae\x69\x7e\xae\xb4\xd2\x66\xba\x71\x1d\x1e\x2f\xf5\x44\x75\x0e\xad\x68\xda\xf6\xe5\x96\x3c\x62\xbf\xfc\xe8\x07\xfa\xef\xff\xd0\xf2\xa6\x39\xd1\xf5\xe6\xd0\xb9\x7f\x1a\xfb\xf1\xe9\xc3\xac\xd3\x1a\xcf\x7d\xfa\x77\x73\xe9\xe1\xf0\xf0\x71\xe9\x15\x87\xe9\x4c\x0f\x95\xe6\xcb\x5a\x13\xcd\x0f\x1b\x04\x97\xb5\x12\x3f\x4c\x6b\x3c\x4c\x34\x7c\x80\x6b\xff\x8c\x0b\x7d\x03\x7d\xa2\xeb\x1d\x4d\xcb\x11\xd7\xe5\xf8\x1d\x15\x77\xc8\xdf\x48\xf8\x84\x7e\xf6\xff\x15\x00\x00\xff\xff\x2a\x32\x0b\x61\xbe\x03\x00\x00")

func resourcesTesting_configYamlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesTesting_configYaml,
		"resources/testing_config.yaml",
	)
}

func resourcesTesting_configYaml() (*asset, error) {
	bytes, err := resourcesTesting_configYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/testing_config.yaml", size: 958, mode: os.FileMode(420), modTime: time.Unix(1539035437, 0)}
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
	"resources/.testing_config.yaml.swp": resourcesTesting_configYamlSwp,
	"resources/default_config.yaml": resourcesDefault_configYaml,
	"resources/testing_config.yaml": resourcesTesting_configYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
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
	"resources": &bintree{nil, map[string]*bintree{
		".testing_config.yaml.swp": &bintree{resourcesTesting_configYamlSwp, map[string]*bintree{}},
		"default_config.yaml": &bintree{resourcesDefault_configYaml, map[string]*bintree{}},
		"testing_config.yaml": &bintree{resourcesTesting_configYaml, map[string]*bintree{}},
	}},
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

