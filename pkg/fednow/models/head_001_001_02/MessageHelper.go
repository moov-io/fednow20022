package head_001_001_02

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type RelatedHelper struct {
	FromMemberId                common.ElementHelper
	ToMemberId                  common.ElementHelper
	BusinessMessageIdentifier   common.ElementHelper
	MessageDefinitionIdentifier common.ElementHelper
	CreatedDateTime             common.ElementHelper
}

func BuildRelatedHelper() RelatedHelper {
	return RelatedHelper{
		FromMemberId: common.ElementHelper{
			Title:         "From Member ID",
			Rules:         "",
			Type:          `pattern: [A-Z0-9]{9,9}`,
			Documentation: `The sending MessagingEndpoint that has created this Business Message for the receiving MessagingEndpoint that will process this Business Message.`,
		},
		ToMemberId: common.ElementHelper{
			Title:         "To Member ID",
			Rules:         "",
			Type:          `pattern: [A-Z0-9]{9,9}`,
			Documentation: `The MessagingEndpoint designated by the sending MessagingEndpoint to be the recipient who will ultimately process this Business Message.`,
		},
		BusinessMessageIdentifier: common.ElementHelper{
			Title:         "Business Message Identifier",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unambiguously identifies the Business Message to the MessagingEndpoint that has created the Business Message.`,
		},
		MessageDefinitionIdentifier: common.ElementHelper{
			Title:         "Message Definition Identifier",
			Rules:         "",
			Type:          `MessageNameIdentification_FRS_1 (based on string) exactLength: 15 pattern: [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}`,
			Documentation: `Contains the MessageIdentifier that defines the BusinessMessage. It must contain a MessageIdentifier published on the ISO 20022 website.`,
		},
		CreatedDateTime: common.ElementHelper{
			Title:         "Created Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time when this Business Message (header) was created. Note Times must be normalized, using the "Z" annotation.`,
		},
	}
}

type MessageHelper struct {
	FromMemberId                common.ElementHelper
	ToMemberId                  common.ElementHelper
	CreatedDateTime             common.ElementHelper
	BusinessMessageIdentifier   common.ElementHelper
	MessageDefinitionIdentifier common.ElementHelper
	MarketPractice              common.MarketPracticeHelper
	BusinessProcessingDate      common.ElementHelper
	CopyDuplicate               common.ElementHelper // Indicates if the message is a copy or duplicate
	RelatedInfo                 RelatedHelper        // Additional related information
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		FromMemberId: common.ElementHelper{
			Title:         "From Member ID",
			Rules:         "",
			Type:          `pattern: [A-Z0-9]{9,9}`,
			Documentation: `The sending MessagingEndpoint that has created this Business Message for the receiving MessagingEndpoint that will process this Business Message.`,
		},
		ToMemberId: common.ElementHelper{
			Title:         "To Member ID",
			Rules:         "",
			Type:          `pattern: [A-Z0-9]{9,9}`,
			Documentation: `The MessagingEndpoint designated by the sending MessagingEndpoint to be the recipient who will ultimately process this Business Message.`,
		},
		BusinessMessageIdentifier: common.ElementHelper{
			Title:         "Business Message Identifier",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unambiguously identifies the Business Message to the MessagingEndpoint that has created the Business Message.`,
		},
		MessageDefinitionIdentifier: common.ElementHelper{
			Title:         "Message Definition Identifier",
			Rules:         "",
			Type:          `MessageNameIdentification_FRS_1 (based on string) exactLength: 15 pattern: [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}`,
			Documentation: `Contains the MessageIdentifier that defines the BusinessMessage. It must contain a MessageIdentifier published on the ISO 20022 website.`,
		},
		CreatedDateTime: common.ElementHelper{
			Title:         "Created Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time when this Business Message (header) was created. Note Times must be normalized, using the "Z" annotation.`,
		},
		MarketPractice: common.BuildMarketPracticeHelper(),
		BusinessProcessingDate: common.ElementHelper{
			Title:         "Business Processing Date",
			Rules:         "",
			Type:          `ISODate (based on date)`,
			Documentation: `Processing date and time indicated by the sender for the receiver of the business message. This date may be different from the date and time provided in the CreationDate.`,
		},
		CopyDuplicate: common.ElementHelper{
			Title:         "Copy Duplicate",
			Rules:         "",
			Type:          `CopyDuplicateCode`,
			Documentation: `Indicates whether the message is a Copy, a Duplicate or a copy of a duplicate of a previously sent ISO 20022 Message.`,
		},
		RelatedInfo: BuildRelatedHelper(),
	}

}
