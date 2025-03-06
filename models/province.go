package models

type ThaiProvince struct {
	ID     int    `json:"id"`
	NameTh string `json:"name_th"`
	NameEn string `json:"name_en"`
}

func (t ThaiProvince) TableName() string {
	return "thai_provinces"
}

type ThaiAmphure struct {
	ID         int    `json:"id"`
	NameTh     string `json:"name_th"`
	NameEn     string `json:"name_en"`
	ProvinceID int    `json:"province_id"`
}

func (t ThaiAmphure) TableName() string {
	return "thai_amphures"
}

type ThaiTambons struct {
	ID        int    `json:"id"`
	NameTh    string `json:"name_th"`
	NameEn    string `json:"name_en"`
	AmphureID int    `json:"amphure_id"`
	ZipCode   int    `json:"zip_code"`
}

func (t ThaiTambons) TableName() string {
	return "thai_tambons"
}
