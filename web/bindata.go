// Code generated by go-bindata.
// sources:
// static/main.css
// templates/_foot.html
// templates/_head.html
// templates/index.html
// templates/redirect.html
// DO NOT EDIT!

package web

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

var _staticMainCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xca\x4f\xa9\x54\xa8\xe6\x52\x50\x48\xcb\xcf\x2b\xd1\x4d\x4b\xcc\xcd\xcc\xa9\xb4\x52\x50\xf2\x48\xcd\x29\x4b\x2d\xc9\x4c\x4e\x54\xf0\x4b\x2d\x4d\x55\xd2\x51\x80\x0b\xe8\x28\x28\x05\xa7\xa6\xe7\xa7\x2a\x84\x7a\x02\x85\x1d\x8b\x32\x13\x73\x74\x14\xd2\x8a\x52\x53\x8b\x13\xf3\x8a\x75\x14\x40\xa4\x6e\x71\x6a\x51\x66\x1a\x50\xa1\x63\x41\x41\x4e\xaa\x82\x73\x7e\x4e\x7e\x91\x82\x6b\x6e\x7e\x56\xa6\x12\x92\x6e\x2c\x22\xc1\x95\xb9\x49\xf9\x39\x4a\xd6\x5c\xb5\x5c\x80\x00\x00\x00\xff\xff\x3e\xe1\x83\x30\x98\x00\x00\x00")

func staticMainCssBytes() ([]byte, error) {
	return bindataRead(
		_staticMainCss,
		"static/main.css",
	)
}

func staticMainCss() (*asset, error) {
	bytes, err := staticMainCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/main.css", size: 152, mode: os.FileMode(420), modTime: time.Unix(1468518898, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_footHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x50\xb0\xd1\x4f\xca\x4f\xa9\xb4\xe3\xb2\xd1\xcf\x28\xc9\xcd\xb1\xe3\x02\x04\x00\x00\xff\xff\xd2\x42\x65\xbd\x12\x00\x00\x00")

func templates_footHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templates_footHtml,
		"templates/_foot.html",
	)
}

func templates_footHtml() (*asset, error) {
	bytes, err := templates_footHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/_foot.html", size: 18, mode: os.FileMode(420), modTime: time.Unix(1468764419, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_headHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\xc1\x6a\xf3\x30\x10\x84\xef\x79\x8a\xfd\x75\xfe\x15\xd3\x5b\x0f\x56\x20\x24\x39\x14\x0a\x2d\x6d\x03\xed\x51\x91\x27\xd5\x12\x59\x72\xa5\x75\x42\xfa\xf4\x75\x62\x02\x2e\x3d\x69\x06\xe9\xd3\xcc\x6e\xfd\x6f\xfd\xb4\x7a\xfb\x78\xde\x90\x97\x36\x2c\x66\xf5\x78\x10\xd5\x1e\xb6\xb9\x88\x41\xb6\x10\x4b\xce\xdb\x5c\x20\x46\xf5\xb2\xd7\xf7\x6a\x7a\xe5\x45\x3a\x8d\xaf\x9e\x8f\x46\xbd\xeb\xed\x52\xaf\x52\xdb\x59\xe1\x5d\x80\x22\x97\xa2\x20\x0e\xdc\xc3\xc6\xa0\xf9\xc4\x2f\x32\xda\x16\x46\x1d\x19\xa7\x2e\x65\x99\x3c\x3e\x71\x23\xde\x34\x38\xb2\x83\xbe\x9a\xff\xc4\x91\x85\x6d\xd0\xc5\xd9\x00\x73\x77\xfb\x48\x58\x02\x16\xe9\xbb\x9f\x73\xaa\xab\xd1\xfd\x89\x68\x50\x5c\xe6\x4e\x38\xc5\x49\xca\x08\x91\xa6\x25\xad\x73\x1f\x0f\x88\xb4\x7d\x79\xa4\x57\x3f\x94\x41\x44\xbe\x45\x04\x8e\x07\xca\x08\x46\x15\x39\x07\x14\x0f\x0c\x65\x7d\xc6\xde\xa8\xaa\xc8\x30\xaa\xab\x5a\xcb\x71\xee\x4a\xb9\x32\x75\x35\xae\xef\x22\x77\xa9\x39\x2f\x66\x3f\x01\x00\x00\xff\xff\x55\xce\x24\xf7\x69\x01\x00\x00")

func templates_headHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templates_headHtml,
		"templates/_head.html",
	)
}

func templates_headHtml() (*asset, error) {
	bytes, err := templates_headHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/_head.html", size: 361, mode: os.FileMode(420), modTime: time.Unix(1468764427, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\x8a\xcf\x48\x4d\x4c\xd1\xcb\x28\xc9\xcd\x51\xaa\xad\xe5\x52\x00\x82\xf0\xd4\x9c\xe4\xfc\xdc\x54\x85\x92\x7c\x85\xfc\xaa\x52\xbd\xcc\x7c\x1d\x85\x44\x85\x94\xa2\xd2\xbc\xec\xd4\x3c\x85\xd0\x20\x1f\x85\xe2\x8c\xfc\xa2\x92\xd4\xbc\xd4\x22\x3d\x2e\x14\x93\xd2\xf2\xf3\x4b\xe0\x26\x01\x02\x00\x00\xff\xff\x93\x8b\x7b\xb5\x64\x00\x00\x00")

func templatesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesIndexHtml,
		"templates/index.html",
	)
}

func templatesIndexHtml() (*asset, error) {
	bytes, err := templatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/index.html", size: 100, mode: os.FileMode(420), modTime: time.Unix(1468764414, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRedirectHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x4f\x41\xcb\x82\x50\x10\xbc\xfb\x2b\x96\xf7\xdd\xf5\xeb\xbe\x3e\xd0\x7c\xf6\x02\xc5\x90\x3d\xd4\x51\x72\xc9\x43\x2a\x98\x05\x21\xfe\xf7\xde\xdb\xea\xd6\x9c\x66\x67\x86\x61\x16\x2d\x95\x85\x0e\x00\xd0\x9a\x24\xf3\xc4\x51\xda\x53\x61\x74\x39\x3e\xb8\x85\x03\x4f\x7d\x33\xf0\x30\x5f\x9f\x18\xbd\x0d\x9f\x8e\xbe\x71\x4c\xab\xec\x04\xe9\x6e\x5b\x15\x55\x1d\xab\xbf\x5c\xa0\x80\xcc\x91\xdc\xf9\x2f\x50\x9f\x62\xbb\xf9\xd5\xea\x54\xb1\xa9\x63\x68\xc7\xf3\xbd\x77\x3a\x74\xcd\x0d\x7a\xc9\x62\x02\xb6\x36\x79\xac\x96\x25\xa4\x66\xba\xf0\xbc\xae\x4a\x77\x3c\x31\x46\x89\x0e\x65\x8d\x1f\xa1\x03\x57\x25\xdf\xbc\x02\x00\x00\xff\xff\xf9\x4e\x03\x95\xd5\x00\x00\x00")

func templatesRedirectHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesRedirectHtml,
		"templates/redirect.html",
	)
}

func templatesRedirectHtml() (*asset, error) {
	bytes, err := templatesRedirectHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/redirect.html", size: 213, mode: os.FileMode(420), modTime: time.Unix(1468764854, 0)}
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
	"static/main.css": staticMainCss,
	"templates/_foot.html": templates_footHtml,
	"templates/_head.html": templates_headHtml,
	"templates/index.html": templatesIndexHtml,
	"templates/redirect.html": templatesRedirectHtml,
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
	"static": &bintree{nil, map[string]*bintree{
		"main.css": &bintree{staticMainCss, map[string]*bintree{}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"_foot.html": &bintree{templates_footHtml, map[string]*bintree{}},
		"_head.html": &bintree{templates_headHtml, map[string]*bintree{}},
		"index.html": &bintree{templatesIndexHtml, map[string]*bintree{}},
		"redirect.html": &bintree{templatesRedirectHtml, map[string]*bintree{}},
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

