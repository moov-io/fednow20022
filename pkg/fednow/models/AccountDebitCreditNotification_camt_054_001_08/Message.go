// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: AccountDebitCreditNotification_camt.054.001.08
// Base Message: camt.054.001.08
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_qR07TZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

package AccountDebitCreditNotification_camt_054_001_08

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/camt_054_001_08"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type MessageModel struct {
	MessageId                      common.CAMTReportType             `json:"message_id,omitempty"`
	CreatedDateTime                time.Time                         `json:"created_date_time,omitzero"`
	Pagenation                     common.MessagePagenation          `json:"pagenation,omitempty"`
	NotificationId                 string                            `json:"notification_id,omitempty"`
	AccountOtherId                 string                            `json:"account_other_id,omitempty"`
	EntryAmount                    common.CurrencyAndAmount          `json:"entry_amount,omitempty"`
	EntryCreditDebitIndicator      common.CdtDbtInd                  `json:"entry_credit_debit_indicator,omitempty"`
	EntryStatus                    common.ReportStatus               `json:"entry_status,omitempty"`
	BankTransactionCode            common.TransactionStatusCode      `json:"bank_transaction_code,omitempty"`
	AdditionalInformationIndicator string                            `json:"additional_information_indicator,omitempty"`
	TransactionDetailRefer         common.TransactionDetailReference `json:"transaction_detail_reference,omitempty"`
	RelatedAgent                   common.RelatedAgents              `json:"related_agent,omitempty"`
	LocalInstrument                string                            `json:"local_instrument,omitempty"`
	RelatedDate                    common.RelatedDates               `json:"related_dates,omitempty"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:camt.054.001.08"

var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &camt_054_001_08.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId",
	"CreatedDateTime",
	"Pagenation.PageNumber",
	"Pagenation.LastPageIndicator",
	"NotificationId",
	"AccountOtherId",
	"EntryAmount.Amount",
	"EntryCreditDebitIndicator",
	"EntryStatus",
	"BankTransactionCode",
	"AdditionalInformationIndicator",
	"TransactionDetailRefer.TransactionId",
	"RelatedAgent.InstructingAgent.PaymentSysCode",
	"RelatedAgent.InstructingAgent.PaymentSysMemberId",
	"RelatedAgent.InstructedAgent.PaymentSysCode",
	"RelatedAgent.InstructedAgent.PaymentSysMemberId",
	"RelatedAgent.ProprietaryType",
	"RelatedAgent.ProprietaryAgent.PaymentSysCode",
	"RelatedAgent.ProprietaryAgent.PaymentSysMemberId",
	"LocalInstrument",
	"RelatedDate.AcceptanceDateTime",
	"RelatedDate.InterbankSettlementDate",
}
var PathMap = map[string]any{
	"BkToCstmrDbtCdtNtfctn.GrpHdr.MsgId":                                                                      "MessageId",
	"BkToCstmrDbtCdtNtfctn.GrpHdr.CreDtTm":                                                                    "CreatedDateTime",
	"BkToCstmrDbtCdtNtfctn.GrpHdr.MsgPgntn.PgNb":                                                              "Pagenation.PageNumber",
	"BkToCstmrDbtCdtNtfctn.GrpHdr.MsgPgntn.LastPgInd":                                                         "Pagenation.LastPageIndicator",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Id":                                                                         "NotificationId",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Acct.Id.Othr.Id":                                                            "AccountOtherId",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.Amt.Value":                                                             "EntryAmount.Amount",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.Amt.Ccy":                                                               "EntryAmount.Currency",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.CdtDbtInd":                                                             "EntryCreditDebitIndicator",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.Sts.Cd":                                                                "EntryStatus",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.BkTxCd.Prtry.Cd":                                                       "BankTransactionCode",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.AddtlInfInd.MsgNmId":                                                   "AdditionalInformationIndicator",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.NtryDtls.TxDtls.Refs.MsgId":                                            "TransactionDetailRefer.TransactionId",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.NtryDtls.TxDtls.Refs.InstrId":                                          "TransactionDetailRefer.InstructionIdentification",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.NtryDtls.TxDtls.Refs.EndToEndId":                                       "TransactionDetailRefer.EndToEndIdentification",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.NtryDtls.TxDtls.Refs.UETR":                                             "TransactionDetailRefer.UETR",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.NtryDtls.TxDtls.RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":  "RelatedAgent.InstructingAgent.PaymentSysCode",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.NtryDtls.TxDtls.RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.MmbId":        "RelatedAgent.InstructingAgent.PaymentSysMemberId",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.NtryDtls.TxDtls.RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":  "RelatedAgent.InstructedAgent.PaymentSysCode",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.NtryDtls.TxDtls.RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.MmbId":        "RelatedAgent.InstructedAgent.PaymentSysMemberId",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.NtryDtls.TxDtls.RltdAgts.Prtry.Tp":                                     "RelatedAgent.ProprietaryType",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.NtryDtls.TxDtls.RltdAgts.Prtry.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "RelatedAgent.ProprietaryAgent.PaymentSysCode",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.NtryDtls.TxDtls.RltdAgts.Prtry.Agt.FinInstnId.ClrSysMmbId.MmbId":       "RelatedAgent.ProprietaryAgent.PaymentSysMemberId",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.NtryDtls.TxDtls.LclInstrm.Prtry":                                       "LocalInstrument",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.NtryDtls.TxDtls.RltdDts.AccptncDtTm":                                   "RelatedDate.AcceptanceDateTime",
	"BkToCstmrDbtCdtNtfctn.Ntfctn.Ntry.NtryDtls.TxDtls.RltdDts.IntrBkSttlmDt":                                 "RelatedDate.InterbankSettlementDate",
}
