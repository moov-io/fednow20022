// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: RequestForPaymentResponse_pain.014.001.07
// Base Message: pain.014.001.07
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7AXZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope
package RequestForPaymentResponse_pain_014_001_07

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/pain_014_001_07"
	"github.com/moov-io/fednow20022/pkg/fednow"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type PaymentConditionStatus struct {
	AcceptedAmount    common.CurrencyAndAmount `json:"accepted_amount,omitempty"`    // Accepted Amount
	GuaranteedPayment bool                     `json:"guaranteed_payment,omitempty"` // Guaranteed Payment
	EarlyPayment      bool                     `json:"early_payment,omitempty"`      // Early Payment
}
type TransactionReference struct {
	Amount                 common.CurrencyAndAmount `json:"amount,omitempty"`                   // Amount
	RequestedExecutionDate fednow.ISODate           `json:"requested_execution_date,omitempty"` // Requested Execution Date
	Creditor               common.TransactionParty  `json:"creditor,omitempty"`                 // Creditor
}
type MessageModel struct {
	MessageId                    string                            `json:"message_id,omitempty"`
	CreatedDateTime              time.Time                         `json:"created_date_time,omitzero"`
	InitiatingPartyName          string                            `json:"initiating_party_name,omitempty"`
	Agents                       common.DebtorAndCreditorAgent     `json:"agents,omitempty"`                          // Debtor and Creditor Agents
	OriginalGroupInfo            common.GroupInformation           `json:"original_group_information,omitempty"`      // Original Group Information
	OriginalPaymentInformationId string                            `json:"original_payment_information_id,omitempty"` // Original Payment Information Identification
	OriginalTransactionInfo      common.TransactionDetailReference `json:"original_transaction_information,omitzero"` // Original Transaction
	TransactionStatus            common.TransactionStatusCode      `json:"transaction_status,omitempty"`              // Transaction Status
	PaymentConditionStatus       PaymentConditionStatus            `json:"payment_condition_status,omitzero"`         // Payment Condition Status
	ReturnReason                 common.ReturnReason               `json:"return_reason,omitzero"`                    // Return Reason
	OriginalTransactionReference TransactionReference              `json:"original_transaction_reference,omitzero"`   // Original Transaction Reference
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:pain.014.001.07"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &pain_014_001_07.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId",
	"CreatedDateTime",
	"InitiatingPartyName",
	"Agents.DebtorAgent.PaymentSysCode",
	"Agents.DebtorAgent.PaymentSysMemberId",
	"Agents.CreditorAgent.PaymentSysCode",
	"Agents.CreditorAgent.PaymentSysMemberId",
	"OriginalGroupInfo.OriginalMessageIdentification",
	"OriginalGroupInfo.OriginalMessageNameIdentification",
	"OriginalGroupInfo.OriginalCreationDateTime",
	"OriginalPaymentInformationId",
	"TransactionStatus",
}
var PathMap = map[string]any{
	"CdtrPmtActvtnReqStsRpt.GrpHdr.MsgId":                                                                        "MessageId",
	"CdtrPmtActvtnReqStsRpt.GrpHdr.CreDtTm":                                                                      "CreatedDateTime",
	"CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.Nm":                                                                  "InitiatingPartyName",
	"CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":                                   "Agents.DebtorAgent.PaymentSysCode",
	"CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId":                                         "Agents.DebtorAgent.PaymentSysMemberId",
	"CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId.Nm":                                                        "Agents.DebtorAgent.BankName",
	"CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":                                   "Agents.CreditorAgent.PaymentSysCode",
	"CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId":                                         "Agents.CreditorAgent.PaymentSysMemberId",
	"CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId.Nm":                                                        "Agents.CreditorAgent.BankName",
	"CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlMsgId":                                                        "OriginalGroupInfo.OriginalMessageIdentification",
	"CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlMsgNmId":                                                      "OriginalGroupInfo.OriginalMessageNameIdentification",
	"CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlCreDtTm":                                                      "OriginalGroupInfo.OriginalCreationDateTime",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.OrgnlPmtInfId":                                                     "OriginalPaymentInformationId",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlEndToEndId":                                       "OriginalTransactionInfo.EndToEndIdentification",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.TxSts":                                                 "TransactionStatus",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.StsRsnInf[0].Rsn.Cd":                                   "ReturnReason.Code",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.StsRsnInf[0].AddtlInf":                                 "ReturnReason.Info",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.StsRsnInf[0].Rsn.Prtry":                                "ReturnReason.Proprietary",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.PmtCondSts.AccptdAmt.Value":                            "PaymentConditionStatus.AcceptedAmount.Amount",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.PmtCondSts.AccptdAmt.Ccy":                              "PaymentConditionStatus.AcceptedAmount.Currency",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.PmtCondSts.GrntedPmt":                                  "PaymentConditionStatus.GuaranteedPayment",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.PmtCondSts.EarlyPmt":                                   "PaymentConditionStatus.EarlyPayment",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlTxRef.Amt.InstdAmt.Value":                         "OriginalTransactionReference.Amount.Amount",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlTxRef.Amt.InstdAmt.Ccy":                           "OriginalTransactionReference.Amount.Currency",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlTxRef.ReqdExctnDt.Dt":                             "OriginalTransactionReference.RequestedExecutionDate",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlTxRef.Cdtr.Nm":                                    "OriginalTransactionReference.Creditor.Agent.BankName",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlTxRef.Cdtr.PstlAdr.StrtNm":                        "OriginalTransactionReference.Creditor.Agent.PostalAddress.StreetName",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlTxRef.Cdtr.PstlAdr.BldgNm":                        "OriginalTransactionReference.Creditor.Agent.PostalAddress.BuildingName",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlTxRef.Cdtr.PstlAdr.BldgNb":                        "OriginalTransactionReference.Creditor.Agent.PostalAddress.BuildingNumber",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlTxRef.Cdtr.PstlAdr.Room":                          "OriginalTransactionReference.Creditor.Agent.PostalAddress.RoomNumber",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlTxRef.Cdtr.PstlAdr.Flr":                           "OriginalTransactionReference.Creditor.Agent.PostalAddress.Floor",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlTxRef.Cdtr.PstlAdr.PstBx":                         "OriginalTransactionReference.Creditor.Agent.PostalAddress.PostBox",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlTxRef.Cdtr.PstlAdr.PstCd":                         "OriginalTransactionReference.Creditor.Agent.PostalAddress.PostalCode",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlTxRef.Cdtr.PstlAdr.TwnNm":                         "OriginalTransactionReference.Creditor.Agent.PostalAddress.TownName",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlTxRef.Cdtr.PstlAdr.CtrySubDvsn":                   "OriginalTransactionReference.Creditor.Agent.PostalAddress.Subdivision",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlTxRef.Cdtr.PstlAdr.Ctry":                          "OriginalTransactionReference.Creditor.Agent.PostalAddress.Country",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlTxRef.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "OriginalTransactionReference.Creditor.PartyAgent.PaymentSysCode",
	"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts.OrgnlTxRef.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId":       "OriginalTransactionReference.Creditor.PartyAgent.PaymentSysMemberId",
}
