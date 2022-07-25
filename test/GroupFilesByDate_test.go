package test

import (
	"github.com/stretchr/testify/assert"
	dirCleaner "github.com/taron-dev/dir-cleaner/dir_cleaner_util"
	"os"
	"testing"
	"time"
)

func Test_GroupFilesByDate_MultipleFiles_GroupedCorrectly(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2022-01-28")
	t2, _ := time.Parse("2006-01-02", "2022-01-29")

	f1 := FileConstructor(t1, "f1.png", 1, 0755, false)
	f2 := FileConstructor(t1, "f2.png", 1, 0755, false)
	f3 := FileConstructor(t2, "f3.png", 1, 0755, false)

	files := []os.FileInfo{f1, f2, f3}

	actualDatesMap := dirCleaner.GroupFilesByDate(files)
	assert.Equal(t, 2, len(actualDatesMap))

	actualFolder1 := actualDatesMap[t1]
	actualFolder2 := actualDatesMap[t2]
	assert.Equal(t, 2, len(actualFolder1))
	assert.Equal(t, 1, len(actualFolder2))

	actualFile1 := actualFolder1[0]
	actualFile2 := actualFolder1[1]
	actualFile3 := actualFolder2[0]
	assert.Equal(t, actualFile1.Name(), f1.name)
	assert.Equal(t, actualFile2.Name(), f2.name)
	assert.Equal(t, actualFile3.Name(), f3.name)
}
