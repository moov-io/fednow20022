package ReturnRequestResponse_camt_029_001_09

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type ResolutionRelatedInformationHelper struct {
	TransactionId             common.ElementHelper
	UETR                      common.ElementHelper
	InterbankSettlementAmount common.CurrencyAndAmountHelper
	InterbankSettlementDate   common.ElementHelper
}

func BuildResolutionRelatedInformationHelper() ResolutionRelatedInformationHelper {
	return ResolutionRelatedInformationHelper{
		TransactionId: common.ElementHelper{
			Title:         "Transaction Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification of the transaction.`,
		},
		UETR: common.ElementHelper{
			Title:         "Unique End-to-End Transaction Reference",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique end-to-end transaction reference.`,
		},
		InterbankSettlementAmount: common.BuildCurrencyAndAmountHelper(),
		InterbankSettlementDate: common.ElementHelper{
			Title:         "Interbank Settlement Date",
			Rules:         "",
			Type:          `ISODate (based on date)`,
			Documentation: `Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.`,
		},
	}
}

type MessageHelper struct {
	MessageId       common.ElementHelper
	CreatedDateTime common.ElementHelper
	Assigners       common.AssignmentsHelper
	ResolvedCaseId  common.ElementHelper
	Creator         common.AgentHelper
	Status          common.ElementHelper
	GroupInfo       common.GroupInformationHelper
	TransactionInfo common.TransactionDetailReferenceHelper
	ReturnReason    common.ReturnReasonHelper
	RelatedInfo     ResolutionRelatedInformationHelper
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
		Assigners: common.BuildAssignmentsHelper(),
		ResolvedCaseId: common.ElementHelper{
			Title:         "Resolved Case Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification of the resolved case.`,
		},
		Creator: common.BuildAgentHelper(),
		Status: common.ElementHelper{
			Title:         "Status",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Current status of the message.`,
		},
		GroupInfo:       common.BuildGroupInformationHelper(),
		TransactionInfo: common.BuildTransactionDetailReferenceHelper(),
		ReturnReason:    common.BuildReturnReasonHelper(),
		RelatedInfo:     BuildResolutionRelatedInformationHelper(),
	}

}
