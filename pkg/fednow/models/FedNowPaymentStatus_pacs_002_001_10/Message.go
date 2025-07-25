// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: FedNowPaymentStatus_pacs.002.001.10
// Base Message: pacs.002.001.10
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7ALZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

package FedNowPaymentStatus_pacs_002_001_10

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/FedNowPaymentStatus_pacs_002_001_10"
	"github.com/moov-io/fednow20022/pkg/fednow"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type MessageModel struct {
	MessageId                        string                            `json:"message_id,omitempty"`
	CreatedDateTime                  time.Time                         `json:"created_date_time,omitzero"`
	GroupInfo                        common.GroupInformation           `json:"group_info,omitempty"`
	OriginalTransactionInfo          common.TransactionDetailReference `json:"payment_identification,omitempty"`
	TransactionStatus                common.TransactionStatusCode      `json:"transaction_status,omitempty"`
	StatusReasonInformation          common.Reason                     `json:"status_reason_information,omitzero"` // This field is not in the original document, but can be added if needed
	AcceptanceDateTime               time.Time                         `json:"acceptance_dateTime,omitzero"`
	EffectiveInterbankSettlementDate fednow.ISODate                    `json:"effective_interbank_settlement_date,omitempty"`
	Agents                           common.RelatedAgents              `json:"agents,omitempty"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &FedNowPaymentStatus_pacs_002_001_10.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId",
	"CreatedDateTime",
	"GroupInfo.OriginalMessageIdentification",
	"GroupInfo.OriginalMessageNameIdentification",
	"GroupInfo.OriginalCreationDateTime",
	"TransactionStatus",
	"Agents.InstructingAgent.PaymentSysCode",
	"Agents.InstructingAgent.PaymentSysMemberId",
	"Agents.InstructedAgent.PaymentSysCode",
	"Agents.InstructedAgent.PaymentSysMemberId",
}
var PathMap = map[string]any{
	"FIToFIPmtStsRpt.GrpHdr.MsgId":                                            "MessageId",
	"FIToFIPmtStsRpt.GrpHdr.CreDtTm":                                          "CreatedDateTime",
	"FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OrgnlCreDtTm":                    "GroupInfo.OriginalCreationDateTime",
	"FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OrgnlMsgId":                      "GroupInfo.OriginalMessageIdentification",
	"FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OrgnlMsgNmId":                    "GroupInfo.OriginalMessageNameIdentification",
	"FIToFIPmtStsRpt.TxInfAndSts.OrgnlInstrId":                                "OriginalTransactionInfo.InstructionIdentification",
	"FIToFIPmtStsRpt.TxInfAndSts.OrgnlEndToEndId":                             "OriginalTransactionInfo.EndToEndIdentification",
	"FIToFIPmtStsRpt.TxInfAndSts.OrgnlUETR":                                   "OriginalTransactionInfo.UETR",
	"FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxId":                                   "OriginalTransactionInfo.TransactionId",
	"FIToFIPmtStsRpt.TxInfAndSts.TxSts":                                       "TransactionStatus",
	"FIToFIPmtStsRpt.TxInfAndSts.StsRsnInf[0].Rsn.Cd":                         "StatusReasonInformation.Code",
	"FIToFIPmtStsRpt.TxInfAndSts.StsRsnInf[0].Rsn.Prtry":                      "StatusReasonInformation.Proprietary",
	"FIToFIPmtStsRpt.TxInfAndSts.AccptncDtTm":                                 "AcceptanceDateTime",
	"FIToFIPmtStsRpt.TxInfAndSts.FctvIntrBkSttlmDt.Dt":                        "EffectiveInterbankSettlementDate",
	"FIToFIPmtStsRpt.TxInfAndSts.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Agents.InstructingAgent.PaymentSysCode",
	"FIToFIPmtStsRpt.TxInfAndSts.InstgAgt.FinInstnId.ClrSysMmbId.MmbId":       "Agents.InstructingAgent.PaymentSysMemberId",
	"FIToFIPmtStsRpt.TxInfAndSts.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Agents.InstructedAgent.PaymentSysCode",
	"FIToFIPmtStsRpt.TxInfAndSts.InstdAgt.FinInstnId.ClrSysMmbId.MmbId":       "Agents.InstructedAgent.PaymentSysMemberId",
}
