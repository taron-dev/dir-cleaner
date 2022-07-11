package main

import (
	"os"
	"testing"
	"time"
)

type testFS struct{}

func (testFS) Stat(name string) (os.FileInfo, error) {
	return fileInfo{}, nil
}
func (testFS) Mkdir(name string, perm os.FileMode) error {
	return nil
}

type fileInfo struct{}

func (fileInfo) Name() string       { return "test" }
func (fileInfo) Size() int64        { return 1 }
func (fileInfo) Mode() os.FileMode  { return 0755 }
func (fileInfo) ModTime() time.Time { return time.Time{} }
func (fileInfo) Sys() any           { return nil }
func (fileInfo) IsDir() bool        { return true }

func isDirectory_pathToDirectory_returnsTrue(t *testing.T) {
	var fs fileSystem = testFS{}
	actual, _ := isNotDirectory("path", fs)
	expected := true

	if expected != actual {
		t.Errorf("got %t, wanted %t", actual, expected)
	}

}
