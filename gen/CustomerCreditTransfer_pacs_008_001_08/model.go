// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08
package CustomerCreditTransfer_pacs_008_001_08

import (
	"encoding/xml"

	"github.com/moov-io/fednow20022/pkg/fednow"
)

// XSD Elements

type Document struct {
	XMLName xml.Name

	FIToFICstmrCdtTrf FIToFICustomerCreditTransferV08 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 FIToFICstmrCdtTrf"`
}

// XSD ComplexType declarations

type AccountIdentification4Choice1 struct {
	IBAN *IBAN2007Identifier             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 IBAN,omitempty"`
	Othr *GenericAccountIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Othr,omitempty"`
}

type ActiveCurrencyAndAmountFedNow1 struct {
	Value ActiveCurrencyAndAmountFedNow1SimpleType `xml:",chardata"`
	Ccy   ActiveCurrencyCodeFixed                  `xml:"Ccy,attr"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value ActiveOrHistoricCurrencyAndAmountSimpleType `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode                `xml:"Ccy,attr"`
}

type AddressType3Choice struct {
	Cd    *AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cd,omitempty"`
	Prtry *GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prtry,omitempty"`
}

type BranchAndFinancialInstitutionIdentification61 struct {
	FinInstnId FinancialInstitutionIdentification181 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 FinInstnId"`
}

type BranchAndFinancialInstitutionIdentification62 struct {
	FinInstnId FinancialInstitutionIdentification182 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 FinInstnId"`
}

type CashAccount381 struct {
	Id   AccountIdentification4Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Id"`
	Tp   *CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Tp,omitempty"`
	Prxy *ProxyAccountIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    *ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cd,omitempty"`
	Prtry *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prtry,omitempty"`
}

type CategoryPurpose1Choice1 struct {
	Prtry *Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prtry,omitempty"`
}

type ClearingSystemIdentification2Choice1 struct {
	Cd *ExternalClearingSystemIdentification1CodeFixed `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cd,omitempty"`
}

type ClearingSystemIdentification3Choice1 struct {
	Cd *ExternalCashClearingSystem1CodeFixed `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cd,omitempty"`
}

type ClearingSystemMemberIdentification21 struct {
	ClrSysId ClearingSystemIdentification2Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 ClrSysId"`
	MmbId    RoutingNumberFRS1                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 MmbId"`
}

type Contact4 struct {
	NmPrfx    *NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 NmPrfx,omitempty"`
	Nm        *Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Nm,omitempty"`
	PhneNb    *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 PhneNb,omitempty"`
	MobNb     *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 MobNb,omitempty"`
	FaxNb     *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 FaxNb,omitempty"`
	EmailAdr  *Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 EmailAdr,omitempty"`
	EmailPurp *Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 EmailPurp,omitempty"`
	JobTitl   *Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 JobTitl,omitempty"`
	Rspnsblty *Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Rspnsblty,omitempty"`
	Dept      *Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Dept,omitempty"`
	Othr      []*OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Othr,omitempty"`
	PrefrdMtd *PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 PrefrdMtd,omitempty"`
}

type CreditTransferTransaction391 struct {
	PmtId          PaymentIdentification71                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 PmtId"`
	PmtTpInf       PaymentTypeInformation281                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 PmtTpInf"`
	IntrBkSttlmAmt ActiveCurrencyAndAmountFedNow1                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 IntrBkSttlmAmt"`
	IntrBkSttlmDt  fednow.ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 IntrBkSttlmDt"`
	ChrgBr         ChargeBearerType1Code1                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 ChrgBr"`
	InstgAgt       BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 InstgAgt"`
	InstdAgt       BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 InstdAgt"`
	UltmtDbtr      *PartyIdentification1351                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 UltmtDbtr,omitempty"`
	InitgPty       *PartyIdentification1351                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 InitgPty,omitempty"`
	Dbtr           PartyIdentification1351                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Dbtr"`
	DbtrAcct       CashAccount381                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 DbtrAcct"`
	DbtrAgt        BranchAndFinancialInstitutionIdentification62 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 DbtrAgt"`
	CdtrAgt        BranchAndFinancialInstitutionIdentification62 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CdtrAgt"`
	Cdtr           PartyIdentification1351                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cdtr"`
	CdtrAcct       CashAccount381                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CdtrAcct"`
	UltmtCdtr      *PartyIdentification1351                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 UltmtCdtr,omitempty"`
	Purp           *Purpose2Choice                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Purp,omitempty"`
	RltdRmtInf     *RemittanceLocation71                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 RltdRmtInf,omitempty"`
	RmtInf         *RemittanceInformation161                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 RmtInf,omitempty"`
}

type CreditorReferenceInformation2 struct {
	Tp  *CreditorReferenceType2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Tp,omitempty"`
	Ref *Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Ref,omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    *DocumentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cd,omitempty"`
	Prtry *Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prtry,omitempty"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CdOrPrtry"`
	Issr      *Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Issr,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     fednow.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 BirthDt"`
	PrvcOfBirth *Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CityOfBirth"`
	CtryOfBirth CountryCode    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CtryOfBirth"`
}

type DatePeriod2 struct {
	FrDt fednow.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 FrDt"`
	ToDt fednow.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 ToDt"`
}

type DiscountAmountAndType1 struct {
	Tp  *DiscountAmountType1Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Amt"`
}

type DiscountAmountType1Choice struct {
	Cd    *ExternalDiscountAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cd,omitempty"`
	Prtry *Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prtry,omitempty"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Amt"`
	CdtDbtInd *CreditDebitCode                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CdtDbtInd,omitempty"`
	Rsn       *Max4Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Rsn,omitempty"`
	AddtlInf  *Max140Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 AddtlInf,omitempty"`
}

type DocumentLineIdentification1 struct {
	Tp     *DocumentLineType1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Tp,omitempty"`
	Nb     *Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Nb,omitempty"`
	RltdDt *fednow.ISODate    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 RltdDt,omitempty"`
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Id"`
	Desc *Max2048Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Desc,omitempty"`
	Amt  *RemittanceAmount3            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Amt,omitempty"`
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CdOrPrtry"`
	Issr      *Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Issr,omitempty"`
}

type DocumentLineType1Choice struct {
	Cd    *ExternalDocumentLineType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cd,omitempty"`
	Prtry *Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prtry,omitempty"`
}

type FIToFICustomerCreditTransferV08 struct {
	GrpHdr      GroupHeader931               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 GrpHdr"`
	CdtTrfTxInf CreditTransferTransaction391 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CdtTrfTxInf"`
}

type FinancialInstitutionIdentification181 struct {
	ClrSysMmbId ClearingSystemMemberIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 ClrSysMmbId"`
}

type FinancialInstitutionIdentification182 struct {
	BICFI       *BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 ClrSysMmbId"`
	LEI         *LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 LEI,omitempty"`
	Nm          *Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Nm,omitempty"`
	PstlAdr     *PostalAddress241                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 PstlAdr,omitempty"`
}

type Garnishment3 struct {
	Tp                GarnishmentType1                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Tp"`
	Grnshee           *PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Grnshee,omitempty"`
	GrnshmtAdmstr     *PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 GrnshmtAdmstr,omitempty"`
	RefNb             *Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 RefNb,omitempty"`
	Dt                *fednow.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Dt,omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 RmtdAmt,omitempty"`
	FmlyMdclInsrncInd *TrueFalseIndicator                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 FmlyMdclInsrncInd,omitempty"`
	MplyeeTermntnInd  *TrueFalseIndicator                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 MplyeeTermntnInd,omitempty"`
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CdOrPrtry"`
	Issr      *Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Issr,omitempty"`
}

type GarnishmentType1Choice struct {
	Cd    *ExternalGarnishmentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cd,omitempty"`
	Prtry *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prtry,omitempty"`
}

type GenericAccountIdentification11 struct {
	Id Max34Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Id"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Issr"`
	SchmeNm *Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 SchmeNm,omitempty"`
	Issr    *Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 SchmeNm,omitempty"`
	Issr    *Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Issr,omitempty"`
}

type GroupHeader931 struct {
	MsgId    MessageIdentificationFedNow1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 MsgId"`
	CreDtTm  fednow.ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CreDtTm"`
	NbOfTxs  Max15NumericTextFixed        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 NbOfTxs"`
	SttlmInf SettlementInstruction71      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 SttlmInf"`
}

type LocalInstrument2Choice1 struct {
	Prtry *Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prtry,omitempty"`
}

type NameAndAddress161 struct {
	Nm  Max140Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Nm"`
	Adr PostalAddress241 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Adr"`
}

type OrganisationIdentification29 struct {
	AnyBIC *AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 AnyBIC,omitempty"`
	LEI    *LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 LEI,omitempty"`
	Othr   []*GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Othr,omitempty"`
}

type OrganisationIdentification291 struct {
	AnyBIC *AnyBICDec2014Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 AnyBIC,omitempty"`
	LEI    *LEIIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 LEI,omitempty"`
	Othr   *GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    *ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cd,omitempty"`
	Prtry *Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prtry,omitempty"`
}

type OtherContact1 struct {
	ChanlTp Max4Text    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 ChanlTp"`
	Id      *Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  *OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 OrgId,omitempty"`
	PrvtId *PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 PrvtId,omitempty"`
}

type Party38Choice1 struct {
	OrgId  *OrganisationIdentification291 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 OrgId,omitempty"`
	PrvtId *PersonIdentification131       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 PrvtId,omitempty"`
}

type PartyIdentification135 struct {
	Nm        *Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Nm,omitempty"`
	PstlAdr   *PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 PstlAdr,omitempty"`
	Id        *Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Id,omitempty"`
	CtryOfRes *CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CtryOfRes,omitempty"`
	CtctDtls  *Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CtctDtls,omitempty"`
}

type PartyIdentification1351 struct {
	Nm        Max140Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Nm"`
	PstlAdr   *PostalAddress241 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 PstlAdr,omitempty"`
	Id        *Party38Choice1   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Id,omitempty"`
	CtryOfRes *CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CtryOfRes,omitempty"`
}

type PaymentIdentification71 struct {
	InstrId    *Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 InstrId,omitempty"`
	EndToEndId Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 EndToEndId"`
	TxId       *Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TxId,omitempty"`
	UETR       *UUIDv4Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 UETR,omitempty"`
}

type PaymentTypeInformation281 struct {
	SvcLvl    *ServiceLevel8Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 LclInstrm"`
	CtgyPurp  CategoryPurpose1Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CtgyPurp"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 DtAndPlcOfBirth,omitempty"`
	Othr            []*GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Othr,omitempty"`
}

type PersonIdentification131 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 DtAndPlcOfBirth,omitempty"`
	Othr            *GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    *ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cd,omitempty"`
	Prtry *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prtry,omitempty"`
}

type PostalAddress24 struct {
	AdrTp       *AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 AdrTp,omitempty"`
	Dept        *Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Dept,omitempty"`
	SubDept     *Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 SubDept,omitempty"`
	StrtNm      *Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 StrtNm,omitempty"`
	BldgNb      *Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 BldgNb,omitempty"`
	BldgNm      *Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 BldgNm,omitempty"`
	Flr         *Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Flr,omitempty"`
	PstBx       *Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 PstBx,omitempty"`
	Room        *Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Room,omitempty"`
	PstCd       *Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 PstCd,omitempty"`
	TwnNm       *Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TwnNm,omitempty"`
	TwnLctnNm   *Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TwnLctnNm,omitempty"`
	DstrctNm    *Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 DstrctNm,omitempty"`
	CtrySubDvsn *Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CtrySubDvsn,omitempty"`
	Ctry        *CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Ctry,omitempty"`
	AdrLine     []*Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 AdrLine,omitempty"`
}

type PostalAddress241 struct {
	Dept        *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Dept,omitempty"`
	SubDept     *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 SubDept,omitempty"`
	StrtNm      *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 StrtNm,omitempty"`
	BldgNb      *Max16Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 BldgNb,omitempty"`
	BldgNm      *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 BldgNm,omitempty"`
	Flr         *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Flr,omitempty"`
	PstBx       *Max16Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 PstBx,omitempty"`
	Room        *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Room,omitempty"`
	PstCd       Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 PstCd"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TwnNm"`
	TwnLctnNm   *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TwnLctnNm,omitempty"`
	DstrctNm    *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CtrySubDvsn"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Ctry"`
	AdrLine     *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 AdrLine,omitempty"`
}

type ProxyAccountIdentification11 struct {
	Tp *ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Tp,omitempty"`
	Id Max256Text               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    *ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cd,omitempty"`
	Prtry *Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prtry,omitempty"`
}

type Purpose2Choice struct {
	Cd    *ExternalPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cd,omitempty"`
	Prtry *Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prtry,omitempty"`
}

type ReferredDocumentInformation7 struct {
	Tp       *ReferredDocumentType4      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Tp,omitempty"`
	Nb       *Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Nb,omitempty"`
	RltdDt   *fednow.ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 RltdDt,omitempty"`
	LineDtls []*DocumentLineInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 LineDtls,omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    *DocumentType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cd,omitempty"`
	Prtry *Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prtry,omitempty"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CdOrPrtry"`
	Issr      *Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Issr,omitempty"`
}

type RemittanceAmount2 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 DuePyblAmt,omitempty"`
	DscntApldAmt      []*DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 DscntApldAmt,omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CdtNoteAmt,omitempty"`
	TaxAmt            []*TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []*DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 RmtdAmt,omitempty"`
}

type RemittanceAmount3 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 DuePyblAmt,omitempty"`
	DscntApldAmt      []*DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 DscntApldAmt,omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CdtNoteAmt,omitempty"`
	TaxAmt            []*TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []*DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 RmtdAmt,omitempty"`
}

type RemittanceInformation161 struct {
	Ustrd *Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Ustrd,omitempty"`
	Strd  []*StructuredRemittanceInformation16 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Strd,omitempty"`
}

type RemittanceLocation71 struct {
	RmtId       *Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 RmtId,omitempty"`
	RmtLctnDtls []*RemittanceLocationData11 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 RmtLctnDtls,omitempty"`
}

type RemittanceLocationData11 struct {
	Mtd        RemittanceLocationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Mtd"`
	ElctrncAdr *Max2048Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 ElctrncAdr,omitempty"`
	PstlAdr    *NameAndAddress161            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 PstlAdr,omitempty"`
}

type ServiceLevel8Choice struct {
	Cd    *ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cd,omitempty"`
	Prtry *Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prtry,omitempty"`
}

type SettlementInstruction71 struct {
	SttlmMtd SettlementMethod1Code1               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 SttlmMtd"`
	ClrSys   ClearingSystemIdentification3Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 ClrSys"`
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []*ReferredDocumentInformation7 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 RfrdDocInf,omitempty"`
	RfrdDocAmt  *RemittanceAmount2              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 RfrdDocAmt,omitempty"`
	CdtrRefInf  *CreditorReferenceInformation2  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CdtrRefInf,omitempty"`
	Invcr       *PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Invcr,omitempty"`
	Invcee      *PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Invcee,omitempty"`
	TaxRmt      *TaxInformation7                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TaxRmt,omitempty"`
	GrnshmtRmt  *Garnishment3                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 GrnshmtRmt,omitempty"`
	AddtlRmtInf []*Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 AddtlRmtInf,omitempty"`
}

type TaxAmount2 struct {
	Rate         *PercentageRate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Rate,omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TaxblBaseAmt,omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TtlAmt,omitempty"`
	Dtls         []*TaxRecordDetails2               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Dtls,omitempty"`
}

type TaxAmountAndType1 struct {
	Tp  *TaxAmountType1Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Amt"`
}

type TaxAmountType1Choice struct {
	Cd    *ExternalTaxAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cd,omitempty"`
	Prtry *Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prtry,omitempty"`
}

type TaxAuthorisation1 struct {
	Titl *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Titl,omitempty"`
	Nm   *Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Nm,omitempty"`
}

type TaxInformation7 struct {
	Cdtr            *TaxParty1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Cdtr,omitempty"`
	Dbtr            *TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Dbtr,omitempty"`
	UltmtDbtr       *TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 UltmtDbtr,omitempty"`
	AdmstnZone      *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 AdmstnZone,omitempty"`
	RefNb           *Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 RefNb,omitempty"`
	Mtd             *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Mtd,omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TtlTaxAmt,omitempty"`
	Dt              *fednow.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Dt,omitempty"`
	SeqNb           *Number                            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 SeqNb,omitempty"`
	Rcrd            []*TaxRecord2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Rcrd,omitempty"`
}

type TaxParty1 struct {
	TaxId  *Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TaxId,omitempty"`
	RegnId *Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 RegnId,omitempty"`
	TaxTp  *Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TaxTp,omitempty"`
}

type TaxParty2 struct {
	TaxId   *Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TaxId,omitempty"`
	RegnId  *Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 RegnId,omitempty"`
	TaxTp   *Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TaxTp,omitempty"`
	Authstn *TaxAuthorisation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Authstn,omitempty"`
}

type TaxPeriod2 struct {
	Yr     *fednow.ISODate       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Yr,omitempty"`
	Tp     *TaxRecordPeriod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Tp,omitempty"`
	FrToDt *DatePeriod2          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 FrToDt,omitempty"`
}

type TaxRecord2 struct {
	Tp       *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Tp,omitempty"`
	Ctgy     *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Ctgy,omitempty"`
	CtgyDtls *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CtgyDtls,omitempty"`
	DbtrSts  *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 DbtrSts,omitempty"`
	CertId   *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 CertId,omitempty"`
	FrmsCd   *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 FrmsCd,omitempty"`
	Prd      *TaxPeriod2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prd,omitempty"`
	TaxAmt   *TaxAmount2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 TaxAmt,omitempty"`
	AddtlInf *Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 AddtlInf,omitempty"`
}

type TaxRecordDetails2 struct {
	Prd *TaxPeriod2                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Prd,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Amt"`
}

// XSD SimpleType declarations

type ActiveCurrencyAndAmountFedNow1SimpleType float64

type ActiveCurrencyCodeFixed string

const ActiveCurrencyCodeFixedUsd ActiveCurrencyCodeFixed = "USD"

type ActiveOrHistoricCurrencyAndAmountSimpleType fednow.Amount

type ActiveOrHistoricCurrencyCode string

type AddressType2Code string

const AddressType2CodeAddr AddressType2Code = "ADDR"
const AddressType2CodePbox AddressType2Code = "PBOX"
const AddressType2CodeHome AddressType2Code = "HOME"
const AddressType2CodeBizz AddressType2Code = "BIZZ"
const AddressType2CodeMlto AddressType2Code = "MLTO"
const AddressType2CodeDlvy AddressType2Code = "DLVY"

type AnyBICDec2014Identifier string

type BICFIDec2014Identifier string

type ChargeBearerType1Code1 string

const ChargeBearerType1Code1Slev ChargeBearerType1Code1 = "SLEV"

type CountryCode string

type CreditDebitCode string

const CreditDebitCodeCrdt CreditDebitCode = "CRDT"
const CreditDebitCodeDbit CreditDebitCode = "DBIT"

type DocumentType3Code string

const DocumentType3CodeRadm DocumentType3Code = "RADM"
const DocumentType3CodeRpin DocumentType3Code = "RPIN"
const DocumentType3CodeFxdr DocumentType3Code = "FXDR"
const DocumentType3CodeDisp DocumentType3Code = "DISP"
const DocumentType3CodePuor DocumentType3Code = "PUOR"
const DocumentType3CodeScor DocumentType3Code = "SCOR"

type DocumentType6Code string

const DocumentType6CodeMsin DocumentType6Code = "MSIN"
const DocumentType6CodeCnfa DocumentType6Code = "CNFA"
const DocumentType6CodeDnfa DocumentType6Code = "DNFA"
const DocumentType6CodeCinv DocumentType6Code = "CINV"
const DocumentType6CodeCren DocumentType6Code = "CREN"
const DocumentType6CodeDebn DocumentType6Code = "DEBN"
const DocumentType6CodeHiri DocumentType6Code = "HIRI"
const DocumentType6CodeSbin DocumentType6Code = "SBIN"
const DocumentType6CodeCmcn DocumentType6Code = "CMCN"
const DocumentType6CodeSoac DocumentType6Code = "SOAC"
const DocumentType6CodeDisp DocumentType6Code = "DISP"
const DocumentType6CodeBold DocumentType6Code = "BOLD"
const DocumentType6CodeVchr DocumentType6Code = "VCHR"
const DocumentType6CodeAroi DocumentType6Code = "AROI"
const DocumentType6CodeTsut DocumentType6Code = "TSUT"
const DocumentType6CodePuor DocumentType6Code = "PUOR"

type Exact4AlphaNumericText string

type ExternalCashAccountType1Code string

type ExternalCashClearingSystem1CodeFixed string

const ExternalCashClearingSystem1CodeFixedFdn ExternalCashClearingSystem1CodeFixed = "FDN"

type ExternalClearingSystemIdentification1CodeFixed string

const ExternalClearingSystemIdentification1CodeFixedUsaba ExternalClearingSystemIdentification1CodeFixed = "USABA"

type ExternalDiscountAmountType1Code string

type ExternalDocumentLineType1Code string

type ExternalGarnishmentType1Code string

type ExternalOrganisationIdentification1Code string

type ExternalPersonIdentification1Code string

type ExternalProxyAccountType1Code string

type ExternalPurpose1Code string

type ExternalServiceLevel1Code string

type ExternalTaxAmountType1Code string

type IBAN2007Identifier string

type LEIIdentifier string

type Max128Text string

type Max140Text string

type Max15NumericTextFixed string

const Max15NumericTextFixed1 Max15NumericTextFixed = "1"

type Max16Text string

type Max2048Text string

type Max256Text string

type Max34Text string

type Max35Text string

type Max4Text string

type Max70Text string

type MessageIdentificationFedNow1 string

type NamePrefix2Code string

const NamePrefix2CodeDoct NamePrefix2Code = "DOCT"
const NamePrefix2CodeMadm NamePrefix2Code = "MADM"
const NamePrefix2CodeMiss NamePrefix2Code = "MISS"
const NamePrefix2CodeMist NamePrefix2Code = "MIST"
const NamePrefix2CodeMiks NamePrefix2Code = "MIKS"

type Number float64

type PercentageRate float64

type PhoneNumber string

type PreferredContactMethod1Code string

const PreferredContactMethod1CodeLett PreferredContactMethod1Code = "LETT"
const PreferredContactMethod1CodeMail PreferredContactMethod1Code = "MAIL"
const PreferredContactMethod1CodePhon PreferredContactMethod1Code = "PHON"
const PreferredContactMethod1CodeFaxx PreferredContactMethod1Code = "FAXX"
const PreferredContactMethod1CodeCell PreferredContactMethod1Code = "CELL"

type RemittanceLocationMethod2Code string

const RemittanceLocationMethod2CodeFaxi RemittanceLocationMethod2Code = "FAXI"
const RemittanceLocationMethod2CodeEdic RemittanceLocationMethod2Code = "EDIC"
const RemittanceLocationMethod2CodeUrid RemittanceLocationMethod2Code = "URID"
const RemittanceLocationMethod2CodeEmal RemittanceLocationMethod2Code = "EMAL"
const RemittanceLocationMethod2CodePost RemittanceLocationMethod2Code = "POST"
const RemittanceLocationMethod2CodeSmsm RemittanceLocationMethod2Code = "SMSM"

type RoutingNumberFRS1 string

type SettlementMethod1Code1 string

const SettlementMethod1Code1Clrg SettlementMethod1Code1 = "CLRG"

type TaxRecordPeriod1Code string

const TaxRecordPeriod1CodeMm01 TaxRecordPeriod1Code = "MM01"
const TaxRecordPeriod1CodeMm02 TaxRecordPeriod1Code = "MM02"
const TaxRecordPeriod1CodeMm03 TaxRecordPeriod1Code = "MM03"
const TaxRecordPeriod1CodeMm04 TaxRecordPeriod1Code = "MM04"
const TaxRecordPeriod1CodeMm05 TaxRecordPeriod1Code = "MM05"
const TaxRecordPeriod1CodeMm06 TaxRecordPeriod1Code = "MM06"
const TaxRecordPeriod1CodeMm07 TaxRecordPeriod1Code = "MM07"
const TaxRecordPeriod1CodeMm08 TaxRecordPeriod1Code = "MM08"
const TaxRecordPeriod1CodeMm09 TaxRecordPeriod1Code = "MM09"
const TaxRecordPeriod1CodeMm10 TaxRecordPeriod1Code = "MM10"
const TaxRecordPeriod1CodeMm11 TaxRecordPeriod1Code = "MM11"
const TaxRecordPeriod1CodeMm12 TaxRecordPeriod1Code = "MM12"
const TaxRecordPeriod1CodeQtr1 TaxRecordPeriod1Code = "QTR1"
const TaxRecordPeriod1CodeQtr2 TaxRecordPeriod1Code = "QTR2"
const TaxRecordPeriod1CodeQtr3 TaxRecordPeriod1Code = "QTR3"
const TaxRecordPeriod1CodeQtr4 TaxRecordPeriod1Code = "QTR4"
const TaxRecordPeriod1CodeHlf1 TaxRecordPeriod1Code = "HLF1"
const TaxRecordPeriod1CodeHlf2 TaxRecordPeriod1Code = "HLF2"

type TrueFalseIndicator bool

type UUIDv4Identifier string
