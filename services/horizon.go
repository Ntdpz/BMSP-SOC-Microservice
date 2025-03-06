package services

import (
	"log"
	"os"
)

func (s Services) GetListFileFromHorizon() []string {

	dirPath := "./horizon"

	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	var fileNames []string

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames
}
