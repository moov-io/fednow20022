package MessageReject_admi_002_001_01

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type MessageHelper struct {
	RelatedReference     common.ElementHelper
	RejectingPartyReason common.ElementHelper
	RejectionDateTime    common.ElementHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		RelatedReference: common.ElementHelper{
			Title:         "Related Reference",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `This should be the reference of the original message that is rejected. `,
		},
		RejectingPartyReason: common.ElementHelper{
			Title:         "Rejecting Party Reason",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Reason of the rejection provided by the rejecting party.`,
		},
		RejectionDateTime: common.ElementHelper{
			Title:         "Rejection Date Time",
			Rules:         "This is the calendar date and time in New York City (Eastern Time) when the event occurred. Time is in 24-hour clock format and includes the offset against	 Coordinated Universal Time (UTC).",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the event occurred.`,
		},
	}

}
