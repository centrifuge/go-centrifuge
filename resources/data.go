// Code generated by go-bindata.
// sources:
// ../build/configs/default_config.yaml
// ../build/configs/testing_config.yaml
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

var _goCentrifugeBuildConfigsDefault_configYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\x49\x73\xdb\xb8\xb6\xc7\xf7\xfa\x14\xe7\x39\x9b\xee\x45\x64\x02\x9c\xb5\xa3\x26\x27\x71\xec\xc8\x96\x87\x8e\x37\xaf\x40\xf0\x50\x42\x4c\x12\x0c\x00\x4d\xfe\xf4\xaf\x40\x52\x8a\x32\x28\xfd\x2a\x7d\xbb\xab\x6e\xdd\xeb\x8d\x55\x20\xf0\xc7\x19\x7e\xe7\x10\xc4\x2b\x18\x63\xce\x56\x85\x81\x0c\xd7\x58\xc8\xba\xc4\xca\x80\x41\x6d\x2a\x34\xc0\x16\x4c\x54\xda\x80\x12\xd5\x33\xa6\xbb\x1e\xc7\xca\x28\x91\xaf\x16\x78\x8d\x66\x23\xd5\xf3\x00\xd4\x4a\x6b\xc1\xaa\xa5\x28\x8a\x5e\x23\x26\x2a\x04\xb3\x44\xc8\x3a\xdd\xaa\x9d\xa9\xc1\x2c\x99\x81\xd1\x41\x01\x4a\x26\x2a\x63\xf5\x7b\xfb\x29\x83\x1e\xc0\x2b\x78\x2f\x39\x2b\x1a\x13\x44\xb5\x00\x2e\x2b\xa3\x18\x37\xc0\xb2\x4c\xa1\xd6\xa8\xa1\x42\xcc\xc0\x48\x48\x11\x34\x1a\xd8\x08\xb3\x04\xac\xd6\xb0\x66\x4a\xb0\xb4\x40\xdd\xef\xc1\x7e\xbd\x95\x04\x10\xd9\x00\x5c\xd7\x6d\x7e\xa3\x59\xa2\xc2\x55\xd9\x79\xf0\x36\x1b\x40\xe4\x46\xed\xb3\x54\x4a\xa3\x8d\x62\xf5\x0c\x51\xe9\x76\xed\x6b\x38\x3b\x17\xb5\x77\x4e\x68\xd8\x77\xfa\x4e\x9f\x9c\x1b\x5e\x9f\xbb\x11\x75\xe8\xb9\xa8\x73\x7d\x7e\x53\xde\xdd\x6c\xd3\xcd\xf3\xea\xe9\xe3\xc7\x71\xbe\x7a\xb9\x4b\xb7\x93\xe4\x16\xef\xae\x47\xef\xe5\xcb\x6e\xe7\xfb\xd1\xfa\xa6\x5a\x3c\xac\x67\x57\x9f\xde\x7f\x7c\x3e\xfb\x13\x51\x77\x2f\xfa\x90\x07\x93\xeb\xa0\x7c\xfe\xfc\x88\x9f\x1e\x2f\x1f\xe9\xe7\xd9\x8a\x04\x7f\xd4\xd9\x85\xfb\xfc\x4e\x92\x3b\xb7\x5c\xb2\xe5\x6c\xe8\xcf\xd1\xaf\x48\x2b\xba\x0f\x55\xb2\x8f\x54\xeb\x80\x75\x1f\x2b\x23\xcc\x6e\xca\xb8\x91\x6a\x37\x80\xb3\xb3\xee\x09\xab\xf8\x52\xaa\x5b\xac\xa5\x16\xdf\x3c\x12\xd5\x5a\x0a\x8e\xf7\x55\xcd\x6c\xf8\xce\xce\x7a\x4d\x76\xae\x98\xa8\x7e\xc8\x4a\x97\x44\xf8\xed\xb6\x85\xe5\xf7\x1e\x1c\xc3\xd1\xda\xf2\x0a\xae\x57\x25\x2a\xc1\xe1\xed\x18\x64\xde\x80\x72\x84\x44\xa7\x71\xc8\x99\x4f\xba\x55\xc3\x7d\x62\xa0\x10\xda\xd8\x95\x95\xcc\xf0\x7b\xa6\x6a\x25\xd7\xa2\x79\x20\x1b\xed\x23\x03\xf6\x86\xfe\x69\xa2\x5d\xbf\x4f\xa9\xdf\xa7\x8e\xd3\xf7\xe8\xb7\xc9\x26\x74\xec\x5e\x4a\xf9\x78\xad\x9f\xf4\x63\x78\x97\xf2\x27\x3f\xba\x0e\xc9\xfd\xcd\xfc\xd2\x1f\x7f\x7a\xfa\x5c\x4e\x9f\xdf\xcc\xde\x6c\xb6\xd3\xcb\xbb\x64\x27\xef\xef\xc7\x51\x96\x9f\xfd\x48\x3e\x0a\xfa\x94\x38\xa7\xe4\xc7\x48\xf5\xe6\x71\xe2\xe6\x54\xbc\x4b\xef\xf1\x26\xbe\xb8\xbf\xbf\x19\xbe\x19\xa9\xc7\xf7\xe9\x90\xb3\xf8\xea\xe2\xea\x73\x5e\xa6\xa3\x85\x5a\xa5\x67\x5d\x8c\x26\x1d\xd8\x87\x4c\xbc\x1d\xc3\x6b\xe8\xb2\x71\x0a\x7d\xaf\x5b\xfc\x9e\xd9\xf0\x40\x86\x75\x21\x77\x98\xc1\xbc\x64\xca\xc0\xa8\x23\x4a\x43\x2e\x55\x13\xd0\x85\x58\x63\xf5\x55\x28\xbf\xa7\x0e\x4e\x62\xe7\x6c\xf3\x28\x72\xd2\x28\x70\x88\xe3\xa6\x99\xe7\x33\x9f\xba\x7e\xe8\x25\x88\x23\x27\x1c\x79\x31\x75\x5c\x92\x7b\x61\x44\x7e\x02\xa8\xb3\x8d\x69\x32\xf6\xbc\xe1\x30\x9a\x52\x77\xec\x67\x84\xc6\x38\x8c\x28\xf3\x9d\xcc\x8d\x82\x28\x1d\x7a\x29\xe1\x38\x25\xd3\x53\x28\x3b\x5b\xee\x25\x11\x0e\x69\x98\x0f\xdd\x09\xa3\x23\x27\xf6\xfd\x69\xc4\xfc\x21\x09\x88\x3f\xa4\x41\x16\xf9\xd3\xd1\x10\x23\xec\xa0\xbf\x94\x6b\xd6\x7a\x7d\x84\x68\x8a\xaa\x62\xc5\x12\xc5\x62\x69\xf4\xaf\xe1\x4d\xff\x22\xde\x5f\x99\xf0\xff\x06\x9c\x38\x5e\x9f\xf8\x5e\x9f\x44\x7d\xff\xbb\x6e\xb6\x27\x70\x9e\x6e\xd3\xcb\x51\xfa\xb4\x8c\xdf\x3d\x18\x7d\xb3\x7b\xb8\xc8\xee\x66\x8a\x79\xb7\xf5\x3c\xf1\x4c\xba\xd6\x01\xab\x08\xf9\xb4\xb9\x48\xe8\xcb\xd9\x0f\xe4\xfd\x3e\x89\xfc\x3e\x75\xc3\x53\x1b\xdc\x94\x94\xcf\x4b\x35\x11\x6c\x7e\xf5\xe0\x2d\xee\xd7\xe1\xe3\xc5\xb2\x5e\xdc\x6e\x64\xb4\x91\xd3\xb9\x7e\xb3\x7c\xba\x48\x2f\x84\xcb\x92\x68\xfb\x73\xc4\x9b\xec\x9c\x04\x9c\xfe\x0d\x84\xff\x04\x70\xe2\x06\x74\xc2\x87\x79\x14\x84\x31\xf5\xdc\x09\xf5\xf2\xc4\x99\x8c\x3c\xea\x67\x14\x89\x93\x38\x11\xa5\x2e\x0f\xc7\x3f\x05\x3c\x24\x91\x33\x0e\x43\x97\x38\x19\xf2\x28\x19\xd2\x28\x61\x91\x43\x27\xdc\x89\xa7\x79\x42\xc7\xd3\xc0\xc3\xd8\x09\xf9\x69\xc0\x49\xe4\x92\xd0\xf1\x22\x12\x78\x51\x8e\x79\x8e\x5e\xec\x39\x53\x77\x9c\x24\x99\xcb\xc2\x94\xa7\xa9\xc3\xfd\x24\x99\x76\x80\xdf\xca\x5a\x1b\xfc\x0e\xf1\x4c\x2e\x6a\x66\xf8\xf2\xd7\xe8\x76\xff\x22\xdd\xfb\xdd\xe1\xb7\xbb\x0f\xe3\x0f\xc0\x15\x32\x83\xa0\x3a\x53\x2d\xe1\x8d\xce\xef\xff\x69\x1d\xbd\x0d\xc0\x29\xe0\xdd\x7f\x96\x77\x27\x73\x63\x32\x09\xa9\x4b\xfd\x11\x66\x23\x8f\x4c\xbc\xc8\xf1\xdd\x49\x18\xd2\x28\x62\x51\x3c\xa5\x13\x97\x10\xe2\xff\x94\x77\x3a\x8a\x9c\x29\x19\xb3\x7c\xcc\x42\x96\x8c\x31\xa5\x23\x12\xfa\x99\x37\xf4\xdc\x24\xf2\x23\x2f\x74\x27\x84\x84\xc4\x3d\xcd\xbb\x17\xa7\x18\xbb\x8e\x33\x72\x83\x51\xee\x53\x37\x4a\xa7\x41\x3c\xf1\x46\x5e\xec\x07\xce\x74\x1a\xe5\xe1\x34\x08\xe9\xc4\x3b\x3a\xc5\xd8\x43\xcb\x31\xef\x30\xfe\x00\xd7\x1f\xee\xe0\x7e\x3e\xf9\x9f\x1e\x00\x96\x29\x53\x9c\x65\xa8\xa4\x9d\xf5\x4b\x25\x40\x9c\x93\x6c\x7e\x85\x0f\x89\xa3\x3e\xa1\xb4\x4f\xc8\xc9\x76\x99\x2c\xdc\x09\x4f\x8c\xfa\xf8\x30\xda\x6e\x5e\x82\xe7\x40\xdf\xc5\xe2\x69\x7e\xfb\x62\x5e\xe2\x71\xb8\xbb\x7f\xa9\x87\xb3\xdb\xc9\xf4\x45\xdd\xcb\x87\xb3\xef\x77\xa0\x1e\xed\x53\x4a\xfa\x84\x9c\xec\xf8\x97\x17\x1b\xb1\xfd\x03\xab\xd5\x1f\xc9\xc3\xe7\xe7\x77\x97\x65\xf5\x66\x9e\xbc\x1b\x7f\x7a\xc9\x43\xbc\xb8\x92\x81\x51\x52\x2c\x9e\xb6\x65\x98\xf8\xb7\x3f\x27\xb4\x6c\xa3\x7b\x8a\x50\xf2\xcf\x12\x9a\x4c\x3d\x3f\xe0\x24\x70\xa3\x80\x05\x5e\x9e\x79\x53\x2f\x0d\x62\x96\x13\x97\x45\xc1\x38\x77\x86\x7e\x40\x13\xe6\x38\x3f\x25\x34\x70\xc3\x61\x34\x72\xc7\x34\x49\xdc\x11\xa7\x4e\x30\x8e\x3d\x9f\xc4\xa9\xef\x45\x31\x75\xa2\x98\xc7\x93\x20\x8c\x63\xe7\x34\xa1\x43\x1f\x3d\xea\x66\x23\x1e\x7a\x4e\x3a\x1c\x45\x4e\x1e\x3b\x01\x71\x5d\x24\x7e\xe0\x90\x3c\x8e\x9c\x38\x8e\x5c\x3f\xf8\x86\xd0\x2f\x48\x1d\x01\xf9\xaf\x86\xf1\xef\x46\xf1\xbf\x20\xfe\x7b\x82\xf8\x0a\xc6\xcc\x30\x98\x1b\xa9\xd8\x02\x7b\xba\xfd\xdf\x7e\xa6\xcf\x98\x59\x36\x91\x29\xec\xc7\xe0\x78\x08\xb9\x28\xb0\x07\x50\x33\xb3\x1c\xc0\xb9\x29\xeb\xf3\x2f\xd7\x05\xff\x9b\x31\xc3\xfa\xcd\xcc\x2c\xb5\xba\x23\x59\xe5\x62\xb1\x52\xcc\x08\x59\x1d\x36\xe0\xcd\xe8\xfc\xd7\xb7\x69\x05\xbe\xdb\x2d\xe1\x5c\xae\x2a\xa3\xe1\x19\x77\xd0\x79\xd1\x63\xdd\xa0\xdd\xe7\x19\x77\x76\x18\x3b\xc5\xfd\x23\xbb\xf6\x6d\x65\x50\xe5\x8c\x23\x6c\x2c\x40\x0d\x08\xc9\xec\x2d\xb0\x2a\x83\x19\x9d\xc1\x1c\xd5\x1a\x55\x73\xb4\xc1\xca\x9e\x5d\x7a\xf6\x54\xf2\x46\x6a\x53\xb1\x12\x07\x70\xf8\xc4\xef\xbd\x82\x99\x54\xa6\x93\xb1\x12\x3f\x5e\x6a\x27\x0d\x20\x72\x22\x6a\xb7\xb7\x55\xfa\xda\xc8\xd7\x35\xa2\x02\x7e\x1c\x35\xdd\xab\x69\xdd\x06\x69\x5e\x23\x17\xf9\x0e\x26\x5b\xd3\x7c\x11\xc0\xdb\xd9\x91\xb5\x56\x14\x38\xab\x20\x45\x50\xc8\xf8\x12\x33\x60\x06\x44\x0e\x29\x2e\x45\x95\xc1\x75\x72\x67\x65\xb0\x5b\xfd\x76\x36\x80\x4d\x7f\xdb\xdf\xf5\x5f\xda\x14\x58\xab\x57\x1a\xb3\x43\x21\x58\xbf\x0b\xb6\x43\x65\x13\xd1\x98\xdb\x94\x71\x33\xfb\x4e\x94\x28\x57\x8d\x9b\x15\xc8\x1a\xab\xee\x16\xa7\x42\xde\x58\x6d\x4f\x77\xd6\x19\xdd\x83\xfd\x70\xb7\x64\x00\x67\xae\xa3\x1b\xec\x6e\x56\xb8\xc2\x6f\xdc\x6d\x76\x67\x7a\x57\xf1\xa5\x92\x95\x5c\x69\x7b\x60\xe4\xa8\xb5\xa8\x16\xbd\xcf\x76\x41\x1b\x8c\xf6\x0e\x4a\xb7\xae\xaf\xca\x14\x95\x6d\x8e\xb6\xe6\x51\xe9\x73\x2e\x2b\x6d\xbb\x66\x77\xfc\xdc\x88\xa2\xb0\x71\x61\x45\x21\x39\x33\x6d\x64\xb4\x61\xca\xac\xea\x1e\xd8\xf5\x8f\xed\x42\xdb\x3f\x9d\x46\x7f\xaa\x10\x35\xac\x6a\x18\xcd\xee\x81\xef\x78\x81\xba\x75\xb6\xdd\x02\x84\x86\x0d\x13\xcd\xe5\x95\xb5\x18\xd7\x68\x49\x82\xee\xf1\x23\x13\x8d\xbf\x57\xf3\xb6\xff\x34\x4d\xbc\xb3\x51\xa1\x51\x02\x75\x63\xcc\xa6\x0b\x37\x03\xc3\xb4\x6d\xe2\xf6\xdf\x6d\x3b\xa1\xe9\xe5\xbd\xa3\xa6\xa7\x9b\xfc\x0b\xfe\x75\xc4\x7a\xfb\x96\xd7\x41\x82\x05\xda\x6e\xb6\x59\x0a\xbe\x3c\xb4\x43\xe8\x58\xb7\x69\x59\x69\xdc\xbf\x47\xa4\x8d\x60\x77\x20\xcf\x40\x54\xcd\x20\x5f\x69\x23\xcb\x6e\x93\x7d\x21\x76\xf7\x7c\x5d\x89\x5d\x37\xcc\x9f\xd9\xc6\x7b\x76\xb8\xcd\x6b\x6a\xbc\x13\x3e\xec\xcb\x0b\x81\x95\x69\xe1\xfc\x6d\x63\xc9\xfc\xbc\x12\x0a\x61\xa3\x41\x2a\x10\x35\xef\xae\xf8\x58\x5a\xa0\xfd\xc9\x9b\x4f\x81\x36\x9a\xf6\xc8\x6f\x17\xde\xdf\xbe\x1f\xc0\xd2\x98\x7a\x70\x7e\x6e\xf3\x57\x2c\xa5\x36\x83\xd8\xf7\xfc\x66\xef\x92\x6d\x45\x69\x5d\xec\xe2\xb9\x60\xd6\x27\xc1\x1b\xbd\x9a\xed\xf6\x01\x56\xac\xd2\xac\xa1\xd3\x7a\xba\x41\xd1\xac\xa6\x0e\x5c\x6c\x50\x40\x25\x37\x3d\xb0\x5a\x17\x4c\xcf\xec\xea\x01\x50\xe7\xf0\xd7\x4c\xbd\x60\x1a\x0a\x51\x8a\xee\x5d\x91\x89\x3c\x47\x65\xbd\x3b\x64\x48\xd6\xb8\xaf\x5a\xb0\x76\xbc\x6f\x66\xef\x6f\x27\x47\xcd\xa7\x4d\x83\x58\xa7\x69\x47\x93\x2c\xbb\xc4\xdd\x00\xdc\xe3\xc1\x5b\x5c\xcb\x67\x6c\xc6\x7d\x7f\x3f\xdc\xbe\x29\x46\xb2\x2c\x85\x6d\x1d\xdf\x8c\xcf\x14\xee\x1f\x91\x2f\x52\x55\x6e\xae\x44\x65\x06\x10\x7f\xf1\x63\x5f\xbb\x46\x36\x08\xb7\xf1\xa9\xbe\xe4\xec\x38\x52\x5d\x76\xb2\xac\xbd\x8c\x65\x90\x16\x92\x3f\x37\x6d\xb1\x4d\x12\x18\x25\x16\x0b\x54\x98\xb5\x95\x6e\x70\x6b\xf6\xf4\xb7\xd5\x1e\x38\xb6\xdc\x4f\x6d\xac\x90\x65\x20\xab\x62\x77\x14\xbc\xc3\x8d\xf4\xde\xa4\x2f\xd2\xb7\xc8\xb2\xaf\xe5\x89\xdf\xa9\x5f\x5b\xc6\x8e\x6d\xaf\xa5\x2c\x6c\x46\x0f\x15\x67\x24\x68\xac\xb2\x6f\x60\x90\xeb\xa6\xc3\x95\x6c\x7b\x28\x3c\xda\x45\xea\xc7\x92\xc2\xbe\x2b\xd6\xac\x68\x74\x77\x6d\x57\x60\xd6\x40\xbe\x52\x0d\x0f\xc7\x2b\x96\x4c\x43\x8a\x58\x41\x86\x06\xb9\x69\xc2\xb4\x17\xb0\xfb\xd9\xf7\x3d\xed\x3c\x18\x0b\xdd\xd4\x41\xa3\xa8\x65\xf9\x5d\x1d\x69\xc8\x24\x54\xd2\x80\x5e\xd5\xb6\x23\x83\xd9\x36\x16\xb1\xda\xc2\x6c\xb6\x33\x29\x8b\x84\xdb\x6e\x39\xa9\xac\x52\x36\x00\xa3\x56\x68\xbb\x08\xab\x76\x90\x61\xba\x5a\x2c\xba\x4e\x6d\x8b\xbb\xe9\x8b\x0b\x09\x76\x93\x5e\xf3\xb4\x6d\x22\x75\xad\x64\xde\x42\xbe\x5f\x62\xdf\x01\x76\x74\x00\x39\x2b\x34\xf6\x7a\x2d\x75\xdd\xe5\x7b\xad\x90\x77\xf0\x35\x1b\xfe\x5f\x00\x00\x00\xff\xff\x56\xad\xf2\xc6\x71\x18\x00\x00")

func goCentrifugeBuildConfigsDefault_configYamlBytes() ([]byte, error) {
	return bindataRead(
		_goCentrifugeBuildConfigsDefault_configYaml,
		"go-centrifuge/build/configs/default_config.yaml",
	)
}

func goCentrifugeBuildConfigsDefault_configYaml() (*asset, error) {
	bytes, err := goCentrifugeBuildConfigsDefault_configYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-centrifuge/build/configs/default_config.yaml", size: 6257, mode: os.FileMode(420), modTime: time.Unix(1555403231, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _goCentrifugeBuildConfigsTesting_configYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xb9\x8e\xe4\x36\x10\xcd\xf5\x15\x04\x1d\x4c\xd2\x07\x4f\xf1\xc8\x1c\x1a\x0b\x3b\xb1\x81\x8d\x8b\x64\xb1\x87\xe8\xd1\x61\x92\x9a\xd9\xc6\x62\xff\xdd\x50\x6f\x8f\xbd\x99\x67\xb3\xaa\xd2\x3b\xaa\xa4\xa7\x88\x73\xaf\x25\x6f\x17\xfc\x03\xfb\xdb\x52\xaf\x9e\x74\x6c\xbd\xcc\x97\x01\xfb\x33\x56\xdc\x26\x3f\x10\x02\x31\x2e\xdb\xdc\xdb\x5e\x13\x32\x41\x99\x3d\xb9\x97\x84\x5c\xf1\xe6\xc9\xd3\x57\x0a\x29\x55\x6c\x8d\x7a\x6a\x5d\x60\x60\x47\x6d\x65\x54\x4a\x29\x88\x39\x19\x1e\xd4\x28\x91\x25\x19\xb5\x06\xe4\x8a\x0b\xd0\xf4\x40\x63\xbd\xad\x7d\xa1\xfe\x2b\x8d\x65\x7d\xc6\x4a\x3d\x05\x6c\x47\x2e\xec\x31\xf6\xba\x03\xee\xe3\x8e\x5f\x3a\xf5\x34\x1a\xe3\xb2\x95\xc6\x25\x63\x58\x72\x22\xe6\xc8\x53\x4a\x0a\x6c\x96\x3c\x69\x60\x90\xa2\xcd\x02\x58\x10\xc0\x15\xe3\xd2\xb0\x24\x47\xc9\xb2\xb4\x91\x45\x0b\xff\xea\xad\x50\x61\x6a\xbb\x6d\x79\xa5\x9e\xca\x31\xf2\xd1\xa2\x91\x21\x3b\xcb\x32\x1a\x1d\x98\x11\x26\x5b\xc7\xc0\x70\x48\xf4\xdb\x81\x5e\x53\xa6\x9e\xb6\xfb\xc2\xf4\xde\xfe\x27\x92\xae\x2f\x38\x53\x2f\xc5\x81\xce\xd4\x8b\x51\x70\xa5\x0e\x74\xa5\x9e\x1f\x68\xa5\xde\x1e\x68\x83\x97\xfd\x80\x84\x3c\x20\x1f\x51\x46\x67\xb9\x53\x2a\x71\x8c\x20\x82\x0d\xc2\xa0\xc2\x11\x59\xd0\x21\x07\x25\x03\x32\x69\x46\xd0\xc9\x5a\xeb\x32\x8c\xc6\x81\xb0\x5c\x88\x7d\x91\x09\xe2\xfe\x2a\x22\x17\x36\x58\xae\xb5\xd6\x01\x38\x42\x32\x11\xd0\xb1\x91\xa1\xb5\x4a\x40\x8e\x60\xa5\x1e\x13\x1b\x95\xd6\x21\x39\xd0\x46\x8b\x00\x63\x8e\x91\x39\x81\x79\x57\x2a\x89\x7a\xaa\x34\xb2\x91\xc1\x78\x4c\x02\xf0\xa8\x64\xb0\x47\x27\x44\x3e\x2a\x65\x85\x53\xce\x25\x69\x12\x3d\xd0\x57\xac\xad\x2c\xfb\x91\xdf\x9e\x1e\x1f\x7e\x85\xd6\xde\x96\x9a\x3c\x79\x7a\x1f\x3d\x32\xe0\xc9\x47\x23\x30\x0c\x25\xe1\xdc\x4b\xbf\xfd\x96\x3c\xa1\xec\xcb\x87\xb3\x33\x0c\xbf\x90\x5f\x1f\xa9\xdc\x33\x48\x5a\x5f\x2a\x5c\x70\xf8\x31\xaa\x57\xbc\xed\x63\xf4\xe4\xdc\xa7\xf5\xfc\xfe\x68\x18\xfe\xde\x70\xc3\x1d\x31\x6f\xd3\xe7\xa5\x5e\xb1\x36\x4f\xc4\x40\xc8\xdb\xbd\xf9\x0c\xa5\xff\x55\x26\xfc\xfd\x4f\x4f\xf8\x30\xec\x32\x3b\x78\x15\xeb\xf7\x1f\x60\xdd\xc2\x4b\x89\x9f\xf6\xe4\x9f\x4e\xe7\xd3\xe9\x1c\xb6\xf2\x92\xce\x15\xdb\xb2\xd5\x88\xed\xbc\x8a\xf5\x13\xde\x4e\xeb\x16\x4e\x2b\x4e\xdf\x39\xb5\xbc\x42\xc7\xff\x27\x5d\x77\xe2\x9d\xd4\xca\x65\x2e\xf3\xe5\x83\x9e\x0f\xf4\xcf\xfb\xfe\x40\x7c\xf7\x1e\x60\x8e\xcf\x4b\x7d\x98\xaf\x15\xe3\x32\x4d\xa5\x7b\xd2\xeb\x86\xff\x04\x00\x00\xff\xff\x1f\xaf\xbe\x5d\x34\x04\x00\x00")

func goCentrifugeBuildConfigsTesting_configYamlBytes() ([]byte, error) {
	return bindataRead(
		_goCentrifugeBuildConfigsTesting_configYaml,
		"go-centrifuge/build/configs/testing_config.yaml",
	)
}

func goCentrifugeBuildConfigsTesting_configYaml() (*asset, error) {
	bytes, err := goCentrifugeBuildConfigsTesting_configYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-centrifuge/build/configs/testing_config.yaml", size: 1076, mode: os.FileMode(420), modTime: time.Unix(1554467633, 0)}
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
	"go-centrifuge/build/configs/default_config.yaml": goCentrifugeBuildConfigsDefault_configYaml,
	"go-centrifuge/build/configs/testing_config.yaml": goCentrifugeBuildConfigsTesting_configYaml,
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
	"go-centrifuge": &bintree{nil, map[string]*bintree{
		"build": &bintree{nil, map[string]*bintree{
			"configs": &bintree{nil, map[string]*bintree{
				"default_config.yaml": &bintree{goCentrifugeBuildConfigsDefault_configYaml, map[string]*bintree{}},
				"testing_config.yaml": &bintree{goCentrifugeBuildConfigsTesting_configYaml, map[string]*bintree{}},
			}},
		}},
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
