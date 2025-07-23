package RequestForPaymentResponse_pain_014_001_07

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type PaymentConditionStatusHelper struct {
	AcceptedAmount    common.CurrencyAndAmountHelper
	GuaranteedPayment common.ElementHelper
	EarlyPayment      common.ElementHelper
}

func BuildPaymentConditionStatusHelper() PaymentConditionStatusHelper {
	return PaymentConditionStatusHelper{
		AcceptedAmount: common.BuildCurrencyAndAmountHelper(),
		GuaranteedPayment: common.ElementHelper{
			Title:         "Guaranteed Payment",
			Rules:         "This indicates whether the payment is guaranteed.",
			Type:          `boolean`,
			Documentation: `Indicates if the payment is guaranteed.`,
		},
		EarlyPayment: common.ElementHelper{
			Title:         "Early Payment",
			Rules:         "This indicates whether early payment is applicable.",
			Type:          `boolean`,
			Documentation: `Indicates if early payment is applicable.`,
		},
	}
}

type TransactionReferenceHelper struct {
	Amount                 common.CurrencyAndAmountHelper
	RequestedExecutionDate common.ElementHelper
	Creditor               common.TransactionPartyHelper
}

func BuildTransactionReferenceHelper() TransactionReferenceHelper {
	return TransactionReferenceHelper{
		Amount: common.BuildCurrencyAndAmountHelper(),
		RequestedExecutionDate: common.ElementHelper{
			Title:         "Requested Execution Date",
			Rules:         "This is the date when the transaction is requested to be executed.",
			Type:          `ISODate (based on date)`,
			Documentation: `Date when the transaction is requested to be executed.`,
		},
		Creditor: common.BuildTransactionPartyHelper(),
	}
}

type MessageHelper struct {
	MessageId                    common.ElementHelper
	CreatedDateTime              common.ElementHelper
	InitiatingPartyName          common.ElementHelper
	Agents                       common.DebtorAndCreditorAgentHelper
	OriginalGroupInfo            common.GroupInformationHelper
	OriginalPaymentInformationId common.ElementHelper
	OriginalTransactionInfo      TransactionReferenceHelper
	TransactionStatus            common.ElementHelper
	PaymentConditionStatus       PaymentConditionStatusHelper
	ReturnReason                 common.ReturnReasonHelper
	OriginalTransactionReference TransactionReferenceHelper
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
		InitiatingPartyName: common.ElementHelper{
			Title:         "Initiating Party Name",
			Rules:         "This is the name of the party initiating the payment.",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Name of the party initiating the payment.`,
		},
		Agents:            common.BuildDebtorAndCreditorAgentHelper(),
		OriginalGroupInfo: common.BuildGroupInformationHelper(),
		OriginalPaymentInformationId: common.ElementHelper{
			Title:         "Original Payment Information ID",
			Rules:         "This is the ID of the original payment information.",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `ID of the original payment information.`,
		},
		OriginalTransactionInfo: BuildTransactionReferenceHelper(),
		TransactionStatus: common.ElementHelper{
			Title:         "Transaction Status",
			Rules:         "This indicates the status of the transaction.",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Status of the transaction.`,
		},
		PaymentConditionStatus:       BuildPaymentConditionStatusHelper(),
		ReturnReason:                 common.BuildReturnReasonHelper(),
		OriginalTransactionReference: BuildTransactionReferenceHelper(),
	}
}
