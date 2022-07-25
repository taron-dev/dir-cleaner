package test

import (
	"github.com/stretchr/testify/assert"
	dirCleaner "github.com/taron-dev/dir-cleaner/dir_cleaner_util"
	"github.com/taron-dev/dir-cleaner/file_system"
	"testing"
)

func Test_IsNotDirectory_pathToFile_returnsTrue(t *testing.T) {
	var fs file_system.FileSystem = FileSystem{}
	actual, _ := dirCleaner.IsNotDirectory("f1.png", fs)

	assert.True(t, actual)
}

func Test_IsNotDirectory_pathToDirectory_returnsFalse(t *testing.T) {
	var fs file_system.FileSystem = FileSystem{}
	actual, _ := dirCleaner.IsNotDirectory("d1", fs)

	assert.False(t, actual)

}
