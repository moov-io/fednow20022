package AccountDebitCreditNotification_camt_054_001_08

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type MessageHelper struct {
	MessageId                      common.ElementHelper
	CreatedDateTime                common.ElementHelper
	Pagenation                     common.MessagePagenationHelper
	NotificationId                 common.ElementHelper
	AccountOtherId                 common.ElementHelper
	EntryAmount                    common.CurrencyAndAmountHelper
	EntryCreditDebitIndicator      common.ElementHelper
	EntryStatus                    common.ElementHelper
	BankTransactionCode            common.ElementHelper
	AdditionalInformationIndicator common.ElementHelper
	TransactionDetailRefer         common.TransactionDetailReferenceHelper
	RelatedAgent                   common.RelatedAgentsHelper
	LocalInstrument                common.ElementHelper
	RelatedDate                    common.RelatedDatesHelper
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
		Pagenation: common.BuildMessagePagenationHelper(),
		NotificationId: common.ElementHelper{
			Title:         "Notification Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the account servicer, to unambiguously identify the account notification.`,
		},
		AccountOtherId: common.ElementHelper{
			Title:         "Account Other Id",
			Rules:         "his is the routing number to which the activity report relates.",
			Type:          `RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}`,
			Documentation: `Identification assigned by an institution.`,
		},
		EntryAmount: common.BuildCurrencyAndAmountHelper(),
		EntryCreditDebitIndicator: common.ElementHelper{
			Title:         "Entry CreditDebit Indicator",
			Rules:         "",
			Type:          `CreditDebitCode (based on string)`,
			Documentation: `Indicates whether the entry is a credit or a debit entry.`,
		},
		EntryStatus: common.ElementHelper{
			Title:         "Entry Status",
			Rules:         "",
			Type:          `ExternalEntryStatus1Code (based on string) minLength: 1 maxLength: 4`,
			Documentation: `Status of an entry on the books of the account servicer.`,
		},
		BankTransactionCode: common.ElementHelper{
			Title:         "Bank Transaction Code",
			Rules:         "",
			Type:          `TransactionStatusCode`,
			Documentation: `Proprietary bank transaction code to identify the underlying transaction.`,
		},
		AdditionalInformationIndicator: common.ElementHelper{
			Title:         "Additional Information Indicator",
			Rules:         "",
			Type:          `MessageNameIdentification_FRS_1 (based on string) exactLength: 15 pattern: [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}`,
			Documentation: `Specifies the message name identifier of the message that will be used to provide additional details.`,
		},
		TransactionDetailRefer: common.BuildTransactionDetailReferenceHelper(),
		RelatedAgent:           common.BuildRelatedAgentsHelper(),
		LocalInstrument: common.ElementHelper{
			Title:         "Local Instrument",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `User community specific instrument.`,
		},
		RelatedDate: common.BuildRelatedDatesHelper(),
	}

}
