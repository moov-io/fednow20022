// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:fednow
package Document_InformationRequest_camt_026_001_07

import (
	"encoding/xml"

	"github.com/moov-io/fednow20022/gen/InformationRequest_camt_026_001_07"
	"github.com/moov-io/fednow20022/gen/head_001_001_02"
)

// XSD Elements

type Message struct {
	XMLName xml.Name   `xml:"Message"`
	Xmlns   []xml.Attr `xml:",attr"`

	AppHdr head_001_001_02.BusinessApplicationHeaderV02 `xml:"urn:fednow AppHdr"`

	InformationRequest *InformationRequest_camt_026_001_07.Document `xml:"urn:fednow InformationRequest,omitempty"`
}

var NamespacePrefixMap = map[string]string{
	"urn:iso:std:iso:20022:tech:xsd:head.001.001.01": "head",
	"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07": "ir",
}

func (v *Message) Body() interface{} {
	if v.InformationRequest != nil {
		return v.InformationRequest
	}
	return nil
}

// XSD ComplexType declarations
