package ParticipantPaymentStatus_pacs_002_001_10

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type MessageHelper struct {
	MessageId               common.ElementHelper
	CreatedDateTime         common.ElementHelper
	OriginalGroupInfo       common.GroupInformationHelper
	TransactionInfo         common.TransactionDetailReferenceHelper
	TransactionStatus       common.ElementHelper
	StatusReasonInformation common.ReasonHelper
	Agents                  common.RelatedAgentsHelper
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
		OriginalGroupInfo: common.BuildGroupInformationHelper(),
		TransactionInfo:   common.BuildTransactionDetailReferenceHelper(),
		TransactionStatus: common.ElementHelper{
			Title:         "Transaction Status",
			Rules:         "This is the status of the transaction.",
			Type:          `TransactionStatusCode`,
			Documentation: `Specifies the status of a transaction, in a coded form.`,
		},
		StatusReasonInformation: common.BuildReasonHelper(),
		Agents:                  common.BuildRelatedAgentsHelper(),
	}

}
