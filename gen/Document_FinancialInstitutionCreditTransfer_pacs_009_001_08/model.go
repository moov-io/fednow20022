// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:fednow
package Document_FinancialInstitutionCreditTransfer_pacs_009_001_08

import (
	"encoding/xml"

	"github.com/moov-io/fednow20022/gen/FinancialInstitutionCreditTransfer_pacs_009_001_08"
	"github.com/moov-io/fednow20022/gen/head_001_001_02"
)

// XSD Elements

type Message struct {
	XMLName xml.Name   `xml:"Message"`
	Xmlns   []xml.Attr `xml:",attr"`

	AppHdr head_001_001_02.BusinessApplicationHeaderV02 `xml:"urn:fednow AppHdr"`

	FinancialInstitutionCreditTransfer *FinancialInstitutionCreditTransfer_pacs_009_001_08.Document `xml:"urn:fednow FinancialInstitutionCreditTransfer,omitempty"`
}

var NamespacePrefixMap = map[string]string{
	"urn:iso:std:iso:20022:tech:xsd:head.001.001.01": "head",
	"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08": "fict",
}

func (v *Message) Body() interface{} {
	if v.FinancialInstitutionCreditTransfer != nil {
		return v.FinancialInstitutionCreditTransfer
	}
	return nil
}

// XSD ComplexType declarations
