package main

import (
	"fmt"
	"github.com/taron-dev/dir-cleaner/dir_cleaner_util"
	"github.com/taron-dev/dir-cleaner/file_system"
	"io/ioutil"
	"log"
)

func main() {
	log.Println("DirCleaner started\nEnter path to directory:")
	var minFilesInDir = 2
	var pathToDirectory = ""

	// read path to directory to clean
	_, err := fmt.Scanf("%s", &pathToDirectory)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Validate path
	pathIsNotDirectory, err := dir_cleaner_util.IsNotDirectory(pathToDirectory, file_system.OsFileSystem)
	if err != nil || pathIsNotDirectory {
		log.Fatal("Path is not directory!", err)
		return
	}

	// read all files in folder
	files, err := ioutil.ReadDir(pathToDirectory)
	if err != nil {
		log.Fatal(err)
		return
	}

	var datesMap = dir_cleaner_util.GroupFilesByDate(files)

	err = dir_cleaner_util.CleanUpFilesToFolders(pathToDirectory, datesMap, minFilesInDir, file_system.OsFileSystem)
	if err != nil {
		log.Fatal("Can't clean up files.", err)
		return
	}
	log.Println("Directory has been cleaned successfully.")
}
