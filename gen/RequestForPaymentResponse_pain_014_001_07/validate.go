// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Validations for urn:iso:std:iso:20022:tech:xsd:pain.014.001.07
package RequestForPaymentResponse_pain_014_001_07

import (
	"fmt"

	"github.com/moov-io/base"
	"github.com/moov-io/fednow20022/pkg/fednow"
)

// XSD Element validations

func (v Document) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Document"
	fednow.AddError(&errs, baseName+".CdtrPmtActvtnReqStsRpt", v.CdtrPmtActvtnReqStsRpt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD ComplexType validations

func (v ActiveCurrencyAndAmountFedNow1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ActiveCurrencyAndAmountFedNow1"

	fednow.AddError(&errs, baseName+".Ccy", v.Ccy.Validate())

	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ActiveOrHistoricCurrencyAndAmountFedNow1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ActiveOrHistoricCurrencyAndAmountFedNow1"

	fednow.AddError(&errs, baseName+".Ccy", v.Ccy.Validate())

	if errs.Empty() {
		return nil
	}
	return errs
}

func (v AmountType4Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "AmountType4Choice1"
	if v.InstdAmt != nil {
		fednow.AddError(&errs, baseName+".InstdAmt", v.InstdAmt.Validate())
	}
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

func (v Contact41) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Contact41"
	if v.Nm != nil {
		fednow.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if v.PhneNb != nil {
		fednow.AddError(&errs, baseName+".PhneNb", v.PhneNb.Validate())
	}
	if v.MobNb != nil {
		fednow.AddError(&errs, baseName+".MobNb", v.MobNb.Validate())
	}
	if v.EmailAdr != nil {
		fednow.AddError(&errs, baseName+".EmailAdr", v.EmailAdr.Validate())
	}
	if v.Dept != nil {
		fednow.AddError(&errs, baseName+".Dept", v.Dept.Validate())
	}
	fednow.AddError(&errs, baseName+".PrefrdMtd", v.PrefrdMtd.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CreditorPaymentActivationRequestStatusReportV07) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CreditorPaymentActivationRequestStatusReportV07"
	fednow.AddError(&errs, baseName+".GrpHdr", v.GrpHdr.Validate())
	fednow.AddError(&errs, baseName+".OrgnlGrpInfAndSts", v.OrgnlGrpInfAndSts.Validate())
	fednow.AddError(&errs, baseName+".OrgnlPmtInfAndSts", v.OrgnlPmtInfAndSts.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v DateAndDateTime2Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DateAndDateTime2Choice"
	if v.Dt != nil {
		fednow.AddError(&errs, baseName+".Dt", v.Dt.Validate())
	}
	if v.DtTm != nil {
		fednow.AddError(&errs, baseName+".DtTm", v.DtTm.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v DateAndPlaceOfBirth1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DateAndPlaceOfBirth1"
	fednow.AddError(&errs, baseName+".BirthDt", v.BirthDt.Validate())
	if v.PrvcOfBirth != nil {
		fednow.AddError(&errs, baseName+".PrvcOfBirth", v.PrvcOfBirth.Validate())
	}
	fednow.AddError(&errs, baseName+".CityOfBirth", v.CityOfBirth.Validate())
	fednow.AddError(&errs, baseName+".CtryOfBirth", v.CtryOfBirth.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v FinancialInstitutionIdentification181) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FinancialInstitutionIdentification181"
	if v.BICFI != nil {
		fednow.AddError(&errs, baseName+".BICFI", v.BICFI.Validate())
	}
	fednow.AddError(&errs, baseName+".ClrSysMmbId", v.ClrSysMmbId.Validate())
	if v.LEI != nil {
		fednow.AddError(&errs, baseName+".LEI", v.LEI.Validate())
	}
	if v.Nm != nil {
		fednow.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if v.PstlAdr != nil {
		fednow.AddError(&errs, baseName+".PstlAdr", v.PstlAdr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericOrganisationIdentification1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericOrganisationIdentification1"
	fednow.AddError(&errs, baseName+".Id", v.Id.Validate())
	if v.SchmeNm != nil {
		fednow.AddError(&errs, baseName+".SchmeNm", v.SchmeNm.Validate())
	}
	if v.Issr != nil {
		fednow.AddError(&errs, baseName+".Issr", v.Issr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericPersonIdentification1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericPersonIdentification1"
	fednow.AddError(&errs, baseName+".Id", v.Id.Validate())
	if v.SchmeNm != nil {
		fednow.AddError(&errs, baseName+".SchmeNm", v.SchmeNm.Validate())
	}
	if v.Issr != nil {
		fednow.AddError(&errs, baseName+".Issr", v.Issr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GroupHeader871) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GroupHeader871"
	fednow.AddError(&errs, baseName+".MsgId", v.MsgId.Validate())
	fednow.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	fednow.AddError(&errs, baseName+".InitgPty", v.InitgPty.Validate())
	fednow.AddError(&errs, baseName+".DbtrAgt", v.DbtrAgt.Validate())
	fednow.AddError(&errs, baseName+".CdtrAgt", v.CdtrAgt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OrganisationIdentification291) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OrganisationIdentification291"
	if v.AnyBIC != nil {
		fednow.AddError(&errs, baseName+".AnyBIC", v.AnyBIC.Validate())
	}
	if v.LEI != nil {
		fednow.AddError(&errs, baseName+".LEI", v.LEI.Validate())
	}
	if v.Othr != nil {
		fednow.AddError(&errs, baseName+".Othr", v.Othr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OrganisationIdentificationSchemeName1Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OrganisationIdentificationSchemeName1Choice"
	if v.Cd != nil {
		fednow.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		fednow.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OriginalGroupInformation301) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OriginalGroupInformation301"
	fednow.AddError(&errs, baseName+".OrgnlMsgId", v.OrgnlMsgId.Validate())
	fednow.AddError(&errs, baseName+".OrgnlMsgNmId", v.OrgnlMsgNmId.Validate())
	fednow.AddError(&errs, baseName+".OrgnlCreDtTm", v.OrgnlCreDtTm.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OriginalPaymentInstruction311) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OriginalPaymentInstruction311"
	fednow.AddError(&errs, baseName+".OrgnlPmtInfId", v.OrgnlPmtInfId.Validate())
	fednow.AddError(&errs, baseName+".TxInfAndSts", v.TxInfAndSts.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OriginalTransactionReference291) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OriginalTransactionReference291"
	if v.Amt != nil {
		fednow.AddError(&errs, baseName+".Amt", v.Amt.Validate())
	}
	fednow.AddError(&errs, baseName+".ReqdExctnDt", v.ReqdExctnDt.Validate())
	fednow.AddError(&errs, baseName+".CdtrAgt", v.CdtrAgt.Validate())
	fednow.AddError(&errs, baseName+".Cdtr", v.Cdtr.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party38Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party38Choice1"
	if v.OrgId != nil {
		fednow.AddError(&errs, baseName+".OrgId", v.OrgId.Validate())
	}
	if v.PrvtId != nil {
		fednow.AddError(&errs, baseName+".PrvtId", v.PrvtId.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification1351) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification1351"
	fednow.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	if v.PstlAdr != nil {
		fednow.AddError(&errs, baseName+".PstlAdr", v.PstlAdr.Validate())
	}
	if v.Id != nil {
		fednow.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if v.CtryOfRes != nil {
		fednow.AddError(&errs, baseName+".CtryOfRes", v.CtryOfRes.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification1352) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification1352"
	fednow.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	if v.Id != nil {
		fednow.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if v.CtctDtls != nil {
		fednow.AddError(&errs, baseName+".CtctDtls", v.CtctDtls.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification1353) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification1353"
	fednow.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	if v.PstlAdr != nil {
		fednow.AddError(&errs, baseName+".PstlAdr", v.PstlAdr.Validate())
	}
	if v.Id != nil {
		fednow.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if v.CtryOfRes != nil {
		fednow.AddError(&errs, baseName+".CtryOfRes", v.CtryOfRes.Validate())
	}
	if v.CtctDtls != nil {
		fednow.AddError(&errs, baseName+".CtctDtls", v.CtctDtls.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PaymentConditionStatus11) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PaymentConditionStatus11"
	fednow.AddError(&errs, baseName+".AccptdAmt", v.AccptdAmt.Validate())
	fednow.AddError(&errs, baseName+".GrntedPmt", v.GrntedPmt.Validate())
	fednow.AddError(&errs, baseName+".EarlyPmt", v.EarlyPmt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PaymentTransaction1041) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PaymentTransaction1041"
	if v.OrgnlInstrId != nil {
		fednow.AddError(&errs, baseName+".OrgnlInstrId", v.OrgnlInstrId.Validate())
	}
	if v.OrgnlEndToEndId != nil {
		fednow.AddError(&errs, baseName+".OrgnlEndToEndId", v.OrgnlEndToEndId.Validate())
	}
	if v.OrgnlUETR != nil {
		fednow.AddError(&errs, baseName+".OrgnlUETR", v.OrgnlUETR.Validate())
	}
	fednow.AddError(&errs, baseName+".TxSts", v.TxSts.Validate())
	if v.StsRsnInf != nil {
		for indx := range v.StsRsnInf {
			fednow.AddError(&errs, baseName+".StsRsnInf", v.StsRsnInf[indx].Validate())
		}
	}
	if v.PmtCondSts != nil {
		fednow.AddError(&errs, baseName+".PmtCondSts", v.PmtCondSts.Validate())
	}
	if v.OrgnlTxRef != nil {
		fednow.AddError(&errs, baseName+".OrgnlTxRef", v.OrgnlTxRef.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PersonIdentification131) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PersonIdentification131"
	if v.DtAndPlcOfBirth != nil {
		fednow.AddError(&errs, baseName+".DtAndPlcOfBirth", v.DtAndPlcOfBirth.Validate())
	}
	if v.Othr != nil {
		fednow.AddError(&errs, baseName+".Othr", v.Othr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PersonIdentificationSchemeName1Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PersonIdentificationSchemeName1Choice"
	if v.Cd != nil {
		fednow.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		fednow.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PostalAddress241) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PostalAddress241"
	if v.Dept != nil {
		fednow.AddError(&errs, baseName+".Dept", v.Dept.Validate())
	}
	if v.SubDept != nil {
		fednow.AddError(&errs, baseName+".SubDept", v.SubDept.Validate())
	}
	if v.StrtNm != nil {
		fednow.AddError(&errs, baseName+".StrtNm", v.StrtNm.Validate())
	}
	if v.BldgNb != nil {
		fednow.AddError(&errs, baseName+".BldgNb", v.BldgNb.Validate())
	}
	if v.BldgNm != nil {
		fednow.AddError(&errs, baseName+".BldgNm", v.BldgNm.Validate())
	}
	if v.Flr != nil {
		fednow.AddError(&errs, baseName+".Flr", v.Flr.Validate())
	}
	if v.PstBx != nil {
		fednow.AddError(&errs, baseName+".PstBx", v.PstBx.Validate())
	}
	if v.Room != nil {
		fednow.AddError(&errs, baseName+".Room", v.Room.Validate())
	}
	fednow.AddError(&errs, baseName+".PstCd", v.PstCd.Validate())
	fednow.AddError(&errs, baseName+".TwnNm", v.TwnNm.Validate())
	if v.TwnLctnNm != nil {
		fednow.AddError(&errs, baseName+".TwnLctnNm", v.TwnLctnNm.Validate())
	}
	if v.DstrctNm != nil {
		fednow.AddError(&errs, baseName+".DstrctNm", v.DstrctNm.Validate())
	}
	fednow.AddError(&errs, baseName+".CtrySubDvsn", v.CtrySubDvsn.Validate())
	fednow.AddError(&errs, baseName+".Ctry", v.Ctry.Validate())
	if v.AdrLine != nil {
		fednow.AddError(&errs, baseName+".AdrLine", v.AdrLine.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v StatusReason6Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "StatusReason6Choice"
	if v.Cd != nil {
		fednow.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		fednow.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v StatusReasonInformation121) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "StatusReasonInformation121"
	if v.Orgtr != nil {
		fednow.AddError(&errs, baseName+".Orgtr", v.Orgtr.Validate())
	}
	fednow.AddError(&errs, baseName+".Rsn", v.Rsn.Validate())
	if v.AddtlInf != nil {
		fednow.AddError(&errs, baseName+".AddtlInf", v.AddtlInf.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD SimpleType validations

func (v ActiveCurrencyAndAmountFedNow1SimpleType) Validate() error {
	if err := fednow.ValidateMinInclusive(float64(v), 0); err != nil {
		return err
	}
	if err := fednow.ValidateFractionDigits(fmt.Sprintf("%v", v), 2); err != nil {
		return err
	}
	if err := fednow.ValidateTotalDigits(fmt.Sprintf("%v", v), 14); err != nil {
		return err
	}
	return nil
}

func (v ActiveCurrencyCodeFixed) Validate() error {
	if err := fednow.ValidateEnumeration(string(v), "USD"); err != nil {
		return err
	}
	return nil
}

func (v ActiveOrHistoricCurrencyAndAmountFedNow1SimpleType) Validate() error {
	if err := fednow.ValidateMinInclusive(float64(v), 0); err != nil {
		return err
	}
	if err := fednow.ValidateFractionDigits(fmt.Sprintf("%v", v), 2); err != nil {
		return err
	}
	if err := fednow.ValidateTotalDigits(fmt.Sprintf("%v", v), 14); err != nil {
		return err
	}
	return nil
}

func (v AnyBICDec2014Identifier) Validate() error {
	if err := fednow.ValidatePattern(string(v), `[A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}`); err != nil {
		return err
	}
	return nil
}

func (v BICFIDec2014Identifier) Validate() error {
	if err := fednow.ValidatePattern(string(v), `[A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}`); err != nil {
		return err
	}
	return nil
}

func (v CountryCode) Validate() error {
	if err := fednow.ValidatePattern(string(v), `[A-Z]{2,2}`); err != nil {
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

func (v ExternalOrganisationIdentification1Code) Validate() error {
	if err := fednow.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fednow.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v ExternalPaymentTransactionStatus1Code) Validate() error {
	if err := fednow.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fednow.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v ExternalPersonIdentification1Code) Validate() error {
	if err := fednow.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fednow.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v ExternalStatusReason1Code) Validate() error {
	if err := fednow.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fednow.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v LEIIdentifier) Validate() error {
	if err := fednow.ValidatePattern(string(v), `[A-Z0-9]{18,18}[0-9]{2,2}`); err != nil {
		return err
	}
	return nil
}

func (v Max105Text) Validate() error {
	if err := fednow.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fednow.ValidateMaxLength(string(v), 105); err != nil {
		return err
	}
	return nil
}

func (v Max140Text) Validate() error {
	if err := fednow.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fednow.ValidateMaxLength(string(v), 140); err != nil {
		return err
	}
	return nil
}

func (v Max16Text) Validate() error {
	if err := fednow.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fednow.ValidateMaxLength(string(v), 16); err != nil {
		return err
	}
	return nil
}

func (v Max2048Text) Validate() error {
	if err := fednow.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fednow.ValidateMaxLength(string(v), 2048); err != nil {
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

func (v Max70Text) Validate() error {
	if err := fednow.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fednow.ValidateMaxLength(string(v), 70); err != nil {
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

func (v PhoneNumber) Validate() error {
	if err := fednow.ValidatePattern(string(v), `\+[0-9]{1,3}-[0-9()+\-]{1,30}`); err != nil {
		return err
	}
	return nil
}

func (v PreferredContactMethod1Code1) Validate() error {
	if err := fednow.ValidateEnumeration(string(v), "MAIL", "PHON", "CELL"); err != nil {
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

func (v TrueFalseIndicator) Validate() error {
	return nil
}

func (v UUIDv4Identifier) Validate() error {
	if err := fednow.ValidatePattern(string(v), `[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}`); err != nil {
		return err
	}
	return nil
}
