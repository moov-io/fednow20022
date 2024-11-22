// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:camt.029.001.09
package ReturnRequestResponse_camt_029_001_09

import (
	"encoding/xml"

	"github.com/moov-io/fednow20022/pkg/fednow"
)

// XSD Elements

type Document struct {
	XMLName xml.Name

	RsltnOfInvstgtn ResolutionOfInvestigationV09 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 RsltnOfInvstgtn"`
}

// XSD ComplexType declarations

type ActiveOrHistoricCurrencyAndAmountFedNow1 struct {
	Value ActiveOrHistoricCurrencyAndAmountFedNow1SimpleType `xml:",chardata"`
	Ccy   ActiveCurrencyCodeFixed                            `xml:"Ccy,attr"`
}

type BranchAndFinancialInstitutionIdentification61 struct {
	FinInstnId FinancialInstitutionIdentification181 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 FinInstnId"`
}

type BranchAndFinancialInstitutionIdentification62 struct {
	FinInstnId FinancialInstitutionIdentification182 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 FinInstnId"`
}

type CancellationStatusReason3Choice struct {
	Cd    *ExternalPaymentCancellationRejection1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Cd,omitempty"`
	Prtry *Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Prtry,omitempty"`
}

type CancellationStatusReason41 struct {
	Orgtr    *PartyIdentification1351        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Orgtr,omitempty"`
	Rsn      CancellationStatusReason3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Rsn"`
	AddtlInf *Max105Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 AddtlInf,omitempty"`
}

type Case51 struct {
	Id    Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Id"`
	Cretr Party40Choice2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Cretr"`
}

type CaseAssignment51 struct {
	Id      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Id"`
	Assgnr  Party40Choice1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Assgnr"`
	Assgne  Party40Choice1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Assgne"`
	CreDtTm fednow.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CreDtTm"`
}

type Charges71 struct {
	Amt ActiveOrHistoricCurrencyAndAmountFedNow1      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Amt"`
	Agt BranchAndFinancialInstitutionIdentification62 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Agt"`
}

type ClearingSystemIdentification2Choice1 struct {
	Cd *ExternalClearingSystemIdentification1CodeFixed `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Cd,omitempty"`
}

type ClearingSystemMemberIdentification21 struct {
	ClrSysId ClearingSystemIdentification2Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 ClrSysId"`
	MmbId    RoutingNumberFRS1                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 MmbId"`
}

type Contact41 struct {
	Nm        *Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Nm,omitempty"`
	PhneNb    *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 PhneNb,omitempty"`
	MobNb     *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 MobNb,omitempty"`
	EmailAdr  *Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 EmailAdr,omitempty"`
	Dept      *Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Dept,omitempty"`
	PrefrdMtd PreferredContactMethod1Code1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 PrefrdMtd"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     fednow.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 BirthDt"`
	PrvcOfBirth *Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CityOfBirth"`
	CtryOfBirth CountryCode    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CtryOfBirth"`
}

type FinancialInstitutionIdentification181 struct {
	ClrSysMmbId ClearingSystemMemberIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 ClrSysMmbId"`
}

type FinancialInstitutionIdentification182 struct {
	BICFI       *BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 ClrSysMmbId"`
	LEI         *LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 LEI,omitempty"`
	Nm          *Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Nm,omitempty"`
	PstlAdr     *PostalAddress241                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 PstlAdr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 SchmeNm,omitempty"`
	Issr    *Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 SchmeNm,omitempty"`
	Issr    *Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Issr,omitempty"`
}

type InvestigationStatus5Choice1 struct {
	Conf *ExternalInvestigationExecutionConfirmation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Conf,omitempty"`
}

type OrganisationIdentification291 struct {
	AnyBIC *AnyBICDec2014Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 AnyBIC,omitempty"`
	LEI    *LEIIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 LEI,omitempty"`
	Othr   *GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    *ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Cd,omitempty"`
	Prtry *Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Prtry,omitempty"`
}

type OriginalGroupInformation291 struct {
	OrgnlMsgId   MessageIdentificationFedNow1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlMsgId"`
	OrgnlMsgNmId MessageNameIdentificationFRS1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlMsgNmId"`
	OrgnlCreDtTm fednow.ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlCreDtTm"`
}

type Party38Choice1 struct {
	OrgId  *OrganisationIdentification291 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgId,omitempty"`
	PrvtId *PersonIdentification131       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 PrvtId,omitempty"`
}

type Party40Choice1 struct {
	Agt *BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Agt,omitempty"`
}

type Party40Choice2 struct {
	Pty *PartyIdentification1351                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Pty,omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification62 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Agt,omitempty"`
}

type PartyIdentification1351 struct {
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Nm"`
	Id       *Party38Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Id,omitempty"`
	CtctDtls *Contact41      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CtctDtls,omitempty"`
}

type PaymentTransaction1021 struct {
	OrgnlGrpInf     OriginalGroupInformation291   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlGrpInf"`
	OrgnlInstrId    *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlEndToEndId,omitempty"`
	OrgnlTxId       *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlTxId,omitempty"`
	OrgnlUETR       *UUIDv4Identifier             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlUETR,omitempty"`
	CxlStsRsnInf    []*CancellationStatusReason41 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CxlStsRsnInf,omitempty"`
	RsltnRltdInf    *ResolutionData11             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 RsltnRltdInf,omitempty"`
}

type PersonIdentification131 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 DtAndPlcOfBirth,omitempty"`
	Othr            *GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    *ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Cd,omitempty"`
	Prtry *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Prtry,omitempty"`
}

type PostalAddress241 struct {
	Dept        *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Dept,omitempty"`
	SubDept     *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 SubDept,omitempty"`
	StrtNm      *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 StrtNm,omitempty"`
	BldgNb      *Max16Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 BldgNb,omitempty"`
	BldgNm      *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 BldgNm,omitempty"`
	Flr         *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Flr,omitempty"`
	PstBx       *Max16Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 PstBx,omitempty"`
	Room        *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Room,omitempty"`
	PstCd       Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 PstCd"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 TwnNm"`
	TwnLctnNm   *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 TwnLctnNm,omitempty"`
	DstrctNm    *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CtrySubDvsn"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Ctry"`
	AdrLine     *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 AdrLine,omitempty"`
}

type ResolutionData11 struct {
	EndToEndId     *Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 EndToEndId,omitempty"`
	TxId           *Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 TxId,omitempty"`
	UETR           *UUIDv4Identifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 UETR,omitempty"`
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmountFedNow1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 IntrBkSttlmAmt"`
	IntrBkSttlmDt  *fednow.ISODate                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 IntrBkSttlmDt,omitempty"`
	ClrChanl       *ClearingChannel2Code1                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 ClrChanl,omitempty"`
	Chrgs          *Charges71                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Chrgs,omitempty"`
}

type ResolutionOfInvestigationV09 struct {
	Assgnmt   CaseAssignment51            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Assgnmt"`
	RslvdCase Case51                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 RslvdCase"`
	Sts       InvestigationStatus5Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Sts"`
	CxlDtls   UnderlyingTransaction221    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CxlDtls"`
}

type UnderlyingTransaction221 struct {
	TxInfAndSts PaymentTransaction1021 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 TxInfAndSts"`
}

// XSD SimpleType declarations

type ActiveCurrencyCodeFixed string

const ActiveCurrencyCodeFixedUsd ActiveCurrencyCodeFixed = "USD"

type ActiveOrHistoricCurrencyAndAmountFedNow1SimpleType float64

type AnyBICDec2014Identifier string

type BICFIDec2014Identifier string

type ClearingChannel2Code1 string

const ClearingChannel2Code1Rtgs ClearingChannel2Code1 = "RTGS"
const ClearingChannel2Code1Rtns ClearingChannel2Code1 = "RTNS"
const ClearingChannel2Code1Mpns ClearingChannel2Code1 = "MPNS"

type CountryCode string

type ExternalClearingSystemIdentification1CodeFixed string

const ExternalClearingSystemIdentification1CodeFixedUsaba ExternalClearingSystemIdentification1CodeFixed = "USABA"

type ExternalInvestigationExecutionConfirmation1Code string

type ExternalOrganisationIdentification1Code string

type ExternalPaymentCancellationRejection1Code string

type ExternalPersonIdentification1Code string

type LEIIdentifier string

type Max105Text string

type Max140Text string

type Max16Text string

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
