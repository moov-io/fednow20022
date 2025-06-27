package FedNowBroadcast_admi_004_001_02

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type MessageHelper struct {
	EventCode        common.ElementHelper
	EventParameters  common.ElementHelper
	EventDescription common.ElementHelper
	EventTime        common.ElementHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		EventCode: common.ElementHelper{
			Title:         "Event Code",
			Rules:         "",
			Type:          `FundEventType`,
			Documentation: `Proprietary code used to specify an event that occurred in a system.`,
		},
		EventParameters: common.ElementHelper{
			Title:         "Event Parameters",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Describes the parameters of an event which occurred in a system.`,
		},
		EventDescription: common.ElementHelper{
			Title:         "Event Description",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 max	Length: 140`,
			Documentation: `Free text used to describe an event which occurred in a system.`,
		},
		EventTime: common.ElementHelper{
			Title:         "Event Time",
			Rules:         "This is the calendar date and time in New York City (Eastern Time) when the event occurred. Time is in 24-hour clock format and includes the offset against	 Coordinated Universal Time (UTC).",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the event occurred.`,
		},
	}

}
