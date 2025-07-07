// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: FinancialInstitutionCreditTransfer_pacs.009.001.08
// Base Message: pacs.009.001.08
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7ANZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope
package FinancialInstitutionCreditTransfer_pacs_009_001_08

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/FinancialInstitutionCreditTransfer_pacs_009_001_08"
	"github.com/moov-io/fednow20022/pkg/fednow"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type MessageModel struct {
	MessageId                 string                            `json:"message_id,omitempty"`
	CreatedDateTime           time.Time                         `json:"created_date_time,omitempty"`
	NumberOfTransactions      string                            `json:"number_of_transactions,omitempty"`
	SettlementInfo            common.SettlementInformation      `json:"settlement_info,omitempty"`
	PaymentIdentification     common.TransactionDetailReference `json:"payment_identification,omitempty"`
	PaymentType               common.PaymentTypeInfo            `json:"payment_type,omitempty"`
	InterbankSettlementAmount common.CurrencyAndAmount          `json:"interbank_settlement_amount,omitempty"`
	InterbankSettlementDate   fednow.ISODate                    `json:"interbank_settlement_date,omitempty"`
	Agents                    common.RelatedAgents              `json:"agents,omitempty"`
	Debtor                    common.TransactionParty           `json:"debtor,omitempty"`
	Creditor                  common.TransactionParty           `json:"creditor,omitempty"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &FinancialInstitutionCreditTransfer_pacs_009_001_08.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId",
	"CreatedDateTime",
	"NumberOfTransactions",
	"SettlementInfo.Method",
	"SettlementInfo.Service",
	"PaymentIdentification.EndToEndIdentification",
	"PaymentType.LocalInstrument",
	"InterbankSettlementAmount",
	"InterbankSettlementDate",
	"Agents.InstructingAgent.PaymentSysCode",
	"Agents.InstructingAgent.PaymentSysMemberId",
	"Agents.InstructedAgent.PaymentSysCode",
	"Agents.InstructedAgent.PaymentSysMemberId",
	"Debtor.PartyAgent.PaymentSysCode",
	"Debtor.PartyAgent.PaymentSysMemberId",
	"Creditor.PartyAgent.PaymentSysCode",
	"Creditor.PartyAgent.PaymentSysMemberId",
}
var PathMap = map[string]any{
	"FICdtTrf.GrpHdr.MsgId":                                            "MessageId",
	"FICdtTrf.GrpHdr.CreDtTm":                                          "CreatedDateTime",
	"FICdtTrf.GrpHdr.NbOfTxs":                                          "NumberOfTransactions",
	"FICdtTrf.GrpHdr.SttlmInf.SttlmMtd":                                "SettlementInfo.Method",
	"FICdtTrf.GrpHdr.SttlmInf.ClrSys.Cd":                               "SettlementInfo.Service",
	"FICdtTrf.CdtTrfTxInf.PmtId.InstrId":                               "PaymentIdentification.InstructionIdentification",
	"FICdtTrf.CdtTrfTxInf.PmtId.EndToEndId":                            "PaymentIdentification.EndToEndIdentification",
	"FICdtTrf.CdtTrfTxInf.PmtId.TxId":                                  "PaymentIdentification.TransactionId",
	"FICdtTrf.CdtTrfTxInf.PmtTpInf.SvcLvl.Cd":                          "PaymentType.ServiceLevel",
	"FICdtTrf.CdtTrfTxInf.PmtTpInf.LclInstrm.Prtry":                    "PaymentType.LocalInstrument",
	"FICdtTrf.CdtTrfTxInf.IntrBkSttlmAmt.Value":                        "InterbankSettlementAmount.Amount",
	"FICdtTrf.CdtTrfTxInf.IntrBkSttlmAmt.Ccy":                          "InterbankSettlementAmount.Currency",
	"FICdtTrf.CdtTrfTxInf.IntrBkSttlmDt":                               "InterbankSettlementDate",
	"FICdtTrf.CdtTrfTxInf.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Agents.InstructingAgent.PaymentSysCode",
	"FICdtTrf.CdtTrfTxInf.InstgAgt.FinInstnId.ClrSysMmbId.MmbId":       "Agents.InstructingAgent.PaymentSysMemberId",
	"FICdtTrf.CdtTrfTxInf.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Agents.InstructedAgent.PaymentSysCode",
	"FICdtTrf.CdtTrfTxInf.InstdAgt.FinInstnId.ClrSysMmbId.MmbId":       "Agents.InstructedAgent.PaymentSysMemberId",
	"FICdtTrf.CdtTrfTxInf.Dbtr.FinInstnId.Nm":                          "Debtor.PartyName",
	"FICdtTrf.CdtTrfTxInf.Dbtr.FinInstnId.ClrSysMmbId.ClrSysId.Cd":     "Debtor.PartyAgent.PaymentSysCode",
	"FICdtTrf.CdtTrfTxInf.Dbtr.FinInstnId.ClrSysMmbId.MmbId":           "Debtor.PartyAgent.PaymentSysMemberId",
	"FICdtTrf.CdtTrfTxInf.Cdtr.FinInstnId.Nm":                          "Creditor.PartyName",
	"FICdtTrf.CdtTrfTxInf.Cdtr.FinInstnId.ClrSysMmbId.ClrSysId.Cd":     "Creditor.PartyAgent.PaymentSysCode",
	"FICdtTrf.CdtTrfTxInf.Cdtr.FinInstnId.ClrSysMmbId.MmbId":           "Creditor.PartyAgent.PaymentSysMemberId",
}
