// Code generated by vfsgen; DO NOT EDIT.

package fixtures

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Fixtures statically implements the virtual filesystem provided to vfsgen.
var Fixtures = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2022, 9, 23, 10, 43, 51, 567531086, time.UTC),
		},
		"/download.yaml": &vfsgen۰CompressedFileInfo{
			name:             "download.yaml",
			modTime:          time.Date(2022, 9, 23, 10, 43, 51, 567531086, time.UTC),
			uncompressedSize: 9992,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xdc\x5a\x5f\x6f\xdb\x38\x12\x7f\xf7\xa7\x20\xfc\x92\xdd\xc3\x4a\x69\xda\x6b\x51\x08\xd8\x07\xd5\x76\x1b\xe3\x12\xc7\xb0\x9d\xdd\x3d\x1c\x0e\x01\x4d\x8e\x6d\x5e\x28\x52\x4b\x52\x8e\x75\x41\xbf\xfb\x81\xd4\x7f\x59\x76\x93\x6c\xeb\xeb\x6e\x1e\x1a\x85\x1c\xce\xfc\xe6\x2f\x87\x64\x3d\xcf\xeb\xe1\x98\xfd\x02\x4a\x33\x29\x02\xb4\xbd\xe8\xdd\x33\x41\x03\x34\xc1\x11\xe8\x18\x13\xe8\x45\x60\x30\xc5\x06\x07\x3d\x84\x38\x5e\x02\xd7\xf6\x0b\x21\x22\x85\x51\x92\x7b\x31\xc7\x02\x82\xe2\x4f\x0e\xca\x8b\xb0\xc0\x6b\x50\x3d\x84\x04\x8e\x20\x40\x31\x56\x98\x27\xda\xd3\xa9\x36\x10\xf5\x0e\xca\x9c\x83\xda\x32\x02\x21\x21\x32\x11\xa6\x21\x38\x63\x94\x31\xf0\x34\xce\x47\x1c\xc0\x4e\xfe\x8f\x8f\x1e\x62\x2b\x84\x05\x45\xfe\x80\x27\xda\x80\xf2\xaf\x73\x7e\xfe\x95\xd3\x02\xfd\x00\xbf\xa3\x1f\x98\xa0\xb0\x3b\x4c\xd3\xcf\x79\xfb\x14\xb6\xe7\xf7\xc9\x12\x94\x00\x03\x7a\xaa\xe4\x96\x51\x50\xfd\x1f\x51\xff\x66\x3a\x9a\xcc\x2f\xc7\x1f\x17\xfd\x1f\xd1\xe7\xcf\x0d\xcd\x34\x90\x44\x31\x93\xfa\x32\x06\xa1\x37\x6c\x65\x7c\x26\xcf\x6b\xfa\x66\xd3\x03\x29\x0c\xec\xcc\x40\x0a\x6d\x14\x66\xc2\xe8\x0e\xdd\x0b\x25\x63\xc5\xb6\x8c\xc3\x1a\xa8\xa7\x09\xe9\x21\x84\x85\x90\x06\x1b\x26\x45\xee\x18\x26\x08\x4f\x28\xf8\x0a\x38\x60\x0d\x4d\xe1\x6c\x19\x79\x84\xcb\x84\xe6\x5e\xa2\x01\xea\x1b\x95\x40\xff\xcb\x4b\x35\xf0\x55\xb1\xca\xdb\xb0\xf5\xc6\xc3\x5b\xcc\x38\x5e\x32\xce\x4c\xfa\x0c\x3e\x4c\xac\x39\x78\x42\x52\xf0\x28\x6c\x81\xcb\x18\x54\x63\x79\x65\x68\x4b\x4f\x41\x13\xc5\x62\xe3\x4c\x7a\x56\xe9\x8f\x30\xe7\xf2\x41\x23\x4c\x08\x68\x8d\x8c\xb4\x03\xa8\x3e\x2f\x28\xda\x48\x6d\x1c\x4f\x84\x56\x80\x4d\xa2\x40\xbb\x71\xb3\x01\x94\x23\xb7\x2b\x55\x22\x10\xb6\x33\x29\x4a\x34\xa8\x9f\xdc\xd7\x5a\xc9\x24\xce\x3e\x57\xfa\x53\xf1\x07\x45\x0f\xcc\x6c\x72\x9e\x76\x6e\x3e\xba\x62\x22\xd9\xb9\x0c\x80\x9d\xf1\x11\xfa\x35\x9c\x4d\xc6\x93\x4f\x01\x32\x1b\xa6\x11\xd3\x4e\x5a\x24\xb5\x41\x0a\x38\xde\x01\x45\xf3\xc1\xc0\xb1\xd2\x1b\x99\x70\x8a\x96\x60\xc5\xd2\x9c\xa9\x14\x3c\x45\x2b\xa9\x10\xc9\xa2\x12\x61\x1a\x31\xc1\x6c\x70\x58\x23\xf8\xe8\x93\xc2\xc2\x38\x18\x88\xe0\xc4\x8d\x9d\xb9\xb5\x9d\xf6\x26\x0a\xb0\x01\xcf\x72\xad\x59\xb9\x9e\xca\x0a\x62\xef\x41\xaa\x7b\x2e\xb1\x8d\x87\xee\x50\xeb\xf7\x9c\xbd\x2f\xa5\x36\x43\xa6\x7e\x91\x3c\x89\x60\xca\x93\x35\x13\x01\xb2\x4c\xab\xe9\xf1\x74\xd0\x1e\x9a\x80\xb1\x02\xda\xc3\xd3\xf1\x70\x6f\x48\x2a\xa3\xeb\x83\xd3\x02\xc5\x48\x13\xcc\x71\x16\x06\x1d\xd3\xd4\xa6\x11\x66\xc2\xc6\x52\x35\x0d\x74\x80\xe3\xcc\xd1\x0c\x74\xd0\xf3\xd0\xd9\xdf\xce\x8a\xa9\x5b\xa1\xf1\x0a\xe6\xa9\x26\x86\x97\x73\x14\x56\x38\xe1\x26\xa4\xcd\x95\x48\x24\x9c\xf7\xf2\x38\xb0\x66\x33\x69\x0c\x01\x9a\x25\x22\xd4\xa1\x48\x7b\x2e\x58\x1c\x93\xac\x0c\x05\xb9\xf3\x3c\xe7\x3c\x5d\x8d\xdb\xb8\xaf\xfd\x19\x61\x4b\xa5\x7b\xb1\x62\x52\xb9\x3c\x72\x82\x14\x60\x7a\x23\x78\x3a\x93\xd2\x7c\x64\x1c\x72\x6a\xb4\xc2\x5c\x43\x4f\xc1\xef\x09\x53\x40\x87\x4a\xc6\x1d\x28\x95\x05\x75\xab\x41\x75\xe0\xd4\xe0\xa2\x35\xaf\x39\x9d\x04\x84\xc8\x28\x9e\x2a\xb9\xb2\x72\x0b\xb3\xe8\x24\x8e\x39\x44\x20\x0c\xe6\x9f\x72\x5d\xf7\xd6\xda\xdc\xa9\xdb\x40\x67\x15\x1d\x67\x15\x3d\x68\x16\xea\x20\xb7\xf4\x53\xc9\xab\xea\xff\xc4\x05\x4c\xac\x15\x68\xed\x89\x35\x13\xbb\x17\x2d\x72\xce\xd3\xb6\x98\x3f\x75\xf9\x1a\x1b\xb8\x07\x88\x0b\xc7\xbf\x60\x5d\x12\x53\x9b\xaf\xe5\x06\xe7\xb9\x54\xed\x6d\x5d\xc6\x95\xfe\x28\x76\x39\x10\xb4\xbd\xed\xa8\x25\x26\x3e\x4e\xcc\x46\x2a\xf6\xdf\xac\x64\xdc\xbf\xd7\x8d\xad\x27\xdf\xef\x66\x92\xc3\xe1\xbd\x26\x28\x36\x72\x95\xe4\x81\x80\x63\x56\xf9\x3e\xc3\x61\xab\x87\x96\x89\x22\xd0\x18\xdc\x82\x5a\xd6\x06\x3c\x24\xa4\x98\xe5\x84\xb7\xb3\xab\xe3\xb4\x2d\x31\x79\x26\xf9\xb5\xbd\x78\x5f\x6c\x69\x2e\xdd\xe4\x98\x55\x3f\xf7\x49\x81\x43\xfe\xb9\x06\xe3\x7e\x73\xa6\xb3\x8f\x18\x1b\xb2\x71\x5f\x99\xf9\xdd\xe7\x83\x1b\xfc\xa3\x78\xce\xb5\xc1\x26\x69\xc1\x2a\x00\xec\xc9\x7d\x99\x34\x83\xf5\x3d\x07\xf3\x9d\xe8\x5e\xa0\x39\x95\xe6\xdf\x91\xda\xcf\xd5\xb9\xd5\x0b\x3f\x2f\x75\x3f\x30\x41\x99\x58\x1f\xc9\xe0\x56\x95\xc9\x13\xda\x53\x92\xc3\x32\x5f\x6c\xbf\x67\xb0\xb2\x6b\x0b\xa5\x8f\xe0\xe8\x21\xb4\x5f\x41\x0e\xd5\x0d\x9d\x2c\xff\x03\xc4\xb8\xd2\xd1\xd9\xe2\x3f\xb3\xb1\x3f\xce\xa4\xd8\x4e\xbe\x74\x36\xf8\x26\xa5\x32\x56\x72\x97\x3a\xbb\x1e\xac\x96\x56\x0c\x08\xc3\x48\xcb\x9e\xed\x28\x92\xf7\x20\x14\x6c\x19\x3c\x74\xc7\x75\x27\xe3\x0e\x3f\xb5\xf8\xe6\xce\xc8\xba\xe5\xa3\xfc\x4f\x1b\x93\x95\xe5\x8a\x88\x3c\xe6\xc0\x6f\x10\xad\x35\xd7\x7d\xff\x01\xab\xc9\x06\x68\xc2\x99\x58\xd7\x9c\xd0\x38\x21\x2d\xec\x99\xa3\xe8\x26\x11\xe1\x58\xeb\xd6\x31\xc3\x9d\x2d\x72\x39\x28\x6f\x4b\x50\x2c\xa9\x76\x27\x0f\x3f\xf7\xe9\x34\x67\x31\xb0\x1c\x8e\x1c\x48\x8b\x3e\x97\x28\x66\x63\x9b\xf7\xb6\x98\x27\x10\xa0\x8b\x57\xc5\xcf\x4b\x22\xea\x0b\xe9\xc6\x01\x53\x50\x1e\x70\x20\x76\x79\xe6\xbd\xa3\x71\x73\x20\x29\xfb\xfd\xfd\x4c\x21\x52\xac\xd8\x3a\xc2\xf1\x81\x2a\x5e\xee\x20\x0f\x65\x39\xaf\x6d\x39\xb5\xcd\xa4\x2a\xf7\xf9\x3e\xf4\x4c\xf1\x47\xf7\x92\x3d\x39\x4f\x62\x0e\x5b\x10\xe6\xab\xe5\xfd\xd3\x12\xbe\xcb\x59\xdf\x30\xd7\x3b\x93\xbc\x33\x60\xbe\xff\x74\xaf\x2a\xac\x0b\x8a\x6b\x1c\x77\xd9\xd9\x15\xb0\x2c\x6c\x8e\x72\x6f\x9e\xfa\x71\x1c\xfb\xcd\x9b\x96\xe2\x72\x67\x99\x06\xe8\x12\x78\xd4\x75\xbb\x64\xc5\xfb\x1b\xe0\x91\xaf\x37\xe7\xf9\x7d\x83\x97\xe1\xd8\xbe\xf6\x98\x58\x29\x7c\x9c\xae\x1b\x5a\xa1\xd1\xc6\x18\x7b\xf4\xdc\xa5\x01\xea\x3f\x3e\xb2\x55\x75\x35\x37\x8f\x81\xf8\x6e\x2a\xb3\xc5\xe7\xcf\x8f\x8f\xe8\x18\x81\x3f\x12\x78\xc9\xc1\x1e\x8f\xec\x41\xe9\x30\xdd\x65\x21\x13\x79\x96\x14\x04\x2d\x7f\xf5\x73\x48\xfa\xff\x81\x49\x1f\x01\x25\xe4\xa9\x11\x4d\xe4\x11\x38\x2e\x04\xc3\xc4\x6c\x4e\x08\x68\x5a\xc8\xec\x86\xb4\x94\xd2\x68\xa3\x70\x3c\x08\x4f\x08\xea\x43\x25\xb5\x1b\x96\xbb\x7a\x1a\x0b\x77\x35\x0c\x25\xf5\x57\x45\x88\x8e\x43\x0c\x3b\x11\x74\xa1\x7d\x51\x35\x52\xc0\x71\xea\xe1\x35\x08\xf3\x17\xa9\x49\x79\x87\x33\x1e\x06\xa8\xe9\xa4\xd6\x6b\x41\xe6\xa3\x67\x3c\x28\x94\x8c\xfb\xa8\xb4\xba\xdb\xb2\x39\x4e\x75\x80\xce\x0e\x48\x0b\x2b\xfd\xb3\x70\x3c\x24\xb3\x46\x88\xfa\x19\xd7\x7e\xcd\xcf\x67\x7b\x0e\xc6\x71\xac\xab\xdd\x7d\x08\x31\x97\x69\x04\x9d\x6f\x31\x35\x37\xff\x59\xfd\xab\x63\x20\x56\x80\x76\x5d\x81\x54\xb9\x30\xdb\x4c\x5d\xd5\x30\x3b\xd4\x6d\x85\x15\xc4\x9c\x11\xac\x03\x74\x91\x95\x3f\x77\x6b\x38\x04\x4c\x39\x13\x30\x07\x22\x05\xb5\x93\xef\x5f\xbd\xea\x21\x64\x20\x8a\x39\x36\x50\xa9\x53\x18\xd3\xfe\xf0\x86\xb0\x4e\x71\x08\x15\x60\x1d\xc1\x6a\xc5\x04\x33\x69\xb5\x44\x48\x0a\xe1\xde\xa8\x85\x99\x5f\x14\x27\x8a\x89\xf5\xbc\x3c\x3c\x8c\xd7\x42\x96\xc3\xa3\x1d\x10\xf7\x8a\x50\x5f\x99\xf1\x9c\xe7\xa6\x59\x80\x8a\x74\x73\xda\x76\x8e\xce\x56\xa3\x5d\x6c\x75\xaf\xdc\xd5\xa6\xba\x87\x34\x68\xbd\xe9\x48\xdd\x41\x89\x90\x8c\x41\x61\xeb\x09\x34\x16\x9d\x04\xee\x6c\xd1\x29\x26\x13\xc5\x99\x48\x76\x4f\x05\x81\x15\xd9\x7c\x1b\x18\x38\xa2\xef\xfe\x7e\x78\x56\x45\xe5\xac\x01\x15\x31\xe1\x82\xfd\x93\xc2\x04\xa6\xa0\x98\xa4\x55\xfc\xbc\xca\xe9\xe2\xfa\x69\x6c\x72\xfc\x00\x96\xad\xd0\x8d\x2e\x74\xb2\xd7\xc7\x66\xef\x74\xcc\x94\xcf\x26\xa5\x3e\x5e\xd1\xf5\x82\xf1\x38\x8b\x98\xa9\x9c\xc5\x22\xbc\x86\x00\xf5\x97\x89\x4e\x97\x72\x17\x5c\xf8\x6f\xde\xf4\xcb\x59\x22\xa3\x08\xdb\xca\xf1\xaf\xbe\xde\xf4\x7f\x42\x7d\x8f\xd8\x7f\x13\xc7\x04\x79\x02\xbd\x7b\xfb\xf6\xcd\xbb\xfe\xbf\xcb\x05\xba\xf9\x02\x5a\x37\x68\xf5\xf8\x94\x3f\xe7\x14\x22\x0e\xa1\x6d\x27\x4c\x1d\xee\xe3\x23\xf2\x87\xf2\x41\x70\x89\xe9\xd0\x56\xc6\x99\x25\x0e\x2d\xed\xd8\x92\xa0\x6c\x6f\xce\xb3\x4b\xad\xb5\x55\xc1\xf3\x22\x49\xe1\x67\xc2\x19\x08\xe3\xb4\xf1\xb8\x5c\x7b\x1c\xb6\xc0\x7f\x7e\x53\xd3\x02\xc4\xb6\x42\x5e\xc0\x99\xde\x0c\xef\x26\xe1\xf5\xa8\xd7\x8a\x9a\x8f\x4a\x46\xcd\xc0\x59\x31\xe0\x34\x3f\xe7\xec\x8d\x4f\xb1\xed\xa9\x8a\x8a\xe1\x5b\xde\x07\x45\xcd\xa7\xe1\xe0\x1b\xc8\xcb\xfe\x1b\x40\x5b\xe8\x75\xf8\xdb\xdd\x70\x1c\x5e\xcd\xbf\x2c\x90\x14\xdd\xc2\x3f\x20\xed\x90\x7b\xa4\x6d\xa8\xff\xb8\x04\x8e\xf0\x6e\xc8\x30\x6f\x57\x0f\xe9\xee\x3e\x30\x6f\x84\x4a\x85\xd5\xe2\xbc\xb9\x5d\xdc\x4d\x67\x37\xbf\xfd\xf3\x2b\xe1\x6d\x1d\xba\xf6\x90\x56\xa7\x86\x3f\x80\xf5\x2e\xbc\x5d\x5c\x8e\x26\x8b\xf1\x20\x5c\x8c\x6f\x26\x27\x82\x5e\x36\xf3\xcf\x43\x7e\xb9\x58\x4c\x4f\x6e\xe2\x17\x58\xd8\xe2\x9c\xff\x29\x62\x61\x72\x73\x52\x98\xf9\x99\xf2\x79\x18\x3f\xdc\xdc\x2c\xe6\x8b\x59\x38\xbd\x1b\x84\x77\x83\xd1\x6c\x71\x22\xb0\xb5\xe3\xdd\xf3\x00\x87\x57\x57\x37\xbf\xde\x8d\x27\xf3\xd1\xe0\x76\x36\xba\x2b\xf1\x9f\x08\x77\xf7\xf9\xef\x69\x2a\x64\x4f\xc0\xd7\x76\x3f\x6f\x34\x1f\xde\xd3\x6a\x68\x64\x17\x66\x05\xfe\x1c\x0c\x39\x6f\x91\x34\x6e\x0a\xeb\x0d\x24\x68\xd3\xea\x75\x48\x9c\xb8\x0b\xde\xa8\xc9\x1f\x22\xa9\xd2\x00\x5d\xbc\x7e\x7f\xcd\x6a\x33\x59\x0f\xd1\xc1\xe1\xed\x21\x0e\x6f\x2f\x5e\x5f\x33\x37\xd2\xab\x6b\x7e\x6c\xcf\x6f\x2b\x5c\x3a\xab\x2e\xf7\xe0\xaa\xff\x05\x00\x00\xff\xff\xf7\xb8\xaf\x00\x08\x27\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/download.yaml"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
