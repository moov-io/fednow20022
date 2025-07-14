// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: PaymentStatusRequest_pacs.028.001.03
// Base Message: pacs.028.001.03
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7AUZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

package PaymentStatusRequest_pacs_028_001_03

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/PaymentStatusRequest_pacs_028_001_03"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type MessageModel struct {
	MessageId               string                            `json:"message_id,omitempty"`
	CreatedDateTime         time.Time                         `json:"created_date_time,omitempty"`
	OriginalGroupInfo       common.GroupInformation           `json:"original_group_info,omitempty"`
	OriginalTransactionInfo common.TransactionDetailReference `json:"payment_identification,omitempty"`
	Agents                  common.RelatedAgents              `json:"agents,omitempty"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &PaymentStatusRequest_pacs_028_001_03.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId",
	"CreatedDateTime",
	"OriginalGroupInfo.OriginalMessageIdentification",
	"OriginalGroupInfo.OriginalMessageNameIdentification",
	"OriginalGroupInfo.OriginalCreationDateTime",
	"Agents.InstructingAgent.PaymentSysCode",
	"Agents.InstructingAgent.PaymentSysMemberId",
	"Agents.InstructedAgent.PaymentSysCode",
	"Agents.InstructedAgent.PaymentSysMemberId",
}
var PathMap = map[string]any{
	"FIToFIPmtStsReq.GrpHdr.MsgId":                                      "MessageId",
	"FIToFIPmtStsReq.GrpHdr.CreDtTm":                                    "CreatedDateTime",
	"FIToFIPmtStsReq.TxInf.OrgnlGrpInf.OrgnlMsgId":                      "OriginalGroupInfo.OriginalMessageIdentification",
	"FIToFIPmtStsReq.TxInf.OrgnlGrpInf.OrgnlMsgNmId":                    "OriginalGroupInfo.OriginalMessageNameIdentification",
	"FIToFIPmtStsReq.TxInf.OrgnlGrpInf.OrgnlCreDtTm":                    "OriginalGroupInfo.OriginalCreationDateTime",
	"FIToFIPmtStsReq.TxInf.OrgnlInstrId":                                "OriginalTransactionInfo.InstructionIdentification",
	"FIToFIPmtStsReq.TxInf.OrgnlEndToEndId":                             "OriginalTransactionInfo.EndToEndIdentification",
	"FIToFIPmtStsReq.TxInf.OrgnlUETR":                                   "OriginalTransactionInfo.UETR",
	"FIToFIPmtStsReq.TxInf.OrgnlTxId":                                   "OriginalTransactionInfo.TransactionId",
	"FIToFIPmtStsReq.TxInf.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Agents.InstructingAgent.PaymentSysCode",
	"FIToFIPmtStsReq.TxInf.InstgAgt.FinInstnId.ClrSysMmbId.MmbId":       "Agents.InstructingAgent.PaymentSysMemberId",
	"FIToFIPmtStsReq.TxInf.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Agents.InstructedAgent.PaymentSysCode",
	"FIToFIPmtStsReq.TxInf.InstdAgt.FinInstnId.ClrSysMmbId.MmbId":       "Agents.InstructedAgent.PaymentSysMemberId",
}
