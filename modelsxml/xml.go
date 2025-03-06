package modelsxml

import (
	"encoding/xml"
)

// กำหนด namespace ต่างๆ
const (
	NamespaceRam = "urn:etda:uncefact:data:standard:TaxInvoice_ReusableAggregateBusinessInformationEntity:2"
	NamespaceRsm = "urn:etda:uncefact:data:standard:TaxInvoice_CrossIndustryInvoice:2"
	NamespaceXsi = "http://www.w3.org/2001/XMLSchema-instance"
)

type Document struct {
	XMLName                     xml.Name                    `xml:"rsm:TaxInvoice_CrossIndustryInvoice"`
	XmlnsRam                    string                      `xml:"xmlns:ram,attr,omitempty"`
	XmlnsRsm                    string                      `xml:"xmlns:rsm,attr,omitempty"`
	XmlnsXsi                    string                      `xml:"xmlns:xsi,attr,omitempty"`
	SchemaLocation              string                      `xml:"xsi:schemaLocation,attr,omitempty"`
	ExchangedDocumentContext    ExchangedDocumentContext    `xml:"rsm:ExchangedDocumentContext"`
	ExchangedDocument           ExchangedDocument           `xml:"rsm:ExchangedDocument"`
	SupplyChainTradeTransaction SupplyChainTradeTransaction `xml:"rsm:SupplyChainTradeTransaction"`
}

type ExchangedDocumentContext struct {
	GuidelineSpecifiedDocumentContextParameter GuidelineSpecifiedDocumentContextParameter `xml:"ram:GuidelineSpecifiedDocumentContextParameter"`
}

type GuidelineSpecifiedDocumentContextParameter struct {
	ID ID `xml:"ram:ID"`
}

type ID struct {
	Value           string `xml:",chardata"`
	SchemeAgencyID  string `xml:"schemeAgencyID,attr,omitempty"`
	SchemeVersionID string `xml:"schemeVersionID,attr,omitempty"`
	SchemeID        string `xml:"schemeID,attr,omitempty"`
}

type ExchangedDocument struct {
	ID               string       `xml:"ram:ID"`
	Name             string       `xml:"ram:Name"`
	TypeCode         string       `xml:"ram:TypeCode"`
	IssueDateTime    string       `xml:"ram:IssueDateTime"`
	CreationDateTime string       `xml:"ram:CreationDateTime"`
	IncludedNote     IncludedNote `xml:"ram:IncludedNote"`
}

type IncludedNote struct {
	Subject string `xml:"ram:Subject"`
	Content string `xml:"ram:Content"`
}

type SupplyChainTradeTransaction struct {
	ApplicableHeaderTradeAgreement   ApplicableHeaderTradeAgreement  `xml:"ram:ApplicableHeaderTradeAgreement"`
	ApplicableHeaderTradeDelivery    ApplicableHeaderTradeDelivery   `xml:"ram:ApplicableHeaderTradeDelivery"`
	ApplicableHeaderTradeSettlement  ApplicableHeaderTradeSettlement `xml:"ram:ApplicableHeaderTradeSettlement"`
	IncludedSupplyChainTradeLineItem []SupplyChainTradeLineItem      `xml:"ram:IncludedSupplyChainTradeLineItem,omitempty"`
}

type ApplicableHeaderTradeAgreement struct {
	SellerTradeParty TradeParty `xml:"ram:SellerTradeParty"`
	BuyerTradeParty  TradeParty `xml:"ram:BuyerTradeParty"`
}

type TradeParty struct {
	Name                     string              `xml:"ram:Name"`
	SpecifiedTaxRegistration TaxRegistration     `xml:"ram:SpecifiedTaxRegistration"`
	DefinedTradeContact      DefinedTradeContact `xml:"ram:DefinedTradeContact"`
	PostalTradeAddress       PostalTradeAddress  `xml:"ram:PostalTradeAddress"`
}

type TaxRegistration struct {
	ID ID `xml:"ram:ID"`
}

type DefinedTradeContact struct {
	PersonName                      string    `xml:"ram:PersonName"`
	EmailURIUniversalCommunication  EmailURI  `xml:"ram:EmailURIUniversalCommunication"`
	TelephoneUniversalCommunication Telephone `xml:"ram:TelephoneUniversalCommunication"`
}

type EmailURI struct {
	URIID string `xml:"ram:URIID"`
}

type Telephone struct {
	CompleteNumber string `xml:"ram:CompleteNumber"`
}

type PostalTradeAddress struct {
	PostcodeCode         string `xml:"ram:PostcodeCode"`
	LineOne              string `xml:"ram:LineOne"`
	CityName             string `xml:"ram:CityName"`
	CitySubDivisionName  string `xml:"ram:CitySubDivisionName"`
	CountryID            string `xml:"ram:CountryID"`
	CountrySubDivisionID string `xml:"ram:CountrySubDivisionID"`
	BuildingNumber       string `xml:"ram:BuildingNumber"`
}

type ApplicableHeaderTradeDelivery struct {
	ShipToTradeParty   TradeParty `xml:"ram:ShipToTradeParty"`
	ShipFromTradeParty TradeParty `xml:"ram:ShipFromTradeParty"`
}

type ApplicableHeaderTradeSettlement struct {
	InvoiceCurrencyCode                             InvoiceCurrencyCode                             `xml:"ram:InvoiceCurrencyCode"`
	ApplicableTradeTax                              ApplicableTradeTax                              `xml:"ram:ApplicableTradeTax"`
	SpecifiedTradePaymentTerms                      SpecifiedTradePaymentTerms                      `xml:"ram:SpecifiedTradePaymentTerms"`
	SpecifiedTradeSettlementHeaderMonetarySummation SpecifiedTradeSettlementHeaderMonetarySummation `xml:"ram:SpecifiedTradeSettlementHeaderMonetarySummation"`
}

type InvoiceCurrencyCode struct {
	Value  string `xml:",chardata"`
	ListID string `xml:"listID,attr,omitempty"`
}

type ApplicableTradeTax struct {
	TypeCode         string `xml:"ram:TypeCode"`
	CalculatedRate   string `xml:"ram:CalculatedRate"`
	BasisAmount      string `xml:"ram:BasisAmount"`
	CalculatedAmount string `xml:"ram:CalculatedAmount"`
}

type SpecifiedTradePaymentTerms struct {
	Description     string `xml:"ram:Description"`
	DueDateDateTime string `xml:"ram:DueDateDateTime"`
	TypeCode        string `xml:"ram:TypeCode"`
}

type SpecifiedTradeSettlementHeaderMonetarySummation struct {
	LineTotalAmount     string `xml:"ram:LineTotalAmount"`
	TaxBasisTotalAmount string `xml:"ram:TaxBasisTotalAmount"`
	TaxTotalAmount      string `xml:"ram:TaxTotalAmount"`
	GrandTotalAmount    string `xml:"ram:GrandTotalAmount"`
}

type SupplyChainTradeLineItem struct {
	AssociatedDocumentLineDocument AssociatedDocumentLineDocument `xml:"ram:AssociatedDocumentLineDocument"`
	SpecifiedTradeProduct          SpecifiedTradeProduct          `xml:"ram:SpecifiedTradeProduct"`
	SpecifiedLineTradeAgreement    SpecifiedLineTradeAgreement    `xml:"ram:SpecifiedLineTradeAgreement"`
	SpecifiedLineTradeDelivery     SpecifiedLineTradeDelivery     `xml:"ram:SpecifiedLineTradeDelivery"`
	SpecifiedLineTradeSettlement   SpecifiedLineTradeSettlement   `xml:"ram:SpecifiedLineTradeSettlement"`
}

type AssociatedDocumentLineDocument struct {
	LineID int `xml:"ram:LineID"`
}

type SpecifiedTradeProduct struct {
	Name string `xml:"ram:Name"`
}

type SpecifiedLineTradeAgreement struct {
	GrossPriceProductTradePrice GrossPriceProductTradePrice `xml:"ram:GrossPriceProductTradePrice"`
}

type GrossPriceProductTradePrice struct {
	ChargeAmount string `xml:"ram:ChargeAmount"`
}

type SpecifiedLineTradeDelivery struct {
	BilledQuantity BilledQuantity `xml:"ram:BilledQuantity"`
}

type BilledQuantity struct {
	Value    int    `xml:",chardata"`
	UnitCode string `xml:"unitCode,attr,omitempty"`
}

type SpecifiedLineTradeSettlement struct {
	ApplicableTradeTax                            ApplicableTradeTax                            `xml:"ram:ApplicableTradeTax"`
	SpecifiedTradeSettlementLineMonetarySummation SpecifiedTradeSettlementLineMonetarySummation `xml:"ram:SpecifiedTradeSettlementLineMonetarySummation"`
}

type SpecifiedTradeSettlementLineMonetarySummation struct {
	TaxTotalAmount                   string         `xml:"ram:TaxTotalAmount"`
	NetLineTotalAmount               MonetaryAmount `xml:"ram:NetLineTotalAmount"`
	NetIncludingTaxesLineTotalAmount MonetaryAmount `xml:"ram:NetIncludingTaxesLineTotalAmount"`
}

type MonetaryAmount struct {
	Value      string `xml:",chardata"`
	CurrencyID string `xml:"currencyID,attr,omitempty"`
}
