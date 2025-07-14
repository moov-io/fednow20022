package RetrievalRequest_admi_006_001_01

// type ResendSearchCriteriaHelper struct {
// 	OriginalMessageNameId common.ElementHelper
// 	BusinessDate          common.ElementHelper
// 	FileReference         common.ElementHelper
// 	RecipientId           common.ElementHelper
// 	RecipientIssuer       common.ElementHelper
// }

// func BuildResendSearchCriteriaHelper() ResendSearchCriteriaHelper {
// 	return ResendSearchCriteriaHelper{
// 		OriginalMessageNameId: common.ElementHelper{
// 			Title:         "Original Message Name ID",
// 			Rules:         "This is the original message identification.",
// 			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
// 			Documentation: `Identification of the original message.`,
// 		},
// 		BusinessDate: common.ElementHelper{
// 			Title:         "Business Date",
// 			Rules:         "This is the business date for the request.",
// 			Type:          `ISODate (based on date)`,
// 			Documentation: `The business date for the request.`,
// 		},
// 		FileReference: common.ElementHelper{
// 			Title:         "File Reference",
// 			Rules:         "This is the reference to the file.",
// 			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
// 			Documentation: `Reference to the file.`,
// 		},
// 		RecipientId: common.ElementHelper{
// 			Title:         "Recipient ID",
// 			Rules:         "This is the identification of the recipient.",
// 			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
// 			Documentation: `Identification of the recipient.`,
// 		},
// 		RecipientIssuer: common.ElementHelper{
// 			Title:         "Recipient Issuer",
// 			Rules:         "This is the issuer of the recipient identification.",
// 			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
// 			Documentation: `Issuer of the recipient identification.`,
// 		},
// 	}
// }

// type MessageHelper struct {
// 	MessageId       common.ElementHelper
// 	CreatedDateTime common.ElementHelper
// 	Criteriaes      ResendSearchCriteriaHelper
// }

// func BuildMessageHelper() MessageHelper {
// 	return MessageHelper{
// 		MessageId: common.ElementHelper{
// 			Title:         "Message ID",
// 			Rules:         "This is the unique identifier for the message.",
// 			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
// 			Documentation: `Unique identifier for the message.`,
// 		},
// 		CreatedDateTime: common.ElementHelper{
// 			Title:         "Created Date Time",
// 			Rules:         "This is the date and time when the message was created.",
// 			Type:          `ISODateTime (based on dateTime)`,
// 			Documentation: `Date and time when the message was created.`,
// 		},
// 		Criteriaes: BuildResendSearchCriteriaHelper(),
// 	}
// }
