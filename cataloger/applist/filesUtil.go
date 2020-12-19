package applist

import (
	"io/ioutil"
	"log"
)

func _files(rootDir string) []string {
	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		log.Fatal(err)
	}
	var filePaths []string

	for _, file := range files {
		fullPath := JoinPath(rootDir, file.Name())
		if file.IsDir() {
			filePaths = append(filePaths, _files(fullPath)...)
		} else {
			if IsExecutable(file) {
				filePaths = append(filePaths, fullPath)
			}
		}
	}

	return filePaths

} //end _files
