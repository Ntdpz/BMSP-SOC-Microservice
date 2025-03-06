package services

import (
	"bmsp-backend-service/models"
	"bmsp-backend-service/modelsxml"
	"bmsp-backend-service/repositories"
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func (s Services) CreateXML(doc models.Document) error {
	// สร้าง Line Items
	tmpItem := []modelsxml.SupplyChainTradeLineItem{}

	buyerCountrySubDivisionID := ""
	buyerBuildingNumber := extractHouseNumber(doc.BuyerLineOne)

	doc.BuyerTaxID = doc.BuyerTaxID + doc.BuyerBranchCode
	doc.SellerTaxID = doc.SellerTaxID + doc.SellerBranchCode

	{
		tambon, amphoe, province := extractAddressComponents(doc.BuyerLineTwo)

		// find province
		provinceData, err := s.repo.GetProvince(repositories.ThaiCityFilter{
			ProvinceNameTh: province,
		})

		if err != nil {
			return err
		}

		if len(provinceData) == 0 {
			return fmt.Errorf("buyer province not found %s", province)
		}

		// find amphure
		// log.Println(amphoe, provinceData[0].ID)
		amphoeData, err := s.repo.GetAmphure(repositories.ThaiCityFilter{
			AmphureNameTh: amphoe,
			ProvinceID:    provinceData[0].ID,
		})

		if err != nil {
			return err
		}

		if len(amphoeData) == 0 {
			return fmt.Errorf("amphure not found")
		}

		// find tambon
		tambonData, err := s.repo.GetTambon(repositories.ThaiCityFilter{
			TambonNameTh: tambon,
			AmphureID:    amphoeData[0].ID,
		})

		if err != nil {
			return err
		}

		if len(tambonData) == 0 {
			return fmt.Errorf("tambon not found")
		}

		if doc.BuyerCityCode == "" {

			doc.BuyerCityCode = fmt.Sprintf("%d", amphoeData[0].ID)
			doc.BuyerCitySubDivisionCode = fmt.Sprintf("%d", tambonData[0].ID)
			doc.BuyerPostcodeCode = fmt.Sprintf("%d", tambonData[0].ZipCode)

		}

		buyerCountrySubDivisionID = fmt.Sprintf("%d", provinceData[0].ID)

		if buyerCountrySubDivisionID == "" {
			buyerCountrySubDivisionID = "10"
		}
	}

	// {
	// 	tambon, amphoe, province := extractAddressComponents(doc.SellerLineTwo)

	// 	// find province
	// 	provinceData, err := s.repo.GetProvince(repositories.ThaiCityFilter{
	// 		ProvinceNameTh: province,
	// 	})

	// 	if err != nil {
	// 		return err
	// 	}

	// 	if len(provinceData) == 0 {
	// 		return fmt.Errorf("seller province not found %s", province)
	// 	}

	// 	// find amphure
	// 	amphoeData, err := s.repo.GetAmphure(repositories.ThaiCityFilter{
	// 		AmphureNameTh: amphoe,
	// 		ProvinceID:    provinceData[0].ID,
	// 	})

	// 	if err != nil {
	// 		return err
	// 	}

	// 	if len(amphoeData) == 0 {
	// 		return fmt.Errorf("amphure not found")
	// 	}

	// 	// find tambon
	// 	tambonData, err := s.repo.GetTambon(repositories.ThaiCityFilter{
	// 		TambonNameTh: tambon,
	// 		AmphureID:    amphoeData[0].ID,
	// 	})

	// 	if err != nil {
	// 		return err
	// 	}

	// 	if len(tambonData) == 0 {
	// 		return fmt.Errorf("tambon not found")
	// 	}

	// 	if doc.SellerCityCode == "" {
	// 		doc.SellerCityCode = fmt.Sprintf("%d", amphoeData[0].ID)
	// 		doc.SellerCitySubDivisionCode = fmt.Sprintf("%d", tambonData[0].ID)
	// 		doc.SellerPostcodeCode = fmt.Sprintf("%d", tambonData[0].ZipCode)
	// 	}

	// 	sellerCountrySubDivisionID = fmt.Sprintf("%d", provinceData[0].ID)

	// 	if sellerCountrySubDivisionID == "" {
	// 		sellerCountrySubDivisionID = "10"
	// 	}
	// }

	for _, item := range doc.DocumentLines {
		tmpItem = append(tmpItem, modelsxml.SupplyChainTradeLineItem{
			AssociatedDocumentLineDocument: modelsxml.AssociatedDocumentLineDocument{
				LineID: item.LineID,
			},
			SpecifiedTradeProduct: modelsxml.SpecifiedTradeProduct{
				Name: item.ProductName,
			},
			SpecifiedLineTradeAgreement: modelsxml.SpecifiedLineTradeAgreement{
				GrossPriceProductTradePrice: modelsxml.GrossPriceProductTradePrice{
					ChargeAmount: fmt.Sprintf("%.5f", item.ChargeAmount),
				},
			},
			SpecifiedLineTradeDelivery: modelsxml.SpecifiedLineTradeDelivery{
				BilledQuantity: modelsxml.BilledQuantity{
					UnitCode: item.UnitName,
					Value:    item.BilledQuantity, // เปลี่ยนจาก Amount เป็น Value ตาม struct ใหม่
				},
			},
			SpecifiedLineTradeSettlement: modelsxml.SpecifiedLineTradeSettlement{
				ApplicableTradeTax: modelsxml.ApplicableTradeTax{
					TypeCode:         "VAT",
					CalculatedRate:   fmt.Sprintf("%.5f", doc.TaxCalculatedRate),
					BasisAmount:      fmt.Sprintf("%.5f", item.NetAmountDue),
					CalculatedAmount: fmt.Sprintf("%.5f", item.TaxTotalAmount),
				},
				SpecifiedTradeSettlementLineMonetarySummation: modelsxml.SpecifiedTradeSettlementLineMonetarySummation{
					TaxTotalAmount: fmt.Sprintf("%.5f", item.TaxTotalAmount),
					NetLineTotalAmount: modelsxml.MonetaryAmount{
						Value:      fmt.Sprintf("%.5f", item.NetAmountDue),
						CurrencyID: "THB",
					},
					NetIncludingTaxesLineTotalAmount: modelsxml.MonetaryAmount{
						Value:      fmt.Sprintf("%.5f", item.NetAmountDue+item.TaxTotalAmount),
						CurrencyID: "THB",
					},
				},
			},
		})
	}

	// สร้าง XML Document
	xmlDoc := modelsxml.Document{
		XMLName: xml.Name{
			Local: "TaxInvoice_CrossIndustryInvoice",
			Space: "rsm",
		},
		// เพิ่ม XML namespace attributes
		XmlnsRam:       "urn:etda:uncefact:data:standard:TaxInvoice_ReusableAggregateBusinessInformationEntity:2",
		XmlnsRsm:       "urn:etda:uncefact:data:standard:TaxInvoice_CrossIndustryInvoice:2",
		XmlnsXsi:       "http://www.w3.org/2001/XMLSchema-instance",
		SchemaLocation: "urn:etda:uncefact:data:standard:TaxInvoice_CrossIndustryInvoice:2",

		// เพิ่ม ExchangedDocumentContext
		ExchangedDocumentContext: modelsxml.ExchangedDocumentContext{
			GuidelineSpecifiedDocumentContextParameter: modelsxml.GuidelineSpecifiedDocumentContextParameter{
				ID: modelsxml.ID{
					Value:           "ER3-2560",
					SchemeAgencyID:  "ETDA",
					SchemeVersionID: "v2.0",
				},
			},
		},

		// ExchangedDocument
		ExchangedDocument: modelsxml.ExchangedDocument{
			ID:               doc.DocumentID,
			Name:             "ใบเสร็จรับเงิน/ใบกำกับภาษี",
			TypeCode:         doc.DocumentTypeCode,
			IssueDateTime:    doc.IssueDateTime,
			CreationDateTime: doc.CreationDateTime,
			IncludedNote: modelsxml.IncludedNote{
				Subject: "หมายเหตุ",
				Content: "บันทึกข้อความ: ",
			},
		},

		SupplyChainTradeTransaction: modelsxml.SupplyChainTradeTransaction{
			ApplicableHeaderTradeAgreement: modelsxml.ApplicableHeaderTradeAgreement{
				SellerTradeParty: modelsxml.TradeParty{
					Name: doc.SellerName,
					SpecifiedTaxRegistration: modelsxml.TaxRegistration{
						ID: modelsxml.ID{
							Value:    doc.SellerTaxID,
							SchemeID: doc.SellerTaxType,
						},
					},
					DefinedTradeContact: modelsxml.DefinedTradeContact{
						PersonName: doc.SellerName,
						EmailURIUniversalCommunication: modelsxml.EmailURI{
							URIID: doc.SellerEmail,
						},
						TelephoneUniversalCommunication: modelsxml.Telephone{
							CompleteNumber: doc.SellerPhoneNumber,
						},
					},
					PostalTradeAddress: modelsxml.PostalTradeAddress{
						PostcodeCode:         doc.SellerPostcodeCode,
						LineOne:              doc.SellerLineOne,
						CityName:             doc.SellerCityCode,
						CitySubDivisionName:  doc.SellerCitySubDivisionCode,
						CountryID:            "TH",
						CountrySubDivisionID: doc.SellerCountrySubDivisionCode,
						BuildingNumber:       doc.SellerBuildingNumber,
					},
				},
				BuyerTradeParty: modelsxml.TradeParty{
					Name: doc.BuyerName,
					SpecifiedTaxRegistration: modelsxml.TaxRegistration{
						ID: modelsxml.ID{
							Value:    doc.BuyerTaxID,
							SchemeID: doc.BuyerTaxType,
						},
					},
					DefinedTradeContact: modelsxml.DefinedTradeContact{
						PersonName: doc.BuyerName,
						EmailURIUniversalCommunication: modelsxml.EmailURI{
							URIID: doc.BuyerEmail,
						},
						TelephoneUniversalCommunication: modelsxml.Telephone{
							CompleteNumber: doc.BuyerPhoneNumber,
						},
					},
					PostalTradeAddress: modelsxml.PostalTradeAddress{
						PostcodeCode:         doc.BuyerPostcodeCode,
						LineOne:              doc.BuyerLineOne,
						CityName:             doc.BuyerCityCode,
						CitySubDivisionName:  doc.BuyerCitySubDivisionCode,
						CountryID:            "TH",
						CountrySubDivisionID: buyerCountrySubDivisionID,
						BuildingNumber:       buyerBuildingNumber,
					},
				},
			},
			ApplicableHeaderTradeDelivery: modelsxml.ApplicableHeaderTradeDelivery{
				ShipToTradeParty: modelsxml.TradeParty{
					Name: doc.BuyerName,
					SpecifiedTaxRegistration: modelsxml.TaxRegistration{
						ID: modelsxml.ID{
							Value:    doc.BuyerTaxID,
							SchemeID: doc.BuyerTaxType,
						},
					},
					DefinedTradeContact: modelsxml.DefinedTradeContact{
						PersonName: doc.BuyerName,
						EmailURIUniversalCommunication: modelsxml.EmailURI{
							URIID: doc.BuyerEmail,
						},
						TelephoneUniversalCommunication: modelsxml.Telephone{
							CompleteNumber: doc.BuyerPhoneNumber,
						},
					},
					PostalTradeAddress: modelsxml.PostalTradeAddress{
						PostcodeCode:         doc.BuyerPostcodeCode,
						LineOne:              doc.BuyerLineOne,
						CityName:             doc.BuyerCityCode,
						CitySubDivisionName:  doc.BuyerCitySubDivisionCode,
						CountryID:            "TH",
						CountrySubDivisionID: buyerCountrySubDivisionID,
						BuildingNumber:       buyerBuildingNumber,
					},
				},
				ShipFromTradeParty: modelsxml.TradeParty{
					Name: doc.SellerName,
					SpecifiedTaxRegistration: modelsxml.TaxRegistration{
						ID: modelsxml.ID{
							Value:    doc.SellerTaxID,
							SchemeID: doc.SellerTaxType,
						},
					},
					DefinedTradeContact: modelsxml.DefinedTradeContact{
						PersonName: doc.SellerName,
						EmailURIUniversalCommunication: modelsxml.EmailURI{
							URIID: doc.SellerEmail,
						},
						TelephoneUniversalCommunication: modelsxml.Telephone{
							CompleteNumber: doc.SellerPhoneNumber,
						},
					},
					PostalTradeAddress: modelsxml.PostalTradeAddress{
						PostcodeCode:         doc.SellerPostcodeCode,
						LineOne:              doc.SellerLineOne,
						CityName:             doc.SellerCityCode,
						CitySubDivisionName:  doc.SellerCitySubDivisionCode,
						CountryID:            "TH",
						CountrySubDivisionID: doc.SellerCountrySubDivisionCode,
						BuildingNumber:       doc.SellerBuildingNumber,
					},
				},
			},
			ApplicableHeaderTradeSettlement: modelsxml.ApplicableHeaderTradeSettlement{
				InvoiceCurrencyCode: modelsxml.InvoiceCurrencyCode{
					Value:  "THB",
					ListID: "ISO 4217 3A",
				},
				ApplicableTradeTax: modelsxml.ApplicableTradeTax{
					TypeCode:         "VAT",
					CalculatedRate:   fmt.Sprintf("%.2f", doc.TaxCalculatedRate),
					BasisAmount:      fmt.Sprintf("%.5f", doc.TaxBasisTotalAmount),
					CalculatedAmount: fmt.Sprintf("%.5f", doc.TaxTotalAmount),
				},
				SpecifiedTradePaymentTerms: modelsxml.SpecifiedTradePaymentTerms{
					Description:     "ชำระภายใน 0",
					DueDateDateTime: doc.DueDateTime,
					TypeCode:        "1",
				},
				SpecifiedTradeSettlementHeaderMonetarySummation: modelsxml.SpecifiedTradeSettlementHeaderMonetarySummation{
					LineTotalAmount:     fmt.Sprintf("%.5f", doc.LineTotalAmount),
					TaxBasisTotalAmount: fmt.Sprintf("%.5f", doc.TaxBasisTotalAmount),
					TaxTotalAmount:      fmt.Sprintf("%.5f", doc.TaxTotalAmount),
					GrandTotalAmount:    fmt.Sprintf("%.5f", doc.GrandTotalAmount),
				},
			},
			IncludedSupplyChainTradeLineItem: tmpItem,
		},
	}

	// แปลงเป็น XML
	xmlData, err := xml.MarshalIndent(xmlDoc, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return err
	}

	// เพิ่ม XML header
	xmlHeader := []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
	xmlData = append(xmlHeader, xmlData...)

	xmlName := fmt.Sprintf("%d_", doc.Id) + doc.DocumentID + ".xml"

	// สร้างโฟลเดอร์ถ้ายังไม่มี
	folderPath := "./xml_" + doc.Customer
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return err
		}
	}

	// เขียน XML ลงในไฟล์
	err = os.WriteFile(folderPath+"/"+xmlName, xmlData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}

	fmt.Println("XML data has been written to " + xmlName)
	return nil
}

func extractAddressComponents(address string) (string, string, string) {
	// สร้าง regex สำหรับจับ "ตำบล/แขวง", "อำเภอ/เขต", และ "จังหวัด"
	subdistrictRegex := regexp.MustCompile(`(ตำบล|แขวง)\s+([\p{L}\p{M}]+)`)
	districtRegex := regexp.MustCompile(`(อำเภอ|เขต)\s+([\p{L}\p{M}]+)`)

	// แยกที่อยู่เป็นคำเพื่อหาจังหวัด (คำสุดท้าย)
	words := strings.Fields(address)
	province := "ไม่พบข้อมูล"
	if len(words) > 0 {
		province = words[len(words)-1] // คำสุดท้ายของที่อยู่เป็นจังหวัด
	}

	// ค้นหาข้อมูลตำบล/แขวง และ อำเภอ/เขต
	subdistrictMatch := subdistrictRegex.FindStringSubmatch(address)
	districtMatch := districtRegex.FindStringSubmatch(address)

	// กำหนดค่าเริ่มต้น
	subdistrict := "ไม่พบข้อมูล"
	district := "ไม่พบข้อมูล"

	// ตรวจสอบผลลัพธ์
	if len(subdistrictMatch) > 2 {
		subdistrict = subdistrictMatch[2]
	}
	if len(districtMatch) > 2 {
		district = districtMatch[2]
	}

	return subdistrict, district, province
}

func extractHouseNumber(text string) string {
	// ใช้ Regex เพื่อจับ "เลขที่" ตามด้วยตัวเลข อาจมี "/" ตามด้วยตัวเลขอีกชุด
	houseNumberRegex := regexp.MustCompile(`เลขที่\s+(\d+(?:/\d+)?)`)

	match := houseNumberRegex.FindStringSubmatch(text)
	if len(match) > 1 {
		return match[1] // match[1] คือเลขที่บ้านที่ต้องการ
	}
	return "" // ถ้าไม่มีให้คืนค่าเป็นค่าว่าง
}
