package dir_cleaner

import (
	"dir_cleaner"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_CleanUpFilesToFolders_MoreFilesThanMinCountForOneDate_FilesMovedToFolder(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2022-01-28")
	f1 := FileConstructor(t1, "f1.png", 1, 0755, false)
	f2 := FileConstructor(t1, "f2.png", 1, 0755, false)

	var testFiles = map[string]TestFile{}
	testFiles["/f1.png"] = *f1
	testFiles["/f2.png"] = *f2

	var fs dir_cleaner.FileSystem = &TestFileSystem{Files: testFiles}

	var datesMap = map[time.Time][]dir_cleaner.File{}
	datesMap[t1] = []dir_cleaner.File{f1, f2}

	err := dir_cleaner.CleanUpFilesToFolders("", datesMap, 1, fs)
	assert.Nil(t, err)
	assert.Equal(t, *f1, testFiles["/2022-01-28/f1.png"])
	assert.Equal(t, *f2, testFiles["/2022-01-28/f2.png"])
}
