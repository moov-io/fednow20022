// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: RequestForPaymentCancellationRequest_camt.055.001.09
// Base Message: camt.055.001.09
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7AWZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

package RequestForPaymentCancellationRequest_camt_055_001_09

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/RequestForPaymentCancellationRequest_camt_055_001_09"
	"github.com/moov-io/fednow20022/pkg/fednow"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type MessageModel struct {
	MessageId                      string                            `json:"message_id,omitempty"`
	CreatedDateTime                time.Time                         `json:"created_date_time,omitempty"`
	Assigners                      common.Assignments                `json:"assigners,omitempty"`
	CaseId                         string                            `json:"case_id,omitempty"`
	Creator                        common.PartyContact               `json:"creator,omitempty"`
	PaymentCancellationId          string                            `json:"payment_cancellation_id,omitempty"`
	GroupInfo                      common.GroupInformation           `json:"group_information,omitempty"`
	TransactionInfo                common.TransactionDetailReference `json:"transaction_information,omitempty"`
	OriginalInstructedAmount       common.CurrencyAndAmount          `json:"original_instructed_amount,omitempty"`
	OriginalRequestedExecutionDate fednow.ISODate                    `json:"original_requested_execution_date,omitempty"`
	Reason                         common.ReturnReason               `json:"reason,omitempty"`
}

func Test() {
	data := MessageModel{}
	doc := RequestForPaymentCancellationRequest_camt_055_001_09.Document{}

	data.MessageId = string(doc.CstmrPmtCxlReq.Assgnmt.Id)
	data.CreatedDateTime = time.Time(doc.CstmrPmtCxlReq.Assgnmt.CreDtTm)
	data.Assigners.Assigner.PaymentSysCode = common.PaymentSystemType(*doc.CstmrPmtCxlReq.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	data.Assigners.Assigner.PaymentSysMemberId = string(doc.CstmrPmtCxlReq.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.MmbId)
	data.Assigners.Assignee.PaymentSysCode = common.PaymentSystemType(*doc.CstmrPmtCxlReq.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	data.Assigners.Assignee.PaymentSysMemberId = string(doc.CstmrPmtCxlReq.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.MmbId)
	data.CaseId = string(doc.CstmrPmtCxlReq.Case.Id)
	data.Creator.Name = string(doc.CstmrPmtCxlReq.Case.Cretr.Pty.Nm)
	data.Creator.PhoneNumber = string(*doc.CstmrPmtCxlReq.Case.Cretr.Pty.CtctDtls.PhneNb)
	data.Creator.EmailAddress = string(*doc.CstmrPmtCxlReq.Case.Cretr.Pty.CtctDtls.EmailAdr)
	data.Creator.PreferredMethod = common.ContactMethod(doc.CstmrPmtCxlReq.Case.Cretr.Pty.CtctDtls.PrefrdMtd)
	data.PaymentCancellationId = string(doc.CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.OrgnlPmtInfId)
	data.GroupInfo.OriginalMessageIdentification = string(doc.CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.OrgnlGrpInf.OrgnlMsgId)
	data.GroupInfo.OriginalMessageNameIdentification = string(doc.CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.OrgnlGrpInf.OrgnlMsgNmId)
	data.GroupInfo.OriginalCreationDateTime = time.Time(doc.CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.OrgnlGrpInf.OrgnlCreDtTm)
	data.TransactionInfo.EndToEndIdentification = string(*doc.CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.TxInf.OrgnlEndToEndId)
	data.OriginalInstructedAmount.Amount = float64(doc.CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.TxInf.OrgnlInstdAmt.Value)
	data.OriginalInstructedAmount.Currency = string(doc.CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.TxInf.OrgnlInstdAmt.Ccy)
	data.OriginalRequestedExecutionDate = fednow.ISODate(*doc.CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.TxInf.OrgnlReqdExctnDt.Dt)
	data.Reason.Code = common.ReturnReasonCode(*doc.CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.TxInf.CxlRsnInf.Rsn.Cd)
	data.Reason.Info = string(*doc.CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.TxInf.CxlRsnInf.AddtlInf)
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:camt.055.001.09"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &RequestForPaymentCancellationRequest_camt_055_001_09.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId",
	"CreatedDateTime",
	"Assigners.Assigner",
	"Assigners.Assignee",
	"CaseId",
	"Creator",
	"PaymentCancellationId",
	"GroupInfo.OriginalMessageIdentification",
	"GroupInfo.OriginalMessageNameIdentification",
	"GroupInfo.OriginalCreationDateTime",
	"OriginalInstructedAmount",
	"OriginalRequestedExecutionDate",
	"Reason",
}
var PathMap = map[string]any{
	"CstmrPmtCxlReq.Assgnmt.Id":                                            "MessageId",
	"CstmrPmtCxlReq.Assgnmt.CreDtTm":                                       "CreatedDateTime",
	"CstmrPmtCxlReq.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Assigners.Assigner.PaymentSysCode",
	"CstmrPmtCxlReq.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.MmbId":       "Assigners.Assigner.PaymentSysMemberId",
	"CstmrPmtCxlReq.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Assigners.Assignee.PaymentSysCode",
	"CstmrPmtCxlReq.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.MmbId":       "Assigners.Assignee.PaymentSysMemberId",
	"CstmrPmtCxlReq.Case.Id":                                               "CaseId",
	"CstmrPmtCxlReq.Case.Cretr.Pty.Nm":                                     "Creator.Name",
	"CstmrPmtCxlReq.Case.Cretr.Pty.CtctDtls.PhneNb":                        "Creator.PhoneNumber",
	"CstmrPmtCxlReq.Case.Cretr.Pty.CtctDtls.EmailAdr":                      "Creator.EmailAddress",
	"CstmrPmtCxlReq.Case.Cretr.Pty.CtctDtls.PrefrdMtd":                     "Creator.PreferredMethod",
	"CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.OrgnlPmtInfId":               "PaymentCancellationId",
	"CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.OrgnlGrpInf.OrgnlMsgId":      "GroupInfo.OriginalMessageIdentification",
	"CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.OrgnlGrpInf.OrgnlMsgNmId":    "GroupInfo.OriginalMessageNameIdentification",
	"CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.OrgnlGrpInf.OrgnlCreDtTm":    "GroupInfo.OriginalCreationDateTime",
	"CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.TxInf.OrgnlEndToEndId":       "TransactionInfo.EndToEndIdentification",
	"CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.TxInf.OrgnlInstdAmt.Value":   "OriginalInstructedAmount.Amount",
	"CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.TxInf.OrgnlInstdAmt.Ccy":     "OriginalInstructedAmount.Currency",
	"CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.TxInf.OrgnlReqdExctnDt.Dt":   "OriginalRequestedExecutionDate",
	"CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.TxInf.CxlRsnInf.Rsn.Cd":      "Reason.Code",
	"CstmrPmtCxlReq.Undrlyg.OrgnlPmtInfAndCxl.TxInf.CxlRsnInf.AddtlInf":    "Reason.Info",
}
