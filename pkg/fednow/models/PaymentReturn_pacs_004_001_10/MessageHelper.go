package PaymentReturn_pacs_004_001_10

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type MessageHelper struct {
	MessageId                         common.ElementHelper
	CreatedDateTime                   common.ElementHelper
	NumberOfTransactions              common.ElementHelper
	SettlementInfo                    common.SettlementInformationHelper
	OriginalGroupInfo                 common.GroupInformationHelper
	OriginalTransactionInfo           common.TransactionDetailReferenceHelper
	OriginalInterbankSettlementAmount common.CurrencyAndAmountHelper
	OriginalInterbankSettlementDate   common.ElementHelper
	ReturnedInterbankSettlementAmount common.CurrencyAndAmountHelper
	InterbankSettlementDate           common.ElementHelper
	ChargeBearer                      common.ElementHelper
	Agents                            common.RelatedAgentsHelper
	Debtor                            common.TransactionPartyHelper
	Creditor                          common.TransactionPartyHelper
	ReturnReasonInformation           common.ReturnReasonHelper
	PaymentType                       common.PaymentTypeInfoHelper
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
		NumberOfTransactions: common.ElementHelper{
			Title:         "Number of Transactions",
			Rules:         "This is the total number of transactions in the message.",
			Type:          `Max15NumericText (based on string) minLength: 1 maxLength: 15`,
			Documentation: `Total number of transactions in the message.`,
		},
		SettlementInfo:                    common.BuildSettlementInformationHelper(),
		OriginalGroupInfo:                 common.BuildGroupInformationHelper(),
		OriginalTransactionInfo:           common.BuildTransactionDetailReferenceHelper(),
		OriginalInterbankSettlementAmount: common.BuildCurrencyAndAmountHelper(),
		OriginalInterbankSettlementDate: common.ElementHelper{
			Title:         "Original Interbank Settlement Date",
			Rules:         "This is the date when the original interbank settlement occurred.",
			Type:          `ISODate (based on date)`,
			Documentation: `Date when the original interbank settlement occurred.`,
		},
		ReturnedInterbankSettlementAmount: common.BuildCurrencyAndAmountHelper(),
		InterbankSettlementDate: common.ElementHelper{
			Title:         "Interbank Settlement Date",
			Rules:         "This is the date when the interbank settlement occurs.",
			Type:          `ISODate (based on date)`,
			Documentation: `Date when the interbank settlement occurs.`,
		},
		ChargeBearer: common.ElementHelper{
			Title:         "Charge Bearer",
			Rules:         "This indicates who bears the charges for the transaction.",
			Type:          `ChargeBearerType`,
			Documentation: `Specifies who bears the charges for the transaction, such as debtor, creditor, or shared.`,
		},
		Agents:                  common.BuildRelatedAgentsHelper(),
		Debtor:                  common.BuildTransactionPartyHelper(),
		Creditor:                common.BuildTransactionPartyHelper(),
		ReturnReasonInformation: common.BuildReturnReasonHelper(),
		PaymentType:             common.BuildPaymentTypeInfoHelper(),
	}

}
