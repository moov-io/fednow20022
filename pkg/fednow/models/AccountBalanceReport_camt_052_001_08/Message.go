// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: AccountBalanceReport_camt.052.001.08
// Base Message: camt.052.001.08
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_qR07SZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

package AccountBalanceReport_camt_052_001_08

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/AccountBalanceReport_camt_052_001_08"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type MessageModel struct {
	MessageId             common.CAMTReportType             `json:"message_id,omitempty"`
	CreatedDateTime       time.Time                         `json:"created_date_time,omitempty"`
	Pagenation            common.MessagePagenation          `json:"pagenation,omitempty"`
	OriginalBusinessQuery common.OriginalBusinessQuery      `json:"original_business_query,omitempty"` // Optional field for original business query
	ReportId              common.ReportType                 `json:"report_id,omitempty"`
	ReportCreateDateTime  time.Time                         `json:"report_create_date_time,omitempty"`
	AccountOtherId        string                            `json:"account_other_id,omitempty"`
	AccountType           string                            `json:"account_type,omitempty"`             // Optional field for account type
	RelateAccountOtherId  string                            `json:"related_account_other_id,omitempty"` // Optional field for related account ID
	Balances              []common.Balance                  `json:"balances,omitempty"`
	TransactionsSummary   []common.TotalsPerBankTransaction `json:"transactions_summary,omitempty"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08"

var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &AccountBalanceReport_camt_052_001_08.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}

var RequireFileds = []string{
	"MessageId", "CreatedDateTime",
	"Pagenation.PageNumber", "Pagenation.LastPageIndicator",
	"OriginalBusinessQuery.MessageIdentification",
	"OriginalBusinessQuery.MessageNameIdentification",
	"OriginalBusinessQuery.CreationDateTime",
	"ReportId", "ReportCreateDateTime",
	"AccountOtherId", "AccountType", "RelateAccountOtherId",
	"TransactionsSummary",
}

var PathMap = map[string]any{
	"BkToCstmrAcctRpt.GrpHdr.MsgId":               "MessageId",
	"BkToCstmrAcctRpt.GrpHdr.CreDtTm":             "CreatedDateTime",
	"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb":       "Pagenation.PageNumber",
	"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd":  "Pagenation.LastPageIndicator",
	"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId":   "OriginalBusinessQuery.MessageIdentification",
	"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId": "OriginalBusinessQuery.MessageNameIdentification",
	"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.CreDtTm": "OriginalBusinessQuery.CreationDateTime",
	"BkToCstmrAcctRpt.Rpt.Id":                     "ReportId",
	"BkToCstmrAcctRpt.Rpt.CreDtTm":                "ReportCreateDateTime",
	"BkToCstmrAcctRpt.Rpt.Acct.Id.Othr.Id":        "AccountOtherId",
	"BkToCstmrAcctRpt.Rpt.Acct.Tp.Prtry":          "AccountType",
	"BkToCstmrAcctRpt.Rpt.RltdAcct.Id.Othr.Id":    "RelateAccountOtherId",
	"BkToCstmrAcctRpt.Rpt.Bal : Balances": map[string]any{
		"Tp.CdOrPrtry.Prtry": "BalanceTypeId",
		"Amt.Value":          "Amount.Amount",
		"Amt.Ccy":            "Amount.Currency",
		"CdtDbtInd":          "CreditDebitIndicator",
		"Dt.DtTm":            "DateTime",
		"CdtLine : CdtLines": map[string]string{
			"Incl":      "Included",
			"Tp.Prtry":  "Type",
			"Amt.Value": "Amount.Amount",
			"Amt.Ccy":   "Amount.Currency",
			"Dt.DtTm":   "DateTime",
		},
	},
	"BkToCstmrAcctRpt.Rpt.TxsSummry.TtlNtriesPerBkTxCd : TransactionsSummary": map[string]string{
		"TtlNetNtry.Amt":       "TotalNetEntries",
		"TtlNetNtry.CdtDbtInd": "CreditDebitIndicator",
		"CdtNtries.NbOfNtries": "CreditEntries.NumberOfEntries",
		"CdtNtries.Sum":        "CreditEntries.Sum",
		"DbtNtries.NbOfNtries": "DebitEntries.NumberOfEntries",
		"DbtNtries.Sum":        "DebitEntries.Sum",
		"BkTxCd.Prtry.Cd":      "BankTransactionCode",
		"Dt.DtTm":              "Date",
	},
}
