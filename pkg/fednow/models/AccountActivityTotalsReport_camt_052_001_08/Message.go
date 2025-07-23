// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: AccountActivityTotalsReport.camt.052.001.08
// Base Message: camt.052.001.08
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_qR07RZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

package AccountActivityTotalsReport_camt_052_001_08

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/AccountActivityTotalsReport_camt_052_001_08"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type ReportModel struct {
	ReportId                           common.ReportType                     `json:"report_id,omitempty"`
	ReportCreateDateTime               time.Time                             `json:"report_create_date_time,omitzero"`
	AccountOtherId                     string                                `json:"account_other_id,omitempty"`
	RelateAccountOtherId               string                                `json:"related_account_other_id,omitempty"` // Optional field for related account ID
	TotalEntries                       string                                `json:"total_entries,omitempty"`
	TotalCreditEntries                 common.NumberAndSumOfTransactions     `json:"total_credit_entries,omitempty"`
	TotalDebitEntries                  common.NumberAndSumOfTransactions     `json:"total_debit_entries,omitempty"`
	TotalEntriesPerBankTransactionCode []common.TotalsPerBankTransactionCode `json:"total_entries_per_bank_transaction_code,omitempty"`
}

type MessageModel struct {
	MessageId             common.CAMTReportType        `json:"message_id,omitempty"`
	CreatedDateTime       time.Time                    `json:"created_date_time,omitempty"`
	Pagenation            common.MessagePagenation     `json:"pagenation,omitempty"`
	OriginalBusinessQuery common.OriginalBusinessQuery `json:"original_business_query,omitempty"` // Optional field for original business query
	Reports               []ReportModel                `json:"reports,omitempty"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08"

var DataFactory = func() any {
	return &MessageModel{}
}

var DocumentFactory = func() common.ISODocument {
	return &AccountActivityTotalsReport_camt_052_001_08.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}

var RequireFileds = []string{
	"MessageId", "CreatedDateTime", "Pagenation.PageNumber", "Pagenation.LastPageIndicator",
}

var PathMap = map[string]any{
	"BkToCstmrAcctRpt.GrpHdr.MsgId":               "MessageId",
	"BkToCstmrAcctRpt.GrpHdr.CreDtTm":             "CreatedDateTime",
	"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb":       "Pagenation.PageNumber",
	"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd":  "Pagenation.LastPageIndicator",
	"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId":   "OriginalBusinessQuery.MessageIdentification",
	"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId": "OriginalBusinessQuery.MessageNameIdentification",
	"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.CreDtTm": "OriginalBusinessQuery.CreationDateTime",
	"BkToCstmrAcctRpt.Rpt : Reports": map[string]any{
		"Id":                                "ReportId",
		"CreDtTm":                           "ReportCreateDateTime",
		"Acct.Id.Othr.Id":                   "AccountOtherId",
		"RltdAcct.Id.Othr.Id":               "RelateAccountOtherId",
		"TxsSummry.TtlNtries.NbOfNtries":    "TotalEntries",
		"TxsSummry.TtlCdtNtries.NbOfNtries": "TotalCreditEntries.NumberOfEntries",
		"TxsSummry.TtlCdtNtries.Sum":        "TotalCreditEntries.Sum",
		"TxsSummry.TtlDbtNtries.NbOfNtries": "TotalDebitEntries.NumberOfEntries",
		"TxsSummry.TtlDbtNtries.Sum":        "TotalDebitEntries.Sum",
		"TxsSummry.TtlNtriesPerBkTxCd : TotalEntriesPerBankTransactionCode": map[string]string{
			"NbOfNtries":      "NumberOfEntries",
			"BkTxCd.Prtry.Cd": "BankTransactionCode",
		},
	},
}
