package AdditionalPaymentInformation_camt_028_001_09

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type MessageHelper struct {
	MessageId                         common.ElementHelper
	Assignment                        common.AssignmentsHelper
	CreatedDateTime                   common.ElementHelper
	CaseId                            common.ElementHelper
	Creator                           common.AgentHelper
	OriginalGroupInfo                 common.GroupInformationHelper
	OriginalTransactionInfo           common.TransactionDetailReferenceHelper
	OriginalInterbankSettlementAmount common.CurrencyAndAmountHelper
	OriginalInterbankSettlementDate   common.ElementHelper
	Information                       common.ElementHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: common.ElementHelper{
			Title:         "Message Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message. Usage: The account servicing institution has to make sure that MessageIdentification is unique per account owner for a pre-agreed period.`,
		},
		Assignment: common.BuildAssignmentsHelper(),
		CreatedDateTime: common.ElementHelper{
			Title:         "Message Identification",
			Rules:         "This is the calendar date and time in New York City (Eastern Time) when the message is created by the Fedwire Funds Service application. Time is in 24-hour clock format and includes the offset against the Coordinated Universal Time (UTC).",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		CaseId: common.ElementHelper{
			Title:         "Case Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Uniquely identifies the case.`,
		},
		Creator:                           common.BuildAgentHelper(),
		OriginalGroupInfo:                 common.BuildGroupInformationHelper(),
		OriginalTransactionInfo:           common.BuildTransactionDetailReferenceHelper(),
		OriginalInterbankSettlementAmount: common.BuildCurrencyAndAmountHelper(),
		OriginalInterbankSettlementDate: common.ElementHelper{
			Title:         "Original Interbank Settlement Date",
			Rules:         "",
			Type:          `ISODate (based on date)`,
			Documentation: `Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.`,
		},
		Information: common.ElementHelper{
			Title:         "Information",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.`,
		},
	}

}
