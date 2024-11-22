// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:fednow
package Document_FedNowSystemResponse_admi_011_001_01

import (
	"encoding/xml"

	"github.com/moov-io/fednow20022/gen/FedNowSystemResponse_admi_011_001_01"
	"github.com/moov-io/fednow20022/gen/head_001_001_02"
)

// XSD Elements

type Message struct {
	XMLName xml.Name   `xml:"Message"`
	Xmlns   []xml.Attr `xml:",attr"`

	AppHdr head_001_001_02.BusinessApplicationHeaderV02 `xml:"urn:fednow AppHdr"`

	FedNowSystemResponse *FedNowSystemResponse_admi_011_001_01.Document `xml:"urn:fednow FedNowSystemResponse,omitempty"`
}

var NamespacePrefixMap = map[string]string{
	"urn:iso:std:iso:20022:tech:xsd:head.001.001.01": "head",
	"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01": "fnsr",
}

func (v *Message) Body() interface{} {
	if v.FedNowSystemResponse != nil {
		return v.FedNowSystemResponse
	}
	return nil
}

// XSD ComplexType declarations
