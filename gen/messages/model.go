// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:tch
package messages

import (
	"encoding/xml"

	"github.com/moov-io/fednow20022/gen/admi_004_001_02"
	"github.com/moov-io/fednow20022/gen/head_001_001_02"
	"github.com/moov-io/fednow20022/gen/pacs_008_001_08"
)

// XSD Elements

type Message struct {
	XMLName xml.Name   `xml:"Message"`
	Xmlns   []xml.Attr `xml:",attr"`

	AppHdr head_001_001_02.BusinessApplicationHeaderV02 `xml:"urn:tch AppHdr"`

	CreditTransfer *pacs_008_001_08.Document `xml:"urn:tch CreditTransfer,omitempty"`

	SystemNotificationEvent *admi_004_001_02.Document `xml:"urn:tch SystemNotificationEvent,omitempty"`
}

var NamespacePrefixMap = map[string]string{
	"urn:iso:std:iso:20022:tech:xsd:head.001.001.01": "head",
	"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02": "ne",
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08": "ct",
}

func (v *Message) Body() interface{} {
	if v.CreditTransfer != nil {
		return v.CreditTransfer
	}
	if v.SystemNotificationEvent != nil {
		return v.SystemNotificationEvent
	}
	return nil
}

// XSD ComplexType declarations