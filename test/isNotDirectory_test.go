package dir_cleaner

import (
	"dir_cleaner"
	"testing"
)

func Test_IsNotDirectory_pathToFile_returnsTrue(t *testing.T) {
	var fs dir_cleaner.FileSystem = testFS{}
	actual, _ := dir_cleaner.IsNotDirectory("/pathToFile", fs)
	expected := true

	if expected != actual {
		t.Errorf("got %t, wanted %t", actual, expected)
	}
}

func Test_IsNotDirectory_pathToDirectory_returnsFalse(t *testing.T) {
	var fs dir_cleaner.FileSystem = testFS{}
	actual, _ := dir_cleaner.IsNotDirectory("/pathToDir/", fs)
	expected := false

	if expected != actual {
		t.Errorf("got %t, wanted %t", actual, expected)
	}
}
