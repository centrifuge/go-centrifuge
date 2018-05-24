// Code generated by go-bindata. DO NOT EDIT.
// sources:
// resources/networks.yaml
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

var _resourcesNetworksYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xd0\xbf\x6e\xdb\x3a\x1c\xc5\xf1\x5d\x4f\x41\x68\xb7\x25\xfe\x35\xa9\x2d\xf7\x36\x2d\xda\xa6\x86\x9d\x1a\x71\xdd\x8d\x22\x7f\xb4\x18\xc5\x94\x42\x52\x89\x95\xa7\x2f\x60\x37\x43\xbb\x74\x3b\xc0\x01\x3e\xc3\x37\x40\x7e\x1d\x62\x9f\x1a\x54\x20\x64\x20\xe4\xe8\xdd\x74\x84\x45\x9c\x52\xf2\x3a\x74\xfe\xe9\x69\x01\xb9\x5b\x44\x1f\x7a\x68\xe7\xa6\x40\x08\x21\x6f\x1b\xc4\xf1\x65\xb6\xc3\x90\x53\x8e\x7a\xdc\x00\xc4\x74\xbd\x11\x5a\xa0\xb2\xf2\x23\xab\x30\x59\x2d\xeb\x65\xbd\xc4\x55\x36\x63\x45\x25\xa9\x49\xe5\x47\x97\xaa\xed\x69\xb7\x3d\xb7\xaf\xfd\xf4\xf3\x70\xf8\xe0\xa6\xb7\x5d\x7b\xbe\xbd\xb9\x87\xdd\xfa\xff\xbb\xe1\x6d\x9e\x39\x97\x2f\xdb\x70\x7c\x78\xd9\x7c\x7b\xbc\x3b\xf4\xe5\x3f\x59\xfa\xce\x3e\x38\x71\xbb\x16\xa7\xfe\x79\x0f\x8f\xfb\xaf\x7b\xf2\xbc\x99\xb0\xf8\x31\xda\x4f\xb4\xff\x32\xe0\x1d\x3d\x75\xba\xdb\xfc\xc7\xbf\x03\x0f\xb8\x2c\x2e\x2e\xe4\x0e\x22\x4c\xa7\xf5\x35\xc5\x67\xdb\x20\x76\x39\xcc\x10\x72\xd4\x26\xdf\x58\x1b\x21\x25\xb8\x46\xba\x06\x80\x90\x7d\x9e\x3f\x6a\x93\x87\x38\x37\xa8\xac\xcf\x35\x97\x0a\x2c\x93\x44\x3b\x69\x85\xac\x95\xab\x09\x71\x06\x63\xad\xa9\x52\xce\x48\x2d\x25\xb5\x9c\x94\x7f\x19\xf7\x70\xf4\x29\xff\x46\x9c\x62\x78\x65\x2c\x58\x4c\x69\xab\x25\x65\x96\x19\x67\x8d\x00\xc6\x89\x6d\x2d\x71\x86\x13\xb1\x62\xf0\x8e\xe8\x60\xba\x21\xfe\x41\xb4\x7c\x85\x1d\xb0\xda\x02\xb1\x02\x83\xe2\xcc\x3a\x65\x34\x03\x25\xb0\x13\x8a\x49\xe9\x30\xe5\xba\x2e\x8b\xa2\x28\x7e\x05\x00\x00\xff\xff\x6a\x3c\x33\xbb\xff\x01\x00\x00")

func resourcesNetworksYamlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesNetworksYaml,
		"resources/networks.yaml",
	)
}

func resourcesNetworksYaml() (*asset, error) {
	bytes, err := resourcesNetworksYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/networks.yaml", size: 511, mode: os.FileMode(420), modTime: time.Unix(1526994034, 0)}
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
	"resources/networks.yaml": resourcesNetworksYaml,
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
		"networks.yaml": &bintree{resourcesNetworksYaml, map[string]*bintree{}},
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

