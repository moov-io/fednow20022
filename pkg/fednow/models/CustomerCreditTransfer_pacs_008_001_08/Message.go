// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: CustomerCreditTransfer_pacs.008.001.08
// Base Message: pacs.008.001.08
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_qR07XZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

package CustomerCreditTransfer_pacs_008_001_08

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/CustomerCreditTransfer_pacs_008_001_08"
	"github.com/moov-io/fednow20022/pkg/fednow"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type MessageModel struct {
	MessageId                    string                              `json:"message_id,omitempty"`
	CreatedDateTime              time.Time                           `json:"created_date_time,omitempty"`
	NumberOfTransactions         string                              `json:"number_of_transactions,omitempty"`
	SettlementInfo               common.SettlementInformation        `json:"settlement_info,omitempty"`
	PaymentIdentification        common.TransactionDetailReference   `json:"payment_identification,omitempty"`
	PaymentType                  common.PaymentTypeInfo              `json:"payment_type,omitempty"`
	InterbankSettlementAmount    common.CurrencyAndAmount            `json:"interbank_settlement_amount,omitempty"`
	InterbankSettlementDate      fednow.ISODate                      `json:"interbank_settlement_date,omitempty"`
	ChargeBearer                 common.ChargeBearerType             `json:"charge_bearer,omitempty"`
	Agents                       common.RelatedAgents                `json:"agents,omitempty"`
	Debtor                       common.TransactionParty             `json:"debtor,omitempty"`
	Creditor                     common.TransactionParty             `json:"creditor,omitempty"`
	PurposeCode                  string                              `json:"purpose_code,omitempty"`
	RemittanceInformation        common.RemittanceInformation        `json:"remittance_information,omitempty"`
	RelatedRemittanceInformation common.RelatedRemittanceInformation `json:"related_remittance_information,omitempty"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &CustomerCreditTransfer_pacs_008_001_08.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId",
	"CreatedDateTime",
	"NumberOfTransactions",
	"SettlementInfo.Method",
	"SettlementInfo.Service",
	"PaymentIdentification.EndToEndIdentification",
	"PaymentType.LocalInstrument",
	"PaymentType.CategoryPurpose",
	"InterbankSettlementAmount",
	"InterbankSettlementDate",
	"ChargeBearer",
	"Agents.InstructingAgent.PaymentSysCode",
	"Agents.InstructingAgent.PaymentSysMemberId",
	"Agents.InstructedAgent.PaymentSysCode",
	"Agents.InstructedAgent.PaymentSysMemberId",
	"Debtor.PartyName",
	"Debtor.PartyAccountId",
	"Debtor.PartyAgent.PaymentSysCode",
	"Debtor.PartyAgent.PaymentSysMemberId",
	"Creditor.PartyName",
	"Creditor.PartyAccountId",
	"Creditor.PartyAgent.PaymentSysCode",
	"Creditor.PartyAgent.PaymentSysMemberId",
}
var PathMap = map[string]any{
	"FIToFICstmrCdtTrf.GrpHdr.MsgId":                                             "MessageId",
	"FIToFICstmrCdtTrf.GrpHdr.CreDtTm":                                           "CreatedDateTime",
	"FIToFICstmrCdtTrf.GrpHdr.NbOfTxs":                                           "NumberOfTransactions",
	"FIToFICstmrCdtTrf.GrpHdr.SttlmInf.SttlmMtd":                                 "SettlementInfo.Method",
	"FIToFICstmrCdtTrf.GrpHdr.SttlmInf.ClrSys.Cd":                                "SettlementInfo.Service",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.PmtId.InstrId":                                "PaymentIdentification.InstructionIdentification",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.PmtId.EndToEndId":                             "PaymentIdentification.EndToEndIdentification",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.PmtId.UETR":                                   "PaymentIdentification.UETR",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.PmtId.TxId":                                   "PaymentIdentification.MessageIdentification",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.PmtTpInf.LclInstrm.Prtry":                     "PaymentType.LocalInstrument",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.PmtTpInf.CtgyPurp.Prtry":                      "PaymentType.CategoryPurpose",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.IntrBkSttlmAmt.Value":                         "InterbankSettlementAmount.Amount",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.IntrBkSttlmAmt.Ccy":                           "InterbankSettlementAmount.Currency",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.IntrBkSttlmDt":                                "InterbankSettlementDate",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.ChrgBr":                                       "ChargeBearer",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":  "Agents.InstructingAgent.PaymentSysCode",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.InstgAgt.FinInstnId.ClrSysMmbId.MmbId":        "Agents.InstructingAgent.PaymentSysMemberId",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":  "Agents.InstructedAgent.PaymentSysCode",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.InstdAgt.FinInstnId.ClrSysMmbId.MmbId":        "Agents.InstructedAgent.PaymentSysMemberId",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.Nm":                                      "Debtor.PartyName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.PstlAdr.StrtNm":                          "Debtor.PartyAddress.StreetName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.PstlAdr.BldgNb":                          "Debtor.PartyAddress.BuildingNumber",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.PstlAdr.BldgNm":                          "Debtor.PartyAddress.BuildingName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.PstlAdr.Flr":                             "Debtor.PartyAddress.Floor",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.PstlAdr.Room":                            "Debtor.PartyAddress.RoomNumber",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.PstlAdr.PstBx":                           "Debtor.PartyAddress.PostBox",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.PstlAdr.PstCd":                           "Debtor.PartyAddress.PostalCode",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.PstlAdr.TwnNm":                           "Debtor.PartyAddress.TownName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.PstlAdr.CtrySubDvsn":                     "Debtor.PartyAddress.Subdivision",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.PstlAdr.Ctry":                            "Debtor.PartyAddress.Country",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAcct.Id.Othr.Id":                          "Debtor.PartyAccountId",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAcct.Prxy.Tp.Cd":                          "Debtor.PartyAccountProxy.Type",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAcct.Prxy.Id":                             "Debtor.PartyAccountProxy.Value",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":   "Debtor.PartyAgent.PaymentSysCode",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId":         "Debtor.PartyAgent.PaymentSysMemberId",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAgt.FinInstnId.Nm":                        "Debtor.PartyAgent.BankName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtDbtr.Nm":                                 "Debtor.UltimatePartyName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtDbtr.PstlAdr.StrtNm":                     "Debtor.UltimatePartyAddress.StreetName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtDbtr.PstlAdr.BldgNb":                     "Debtor.UltimatePartyAddress.BuildingNumber",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtDbtr.PstlAdr.BldgNm":                     "Debtor.UltimatePartyAddress.BuildingName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtDbtr.PstlAdr.Flr":                        "Debtor.UltimatePartyAddress.Floor",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtDbtr.PstlAdr.Room":                       "Debtor.UltimatePartyAddress.RoomNumber",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtDbtr.PstlAdr.PstBx":                      "Debtor.UltimatePartyAddress.PostBox",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtDbtr.PstlAdr.PstCd":                      "Debtor.UltimatePartyAddress.PostalCode",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtDbtr.PstlAdr.TwnNm":                      "Debtor.UltimatePartyAddress.TownName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtDbtr.PstlAdr.CtrySubDvsn":                "Debtor.UltimatePartyAddress.Subdivision",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtDbtr.PstlAdr.Ctry":                       "Debtor.UltimatePartyAddress.Country",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.Nm":                                      "Creditor.PartyName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.PstlAdr.StrtNm":                          "Creditor.PartyAddress.StreetName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.PstlAdr.BldgNb":                          "Creditor.PartyAddress.BuildingNumber",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.PstlAdr.BldgNm":                          "Creditor.PartyAddress.BuildingName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.PstlAdr.Flr":                             "Creditor.PartyAddress.Floor",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.PstlAdr.Room":                            "Creditor.PartyAddress.RoomNumber",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.PstlAdr.PstBx":                           "Creditor.PartyAddress.PostBox",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.PstlAdr.PstCd":                           "Creditor.PartyAddress.PostalCode",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.PstlAdr.TwnNm":                           "Creditor.PartyAddress.TownName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.PstlAdr.CtrySubDvsn":                     "Creditor.PartyAddress.Subdivision",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.PstlAdr.Ctry":                            "Creditor.PartyAddress.Country",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.CdtrAcct.Id.Othr.Id":                          "Creditor.PartyAccountId",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.CdtrAcct.Prxy.Tp.Cd":                          "Creditor.PartyAccountProxy.Type",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.CdtrAcct.Prxy.Id":                             "Creditor.PartyAccountProxy.Value",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":   "Creditor.PartyAgent.PaymentSysCode",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId":         "Creditor.PartyAgent.PaymentSysMemberId",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.CdtrAgt.FinInstnId.Nm":                        "Creditor.PartyAgent.BankName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr.Nm":                                 "Creditor.UltimatePartyName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr.PstlAdr.StrtNm":                     "Creditor.UltimatePartyAddress.StreetName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr.PstlAdr.BldgNb":                     "Creditor.UltimatePartyAddress.BuildingNumber",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr.PstlAdr.BldgNm":                     "Creditor.UltimatePartyAddress.BuildingName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr.PstlAdr.Flr":                        "Creditor.UltimatePartyAddress.Floor",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr.PstlAdr.Room":                       "Creditor.UltimatePartyAddress.RoomNumber",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr.PstlAdr.PstBx":                      "Creditor.UltimatePartyAddress.PostBox",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr.PstlAdr.PstCd":                      "Creditor.UltimatePartyAddress.PostalCode",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr.PstlAdr.TwnNm":                      "Creditor.UltimatePartyAddress.TownName",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr.PstlAdr.CtrySubDvsn":                "Creditor.UltimatePartyAddress.Subdivision",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr.PstlAdr.Ctry":                       "Creditor.UltimatePartyAddress.Country",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.Purp.Cd":                                      "PurposeCode",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.RmtInf.Ustrd":                                 "RemittanceInformation.Unstructured",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.RmtInf.Strd[0].RfrdDocInf[0].Tp.CdOrPrtry.Cd": "RemittanceInformation.ReferredDocumentInformation",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.RmtInf.Strd[0].RfrdDocInf[0].Nb":              "RemittanceInformation.Number",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.RmtInf.Strd[0].RfrdDocInf[0].RltdDt":          "RemittanceInformation.RelatedDate",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.RltdRmtInf.RmtId":                             "RelatedRemittanceInformation.RemittanceIdentification",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.RltdRmtInf.RmtLctnDtls[0].Mtd":                "RelatedRemittanceInformation.Method",
	"FIToFICstmrCdtTrf.CdtTrfTxInf.RltdRmtInf.RmtLctnDtls[0].ElctrncAdr":         "RelatedRemittanceInformation.ElectronicAddress",
}
