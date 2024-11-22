// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:fednow
package Document_RetrievalRequest_admi_006_001_01

import (
	"fmt"

	"encoding/xml"
)

func NewAdmi006Message() *Message {
	message := &Message{}
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name:  xml.Name{Local: "xmlns"},
		Value: "urn:fednow",
	})
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name:  xml.Name{Local: "xmlns:head"},
		Value: "urn:iso:std:iso:20022:tech:xsd:head.001.001.01",
	})
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: fmt.Sprintf("xmlns:%s", NamespacePrefixMap["urn:iso:std:iso:20022:tech:xsd:admi.006.001.01"]),
		},
		Value: "urn:iso:std:iso:20022:tech:xsd:admi.006.001.01",
	})
	return message
}
