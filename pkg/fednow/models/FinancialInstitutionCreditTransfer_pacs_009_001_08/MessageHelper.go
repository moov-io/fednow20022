package FinancialInstitutionCreditTransfer_pacs_009_001_08

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type MessageHelper struct {
	MessageId                 common.ElementHelper
	CreatedDateTime           common.ElementHelper
	NumberOfTransactions      common.ElementHelper
	SettlementInfo            common.SettlementInformationHelper
	PaymentIdentification     common.TransactionDetailReferenceHelper
	PaymentType               common.PaymentTypeInfoHelper
	InterbankSettlementAmount common.CurrencyAndAmountHelper
	InterbankSettlementDate   common.ElementHelper
	Agents                    common.RelatedAgentsHelper
	Debtor                    common.TransactionPartyHelper
	Creditor                  common.TransactionPartyHelper
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
		NumberOfTransactions: common.ElementHelper{
			Title:         "Number of Transactions",
			Rules:         "The number of transactions in the message.",
			Type:          `Max15NumericText (based on string) minLength: 1 maxLength: 15`,
			Documentation: `The number of transactions in the message.`,
		},
		SettlementInfo:            common.BuildSettlementInformationHelper(),
		PaymentIdentification:     common.BuildTransactionDetailReferenceHelper(),
		PaymentType:               common.BuildPaymentTypeInfoHelper(),
		InterbankSettlementAmount: common.BuildCurrencyAndAmountHelper(),
		InterbankSettlementDate: common.ElementHelper{
			Title:         "Interbank Settlement Date",
			Rules:         "",
			Type:          `ISODate (based on date)`,
			Documentation: `Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.`,
		},
		Agents:   common.BuildRelatedAgentsHelper(),
		Debtor:   common.BuildTransactionPartyHelper(),
		Creditor: common.BuildTransactionPartyHelper(),
	}

}
