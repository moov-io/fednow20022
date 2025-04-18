// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Validations for urn:fednow
package Message_AccountActivityTotalsReport_camt_052_001_08

import (
	"github.com/moov-io/base"
	"github.com/moov-io/fednow20022/pkg/fednow"
)

// XSD Element validations

func (v Message) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Message"
	fednow.AddError(&errs, baseName+".AppHdr", v.AppHdr.Validate())
	if v.AccountActivityTotalsReport != nil {
		fednow.AddError(&errs, baseName+".AccountActivityTotalsReport", v.AccountActivityTotalsReport.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD ComplexType validations
