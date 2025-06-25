package AccountReportingRequest_camt_060_001_05

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type MessageHelper struct {
	MessageId              common.ElementHelper
	CreatedDateTime        common.ElementHelper
	ReportingRequestId     common.ElementHelper
	RequestedMessageNameId common.ElementHelper
	AccountOtherId         common.ElementHelper
	AccountType            common.ElementHelper
	AccountOwner           common.AgentHelper
	ReportingPeriod        common.PeriodDateAndTimeHelper
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
		ReportingRequestId: common.ElementHelper{
			Title:         "Reporting Request Id",
			Rules:         "",
			Type:          `CAMTReportType`,
			Documentation: `The codes are used to identify the requested report. `,
		},
		RequestedMessageNameId: common.ElementHelper{
			Title:         "Requested Message Name Id",
			Rules:         "",
			Type:          `MessageNameIdentification_FRS_1 (based on string) exactLength: 15 pattern: [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}`,
			Documentation: `Specifies the type of the requested reporting message.`,
		},
		AccountOtherId: common.ElementHelper{
			Title:         "Account Other Id",
			Rules:         "This is the routing number to which the activity report relates.",
			Type:          `RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}`,
			Documentation: `Identification assigned by an institution.`,
		},
		AccountType: common.ElementHelper{
			Title:         "Account Type",
			Rules:         "",
			Type:          `AccountType`,
			Documentation: `Specifies the nature, or use of the account.`,
		},
		AccountOwner:    common.BuildAgentHelper(),
		ReportingPeriod: common.BuildPeriodDateAndTimeHelper(),
	}

}
