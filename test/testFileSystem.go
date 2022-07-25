package test

import (
	"fmt"
	"os"
	"time"
)

type FileSystem struct {
	Files map[string]File
}

func (fs FileSystem) getFiles() map[string]File {
	return fs.Files
}

func (FileSystem) Stat(name string) (os.FileInfo, error) {
	t1, _ := time.Parse("2006-01-02", "2020-01-28")
	t2, _ := time.Parse("2006-01-02", "2020-01-29")

	f1 := FileConstructor(t1, "f1.png", 1, 0755, false)
	f2 := FileConstructor(t1, "f2.png", 1, 0755, false)
	f3 := FileConstructor(t2, "f3.png", 1, 0755, false)
	d1 := FileConstructor(t2, "d1", 1, 0755, true)

	switch {
	case name == "f1.png":
		return f1, nil
	case name == "f2.png":
		return f2, nil
	case name == "f3.png":
		return f3, nil
	case name == "d1":
		return d1, nil
	}
	return nil, nil
}

func (fs FileSystem) Mkdir(name string, perm os.FileMode) error {
	_, ok := fs.Files[name]
	if ok {
		return fmt.Errorf(name, "already exists")
	} else {
		fs.Files[name] = *FileConstructor(time.Now(), name, 1, perm, true)
		return nil
	}
}

func (fs FileSystem) Rename(oldLocation string, newLocation string) error {
	oldFile, oldLocationOk := fs.Files[oldLocation]
	_, newLocationOk := fs.Files[newLocation]
	if oldLocationOk && !newLocationOk {
		fs.Files[newLocation] = oldFile
		delete(fs.Files, oldLocation)
		return nil
	} else {
		return fmt.Errorf(oldLocation, "does not exists. Or", newLocation, "already exists")
	}
}

type File struct {
	time  time.Time
	name  string
	size  int64
	mode  os.FileMode
	isDir bool
}

func (f File) Name() string       { return f.name }
func (f File) Size() int64        { return f.size }
func (f File) Mode() os.FileMode  { return f.mode }
func (f File) ModTime() time.Time { return f.time }
func (f File) Sys() any           { return nil }
func (f File) IsDir() bool        { return f.isDir }

func FileConstructor(
	time time.Time,
	name string,
	size int64,
	mode os.FileMode,
	isDir bool) *File {
	return &File{
		time:  time,
		name:  name,
		size:  size,
		mode:  mode,
		isDir: isDir,
	}
}
