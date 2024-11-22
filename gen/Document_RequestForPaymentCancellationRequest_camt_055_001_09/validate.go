// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Validations for urn:fednow
package Document_RequestForPaymentCancellationRequest_camt_055_001_09

import (
	"github.com/moov-io/base"
	"github.com/moov-io/fednow20022/pkg/fednow"
)

// XSD Element validations

func (v Message) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Message"
	fednow.AddError(&errs, baseName+".AppHdr", v.AppHdr.Validate())
	if v.RequestForPaymentCancellationRequestResponse != nil {
		fednow.AddError(&errs, baseName+".RequestForPaymentCancellationRequestResponse", v.RequestForPaymentCancellationRequestResponse.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD ComplexType validations
