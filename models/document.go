package models

import (
	"time"

	"gorm.io/gorm"
)

type Document struct {
	Id                           int            `json:"id" gorm:"primaryKey;autoIncrement;column:id" xml:"-"`
	Customer                     string         `json:"customer" gorm:"column:customer" xml:"-"`
	DocumentID                   string         `json:"documentId" gorm:"column:document_id" xml:"ram:DocumentID"`
	DocumentTypeCode             string         `json:"documentTypeCode" gorm:"column:document_type_code"`
	IssueDateTime                string         `json:"issueDateTime" gorm:"column:issue_date_time"`
	TemplateCode                 string         `json:"templateCode" gorm:"column:template_code"`
	DocumentName                 string         `json:"documentName" gorm:"column:document_name"`
	IssuerDateTime               string         `json:"issuerDateTime" gorm:"column:issuer_date_time"`
	IssuerAssignedID             string         `json:"issuerAssignedId" gorm:"column:issuer_assigned_id"`
	ReferenceTypeCode            string         `json:"referenceTypeCode" gorm:"column:reference_type_code"`
	CreationDateTime             string         `json:"creationDateTime" gorm:"column:creation_date_time"`
	PurposeCode                  string         `json:"purposeCode" gorm:"column:purpose_code"`
	PurposeRemark                string         `json:"purposeRemark" gorm:"column:purpose_remark"`
	TermValue                    int            `json:"termValue" gorm:"column:term_value"`
	DueDateTime                  string         `json:"dueDateTime" gorm:"column:due_date_time"`
	TaxTypeCode                  string         `json:"taxTypeCode" gorm:"column:tax_type_code"`
	TaxCalculatedRate            float64        `json:"taxCalculatedRate" gorm:"column:tax_calculated_rate"`
	TaxCalculatedAmount          float64        `json:"taxCalculatedAmount" gorm:"column:tax_calculated_amount"`
	WhtPercent                   float64        `json:"whtPercent" gorm:"column:wht_percent"`
	OriginalInformationAmount    float64        `json:"originalInformationAmount" gorm:"column:original_information_amount"`
	LineTotalAmount              float64        `json:"lineTotalAmount" gorm:"column:line_total_amount"`
	DifferenceInformationAmount  float64        `json:"differenceInformationAmount" gorm:"column:difference_information_amount"`
	AllowanceTotalAmount         float64        `json:"allowanceTotalAmount" gorm:"column:allowance_total_amount"`
	TaxBasisTotalAmount          float64        `json:"taxBasisTotalAmount" gorm:"column:tax_basis_total_amount"`
	TaxTotalAmount               float64        `json:"taxTotalAmount" gorm:"column:tax_total_amount"`
	GrandTotalAmount             float64        `json:"grandTotalAmount" gorm:"column:grand_total_amount"`
	SellerTaxID                  string         `json:"sellerTaxId" gorm:"column:seller_tax_id"`
	SellerBranchCode             string         `json:"sellerBranchCode" gorm:"column:seller_branch_code"`
	SellerName                   string         `json:"sellerName" gorm:"column:seller_name"`
	SellerBranchName             string         `json:"sellerBranchName" gorm:"column:seller_branch_name"`
	SellerTaxType                string         `json:"sellerTaxType" gorm:"column:seller_tax_type"`
	SellerEmail                  string         `json:"sellerEmail" gorm:"column:seller_email"`
	SellerPhoneNumber            string         `json:"sellerPhoneNumber" gorm:"column:seller_phone_number"`
	SellerBuildingNumber         string         `json:"sellerBuildingNumber" gorm:"column:seller_building_number"`
	SellerLineOne                string         `json:"sellerLineOne" gorm:"column:seller_line_one"`
	SellerLineTwo                string         `json:"sellerLineTwo" gorm:"column:seller_line_two"`
	SellerCountrySubDivisionCode string         `json:"sellerCountrySubDivisionCode" gorm:"column:seller_country_sub_division_code"`
	SellerCountrySubDivisionName string         `json:"sellerCountrySubDivisionName" gorm:"column:seller_country_sub_division_name"`
	SellerCityCode               string         `json:"sellerCityCode" gorm:"column:seller_city_code"`
	SellerCityName               string         `json:"sellerCityName" gorm:"column:seller_city_name"`
	SellerCitySubDivisionCode    string         `json:"sellerCitySubDivisionCode" gorm:"column:seller_city_sub_division_code"`
	SellerCitySubDivisionName    string         `json:"sellerCitySubDivisionName" gorm:"column:seller_city_sub_division_name"`
	SellerPostcodeCode           string         `json:"sellerPostcodeCode" gorm:"column:seller_postcode_code"`
	BuyerTaxID                   string         `json:"buyerTaxId" gorm:"column:buyer_tax_id"`
	BuyerBranchCode              string         `json:"buyerBranchCode" gorm:"column:buyer_branch_code"`
	BuyerName                    string         `json:"buyerName" gorm:"column:buyer_name"`
	BuyerBranchName              string         `json:"buyerBranchName" gorm:"column:buyer_branch_name"`
	BuyerTaxType                 string         `json:"buyerTaxType" gorm:"column:buyer_tax_type"`
	BuyerEmail                   string         `json:"buyerEmail" gorm:"column:buyer_email"`
	BuyerPhoneNumber             string         `json:"buyerPhoneNumber" gorm:"column:buyer_phone_number"`
	BuyerBuildingNumber          string         `json:"buyerBuildingNumber" gorm:"column:buyer_building_number"`
	BuyerLineOne                 string         `json:"buyerLineOne" gorm:"column:buyer_line_one"`
	BuyerLineTwo                 string         `json:"buyerLineTwo" gorm:"column:buyer_line_two"`
	BuyerCountrySubDivisionCode  string         `json:"buyerCountrySubDivisionCode" gorm:"column:buyer_country_sub_division_code"`
	BuyerCountrySubDivisionName  string         `json:"buyerCountrySubDivisionName" gorm:"column:buyer_country_sub_division_name"`
	BuyerCityCode                string         `json:"buyerCityCode" gorm:"column:buyer_city_code"`
	BuyerCityName                string         `json:"buyerCityName" gorm:"column:buyer_city_name"`
	BuyerCitySubDivisionCode     string         `json:"buyerCitySubDivisionCode" gorm:"column:buyer_city_sub_division_code"`
	BuyerCitySubDivisionName     string         `json:"buyerCitySubDivisionName" gorm:"column:buyer_city_sub_division_name"`
	BuyerPostcodeCode            string         `json:"buyerPostcodeCode" gorm:"column:buyer_postcode_code"`
	Remark                       string         `json:"remark" gorm:"column:remark"`
	Channel                      string         `json:"channel" gorm:"column:channel"`
	DocumentLines                []DocumentLine `json:"documentLine" gorm:"foreignKey:DocumentRawId;references:Id;index:idx_document_lines"`
	Status                       string         `json:"status" gorm:"column:status"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type DocumentLine struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	DocumentID           string  `json:"documentId" gorm:"column:document_id"`
	DocumentRawId        int     `json:"documentRawId" gorm:"column:document_raw_id" xml:"-"`
	LineID               int     `json:"lineId" gorm:"column:line_id"`
	ProductName          string  `json:"productName" gorm:"column:product_name"`
	UnitName             string  `json:"unitName" gorm:"column:unit_name"`
	ChargeAmount         float64 `json:"chargeAmount" gorm:"column:charge_amount"`
	BilledQuantity       int     `json:"billedQuantity" gorm:"column:billed_quantity"`
	TaxCalculatedAmount  float64 `json:"taxCalculatedAmount" gorm:"column:tax_calculated_amount"`
	LineTotalAmount      float64 `json:"lineTotalAmount" gorm:"column:line_total_amount"`
	AllowanceTotalAmount float64 `json:"allowanceTotalAmount" gorm:"column:allowance_total_amount"`
	TaxBasisTotalAmount  float64 `json:"taxBasisTotalAmount" gorm:"column:tax_basis_total_amount"`
	TaxTotalAmount       float64 `json:"taxTotalAmount" gorm:"column:tax_total_amount"`
	GrandTotalAmount     float64 `json:"grandTotalAmount" gorm:"column:grand_total_amount"`
	WhtPercent           float64 `json:"whtPercent" gorm:"column:wht_percent"`
	WhtBaht              float64 `json:"whtBaht" gorm:"column:wht_baht"`
	NetAmountDue         float64 `json:"netAmountDue" gorm:"column:net_amount_due"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}
