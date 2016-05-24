// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package assets contains template strings and other assets for the allocatord web interface.
//
// This package is auto-generated by "jiri go generate v.io/x/ref/services/allocator/allocatord"
// which in-turn uses https://github.com/jteeuwen/go-bindata/
// Code generated by go-bindata.
// sources:
// bad-request.tmpl.html
// dash.js
// dashboard.tmpl.html
// error.tmpl.html
// head.tmpl.html
// header.tmpl.html
// home.tmpl.html
// root.tmpl.html
// style.css
// DO NOT EDIT!

package assets

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

var _badRequestTmplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8e\x41\x0a\x02\x31\x0c\x00\xef\xfb\x8a\xf8\x00\x5b\xbc\x4a\xe8\x41\xf0\x2e\xfe\xa0\x9a\x40\x17\x56\x5b\x6b\x44\x96\xd2\xbf\xdb\x36\x7a\xf0\x14\x3a\x9d\x21\xc1\x0d\xc5\xab\xac\x89\x21\xc8\x6d\x71\x13\xfe\x06\x7b\x72\x13\x00\xca\x2c\x0b\xbb\x83\x27\x38\xf3\xe3\xc5\x4f\x41\xab\x68\x42\xab\x12\x5e\x22\xad\xc3\x0d\xbb\x7f\xb1\xbd\x3b\x4e\x99\xfb\x04\x28\xc5\x9c\xbc\x84\x5a\x3b\xb5\x5f\x5c\xca\x7b\x96\x00\xe6\x98\x73\xcc\xb0\xd5\xcf\xa4\xc1\x60\xfb\xde\xd5\x6a\x34\xd2\x84\xef\xd4\x44\xb4\xba\xba\x6d\x1a\x57\x7f\x02\x00\x00\xff\xff\xe0\x6f\x86\xe4\xcd\x00\x00\x00")

func badRequestTmplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_badRequestTmplHtml,
		"bad-request.tmpl.html",
	)
}

func badRequestTmplHtml() (*asset, error) {
	bytes, err := badRequestTmplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bad-request.tmpl.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dashJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x57\xdd\x53\xdb\x38\x10\x7f\xe7\xaf\xd8\xb6\xdc\xd8\xb9\x0b\x26\xa1\x0c\xed\x84\xe9\x74\x20\x70\x25\x73\xe5\xa3\x24\xf4\x1e\x18\xa6\x23\x2c\x25\x56\x2b\x4b\xae\x24\x03\xe9\x4d\xfe\xf7\x5b\xf9\xdb\x89\xe1\xe6\x78\x30\xf1\xee\xfe\x76\x57\xfb\xa5\xf5\xee\x2e\x8c\x55\xb2\xd4\x7c\x11\x59\xd8\x1b\x0c\x0f\x60\x16\x31\xf8\x4a\x24\xa1\x3c\x8d\xe1\x28\xb5\x91\xd2\x26\x80\x23\x21\x20\x13\x32\xa0\x99\x61\xfa\x81\xd1\x60\x0b\xc1\x37\x86\x81\x9a\x83\x8d\xb8\x01\xa3\x52\x1d\x32\x08\x15\x65\x80\xaf\x0b\xf5\xc0\xb4\x64\x14\xee\x97\x40\xe0\x78\x7a\xb2\x63\xec\x52\x30\x87\x12\x3c\x64\x12\x91\x36\x22\x16\x42\x22\xe1\x9e\xc1\x5c\xa5\x92\x02\x97\x48\x64\xf0\x79\x32\x3e\xbd\x98\x9e\xc2\x9c\x0b\x16\x6c\x6d\x3d\x10\x0d\x94\x98\x08\x3e\xc0\x3c\x95\xa1\xe5\x4a\xfa\x3d\xf8\x67\x0b\x00\x95\x9d\x10\x4b\x20\x35\x68\xc8\x2a\x58\x30\xc9\x34\xb1\xe8\x45\x44\xb4\x35\x01\x8a\x38\xf0\xf8\xec\xe8\x7a\x36\x45\xf8\x2d\x12\x20\x43\xba\x3f\xcb\xad\x60\x23\xf0\x3e\x23\x42\x86\x4b\xf0\x63\xd3\xf3\xfa\x05\x93\xa2\xde\xbf\xd8\xb2\x66\x57\x1c\x4e\x91\x28\x0a\x62\x46\x5b\xf5\x3b\xf5\x7e\xb9\x9a\x82\xaf\xd9\xcf\xdd\x4e\xb5\x5f\x12\xd3\x56\xf9\x13\x09\x2f\xa9\x1b\x5f\xdd\x60\xc0\xc9\x82\x81\xff\x5b\x97\xc2\xe9\xd2\xa0\x48\x26\x71\x15\xda\xb6\xee\x30\x49\x77\x52\xc7\xd9\x49\x90\xf5\x92\x95\x73\x16\x2b\xbd\x2c\x0d\x7d\x3a\x7e\xc6\x12\x8a\x65\x22\xc7\x4b\xcb\xd6\xce\x11\xb3\xb8\xb0\x75\xdf\x62\x9a\x90\x38\x0b\xc3\xdd\xe1\x60\x6f\xbf\x7e\xbc\xe4\xcc\x09\x37\x3f\xfe\xd3\x15\x27\xf4\x9c\x2f\x14\x79\xff\xdf\x19\x7c\xde\x1d\x6e\x15\xc5\x73\x72\x73\x7d\x34\x9b\x5c\x5e\x4c\xbf\xcd\x2e\xbf\x4d\x4f\xc7\x97\x17\x27\xae\x94\x72\x57\xbd\x61\xe4\x8d\xe0\xed\xc1\x60\x90\x2b\xf6\xf6\xca\x77\xf8\x1d\xf6\x0a\xda\x7e\x83\xb6\x5f\xd0\x0e\x1a\xb4\x83\x82\x36\x6c\x82\x87\x25\x7a\x48\x1b\x1a\x4b\xf8\xbb\x16\x11\x1f\xef\x90\xbe\xaa\x7c\xce\x0a\xfe\xdb\xdf\x93\x93\xd9\x19\xba\xfa\x6e\x30\x38\x6c\x31\xce\x4e\x27\x9f\xce\x66\xc8\xd9\x7b\xbb\xc6\x19\x5f\x7e\xbe\xbc\x46\x86\xf7\x66\x30\x78\xff\xf6\xfd\x9f\x5e\xa5\x92\xa6\xd8\x58\xd8\x79\x13\x39\x65\xa1\x92\xd4\xa0\x94\x73\xa0\xc4\xc7\xd8\xc0\x96\xd1\x0b\x12\x33\x87\xf7\x4a\x7a\x68\xf4\xbc\x20\x20\xa5\xec\x5f\x60\x4f\x56\x93\xd0\x5e\x11\x8d\x00\xcb\xb4\x29\x3a\x3a\x07\x6d\xa7\x5a\x20\x68\x3b\xc0\xff\x7e\xef\x30\xa3\xb7\x0d\x38\x89\x20\x71\x68\xdf\x93\x5e\x21\x52\xd8\x6a\xf2\x1c\xa9\x64\x6b\x66\x53\x2d\x5b\x8a\x5e\x7d\xf8\x00\x38\x78\xd8\x9c\xe3\xa8\x72\x52\xab\xad\x7c\xac\x4c\x19\xce\xba\x34\x81\x88\x48\x2a\xd0\x3d\x1c\x50\xda\x4d\x15\xb9\xe0\x72\x51\x05\x23\x68\x1e\xc9\xa0\xfa\xe4\xa4\xe0\x8c\x95\xb4\x5a\x09\x84\x56\x07\xdb\xf6\xbd\x37\x25\xd0\x78\xbd\x20\x8c\xb8\xa0\x9a\xe1\x2c\x0b\x42\x1c\x89\x3f\xfc\xb5\xe1\xe6\xfe\xd0\x13\x0c\x11\xce\xca\x12\x08\x82\xdc\x33\x01\xe8\x95\xb3\xb7\x99\x95\xa0\xec\x90\x8e\x74\x75\x55\xf2\xed\xb6\xef\xa6\x77\x2f\xb0\x98\x11\xbf\x77\x77\x58\xe0\xd3\x04\x7b\x8c\xb9\xd8\xd7\x9e\xdc\x64\x34\xb8\x99\x94\x46\xf0\x44\x41\x69\x68\x87\x5b\x16\x07\x86\x09\x16\x62\x70\xf1\x78\x56\x2d\x16\x82\x8d\x05\x31\xc6\xf7\x6a\xfa\x61\x85\x2d\xec\xbe\x24\xb6\xea\x35\x73\xf2\xc9\xe5\xc4\xb5\x7e\x76\xfc\xdc\x43\xd3\x18\xf5\x55\x26\x4a\xe7\x1b\x81\x17\x0a\xaf\x34\xb9\xd8\xc9\xc2\x87\xde\x99\x48\x3d\x96\x95\xe5\x2a\x2e\x2b\x17\x53\xf5\x35\x80\x1c\x35\x0b\xa5\x1a\x3c\xa3\xcd\xc8\x96\x3c\x57\x6a\xa3\xec\x99\xfb\x7e\x58\x19\x67\x5a\x2b\xbd\x13\x9b\x05\x1a\x8e\x38\x65\xa5\xe1\xed\x80\x7c\x27\x4f\x78\x6c\x4b\xac\xf9\xe8\xc1\x1f\x48\xc9\xeb\x36\x77\xa7\xd7\x2b\x54\x03\x04\x54\x49\x56\x57\x88\x8b\x42\x5d\x25\x75\xc2\xc6\x59\x30\x72\xf6\x61\xc5\x5d\x35\xf4\xcc\x09\x17\x8d\x4a\x6b\xaa\x58\x77\xb5\x19\xa3\x75\x35\xa1\x8a\x13\x81\x8d\xdb\x55\xb4\xcf\xc5\xbc\x79\xf4\xcd\xec\xde\x14\xf9\x24\xb8\x6b\x3c\x9b\xd3\xe6\xf9\x0a\x7b\xf9\xe5\x1e\x60\x7f\x9e\x92\x30\xaa\xfd\xc9\x74\xd4\x4e\x65\xb3\x08\xf3\x2b\xd9\x63\xc3\xcd\x85\x52\x58\x7d\xc1\x03\x37\x29\x11\xfc\x57\xde\xd4\x47\x9a\x91\xcc\x90\x4f\x55\x98\xc6\x4c\xda\x60\xc1\xec\xa9\x60\xee\xe7\xf1\x72\x42\x73\xdd\x01\xa7\xbd\xea\x34\x4e\xbd\x4a\xb2\xd6\x6e\x14\x51\x75\x93\xe5\x80\xec\xa5\xdf\xe6\xcd\xb0\xef\xa6\x6e\x33\x1a\xb5\xe2\x37\xc7\xf9\x31\xe5\xbf\xdc\x15\xb5\x5f\x23\x56\xf5\xcf\x47\x4e\x6d\x34\x6a\x4e\xfa\x9a\x17\x31\xb7\xab\x8d\x5a\xd3\xbe\xe6\x0a\x86\x6b\x12\x6d\x9b\x4b\x94\xe1\xce\x79\xbc\x32\x25\x16\x9a\xd7\x65\x11\xf7\x3e\xce\x4c\x1b\x37\x68\xbf\x62\x13\x28\xa1\xf4\xa8\x79\x9b\x34\xd8\xab\x2e\xb5\xd1\xd1\x13\x5f\xd3\x1a\x73\xf9\x95\x88\x14\x0f\x8f\xd9\x72\x1b\x1e\xcb\x32\x1e\x9c\x73\x39\xe3\x38\xb5\xf1\x76\x1c\x0c\x06\xbd\x7e\x13\x42\x9e\xba\x21\xe4\xe9\x39\xc8\x42\x73\x2a\x70\xee\x9b\xcd\x33\x60\xdf\x8f\x60\x67\xd8\x6f\x91\x53\xc9\xed\x86\xac\xdb\x45\x96\x8e\x8a\xf5\x17\x13\x44\xdd\x7a\xe7\xe7\xe7\x40\xa9\x77\xb7\xea\xaf\x49\x46\xb8\x21\xb7\x45\xa3\x51\x1c\x03\xf1\xfa\xe0\x45\xf8\x6f\x13\x11\x73\x21\xb8\xc9\xc7\x4c\x17\xf0\x6e\xd5\x02\xb4\xf0\x8d\x60\x17\xbf\x56\x65\xb1\x86\x01\xd5\xe4\xd1\xc7\x32\x70\xdb\xf3\x8c\xdc\x8b\x3c\x5a\xb7\x79\x99\x16\xdb\xd5\x5d\xbf\x28\xdb\x6c\x57\xea\xf5\xcb\xfa\x7e\x66\x36\xe7\x9b\xb7\x81\x4a\x25\xa8\xfb\xef\x38\xcd\x61\xae\x55\x9c\xad\xf5\x0b\xfe\xc0\x24\x96\x1a\x97\x6b\xbd\xdd\x72\x24\xe7\xf7\xf3\x0d\xad\xb9\x14\x50\x0b\x79\xff\x76\x77\x6d\xad\xa1\x70\x90\xda\x80\x50\x3a\x56\x22\x8d\xa5\xef\xb9\xe9\x61\xb1\x12\x5c\xb0\xbd\x4e\x09\x99\xc6\xf7\x4c\x37\xf9\x7c\x0e\x85\x3b\xf5\x1c\xc9\x31\xd7\xea\xd1\x14\xac\x20\x26\x49\x3d\x75\x12\xdb\x9c\x83\xc5\xc6\x71\xdb\xca\x52\x55\xa0\x89\x0d\x5c\x6d\xe2\xf4\x8f\x93\xae\x02\x2d\xb6\x54\xf8\x08\x28\x99\x55\x37\x4a\xe5\xa4\x51\x45\xaa\xe4\xab\xab\x7b\x55\x4e\xa5\x15\x30\x81\xab\xc3\x86\xe7\xfe\x6d\xe5\x02\x66\x75\x70\x57\x8a\x37\x97\x24\x6a\xab\xec\x56\x69\xe2\xd8\x00\xd5\x98\x77\xb1\x79\xd5\xb1\xc5\xd5\xc7\x47\x3f\x71\x84\x7a\xaf\xe5\x6b\xf7\x3d\x88\xdf\x41\x29\xd7\xf8\x99\xe6\x16\x29\x2e\xf1\xd0\x12\xbf\x16\x8b\x2b\x16\x24\xe2\xeb\xdd\x20\xf7\xa1\xe1\xd4\x76\xd7\x35\xf3\xec\xbe\xb5\xb9\xc3\xe4\xef\xf5\x0a\xc3\xf0\xdb\x74\xe9\xc6\x4c\x6a\x59\x50\xab\x9b\xa0\x2f\xfa\x81\x08\x3f\x47\xf6\x01\xf7\x5b\xcc\xc9\x46\xbd\x17\x31\x2a\x02\x81\x51\x19\x65\xcf\x7c\x05\x5f\x39\x83\xff\x06\x00\x00\xff\xff\x35\xc1\x74\xa9\x59\x0f\x00\x00")

func dashJsBytes() ([]byte, error) {
	return bindataRead(
		_dashJs,
		"dash.js",
	)
}

func dashJs() (*asset, error) {
	bytes, err := dashJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dash.js", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dashboardTmplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x55\x5d\x4f\xdb\x30\x14\x7d\xe7\x57\x78\x79\x09\x13\x8d\xdd\x02\xdb\xa4\x2e\xed\x34\x01\x9b\x90\x10\xa0\xb1\x97\x69\xda\x83\x6b\xdf\x36\x6e\x9d\x38\xd8\x37\x85\x88\xf1\xdf\xe7\x7c\xb4\xb4\x50\x2a\x2a\x4d\x7d\xa8\x7d\x7d\xee\xb9\x1f\xc7\xd7\x89\xdf\x9d\x5e\x9d\xfc\xfc\x75\x7d\x46\x12\x4c\xf5\x70\x2f\x6e\xfe\x08\x89\x13\xe0\xb2\x5a\xf8\x25\x2a\xd4\x30\x7c\x78\xa0\x37\x60\xe7\x60\x2f\x79\x0a\x8f\x8f\xe4\x94\xbb\x64\x64\xb8\x95\x31\x6b\x00\x7b\x0d\x5a\xab\x6c\x46\x12\x0b\xe3\x41\xc8\xd8\xd8\x64\xe8\xe8\xc4\x98\x89\x06\x9e\x2b\x47\x85\x49\x99\x70\xee\xcb\x98\xa7\x4a\x97\x83\x1b\x53\x58\x01\x07\x27\x46\xc2\xc1\xb5\x35\xfd\xe3\x6e\xb7\xf3\xa1\xdb\xfd\xfb\xc3\x8c\x0c\x9a\xbe\x5f\x76\xbc\x49\x21\xd7\x4a\x74\x8e\x9a\xc3\xa7\x5d\xbb\xf2\x88\xb0\x8e\x4d\x88\x05\x3d\x08\x1d\x96\x1a\x5c\x02\x80\x0b\x33\x96\x39\x0c\x42\x84\x7b\xac\x82\x87\x6d\x5d\x4e\x58\x95\x23\x71\x56\x0c\x82\x04\x31\x77\x7d\xc6\xf8\x94\xdf\x3f\xcf\xb7\xb2\x31\xad\x46\x8e\x4d\x6f\x0b\xb0\x25\x3b\xa4\xfe\xd7\x6e\x68\xaa\x32\x3a\x75\xc1\x30\x66\x0d\xdf\x16\x72\x21\xb3\xa9\xa7\xd4\xa6\x90\x63\xcd\x2d\x6c\x64\x8f\x0a\xab\xa3\x9c\x5b\x07\xd6\x07\x3a\xa2\x3d\x96\x7b\xcb\xff\x0f\x23\x8c\x99\x29\x60\x3d\x7a\xec\x43\xb4\xb5\x34\xb6\xd7\x62\x3d\x29\x1b\x38\xe4\xa8\x04\xab\x3b\x4d\x7d\x4b\x83\xba\xf3\xc1\x53\xe7\x83\x0d\xf9\xb5\x4e\xd2\x5f\x9c\x2d\xa5\xd4\x5a\x05\xb5\x56\x53\x3e\xe7\x8d\x35\x58\xaf\xf0\xee\xee\x8e\x4e\x1a\xba\xe6\x46\x25\xdc\xa2\x63\xda\x70\x09\x76\x67\xee\x61\x7b\x4b\x1a\xd9\x69\x43\x46\x2b\xb2\xfd\x50\x14\xd6\x42\x86\x61\x87\x3c\xe4\x5c\xcc\xf8\x04\x5c\x9f\xfc\x0e\x85\xb1\x50\xe3\xc2\x3f\x8f\xef\x3f\x6f\xf4\x77\x80\x57\xd9\x85\x27\x39\xe1\x5a\x8f\xbc\xef\x7e\x5d\xb8\xca\x14\xb6\x1e\xab\x39\xc6\x6c\x31\x70\xf1\xc8\xc8\xb2\x4d\x5b\xaa\x39\x51\x72\x10\x08\x3f\x46\x5c\x65\x60\x97\xb9\x2e\x8f\x2a\xb7\x15\xfb\xca\x49\x3d\x95\xc1\xb6\xb9\xf5\xc8\x75\xbf\x0a\x7c\x9e\xf9\xc6\x66\xc2\x43\x37\x00\x6a\x62\x48\xb9\xd2\x35\xf1\x59\xb5\x7a\x06\x5c\xdf\x2c\x2b\xa8\x9b\xf2\x3c\x4d\xa1\xb9\x73\xed\x61\x50\xe3\x34\x47\xc8\x44\x59\x09\xf8\x32\xf8\x4b\xf8\x6d\xee\xde\x0a\x15\x79\x11\x15\xce\x0b\x18\xe5\x02\xdf\xea\x94\x42\xda\x3a\x8d\x4a\x84\x37\xc7\x92\xca\xcd\xb6\xf8\xad\x6c\x56\x97\x8b\x5e\xc9\xc2\xfa\x9b\x6d\x32\x17\x6d\xd1\xbd\xba\x9e\x2a\x9b\x44\x9a\x8f\xc0\x8b\x71\x71\xf5\xf5\xf4\xfc\xf2\x3b\xa5\x74\x73\xff\x97\x9c\xaf\x48\xb0\x38\x8f\x14\x42\x4a\x1c\x68\x10\x08\x32\x18\xf6\x92\x2d\x15\xaf\x39\x05\xc3\xc3\x1d\xb0\xc7\x3b\x60\x3f\xee\x80\xed\xed\x92\x44\x6f\xd3\x0c\xbc\x82\xfd\x24\x77\x91\x10\xac\x35\x36\x4a\xdd\x64\xd9\xee\x6f\x7e\x52\x40\x12\x34\xfe\xa5\x44\xab\x60\x0e\x44\x72\xe4\x94\x5c\xfb\x2f\x8d\x03\x82\xb6\x24\x7c\xe2\xc5\x26\xd5\x08\x58\xba\x4e\x1d\xb3\xe6\x55\xf0\xcf\x44\xfd\x81\xfe\x17\x00\x00\xff\xff\x06\x5c\x3d\xaa\xb8\x07\x00\x00")

func dashboardTmplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_dashboardTmplHtml,
		"dashboard.tmpl.html",
	)
}

func dashboardTmplHtml() (*asset, error) {
	bytes, err := dashboardTmplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dashboard.tmpl.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _errorTmplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8f\x41\x0a\xc2\x30\x10\x45\xf7\x3d\xc5\xd8\xbd\x0d\xdd\xca\x18\x70\x21\xb8\x10\x04\xc1\x03\xd4\x66\x24\x85\xb6\x29\x43\x44\x4a\xc8\xdd\x4d\x32\x76\xe1\xea\x93\xff\xdf\x83\x09\xee\x8c\xeb\xfd\xba\x10\x58\x3f\x8d\xba\xc2\x2d\xa8\x33\xba\x02\x40\x3f\xf8\x91\xf4\x99\xd9\x31\x2a\x79\x54\xa8\x64\xc6\xa7\x33\x6b\xa1\x6c\xab\x4f\x33\x14\x0a\x6e\x7d\xff\x66\x26\x93\xa8\xb6\x8c\x0b\x53\x4e\x80\x10\x9a\xc7\xfd\x1a\x63\x2e\xd5\xaf\x0d\xe1\x33\x78\x0b\x8d\xb8\x7b\x19\x17\xe1\x4b\x77\xc8\x5a\x8c\x8d\x48\xa2\xd0\x6c\xfe\x40\xec\xc0\x32\xbd\x8e\x75\x42\x2f\x6e\xa2\x18\x6b\x9d\x13\x55\xa7\x37\x0f\x95\x5c\x9b\xce\x2a\x5f\xfc\x06\x00\x00\xff\xff\x96\xa3\xfc\xdf\xfa\x00\x00\x00")

func errorTmplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_errorTmplHtml,
		"error.tmpl.html",
	)
}

func errorTmplHtml() (*asset, error) {
	bytes, err := errorTmplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "error.tmpl.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _headTmplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\x51\x6f\xda\x3c\x14\x7d\xff\x7e\x05\xf2\xf7\xc0\x43\x09\x09\x01\x42\xa9\x96\x4e\x53\x9f\xf6\x56\x6d\xfb\x03\x8e\x73\x81\xab\x26\x71\x14\x3b\xc5\x2c\xe3\xbf\xef\x3a\x46\x30\x68\xbb\x91\x08\x24\x2b\xf6\xf5\x3d\xe7\xdc\x63\x5b\x36\x4d\x93\xc2\x0a\x0b\x18\xb0\x0d\xf0\x94\xed\xf7\xff\x0d\x06\x9f\x72\xd0\x9c\xbe\xf4\x2b\x78\x0e\x31\x7b\x45\xd8\x96\xb2\xd2\xcc\x05\x85\x2c\x34\x14\x3a\x66\x5b\x4c\xf5\x26\x4e\xe1\x15\x05\x78\xed\x60\xe4\x32\x8e\x3f\x2c\x50\x23\xcf\x3c\x25\x78\x06\xf1\xe4\x72\x3a\xe7\x06\xf3\x3a\xff\x68\xba\x56\x50\xb5\x73\x3c\xa1\xe9\x42\xbe\x81\x13\x7d\x4e\xec\x35\xb2\xc7\xf7\xea\xe6\x65\x99\x81\x97\xcb\x04\xe9\xb3\x85\xc4\xa3\x80\x27\x78\x69\xe9\x2e\xbd\xec\x40\x11\xc9\xd5\x2c\x4a\x73\x5d\x2b\x2f\xe1\x54\xa0\xde\xbd\xa5\x4b\x32\x2e\x5e\x2c\xa1\x65\xcc\xb0\x78\x19\x6c\x2a\x58\xc5\x43\xdf\x5f\x51\x8a\x1a\xaf\xa5\x5c\x67\xc0\x4b\x54\x63\x21\x73\x5f\x28\xf5\x79\xc5\x73\xcc\x76\xf1\x77\x59\x57\x02\xee\x9e\x64\x0a\x77\xcf\x95\x7c\x98\x05\xc1\x68\x1e\x04\xbf\xbe\xc9\x44\x6a\xf9\x40\xdd\x11\x85\x50\xf3\x0c\xc5\x68\xea\x26\x4f\xa3\x43\x8f\x32\x86\x6d\x45\x15\x64\xf1\xb0\xad\x50\x6d\x00\xb4\x0b\xea\x5d\x09\xf1\x50\x83\xd1\x56\x78\xf8\x78\x2a\xd2\xa6\xb3\x53\x3a\x73\x55\xb3\xa6\x19\x7f\x51\x0a\xb4\x7a\xa6\x21\x9a\xfd\xde\xc7\x94\x6c\xa2\xde\x8d\x89\x80\x5d\x12\xb8\x15\xd3\xb2\x16\x1b\x0f\x69\x49\xd8\x40\xe1\x4f\x50\x31\x9b\x2f\xcc\x7c\xf1\x31\xe9\x8a\xbf\xda\x74\xe5\x5f\x12\x78\x2d\x70\x5c\x16\x6b\xb7\xd1\x57\x48\x4d\x26\x33\x43\xad\x8f\xd8\x01\xda\x49\x6e\x11\x9a\x45\xd8\x47\xac\x05\x76\x73\x36\xa3\xf2\x66\xfd\x9c\x39\x68\x27\xb9\x28\x30\x51\xd0\x47\xac\x05\x76\x73\x16\x06\x86\x5a\x2f\x67\x0e\xda\x6d\xcf\x22\xb3\x88\x7a\xed\x99\x05\x76\x73\x36\x0f\x0d\xb5\x5e\xce\x1c\xb4\x9b\xdc\x3d\xad\xc6\x7d\xbf\x85\x74\xd0\xf7\xe4\x9c\x44\x7b\x79\x30\xba\x77\xd7\xe0\xdb\xa4\x7f\x6b\x1c\x3a\xde\x64\x49\x4e\x96\xce\xc9\xb1\x52\x17\xbb\xb9\x14\x1d\xbe\xc9\xe1\xf4\x1d\xa5\x5c\xec\xd6\x52\xcb\xc8\x2c\xa3\x33\xa1\x36\x72\x7b\x47\x54\xfc\x85\x1f\x8a\xdc\x5a\x66\x1a\x9a\xe9\xf9\x0e\xb5\x91\xd3\xfb\x7a\x78\x14\x73\x65\x4f\x0e\x0a\xae\x91\x50\x3f\xe8\x69\x7c\x92\x99\xac\xd8\xe9\x05\xfc\x3f\xe5\xf3\x69\x28\xae\x40\x7e\xb5\x75\xfe\x81\xfc\x4b\x99\xb9\xd2\xf6\x15\x3e\xbf\xc7\x9a\x06\x8a\x94\xfe\xb9\xfc\x0e\x00\x00\xff\xff\x0c\x31\x0b\x6b\xcc\x08\x00\x00")

func headTmplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_headTmplHtml,
		"head.tmpl.html",
	)
}

func headTmplHtml() (*asset, error) {
	bytes, err := headTmplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "head.tmpl.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _headerTmplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\x31\x6e\xc6\x20\x0c\x46\xf7\x9e\xc2\x72\xe7\x86\x0b\x90\x48\x1d\xba\x76\xa9\xd4\xdd\x0a\x26\x41\x22\x50\x01\xcd\x62\x71\xf7\xd2\x26\xa4\xbf\x94\x0d\x7f\xcf\x96\x1f\x16\x31\x6c\x5d\x60\xc0\x95\xc9\x70\xc2\x5a\x9f\x00\xf4\x51\x4c\xed\xd9\x8a\x40\x3b\xcc\x9e\x72\x1e\xd1\xb3\x2d\x78\xc4\x0d\x10\xac\x89\xed\x88\xcf\x78\xf1\xb8\x44\x9c\x3e\x29\x90\x71\xdf\x9b\x56\x74\xf5\xe6\x2f\x0a\xbd\x2b\x73\xda\xdd\xcc\x2f\x81\x36\xc6\x49\xa4\xb8\xe2\x19\x86\x8f\x16\x73\x7a\x6f\x61\xad\xf0\xea\x7d\x9c\xa9\xc4\xa4\xd5\xef\xe8\xa9\xa2\x9a\xcb\xdd\x2a\xb9\x65\xfd\xd7\x12\x71\x16\x86\xb7\x8d\x9c\xff\xfb\xcc\x4d\xb6\x6d\xec\xf8\xc1\x50\x84\x83\x39\x07\xae\x3d\x5a\xf5\x4b\x74\xfc\x13\x00\x00\xff\xff\xe2\x89\x41\x0f\x31\x01\x00\x00")

func headerTmplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_headerTmplHtml,
		"header.tmpl.html",
	)
}

func headerTmplHtml() (*asset, error) {
	bytes, err := headerTmplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "header.tmpl.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _homeTmplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x56\xdf\x93\x9b\x36\x10\x7e\xf7\x5f\xb1\xa1\x37\x35\xee\x9d\xe1\xd2\xbc\xdd\x61\x66\x92\x26\x33\xcd\x4c\x7c\xcd\x24\x97\xce\xe4\xa9\x23\x60\x6d\x94\x03\xc9\x95\x84\x1d\xca\xf0\xbf\x77\x85\x00\xdb\x77\x76\x9a\xb6\x7e\x00\x0b\x7d\xfa\xf4\xed\x0f\xed\x2a\x7a\x96\xc9\xd4\xd4\x1b\x84\xdc\x94\x45\x3c\x89\x86\x17\xb2\x2c\x9e\x00\x44\x86\x9b\x02\xe3\xdf\x99\x60\x19\xaf\x4a\x68\x9a\xee\x03\x04\x1f\x51\x6d\x51\xdd\xb1\x12\xdb\x16\x5e\x16\x85\x4c\x99\x91\x2a\x0a\x1d\x9e\x56\x12\x12\xcb\x4d\xc1\x0c\x82\x67\xd9\x3c\x08\xda\x76\x12\x85\x8e\x39\x7a\x36\x9f\xc3\xdd\x6f\xf7\x6f\x6e\x60\x87\x30\x4d\xa4\x52\x72\x37\x05\x06\x85\x34\x20\x57\x60\x72\x04\x6d\xea\x82\x8b\xb5\x1d\xf2\x0c\x05\x31\xd7\xd9\x15\xe4\x28\x52\xb4\xf3\x13\x8d\x69\xa5\xe8\xe3\x5c\xa1\xdd\x26\x03\x83\xaa\xe4\x42\x16\x72\x5d\x03\x17\x23\x07\x02\x16\x58\x12\x81\xbe\x82\x82\x3f\xe0\xc8\x06\x4c\x64\x93\xa4\x40\xad\x69\x9b\x00\xe6\x73\xd2\x95\xc8\xac\x86\xb4\x60\x5a\x2f\xbc\x01\xb7\x51\x72\x4b\xff\xd5\xbc\x60\xb5\xac\x8c\x77\xca\x3c\x54\xce\x40\x72\x99\x4e\x15\xdf\x18\x0b\x02\x58\x55\x22\x35\x5c\x0a\x48\x73\x26\xd6\xf8\xca\x08\x9f\x93\x11\x02\x77\xbf\xde\x2f\xdf\xcd\xa0\xe9\x50\x00\x14\x85\xca\x4a\x0c\xd6\x68\xde\x38\xb5\xaf\xea\xb7\x19\x81\x67\x01\x17\x02\x95\x85\xc3\x02\xbc\x68\x25\x85\x81\x94\x8c\x54\x8b\xe9\x5a\xb1\x7a\x1a\x7b\x97\x3d\xdd\xa5\x17\x85\x76\x3a\xf6\x6e\x3b\xda\x4e\x4d\xb8\x97\x13\x95\x8c\x0b\xa7\x2b\xca\x9f\x0f\x56\x6e\xd8\x1a\xe7\x5d\x84\xe2\x5f\x14\x5a\x83\xc8\x2d\xb0\xa4\x88\xaf\x11\xde\x0a\x6d\x18\x39\x5c\xdb\x28\x9c\x0c\x3e\x45\xf4\x79\xcf\x99\xf1\xed\x40\x3a\x78\x55\xcf\x0b\xae\x9d\xc7\xce\x43\x7a\xff\x0d\xa0\x4e\x5d\xfc\x59\x56\x8a\xa2\xd8\x6f\xbf\xdf\xc5\x9a\x44\x2c\x6e\xd0\x34\xca\xba\x15\x2e\xb8\xc8\xf0\xeb\x15\x5c\xf4\x91\x86\x9b\x05\x04\xa3\xf6\x2e\x2c\xe7\xb7\xe7\x14\xc9\xc3\xcd\x4f\x80\xe6\x19\x1a\xc6\x0b\x7d\x00\xb3\x2a\x5f\xc4\x4d\xf3\x50\x25\x68\x3d\x01\xc1\xe8\x8f\x17\x47\xa8\x4d\x1c\x25\xce\xb3\x36\x0f\xee\x79\x89\x37\x51\x98\xc4\x3f\x8a\x44\x6f\x6e\xdd\x33\xd2\x1b\x26\x86\x3d\x2b\xc1\xbf\x1a\x42\x79\x90\x31\xc3\xe6\xc3\x70\xd1\x34\xc1\xc0\x62\x49\x82\x4f\x34\xd1\xb6\xf1\xe3\xcf\x1f\x8d\x22\xc1\x56\x87\x25\x3d\x52\x12\x6e\x1e\x09\x3b\x18\xd1\x38\x89\x97\xb2\x22\xdf\x59\x3b\xac\xc2\x28\x51\xe1\x23\x08\xb7\xdb\xd9\xf9\x0f\x52\x9a\xb6\x0d\xa3\x90\xc7\xd1\xce\xe2\x28\x12\xf6\x14\xf2\xed\x91\x33\xbe\xb5\xfb\x13\x1f\xa7\x6c\x4b\x76\x1c\xfb\xb8\x93\xf5\xaa\x47\xc0\x7b\x66\xe8\x8c\x0b\x7d\x5a\xdd\x90\x0c\xc1\x80\x1f\xe0\x47\x3a\x1c\x92\x8e\xea\x29\x02\x14\xd9\x79\xd1\x07\x79\xd7\x0d\x4f\xe5\x89\xc2\xad\x7c\xc0\xec\x38\x4f\x18\xe4\x0a\x57\x0b\x8f\xb6\x7d\xcd\x74\x9e\x48\xa6\xb2\x4f\x1f\xde\xb5\xad\x37\xae\xaf\x8c\x91\x62\xbe\xa1\x01\x79\xd0\x03\xc3\x14\x15\x82\x85\xf7\x47\x52\x30\xf1\xe0\xc5\xe3\xb2\x28\x64\xff\x4d\x10\x15\xbd\x85\x97\xa1\x36\x4a\xd6\x54\x85\x9a\xc6\x1d\x19\xd2\x70\x56\xaa\x03\x7f\x5b\xa8\x14\x69\xc1\xd3\x87\x85\xb7\xaf\x6f\xd3\x53\xbb\x4c\xaf\x60\xda\x13\xda\x40\xfa\x86\x3d\x50\x4d\x61\xb0\xc2\x1d\x50\x1d\x97\x22\xd3\x33\x08\x82\x60\x3a\x23\x63\x1d\xee\xbc\xa9\x47\x05\x00\x0b\xbd\x4f\xb4\x3b\x29\x10\x56\x94\xc5\x59\x30\x79\x1c\xd1\x7d\xbe\xbb\xe3\x66\x1d\x92\x76\x05\x8f\x94\x7a\xf1\xa1\xe9\xae\x0e\xfe\x7b\xcb\x47\x3a\x6b\xae\x3b\x95\xff\x64\x6c\x5f\x72\xef\x70\x67\xed\x3d\x3c\xb4\x63\xee\x35\xcd\x8e\x9b\x1c\x82\x25\x85\xd4\xd6\xe4\xf9\x60\x50\xae\xe2\xc7\x96\x75\xdd\x41\xf3\xbf\x70\xe1\xfd\xec\xc5\x2e\xcf\x5d\x4f\x78\xcc\x79\xe0\x98\x11\x10\x85\x43\x87\xe8\x9b\x18\x68\x95\x76\x1e\x79\xa9\x35\x1a\xfd\x9e\x1c\x64\x8b\x4e\x38\x74\xc6\xb0\x94\x5d\xd7\xfa\x42\xa7\xf6\xa8\xd3\x7c\xe7\xf2\x2f\x7f\x56\xa8\xea\x73\xcb\xed\xdf\xb1\x81\x12\x81\xad\x6f\xf7\xf8\xd5\xf8\xb6\xca\x0f\xbd\x73\xcb\x14\xd8\xfa\x48\xb5\xbe\xdc\x50\x87\xb4\x73\x81\xad\x9c\xfe\xbe\x92\xce\x6e\x47\x68\x49\x10\x27\xda\x1f\x57\xfd\xf4\xfc\xfa\xfa\x3a\xb8\x3e\x40\xb9\x6b\xc3\x11\x59\xf7\x69\x60\xe2\x2b\xf0\x7b\xcc\x82\x9a\x32\x4b\xb4\x2c\x2a\x43\xd3\x63\x43\xef\x56\xda\xcb\x94\xef\x0d\x99\x35\xfd\x61\x3a\x66\xcd\x54\xa1\xa9\x94\x80\x15\xa3\xf4\xbd\xa5\x1e\x0e\x97\x50\x06\x2b\xa9\x4a\x66\x7c\x6f\xb9\x5c\xc2\xeb\xd7\x57\xf0\x99\x7e\x90\xdf\x94\xe5\x8d\xa6\xfc\x21\xf6\x4b\xba\x00\x50\x9a\x0c\x32\xe0\xa9\xc0\x2b\xf0\x56\x4a\x96\x77\x72\x37\x80\x5a\xb0\x47\xe4\x7f\x0a\x73\x94\xfe\x77\x2a\xd8\xfb\x63\x7f\x11\x69\x27\xf4\xb8\xf0\x87\x8b\xce\x2c\xa0\xc4\xcf\x6a\x7f\x88\xaf\x3f\xb8\xee\xc2\xf7\x82\x7d\xe0\x02\x64\x69\xfe\x14\x04\x10\x86\xd0\x09\xb6\x07\xcc\x5e\xf4\xf6\x29\xa0\x73\x59\x15\x74\x19\x94\xeb\x75\xd1\x5d\x14\x21\xe3\x9a\x6e\x6a\x35\x38\xf7\x06\x3d\xc3\x85\x6f\x72\xae\x67\x41\x47\x73\xb4\xc5\x51\xae\xf5\xb0\xd9\x2d\xb4\xa3\xcd\x27\xe7\x9d\xa5\xdd\xdb\x3d\xf7\x19\x4d\xfd\x8a\xae\x95\xf6\xed\xae\xd7\x7f\x07\x00\x00\xff\xff\x5f\x52\x1c\x23\x76\x0b\x00\x00")

func homeTmplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_homeTmplHtml,
		"home.tmpl.html",
	)
}

func homeTmplHtml() (*asset, error) {
	bytes, err := homeTmplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "home.tmpl.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rootTmplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x90\xcd\x6a\xc3\x30\x10\x84\xef\x79\x8a\xad\xee\xb1\xc8\x5d\x11\x94\x5e\x5a\x68\x7b\x29\xf4\xbe\xb5\x36\xb6\x40\xd2\x0a\x59\x09\x18\xe3\x77\xaf\x7e\x92\x5b\x4e\x3b\x98\x99\xcf\x33\x52\x2f\x86\xc7\xbc\x46\x82\x39\x7b\xa7\x0f\xea\x71\x08\x8d\x3e\x00\xa8\x6c\xb3\x23\xfd\x8b\x01\x8d\xbd\x7a\xd8\xb6\xf6\x01\x86\x1f\x4a\x37\x4a\xdf\xe8\x69\xdf\xe1\xd5\x39\x1e\x31\x73\x52\xb2\xfb\x4b\xb2\x38\xc9\x47\x87\x99\x40\x54\x9a\x80\x61\xdf\x0f\x4a\x76\xb2\xfa\x63\xb3\xc2\xe8\x70\x59\xce\xc2\x1a\x0a\x25\xb7\xc6\xc4\xb7\xa2\xd3\xd1\xe1\xca\xd7\x2c\x9e\x61\x28\x35\x50\xeb\xe6\xd1\x86\x2a\x8a\x9c\x4f\x0f\x58\xc4\x89\x8e\xed\x87\xfa\x2d\x51\xcd\x61\x30\xf0\x55\x06\x4c\x04\x1f\x61\xc9\x18\x46\x5a\x80\x2f\xcf\xb7\x94\x82\xa7\x3b\x33\xf6\x5b\x14\xc2\x9c\xe8\x72\x16\xdb\x36\xbc\x73\x75\x09\xfd\xc9\x13\xd8\xa0\x24\xde\xcd\xb2\xb9\x95\xec\x9d\x94\xac\xfb\xea\xed\xef\xf9\x1f\x00\x00\xff\xff\x9f\x7c\x3e\xfd\x67\x01\x00\x00")

func rootTmplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_rootTmplHtml,
		"root.tmpl.html",
	)
}

func rootTmplHtml() (*asset, error) {
	bytes, err := rootTmplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "root.tmpl.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _styleCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x55\xcd\x8e\xdb\x36\x10\xbe\xfb\x29\x06\x30\x0a\x24\x41\xe8\xc8\x8e\xb1\xde\x7a\x4f\xae\x9b\x05\x0a\x14\x3d\x34\x6d\xef\x94\x38\x92\xd8\x50\xa4\x40\x52\x6b\x3b\x8b\xbe\x7b\x87\x14\x25\xff\x69\x11\xc3\xb6\xc0\xd1\xfc\x7e\xf3\xcd\xf0\xd3\x07\xd8\x9b\xf6\x64\x65\x55\x7b\x58\x65\xcb\x07\xf8\xab\x46\xf8\x87\x6b\x2e\x64\xd7\xc0\xae\xf3\xb5\xb1\x6e\x01\x3b\xa5\x20\x2a\x39\xb0\xe8\xd0\xbe\xa0\x58\xc0\x87\x4f\x33\xb2\xff\xdb\x21\x98\x12\x7c\x2d\x1d\x38\xd3\xd9\x02\xa1\x30\x02\x81\x8e\x95\x79\x41\xab\x51\x40\x7e\x02\x0e\xbf\x7c\xfd\x95\x39\x7f\x52\x98\x0c\x95\x2c\x50\x93\xb1\xaf\xb9\x87\x82\x6b\xc8\x11\x4a\xd3\x69\x01\x52\x93\x10\xe1\xf7\xdf\xf6\x5f\xfe\xf8\xfa\x05\x4a\xa9\x30\x46\x9b\xe5\x46\x9c\xe0\x75\x06\xa4\xa7\x3d\x2b\x79\x23\xd5\x69\x0b\x7f\x9a\xdc\x78\xf3\x34\x88\x0f\x18\x12\xdd\xc2\x3a\xcb\x46\x59\x8c\xbb\x05\x6d\x6c\xc3\xd5\x28\x2d\x8c\x32\x76\x0b\xf3\xcf\x9b\xf5\x66\xfd\x1c\xc4\x64\x9c\x7f\x93\xe4\x3a\x1a\x35\xc6\x50\x59\xba\xda\x02\xd7\x5e\x72\x25\xb9\x43\x71\xf6\x29\xbf\x93\xcb\xe5\xaa\x3d\x06\x91\x92\x1a\x59\x9d\x42\xaf\xb2\x5e\xd8\x72\x21\xa2\x7d\x3a\x37\xdc\x56\x52\x8f\xc7\x00\x4f\xa9\xcc\x61\x0b\xb5\x14\x02\xf5\xd3\xec\xbf\xd9\x6c\x5e\x90\x73\x4e\xde\x6c\xac\xb4\x35\x4e\x7a\x69\xc8\xa8\x94\xc7\x3e\xba\x90\xae\x55\x9c\x0a\x2f\x15\x46\x3f\xe1\xc9\x84\xb4\x58\xf4\x9a\x54\x57\xd7\xe8\xf0\x66\x48\x68\x99\x65\x3f\x85\xf3\x41\x0a\x5f\x0f\xc7\x10\xac\x46\x2e\x52\xa4\x3b\xb7\xff\x76\xce\xcb\xf2\xc4\x42\x42\xa8\xc9\x8b\x6b\x79\x81\x2c\x47\x7f\x40\x8c\xee\x73\x5e\x7c\xab\x6c\xe8\xd9\x88\x65\x96\x3d\x7e\x7e\x8c\x58\xe6\xe6\xc8\x5c\xcd\x45\xa8\xcf\x56\x39\x7f\x97\x7d\x84\xf4\x5d\xac\xb2\xf7\x01\x04\x20\xf0\x60\x4d\xbf\x04\xc8\x55\x7a\xe7\xec\xd7\x8f\x13\x18\x0f\xc2\x14\xf8\x50\x4b\x8f\xb7\xbd\x79\xe8\x55\xa6\x68\xf1\x9d\x49\x2d\xf0\xb8\x85\x9f\xe9\xf3\x74\x05\x34\xcf\x1d\x01\xd8\x7b\xf3\xa6\x1d\xdb\xa5\xb0\xf4\xe9\x70\x01\x9d\x90\x2f\x11\xbe\xbe\xb7\xac\x57\x5a\x5e\x35\x9c\xd9\xa1\x0b\x37\xb6\x73\x2f\x3d\xcd\xc3\xeb\xb9\xf4\xcf\xd9\x9d\x0e\x36\x5c\xaa\x09\x1d\x4a\x0e\x8f\x9e\x11\x2d\x2b\x4a\x3a\x86\x88\x86\x8b\xa6\x62\x27\xc6\x8f\x34\x80\x41\xe1\x23\x04\xc1\xf1\x2c\x38\x0f\x50\x42\x69\x99\x22\x06\xbd\x0b\x3e\xdc\x01\x99\x54\x38\xb1\xec\x05\x99\xe0\x9e\xb7\x46\xea\x7b\x7f\xab\xa1\x82\xa2\xe6\x96\x36\xc6\x24\xb9\x22\x67\x0f\x96\x13\xbc\xe1\xff\x72\x1c\x18\xa9\xf1\xae\x9f\xe8\x58\xdd\x99\x81\xd1\xca\x79\x72\xfb\x83\x96\x0d\xec\xa0\xcd\xe0\x4d\x73\xee\xda\x22\xe6\x74\xd9\xb0\xa8\xbe\x1a\x51\x17\x9d\xe5\xc1\xa7\x63\x6f\xcd\xe1\x65\xac\x0b\x4a\xdc\xc4\x9a\x1a\x8e\x91\xa3\x86\x06\x49\x7a\x2a\x33\x5b\x3c\xf6\x61\x95\xe1\x61\x53\x30\xc5\x73\xec\x9b\x9d\x76\x47\x22\xd4\xc3\x15\xd9\xe7\xcf\xcf\xeb\xe7\xe5\xee\x26\xe3\x69\xa0\x6f\x77\x52\x6e\x2c\x75\xb8\x2f\x9b\x3a\x4f\x3b\x5b\x49\x01\xf3\x34\x05\xe9\xed\x40\xd8\x9b\xf7\x01\xc0\x21\x1c\xa3\x5a\x9a\xe9\xde\x8f\x4c\x5d\xad\xdb\xc9\x1c\x26\x97\xe5\x25\x9b\xe9\x66\xf0\x68\x63\xc5\x9d\x75\xa1\xe4\x48\xb5\x24\x4a\x20\x6c\x36\x9b\x89\x94\x16\x0e\x15\x2d\x42\xba\x75\x5e\x2f\xcc\x05\x96\xbc\x53\xfe\x87\x4b\xeb\xba\x53\x77\xce\xb7\xda\xf8\x77\x63\x84\xf7\xdb\x3a\x70\x36\x06\x9a\xf0\xba\xdf\xef\xfb\x0e\xa1\xb5\xc6\xb2\xc6\x55\x6f\x6d\xf4\x04\xf7\x34\x8f\xc6\x4b\x64\x58\x96\xe9\xf9\x46\x2d\xbb\x5d\x96\xf5\x2b\xee\x96\x75\x23\x35\xb4\xd1\xb1\xb8\xff\x03\x00\x00\xff\xff\xd0\x09\x6a\xf4\xfe\x07\x00\x00")

func styleCssBytes() ([]byte, error) {
	return bindataRead(
		_styleCss,
		"style.css",
	)
}

func styleCss() (*asset, error) {
	bytes, err := styleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "style.css", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
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
	"bad-request.tmpl.html": badRequestTmplHtml,
	"dash.js":               dashJs,
	"dashboard.tmpl.html":   dashboardTmplHtml,
	"error.tmpl.html":       errorTmplHtml,
	"head.tmpl.html":        headTmplHtml,
	"header.tmpl.html":      headerTmplHtml,
	"home.tmpl.html":        homeTmplHtml,
	"root.tmpl.html":        rootTmplHtml,
	"style.css":             styleCss,
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
	"bad-request.tmpl.html": &bintree{badRequestTmplHtml, map[string]*bintree{}},
	"dash.js":               &bintree{dashJs, map[string]*bintree{}},
	"dashboard.tmpl.html":   &bintree{dashboardTmplHtml, map[string]*bintree{}},
	"error.tmpl.html":       &bintree{errorTmplHtml, map[string]*bintree{}},
	"head.tmpl.html":        &bintree{headTmplHtml, map[string]*bintree{}},
	"header.tmpl.html":      &bintree{headerTmplHtml, map[string]*bintree{}},
	"home.tmpl.html":        &bintree{homeTmplHtml, map[string]*bintree{}},
	"root.tmpl.html":        &bintree{rootTmplHtml, map[string]*bintree{}},
	"style.css":             &bintree{styleCss, map[string]*bintree{}},
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
