package repositories

import "bmsp-backend-service/models"

func (r Repositories) GetDocumentList(customerName string) ([]models.Document, error) {
	var documents []models.Document
	if err := r.db.Where("customer = ?", customerName).Preload("DocumentLines").Find(&documents).Error; err != nil {
		return nil, err
	}
	return documents, nil
}

func (r Repositories) GetDocument(id int) (models.Document, error) {
	var document models.Document
	if err := r.db.Where("id = ?", id).Preload("DocumentLines").First(&document).Error; err != nil {
		return document, err
	}
	return document, nil
}

func (r Repositories) CreateDocument(document models.Document) error {
	if err := r.db.Create(&document).Error; err != nil {
		return err
	}
	return nil
}

func (r Repositories) UpdateStatusDocument(id int, status string) error {
	var document models.Document
	if err := r.db.Where("id = ?", id).First(&document).Error; err != nil {
		return err
	}
	document.Status = status
	if err := r.db.Save(&document).Error; err != nil {
		return err
	}
	return nil
}
