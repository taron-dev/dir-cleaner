package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/taron-dev/dir-cleaner/dir_cleaner_util"

	"github.com/taron-dev/dir-cleaner/file_system"
	"testing"
	"time"
)

func Test_CleanUpFilesToFolders_MoreFilesThanMinCountForOneDate_FilesMovedToFolder(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2022-01-28")
	f1 := FileConstructor(t1, "f1.png", 1, 0755, false)
	f2 := FileConstructor(t1, "f2.png", 1, 0755, false)

	var testFiles = map[string]File{}
	testFiles["/f1.png"] = *f1
	testFiles["/f2.png"] = *f2

	var fs file_system.FileSystem = &FileSystem{Files: testFiles}

	var datesMap = map[time.Time][]file_system.File{}
	datesMap[t1] = []file_system.File{f1, f2}

	err := dir_cleaner_util.CleanUpFilesToFolders("", datesMap, 1, fs)
	assert.Nil(t, err)
	assert.Equal(t, *f1, testFiles["/2022-01-28/f1.png"])
	assert.Equal(t, *f2, testFiles["/2022-01-28/f2.png"])
}

func Test_CleanUpFilesToFolders_LessFilesThanMinCountForOneDate_FilesNotMovedToFolder(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2022-01-28")
	f1 := FileConstructor(t1, "f1.png", 1, 0755, false)
	f2 := FileConstructor(t1, "f2.png", 1, 0755, false)

	var testFiles = map[string]File{}
	testFiles["/f1.png"] = *f1
	testFiles["/f2.png"] = *f2

	var fs file_system.FileSystem = &FileSystem{Files: testFiles}

	var datesMap = map[time.Time][]file_system.File{}
	datesMap[t1] = []file_system.File{f1, f2}

	err := dir_cleaner_util.CleanUpFilesToFolders("", datesMap, 2, fs)
	assert.Nil(t, err)
	assert.Equal(t, *f1, testFiles["/f1.png"])
	assert.Equal(t, *f2, testFiles["/f2.png"])
}
