// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:admi.006.001.01
package RetrievalRequest_admi_006_001_01

import (
	"encoding/xml"

	"github.com/moov-io/fednow20022/pkg/fednow"
)

// XSD Elements

type Document struct {
	XMLName xml.Name

	RsndReq ResendRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.006.001.01 RsndReq"`
}

// XSD ComplexType declarations

type GenericIdentification361 struct {
	Id   ConnectionPartyIdentifierFedNow1 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.006.001.01 Id"`
	Issr Max35TextFixed                   `xml:"urn:iso:std:iso:20022:tech:xsd:admi.006.001.01 Issr"`
}

type MessageHeader71 struct {
	MsgId   MessageIdentificationFedNow1 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.006.001.01 MsgId"`
	CreDtTm fednow.ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:admi.006.001.01 CreDtTm"`
}

type PartyIdentification120Choice1 struct {
	PrtryId *GenericIdentification361 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.006.001.01 PrtryId,omitempty"`
}

type PartyIdentification1361 struct {
	Id PartyIdentification120Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.006.001.01 Id"`
}

type ResendRequestV01 struct {
	MsgHdr      MessageHeader71          `xml:"urn:iso:std:iso:20022:tech:xsd:admi.006.001.01 MsgHdr"`
	RsndSchCrit []ResendSearchCriteria21 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.006.001.01 RsndSchCrit"`
}

type ResendSearchCriteria21 struct {
	BizDt        fednow.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:admi.006.001.01 BizDt"`
	OrgnlMsgNmId MessageNameIdentificationFRS1 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.006.001.01 OrgnlMsgNmId"`
	FileRef      Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:admi.006.001.01 FileRef"`
	Rcpt         PartyIdentification1361       `xml:"urn:iso:std:iso:20022:tech:xsd:admi.006.001.01 Rcpt"`
}

// XSD SimpleType declarations

type ConnectionPartyIdentifierFedNow1 string

type Max35Text string

type Max35TextFixed string

const Max35TextFixedNa Max35TextFixed = "NA"

type MessageIdentificationFedNow1 string

type MessageNameIdentificationFRS1 string
