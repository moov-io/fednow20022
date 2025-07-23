package AdditionalPaymentInformation_camt_028_001_09

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/AdditionalPaymentInformation_camt_028_001_09"
	"github.com/moov-io/fednow20022/pkg/fednow"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: AdditionalPaymentInformation_camt.028.001.09
// Base Message: camt.028.001.09
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_qR07VZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

type MessageModel struct {
	MessageId                         string                            `json:"message_id,omitempty"`
	Assignment                        common.Assignments                `json:"assignment,omitempty"`
	CreatedDateTime                   time.Time                         `json:"created_date_time,omitzero"`
	CaseId                            string                            `json:"case_id,omitempty"`
	Creator                           common.Agent                      `json:"creator,omitempty"`
	OriginalGroupInfo                 common.GroupInformation           `json:"original_group_info,omitempty"`
	OriginalTransactionInfo           common.TransactionDetailReference `json:"original_transaction_info,omitempty"`
	OriginalInterbankSettlementAmount common.CurrencyAndAmount          `json:"original_interbank_settlement_amount,omitempty"`
	OriginalInterbankSettlementDate   fednow.ISODate                    `json:"original_interbank_settlement_date,omitempty"`
	Information                       string                            `json:"information,omitempty"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:camt.028.001.09"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &AdditionalPaymentInformation_camt_028_001_09.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId",
	"CreatedDateTime",
	"Assignment.Assignee",
	"Assignment.Assigner",
	"CaseId",
	"OriginalGroupInfo",
	"Creator",
	"Information",
}
var PathMap = map[string]any{
	"AddtlPmtInf.Assgnmt.Id": "MessageId",
	"AddtlPmtInf.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Assignment.Assignee.PaymentSysCode",
	"AddtlPmtInf.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.MmbId":       "Assignment.Assignee.PaymentSysMemberId",
	"AddtlPmtInf.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Assignment.Assigner.PaymentSysCode",
	"AddtlPmtInf.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.MmbId":       "Assignment.Assigner.PaymentSysMemberId",
	"AddtlPmtInf.Assgnmt.CreDtTm":                                       "CreatedDateTime",
	"AddtlPmtInf.Case.Id":                                               "CaseId",
	"AddtlPmtInf.Case.Cretr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":     "Creator.PaymentSysCode",
	"AddtlPmtInf.Case.Cretr.Agt.FinInstnId.ClrSysMmbId.MmbId":           "Creator.PaymentSysMemberId",
	"AddtlPmtInf.Undrlyg.IntrBk.OrgnlGrpInf.OrgnlMsgId":                 "OriginalGroupInfo.OriginalMessageIdentification",
	"AddtlPmtInf.Undrlyg.IntrBk.OrgnlGrpInf.OrgnlMsgNmId":               "OriginalGroupInfo.OriginalMessageNameIdentification",
	"AddtlPmtInf.Undrlyg.IntrBk.OrgnlGrpInf.OrgnlCreDtTm":               "OriginalGroupInfo.OriginalCreationDateTime",
	"AddtlPmtInf.Undrlyg.IntrBk.OrgnlTxId":                              "OriginalTransactionInfo.TransactionId",
	"AddtlPmtInf.Undrlyg.IntrBk.OrgnlInstrId":                           "OriginalTransactionInfo.InstructionIdentification",
	"AddtlPmtInf.Undrlyg.IntrBk.OrgnlEndToEndId":                        "OriginalTransactionInfo.EndToEndIdentification",
	"AddtlPmtInf.Undrlyg.IntrBk.OrgnlIntrBkSttlmAmt.Value":              "OriginalInterbankSettlementAmount.Amount",
	"AddtlPmtInf.Undrlyg.IntrBk.OrgnlIntrBkSttlmAmt.Ccy":                "OriginalInterbankSettlementAmount.Currency",
	"AddtlPmtInf.Undrlyg.IntrBk.OrgnlIntrBkSttlmDt":                     "OriginalInterbankSettlementDate",
	"AddtlPmtInf.Inf.RmtInf.Ustrd":                                      "Information",
}
