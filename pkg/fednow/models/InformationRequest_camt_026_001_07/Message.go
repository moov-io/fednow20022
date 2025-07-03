// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: InformationRequest_camt.026.001.07
// Base Message: camt.026.001.07
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7APZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope
package InformationRequest_camt_026_001_07

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/InformationRequest_camt_026_001_07"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type Justification struct {
	MissingInformation           common.MissingCodeAndInfo `json:"missing_information,omitempty"`            // Code and information about the missing or incorrect information
	IncorrectInformation         common.MissingCodeAndInfo `json:"incorrect_information,omitempty"`          // Code and
	PossibleDuplicateInstruction bool                      `json:"possible_duplicate_instruction,omitempty"` // Code and information about the possible duplicate instruction
}

type MessageModel struct {
	AssignmentId    string                   `json:"assignment_id,omitempty"`
	CreatedDateTime time.Time                `json:"created_date_time,omitempty"`
	Assignments     common.Assignments       `json:"assignments,omitempty"`
	CaseId          string                   `json:"case_id,omitempty"`
	Creator         common.Creator           `json:"creator,omitempty"`       // Creator of the message
	Initiation      common.PaymentInfomation `json:"initiation,omitempty"`    // Information about the payment initiation
	Interbank       common.PaymentInfomation `json:"interbank,omitempty"`     // Information about the interbank payment
	Justification   Justification            `json:"justification,omitempty"` // Justification for the information request
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:camt.026.001.07"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &InformationRequest_camt_026_001_07.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"AssignmentId",
	"Assignments",
	"Assignments.Assigner",
	"Assignments.Assignee",
	"CreatedDateTime",
	"CaseId",
	"Creator",
	"Justification",
}
var PathMap = map[string]any{
	"UblToApply.Assgnmt.Id":      "AssignmentId",
	"UblToApply.Assgnmt.CreDtTm": "CreatedDateTime",
	"UblToApply.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Assignments.Assigner.PaymentSysCode",
	"UblToApply.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.MmbId":       "Assignments.Assigner.PaymentSysMemberId",
	"UblToApply.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Assignments.Assignee.PaymentSysCode",
	"UblToApply.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.MmbId":       "Assignments.Assignee.PaymentSysMemberId",
	"UblToApply.Case.Id":           "CaseId",
	"UblToApply.Case.Cretr.Pty.Nm": "Creator.Name",
	"UblToApply.Case.Cretr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Creator.PaymentSysCode",
	"UblToApply.Case.Cretr.Agt.FinInstnId.ClrSysMmbId.MmbId":       "Creator.PaymentSysMemberId",
	"UblToApply.Case.Cretr.Pty.CtctDtls.Nm":                        "Creator.Contact.Name",
	"UblToApply.Case.Cretr.Pty.CtctDtls.PhneNb":                    "Creator.Contact.Phone",
	"UblToApply.Case.Cretr.Pty.CtctDtls.EmailAdr":                  "Creator.Contact.Email",
	"UblToApply.Case.Cretr.Pty.CtctDtls.PrefrdMtd":                 "Creator.Contact.PreferredMethod",

	// Initiation
	"UblToApply.Undrlyg.Initn.OrgnlGrpInf.OrgnlMsgId":   "Initiation.OriginalGroupInfo.OriginalMessageIdentification",
	"UblToApply.Undrlyg.Initn.OrgnlGrpInf.OrgnlMsgNmId": "Initiation.OriginalGroupInfo.OriginalMessageNameIdentification",
	"UblToApply.Undrlyg.Initn.OrgnlGrpInf.OrgnlCreDtTm": "Initiation.OriginalGroupInfo.OriginalCreationDateTime",
	"UblToApply.Undrlyg.Initn.OrgnlInstrId":             "Initiation.OriginalTransactionDetail.InstructionIdentification",
	"UblToApply.Undrlyg.Initn.OrgnlEndToEndId":          "Initiation.OriginalTransactionDetail.EndToEndIdentification",
	"UblToApply.Undrlyg.Initn.OrgnlPmtInfId":            "Initiation.OriginalTransactionDetail.MessageIdentification",
	"UblToApply.Undrlyg.Initn.OrgnlInstdAmt.Value":      "Initiation.SettlementAmount.Amount",
	"UblToApply.Undrlyg.Initn.OrgnlInstdAmt.Ccy":        "Initiation.SettlementAmount.Currency",
	"UblToApply.Undrlyg.Initn.ReqdExctnDt.Dt":           "Initiation.SettlementDate",

	// Interbank
	"UblToApply.Undrlyg.IntrBk.OrgnlGrpInf.OrgnlMsgId":    "Interbank.OriginalGroupInfo.OriginalMessageIdentification",
	"UblToApply.Undrlyg.IntrBk.OrgnlGrpInf.OrgnlMsgNmId":  "Interbank.OriginalGroupInfo.OriginalMessageNameIdentification",
	"UblToApply.Undrlyg.IntrBk.OrgnlGrpInf.OrgnlCreDtTm":  "Interbank.OriginalGroupInfo.OriginalCreationDateTime",
	"UblToApply.Undrlyg.IntrBk.OrgnlInstrId":              "Interbank.OriginalTransactionDetail.InstructionIdentification",
	"UblToApply.Undrlyg.IntrBk.OrgnlEndToEndId":           "Interbank.OriginalTransactionDetail.EndToEndIdentification",
	"UblToApply.Undrlyg.IntrBk.OrgnlTxId":                 "Interbank.OriginalTransactionDetail.MessageIdentification",
	"UblToApply.Undrlyg.IntrBk.OrgnlIntrBkSttlmAmt.Value": "Interbank.SettlementAmount.Amount",
	"UblToApply.Undrlyg.IntrBk.OrgnlIntrBkSttlmAmt.Ccy":   "Interbank.SettlementAmount.Currency",
	"UblToApply.Undrlyg.IntrBk.OrgnlIntrBkSttlmDt":        "Interbank.SettlementDate",

	// Justification
	"UblToApply.Justfn.MssngOrIncrrctInf.IncrrctInf[0].Cd":              "Justification.IncorrectInformation.Code",
	"UblToApply.Justfn.MssngOrIncrrctInf.IncrrctInf[0].AddtlIncrrctInf": "Justification.IncorrectInformation.Info",
	"UblToApply.Justfn.MssngOrIncrrctInf.MssngInf[0].Cd":                "Justification.MissingInformation.Code",
	"UblToApply.Justfn.MssngOrIncrrctInf.MssngInf[0].AddtlMssngInf":     "Justification.MissingInformation.Info",
	"UblToApply.Justfn.PssblDplctInstr":                                 "Justification.PossibleDuplicateInstruction",
}
