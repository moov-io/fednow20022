// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: AccountReportingRequest_camt.060.001.05
// Base Message: camt.060.001.05
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_qR07UZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope
package AccountReportingRequest_camt_060_001_05

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/AccountReportingRequest_camt_060_001_05"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type MessageModel struct {
	MessageId              string                   `json:"message_id,omitempty"`
	CreatedDateTime        time.Time                `json:"created_date_time,omitzero"`
	ReportingRequestId     common.CAMTReportType    `json:"reporting_request_id,omitempty"`
	RequestedMessageNameId string                   `json:"requested_message_name_id,omitempty"`
	AccountOtherId         string                   `json:"account_other_id,omitempty"`
	AccountType            common.AccountType       `json:"account_type,omitempty"`
	AccountOwner           common.Agent             `json:"account_owner,omitempty"`
	ReportingPeriod        common.PeriodDateAndTime `json:"reporting_period,omitzero"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:camt.060.001.05"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &AccountReportingRequest_camt_060_001_05.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId",
	"CreatedDateTime",
	"ReportingRequestId",
	"RequestedMessageNameId",
	"AccountOtherId",
	"AccountOwner",
}
var PathMap = map[string]any{
	"AcctRptgReq.GrpHdr.MsgId":                                      "MessageId",
	"AcctRptgReq.GrpHdr.CreDtTm":                                    "CreatedDateTime",
	"AcctRptgReq.RptgReq.Id":                                        "ReportingRequestId",
	"AcctRptgReq.RptgReq.ReqdMsgNmId":                               "RequestedMessageNameId",
	"AcctRptgReq.RptgReq.Acct.Id.Othr.Id":                           "AccountOtherId",
	"AcctRptgReq.RptgReq.Acct.Tp.Prtry":                             "AccountType",
	"AcctRptgReq.RptgReq.AcctOwnr.Agt.FinInstnId.ClrSysMmbId.MmbId": "AccountOwner.PaymentSysMemberId",
	"AcctRptgReq.RptgReq.RptgPrd.FrToDt.FrDt":                       "ReportingPeriod.FromDate",
	"AcctRptgReq.RptgReq.RptgPrd.FrToDt.ToDt":                       "ReportingPeriod.ToDate",
	"AcctRptgReq.RptgReq.RptgPrd.Tp":                                "ReportingPeriod.Type",
}
