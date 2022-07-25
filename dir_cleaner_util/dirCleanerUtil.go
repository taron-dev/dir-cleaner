package dir_cleaner_util

import (
	"github.com/taron-dev/dir-cleaner/file_system"
	"io/fs"
	"log"
	"time"
)

func IsNotDirectory(path string, fs file_system.FileSystem) (bool, error) {
	fileInfo, err := fs.Stat(path)
	if err != nil {
		return true, err
	}

	return !fileInfo.IsDir(), err
}

func GroupFilesByDate(files []fs.FileInfo) map[time.Time][]file_system.File {
	// build map with date as key and list of files as value
	var datesMap = map[time.Time][]file_system.File{}
	for _, file := range files {
		if file.IsDir() == false {
			// remove hours minutes and seconds from File date
			hours := -time.Duration(file.ModTime().Hour())
			minutes := -time.Duration(file.ModTime().Minute())
			seconds := -time.Duration(file.ModTime().Second())
			keyTime := file.ModTime().Add(time.Hour*hours + time.Minute*minutes + time.Second*seconds)

			if datesMap[keyTime] == nil {
				datesMap[keyTime] = []file_system.File{file.(file_system.File)}
			} else {
				var actualKeyFiles = datesMap[keyTime]
				datesMap[keyTime] = append(actualKeyFiles, file.(file_system.File))
			}
		}
	}
	return datesMap
}

func CleanUpFilesToFolders(pathToDirectory string, datesFilesMap map[time.Time][]file_system.File, minFilesInDir int, customFileSystem file_system.FileSystem) error {
	for keyDate, listOfFiles := range datesFilesMap {
		createFolderForFiles := len(listOfFiles) > minFilesInDir
		if createFolderForFiles {
			// create folders
			dirName := pathToDirectory + "/" + keyDate.Format("2006-01-02")
			err := customFileSystem.Mkdir(dirName, 0774)
			if err != nil {
				log.Println("Can't create directory "+dirName, err)
			}

			// move files to specific folder
			for _, file := range listOfFiles {
				oldLocation := pathToDirectory + "/" + file.Name()
				newLocation := dirName + "/" + file.Name()
				err := customFileSystem.Rename(oldLocation, newLocation)
				if err != nil {
					log.Println("Can't move File from "+oldLocation+" to "+newLocation, err)
					return err
				}
			}
		}
	}
	return nil
}
