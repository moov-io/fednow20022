// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: AccountActivityDetailsReport.camt.052.001.08
// Base Message: camt.052.001.08
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_qR07QZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

package AccountActivityDetailsReport_camt_052_001_08

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/camt_052_001_08"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type MessageModel struct {
	MessageId                          common.CAMTReportType                 `json:"message_id,omitempty"`
	CreatedDateTime                    time.Time                             `json:"created_date_time,omitzero"`
	Pagenation                         common.MessagePagenation              `json:"pagenation,omitempty"`
	OriginalBusinessQuery              common.OriginalBusinessQuery          `json:"original_business_query,omitzero"` // Optional field for original business query
	ReportId                           common.ReportType                     `json:"report_id,omitempty"`
	ReportCreateDateTime               time.Time                             `json:"report_create_date_time,omitzero"`
	AccountOtherId                     string                                `json:"account_other_id,omitempty"`
	RelateAccountOtherId               string                                `json:"related_account_other_id,omitempty"` // Optional field for related account ID
	TotalEntries                       string                                `json:"total_entries,omitempty"`
	TotalCreditEntries                 common.NumberAndSumOfTransactions     `json:"total_credit_entries,omitempty"`
	TotalDebitEntries                  common.NumberAndSumOfTransactions     `json:"total_debit_entries,omitempty"`
	TotalEntriesPerBankTransactionCode []common.TotalsPerBankTransactionCode `json:"total_entries_per_bank_transaction_code,omitempty"`
	EntryDetails                       []common.Entry                        `json:"entry_details,omitempty"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08"

var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &camt_052_001_08.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}

var RequireFileds = []string{
	"MessageId", "CreatedDateTime", "Pagenation.PageNumber", "Pagenation.LastPageIndicator",
}

var PathMap = map[string]any{
	// "BkToCstmrAcctRpt.GrpHdr.MsgId":                             "MessageId",
	"BkToCstmrAcctRpt.GrpHdr.CreDtTm": "CreatedDateTime",
	// "BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb":                     "Pagenation.PageNumber",
	// "BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd":                "Pagenation.LastPageIndicator",
	// "BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId":                 "OriginalBusinessQuery.MessageIdentification",
	// "BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId":               "OriginalBusinessQuery.MessageNameIdentification",
	// "BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.CreDtTm":               "OriginalBusinessQuery.CreationDateTime",
	// "BkToCstmrAcctRpt.Rpt[0].Id":                                "ReportId",
	// "BkToCstmrAcctRpt.Rpt[0].CreDtTm":                           "ReportCreateDateTime",
	// "BkToCstmrAcctRpt.Rpt[0].Acct.Id.Othr.Id":                   "AccountOtherId",
	// "BkToCstmrAcctRpt.Rpt[0].RltdAcct.Id.Othr.Id":               "RelateAccountOtherId",
	// "BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtries.NbOfNtries":    "TotalEntries",
	// "BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlCdtNtries.NbOfNtries": "TotalCreditEntries.NumberOfEntries",
	// "BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlCdtNtries.Sum":        "TotalCreditEntries.Sum",
	// "BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlDbtNtries.NbOfNtries": "TotalDebitEntries.NumberOfEntries",
	// "BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlDbtNtries.Sum":        "TotalDebitEntries.Sum",
	// "BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd : TotalEntriesPerBankTransactionCode": map[string]string{
	// 	"NbOfNtries":      "NumberOfEntries",
	// 	"BkTxCd.Prtry.Cd": "BankTransactionCode",
	// },
	// "BkToCstmrAcctRpt.Rpt[0].Ntry : EntryDetails": map[string]any{
	// 	"Amt.Value":                       "Amount.Amount",
	// 	"Amt.Ccy":                         "Amount.Currency",
	// 	"CdtDbtInd":                       "CreditDebitIndicator",
	// 	"Sts.Cd":                          "Status",
	// 	"BkTxCd.Prtry.Cd":                 "BankTransactionCode",
	// 	"AddtlInfInd.MsgNmId":             "MessageNameId",
	// 	"NtryDtls.TxDtls.Refs.MsgId":      "EntryDetails.MessageId",
	// 	"NtryDtls.TxDtls.Refs.InstrId":    "EntryDetails.InstructionId",
	// 	"NtryDtls.TxDtls.Refs.EndToEndId": "EntryDetails.EndToEndId",
	// 	"NtryDtls.TxDtls.Refs.UETR":       "EntryDetails.UniqueTransactionReference",
	// 	"NtryDtls.TxDtls.Refs.TxId":       "EntryDetails.TransactionIdentification",
	// 	"NtryDtls.TxDtls.RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "EntryDetails.InstructingAgent.PaymentSysCode",
	// 	"NtryDtls.TxDtls.RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.MmbId":       "EntryDetails.InstructingAgent.PaymentSysMemberId",
	// 	"NtryDtls.TxDtls.RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "EntryDetails.InstructedAgent.PaymentSysCode",
	// 	"NtryDtls.TxDtls.RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.MmbId":       "EntryDetails.InstructedAgent.PaymentSysMemberId",
	// 	"NtryDtls.TxDtls.RltdAgts.Prtry[0].Tp":                                 "EntryDetails.ProprietaryAgent.Type",
	// 	"NtryDtls.TxDtls.RltdAgts.Prtry[0].Agt.FinInstnId.ClrSysMmbId.MmbId":   "EntryDetails.ProprietaryAgent.PaymentSysMemberId",
	// 	"NtryDtls.TxDtls.LclInstrm.Prtry":                                      "EntryDetails.LocalInstrumentChoice",
	// 	"NtryDtls.TxDtls.AddtlTxInf":                                           "EntryDetails.AdditionalTransactionInformation",
	// 	"NtryDtls.TxDtls.RltdDts.AccptncDtTm":                                  "EntryDetails.AcceptanceDateTime",
	// 	"NtryDtls.TxDtls.RltdDts.Prtry : EntryDetails.TypeAndDateTimes": map[string]string{
	// 		"Tp":      "RelatedDatesProprietary",
	// 		"Dt.DtTm": "RelatedDateTime",
	// 	},
	// },
}
