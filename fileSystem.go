package dir_cleaner

import (
	"os"
	"time"
)

var fileSystem FileSystem = osFS{}

// osFS implements FileSystem using the local disk.
type osFS struct{}

type FileSystem interface {
	Stat(name string) (os.FileInfo, error)
	Mkdir(name string, perm os.FileMode) error
	Rename(oldLocation string, newLocation string) error
}

type File interface {
	Name() string
	ModTime() time.Time
	IsDir() bool
}

func (osFS) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (osFS) Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

func (osFS) Rename(oldLocation string, newLocation string) error {
	return os.Rename(oldLocation, newLocation)
}

