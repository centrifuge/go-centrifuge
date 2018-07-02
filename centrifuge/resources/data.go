// Code generated by go-bindata.
// sources:
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

var _resourcesDefault_configYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x94\xcd\x6e\xe3\x36\x14\x85\xf7\x7a\x0a\xc2\xfb\xc6\xfc\x13\x49\x69\x97\xb6\x69\x31\x6d\x1a\x38\xae\x3b\x69\xba\xe3\xcf\xa5\xcd\xb1\x44\x7a\x44\x2a\x91\xe7\xe9\x0b\xd9\x93\x99\xa6\x40\xdc\x7d\x31\x5a\x11\xe7\x10\xe7\xe2\x9e\x4f\x52\x70\x10\x4b\x28\xc7\x77\xae\x45\xa4\xaa\x2c\xc4\x32\x04\x3f\x6e\xe1\x0e\xca\x73\x1a\xf6\x2d\xfa\x2a\xad\xc7\x9c\x83\x8e\xbb\xd0\x75\x37\x65\xb7\x0e\x71\x0f\xe6\x58\xc5\xf3\xc5\xdc\xa2\x0a\xa1\x02\xb9\x84\xb8\x3d\x9d\x11\x0a\xae\x45\x8c\xb1\xd3\x19\xca\x0e\x06\x18\xfb\xcf\xb9\xf3\x3c\xc5\xd4\xd9\x33\x29\x95\x5c\x06\x7d\x58\x01\x0c\xb9\x3d\x69\x08\x7d\x87\x16\xcb\x70\xe0\x4b\x42\xe5\x15\xbe\xc2\x57\x64\x59\xec\x61\xc9\x14\xc5\x74\x19\x0e\x3e\x2f\xef\xfb\xcd\xfd\x64\x9e\xf7\xe3\x5f\x8f\x8f\x3f\xfa\xf1\xd3\xc6\x4c\x37\xd7\x6b\xd8\xdc\xfd\x70\x9b\x3e\x1d\x8f\x75\xad\x9e\xee\xe3\xf6\xfd\xd3\xea\xb7\x0f\xb7\x8f\xfb\xc5\x7f\xc6\xb2\x97\xd8\xf7\x5e\xdc\xdc\x89\x7e\xff\xf1\x01\x3e\x3c\xfc\xfa\x40\x3f\xae\x46\x22\xfe\x3c\xb8\x9f\xd9\xfe\x97\x44\x36\xac\xdf\xe9\xdd\xea\xfb\xfa\x77\xa8\x23\x39\xc7\xda\x14\xcb\xa0\x6d\xb9\x76\x6e\x80\x9c\xe1\xcb\x12\x2f\xfd\xfe\xa4\x6d\x49\xc3\xb1\x45\x8b\xc5\xbf\x9c\x35\x6c\x43\x2e\xaf\x2c\x1d\xed\x2e\x0d\xaf\x8c\x0a\xbd\x09\xe2\x5d\x2c\xb0\x1d\x74\x09\x29\xb6\x5f\x6a\x27\x98\x7c\xab\xfd\x9f\xb5\xe3\x49\x32\xca\x88\xe1\x0d\xc3\x42\x5b\x23\xb4\xc1\x92\x6b\x2c\x99\x54\x9e\x49\xe1\x6d\x63\x9c\xa1\x92\x88\x0b\x80\xf0\xe4\x84\xb2\x02\x4b\x63\x2c\xc5\x0a\x63\x2c\x54\xc3\x08\xc7\x0d\x11\x5a\x28\x4a\xbc\x93\x84\x19\xda\xd0\x37\x51\xe2\xc9\xd4\x92\x78\xe0\xd8\x01\x75\x82\x40\x53\x73\xe7\x1b\xab\x39\x34\x82\x78\xd1\x70\xa5\x3c\x61\xb5\xc6\x17\xa1\x7f\xfe\xfa\xbe\x02\xaf\xc9\xff\x80\xe9\x5b\xaf\x2c\x7f\x83\x36\xba\x80\x5b\xd5\x86\x51\x2f\x35\xf3\x1c\x73\x45\x3c\xa1\x8c\x71\xcc\x89\x90\xd8\x2a\x6b\x00\x4b\x2f\x9d\x6c\xec\x45\xdc\x35\xd7\xc0\x24\xf3\xb8\x11\x5e\x7b\xea\x8c\x30\x4a\x73\x21\x89\xb4\xd8\x34\x0a\xac\xd7\x58\xd6\xce\x5d\xc0\xcd\x38\xd5\x92\x1b\xa6\xa8\x23\x8d\xd3\x82\x37\x4a\x19\x26\x85\xc3\xc0\xb5\xa8\x85\x91\xc6\xeb\x53\x44\xf5\xb2\xf9\x8c\xcd\x81\xd7\x63\x57\xae\xad\x4d\x63\x2c\x77\xba\x87\x16\x2d\x7a\x1d\xe2\x3c\x2a\x26\x07\x7f\xac\x6f\x5b\xf4\x9c\xdb\xe5\xb2\x4b\x56\x77\xbb\x94\x4b\xdb\xd4\x5c\x54\x08\x6d\x75\x5e\x0d\xc1\xc2\xfc\x1b\x78\x79\xce\xf2\x6d\xe8\x43\x69\x11\x97\x84\x32\xa5\xaa\x73\xa7\x30\x95\x07\x1d\xca\x26\xf4\x90\xc6\xd2\xa2\x85\xc0\x38\xcf\x63\x7a\x3d\xad\xa1\x0c\x61\x6e\x9a\x9e\x22\x42\x2c\x30\x3c\xe9\x6e\x96\xe7\xf5\xe8\xe9\x5e\x99\x56\x29\x75\xd7\xd6\x42\xce\x37\x51\x9b\x0e\x5c\x8b\xca\x30\x42\x55\xfd\x1d\x00\x00\xff\xff\x36\xda\x4f\x39\x60\x06\x00\x00")

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

	info := bindataFileInfo{name: "resources/default_config.yaml", size: 1632, mode: os.FileMode(420), modTime: time.Unix(1530177056, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesTesting_configYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xcb\x6a\x2c\x3d\x0c\x84\xf7\x79\x8a\x46\x9b\x6c\x7a\x40\xb2\xe5\xf6\xe5\x6d\x64\x49\x26\xc3\x64\x2e\x74\x4f\x92\x3f\x84\xbc\xfb\x4f\x0f\xc3\x39\xdb\xb3\x93\xed\xe2\x73\x15\x2a\xbf\xbf\xf9\xea\x1f\xe7\xf6\x32\x4d\xa2\x7a\xfd\xb8\xdc\xb7\x7d\x9e\xa6\xb3\x1c\x2f\x6d\x7a\x8c\xd3\x74\xf2\xef\x36\xbd\xfe\x80\x98\xad\xbe\x6d\xd0\x80\x3b\xf5\xc2\x51\x46\xce\x52\x46\x1e\x3c\x50\x2c\x10\x27\xac\xa9\xc6\xec\xd1\x2a\x7a\xb4\x02\x33\xe8\xfa\x7d\xbb\x5f\xa1\xfd\x80\x1e\x6f\x6f\xbe\x42\x03\xf1\xed\x40\xa1\x1c\xf4\xbe\xee\x82\xc7\xf5\xdd\xff\xbb\xef\x4f\x99\x73\x5d\x54\x92\x77\x76\xe5\x65\x64\xd5\x31\xaa\x46\xcc\x98\x78\x09\x28\x8e\x25\x97\x40\xd6\xc9\x74\xc4\x6c\x91\x3b\x53\xac\x39\x0f\x0d\xc9\xfe\xf0\x6e\xb2\xca\x79\xdb\xbf\x3d\x7e\x42\x83\x51\x82\x8e\xa8\x22\xd1\x3a\x55\xa5\x54\x5c\x0a\xf3\x08\xc5\x33\x95\xdc\xe1\x77\x86\x93\x0d\x68\xb0\x3d\x0c\xc3\xe3\xf8\x17\x62\xa7\x77\xbf\x40\x8b\x61\x86\x0b\xb4\xb0\x04\x62\x9e\xe1\x06\x8d\x66\x58\xa1\x95\x19\x36\x79\xdf\x03\x10\x05\x4a\x85\xb4\xf7\xb0\x74\xd4\x84\x58\x33\x26\x37\x64\x4c\xb9\x18\x53\xad\xd4\xa9\x47\xaf\xd1\x58\x71\xa9\x38\x0a\x15\xe1\x52\x84\x22\xca\xd8\x8d\x9c\x45\xa1\x41\x8a\x0b\xc5\xc4\xa4\x38\x7c\x84\x8c\x69\x18\x95\xca\x3a\xac\x56\x43\x11\x5f\x50\x3a\x7a\x48\xb5\xba\xa8\x22\xd9\x92\xc4\xca\xa8\xb1\xe6\xa0\x94\x76\xd2\xd1\x76\x50\x42\xce\x18\xed\x90\x4b\x48\x07\x76\xec\x87\x52\xc5\x0e\x16\x58\x12\xeb\x12\x39\x0a\xcc\xf0\xe9\xeb\x76\xbc\xee\x21\x7f\x5f\x9f\x8b\xbf\xc9\xb6\x7d\x5d\x57\x6b\xd3\xf0\xcb\xfa\x35\x22\x5f\xd6\xa8\xf6\xbe\x9d\x4f\x4f\xc5\xb3\x12\x6d\xfa\xd7\x46\xbc\xbc\xfc\x1f\x00\x00\xff\xff\x28\xb6\x4d\x04\x75\x02\x00\x00")

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

	info := bindataFileInfo{name: "resources/testing_config.yaml", size: 629, mode: os.FileMode(420), modTime: time.Unix(1529575897, 0)}
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

