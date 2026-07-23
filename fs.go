package gin

import (
	"net/http"
	"os"
)

type OnlyFilesFS struct {
	FileSystem http.FileSystem
}

func (o OnlyFilesFS) Open(name string) (http.File, error) {
	_ = "STUB: not implemented"
	return *new(http.File), nil
}

type neutralizedReaddirFile struct {
	http.File
}

func (n neutralizedReaddirFile) Readdir(_ int) ([]os.FileInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Dir(root string, listDirectory bool) http.FileSystem {
	_ = "STUB: not implemented"
	return *new(http.FileSystem)
}
