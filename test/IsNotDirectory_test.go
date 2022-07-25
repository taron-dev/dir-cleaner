package dir_cleaner

import (
	"dir_cleaner"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsNotDirectory_pathToFile_returnsTrue(t *testing.T) {
	var fs dir_cleaner.FileSystem = TestFileSystem{}
	actual, _ := dir_cleaner.IsNotDirectory("f1.png", fs)

	assert.True(t, actual)
}

func Test_IsNotDirectory_pathToDirectory_returnsFalse(t *testing.T) {
	var fs dir_cleaner.FileSystem = TestFileSystem{}
	actual, _ := dir_cleaner.IsNotDirectory("d1", fs)

	assert.False(t, actual)

}
