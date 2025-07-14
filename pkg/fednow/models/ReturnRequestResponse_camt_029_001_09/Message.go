// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: ReturnRequestResponse_camt.029.001.09
// Base Message: camt.029.001.09
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7AaZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope
package ReturnRequestResponse_camt_029_001_09

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/ReturnRequestResponse_camt_029_001_09"
	"github.com/moov-io/fednow20022/pkg/fednow"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type ResolutionRelatedInformation struct {
	TransactionId             string                   `json:"transaction_id,omitempty"`
	UETR                      string                   `json:"uetr,omitempty"`
	InterbankSettlementAmount common.CurrencyAndAmount `json:"interbank_settlement_amount,omitempty"`
	InterbankSettlementDate   fednow.ISODate           `json:"interbank_settlement_date,omitempty"`
}
type MessageModel struct {
	MessageId       string                                   `json:"message_id,omitempty"`
	CreatedDateTime time.Time                                `json:"created_date_time,omitempty"`
	Assigners       common.Assignments                       `json:"assigners,omitempty"`
	ResolvedCaseId  string                                   `json:"case_id,omitempty"`
	Creator         common.Agent                             `json:"creator,omitempty"`
	Status          common.InvestigationExecutionConfirmCode `json:"status,omitempty"` // Added status field
	GroupInfo       common.GroupInformation                  `json:"group_information,omitempty"`
	TransactionInfo common.TransactionDetailReference        `json:"transaction_information,omitempty"`
	ReturnReason    common.ReturnReason                      `json:"return_reason,omitempty"`
	RelatedInfo     ResolutionRelatedInformation             `json:"related_information,omitempty"`
}

func Test() {
	data := MessageModel{}
	msg := ReturnRequestResponse_camt_029_001_09.Document{}

	data.MessageId = string(msg.RsltnOfInvstgtn.Assgnmt.Id)
	data.CreatedDateTime = time.Time(msg.RsltnOfInvstgtn.Assgnmt.CreDtTm)
	data.Assigners.Assigner.PaymentSysCode = common.PaymentSystemType(*msg.RsltnOfInvstgtn.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	data.Assigners.Assigner.PaymentSysMemberId = string(msg.RsltnOfInvstgtn.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.MmbId)
	data.Assigners.Assignee.PaymentSysCode = common.PaymentSystemType(*msg.RsltnOfInvstgtn.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	data.Assigners.Assignee.PaymentSysMemberId = string(msg.RsltnOfInvstgtn.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.MmbId)
	data.ResolvedCaseId = string(msg.RsltnOfInvstgtn.RslvdCase.Id)
	data.Creator.PaymentSysCode = common.PaymentSystemType(*msg.RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	data.Creator.PaymentSysMemberId = string(msg.RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.ClrSysMmbId.MmbId)
	data.Status = common.InvestigationExecutionConfirmCode(*msg.RsltnOfInvstgtn.Sts.Conf)
	data.GroupInfo.OriginalMessageIdentification = string(msg.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlGrpInf.OrgnlMsgId)
	data.GroupInfo.OriginalMessageNameIdentification = string(msg.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlGrpInf.OrgnlMsgNmId)
	data.GroupInfo.OriginalCreationDateTime = time.Time(msg.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlGrpInf.OrgnlCreDtTm)
	data.TransactionInfo.InstructionIdentification = string(*msg.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlInstrId)
	data.TransactionInfo.EndToEndIdentification = string(*msg.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlEndToEndId)
	data.TransactionInfo.UETR = string(*msg.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlUETR)
	data.RelatedInfo.TransactionId = string(*msg.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.RsltnRltdInf.TxId)
	data.RelatedInfo.UETR = string(*msg.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.RsltnRltdInf.UETR)
	data.RelatedInfo.InterbankSettlementAmount.Amount = float64(msg.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.RsltnRltdInf.IntrBkSttlmAmt.Value)
	data.RelatedInfo.InterbankSettlementAmount.Currency = string(msg.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.RsltnRltdInf.IntrBkSttlmAmt.Ccy)
	data.RelatedInfo.InterbankSettlementDate = fednow.ISODate(*msg.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.RsltnRltdInf.IntrBkSttlmDt)
	data.ReturnReason.Code = common.ReturnReasonCode(*msg.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.CxlStsRsnInf[0].Rsn.Cd)
	data.ReturnReason.Info = string(*msg.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.CxlStsRsnInf[0].AddtlInf)
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:camt.029.001.09"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &ReturnRequestResponse_camt_029_001_09.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId",
	"CreatedDateTime",
	"Assigners.Assigner",
	"Assigners.Assignee",
	"ResolvedCaseId",
	"Creator",
	"Status",
	"GroupInfo.OriginalMessageIdentification",
	"GroupInfo.OriginalMessageNameIdentification",
	"GroupInfo.OriginalCreationDateTime",
}
var PathMap = map[string]any{

	"RsltnOfInvstgtn.Assgnmt.Id":                                             "MessageId",
	"RsltnOfInvstgtn.Assgnmt.CreDtTm":                                        "CreatedDateTime",
	"RsltnOfInvstgtn.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":  "Assigners.Assigner.PaymentSysCode",
	"RsltnOfInvstgtn.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.MmbId":        "Assigners.Assigner.PaymentSysMemberId",
	"RsltnOfInvstgtn.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":  "Assigners.Assignee.PaymentSysCode",
	"RsltnOfInvstgtn.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.MmbId":        "Assigners.Assignee.PaymentSysMemberId",
	"RsltnOfInvstgtn.RslvdCase.Id":                                           "ResolvedCaseId",
	"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Creator.PaymentSysCode",
	"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.ClrSysMmbId.MmbId":       "Creator.PaymentSysMemberId",
	"RsltnOfInvstgtn.Sts.Conf":                                               "Status",
	"RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlGrpInf.OrgnlMsgId":             "GroupInfo.OriginalMessageIdentification",
	"RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlGrpInf.OrgnlMsgNmId":           "GroupInfo.OriginalMessageNameIdentification",
	"RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlGrpInf.OrgnlCreDtTm":           "GroupInfo.OriginalCreationDateTime",
	"RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlInstrId":                       "TransactionInfo.InstructionIdentification",
	"RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlEndToEndId":                    "TransactionInfo.EndToEndIdentification",
	"RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlUETR":                          "TransactionInfo.UETR",
	"RsltnOfInvstgtn.CxlDtls.TxInfAndSts.RsltnRltdInf.TxId":                  "RelatedInfo.TransactionId",
	"RsltnOfInvstgtn.CxlDtls.TxInfAndSts.RsltnRltdInf.UETR":                  "RelatedInfo.UETR",
	"RsltnOfInvstgtn.CxlDtls.TxInfAndSts.RsltnRltdInf.IntrBkSttlmAmt.Value":  "RelatedInfo.InterbankSettlementAmount.Amount",
	"RsltnOfInvstgtn.CxlDtls.TxInfAndSts.RsltnRltdInf.IntrBkSttlmAmt.Ccy":    "RelatedInfo.InterbankSettlementAmount.Currency",
	"RsltnOfInvstgtn.CxlDtls.TxInfAndSts.RsltnRltdInf.IntrBkSttlmDt":         "RelatedInfo.InterbankSettlementDate",
	"RsltnOfInvstgtn.CxlDtls.TxInfAndSts.CxlStsRsnInf[0].Rsn.Cd":             "ReturnReason.Code",
	"RsltnOfInvstgtn.CxlDtls.TxInfAndSts.CxlStsRsnInf[0].AddtlInf":           "ReturnReason.Info",
}
