// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:camt.026.001.07
package InformationRequest_camt_026_001_07

import (
	"encoding/xml"

	"github.com/moov-io/fednow20022/pkg/fednow"
)

// XSD Elements

type Document struct {
	XMLName xml.Name

	UblToApply UnableToApplyV07 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 UblToApply"`
}

// XSD ComplexType declarations

type ActiveCurrencyAndAmountFedNow1 struct {
	Value ActiveCurrencyAndAmountFedNow1SimpleType `xml:",chardata"`
	Ccy   ActiveCurrencyCodeFixed                  `xml:"Ccy,attr"`
}

type ActiveOrHistoricCurrencyAndAmountFedNow1 struct {
	Value ActiveOrHistoricCurrencyAndAmountFedNow1SimpleType `xml:",chardata"`
	Ccy   ActiveCurrencyCodeFixed                            `xml:"Ccy,attr"`
}

type BranchAndFinancialInstitutionIdentification61 struct {
	FinInstnId FinancialInstitutionIdentification181 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 FinInstnId"`
}

type BranchAndFinancialInstitutionIdentification62 struct {
	FinInstnId FinancialInstitutionIdentification182 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 FinInstnId"`
}

type Case51 struct {
	Id    Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Id"`
	Cretr Party40Choice2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Cretr"`
}

type CaseAssignment51 struct {
	Id      MessageIdentificationFedNow1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Id"`
	Assgnr  Party40Choice1               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Assgnr"`
	Assgne  Party40Choice1               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Assgne"`
	CreDtTm fednow.ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 CreDtTm"`
}

type ClearingSystemIdentification2Choice1 struct {
	Cd *ExternalClearingSystemIdentification1CodeFixed `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Cd,omitempty"`
}

type ClearingSystemMemberIdentification21 struct {
	ClrSysId ClearingSystemIdentification2Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 ClrSysId"`
	MmbId    RoutingNumberFRS1                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 MmbId"`
}

type Contact41 struct {
	Nm        *Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Nm,omitempty"`
	PhneNb    *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 PhneNb,omitempty"`
	MobNb     *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 MobNb,omitempty"`
	EmailAdr  *Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 EmailAdr,omitempty"`
	Dept      *Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Dept,omitempty"`
	PrefrdMtd PreferredContactMethod1Code1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 PrefrdMtd"`
}

type DateAndDateTime2Choice struct {
	Dt   *fednow.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Dt,omitempty"`
	DtTm *fednow.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 DtTm,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     fednow.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 BirthDt"`
	PrvcOfBirth *Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 CityOfBirth"`
	CtryOfBirth CountryCode    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 CtryOfBirth"`
}

type FinancialInstitutionIdentification181 struct {
	ClrSysMmbId ClearingSystemMemberIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 ClrSysMmbId"`
}

type FinancialInstitutionIdentification182 struct {
	BICFI       *BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 ClrSysMmbId"`
	LEI         *LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 LEI,omitempty"`
	Nm          *Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Nm,omitempty"`
	PstlAdr     *PostalAddress241                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 PstlAdr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 SchmeNm,omitempty"`
	Issr    *Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 SchmeNm,omitempty"`
	Issr    *Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Issr,omitempty"`
}

type MissingOrIncorrectInformation31 struct {
	AMLReq     *AMLIndicator               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 AMLReq,omitempty"`
	MssngInf   []*UnableToApplyMissing11   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 MssngInf,omitempty"`
	IncrrctInf []*UnableToApplyIncorrect11 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 IncrrctInf,omitempty"`
}

type OrganisationIdentification291 struct {
	AnyBIC *AnyBICDec2014Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 AnyBIC,omitempty"`
	LEI    *LEIIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 LEI,omitempty"`
	Othr   *GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    *ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Cd,omitempty"`
	Prtry *Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Prtry,omitempty"`
}

type Party38Choice1 struct {
	OrgId  *OrganisationIdentification291 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgId,omitempty"`
	PrvtId *PersonIdentification131       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 PrvtId,omitempty"`
}

type Party40Choice1 struct {
	Agt *BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Agt,omitempty"`
}

type Party40Choice2 struct {
	Pty *PartyIdentification1351                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Pty,omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification62 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Agt,omitempty"`
}

type PartyIdentification1351 struct {
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Nm"`
	Id       *Party38Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Id,omitempty"`
	CtctDtls *Contact41      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 CtctDtls,omitempty"`
}

type PersonIdentification131 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 DtAndPlcOfBirth,omitempty"`
	Othr            *GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    *ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Cd,omitempty"`
	Prtry *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Prtry,omitempty"`
}

type PostalAddress241 struct {
	Dept        *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Dept,omitempty"`
	SubDept     *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 SubDept,omitempty"`
	StrtNm      *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 StrtNm,omitempty"`
	BldgNb      *Max16Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 BldgNb,omitempty"`
	BldgNm      *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 BldgNm,omitempty"`
	Flr         *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Flr,omitempty"`
	PstBx       *Max16Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 PstBx,omitempty"`
	Room        *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Room,omitempty"`
	PstCd       Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 PstCd"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 TwnNm"`
	TwnLctnNm   *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 TwnLctnNm,omitempty"`
	DstrctNm    *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 CtrySubDvsn"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Ctry"`
	AdrLine     *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 AdrLine,omitempty"`
}

type UnableToApplyIncorrect11 struct {
	Cd              UnableToApplyIncorrectInformation4Code1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Cd"`
	AddtlIncrrctInf *Max140Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 AddtlIncrrctInf,omitempty"`
}

type UnableToApplyJustification3Choice1 struct {
	MssngOrIncrrctInf *MissingOrIncorrectInformation31 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 MssngOrIncrrctInf,omitempty"`
	PssblDplctInstr   *TrueFalseIndicator              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 PssblDplctInstr,omitempty"`
}

type UnableToApplyMissing11 struct {
	Cd            UnableToApplyMissingInformation3Code1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Cd"`
	AddtlMssngInf *Max140Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 AddtlMssngInf,omitempty"`
}

type UnableToApplyV07 struct {
	Assgnmt CaseAssignment51                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Assgnmt"`
	Case    Case51                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Case"`
	Undrlyg UnderlyingTransaction5Choice1      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Undrlyg"`
	Justfn  UnableToApplyJustification3Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Justfn"`
}

type UnderlyingGroupInformation11 struct {
	OrgnlMsgId   Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlMsgId"`
	OrgnlMsgNmId MessageNameIdentificationFRS1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlMsgNmId"`
	OrgnlCreDtTm fednow.ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlCreDtTm"`
}

type UnderlyingPaymentInstruction51 struct {
	OrgnlGrpInf     UnderlyingGroupInformation11             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlGrpInf"`
	OrgnlPmtInfId   Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlPmtInfId"`
	OrgnlInstrId    *Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlEndToEndId"`
	OrgnlUETR       *UUIDv4Identifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlUETR,omitempty"`
	OrgnlInstdAmt   ActiveOrHistoricCurrencyAndAmountFedNow1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlInstdAmt"`
	ReqdExctnDt     DateAndDateTime2Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 ReqdExctnDt"`
}

type UnderlyingPaymentTransaction41 struct {
	OrgnlGrpInf         UnderlyingGroupInformation11   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlGrpInf"`
	OrgnlInstrId        *Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId     *Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlEndToEndId,omitempty"`
	OrgnlTxId           *Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlTxId,omitempty"`
	OrgnlUETR           *UUIDv4Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlUETR,omitempty"`
	OrgnlIntrBkSttlmAmt ActiveCurrencyAndAmountFedNow1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlIntrBkSttlmAmt"`
	OrgnlIntrBkSttlmDt  fednow.ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlIntrBkSttlmDt"`
}

type UnderlyingTransaction5Choice1 struct {
	Initn  *UnderlyingPaymentInstruction51 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Initn,omitempty"`
	IntrBk *UnderlyingPaymentTransaction41 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 IntrBk,omitempty"`
}

// XSD SimpleType declarations

type AMLIndicator bool

type ActiveCurrencyAndAmountFedNow1SimpleType float64

type ActiveCurrencyCodeFixed string

const ActiveCurrencyCodeFixedUsd ActiveCurrencyCodeFixed = "USD"

type ActiveOrHistoricCurrencyAndAmountFedNow1SimpleType float64

type AnyBICDec2014Identifier string

type BICFIDec2014Identifier string

type CountryCode string

type ExternalClearingSystemIdentification1CodeFixed string

const ExternalClearingSystemIdentification1CodeFixedUsaba ExternalClearingSystemIdentification1CodeFixed = "USABA"

type ExternalOrganisationIdentification1Code string

type ExternalPersonIdentification1Code string

type LEIIdentifier string

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

type TrueFalseIndicator bool

type UUIDv4Identifier string

type UnableToApplyIncorrectInformation4Code1 string

const UnableToApplyIncorrectInformation4Code1In02 UnableToApplyIncorrectInformation4Code1 = "IN02"
const UnableToApplyIncorrectInformation4Code1In03 UnableToApplyIncorrectInformation4Code1 = "IN03"
const UnableToApplyIncorrectInformation4Code1In06 UnableToApplyIncorrectInformation4Code1 = "IN06"
const UnableToApplyIncorrectInformation4Code1In07 UnableToApplyIncorrectInformation4Code1 = "IN07"
const UnableToApplyIncorrectInformation4Code1In08 UnableToApplyIncorrectInformation4Code1 = "IN08"
const UnableToApplyIncorrectInformation4Code1In13 UnableToApplyIncorrectInformation4Code1 = "IN13"
const UnableToApplyIncorrectInformation4Code1In14 UnableToApplyIncorrectInformation4Code1 = "IN14"
const UnableToApplyIncorrectInformation4Code1In15 UnableToApplyIncorrectInformation4Code1 = "IN15"
const UnableToApplyIncorrectInformation4Code1In16 UnableToApplyIncorrectInformation4Code1 = "IN16"
const UnableToApplyIncorrectInformation4Code1Mm20 UnableToApplyIncorrectInformation4Code1 = "MM20"
const UnableToApplyIncorrectInformation4Code1Mm21 UnableToApplyIncorrectInformation4Code1 = "MM21"
const UnableToApplyIncorrectInformation4Code1Mm25 UnableToApplyIncorrectInformation4Code1 = "MM25"
const UnableToApplyIncorrectInformation4Code1Mm26 UnableToApplyIncorrectInformation4Code1 = "MM26"
const UnableToApplyIncorrectInformation4Code1Mm27 UnableToApplyIncorrectInformation4Code1 = "MM27"
const UnableToApplyIncorrectInformation4Code1Mm28 UnableToApplyIncorrectInformation4Code1 = "MM28"
const UnableToApplyIncorrectInformation4Code1Mm29 UnableToApplyIncorrectInformation4Code1 = "MM29"
const UnableToApplyIncorrectInformation4Code1Mm30 UnableToApplyIncorrectInformation4Code1 = "MM30"
const UnableToApplyIncorrectInformation4Code1Mm31 UnableToApplyIncorrectInformation4Code1 = "MM31"
const UnableToApplyIncorrectInformation4Code1Mm32 UnableToApplyIncorrectInformation4Code1 = "MM32"
const UnableToApplyIncorrectInformation4Code1Mm35 UnableToApplyIncorrectInformation4Code1 = "MM35"
const UnableToApplyIncorrectInformation4Code1In38 UnableToApplyIncorrectInformation4Code1 = "IN38"
const UnableToApplyIncorrectInformation4Code1In39 UnableToApplyIncorrectInformation4Code1 = "IN39"
const UnableToApplyIncorrectInformation4Code1Narr UnableToApplyIncorrectInformation4Code1 = "NARR"

type UnableToApplyMissingInformation3Code1 string

const UnableToApplyMissingInformation3Code1Ms01 UnableToApplyMissingInformation3Code1 = "MS01"
const UnableToApplyMissingInformation3Code1Narr UnableToApplyMissingInformation3Code1 = "NARR"
