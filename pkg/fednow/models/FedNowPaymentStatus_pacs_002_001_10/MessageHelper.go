package FedNowPaymentStatus_pacs_002_001_10

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type MessageHelper struct {
	MessageId                        common.ElementHelper
	CreatedDateTime                  common.ElementHelper
	GroupInfo                        common.GroupInformationHelper
	OriginalTransactionInfo          common.TransactionDetailReferenceHelper
	TransactionStatus                common.ElementHelper
	StatusReasonInformation          common.ReasonHelper // This field is not in the original document, but can be added if needed
	AcceptanceDateTime               common.ElementHelper
	EffectiveInterbankSettlementDate common.ElementHelper
	Agents                           common.RelatedAgentsHelper
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
		GroupInfo:               common.BuildGroupInformationHelper(),
		OriginalTransactionInfo: common.BuildTransactionDetailReferenceHelper(),
		TransactionStatus: common.ElementHelper{
			Title:         "Transaction Status",
			Rules:         "The status of the transaction.",
			Type:          `TransactionStatusCode`,
			Documentation: `The status of the transaction.`,
		},
		StatusReasonInformation: common.BuildReasonHelper(), // This field is not in the original document, but can be added if needed
		AcceptanceDateTime: common.ElementHelper{
			Title:         "Acceptance DateTime",
			Rules:         "The date and time when the transaction was accepted.",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `The date and time when the transaction was accepted.`,
		},
		EffectiveInterbankSettlementDate: common.ElementHelper{
			Title: "Effective Interbank Settlement Date",
			Rules: "",
			Type:  `ISODate (based on date)`,
			Documentation: `Date, as provided in the original transaction, on which the amount of
			 money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.`,
		},
		Agents: common.BuildRelatedAgentsHelper(),
	}

}
