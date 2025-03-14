// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10
package PaymentReturn_pacs_004_001_10

import (
	"encoding/xml"

	"github.com/moov-io/fednow20022/pkg/fednow"
)

// XSD Elements

type Document struct {
	XMLName xml.Name

	PmtRtr PaymentReturnV10 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PmtRtr"`
}

// XSD ComplexType declarations

type AccountIdentification4Choice1 struct {
	IBAN *IBAN2007Identifier             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 IBAN,omitempty"`
	Othr *GenericAccountIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Othr,omitempty"`
}

type ActiveCurrencyAndAmountFedNow1 struct {
	Value ActiveCurrencyAndAmountFedNow1SimpleType `xml:",chardata"`
	Ccy   ActiveCurrencyCodeFixed                  `xml:"Ccy,attr"`
}

type ActiveOrHistoricCurrencyAndAmountFedNow1 struct {
	Value ActiveOrHistoricCurrencyAndAmountFedNow1SimpleType `xml:",chardata"`
	Ccy   ActiveCurrencyCodeFixed                            `xml:"Ccy,attr"`
}

type BranchAndFinancialInstitutionIdentification61 struct {
	FinInstnId FinancialInstitutionIdentification181 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 FinInstnId"`
}

type BranchAndFinancialInstitutionIdentification62 struct {
	FinInstnId FinancialInstitutionIdentification182 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 FinInstnId"`
}

type CashAccount381 struct {
	Id   AccountIdentification4Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Id"`
	Tp   *CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Tp,omitempty"`
	Prxy *ProxyAccountIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    *ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cd,omitempty"`
	Prtry *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Prtry,omitempty"`
}

type ClearingSystemIdentification2Choice1 struct {
	Cd *ExternalClearingSystemIdentification1CodeFixed `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cd,omitempty"`
}

type ClearingSystemIdentification3Choice1 struct {
	Cd *ExternalCashClearingSystem1CodeFixed `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cd,omitempty"`
}

type ClearingSystemMemberIdentification21 struct {
	ClrSysId ClearingSystemIdentification2Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 ClrSysId"`
	MmbId    RoutingNumberFRS1                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 MmbId"`
}

type Contact41 struct {
	Nm        *Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Nm,omitempty"`
	PhneNb    *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PhneNb,omitempty"`
	MobNb     *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 MobNb,omitempty"`
	EmailAdr  *Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 EmailAdr,omitempty"`
	Dept      *Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Dept,omitempty"`
	PrefrdMtd PreferredContactMethod1Code1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PrefrdMtd"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     fednow.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 BirthDt"`
	PrvcOfBirth *Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CityOfBirth"`
	CtryOfBirth CountryCode    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CtryOfBirth"`
}

type FinancialInstitutionIdentification181 struct {
	ClrSysMmbId ClearingSystemMemberIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 ClrSysMmbId"`
}

type FinancialInstitutionIdentification182 struct {
	BICFI       *BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 ClrSysMmbId"`
	LEI         *LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 LEI,omitempty"`
	Nm          *Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Nm,omitempty"`
	PstlAdr     *PostalAddress241                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PstlAdr,omitempty"`
}

type GenericAccountIdentification11 struct {
	Id Max34Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Id"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 SchmeNm,omitempty"`
	Issr    *Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 SchmeNm,omitempty"`
	Issr    *Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Issr,omitempty"`
}

type GroupHeader901 struct {
	MsgId    MessageIdentificationFedNow1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 MsgId"`
	CreDtTm  fednow.ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CreDtTm"`
	NbOfTxs  Max15NumericTextFixed        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 NbOfTxs"`
	SttlmInf SettlementInstruction71      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 SttlmInf"`
}

type LocalInstrument2Choice1 struct {
	Prtry *Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Prtry,omitempty"`
}

type OrganisationIdentification291 struct {
	AnyBIC *AnyBICDec2014Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 AnyBIC,omitempty"`
	LEI    *LEIIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 LEI,omitempty"`
	Othr   *GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    *ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cd,omitempty"`
	Prtry *Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Prtry,omitempty"`
}

type OriginalGroupInformation291 struct {
	OrgnlMsgId   Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlMsgId"`
	OrgnlMsgNmId MessageNameIdentificationFRS1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlMsgNmId"`
	OrgnlCreDtTm fednow.ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlCreDtTm"`
}

type OriginalTransactionReference321 struct {
	PmtTpInf PaymentTypeInformation271 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PmtTpInf"`
}

type Party38Choice1 struct {
	OrgId  *OrganisationIdentification291 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgId,omitempty"`
	PrvtId *PersonIdentification131       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PrvtId,omitempty"`
}

type Party40Choice1 struct {
	Pty *PartyIdentification1351 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Pty,omitempty"`
}

type Party40Choice2 struct {
	Pty *PartyIdentification1351                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Pty,omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification62 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Agt,omitempty"`
}

type PartyIdentification1351 struct {
	Nm        Max140Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Nm"`
	PstlAdr   *PostalAddress241 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PstlAdr,omitempty"`
	Id        *Party38Choice1   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Id,omitempty"`
	CtryOfRes *CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CtryOfRes,omitempty"`
}

type PartyIdentification1352 struct {
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Nm"`
	Id       *Party38Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Id,omitempty"`
	CtctDtls *Contact41      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CtctDtls,omitempty"`
}

type PaymentReturnReason61 struct {
	Orgtr    *PartyIdentification1352 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Orgtr,omitempty"`
	Rsn      ReturnReason5Choice1     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Rsn"`
	AddtlInf *Max105Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 AddtlInf,omitempty"`
}

type PaymentReturnV10 struct {
	GrpHdr GroupHeader901         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 GrpHdr"`
	TxInf  PaymentTransaction1181 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 TxInf"`
}

type PaymentTransaction1181 struct {
	RtrId               *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 RtrId,omitempty"`
	OrgnlGrpInf         OriginalGroupInformation291                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlGrpInf"`
	OrgnlInstrId        *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId     *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlEndToEndId,omitempty"`
	OrgnlTxId           *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlTxId,omitempty"`
	OrgnlUETR           *UUIDv4Identifier                             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlUETR,omitempty"`
	OrgnlIntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmountFedNow1      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlIntrBkSttlmAmt"`
	OrgnlIntrBkSttlmDt  fednow.ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlIntrBkSttlmDt"`
	RtrdIntrBkSttlmAmt  ActiveCurrencyAndAmountFedNow1                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 RtrdIntrBkSttlmAmt"`
	IntrBkSttlmDt       fednow.ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 IntrBkSttlmDt"`
	CompstnAmt          *ActiveOrHistoricCurrencyAndAmountFedNow1     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CompstnAmt,omitempty"`
	ChrgBr              ChargeBearerType1Code1                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 ChrgBr"`
	InstgAgt            BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 InstgAgt"`
	InstdAgt            BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 InstdAgt"`
	RtrChain            TransactionParties81                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 RtrChain"`
	RtrRsnInf           PaymentReturnReason61                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 RtrRsnInf"`
	OrgnlTxRef          OriginalTransactionReference321               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlTxRef"`
}

type PaymentTypeInformation271 struct {
	LclInstrm LocalInstrument2Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 LclInstrm"`
}

type PersonIdentification131 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 DtAndPlcOfBirth,omitempty"`
	Othr            *GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    *ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cd,omitempty"`
	Prtry *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Prtry,omitempty"`
}

type PostalAddress241 struct {
	Dept        *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Dept,omitempty"`
	SubDept     *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 SubDept,omitempty"`
	StrtNm      *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 StrtNm,omitempty"`
	BldgNb      *Max16Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 BldgNb,omitempty"`
	BldgNm      *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 BldgNm,omitempty"`
	Flr         *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Flr,omitempty"`
	PstBx       *Max16Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PstBx,omitempty"`
	Room        *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Room,omitempty"`
	PstCd       Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PstCd"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 TwnNm"`
	TwnLctnNm   *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 TwnLctnNm,omitempty"`
	DstrctNm    *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CtrySubDvsn"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Ctry"`
	AdrLine     *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 AdrLine,omitempty"`
}

type ProxyAccountIdentification11 struct {
	Tp *ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Tp,omitempty"`
	Id Max256Text               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    *ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cd,omitempty"`
	Prtry *Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Prtry,omitempty"`
}

type ReturnReason5Choice1 struct {
	Cd *ExternalReturnReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cd,omitempty"`
}

type SettlementInstruction71 struct {
	SttlmMtd SettlementMethod1Code1               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 SttlmMtd"`
	ClrSys   ClearingSystemIdentification3Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 ClrSys"`
}

type TransactionParties81 struct {
	UltmtDbtr *Party40Choice1                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 UltmtDbtr,omitempty"`
	Dbtr      Party40Choice2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Dbtr"`
	DbtrAcct  *CashAccount381                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 DbtrAcct,omitempty"`
	InitgPty  *Party40Choice1                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 InitgPty,omitempty"`
	DbtrAgt   *BranchAndFinancialInstitutionIdentification62 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 DbtrAgt,omitempty"`
	CdtrAgt   BranchAndFinancialInstitutionIdentification62  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CdtrAgt"`
	Cdtr      Party40Choice1                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cdtr"`
	CdtrAcct  CashAccount381                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CdtrAcct"`
	UltmtCdtr *Party40Choice1                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 UltmtCdtr,omitempty"`
}

// XSD SimpleType declarations

type ActiveCurrencyAndAmountFedNow1SimpleType float64

type ActiveCurrencyCodeFixed string

const ActiveCurrencyCodeFixedUsd ActiveCurrencyCodeFixed = "USD"

type ActiveOrHistoricCurrencyAndAmountFedNow1SimpleType float64

type AnyBICDec2014Identifier string

type BICFIDec2014Identifier string

type ChargeBearerType1Code1 string

const ChargeBearerType1Code1Slev ChargeBearerType1Code1 = "SLEV"

type CountryCode string

type ExternalCashAccountType1Code string

type ExternalCashClearingSystem1CodeFixed string

const ExternalCashClearingSystem1CodeFixedFdn ExternalCashClearingSystem1CodeFixed = "FDN"

type ExternalClearingSystemIdentification1CodeFixed string

const ExternalClearingSystemIdentification1CodeFixedUsaba ExternalClearingSystemIdentification1CodeFixed = "USABA"

type ExternalOrganisationIdentification1Code string

type ExternalPersonIdentification1Code string

type ExternalProxyAccountType1Code string

type ExternalReturnReason1Code string

type IBAN2007Identifier string

type LEIIdentifier string

type Max105Text string

type Max140Text string

type Max15NumericTextFixed string

const Max15NumericTextFixed1 Max15NumericTextFixed = "1"

type Max16Text string

type Max2048Text string

type Max256Text string

type Max34Text string

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

type SettlementMethod1Code1 string

const SettlementMethod1Code1Clrg SettlementMethod1Code1 = "CLRG"

type UUIDv4Identifier string
