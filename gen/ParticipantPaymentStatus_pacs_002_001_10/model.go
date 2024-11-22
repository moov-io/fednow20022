// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10
package ParticipantPaymentStatus_pacs_002_001_10

import (
	"encoding/xml"

	"github.com/moov-io/fednow20022/pkg/fednow"
)

// XSD Elements

type Document struct {
	XMLName xml.Name

	FIToFIPmtStsRpt FIToFIPaymentStatusReportV10 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 FIToFIPmtStsRpt"`
}

// XSD ComplexType declarations

type BranchAndFinancialInstitutionIdentification61 struct {
	FinInstnId FinancialInstitutionIdentification181 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 FinInstnId"`
}

type ClearingSystemIdentification2Choice1 struct {
	Cd *ExternalClearingSystemIdentification1CodeFixed `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Cd,omitempty"`
}

type ClearingSystemMemberIdentification21 struct {
	ClrSysId ClearingSystemIdentification2Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 ClrSysId"`
	MmbId    RoutingNumberFRS1                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 MmbId"`
}

type Contact41 struct {
	Nm        *Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Nm,omitempty"`
	PhneNb    *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 PhneNb,omitempty"`
	MobNb     *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 MobNb,omitempty"`
	EmailAdr  *Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 EmailAdr,omitempty"`
	Dept      *Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Dept,omitempty"`
	PrefrdMtd PreferredContactMethod1Code1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 PrefrdMtd"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     fednow.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 BirthDt"`
	PrvcOfBirth *Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 CityOfBirth"`
	CtryOfBirth CountryCode    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 CtryOfBirth"`
}

type FIToFIPaymentStatusReportV10 struct {
	GrpHdr      GroupHeader911         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 GrpHdr"`
	TxInfAndSts PaymentTransaction1101 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 TxInfAndSts"`
}

type FinancialInstitutionIdentification181 struct {
	ClrSysMmbId ClearingSystemMemberIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 ClrSysMmbId"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 SchmeNm,omitempty"`
	Issr    *Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 SchmeNm,omitempty"`
	Issr    *Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Issr,omitempty"`
}

type GroupHeader911 struct {
	MsgId   MessageIdentificationFedNow1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 MsgId"`
	CreDtTm fednow.ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 CreDtTm"`
}

type OrganisationIdentification291 struct {
	AnyBIC *AnyBICDec2014Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 AnyBIC,omitempty"`
	LEI    *LEIIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 LEI,omitempty"`
	Othr   *GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    *ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Cd,omitempty"`
	Prtry *Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Prtry,omitempty"`
}

type OriginalGroupInformation291 struct {
	OrgnlMsgId   Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlMsgId"`
	OrgnlMsgNmId MessageNameIdentificationFRS1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlMsgNmId"`
	OrgnlCreDtTm fednow.ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlCreDtTm"`
}

type Party38Choice1 struct {
	OrgId  *OrganisationIdentification291 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgId,omitempty"`
	PrvtId *PersonIdentification131       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 PrvtId,omitempty"`
}

type PartyIdentification1351 struct {
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Nm"`
	Id       *Party38Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Id,omitempty"`
	CtctDtls *Contact41      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 CtctDtls,omitempty"`
}

type PaymentTransaction1101 struct {
	OrgnlGrpInf     OriginalGroupInformation291                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlGrpInf"`
	OrgnlInstrId    *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlEndToEndId,omitempty"`
	OrgnlTxId       *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlTxId,omitempty"`
	OrgnlUETR       *UUIDv4Identifier                             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlUETR,omitempty"`
	TxSts           ExternalPaymentTransactionStatus1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 TxSts"`
	StsRsnInf       *StatusReasonInformation121                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 StsRsnInf,omitempty"`
	InstgAgt        BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 InstgAgt"`
	InstdAgt        BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 InstdAgt"`
}

type PersonIdentification131 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 DtAndPlcOfBirth,omitempty"`
	Othr            *GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    *ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Cd,omitempty"`
	Prtry *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Prtry,omitempty"`
}

type StatusReason6Choice1 struct {
	Cd *ExternalStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Cd,omitempty"`
}

type StatusReasonInformation121 struct {
	Orgtr    *PartyIdentification1351 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Orgtr,omitempty"`
	Rsn      StatusReason6Choice1     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Rsn"`
	AddtlInf *Max105Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 AddtlInf,omitempty"`
}

// XSD SimpleType declarations

type AnyBICDec2014Identifier string

type CountryCode string

type ExternalClearingSystemIdentification1CodeFixed string

const ExternalClearingSystemIdentification1CodeFixedUsaba ExternalClearingSystemIdentification1CodeFixed = "USABA"

type ExternalOrganisationIdentification1Code string

type ExternalPaymentTransactionStatus1Code string

type ExternalPersonIdentification1Code string

type ExternalStatusReason1Code string

type LEIIdentifier string

type Max105Text string

type Max140Text string

type Max2048Text string

type Max35Text string

type Max70Text string

type MessageIdentificationFedNow1 string

type MessageNameIdentificationFRS1 string

type PhoneNumber string

type PreferredContactMethod1Code1 string

const PreferredContactMethod1Code1Mail PreferredContactMethod1Code1 = "MAIL"
const PreferredContactMethod1Code1Phon PreferredContactMethod1Code1 = "PHON"
const PreferredContactMethod1Code1Cell PreferredContactMethod1Code1 = "CELL"

type RoutingNumberFRS1 string

type UUIDv4Identifier string
