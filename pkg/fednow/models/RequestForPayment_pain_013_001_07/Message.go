// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: RequestForPayment_pain.013.001.07
// Base Message: pain.013.001.07
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7AYZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

package RequestForPayment_pain_013_001_07

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/RequestForPayment_pain_013_001_07"
	"github.com/moov-io/fednow20022/pkg/fednow"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type PaymentCondition struct {
	AmountModificationAllowed  bool `json:"amount_modification_allowed,omitempty"`
	EarlyPaymentAllowed        bool `json:"early_payment_allowed,omitempty"`
	GuaranteedPaymentRequested bool `json:"guaranteed_payment_requested,omitempty"`
}
type MessageModel struct {
	MessageId                           string                              `json:"message_id,omitempty"`
	CreatedDateTime                     time.Time                           `json:"created_date_time,omitzero"`
	NumberOfTransactions                string                              `json:"number_of_transactions,omitempty"`
	InitiatingPartyName                 string                              `json:"initiating_party_name,omitempty"`
	PaymentInformationId                string                              `json:"payment_information_id,omitempty"`
	PaymentMethod                       common.PaymentMethodCode            `json:"payment_method,omitempty"`
	PaymentCondition                    PaymentCondition                    `json:"payment_condition,omitzero"`
	RequestedExecutionDate              fednow.ISODate                      `json:"requested_execution_date,omitempty"`
	ExpiryDate                          fednow.ISODate                      `json:"expiry_date,omitempty"`
	Debtor                              common.TransactionParty             `json:"debtor,omitempty"`
	CreditTransferId                    common.TransactionDetailReference   `json:"credit_transfer_id,omitempty"`
	CreditTransferPaymentType           common.PaymentTypeInfo              `json:"credit_transfer_payment_type,omitempty"`
	CreditTransferAmount                common.CurrencyAndAmount            `json:"credit_transfer_amount,omitempty"`
	CreditTransferChargeBearer          common.ChargeBearerType             `json:"charge_bearer,omitempty"`
	CreditTransferCreditor              common.TransactionParty             `json:"creditor,omitempty"`
	CreditTransferPurposeCode           string                              `json:"purpose_code,omitempty"`
	CreditTransferRemitInfo             common.RemittanceInformation        `json:"credit_transfer_remit_info,omitempty"`
	CreditTransferRelatedRemittanceInfo common.RelatedRemittanceInformation `json:"related_remittance_information,omitzero"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:pain.013.001.07"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &RequestForPayment_pain_013_001_07.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId",
	"CreatedDateTime",
	"NumberOfTransactions",
	"InitiatingPartyName",
	"PaymentInformationId",
	"PaymentMethod",
	"RequestedExecutionDate",
	"ExpiryDate",
	"Debtor.Agent.BankName",
	"Debtor.PartyAccountId",
	"Debtor.PartyAgent.PaymentSysCode",
	"Debtor.PartyAgent.PaymentSysMemberId",
	"CreditTransferId.EndToEndIdentification",
	"CreditTransferAmount",
	"CreditTransferChargeBearer",
	"CreditTransferCreditor.PartyAgent.PaymentSysCode",
	"CreditTransferCreditor.PartyAgent.PaymentSysMemberId",
	"CreditTransferCreditor.Agent.BankName",
	"CreditTransferCreditor.PartyAccountId",
}
var PathMap = map[string]any{
	"CdtrPmtActvtnReq.GrpHdr.MsgId":                                                 "MessageId",
	"CdtrPmtActvtnReq.GrpHdr.CreDtTm":                                               "CreatedDateTime",
	"CdtrPmtActvtnReq.GrpHdr.NbOfTxs":                                               "NumberOfTransactions",
	"CdtrPmtActvtnReq.GrpHdr.InitgPty.Nm":                                           "InitiatingPartyName",
	"CdtrPmtActvtnReq.PmtInf.PmtInfId":                                              "PaymentInformationId",
	"CdtrPmtActvtnReq.PmtInf.PmtMtd":                                                "PaymentMethod",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.PmtCond.AmtModAllwd":                          "PaymentCondition.AmountModificationAllowed",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.PmtCond.EarlyPmtAllwd":                        "PaymentCondition.EarlyPaymentAllowed",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.PmtCond.GtdPmtReqd":                           "PaymentCondition.GuaranteedPaymentRequested",
	"CdtrPmtActvtnReq.PmtInf.ReqdExctnDt.Dt":                                        "RequestedExecutionDate",
	"CdtrPmtActvtnReq.PmtInf.XpryDt.Dt":                                             "ExpiryDate",
	"CdtrPmtActvtnReq.PmtInf.Dbtr.Nm":                                               "Debtor.Agent.BankName",
	"CdtrPmtActvtnReq.PmtInf.Dbtr.PstlAdr.StrtNm":                                   "Debtor.Agent.PostalAddress.StreetName",
	"CdtrPmtActvtnReq.PmtInf.Dbtr.PstlAdr.BldgNb":                                   "Debtor.Agent.PostalAddress.BuildingNumber",
	"CdtrPmtActvtnReq.PmtInf.Dbtr.PstlAdr.BldgNm":                                   "Debtor.Agent.PostalAddress.BuildingName",
	"CdtrPmtActvtnReq.PmtInf.Dbtr.PstlAdr.Flr":                                      "Debtor.Agent.PostalAddress.Floor",
	"CdtrPmtActvtnReq.PmtInf.Dbtr.PstlAdr.Room":                                     "Debtor.Agent.PostalAddress.RoomNumber",
	"CdtrPmtActvtnReq.PmtInf.Dbtr.PstlAdr.PstBx":                                    "Debtor.Agent.PostalAddress.PostBox",
	"CdtrPmtActvtnReq.PmtInf.Dbtr.PstlAdr.TwnNm":                                    "Debtor.Agent.PostalAddress.TownName",
	"CdtrPmtActvtnReq.PmtInf.Dbtr.PstlAdr.PstCd":                                    "Debtor.Agent.PostalAddress.PostalCode",
	"CdtrPmtActvtnReq.PmtInf.Dbtr.PstlAdr.CtrySubDvsn":                              "Debtor.Agent.PostalAddress.Subdivision",
	"CdtrPmtActvtnReq.PmtInf.Dbtr.PstlAdr.Ctry":                                     "Debtor.Agent.PostalAddress.Country",
	"CdtrPmtActvtnReq.PmtInf.DbtrAcct.Id.Othr.Id":                                   "Debtor.PartyAccountId",
	"CdtrPmtActvtnReq.PmtInf.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":            "Debtor.PartyAgent.PaymentSysCode",
	"CdtrPmtActvtnReq.PmtInf.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId":                  "Debtor.PartyAgent.PaymentSysMemberId",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.PmtId.EndToEndId":                             "CreditTransferId.EndToEndIdentification",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.PmtTpInf.LclInstrm.Prtry":                     "CreditTransferPaymentType.LocalInstrument",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.PmtTpInf.CtgyPurp.Prtry":                      "CreditTransferPaymentType.CategoryPurpose",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.Amt.InstdAmt.Value":                           "CreditTransferAmount.Amount",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.Amt.InstdAmt.Ccy":                             "CreditTransferAmount.Currency",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.ChrgBr":                                       "CreditTransferChargeBearer",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":   "CreditTransferCreditor.PartyAgent.PaymentSysCode",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId":         "CreditTransferCreditor.PartyAgent.PaymentSysMemberId",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.CdtrAgt.FinInstnId.Nm":                        "CreditTransferCreditor.PartyAgent.BankName",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.Cdtr.Nm":                                      "CreditTransferCreditor.Agent.BankName",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.Cdtr.PstlAdr.StrtNm":                          "CreditTransferCreditor.Agent.PostalAddress.StreetName",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.Cdtr.PstlAdr.BldgNb":                          "CreditTransferCreditor.Agent.PostalAddress.BuildingNumber",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.Cdtr.PstlAdr.BldgNm":                          "CreditTransferCreditor.Agent.PostalAddress.BuildingName",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.Cdtr.PstlAdr.Flr":                             "CreditTransferCreditor.Agent.PostalAddress.Floor",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.Cdtr.PstlAdr.Room":                            "CreditTransferCreditor.Agent.PostalAddress.RoomNumber",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.Cdtr.PstlAdr.PstBx":                           "CreditTransferCreditor.Agent.PostalAddress.PostBox",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.Cdtr.PstlAdr.TwnNm":                           "CreditTransferCreditor.Agent.PostalAddress.TownName",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.Cdtr.PstlAdr.PstCd":                           "CreditTransferCreditor.Agent.PostalAddress.PostalCode",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.Cdtr.PstlAdr.CtrySubDvsn":                     "CreditTransferCreditor.Agent.PostalAddress.Subdivision",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.Cdtr.PstlAdr.Ctry":                            "CreditTransferCreditor.Agent.PostalAddress.Country",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.CdtrAcct.Id.Othr.Id":                          "CreditTransferCreditor.PartyAccountId",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.Purp.Cd":                                      "CreditTransferPurposeCode",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.RmtInf.Strd[0].RfrdDocInf[0].Tp.CdOrPrtry.Cd": "CreditTransferRemitInfo.ReferredDocumentInformation",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.RmtInf.Strd[0].RfrdDocInf[0].Nb":              "CreditTransferRemitInfo.Number",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.RmtInf.Strd[0].RfrdDocInf[0].RltdDt":          "CreditTransferRemitInfo.RelatedDate",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.RltdRmtInf.RmtId":                             "CreditTransferRelatedRemittanceInfo.RemittanceIdentification",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.RltdRmtInf.RmtLctnDtls[0].Mtd":                "CreditTransferRelatedRemittanceInfo.Method",
	"CdtrPmtActvtnReq.PmtInf.CdtTrfTx.RltdRmtInf.RmtLctnDtls[0].ElctrncAdr":         "CreditTransferRelatedRemittanceInfo.ElectronicAddress",
}
