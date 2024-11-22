// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:fednow
package Document_CustomerCreditTransfer_pacs_008_001_08

import (
	"encoding/xml"

	"github.com/moov-io/fednow20022/gen/CustomerCreditTransfer_pacs_008_001_08"
	"github.com/moov-io/fednow20022/gen/head_001_001_02"
)

// XSD Elements

type Message struct {
	XMLName xml.Name   `xml:"Message"`
	Xmlns   []xml.Attr `xml:",attr"`

	AppHdr head_001_001_02.BusinessApplicationHeaderV02 `xml:"urn:fednow AppHdr"`

	CustomerCreditTransfer *CustomerCreditTransfer_pacs_008_001_08.Document `xml:"urn:fednow CustomerCreditTransfer,omitempty"`
}

var NamespacePrefixMap = map[string]string{
	"urn:iso:std:iso:20022:tech:xsd:head.001.001.01": "head",
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08": "cct",
}

func (v *Message) Body() interface{} {
	if v.CustomerCreditTransfer != nil {
		return v.CustomerCreditTransfer
	}
	return nil
}

// XSD ComplexType declarations
