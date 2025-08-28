package ReceiptAcknowledgement_admi_007_001_01

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type MessageHelper struct {
	MessageId           common.ElementHelper
	CreatedDateTime     common.ElementHelper
	RelatedReference    common.ElementHelper
	MessageName         common.ElementHelper
	RequestHandlingCode common.ElementHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: common.ElementHelper{
			Title:         "Message ID",
			Rules:         "This is the unique identifier for the message.",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identifier for the message.`,
		},
		CreatedDateTime: common.ElementHelper{
			Title:         "Created Date Time",
			Rules:         "This is the date and time when the message was created.",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time when the message was created.`,
		},
		RelatedReference: common.ElementHelper{
			Title:         "Related Reference",
			Rules:         "This is the reference to the related message.",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unambiguous reference to a previous message having a business relevance with this message.`,
		},
		MessageName: common.ElementHelper{
			Title:         "Message Name",
			Rules:         "This is the name of the message being acknowledged.",
			Type:          `MessageNameIdentification_FRS_1 (based on string) exactLength: 15 pattern: [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}`,
			Documentation: `Name of the message which contained the given additional reference as its message reference.`,
		},
		RequestHandlingCode: common.ElementHelper{
			Title:         "Request Handling Code",
			Rules:         "This is the code indicating how the request should be handled.",
			Type:          `MessageHandlingCode`,
			Documentation: `Specifies the status of the request, for example the result of the schema validation or a business processing result/error.`,
		},
	}

}
