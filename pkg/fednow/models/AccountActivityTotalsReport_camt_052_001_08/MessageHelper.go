package AccountActivityTotalsReport_camt_052_001_08

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type ReportModelHelper struct {
	ReportId                           common.ElementHelper
	ReportCreateDateTime               common.ElementHelper
	AccountOtherId                     common.ElementHelper
	RelateAccountOtherId               common.ElementHelper
	TotalEntries                       common.ElementHelper
	TotalCreditEntries                 common.NumberAndSumOfTransactionsHelper
	TotalDebitEntries                  common.NumberAndSumOfTransactionsHelper
	TotalEntriesPerBankTransactionCode common.TotalsPerBankTransactionCodeHelper
}

func BuildReportModelHelper() ReportModelHelper {
	return ReportModelHelper{
		ReportId: common.ElementHelper{
			Title:         "Report Type Id",
			Rules:         "",
			Type:          `ReportType(EveryDay, Intraday)`,
			Documentation: `Unique identification, as assigned by the account servicer, to unambiguously identify the account report.`,
		},
		ReportCreateDateTime: common.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "This is the Fedwire Funds Service funds-transfer business day. Note: Time will be defaulted to 00:00:00 in New York City (Eastern Time) with the offset against the Coordinated Universal Time (UTC).",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the report was created.`,
		},
		AccountOtherId: common.ElementHelper{
			Title:         "Account Other Id",
			Rules:         "his is the routing number to which the activity report relates.",
			Type:          `RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}`,
			Documentation: `Identification assigned by an institution.`,
		},
		RelateAccountOtherId: common.ElementHelper{
			Title:         "Related Account Other Id",
			Rules:         "",
			Type:          `RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}`,
			Documentation: `This component only occurs in a Correspondent Account Activity Details Report (CADR) to identify the correspondent's respondent for which activity details are reported. The component is not present in a regular Account Activity Details Report (AADR).`,
		},
		TotalEntries: common.ElementHelper{
			Title:         "Number Of Entries",
			Rules:         "",
			Type:          `Max15NumericText (based on string) pattern: [0-9]{1,15}`,
			Documentation: `Number of individual entries included in the report.`,
		},
		TotalCreditEntries:                 common.BuildNumberAndSumOfTransactionsHelper(),
		TotalDebitEntries:                  common.BuildNumberAndSumOfTransactionsHelper(),
		TotalEntriesPerBankTransactionCode: common.BuildTotalsPerBankTransactionCodeHelper(),
	}
}

type MessageHelper struct {
	MessageId             common.ElementHelper
	CreatedDateTime       common.ElementHelper
	Pagenation            common.MessagePagenationHelper
	OriginalBusinessQuery common.OriginalBusinessQueryHelper // Optional field for original business query
	Reports               ReportModelHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: common.ElementHelper{
			Title:         "Message Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message. Usage: The account servicing institution has to make sure that MessageIdentification is unique per account owner for a pre-agreed period.`,
		},
		CreatedDateTime: common.ElementHelper{
			Title:         "Message Identification",
			Rules:         "This is the calendar date and time in New York City (Eastern Time) when the message is created by the Fedwire Funds Service application. Time is in 24-hour clock format and includes the offset against the Coordinated Universal Time (UTC).",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		Pagenation:            common.BuildMessagePagenationHelper(),
		OriginalBusinessQuery: common.BuildOriginalBusinessQueryHelper(), // Optional field for original business query
		Reports:               BuildReportModelHelper(),
	}

}
