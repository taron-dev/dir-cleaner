package dir_cleaner

import (
	"fmt"
	"os"
	"time"
)

type testFS struct{}

func (testFS) Stat(name string) (os.FileInfo, error) {
	fmt.Println("Test stat: ", name)
	if name == "/pathToFile" {
		return fileInfo1{}, nil
	} else if name == "/pathToDir/" {
		return fileInfo2{}, nil
	}
	return nil, nil
}

func (testFS) Mkdir(name string, perm os.FileMode) error {
	return nil
}

type fileInfo1 struct{}

func (fileInfo1) Name() string       { return "test" }
func (fileInfo1) Size() int64        { return 1 }
func (fileInfo1) Mode() os.FileMode  { return 0755 }
func (fileInfo1) ModTime() time.Time { return time.Time{} }
func (fileInfo1) Sys() any           { return nil }
func (fileInfo1) IsDir() bool        { return false }

type fileInfo2 struct{}

func (fileInfo2) Name() string       { return "test/" }
func (fileInfo2) Size() int64        { return 1 }
func (fileInfo2) Mode() os.FileMode  { return 0755 }
func (fileInfo2) ModTime() time.Time { return time.Time{} }
func (fileInfo2) Sys() any           { return nil }
func (fileInfo2) IsDir() bool        { return true }

type testFile struct {
	time  time.Time
	name  string
	size  int64
	mode  os.FileMode
	isDir bool
}

func (f testFile) Name() string       { return f.name }
func (f testFile) Size() int64        { return f.size }
func (f testFile) Mode() os.FileMode  { return f.mode }
func (f testFile) ModTime() time.Time { return f.time }
func (f testFile) Sys() any           { return nil }
func (f testFile) IsDir() bool        { return f.isDir }

func FileConstructor(
	time time.Time,
	name string,
	size int64,
	mode os.FileMode,
	isDir bool) *testFile {
	return &testFile{
		time:  time,
		name:  name,
		size:  size,
		mode:  mode,
		isDir: isDir,
	}
}
