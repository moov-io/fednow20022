// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: PaymentReturn_pacs.004.001.10
// Base Message: pacs.004.001.10
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7ATZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

package PaymentReturn_pacs_004_001_10

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/PaymentReturn_pacs_004_001_10"
	"github.com/moov-io/fednow20022/pkg/fednow"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type MessageModel struct {
	MessageId                         string                            `json:"message_id,omitempty"`
	CreatedDateTime                   time.Time                         `json:"created_date_time,omitzero"`
	NumberOfTransactions              string                            `json:"number_of_transactions,omitempty"`
	SettlementInfo                    common.SettlementInformation      `json:"settlement_info,omitempty"`
	OriginalGroupInfo                 common.GroupInformation           `json:"original_group_info,omitempty"`
	OriginalTransactionInfo           common.TransactionDetailReference `json:"payment_identification,omitempty"`
	OriginalInterbankSettlementAmount common.CurrencyAndAmount          `json:"original_interbank_settlement_amount,omitempty"`
	OriginalInterbankSettlementDate   fednow.ISODate                    `json:"original_interbank_settlement_date,omitempty"`
	ReturnedInterbankSettlementAmount common.CurrencyAndAmount          `json:"returned_interbank_settlement_amount,omitempty"`
	InterbankSettlementDate           fednow.ISODate                    `json:"interbank_settlement_date_time,omitempty"`
	ChargeBearer                      common.ChargeBearerType           `json:"charge_bearer,omitempty"`
	Agents                            common.RelatedAgents              `json:"agents,omitempty"`
	Debtor                            common.TransactionParty           `json:"debtor,omitempty"`
	Creditor                          common.TransactionParty           `json:"creditor,omitempty"`
	ReturnReasonInformation           common.ReturnReason               `json:"return_reason_information,omitempty"`
	PaymentType                       common.PaymentTypeInfo            `json:"payment_type,omitempty"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &PaymentReturn_pacs_004_001_10.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId",
	"CreatedDateTime",
	"NumberOfTransactions",
	"SettlementInfo.Method",
	"SettlementInfo.Service",
	"OriginalGroupInfo.OriginalMessageIdentification",
	"OriginalGroupInfo.OriginalMessageNameIdentification",
	"OriginalGroupInfo.OriginalCreationDateTime",
	"OriginalInterbankSettlementAmount",
	"OriginalInterbankSettlementDate",
	"ReturnedInterbankSettlementAmount",
	"InterbankSettlementDate",
	"ChargeBearer",
	"Agents.InstructingAgent.PaymentSysCode",
	"Agents.InstructingAgent.PaymentSysMemberId",
	"Agents.InstructedAgent.PaymentSysCode",
	"Agents.InstructedAgent.PaymentSysMemberId",
	"Debtor",
	"Creditor.PartyAgent.PaymentSysCode",
	"Creditor.PartyAgent.PaymentSysMemberId",
	"Creditor",
	"Creditor.PartyAccountId",
	"ReturnReasonInformation",
	"PaymentType",
}
var PathMap = map[string]any{
	"PmtRtr.GrpHdr.MsgId":                                               "MessageId",
	"PmtRtr.GrpHdr.CreDtTm":                                             "CreatedDateTime",
	"PmtRtr.GrpHdr.NbOfTxs":                                             "NumberOfTransactions",
	"PmtRtr.GrpHdr.SttlmInf.SttlmMtd":                                   "SettlementInfo.Method",
	"PmtRtr.GrpHdr.SttlmInf.ClrSys.Cd":                                  "SettlementInfo.Service",
	"PmtRtr.TxInf.OrgnlGrpInf.OrgnlMsgId":                               "OriginalGroupInfo.OriginalMessageIdentification",
	"PmtRtr.TxInf.OrgnlGrpInf.OrgnlMsgNmId":                             "OriginalGroupInfo.OriginalMessageNameIdentification",
	"PmtRtr.TxInf.OrgnlGrpInf.OrgnlCreDtTm":                             "OriginalGroupInfo.OriginalCreationDateTime",
	"PmtRtr.TxInf.OrgnlInstrId":                                         "OriginalTransactionInfo.InstructionIdentification",
	"PmtRtr.TxInf.OrgnlEndToEndId":                                      "OriginalTransactionInfo.EndToEndIdentification",
	"PmtRtr.TxInf.OrgnlUETR":                                            "OriginalTransactionInfo.UETR",
	"PmtRtr.TxInf.OrgnlTxId":                                            "OriginalTransactionInfo.TransactionId",
	"PmtRtr.TxInf.OrgnlIntrBkSttlmAmt.Value":                            "OriginalInterbankSettlementAmount.Amount",
	"PmtRtr.TxInf.OrgnlIntrBkSttlmAmt.Ccy":                              "OriginalInterbankSettlementAmount.Currency",
	"PmtRtr.TxInf.OrgnlIntrBkSttlmDt":                                   "OriginalInterbankSettlementDate",
	"PmtRtr.TxInf.RtrdIntrBkSttlmAmt.Value":                             "ReturnedInterbankSettlementAmount.Amount",
	"PmtRtr.TxInf.RtrdIntrBkSttlmAmt.Ccy":                               "ReturnedInterbankSettlementAmount.Currency",
	"PmtRtr.TxInf.IntrBkSttlmDt":                                        "InterbankSettlementDate",
	"PmtRtr.TxInf.ChrgBr":                                               "ChargeBearer",
	"PmtRtr.TxInf.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":          "Agents.InstructingAgent.PaymentSysCode",
	"PmtRtr.TxInf.InstgAgt.FinInstnId.ClrSysMmbId.MmbId":                "Agents.InstructingAgent.PaymentSysMemberId",
	"PmtRtr.TxInf.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":          "Agents.InstructedAgent.PaymentSysCode",
	"PmtRtr.TxInf.InstdAgt.FinInstnId.ClrSysMmbId.MmbId":                "Agents.InstructedAgent.PaymentSysMemberId",
	"PmtRtr.TxInf.RtrChain.Dbtr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Debtor.Agent.PaymentSysCode",
	"PmtRtr.TxInf.RtrChain.Dbtr.Agt.FinInstnId.ClrSysMmbId.MmbId":       "Debtor.Agent.PaymentSysMemberId",
	"PmtRtr.TxInf.RtrChain.Dbtr.Agt.FinInstnId.Nm":                      "Debtor.Agent.BankName",
	"PmtRtr.TxInf.RtrChain.Dbtr.Pty.Nm":                                 "Debtor.PartyName",
	"PmtRtr.TxInf.RtrChain.Dbtr.Pty.PstlAdr.StrtNm":                     "Debtor.PartyAddress.StreetName",
	"PmtRtr.TxInf.RtrChain.Dbtr.Pty.PstlAdr.BldgNb":                     "Debtor.PartyAddress.BuildingNumber",
	"PmtRtr.TxInf.RtrChain.Dbtr.Pty.PstlAdr.BldgNm":                     "Debtor.PartyAddress.BuildingName",
	"PmtRtr.TxInf.RtrChain.Dbtr.Pty.PstlAdr.Flr":                        "Debtor.PartyAddress.Floor",
	"PmtRtr.TxInf.RtrChain.Dbtr.Pty.PstlAdr.PstBx":                      "Debtor.PartyAddress.PostBox",
	"PmtRtr.TxInf.RtrChain.Dbtr.Pty.PstlAdr.Room":                       "Debtor.PartyAddress.RoomNumber",
	"PmtRtr.TxInf.RtrChain.Dbtr.Pty.PstlAdr.PstCd":                      "Debtor.PartyAddress.PostalCode",
	"PmtRtr.TxInf.RtrChain.Dbtr.Pty.PstlAdr.TwnNm":                      "Debtor.PartyAddress.TownName",
	"PmtRtr.TxInf.RtrChain.Dbtr.Pty.PstlAdr.CtrySubDvsn":                "Debtor.PartyAddress.Subdivision",
	"PmtRtr.TxInf.RtrChain.Dbtr.Pty.PstlAdr.Ctry":                       "Debtor.PartyAddress.Country",
	"PmtRtr.TxInf.RtrChain.DbtrAcct.Id.Othr.Id":                         "Debtor.PartyAccountId",
	"PmtRtr.TxInf.RtrChain.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":  "Debtor.PartyAgent.PaymentSysCode",
	"PmtRtr.TxInf.RtrChain.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId":        "Debtor.PartyAgent.PaymentSysMemberId",
	"PmtRtr.TxInf.RtrChain.DbtrAgt.FinInstnId.Nm":                       "Debtor.PartyAgent.BankName",
	"PmtRtr.TxInf.RtrChain.Cdtr.Pty.Nm":                                 "Creditor.PartyName",
	"PmtRtr.TxInf.RtrChain.Cdtr.Pty.PstlAdr.StrtNm":                     "Creditor.PartyAddress.StreetName",
	"PmtRtr.TxInf.RtrChain.Cdtr.Pty.PstlAdr.BldgNb":                     "Creditor.PartyAddress.BuildingNumber",
	"PmtRtr.TxInf.RtrChain.Cdtr.Pty.PstlAdr.BldgNm":                     "Creditor.PartyAddress.BuildingName",
	"PmtRtr.TxInf.RtrChain.Cdtr.Pty.PstlAdr.Flr":                        "Creditor.PartyAddress.Floor",
	"PmtRtr.TxInf.RtrChain.Cdtr.Pty.PstlAdr.PstBx":                      "Creditor.PartyAddress.PostBox",
	"PmtRtr.TxInf.RtrChain.Cdtr.Pty.PstlAdr.Room":                       "Creditor.PartyAddress.RoomNumber",
	"PmtRtr.TxInf.RtrChain.Cdtr.Pty.PstlAdr.PstCd":                      "Creditor.PartyAddress.PostalCode",
	"PmtRtr.TxInf.RtrChain.Cdtr.Pty.PstlAdr.TwnNm":                      "Creditor.PartyAddress.TownName",
	"PmtRtr.TxInf.RtrChain.Cdtr.Pty.PstlAdr.CtrySubDvsn":                "Creditor.PartyAddress.Subdivision",
	"PmtRtr.TxInf.RtrChain.Cdtr.Pty.PstlAdr.Ctry":                       "Creditor.PartyAddress.Country",
	"PmtRtr.TxInf.RtrChain.CdtrAcct.Id.Othr.Id":                         "Creditor.PartyAccountId",
	"PmtRtr.TxInf.RtrChain.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":  "Creditor.PartyAgent.PaymentSysCode",
	"PmtRtr.TxInf.RtrChain.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId":        "Creditor.PartyAgent.PaymentSysMemberId",
	"PmtRtr.TxInf.RtrChain.CdtrAgt.FinInstnId.Nm":                       "Creditor.PartyAgent.BankName",
	"PmtRtr.TxInf.RtrRsnInf.Rsn.Cd":                                     "ReturnReasonInformation.Code",
	"PmtRtr.TxInf.RtrRsnInf.AddtlInf":                                   "ReturnReasonInformation.Info",
	"PmtRtr.TxInf.OrgnlTxRef.PmtTpInf.LclInstrm.Prtry":                  "PaymentType.LocalInstrument",
}
