// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Validations for urn:iso:std:iso:20022:tech:xsd:camt.052.001.08
package AccountBalanceReport_camt_052_001_08

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
	fednow.AddError(&errs, baseName+".RltdAcct", v.RltdAcct.Validate())
	for indx := range v.Bal {
		fednow.AddError(&errs, baseName+".Bal", v.Bal[indx].Validate())
	}
	fednow.AddError(&errs, baseName+".TxsSummry", v.TxsSummry.Validate())
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

func (v AmountAndDirection35) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "AmountAndDirection35"
	fednow.AddError(&errs, baseName+".Amt", v.Amt.Validate())
	fednow.AddError(&errs, baseName+".CdtDbtInd", v.CdtDbtInd.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BalanceType10Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BalanceType10Choice1"
	if v.Prtry != nil {
		fednow.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BalanceType131) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BalanceType131"
	fednow.AddError(&errs, baseName+".CdOrPrtry", v.CdOrPrtry.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BankToCustomerAccountReportV08) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BankToCustomerAccountReportV08"
	fednow.AddError(&errs, baseName+".GrpHdr", v.GrpHdr.Validate())
	fednow.AddError(&errs, baseName+".Rpt", v.Rpt.Validate())
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
	fednow.AddError(&errs, baseName+".Tp", v.Tp.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CashAccountType2Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CashAccountType2Choice1"
	if v.Prtry != nil {
		fednow.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CashBalance81) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CashBalance81"
	fednow.AddError(&errs, baseName+".Tp", v.Tp.Validate())
	if v.CdtLine != nil {
		for indx := range v.CdtLine {
			fednow.AddError(&errs, baseName+".CdtLine", v.CdtLine[indx].Validate())
		}
	}
	fednow.AddError(&errs, baseName+".Amt", v.Amt.Validate())
	fednow.AddError(&errs, baseName+".CdtDbtInd", v.CdtDbtInd.Validate())
	fednow.AddError(&errs, baseName+".Dt", v.Dt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CreditLine31) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CreditLine31"
	fednow.AddError(&errs, baseName+".Incl", v.Incl.Validate())
	fednow.AddError(&errs, baseName+".Tp", v.Tp.Validate())
	fednow.AddError(&errs, baseName+".Amt", v.Amt.Validate())
	fednow.AddError(&errs, baseName+".Dt", v.Dt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CreditLineType1Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CreditLineType1Choice1"
	if v.Prtry != nil {
		fednow.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
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
	fednow.AddError(&errs, baseName+".OrgnlBizQry", v.OrgnlBizQry.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v NumberAndSumOfTransactions11) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "NumberAndSumOfTransactions11"
	if v.NbOfNtries != nil {
		fednow.AddError(&errs, baseName+".NbOfNtries", v.NbOfNtries.Validate())
	}
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
	fednow.AddError(&errs, baseName+".MsgNmId", v.MsgNmId.Validate())
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

func (v ProprietaryBankTransactionCodeStructure11) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ProprietaryBankTransactionCodeStructure11"
	fednow.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v TotalTransactions61) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "TotalTransactions61"
	for indx := range v.TtlNtriesPerBkTxCd {
		fednow.AddError(&errs, baseName+".TtlNtriesPerBkTxCd", v.TtlNtriesPerBkTxCd[indx].Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v TotalsPerBankTransactionCode51) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "TotalsPerBankTransactionCode51"
	fednow.AddError(&errs, baseName+".TtlNetNtry", v.TtlNetNtry.Validate())
	fednow.AddError(&errs, baseName+".CdtNtries", v.CdtNtries.Validate())
	fednow.AddError(&errs, baseName+".DbtNtries", v.DbtNtries.Validate())
	fednow.AddError(&errs, baseName+".BkTxCd", v.BkTxCd.Validate())
	fednow.AddError(&errs, baseName+".Dt", v.Dt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD SimpleType validations

func (v AccountReportingFedNow1) Validate() error {
	if err := fednow.ValidateEnumeration(string(v), "ABAR"); err != nil {
		return err
	}
	return nil
}

func (v AccountTypeFRS1) Validate() error {
	if err := fednow.ValidateEnumeration(string(v), "S", "M"); err != nil {
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

func (v BalanceReportFRS1) Validate() error {
	if err := fednow.ValidateEnumeration(string(v), "ABMS"); err != nil {
		return err
	}
	return nil
}

func (v BalanceTypeFRS1) Validate() error {
	if err := fednow.ValidateEnumeration(string(v), "AVAL", "DLOD", "OBFL", "OBNL", "OBPL", "ABAL", "AVLD"); err != nil {
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

func (v CreditLineTypeFRS1) Validate() error {
	if err := fednow.ValidateEnumeration(string(v), "COLL", "CLOD", "ULOD", "NCAP", "CCAP"); err != nil {
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

func (v NonNegativeDecimalNumber) Validate() error {
	if err := fednow.ValidateMinInclusive(float64(v), 0); err != nil {
		return err
	}
	if err := fednow.ValidateFractionDigits(fmt.Sprintf("%v", v), 17); err != nil {
		return err
	}
	if err := fednow.ValidateTotalDigits(fmt.Sprintf("%v", v), 18); err != nil {
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

func (v TransactionsSummaryTypeFRS1) Validate() error {
	if err := fednow.ValidateEnumeration(string(v), "FDWS", "NSSE", "AVOT", "UVOT", "MEMO", "FDWF", "FDAP", "FDNF"); err != nil {
		return err
	}
	return nil
}

func (v TrueFalseIndicator) Validate() error {
	return nil
}

func (v YesNoIndicator) Validate() error {
	return nil
}
