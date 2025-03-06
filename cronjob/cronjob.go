package cronjob

// import (
// 	"bmsp-backend-service/db"
// 	"bmsp-backend-service/models"
// 	"bmsp-backend-service/modelsxml"
// 	"encoding/json"
// 	"encoding/xml"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// 	"time"
// )

// func StartCronjobBuzzebeeJson() {
// 	log.Println("StartCronjobBuzzebeeJson")
// 	for {

// 		// load list file from horizon
// 		files, err := os.ReadDir("./buzzebee")
// 		if err != nil {

// 		}

// 		for _, file := range files {

// 			// check if file is json
// 			if file.IsDir() {
// 				continue
// 			}

// 			if strings.Contains(file.Name(), ".json") == false {
// 				continue
// 			}

// 			jsonData, err := os.ReadFile("./buzzebee/" + file.Name())
// 			if err != nil {

// 				log.Println(err.Error())
// 				continue
// 			}

// 			// แปลง JSON เป็น struct
// 			var doc models.Document
// 			err = json.Unmarshal([]byte(jsonData), &doc)
// 			if err != nil {
// 				log.Println(err.Error())
// 				continue
// 			}

// 			// แปลง struct เป็น JSON
// 			jsonData, err = json.Marshal(doc)
// 			if err != nil {
// 				log.Println(err.Error())
// 				continue
// 			}

// 			doc.Customer = "buzzebee"
// 			doc.Status = "waiting"

// 			err = db.DBPg.Create(&doc).Error

// 			if err == nil {
// 				os.Remove("./buzzebee/" + file.Name())
// 			}

// 		}

// 		time.Sleep(5 * time.Second)
// 	}
// }

// func StartCronjobBuzzebeeJsonToXML() {
// 	log.Println("StartCronjobBuzzebeeJsonToXML")

// 	for {
// 		// find all document waiting
// 		var docs []models.Document
// 		if err := db.DBPg.Where("status = ?", "waiting").Preload("DocumentLines").Find(&docs).Error; err != nil {
// 			log.Println(err.Error())
// 			continue
// 		}

// 		// convert json to xml
// 		for _, doc := range docs {

// 			log.Printf("%#+v\n", doc)

// 			tmpItem := []modelsxml.SupplyChainTradeLineItem{}

// 			for _, item := range doc.DocumentLines {
// 				tmpItem = append(tmpItem, modelsxml.SupplyChainTradeLineItem{
// 					AssociatedDocumentLineDocument: modelsxml.AssociatedDocumentLineDocument{
// 						LineID: item.LineID,
// 					},
// 					SpecifiedTradeProduct: modelsxml.SpecifiedTradeProduct{
// 						Name: item.ProductName,
// 					},
// 					SpecifiedLineTradeAgreement: modelsxml.SpecifiedLineTradeAgreement{
// 						GrossPriceProductTradePrice: modelsxml.GrossPriceProductTradePrice{
// 							ChargeAmount: item.ChargeAmount,
// 						},
// 					},
// 					SpecifiedLineTradeDelivery: modelsxml.SpecifiedLineTradeDelivery{
// 						BilledQuantity: modelsxml.BilledQuantity{
// 							UnitCode: item.UnitName,
// 							Value:    item.BilledQuantity, // เปลี่ยนจาก Amount เป็น Value ตาม struct ใหม่
// 						},
// 					},
// 					SpecifiedLineTradeSettlement: modelsxml.SpecifiedLineTradeSettlement{
// 						ApplicableTradeTax: modelsxml.ApplicableTradeTax{
// 							TypeCode:         "VAT",
// 							CalculatedRate:   7.00,
// 							BasisAmount:      item.NetAmountDue,
// 							CalculatedAmount: item.TaxTotalAmount,
// 						},
// 						SpecifiedTradeSettlementLineMonetarySummation: modelsxml.SpecifiedTradeSettlementLineMonetarySummation{
// 							TaxTotalAmount: item.TaxTotalAmount,
// 							NetLineTotalAmount: modelsxml.MonetaryAmount{
// 								Value:      item.NetAmountDue,
// 								CurrencyID: "THB",
// 							},
// 							NetIncludingTaxesLineTotalAmount: modelsxml.MonetaryAmount{
// 								Value:      item.NetAmountDue + item.TaxTotalAmount,
// 								CurrencyID: "THB",
// 							},
// 						},
// 					},
// 				})
// 			}

// 			// สร้าง XML Document
// 			xmlDoc := modelsxml.Document{
// 				XMLName: xml.Name{
// 					Local: "TaxInvoice_CrossIndustryInvoice",
// 					Space: "rsm",
// 				},
// 				// เพิ่ม XML namespace attributes
// 				XmlnsRam:       "urn:etda:uncefact:data:standard:TaxInvoice_ReusableAggregateBusinessInformationEntity:2",
// 				XmlnsRsm:       "urn:etda:uncefact:data:standard:TaxInvoice_CrossIndustryInvoice:2",
// 				XmlnsXsi:       "http://www.w3.org/2001/XMLSchema-instance",
// 				SchemaLocation: "urn:etda:uncefact:data:standard:TaxInvoice_CrossIndustryInvoice:2",

// 				// เพิ่ม ExchangedDocumentContext
// 				ExchangedDocumentContext: modelsxml.ExchangedDocumentContext{
// 					GuidelineSpecifiedDocumentContextParameter: modelsxml.GuidelineSpecifiedDocumentContextParameter{
// 						ID: modelsxml.ID{
// 							Value:           "ER3-2560",
// 							SchemeAgencyID:  "ETDA",
// 							SchemeVersionID: "v2.0",
// 						},
// 					},
// 				},

// 				// ExchangedDocument
// 				ExchangedDocument: modelsxml.ExchangedDocument{
// 					ID:               doc.DocumentID,
// 					Name:             "ใบเสร็จรับเงิน/ใบกำกับภาษี",
// 					TypeCode:         doc.DocumentTypeCode,
// 					IssueDateTime:    doc.IssueDateTime,
// 					CreationDateTime: doc.CreationDateTime,
// 					IncludedNote: modelsxml.IncludedNote{
// 						Subject: "หมายเหตุ",
// 						Content: "บันทึกข้อความ: ",
// 					},
// 				},

// 				SupplyChainTradeTransaction: modelsxml.SupplyChainTradeTransaction{
// 					ApplicableHeaderTradeAgreement: modelsxml.ApplicableHeaderTradeAgreement{
// 						SellerTradeParty: modelsxml.TradeParty{
// 							Name: doc.SellerName,
// 							SpecifiedTaxRegistration: modelsxml.TaxRegistration{
// 								ID: modelsxml.ID{
// 									Value:    doc.SellerTaxID,
// 									SchemeID: "TXID",
// 								},
// 							},
// 							DefinedTradeContact: modelsxml.DefinedTradeContact{
// 								PersonName: doc.SellerName,
// 								EmailURIUniversalCommunication: modelsxml.EmailURI{
// 									URIID: doc.SellerEmail,
// 								},
// 								TelephoneUniversalCommunication: modelsxml.Telephone{
// 									CompleteNumber: doc.SellerPhoneNumber,
// 								},
// 							},
// 							PostalTradeAddress: modelsxml.PostalTradeAddress{
// 								PostcodeCode:         doc.SellerPostcodeCode,
// 								LineOne:              doc.SellerLineOne,
// 								CityName:             doc.SellerCityName,
// 								CitySubDivisionName:  doc.SellerCitySubDivisionName,
// 								CountryID:            "TH",
// 								CountrySubDivisionID: "10",
// 								BuildingNumber:       doc.SellerBuildingNumber,
// 							},
// 						},
// 						BuyerTradeParty: modelsxml.TradeParty{
// 							Name: doc.BuyerName,
// 							SpecifiedTaxRegistration: modelsxml.TaxRegistration{
// 								ID: modelsxml.ID{
// 									Value:    doc.BuyerTaxID,
// 									SchemeID: "TXID",
// 								},
// 							},
// 							DefinedTradeContact: modelsxml.DefinedTradeContact{
// 								PersonName: doc.BuyerName,
// 								EmailURIUniversalCommunication: modelsxml.EmailURI{
// 									URIID: doc.BuyerEmail,
// 								},
// 								TelephoneUniversalCommunication: modelsxml.Telephone{
// 									CompleteNumber: doc.BuyerPhoneNumber,
// 								},
// 							},
// 							PostalTradeAddress: modelsxml.PostalTradeAddress{
// 								PostcodeCode:         doc.BuyerPostcodeCode,
// 								LineOne:              doc.BuyerLineOne,
// 								CityName:             doc.BuyerCityName,
// 								CitySubDivisionName:  doc.BuyerCitySubDivisionName,
// 								CountryID:            "TH",
// 								CountrySubDivisionID: "10",
// 								BuildingNumber:       doc.BuyerBuildingNumber,
// 							},
// 						},
// 					},
// 					ApplicableHeaderTradeDelivery: modelsxml.ApplicableHeaderTradeDelivery{
// 						ShipToTradeParty: modelsxml.TradeParty{
// 							Name: doc.BuyerName,
// 							SpecifiedTaxRegistration: modelsxml.TaxRegistration{
// 								ID: modelsxml.ID{
// 									Value:    doc.BuyerTaxID,
// 									SchemeID: "TXID",
// 								},
// 							},
// 							DefinedTradeContact: modelsxml.DefinedTradeContact{
// 								PersonName: doc.BuyerName,
// 								EmailURIUniversalCommunication: modelsxml.EmailURI{
// 									URIID: doc.BuyerEmail,
// 								},
// 								TelephoneUniversalCommunication: modelsxml.Telephone{
// 									CompleteNumber: doc.BuyerPhoneNumber,
// 								},
// 							},
// 							PostalTradeAddress: modelsxml.PostalTradeAddress{
// 								PostcodeCode:         doc.BuyerPostcodeCode,
// 								LineOne:              doc.BuyerLineOne,
// 								CityName:             doc.BuyerCityName,
// 								CitySubDivisionName:  doc.BuyerCitySubDivisionName,
// 								CountryID:            "TH",
// 								CountrySubDivisionID: "10",
// 								BuildingNumber:       doc.BuyerBuildingNumber,
// 							},
// 						},
// 						ShipFromTradeParty: modelsxml.TradeParty{
// 							Name: doc.SellerName,
// 							SpecifiedTaxRegistration: modelsxml.TaxRegistration{
// 								ID: modelsxml.ID{
// 									Value:    doc.SellerTaxID,
// 									SchemeID: "TXID",
// 								},
// 							},
// 							DefinedTradeContact: modelsxml.DefinedTradeContact{
// 								PersonName: doc.SellerName,
// 								EmailURIUniversalCommunication: modelsxml.EmailURI{
// 									URIID: doc.SellerEmail,
// 								},
// 								TelephoneUniversalCommunication: modelsxml.Telephone{
// 									CompleteNumber: doc.SellerPhoneNumber,
// 								},
// 							},
// 							PostalTradeAddress: modelsxml.PostalTradeAddress{
// 								PostcodeCode:         doc.SellerPostcodeCode,
// 								LineOne:              doc.SellerLineOne,
// 								CityName:             doc.SellerCityName,
// 								CitySubDivisionName:  doc.SellerCitySubDivisionName,
// 								CountryID:            "TH",
// 								CountrySubDivisionID: "10",
// 								BuildingNumber:       doc.SellerBuildingNumber,
// 							},
// 						},
// 					},
// 					ApplicableHeaderTradeSettlement: modelsxml.ApplicableHeaderTradeSettlement{
// 						InvoiceCurrencyCode: modelsxml.InvoiceCurrencyCode{
// 							Value:  "THB",
// 							ListID: "ISO 4217 3A",
// 						},
// 						ApplicableTradeTax: modelsxml.ApplicableTradeTax{
// 							TypeCode:         "VAT",
// 							CalculatedRate:   7.00,
// 							BasisAmount:      doc.TaxBasisTotalAmount,
// 							CalculatedAmount: doc.TaxTotalAmount,
// 						},
// 						SpecifiedTradePaymentTerms: modelsxml.SpecifiedTradePaymentTerms{
// 							Description:     "ชำระภายใน 0",
// 							DueDateDateTime: doc.DueDateTime,
// 							TypeCode:        "1",
// 						},
// 						SpecifiedTradeSettlementHeaderMonetarySummation: modelsxml.SpecifiedTradeSettlementHeaderMonetarySummation{
// 							LineTotalAmount:     doc.LineTotalAmount,
// 							TaxBasisTotalAmount: doc.TaxBasisTotalAmount,
// 							TaxTotalAmount:      doc.TaxTotalAmount,
// 							GrandTotalAmount:    doc.GrandTotalAmount,
// 						},
// 					},
// 					IncludedSupplyChainTradeLineItem: tmpItem,
// 				},
// 			}

// 			// แปลงเป็น XML
// 			xmlData, err := xml.MarshalIndent(xmlDoc, "", "  ")
// 			if err != nil {
// 				fmt.Println("Error marshalling:", err)
// 				continue
// 			}

// 			// เพิ่ม XML header
// 			xmlHeader := []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
// 			xmlData = append(xmlHeader, xmlData...)

// 			xmlName := fmt.Sprintf("%d_", doc.Id) + doc.DocumentID + ".xml"

// 			// เขียน XML ลงในไฟล์
// 			err = os.WriteFile("./xml_buzzebee/"+xmlName, xmlData, 0644)
// 			if err != nil {
// 				fmt.Println("Error writing file:", err)
// 				continue
// 			}

// 			fmt.Println("XML data has been written to " + xmlName)

// 			// update status to sent
// 			db.DBPg.Model(&models.Document{}).Where("id = ?", doc.Id).Update("status", "sent")

// 		}

// 		time.Sleep(5 * time.Second)
// 	}

// }
