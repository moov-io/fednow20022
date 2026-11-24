//go:build ignore
// +build ignore

// fixup_xmlnames.go sets XML element names in the generated models.
//
// moovio_xsd2go writes `xml:",any"` for a complex type that has only one
// element, and it writes no name on the XMLName field of that type. The
// committed models use the element names instead, so that decoding and
// encoding use the correct elements. This program applies those names after
// the models are generated and formatted. Each replacement must match exactly
// one time, or the program fails.
//
// It also removes the unexported innerXml fields that moovio_xsd2go writes.
// encoding/xml ignores unexported fields, so these fields are never set or
// written, and go vet reports each of them ("struct field innerXml has xml
// tag but is not exported").
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type replacement struct {
	old, new string
}

var fixups = []struct {
	path         string
	replacements []replacement
}{
	{
		path: "gen/admi_002_001_01/models.go",
		replacements: []replacement{
			{
				old: "\tRef Max35Text `xml:\",any\"`",
				new: "\tRef Max35Text `xml:\"Ref\"`",
			},
		},
	},
	{
		path: "gen/admi_004_001_02/models.go",
		replacements: []replacement{
			{
				old: "\tSysEvtNtfctn SystemEventNotificationV02 `xml:\",any\"`",
				new: "\tSysEvtNtfctn SystemEventNotificationV02 `xml:\"SysEvtNtfctn\"`",
			},
			{
				old: `type Event2 struct {
	XMLName xml.Name`,
				new: "type Event2 struct {\n\tXMLName xml.Name `xml:\"EvtInf\"`",
			},
			{
				old: `type SystemEventNotificationV02 struct {
	XMLName xml.Name`,
				new: "type SystemEventNotificationV02 struct {\n\tXMLName xml.Name `xml:\"SysEvtNtfctn\"`",
			},
			{
				old: "\tEvtInf Event2 `xml:\",any\"`",
				new: "\tEvtInf Event2 `xml:\"EvtInf\"`",
			},
		},
	},
	{
		path: "gen/admi_006_001_01/models.go",
		replacements: []replacement{
			{
				old: "\tRsndReq ResendRequestV01 `xml:\",any\"`",
				new: "\tRsndReq ResendRequestV01 `xml:\"RsndReq\"`",
			},
			{
				old: `type ResendRequestV01 struct {
	XMLName xml.Name`,
				new: "type ResendRequestV01 struct {\n\tXMLName xml.Name `xml:\"RsndReq\"`",
			},
		},
	},
	{
		path: "gen/admi_007_001_01/models.go",
		replacements: []replacement{
			{
				old: "\tRctAck ReceiptAcknowledgementV01 `xml:\",any\"`",
				new: "\tRctAck ReceiptAcknowledgementV01 `xml:\"RctAck\"`",
			},
			{
				old: `type ReceiptAcknowledgementV01 struct {
	XMLName xml.Name`,
				new: "type ReceiptAcknowledgementV01 struct {\n\tXMLName xml.Name `xml:\"RctAck\"`",
			},
		},
	},
	{
		path: "gen/admi_011_001_01/models.go",
		replacements: []replacement{
			{
				old: "\tSysEvtAck SystemEventAcknowledgementV01 `xml:\",any\"`",
				new: "\tSysEvtAck SystemEventAcknowledgementV01 `xml:\"SysEvtAck\"`",
			},
			{
				old: `type SystemEventAcknowledgementV01 struct {
	XMLName xml.Name`,
				new: "type SystemEventAcknowledgementV01 struct {\n\tXMLName xml.Name `xml:\"SysEvtAck\"`",
			},
		},
	},
	{
		path: "gen/admi_998_001_02/models.go",
		replacements: []replacement{
			{
				old: "\tAdmstnPrtryMsg AdministrationProprietaryMessageV02 `xml:\",any\"`",
				new: "\tAdmstnPrtryMsg AdministrationProprietaryMessageV02 `xml:\"AdmstnPrtryMsg\"`",
			},
			{
				old: `type AdministrationProprietaryMessageV02 struct {
	XMLName xml.Name`,
				new: "type AdministrationProprietaryMessageV02 struct {\n\tXMLName xml.Name `xml:\"AdmstnPrtryMsg\"`",
			},
			{
				old: "\tRef Max35Text `xml:\",any\"`",
				new: "\tRef Max35Text `xml:\"Ref\"`",
			},
		},
	},
	{
		path: "gen/camt_026_001_07/models.go",
		replacements: []replacement{
			{
				old: "\tUblToApply UnableToApplyV07 `xml:\",any\"`",
				new: "\tUblToApply UnableToApplyV07 `xml:\"UblToApply\"`",
			},
			{
				old: `type UnableToApplyV07 struct {
	XMLName xml.Name`,
				new: "type UnableToApplyV07 struct {\n\tXMLName xml.Name `xml:\"UblToApply\"`",
			},
		},
	},
	{
		path: "gen/camt_028_001_09/models.go",
		replacements: []replacement{
			{
				old: "\tAddtlPmtInf AdditionalPaymentInformationV09 `xml:\",any\"`",
				new: "\tAddtlPmtInf AdditionalPaymentInformationV09 `xml:\"AddtlPmtInf\"`",
			},
			{
				old: `type AdditionalPaymentInformationV09 struct {
	XMLName xml.Name`,
				new: "type AdditionalPaymentInformationV09 struct {\n\tXMLName xml.Name `xml:\"AddtlPmtInf\"`",
			},
		},
	},
	{
		path: "gen/camt_029_001_09/models.go",
		replacements: []replacement{
			{
				old: "\tRsltnOfInvstgtn ResolutionOfInvestigationV09 `xml:\",any\"`",
				new: "\tRsltnOfInvstgtn ResolutionOfInvestigationV09 `xml:\"RsltnOfInvstgtn\"`",
			},
			{
				old: `type ResolutionOfInvestigationV09 struct {
	XMLName xml.Name`,
				new: "type ResolutionOfInvestigationV09 struct {\n\tXMLName xml.Name `xml:\"RsltnOfInvstgtn\"`",
			},
		},
	},
	{
		path: "gen/camt_052_001_08/models.go",
		replacements: []replacement{
			{
				old: "\tBkToCstmrAcctRpt BankToCustomerAccountReportV08 `xml:\",any\"`",
				new: "\tBkToCstmrAcctRpt BankToCustomerAccountReportV08 `xml:\"BkToCstmrAcctRpt\"`",
			},
			{
				old: `type BankToCustomerAccountReportV08 struct {
	XMLName xml.Name`,
				new: "type BankToCustomerAccountReportV08 struct {\n\tXMLName xml.Name `xml:\"BkToCstmrAcctRpt\"`",
			},
		},
	},
	{
		path: "gen/camt_054_001_08/models.go",
		replacements: []replacement{
			{
				old: "\tBkToCstmrDbtCdtNtfctn BankToCustomerDebitCreditNotificationV08 `xml:\",any\"`",
				new: "\tBkToCstmrDbtCdtNtfctn BankToCustomerDebitCreditNotificationV08 `xml:\"BkToCstmrDbtCdtNtfctn\"`",
			},
			{
				old: `type BankToCustomerDebitCreditNotificationV08 struct {
	XMLName xml.Name`,
				new: "type BankToCustomerDebitCreditNotificationV08 struct {\n\tXMLName xml.Name `xml:\"BkToCstmrDbtCdtNtfctn\"`",
			},
		},
	},
	{
		path: "gen/camt_055_001_09/models.go",
		replacements: []replacement{
			{
				old: "\tCstmrPmtCxlReq CustomerPaymentCancellationRequestV09 `xml:\",any\"`",
				new: "\tCstmrPmtCxlReq CustomerPaymentCancellationRequestV09 `xml:\"CstmrPmtCxlReq\"`",
			},
			{
				old: `type CustomerPaymentCancellationRequestV09 struct {
	XMLName xml.Name`,
				new: "type CustomerPaymentCancellationRequestV09 struct {\n\tXMLName xml.Name `xml:\"CstmrPmtCxlReq\"`",
			},
		},
	},
	{
		path: "gen/camt_056_001_08/models.go",
		replacements: []replacement{
			{
				old: "\tFIToFIPmtCxlReq FIToFIPaymentCancellationRequestV08 `xml:\",any\"`",
				new: "\tFIToFIPmtCxlReq FIToFIPaymentCancellationRequestV08 `xml:\"FIToFIPmtCxlReq\"`",
			},
			{
				old: `type FIToFIPaymentCancellationRequestV08 struct {
	XMLName xml.Name`,
				new: "type FIToFIPaymentCancellationRequestV08 struct {\n\tXMLName xml.Name `xml:\"FIToFIPmtCxlReq\"`",
			},
		},
	},
	{
		path: "gen/camt_060_001_05/models.go",
		replacements: []replacement{
			{
				old: "\tAcctRptgReq AccountReportingRequestV05 `xml:\",any\"`",
				new: "\tAcctRptgReq AccountReportingRequestV05 `xml:\"AcctRptgReq\"`",
			},
			{
				old: `type AccountReportingRequestV05 struct {
	XMLName xml.Name`,
				new: "type AccountReportingRequestV05 struct {\n\tXMLName xml.Name `xml:\"AcctRptgReq\"`",
			},
		},
	},
	{
		path: "gen/pacs_002_001_10/models.go",
		replacements: []replacement{
			{
				old: "\tFIToFIPmtStsRpt FIToFIPaymentStatusReportV10 `xml:\",any\"`",
				new: "\tFIToFIPmtStsRpt FIToFIPaymentStatusReportV10 `xml:\"FIToFIPmtStsRpt\"`",
			},
			{
				old: `type FIToFIPaymentStatusReportV10 struct {
	XMLName xml.Name`,
				new: "type FIToFIPaymentStatusReportV10 struct {\n\tXMLName xml.Name `xml:\"FIToFIPmtStsRpt\"`",
			},
		},
	},
	{
		path: "gen/pacs_004_001_10/models.go",
		replacements: []replacement{
			{
				old: "\tPmtRtr PaymentReturnV10 `xml:\",any\"`",
				new: "\tPmtRtr PaymentReturnV10 `xml:\"PmtRtr\"`",
			},
			{
				old: `type PaymentReturnV10 struct {
	XMLName xml.Name`,
				new: "type PaymentReturnV10 struct {\n\tXMLName xml.Name `xml:\"PmtRtr\"`",
			},
		},
	},
	{
		path: "gen/pacs_008_001_08/models.go",
		replacements: []replacement{
			{
				old: "\tFIToFICstmrCdtTrf FIToFICustomerCreditTransferV08 `xml:\",any\"`",
				new: "\tFIToFICstmrCdtTrf FIToFICustomerCreditTransferV08 `xml:\"FIToFICstmrCdtTrf\"`",
			},
			{
				old: `type FIToFICustomerCreditTransferV08 struct {
	XMLName xml.Name`,
				new: "type FIToFICustomerCreditTransferV08 struct {\n\tXMLName xml.Name `xml:\"FIToFICstmrCdtTrf\"`",
			},
		},
	},
	{
		path: "gen/pacs_009_001_08/models.go",
		replacements: []replacement{
			{
				old: "\tFICdtTrf FinancialInstitutionCreditTransferV08 `xml:\",any\"`",
				new: "\tFICdtTrf FinancialInstitutionCreditTransferV08 `xml:\"FICdtTrf\"`",
			},
			{
				old: `type FinancialInstitutionCreditTransferV08 struct {
	XMLName xml.Name`,
				new: "type FinancialInstitutionCreditTransferV08 struct {\n\tXMLName xml.Name `xml:\"FICdtTrf\"`",
			},
		},
	},
	{
		path: "gen/pacs_028_001_03/models.go",
		replacements: []replacement{
			{
				old: "\tFIToFIPmtStsReq FIToFIPaymentStatusRequestV03 `xml:\",any\"`",
				new: "\tFIToFIPmtStsReq FIToFIPaymentStatusRequestV03 `xml:\"FIToFIPmtStsReq\"`",
			},
			{
				old: `type FIToFIPaymentStatusRequestV03 struct {
	XMLName xml.Name`,
				new: "type FIToFIPaymentStatusRequestV03 struct {\n\tXMLName xml.Name `xml:\"FIToFIPmtStsReq\"`",
			},
		},
	},
	{
		path: "gen/pain_013_001_07/models.go",
		replacements: []replacement{
			{
				old: "\tCdtrPmtActvtnReq CreditorPaymentActivationRequestV07 `xml:\",any\"`",
				new: "\tCdtrPmtActvtnReq CreditorPaymentActivationRequestV07 `xml:\"CdtrPmtActvtnReq\"`",
			},
			{
				old: `type CreditorPaymentActivationRequestV07 struct {
	XMLName xml.Name`,
				new: "type CreditorPaymentActivationRequestV07 struct {\n\tXMLName xml.Name `xml:\"CdtrPmtActvtnReq\"`",
			},
		},
	},
	{
		path: "gen/pain_014_001_07/models.go",
		replacements: []replacement{
			{
				old: "\tCdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV07 `xml:\",any\"`",
				new: "\tCdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV07 `xml:\"CdtrPmtActvtnReqStsRpt\"`",
			},
			{
				old: `type CreditorPaymentActivationRequestStatusReportV07 struct {
	XMLName xml.Name`,
				new: "type CreditorPaymentActivationRequestStatusReportV07 struct {\n\tXMLName xml.Name `xml:\"CdtrPmtActvtnReqStsRpt\"`",
			},
		},
	},
	{
		path: "gen/sup_FedNowParticipantFile_admi_998_001_02/models.go",
		replacements: []replacement{
			{
				old: "\tAdmi998SuplDataV01 Admi998SuplDataV01 `xml:\",any\"`",
				new: "\tAdmi998SuplDataV01 Admi998SuplDataV01 `xml:\"admi998SuplDataV01\"`",
			},
			{
				old: `type FedNowParticipantFile1 struct {
	XMLName xml.Name`,
				new: "type FedNowParticipantFile1 struct {\n\tXMLName xml.Name `xml:\"PtcptFile\"`",
			},
			{
				old: `type Admi998SuplDataV01 struct {
	XMLName xml.Name`,
				new: "type Admi998SuplDataV01 struct {\n\tXMLName xml.Name `xml:\"admi998SuplDataV01\"`",
			},
			{
				old: "\tPtcptFile FedNowParticipantFile1 `xml:\",any\"`",
				new: "\tPtcptFile FedNowParticipantFile1 `xml:\"PtcptFile\"`",
			},
		},
	},
}

func main() {
	for _, f := range fixups {
		data, err := os.ReadFile(f.path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: reading %s: %v\n", f.path, err)
			os.Exit(1)
		}
		content := string(data)
		for _, r := range f.replacements {
			if n := strings.Count(content, r.old); n != 1 {
				fmt.Fprintf(os.Stderr, "ERROR: %s: found %d matches, want 1:\n%s\n", f.path, n, r.old)
				os.Exit(1)
			}
			content = strings.Replace(content, r.old, r.new, 1)
		}
		if err := os.WriteFile(f.path, []byte(content), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: writing %s: %v\n", f.path, err)
			os.Exit(1)
		}
	}
	if err := removeInnerXML("gen"); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("XML name fixups completed.")
}

const innerXMLField = "\tinnerXml string `xml:\",innerxml\"`\n"

// removeInnerXML removes the unexported innerXml field from every generated
// models.go file under dir. Run gofmt after it.
func removeInnerXML(dir string) error {
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Base(path) != "models.go" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		content := string(data)
		if !strings.Contains(content, innerXMLField) {
			return nil
		}
		content = strings.ReplaceAll(content, "\n"+innerXMLField, "")
		content = strings.ReplaceAll(content, innerXMLField, "")
		return os.WriteFile(path, []byte(content), 0644)
	})
}
