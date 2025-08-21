// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: RequestForPaymentCancellationRequestResponse_camt.029.001.09
// Base Message: camt.029.001.09
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7AVZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope
package RequestForPaymentCancellationRequestResponse_camt_029_001_09

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/RequestForPaymentCancellationRequestResponse_camt_029_001_09"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type CancellationStatusReason struct {
	OriginatorName string                `json:"originator_name,omitempty"`
	ReasonCode     common.ReasonCodeType `json:"reason_code,omitempty"`
}

type MessageModel struct {
	MessageId             string                             `json:"message_id,omitempty"`
	CreatedDateTime       time.Time                          `json:"created_date_time,omitzero"`
	Assigners             common.Assignments                 `json:"assigners,omitempty"`
	ResolvedCaseId        string                             `json:"case_id,omitempty"`
	Creator               common.PartyContact                `json:"creator,omitempty"`
	Status                common.CancellationRequestResponse `json:"status,omitempty"` // Added status field
	PaymentCancellationId string                             `json:"payment_cancellation_id,omitempty"`
	GroupInfo             common.GroupInformation            `json:"group_information,omitzero"`
	TransactionInfo       common.TransactionDetailReference  `json:"transaction_information,omitzero"`
	Reason                CancellationStatusReason           `json:"reason,omitzero"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:camt.029.001.09"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &RequestForPaymentCancellationRequestResponse_camt_029_001_09.Document{
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
	"PaymentCancellationId",
	"GroupInfo.OriginalMessageIdentification",
	"GroupInfo.OriginalMessageNameIdentification",
	"GroupInfo.OriginalCreationDateTime",
}
var PathMap = map[string]any{
	"RsltnOfInvstgtn.Assgnmt.Id":                                                     "MessageId",
	"RsltnOfInvstgtn.Assgnmt.CreDtTm":                                                "CreatedDateTime",
	"RsltnOfInvstgtn.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":          "Assigners.Assigner.PaymentSysCode",
	"RsltnOfInvstgtn.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.MmbId":                "Assigners.Assigner.PaymentSysMemberId",
	"RsltnOfInvstgtn.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":          "Assigners.Assignee.PaymentSysCode",
	"RsltnOfInvstgtn.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.MmbId":                "Assigners.Assignee.PaymentSysMemberId",
	"RsltnOfInvstgtn.RslvdCase.Id":                                                   "ResolvedCaseId",
	"RsltnOfInvstgtn.RslvdCase.Cretr.Pty.Nm":                                         "Creator.Name",
	"RsltnOfInvstgtn.RslvdCase.Cretr.Pty.CtctDtls.PhneNb":                            "Creator.PhoneNumber",
	"RsltnOfInvstgtn.RslvdCase.Cretr.Pty.CtctDtls.EmailAdr":                          "Creator.EmailAddress",
	"RsltnOfInvstgtn.RslvdCase.Cretr.Pty.CtctDtls.PrefrdMtd":                         "Creator.PreferredMethod",
	"RsltnOfInvstgtn.Sts.Conf":                                                       "Status",
	"RsltnOfInvstgtn.CxlDtls.OrgnlPmtInfAndSts.OrgnlPmtInfId":                        "PaymentCancellationId",
	"RsltnOfInvstgtn.CxlDtls.OrgnlPmtInfAndSts.OrgnlGrpInf.OrgnlMsgId":               "GroupInfo.OriginalMessageIdentification",
	"RsltnOfInvstgtn.CxlDtls.OrgnlPmtInfAndSts.OrgnlGrpInf.OrgnlMsgNmId":             "GroupInfo.OriginalMessageNameIdentification",
	"RsltnOfInvstgtn.CxlDtls.OrgnlPmtInfAndSts.OrgnlGrpInf.OrgnlCreDtTm":             "GroupInfo.OriginalCreationDateTime",
	"RsltnOfInvstgtn.CxlDtls.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlEndToEndId":          "TransactionInfo.EndToEndIdentification",
	"RsltnOfInvstgtn.CxlDtls.OrgnlPmtInfAndSts.TxInfAndSts.CxlStsRsnInf[0].Orgtr.Nm": "Reason.OriginatorName",
	"RsltnOfInvstgtn.CxlDtls.OrgnlPmtInfAndSts.TxInfAndSts.CxlStsRsnInf[0].Rsn.Cd":   "Reason.ReasonCode",
}
