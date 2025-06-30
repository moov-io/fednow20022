package FedNowSystemResponse_admi_011_001_01

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type MessageHelper struct {
	MessageId      common.ElementHelper
	EventCode      common.ElementHelper
	EventParameter common.ElementHelper
	EventTime      common.ElementHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: common.ElementHelper{
			Title:         "Message ID",
			Rules:         "This is a unique identifier for the message.",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identifier for the message.`,
		},
		EventCode: common.ElementHelper{
			Title:         "Event Code",
			Rules:         "",
			Type:          `FundEventType`,
			Documentation: `Proprietary code used to specify an event that occurred in a system.`,
		},
		EventParameter: common.ElementHelper{
			Title:         "Event Parameter",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Describes the parameters of an event which occurred in a system.`,
		},

		EventTime: common.ElementHelper{
			Title:         "Event Time",
			Rules:         "This is the calendar date and time in New York City (Eastern Time) when the event occurred. Time is in 24-hour clock format and includes the offset against	 Coordinated Universal Time (UTC).",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the event occurred.`,
		},
	}

}
