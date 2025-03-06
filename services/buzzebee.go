package services

import (
	"bmsp-backend-service/models"
	"log"
	"os"
)

func (s Services) GetListFileFromBuzzebee() []string {

	dirPath := "./buzzebee"

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

func (s Services) GetLenFileAndLenDBFromBuzzebee() (int, int, int) {

	dirPath := "./xml_buzzebee"

	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	var fileNames []string

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	lenFile := len(fileNames)

	lenSentDb, lenWaitingDb := s.repo.GetLen("buzzebee")

	return lenFile, int(lenSentDb), int(lenWaitingDb)
}

func (s Services) GetDocumentListFromBuzzebee() ([]models.Document, error) {

	return s.repo.GetDocumentList("buzzebee")

}

func (s Services) GetDocumentFromBuzzebee(id int) (models.Document, error) {

	return s.repo.GetDocument(id)

}
