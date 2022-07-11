package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var customFileSystem fileSystem = osFS{}

type fileSystem interface {
	Stat(name string) (os.FileInfo, error)
	Mkdir(name string, perm os.FileMode) error
}

type file interface {
	Name() string
	Stat() (os.FileInfo, error)
	ModTime() time.Time
	IsDir() bool
}

// osFS implements fileSystem using the local disk.
type osFS struct{}

func (osFS) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }
func (osFS) Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

func main() {
	fmt.Println("DirCleaner started\nEnter path to directory:")
	var minFilesInDir = 2
	var pathToDirectory = ""

	// read path to directory to clean
	_, err := fmt.Scanf("%s", &pathToDirectory)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Validate path
	// skip if length is too short (something stupid)
	if len(pathToDirectory) < 5 {
		log.Fatal("Path is too short!")
		return
	}
	pathIsNotDirectory, err := isNotDirectory(pathToDirectory, customFileSystem)
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

	var datesMap = groupFilesByDate(files)

	err = cleanUpFilesToFolders(pathToDirectory, datesMap, minFilesInDir)
	if err != nil {
		log.Fatal("Can't clean up files.", err)
		return
	}
	fmt.Println("Directory has been cleaned successfully.")
}

func isNotDirectory(path string, fs fileSystem) (bool, error) {
	fileInfo, err := fs.Stat(path)
	if err != nil {
		return true, err
	}

	return !fileInfo.IsDir(), err
}

func groupFilesByDate(files []fs.FileInfo) map[time.Time][]fs.FileInfo {
	// build map with date as key and list of files as value
	var datesMap = map[time.Time][]fs.FileInfo{}
	for _, file := range files {
		if file.IsDir() == false {
			// remove hours minutes and seconds from file date
			hours := -time.Duration(file.ModTime().Hour())
			minutes := -time.Duration(file.ModTime().Minute())
			seconds := -time.Duration(file.ModTime().Second())
			keyTime := file.ModTime().Add(time.Hour*hours + time.Minute*minutes + time.Second*seconds)

			if datesMap[keyTime] == nil {
				datesMap[keyTime] = []fs.FileInfo{file}
			} else {
				var actualKeyFiles = datesMap[keyTime]
				datesMap[keyTime] = append(actualKeyFiles, file)
			}
		}
	}
	return datesMap
}

func cleanUpFilesToFolders(pathToDirectory string, datesFilesMap map[time.Time][]fs.FileInfo, minFilesInDir int) error {
	for keyDate, listOfFiles := range datesFilesMap {
		// create folders
		var dirName = ""
		createFolderForFiles := len(listOfFiles) > minFilesInDir
		if createFolderForFiles {
			dirName = pathToDirectory + "/" + keyDate.Format("2006-01-02")
			err := os.Mkdir(dirName, 0774)
			if err != nil {
				log.Println("Can't create directory "+dirName, err)
				dirName = ""
			}
		}

		// move files to specific folder
		if dirName != "" {
			for _, file := range listOfFiles {
				oldLocation := pathToDirectory + "/" + file.Name()
				newLocation := dirName + "/" + file.Name()
				err := os.Rename(oldLocation, newLocation)
				if err != nil {
					log.Println("Can't move file from "+oldLocation+" to "+newLocation, err)
					return err
				}
			}
		}
	}
	return nil
}
