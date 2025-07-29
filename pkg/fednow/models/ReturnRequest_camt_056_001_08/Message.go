// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: ReturnRequest_camt.056.001.08
// Base Message: camt.056.001.08
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7AbZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

package ReturnRequest_camt_056_001_08

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/ReturnRequest_camt_056_001_08"
	"github.com/moov-io/fednow20022/pkg/fednow"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type CancellationReasonInfo struct {
	Contact common.PartyContact `json:"code,omitzero"`   // Cancellation reason code
	Reason  common.ReturnReason `json:"reason,omitzero"` // Reason for cancellation
}
type Creator struct {
	Agent common.Agent        `json:"agent,omitzero"` // Agent details
	Party common.PartyContact `json:"party,omitzero"` // Party details
}
type MessageModel struct {
	MessageId                         string                            `json:"message_id,omitempty"`
	CreatedDateTime                   time.Time                         `json:"created_date_time,omitzero"`
	Assignment                        common.Assignments                `json:"assignment,omitempty"`                           // Assignment details
	CaseId                            string                            `json:"case_id,omitempty"`                              // Case ID
	CaseCreator                       Creator                           `json:"case_creator,omitempty"`                         // Case creator details
	OriginalGroup                     common.GroupInformation           `json:"original_group,omitempty"`                       // Original group information
	OriginalTransaction               common.TransactionDetailReference `json:"original_transaction,omitempty"`                 // Original transaction details
	OriginalInterbankSettlementAmount common.CurrencyAndAmount          `json:"original_interbank_settlement_amount,omitempty"` // Original interbank settlement amount
	OriginalInterbankSettlementDate   fednow.ISODate                    `json:"original_interbank_settlement_date,omitempty"`   // Original interbank settlement date
	CancelReason                      CancellationReasonInfo            `json:"cancel_reason,omitzero"`                         // Cancellation reason information
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:camt.056.001.08"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &ReturnRequest_camt_056_001_08.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId",
	"CreatedDateTime",
	"Assignment.Assigner",
	"Assignment.Assignee",
	"CaseId",
	"CaseCreator",
	"OriginalGroup.OriginalMessageIdentification",
	"OriginalGroup.OriginalMessageNameIdentification",
	"OriginalGroup.OriginalCreationDateTime",
	"OriginalInterbankSettlementAmount",
	"OriginalInterbankSettlementDate",
	"CancelReason",
}
var PathMap = map[string]any{
	"FIToFIPmtCxlReq.Assgnmt.Id":                                            "MessageId",
	"FIToFIPmtCxlReq.Assgnmt.CreDtTm":                                       "CreatedDateTime",
	"FIToFIPmtCxlReq.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Assignment.Assigner.PaymentSysCode",
	"FIToFIPmtCxlReq.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.MmbId":       "Assignment.Assigner.PaymentSysMemberId",
	"FIToFIPmtCxlReq.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Assignment.Assignee.PaymentSysCode",
	"FIToFIPmtCxlReq.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.MmbId":       "Assignment.Assignee.PaymentSysMemberId",
	"FIToFIPmtCxlReq.Case.Id":                                               "CaseId",
	"FIToFIPmtCxlReq.Case.Cretr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":     "CaseCreator.Agent.PaymentSysCode",
	"FIToFIPmtCxlReq.Case.Cretr.Agt.FinInstnId.ClrSysMmbId.MmbId":           "CaseCreator.Agent.PaymentSysMemberId",
	"FIToFIPmtCxlReq.Case.Cretr.Pty.Nm":                                     "CaseCreator.Party.Name",
	"FIToFIPmtCxlReq.Case.Cretr.Pty.CtctDtls.PhneNb":                        "CaseCreator.Party.PhoneNumber",
	"FIToFIPmtCxlReq.Case.Cretr.Pty.CtctDtls.EmailAdr":                      "CaseCreator.Party.EmailAddress",
	"FIToFIPmtCxlReq.Case.Cretr.Pty.CtctDtls.MobNb":                         "CaseCreator.Party.MobileNumber",
	"FIToFIPmtCxlReq.Case.Cretr.Pty.CtctDtls.Dept":                          "CaseCreator.Party.Department",
	"FIToFIPmtCxlReq.Case.Cretr.Pty.CtctDtls.PrefrdMtd":                     "CaseCreator.Party.PreferredMethod",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlGrpInf.OrgnlMsgId":                  "OriginalGroup.OriginalMessageIdentification",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlGrpInf.OrgnlMsgNmId":                "OriginalGroup.OriginalMessageNameIdentification",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlGrpInf.OrgnlCreDtTm":                "OriginalGroup.OriginalCreationDateTime",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlInstrId":                            "OriginalTransaction.InstructionIdentification",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlEndToEndId":                         "OriginalTransaction.EndToEndIdentification",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlUETR":                               "OriginalTransaction.UETR",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlIntrBkSttlmAmt.Value":               "OriginalInterbankSettlementAmount.Amount",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlIntrBkSttlmAmt.Ccy":                 "OriginalInterbankSettlementAmount.Currency",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlIntrBkSttlmDt":                      "OriginalInterbankSettlementDate",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.CxlRsnInf.Orgtr.Nm":                      "CancelReason.Contact.Name",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.CxlRsnInf.Orgtr.CtctDtls.PhneNb":         "CancelReason.Contact.PhoneNumber",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.CxlRsnInf.Orgtr.CtctDtls.EmailAdr":       "CancelReason.Contact.EmailAddress",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.CxlRsnInf.Orgtr.CtctDtls.MobNb":          "CancelReason.Contact.MobileNumber",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.CxlRsnInf.Orgtr.CtctDtls.Dept":           "CancelReason.Contact.Department",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.CxlRsnInf.Orgtr.CtctDtls.PrefrdMtd":      "CancelReason.Contact.PreferredMethod",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.CxlRsnInf.Rsn.Cd":                        "CancelReason.Reason.Code",
	"FIToFIPmtCxlReq.Undrlyg.TxInf.CxlRsnInf.AddtlInf":                      "CancelReason.Reason.Info",
}
