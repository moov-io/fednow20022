// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Validations for urn:iso:std:iso:20022:tech:xsd:camt.052.001.08
package AccountActivityDetailsReport_camt_052_001_08

import (
	"fmt"

	"github.com/moov-io/base"
	"github.com/moov-io/fednow20022/pkg/fednow"
)

// XSD Element validations

func (v Document) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Document"
	fednow.AddError(&errs, baseName+".BkToCstmrAcctRpt", v.BkToCstmrAcctRpt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD ComplexType validations

func (v AccountIdentification4Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "AccountIdentification4Choice1"
	if v.Othr != nil {
		fednow.AddError(&errs, baseName+".Othr", v.Othr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v AccountReport251) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "AccountReport251"
	fednow.AddError(&errs, baseName+".Id", v.Id.Validate())
	fednow.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	fednow.AddError(&errs, baseName+".Acct", v.Acct.Validate())
	if v.RltdAcct != nil {
		fednow.AddError(&errs, baseName+".RltdAcct", v.RltdAcct.Validate())
	}
	if v.TxsSummry != nil {
		fednow.AddError(&errs, baseName+".TxsSummry", v.TxsSummry.Validate())
	}
	if v.Ntry != nil {
		for indx := range v.Ntry {
			fednow.AddError(&errs, baseName+".Ntry", v.Ntry[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ActiveOrHistoricCurrencyAndAmount) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ActiveOrHistoricCurrencyAndAmount"

	fednow.AddError(&errs, baseName+".Ccy", v.Ccy.Validate())

	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BankToCustomerAccountReportV08) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BankToCustomerAccountReportV08"
	fednow.AddError(&errs, baseName+".GrpHdr", v.GrpHdr.Validate())
	for indx := range v.Rpt {
		fednow.AddError(&errs, baseName+".Rpt", v.Rpt[indx].Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BankTransactionCodeStructure41) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BankTransactionCodeStructure41"
	fednow.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BankTransactionCodeStructure42) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BankTransactionCodeStructure42"
	fednow.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BranchAndFinancialInstitutionIdentification61) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BranchAndFinancialInstitutionIdentification61"
	fednow.AddError(&errs, baseName+".FinInstnId", v.FinInstnId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BranchAndFinancialInstitutionIdentification62) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BranchAndFinancialInstitutionIdentification62"
	fednow.AddError(&errs, baseName+".FinInstnId", v.FinInstnId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CashAccount381) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CashAccount381"
	fednow.AddError(&errs, baseName+".Id", v.Id.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CashAccount391) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CashAccount391"
	fednow.AddError(&errs, baseName+".Id", v.Id.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ClearingSystemIdentification2Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ClearingSystemIdentification2Choice1"
	if v.Cd != nil {
		fednow.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ClearingSystemMemberIdentification21) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ClearingSystemMemberIdentification21"
	fednow.AddError(&errs, baseName+".ClrSysId", v.ClrSysId.Validate())
	fednow.AddError(&errs, baseName+".MmbId", v.MmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ClearingSystemMemberIdentification22) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ClearingSystemMemberIdentification22"
	fednow.AddError(&errs, baseName+".MmbId", v.MmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v DateAndDateTime2Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DateAndDateTime2Choice1"
	if v.DtTm != nil {
		fednow.AddError(&errs, baseName+".DtTm", v.DtTm.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v EntryDetails91) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "EntryDetails91"
	fednow.AddError(&errs, baseName+".TxDtls", v.TxDtls.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v EntryStatus1Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "EntryStatus1Choice1"
	if v.Cd != nil {
		fednow.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v EntryTransaction101) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "EntryTransaction101"
	fednow.AddError(&errs, baseName+".Refs", v.Refs.Validate())
	fednow.AddError(&errs, baseName+".RltdAgts", v.RltdAgts.Validate())
	if v.LclInstrm != nil {
		fednow.AddError(&errs, baseName+".LclInstrm", v.LclInstrm.Validate())
	}
	fednow.AddError(&errs, baseName+".RltdDts", v.RltdDts.Validate())
	if v.AddtlTxInf != nil {
		fednow.AddError(&errs, baseName+".AddtlTxInf", v.AddtlTxInf.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v FinancialInstitutionIdentification181) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FinancialInstitutionIdentification181"
	fednow.AddError(&errs, baseName+".ClrSysMmbId", v.ClrSysMmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v FinancialInstitutionIdentification182) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FinancialInstitutionIdentification182"
	fednow.AddError(&errs, baseName+".ClrSysMmbId", v.ClrSysMmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericAccountIdentification11) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericAccountIdentification11"
	fednow.AddError(&errs, baseName+".Id", v.Id.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GroupHeader811) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GroupHeader811"
	fednow.AddError(&errs, baseName+".MsgId", v.MsgId.Validate())
	fednow.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	fednow.AddError(&errs, baseName+".MsgPgntn", v.MsgPgntn.Validate())
	if v.OrgnlBizQry != nil {
		fednow.AddError(&errs, baseName+".OrgnlBizQry", v.OrgnlBizQry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v LocalInstrument2Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "LocalInstrument2Choice1"
	if v.Prtry != nil {
		fednow.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v MessageIdentification21) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "MessageIdentification21"
	fednow.AddError(&errs, baseName+".MsgNmId", v.MsgNmId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v NumberAndSumOfTransactions11) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "NumberAndSumOfTransactions11"
	fednow.AddError(&errs, baseName+".NbOfNtries", v.NbOfNtries.Validate())
	fednow.AddError(&errs, baseName+".Sum", v.Sum.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OriginalBusinessQuery11) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OriginalBusinessQuery11"
	fednow.AddError(&errs, baseName+".MsgId", v.MsgId.Validate())
	if v.MsgNmId != nil {
		fednow.AddError(&errs, baseName+".MsgNmId", v.MsgNmId.Validate())
	}
	fednow.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Pagination1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Pagination1"
	fednow.AddError(&errs, baseName+".PgNb", v.PgNb.Validate())
	fednow.AddError(&errs, baseName+".LastPgInd", v.LastPgInd.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ProprietaryAgent41) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ProprietaryAgent41"
	fednow.AddError(&errs, baseName+".Tp", v.Tp.Validate())
	fednow.AddError(&errs, baseName+".Agt", v.Agt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ProprietaryBankTransactionCodeStructure11) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ProprietaryBankTransactionCodeStructure11"
	fednow.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ProprietaryBankTransactionCodeStructure12) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ProprietaryBankTransactionCodeStructure12"
	fednow.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ProprietaryDate31) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ProprietaryDate31"
	fednow.AddError(&errs, baseName+".Tp", v.Tp.Validate())
	fednow.AddError(&errs, baseName+".Dt", v.Dt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ReportEntry101) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ReportEntry101"
	fednow.AddError(&errs, baseName+".Amt", v.Amt.Validate())
	fednow.AddError(&errs, baseName+".CdtDbtInd", v.CdtDbtInd.Validate())
	fednow.AddError(&errs, baseName+".Sts", v.Sts.Validate())
	fednow.AddError(&errs, baseName+".BkTxCd", v.BkTxCd.Validate())
	fednow.AddError(&errs, baseName+".AddtlInfInd", v.AddtlInfInd.Validate())
	fednow.AddError(&errs, baseName+".NtryDtls", v.NtryDtls.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v TotalTransactions61) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "TotalTransactions61"
	fednow.AddError(&errs, baseName+".TtlCdtNtries", v.TtlCdtNtries.Validate())
	fednow.AddError(&errs, baseName+".TtlDbtNtries", v.TtlDbtNtries.Validate())
	if v.TtlNtriesPerBkTxCd != nil {
		for indx := range v.TtlNtriesPerBkTxCd {
			fednow.AddError(&errs, baseName+".TtlNtriesPerBkTxCd", v.TtlNtriesPerBkTxCd[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v TotalsPerBankTransactionCode51) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "TotalsPerBankTransactionCode51"
	fednow.AddError(&errs, baseName+".NbOfNtries", v.NbOfNtries.Validate())
	fednow.AddError(&errs, baseName+".BkTxCd", v.BkTxCd.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v TransactionAgents51) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "TransactionAgents51"
	fednow.AddError(&errs, baseName+".InstgAgt", v.InstgAgt.Validate())
	fednow.AddError(&errs, baseName+".InstdAgt", v.InstdAgt.Validate())
	if v.Prtry != nil {
		for indx := range v.Prtry {
			fednow.AddError(&errs, baseName+".Prtry", v.Prtry[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v TransactionDates31) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "TransactionDates31"
	if v.AccptncDtTm != nil {
		fednow.AddError(&errs, baseName+".AccptncDtTm", v.AccptncDtTm.Validate())
	}
	for indx := range v.Prtry {
		fednow.AddError(&errs, baseName+".Prtry", v.Prtry[indx].Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v TransactionReferences61) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "TransactionReferences61"
	fednow.AddError(&errs, baseName+".MsgId", v.MsgId.Validate())
	if v.InstrId != nil {
		fednow.AddError(&errs, baseName+".InstrId", v.InstrId.Validate())
	}
	if v.EndToEndId != nil {
		fednow.AddError(&errs, baseName+".EndToEndId", v.EndToEndId.Validate())
	}
	if v.UETR != nil {
		fednow.AddError(&errs, baseName+".UETR", v.UETR.Validate())
	}
	if v.TxId != nil {
		fednow.AddError(&errs, baseName+".TxId", v.TxId.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD SimpleType validations

func (v AccountReportingFedNow1) Validate() error {
	if err := fednow.ValidateEnumeration(string(v), "AADR", "CADR"); err != nil {
		return err
	}
	return nil
}

func (v ActiveOrHistoricCurrencyCode) Validate() error {
	if err := fednow.ValidatePattern(string(v), `[A-Z]{3,3}`); err != nil {
		return err
	}
	return nil
}

func (v BankTransactionCodeFedNow1) Validate() error {
	if err := fednow.ValidateEnumeration(string(v), "SENT", "RCVD", "DBIT", "CRDT", "RJTS", "RJTR"); err != nil {
		return err
	}
	return nil
}

func (v BankTransactionCodeFedNow11) Validate() error {
	if err := fednow.ValidateEnumeration(string(v), "SENT", "RCVD", "RJTS", "RJTR"); err != nil {
		return err
	}
	return nil
}

func (v ConnectionPartyIdentifierFedNow1) Validate() error {
	if err := fednow.ValidatePattern(string(v), `[A-Z0-9]{9,9}`); err != nil {
		return err
	}
	return nil
}

func (v CreditDebitCode) Validate() error {
	if err := fednow.ValidateEnumeration(string(v), "CRDT", "DBIT"); err != nil {
		return err
	}
	return nil
}

func (v DecimalNumber) Validate() error {
	if err := fednow.ValidateFractionDigits(fmt.Sprintf("%v", v), 17); err != nil {
		return err
	}
	if err := fednow.ValidateTotalDigits(fmt.Sprintf("%v", v), 18); err != nil {
		return err
	}
	return nil
}

func (v ExternalClearingSystemIdentification1CodeFixed) Validate() error {
	if err := fednow.ValidateEnumeration(string(v), "USABA"); err != nil {
		return err
	}
	return nil
}

func (v ExternalEntryStatus1Code) Validate() error {
	if err := fednow.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fednow.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v Max15NumericText) Validate() error {
	if err := fednow.ValidatePattern(string(v), `[0-9]{1,15}`); err != nil {
		return err
	}
	return nil
}

func (v Max35Text) Validate() error {
	if err := fednow.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fednow.ValidateMaxLength(string(v), 35); err != nil {
		return err
	}
	return nil
}

func (v Max500Text) Validate() error {
	if err := fednow.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fednow.ValidateMaxLength(string(v), 500); err != nil {
		return err
	}
	return nil
}

func (v Max5NumericText) Validate() error {
	if err := fednow.ValidatePattern(string(v), `[0-9]{1,5}`); err != nil {
		return err
	}
	return nil
}

func (v MessageNameIdentificationFRS1) Validate() error {
	if err := fednow.ValidatePattern(string(v), `[a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}`); err != nil {
		return err
	}
	if err := fednow.ValidateLength(string(v), 15); err != nil {
		return err
	}
	return nil
}

func (v ReportDatesFedNow1) Validate() error {
	if err := fednow.ValidateEnumeration(string(v), "BPRD", "SDTM"); err != nil {
		return err
	}
	return nil
}

func (v ReportTimingFRS1) Validate() error {
	if err := fednow.ValidateEnumeration(string(v), "EDAY"); err != nil {
		return err
	}
	return nil
}

func (v RoutingNumberFRS1) Validate() error {
	if err := fednow.ValidatePattern(string(v), `[0-9]{9,9}`); err != nil {
		return err
	}
	return nil
}

func (v UUIDv4Identifier) Validate() error {
	if err := fednow.ValidatePattern(string(v), `[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}`); err != nil {
		return err
	}
	return nil
}

func (v YesNoIndicator) Validate() error {
	return nil
}
