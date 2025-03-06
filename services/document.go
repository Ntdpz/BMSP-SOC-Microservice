package services

import "bmsp-backend-service/models"

func (s Services) UpdateDocument(id int, status string) error {

	return s.repo.UpdateStatusDocument(id, status)
}

func (s Services) CreateDocument(document models.Document) error {

	return s.repo.CreateDocument(document)

}
