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

var _goCentrifugeBuildConfigsDefault_configYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x58\x59\x73\x1b\x37\x12\x7e\xe7\xaf\xe8\xa2\x5e\x92\xaa\x1d\x6a\xee\x83\x55\xa9\x2d\x9e\xb6\x63\x59\xa1\xae\x28\xd6\xcb\x1a\x83\xe9\x21\x61\x0d\x81\x31\x80\xe1\xe1\x5f\xbf\x05\xcc\x90\xa6\x2d\x4b\xd9\x4d\x2a\xd9\x87\xf5\x8b\xcc\x1e\x74\xa3\x8f\xaf\x3f\x34\x70\x06\x53\x2c\x49\x53\x69\x28\x70\x83\x95\xa8\xd7\xc8\x35\x68\x54\x9a\xa3\x06\xb2\x24\x8c\x2b\x0d\x92\xf1\x47\xcc\xf7\x3d\x8a\x5c\x4b\x56\x36\x4b\xbc\x44\xbd\x15\xf2\x71\x08\xb2\x51\x8a\x11\xbe\x62\x55\xd5\xb3\xc6\x18\x47\xd0\x2b\x84\xa2\xb3\xcb\xdb\x95\x0a\xf4\x8a\x68\x98\x1c\x2d\xc0\x9a\x30\xae\x8d\xfd\xde\x61\xc9\xb0\x07\x70\x06\x17\x82\x92\xca\xba\xc0\xf8\x12\xa8\xe0\x5a\x12\xaa\x81\x14\x85\x44\xa5\x50\x01\x47\x2c\x40\x0b\xc8\x11\x14\x6a\xd8\x32\xbd\x02\xe4\x1b\xd8\x10\xc9\x48\x5e\xa1\x1a\xf4\xe0\xa0\x6f\x4c\x02\xb0\x62\x08\x41\x10\xd8\xff\xa3\x5e\xa1\xc4\x66\xdd\x45\xf0\xa6\x18\x42\x1a\xa4\xed\xb7\x5c\x08\xad\xb4\x24\xf5\x02\x51\xaa\x56\xd7\x81\xfe\x39\xab\xc3\x73\xcf\x4f\x06\xee\xc0\x1d\x78\xe7\x9a\xd6\xe7\x41\xea\xbb\xfe\x39\xab\x4b\x75\x7e\xb5\xbe\xbd\xda\xe5\xdb\xc7\xe6\xe1\xfd\xfb\x69\xd9\x7c\xbe\xcd\x77\xb3\xd1\x35\xde\x5e\x4e\x2e\xc4\xe7\xfd\x3e\x8a\xd2\xcd\x15\x5f\xfe\xba\x59\xbc\xfb\x78\xf1\xfe\xb1\xff\x3b\x46\x83\x83\xd1\x5f\xcb\x78\x76\x19\xaf\x1f\x3f\xdd\xe3\xc7\xfb\xb7\xf7\xfe\xa7\x45\xe3\xc5\xbf\xd5\xc5\xab\xe0\xf1\x67\xe1\xdd\x06\xeb\x15\x59\x2d\xc6\xd1\x0d\x46\xdc\x6b\x8d\x1e\x52\x35\x3a\x64\xaa\x0d\xc0\x84\x8f\x5c\x33\xbd\x9f\x13\xaa\x85\xdc\x0f\xa1\xdf\xef\xbe\x10\x4e\x57\x42\x5e\x63\x2d\x14\xfb\xe6\x13\xe3\x1b\xc1\x28\xde\xf1\x9a\x98\xf4\xf5\xfb\x3d\x5b\x9d\x77\x84\xf1\xef\x62\xa5\x2b\x22\xfc\x70\xdd\x82\xe5\xc7\x1e\x9c\x82\xa3\xf5\xe5\x0c\x2e\x9b\x35\x4a\x46\xe1\xcd\x14\x44\x69\x81\x72\x02\x89\xce\xc6\xb1\x66\x91\xd7\x69\x8d\x0f\x85\x81\x8a\x29\x6d\x34\xb9\x28\xf0\x29\xa6\x6a\x29\x36\xcc\x7e\x10\xd6\xf6\x89\x03\x07\x47\x7f\xb7\xd0\x41\x34\xf0\xfd\x68\xe0\xbb\xee\x20\xf4\xbf\x2d\xb6\xe7\x4f\x83\xb7\x42\xdc\x5f\x30\x46\xaf\x7e\xdd\xde\xae\x6e\xc7\xef\xe3\xdd\x5b\xba\x10\x17\x65\x7c\x7d\xf5\xfe\xe7\x79\xbd\x2d\x3d\x99\x44\xdb\x8b\x9d\xff\x70\x1d\xd4\x93\xc2\xeb\x7f\xcf\x7c\x1a\x0f\x7c\xcf\x7d\xce\xfc\xd5\xc3\xbb\x51\xfa\x6a\xf1\x5a\x6e\x66\x0f\xe3\x6c\x5b\x3c\x8a\x3b\x3a\x1a\xad\x27\x0f\xaf\xeb\x0c\xf7\xfb\x87\xf0\x66\x96\x2e\xe7\x32\x58\xdd\x5e\xfe\xd6\xef\x72\x34\xeb\x80\x7d\xac\xc4\x9b\x29\x38\xd0\x55\xe3\x39\xe8\x87\x9d\xf2\x05\x31\xe9\x81\x02\xeb\x4a\xec\xb1\x80\x9b\x35\x91\x1a\x26\x1d\xa2\x14\x94\x42\xda\x84\x2e\xd9\x06\xf9\x57\xa9\x7c\x8a\x3a\x78\x16\x76\xee\x2e\xc3\x7c\x3a\x4d\xc7\x49\xee\x63\xe2\x7a\x34\x2b\x93\x22\x73\xe7\xa9\x47\xb2\xc4\x8b\xa3\xd8\xa7\xa3\x49\x14\x85\x79\xf0\x02\x40\xdd\x5d\x10\x63\x90\xa7\x79\x56\x96\x65\xec\x8e\x3c\x74\x03\xaf\x2c\xdc\x24\x4e\x3c\x3f\x4b\x0b\x37\x4a\xc2\x88\x96\x63\xe2\x3f\x07\x65\x77\x37\xf3\x62\x3f\x1c\xcd\xc2\x24\x8e\xb1\x0c\xc8\x28\x0c\x5c\xd7\xf3\xfd\x69\x34\x1b\x91\x88\x84\x38\x4a\xca\x49\x94\x25\x69\x07\xfa\xb7\x62\x43\xda\xa8\x4f\x20\x9a\xa3\xe4\xa4\x5a\x21\x5b\xae\x74\x07\xa1\xb3\xb3\xb3\x2e\x9f\xad\xc6\x7c\x74\xd5\xfd\x76\xe0\xde\xd0\x14\xe3\x65\x23\x09\xec\x45\x03\x4b\xc3\xaf\x1c\x50\x4a\x21\x0d\x38\x6e\x57\x4c\x81\xc4\x4f\x8d\xd9\x85\x29\xe0\x42\x83\x6a\xea\x5a\x48\x8d\x05\xe4\x48\x49\xa3\xd0\x68\x4a\x8b\x7d\xb3\x44\x36\x9c\x1b\x8e\xb4\x0c\xa8\x34\xd1\xa6\x01\x1a\x23\x1a\xc0\x75\xc3\x5b\xb9\xe3\x74\xb2\x9f\x88\xa4\x2b\xb6\xc1\x41\xff\x1f\x9d\x53\x00\x5b\xd3\x3f\x5a\x40\x21\xfe\x69\x35\x08\x54\x96\x7d\x6b\x22\x99\xde\xb7\x1b\x59\x2b\x8f\x36\x1e\x5c\x0e\xdb\x9f\x1f\xba\x05\x8e\x43\x57\x84\xf1\x9f\xda\xcf\x8e\x63\xbc\xfd\x29\x70\x03\x37\x04\xc7\xd9\x12\x59\x77\x7f\x9c\x9c\x48\xc9\x50\x42\x14\xa7\xae\xeb\xba\xe0\x38\x5c\x38\x84\x53\x86\x5c\x3b\x79\x25\xe8\xa3\x6a\x65\x0a\xe5\x06\x9d\xca\x24\x15\x1c\x67\x4d\x76\x4e\x6d\x5a\x14\xfc\xc8\x28\x29\x4e\x6a\xb5\x12\xba\x13\x5a\xd9\x9a\xf1\xaf\x7e\x1a\x9f\x09\xd5\x6c\x83\xe0\x38\x06\x9a\x26\x45\xa2\x2c\x9f\x66\x02\x1c\xa7\xc8\x1d\x2a\xd6\xb5\x59\x2f\x38\x28\x55\x98\x90\x08\x5d\xa1\xa3\xd8\x67\x84\xd0\xcd\x62\x70\x9c\x8f\x4a\x70\x59\x53\x67\x25\x94\x56\x40\xaa\xea\x44\xc6\xb8\x46\x59\x12\x8a\x46\xfe\xe1\xeb\x72\x3f\x4d\xe6\xf7\x2a\x3f\x36\xe1\x63\x61\x3a\x89\x63\xeb\x88\x16\x70\x8f\xf9\x8d\x91\x6b\x05\x36\x27\x12\x4a\x29\xd6\xd0\x70\x2d\x1b\x65\x20\x21\x24\x5b\x32\x3e\x84\xc1\xa0\xff\x6c\x3d\x4d\xcb\x3e\xa9\xe5\x07\xc7\x69\xb8\x22\x25\x3a\xb8\xab\x85\xc2\x0f\x50\x56\x64\xf9\x0d\x80\xff\x3b\x9e\xf6\xff\x24\x4f\x7f\xd5\x4b\xff\x31\x53\x7b\x6e\x38\xf0\xa2\x70\xe0\xa5\x83\xe8\xc9\xb1\x7c\xa0\xd2\x85\x8a\x19\xc1\xbb\x66\xfe\x70\xd9\x78\xaf\x76\x1b\xb5\x1f\xdf\xde\xc8\x5b\x95\x6d\xf4\x38\xce\xf5\xbb\x11\x7f\x3d\x17\x17\x1f\xf3\xc7\xcf\x13\xd2\xff\x8e\xf9\x68\xe0\xa5\xd1\xc0\x0f\x92\x67\x37\x98\xbc\xa2\x5b\x76\xfb\x51\xbc\xbd\x7f\x5d\x8e\x49\x98\xfa\x77\x0b\x4d\xf0\x6e\x77\x79\xb1\x2d\xd2\xcf\x39\x1f\x7b\x37\xc9\x16\x47\x0f\x77\xbb\x87\x97\xb9\xda\x92\xc6\xb3\x4c\xed\xff\x05\x54\xfd\x02\x53\x07\x71\x9c\x87\x9e\x17\xfb\x24\x0a\xca\x22\x89\x8a\x2c\x2a\x02\xaf\x98\xc6\x85\x37\x09\x8b\x34\xc8\xa7\x61\x10\x8f\xc7\xf8\x22\x53\x27\xd1\x98\xd0\x71\xe6\x8f\x02\x17\x93\x59\xe9\x4d\x83\x94\xfa\x24\xf0\xc6\xe1\xdc\xcf\x82\xa8\x08\xb2\x79\x9c\xa4\xf4\x05\xa6\x4e\x93\xd4\x45\xcf\x0b\x03\x37\x88\xf3\x90\x66\x49\x56\xce\xfc\x28\xf0\xfc\x94\x84\xa9\x9b\x84\x6e\x16\x94\x5e\x96\x26\x1d\x53\x5f\x8b\x5a\x69\x7c\xc2\xd5\x85\x58\xd6\x44\xd3\xd5\x1f\x9b\x42\x82\x3f\x89\xee\xc3\xee\xf0\xc3\xed\x2f\xd3\x5f\x80\x4a\x34\x54\x2d\x3b\x57\x0d\xc2\xad\x9d\x1f\x9f\x05\xfc\x5f\x3e\x9c\xfc\xef\xc6\x93\x36\x09\xcf\x81\x3e\xf8\x7b\x31\x3f\x4f\x63\x2f\x49\x72\x3f\x89\x66\xe1\x9c\xe6\xfe\x28\x8f\x46\xc1\x9c\xce\xa3\x59\x18\x13\x92\x47\x5e\x99\x04\xd3\x84\xc4\x2f\x62\xbe\x8c\x03\xb7\x48\xf3\x59\x14\xc7\xe1\x24\x99\x8f\x3d\x3f\x08\x27\xb3\x32\x4c\xf3\x78\x94\xf8\x19\x9d\x61\x30\x4e\xbc\x69\xfa\x3c\xe6\x33\x44\x4c\xe7\xf9\x2c\x98\xf9\x19\xcd\x62\x3a\xc9\xc2\x70\xee\x22\x4e\xbc\x28\x0b\x4b\xf4\xf2\x38\x08\x8a\x91\x69\xbd\x2f\x23\xb9\x19\xc1\xbf\x01\x3d\xae\x73\x22\x29\x29\x50\x0a\xf3\xe5\x0f\x61\xdf\x73\xff\x0f\x41\x69\x6e\xa6\x07\xfc\x7c\x07\x94\xde\xdf\x0b\x4a\xdf\x8f\x12\xe2\x8e\x93\x59\xe4\x05\x51\xea\xd3\x20\x1e\x7b\x7e\x42\xc3\x99\x5b\x78\xa1\x3f\x4f\xc7\x59\xea\x27\x01\x79\x79\x64\x8e\x7d\x8c\xd2\x71\x90\x65\xa4\xa0\x6e\x10\x60\x8c\xa1\x1b\x66\x49\x3a\xf5\xe2\x69\x14\x7a\x7e\x50\x86\x13\xcf\x1d\xd3\xe7\x41\x39\x8e\xf3\x24\x9e\x4e\xa6\xf1\x28\x8c\xe2\x69\x54\x64\xf3\x72\x9e\x4c\xbd\x74\x1e\xc5\xfe\x34\x48\xa3\xb0\xf4\x3c\x12\x4d\xb2\x7e\xaf\x77\x06\x53\xa2\x09\xdc\x68\x21\xc9\x12\x7b\xaa\xfd\xdb\xde\xee\x17\x44\xaf\x6c\x66\x2a\x73\x87\x9c\x8e\xa1\x64\x15\xf6\x00\x6a\xa2\x57\x43\x38\xd7\xeb\xfa\xfc\xcb\x2b\xc3\xbf\x0a\xa2\xc9\xc0\xae\x2c\x72\x63\x77\x22\x78\xc9\x96\x8d\x24\x76\x4a\x3a\x6c\x40\xad\xf4\xe6\x8f\x6f\xd3\x1a\x78\xb2\xdb\x88\x52\xd1\x70\xad\xe0\x11\xf7\xd0\x45\xd1\x23\x9d\xd0\xec\xf3\x88\x7b\x23\xc6\xce\xe2\xe1\x93\xd1\x7d\x73\x1c\x0b\xb7\x06\x40\x16\x08\xa3\xc5\x1b\x20\xbc\x80\x85\xbf\x80\x9b\x76\xa6\x33\x07\x09\x72\x73\x52\xf4\xcc\x19\xf0\x5a\x28\xcd\xc9\x1a\x87\x70\x7c\x19\xe8\x9d\xc1\x42\x48\xdd\x99\x31\x26\xbe\xaf\x6a\x16\x0d\x21\x75\x53\xdf\x6c\x6f\xfa\xd4\xd1\xc2\x0e\xc6\x40\x4f\xb3\xa6\x7a\xb5\x5f\xb7\x49\xba\xa9\x91\xb2\x72\x0f\xb3\x9d\xb6\xf3\x17\xbc\x59\x9c\x78\x6b\x07\x46\x4a\x38\xe4\x08\x12\xcd\x4c\x5c\x00\xd1\xc0\x4a\xc8\x71\xc5\x78\x01\x97\xa3\x5b\x63\x06\x3b\xed\x37\x8b\x21\x6c\x07\xbb\xc1\x7e\xf0\xb9\x2d\x81\xf1\xba\x51\x58\x1c\x1b\xc1\xc4\x5d\x91\x3d\x4a\x53\x08\xeb\xae\x6d\x63\xbb\xfa\x96\xad\x51\x34\x36\x4c\x0e\xa2\x46\xde\x3d\xfe\x74\x13\xb1\x3d\x4b\xed\x94\xdf\x83\x83\xb8\x53\x19\x42\x3f\x70\x95\x85\xdd\x55\x83\x0d\x7e\x13\xae\xdd\x9d\xa8\x3d\xa7\x2b\x29\xb8\x68\x94\x39\x9e\x29\x2a\xc5\xf8\xb2\xf7\xc9\x28\xb4\xc9\x68\x9f\xae\x54\x1b\x7a\xb3\xce\x51\x1a\x7a\x34\x3d\x8f\x52\x9d\x53\xc1\x95\xe1\xcd\xee\xb0\xdf\xb2\xaa\x32\x79\x21\x95\x19\xf1\x75\x9b\x19\xa5\x89\xd4\x4d\xdd\x03\xa3\x7f\xdf\x2a\x0e\xa1\x0d\x6f\x2e\x11\x15\x34\x35\x4c\x16\x77\x40\xf7\xb4\x42\xd5\x86\xda\x6e\x60\x6e\x73\x5b\xc2\xec\x8b\x97\xf1\x17\x37\x68\x70\x04\xdd\xe7\x7b\xc2\x6c\xb4\xef\x6e\x5a\xf6\xb1\x24\xde\x79\x28\x51\x4b\x86\xf6\x56\x22\xb6\x5d\xb2\x09\x68\xa2\x0c\x89\x9b\x3f\xd7\xed\x02\xcb\xe5\xbd\x13\xca\x53\xb6\xfa\x8c\x7e\x9d\xaf\xde\x81\xf0\x3a\x88\x60\x85\x86\xcb\xb6\x2b\x46\x57\x47\x32\x84\x0e\xe9\xa6\x28\xe6\x56\xda\x9d\x23\xc2\xe4\xaf\x1b\x7e\x0a\x60\xed\xf5\x83\x36\x4a\x8b\x75\xb7\xc9\xa1\x0d\xbb\xc7\xc1\xae\xc1\x2e\x2d\xe2\xfb\x86\x76\xfb\xc7\x27\x40\xdb\xe1\x9d\xe1\xe3\xbe\xb4\x32\x17\xc6\x16\x9a\x3f\x6c\xd1\xde\x97\x99\x44\xd8\x2a\x10\x12\x58\x4d\xbb\x77\x41\x92\x57\x68\xfe\x4b\xed\xd8\xd5\x66\xd3\x8c\x57\x46\xf1\xee\xfa\x62\x08\x2b\xad\xeb\xe1\xf9\xb9\xbd\xa0\x99\x5b\xdd\x30\x8b\xc2\xc8\xee\xbd\x26\x3b\xb6\x36\x21\x76\xf9\x5c\x12\x13\x13\xa3\xd6\x5e\x4d\xf6\x87\x04\x4b\xc2\x55\x77\x6d\x64\x1c\xb6\xc8\xac\xb6\xef\xc2\xab\x2d\x32\xe0\x62\xdb\x03\x63\xeb\x15\x51\x0b\xa3\x3d\x04\xdf\x3d\xfe\xb3\x4b\x5f\x11\x05\x15\x5b\xb3\xee\xa4\x28\x58\x59\xa2\x34\xd1\x1d\x2b\x24\x6a\x3c\xf4\x2c\x18\x3f\x2e\xec\xea\xc3\x93\xe6\xc4\x8e\x91\xa6\xac\x07\x9b\x46\x3a\x2a\x8a\xb7\xb8\x1f\x42\x70\x2a\xbc\xc6\x8d\x78\x44\x2b\x8f\xa2\x83\xb8\x3d\x27\x26\x62\xbd\x66\x86\x38\xbe\x91\x2f\x24\x1e\x3e\x79\x5f\x4c\xf1\x52\xbf\x63\x5c\x0f\x21\xfb\x12\xc7\xa1\x73\xb5\xb0\x10\x6e\xf3\xc3\xbf\xd4\xec\x34\x53\x5d\x75\x8a\xa2\x7d\xc1\x25\x60\xaf\xfe\x96\x14\xdb\x22\x81\x96\x6c\xb9\x44\x89\x45\xdb\xe7\x1a\x77\xfa\x80\xfe\xb6\xd7\x63\xd7\x34\xfb\x73\x1b\x4b\x24\x05\x08\x5e\xed\x4f\x92\x77\x7c\xc6\x3e\xb8\xf4\xc5\xf4\x35\x92\xe2\x6b\xf3\x5e\xd4\x59\xbf\x34\x18\x3b\xf5\xbd\x16\xa2\x32\x15\x3d\x76\x9c\x16\xa0\x90\x17\xdf\x80\x41\x6c\x2c\xbf\xad\xc9\xee\xd8\x78\x7e\x97\xa9\xef\x9b\xb4\x0f\x08\x1b\x52\x59\xbb\xfb\x96\x15\x88\x71\x90\x36\xd2\xe2\xe1\x54\x63\x45\x14\xe4\x88\x1c\x0a\xd4\x48\xb5\x4d\xd3\xc1\x80\xd9\xcf\x9c\xf6\x7e\x17\xc1\x94\x29\xdb\x07\xd6\xa2\x12\xeb\x27\x7d\xa4\xa0\x10\xa7\xef\x4c\xa0\x77\xd6\x23\x52\x1b\x30\xeb\xdd\x42\x88\x6a\x44\x0d\x57\xce\xb8\xb1\x54\x0c\x41\xcb\x06\x0d\x8b\x10\xbe\x87\x02\xf3\x66\xb9\xec\x78\xda\x34\xb7\x65\xc5\xa5\x00\xb3\x49\xcf\x7e\x6d\x49\xa4\xae\xa5\x28\x5b\x90\x1f\x54\xcc\x09\x60\xa4\x43\x28\x49\xa5\xb0\xd7\x6b\x51\xd7\xbd\xd8\xd7\x12\x69\x07\x3e\xbb\xe1\xbf\x03\x00\x00\xff\xff\xae\x2d\xc0\x71\xa6\x18\x00\x00")

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

	info := bindataFileInfo{name: "go-centrifuge/build/configs/default_config.yaml", size: 6310, mode: os.FileMode(420), modTime: time.Unix(1554381575, 0)}
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

	info := bindataFileInfo{name: "go-centrifuge/build/configs/testing_config.yaml", size: 1076, mode: os.FileMode(420), modTime: time.Unix(1552945185, 0)}
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
