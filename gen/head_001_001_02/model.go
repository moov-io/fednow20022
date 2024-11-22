// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:head.001.001.02
package head_001_001_02

import (
	"encoding/xml"

	"github.com/moov-io/fednow20022/pkg/fednow"
)

// XSD Elements

type AppHdr struct {
	XMLName xml.Name

	Fr Party44Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Fr"`

	To Party44Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 To"`

	BizMsgIdr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BizMsgIdr"`

	MsgDefIdr MessageNameIdentificationFRS1 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 MsgDefIdr"`

	BizSvc *Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BizSvc,omitempty"`

	MktPrctc ImplementationSpecification11 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 MktPrctc"`

	CreDt fednow.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 CreDt"`

	BizPrcgDt *fednow.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BizPrcgDt,omitempty"`

	CpyDplct *CopyDuplicate1Code1 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 CpyDplct,omitempty"`

	Rltd *BusinessApplicationHeader51 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Rltd,omitempty"`
}

// XSD ComplexType declarations

type BranchAndFinancialInstitutionIdentification61 struct {
	FinInstnId FinancialInstitutionIdentification181 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 FinInstnId"`
}

type BusinessApplicationHeader51 struct {
	Fr        Party44Choice1                `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Fr"`
	To        Party44Choice1                `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 To"`
	BizMsgIdr Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BizMsgIdr"`
	MsgDefIdr MessageNameIdentificationFRS1 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 MsgDefIdr"`
	BizSvc    *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BizSvc,omitempty"`
	CreDt     fednow.ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 CreDt"`
}

type BusinessApplicationHeaderV02 struct {
	Fr        Party44Choice1                `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Fr"`
	To        Party44Choice1                `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 To"`
	BizMsgIdr Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BizMsgIdr"`
	MsgDefIdr MessageNameIdentificationFRS1 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 MsgDefIdr"`
	BizSvc    *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BizSvc,omitempty"`
	MktPrctc  ImplementationSpecification11 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 MktPrctc"`
	CreDt     fednow.ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 CreDt"`
	BizPrcgDt *fednow.ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BizPrcgDt,omitempty"`
	CpyDplct  *CopyDuplicate1Code1          `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 CpyDplct,omitempty"`
	Rltd      *BusinessApplicationHeader51  `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Rltd,omitempty"`
}

type ClearingSystemMemberIdentification21 struct {
	MmbId ConnectionPartyIdentifierFedNow1 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 MmbId"`
}

type FinancialInstitutionIdentification181 struct {
	ClrSysMmbId ClearingSystemMemberIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 ClrSysMmbId"`
}

type ImplementationSpecification11 struct {
	Regy Max350TextFixed                     `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Regy"`
	Id   MarketPracticeIdentificationFedNow1 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Id"`
}

type Party44Choice1 struct {
	FIId *BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 FIId,omitempty"`
}

// XSD SimpleType declarations

type ConnectionPartyIdentifierFedNow1 string

type CopyDuplicate1Code1 string

const CopyDuplicate1Code1Dupl CopyDuplicate1Code1 = "DUPL"

type MarketPracticeIdentificationFedNow1 string

type Max350TextFixed string

const Max350TextFixedWww2SwiftCommystandardsgroupfederalReserveFinancialServicesfednowService Max350TextFixed = "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/FedNow_Service"

type Max35Text string

type MessageNameIdentificationFRS1 string
