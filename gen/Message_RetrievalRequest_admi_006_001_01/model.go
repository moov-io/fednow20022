// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:fednow
package Message_RetrievalRequest_admi_006_001_01

import (
	"encoding/xml"

	"github.com/moov-io/fednow20022/gen/RetrievalRequest_admi_006_001_01"
	"github.com/moov-io/fednow20022/gen/head_001_001_02"
)

// XSD Elements

type Message struct {
	XMLName xml.Name   `xml:"Message"`
	Xmlns   []xml.Attr `xml:",attr"`

	AppHdr head_001_001_02.BusinessApplicationHeaderV02 `xml:"urn:fednow AppHdr"`

	RetrievalRequest *RetrievalRequest_admi_006_001_01.Document `xml:"urn:fednow RetrievalRequest,omitempty"`
}

var NamespacePrefixMap = map[string]string{
	"urn:iso:std:iso:20022:tech:xsd:head.001.001.01": "head",
	"urn:iso:std:iso:20022:tech:xsd:admi.006.001.01": "rr",
}

func (v *Message) Body() interface{} {
	if v.RetrievalRequest != nil {
		return v.RetrievalRequest
	}
	return nil
}

// XSD ComplexType declarations
