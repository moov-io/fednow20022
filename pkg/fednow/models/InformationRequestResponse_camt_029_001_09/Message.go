// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: InformationRequestResponse_camt.029.001.09
// Base Message: camt.029.001.09
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7AOZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope
package InformationRequestResponse_camt_029_001_09

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/InformationRequestResponse_camt_029_001_09"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type MessageModel struct {
	AssignmentId                 string                                        `json:"assignment_id,omitempty"`
	CreatedDateTime              time.Time                                     `json:"created_date_time,omitzero"`
	Assignments                  common.Assignments                            `json:"assignments,omitempty"`
	CaseId                       string                                        `json:"case_id,omitempty"`
	Creator                      common.Creator                                `json:"creator,omitempty"`                       // Creator of the message
	Status                       common.InvestigationExecutionConfirmationCode `json:"status,omitempty"`                        // Status of the investigation execution confirmation
	Interbank                    common.PaymentInfomation                      `json:"interbank,omitempty"`                     // Information about the interbank payment
	ResolutionRelatedInformation common.TransactionDetailReference             `json:"resolution_related_information,omitzero"` // Reference to the resolution related information
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:camt.029.001.09"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &InformationRequestResponse_camt_029_001_09.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"AssignmentId",
	"Assignments.Assigner",
	"Assignments.Assignee",
	"CreatedDateTime",
	"CaseId",
	"Creator",
	"Status",
}
var PathMap = map[string]any{
	"RsltnOfInvstgtn.Assgnmt.Id":                                             "AssignmentId",
	"RsltnOfInvstgtn.Assgnmt.CreDtTm":                                        "CreatedDateTime",
	"RsltnOfInvstgtn.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":  "Assignments.Assigner.PaymentSysCode",
	"RsltnOfInvstgtn.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.MmbId":        "Assignments.Assigner.PaymentSysMemberId",
	"RsltnOfInvstgtn.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":  "Assignments.Assignee.PaymentSysCode",
	"RsltnOfInvstgtn.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.MmbId":        "Assignments.Assignee.PaymentSysMemberId",
	"RsltnOfInvstgtn.RslvdCase.Id":                                           "CaseId",
	"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Creator.PaymentSysCode",
	"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.ClrSysMmbId.MmbId":       "Creator.PaymentSysMemberId",
	"RsltnOfInvstgtn.RslvdCase.Cretr.Pty.Nm":                                 "Creator.Name",
	"RsltnOfInvstgtn.RslvdCase.Cretr.Pty.CtctDtls.Nm":                        "Creator.Contact.Name",
	"RsltnOfInvstgtn.RslvdCase.Cretr.Pty.CtctDtls.PhneNb":                    "Creator.Contact.Phone",
	"RsltnOfInvstgtn.RslvdCase.Cretr.Pty.CtctDtls.EmailAdr":                  "Creator.Contact.Email",
	"RsltnOfInvstgtn.RslvdCase.Cretr.Pty.CtctDtls.PrefrdMtd":                 "Creator.Contact.PreferredMethod",
	"RsltnOfInvstgtn.Sts.Conf":                                               "Status",
	// Interbank
	"RsltnOfInvstgtn.CrrctnTx.IntrBk.GrpHdr.MsgId":         "Interbank.OriginalGroupInfo.OriginalMessageIdentification",
	"RsltnOfInvstgtn.CrrctnTx.IntrBk.GrpHdr.MsgNmId":       "Interbank.OriginalGroupInfo.OriginalMessageNameIdentification",
	"RsltnOfInvstgtn.CrrctnTx.IntrBk.GrpHdr.CreDtTm":       "Interbank.OriginalGroupInfo.OriginalCreationDateTime",
	"RsltnOfInvstgtn.CrrctnTx.IntrBk.InstrId":              "Interbank.OriginalTransactionDetail.InstructionIdentification",
	"RsltnOfInvstgtn.CrrctnTx.IntrBk.EndToEndId":           "Interbank.OriginalTransactionDetail.EndToEndIdentification",
	"RsltnOfInvstgtn.CrrctnTx.IntrBk.TxId":                 "Interbank.OriginalTransactionDetail.TransactionId",
	"RsltnOfInvstgtn.CrrctnTx.IntrBk.IntrBkSttlmAmt.Value": "Interbank.SettlementAmount.Amount",
	"RsltnOfInvstgtn.CrrctnTx.IntrBk.IntrBkSttlmAmt.Ccy":   "Interbank.SettlementAmount.Currency",
	"RsltnOfInvstgtn.CrrctnTx.IntrBk.IntrBkSttlmDt":        "Interbank.SettlementDate",
	// Resolution Related Information
	"RsltnOfInvstgtn.RsltnRltdInf.EndToEndId": "ResolutionRelatedInformation.EndToEndIdentification",
	"RsltnOfInvstgtn.RsltnRltdInf.TxId":       "ResolutionRelatedInformation.TransactionId",
}
