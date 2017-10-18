// Code generated by go-bindata.
// sources:
// web/your-404-page.tmpl
// web/your-500-page.tmpl
// DO NOT EDIT!

package jsonfaas

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

var _webYour404PageTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x92\x41\x6f\x13\x31\x10\x85\xcf\xe4\x57\x4c\x7c\x6d\xbc\x26\x0a\xd0\x16\xd9\x2b\xb5\x34\xa8\xaa\x80\x50\x28\x42\x70\x33\xeb\xc9\x7a\x12\xaf\xbd\xb1\x27\x89\xc2\xaf\x47\xd9\x04\xb5\x27\xcf\xcc\x93\x9e\xbe\xf7\x64\x3d\xbe\x5b\x7c\x78\xfa\xf5\x75\x0e\x9e\xbb\x50\x8f\xf4\xf1\x81\x60\x63\x6b\x04\x46\x51\x8f\x00\xb4\x47\xeb\x8e\x03\x80\x1e\x4b\x09\xdf\x70\xb3\xa5\x8c\x0e\x3a\x64\x0b\x6c\xdb\x02\x52\x9e\xf5\xe1\xd4\x78\x9b\x0b\xb2\x11\x5b\x5e\xca\x2b\xf1\x52\x8a\xb6\x43\x23\x76\x84\xfb\x3e\x65\x16\xd0\xa4\xc8\x18\xd9\x88\x3d\x39\xf6\xc6\xe1\x8e\x1a\x94\xc3\x32\x01\x8a\xc4\x64\x83\x2c\x8d\x0d\x68\xa6\x13\x28\x3e\x53\x5c\x4b\x4e\x72\x49\x6c\x62\x3a\x5b\xbf\xd2\x4c\x1c\xb0\xbe\x0d\x36\xae\xa1\xb7\x2d\x6a\x75\xba\x1c\xe9\xd5\x7f\x7c\xfd\x27\xb9\xc3\x19\xc6\x4f\xeb\x7b\x0c\x21\x4d\x60\x9f\x72\x70\x63\xad\xfc\xb4\x1e\x3d\x67\x5c\x3d\x6e\x31\x1f\x60\x49\xb9\xf0\x04\xd8\x63\x84\x27\x64\x8f\xf9\xbc\xdc\xa6\xc4\x85\xb3\xed\xe1\xe1\x7b\xf5\x1c\xbf\x34\x99\x7a\x86\x92\x1b\x23\x3c\x73\x5f\xde\x2b\xd5\x24\x87\xd5\x6a\x73\xf4\xab\x9a\xd4\xa9\xd3\x28\x67\xd5\xb4\x9a\x56\x25\x50\x57\x75\x14\xab\x55\x11\x40\x91\xb1\xcd\xc4\x07\x23\x8a\xb7\xb3\xab\x37\xf2\xe6\xf2\xe3\xef\xd5\xe5\xee\xc2\xa9\xe2\xba\xcf\x9b\x5e\xc5\xc5\xe3\x3e\xd0\xa7\xdd\x8f\xf2\xb0\xbc\xbb\xff\x79\xb1\xbe\x5e\x74\xad\xb2\x6a\xee\xf1\xc6\xb5\xfc\xf7\x4b\x99\xf9\x7e\x69\xdb\x77\x73\x77\xfd\xf6\x75\x14\xd0\xe4\x54\x4a\xca\xd4\x52\x34\xc2\xc6\x14\x0f\x5d\xda\x16\x51\x6b\x75\x62\x1d\xc0\x87\x9a\x4e\xed\x68\x35\xfc\x83\x7f\x01\x00\x00\xff\xff\xb5\x60\xf2\x4b\x17\x02\x00\x00")

func webYour404PageTmplBytes() ([]byte, error) {
	return bindataRead(
		_webYour404PageTmpl,
		"web/your-404-page.tmpl",
	)
}

func webYour404PageTmpl() (*asset, error) {
	bytes, err := webYour404PageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/your-404-page.tmpl", size: 535, mode: os.FileMode(493), modTime: time.Unix(1508362567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webYour500PageTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x92\x41\x6f\x13\x31\x10\x85\xcf\xe4\x57\x4c\x7c\x6d\xbc\x26\x0a\xd0\x16\xd9\x2b\xb5\x34\xa8\xaa\x80\x50\x28\x42\x70\x33\xeb\xc9\x7a\x12\xaf\xbd\xb1\x27\x89\xc2\xaf\x47\xd9\x04\xb5\x27\xcf\xcc\x93\x9e\xbe\xf7\x64\x3d\xbe\x5b\x7c\x78\xfa\xf5\x75\x0e\x9e\xbb\x50\x8f\xf4\xf1\x81\x60\x63\x6b\x04\x46\x51\x8f\x00\xb4\x47\xeb\x8e\x03\x80\x1e\x4b\x09\xdf\x70\xb3\xa5\x8c\x0e\x3a\x64\x0b\x6c\xdb\x02\x52\x9e\xf5\xe1\xd4\x78\x9b\x0b\xb2\x11\x5b\x5e\xca\x2b\xf1\x52\x8a\xb6\x43\x23\x76\x84\xfb\x3e\x65\x16\xd0\xa4\xc8\x18\xd9\x88\x3d\x39\xf6\xc6\xe1\x8e\x1a\x94\xc3\x32\x01\x8a\xc4\x64\x83\x2c\x8d\x0d\x68\xa6\x13\x28\x3e\x53\x5c\x4b\x4e\x72\x49\x6c\x62\x3a\x5b\xbf\xd2\x4c\x1c\xb0\xbe\x0d\x36\xae\xa1\xb7\x2d\x6a\x75\xba\x1c\xe9\xd5\x7f\x7c\xfd\x27\xb9\xc3\x19\xc6\x4f\xeb\x7b\x0c\x21\x4d\x60\x9f\x72\x70\x63\xad\xfc\xb4\x1e\x3d\x67\x5c\x3d\x6e\x31\x1f\x60\x49\xb9\xf0\x04\xd8\x63\x84\x27\x64\x8f\xf9\xbc\xdc\xa6\xc4\x85\xb3\xed\xe1\xe1\x7b\xf5\x1c\xbf\x34\x99\x7a\x86\x92\x1b\x23\x3c\x73\x5f\xde\x2b\xd5\x24\x87\xd5\x6a\x73\xf4\xab\x9a\xd4\xa9\xd3\x28\x67\xd5\xb4\x9a\x56\x25\x50\x57\x75\x14\xab\x55\x11\x40\x91\xb1\xcd\xc4\x07\x23\x8a\xb7\xb3\xab\x37\xf2\xe6\xf2\xe3\xef\xd5\xe5\xee\xc2\xa9\xe2\xba\xcf\x9b\x5e\xc5\xc5\xe3\x3e\xd0\xa7\xdd\x8f\xf2\xb0\xbc\xbb\xff\x79\xb1\xbe\x5e\x74\xad\xb2\x6a\xee\xf1\xc6\xb5\xfc\xf7\x4b\x99\xf9\x7e\x69\xdb\x77\x73\x77\xfd\xf6\x75\x14\xd0\xe4\x54\x4a\xca\xd4\x52\x34\xc2\xc6\x14\x0f\x5d\xda\x16\x51\x6b\x75\x62\x1d\xc0\x87\x9a\x4e\xed\x68\x35\xfc\x83\x7f\x01\x00\x00\xff\xff\xb5\x60\xf2\x4b\x17\x02\x00\x00")

func webYour500PageTmplBytes() ([]byte, error) {
	return bindataRead(
		_webYour500PageTmpl,
		"web/your-500-page.tmpl",
	)
}

func webYour500PageTmpl() (*asset, error) {
	bytes, err := webYour500PageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/your-500-page.tmpl", size: 535, mode: os.FileMode(493), modTime: time.Unix(1508362567, 0)}
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
	"web/your-404-page.tmpl": webYour404PageTmpl,
	"web/your-500-page.tmpl": webYour500PageTmpl,
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
	"web": &bintree{nil, map[string]*bintree{
		"your-404-page.tmpl": &bintree{webYour404PageTmpl, map[string]*bintree{}},
		"your-500-page.tmpl": &bintree{webYour500PageTmpl, map[string]*bintree{}},
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

