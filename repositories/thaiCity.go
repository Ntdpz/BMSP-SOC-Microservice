package repositories

import "bmsp-backend-service/models"

type ThaiCityFilter struct {
	ProvinceNameTh string `json:"province_name_th"`
	ProvinceNameEn string `json:"province_name_en"`
	ProvinceID     int    `json:"province_id"`

	AmphureNameTh string `json:"amphure_name_th"`
	AmphureNameEn string `json:"amphure_name_en"`
	AmphureID     int    `json:"amphure_id"`

	TambonNameTh string `json:"tambon_name_th"`
	TambonNameEn string `json:"tambon_name_en"`
	TambonID     int    `json:"tambon_id"`
}

func (r Repositories) GetProvince(filter ThaiCityFilter) ([]models.ThaiProvince, error) {
	var provinces []models.ThaiProvince

	if filter.ProvinceNameTh != "" {
		if err := r.db.Where("name_th = ?", filter.ProvinceNameTh).Find(&provinces).Error; err != nil {
			return nil, err
		}

	}

	if filter.ProvinceNameEn != "" {
		if err := r.db.Where("name_en = ?", filter.ProvinceNameEn).Find(&provinces).Error; err != nil {
			return nil, err
		}
	}

	return provinces, nil
}

func (r Repositories) GetAmphure(filter ThaiCityFilter) ([]models.ThaiAmphure, error) {
	var amphures []models.ThaiAmphure

	pipline := r.db

	if filter.AmphureNameTh != "" {

		pipline = pipline.Where("name_th LIKE ?", "%"+filter.AmphureNameTh+"%")

	}

	if filter.AmphureNameEn != "" {

		pipline = pipline.Where("name_en LIKE ?", "%"+filter.AmphureNameEn+"%")
	}

	if filter.ProvinceID != 0 {
		pipline = pipline.Where("province_id = ?", filter.ProvinceID)
	}

	if err := pipline.Find(&amphures).Error; err != nil {
		return nil, err
	}

	return amphures, nil
}

func (r Repositories) GetTambon(filter ThaiCityFilter) ([]models.ThaiTambons, error) {
	var tambons []models.ThaiTambons

	pipline := r.db

	if filter.TambonNameTh != "" {

		pipline = pipline.Where("name_th LIKE ?", "%"+filter.TambonNameTh+"%")

	}

	if filter.TambonNameEn != "" {

		pipline = pipline.Where("name_en LIKE ?", "%"+filter.TambonNameEn+"%")
	}

	if filter.AmphureID != 0 {
		pipline = pipline.Where("amphure_id = ?", filter.AmphureID)
	}

	if err := pipline.Find(&tambons).Error; err != nil {
		return nil, err
	}

	return tambons, nil
}
