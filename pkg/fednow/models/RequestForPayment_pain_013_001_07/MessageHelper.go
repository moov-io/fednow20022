package RequestForPayment_pain_013_001_07

// type PaymentConditionHelper struct {
// 	AmountModificationAllowed  common.ElementHelper
// 	EarlyPaymentAllowed        common.ElementHelper
// 	GuaranteedPaymentRequested common.ElementHelper
// }

// func BuildPaymentConditionHelper() PaymentConditionHelper {
// 	return PaymentConditionHelper{
// 		AmountModificationAllowed: common.ElementHelper{
// 			Title:         "Amount Modification Allowed",
// 			Rules:         "Indicates if the amount can be modified.",
// 			Type:          `Boolean`,
// 			Documentation: `Indicates whether the amount can be modified.`,
// 		},
// 		EarlyPaymentAllowed: common.ElementHelper{
// 			Title:         "Early Payment Allowed",
// 			Rules:         "Indicates if early payment is allowed.",
// 			Type:          `Boolean`,
// 			Documentation: `Indicates whether early payment is allowed.`,
// 		},
// 		GuaranteedPaymentRequested: common.ElementHelper{
// 			Title:         "Guaranteed Payment Requested",
// 			Rules:         "Indicates if a guaranteed payment is requested.",
// 			Type:          `Boolean`,
// 			Documentation: `Indicates whether a guaranteed payment is requested.`,
// 		},
// 	}
// }

// type MessageHelper struct {
// 	MessageId                           common.ElementHelper
// 	CreatedDateTime                     common.ElementHelper
// 	NumberOfTransactions                common.ElementHelper
// 	InitiatingPartyName                 common.ElementHelper
// 	PaymentInformationId                common.ElementHelper
// 	PaymentMethod                       common.ElementHelper
// 	PaymentCondition                    PaymentConditionHelper
// 	RequestedExecutionDate              common.ElementHelper
// 	ExpiryDate                          common.ElementHelper
// 	Debtor                              common.TransactionPartyHelper
// 	CreditTransferId                    common.TransactionDetailReferenceHelper
// 	CreditTransferPaymentType           common.ElementHelper
// 	CreditTransferAmount                common.CurrencyAndAmountHelper
// 	CreditTransferChargeBearer          common.ElementHelper
// 	CreditTransferCreditor              common.TransactionPartyHelper
// 	CreditTransferPurposeCode           common.ElementHelper
// 	CreditTransferRemitInfo             common.RemittanceInformationHelper
// 	CreditTransferRelatedRemittanceInfo common.RelatedRemittanceInformationHelper
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
// 		NumberOfTransactions: common.ElementHelper{
// 			Title:         "Number of Transactions",
// 			Rules:         "This is the number of transactions included in the message.",
// 			Type:          `Max15NumericText (based on string) minLength: 1 maxLength: 15`,
// 			Documentation: `Number of transactions included in the message.`,
// 		},
// 		InitiatingPartyName: common.ElementHelper{
// 			Title:         "Initiating Party Name",
// 			Rules:         "This is the name of the party initiating the payment.",
// 			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
// 			Documentation: `Name of the party initiating the payment.`,
// 		},
// 		PaymentInformationId: common.ElementHelper{
// 			Title:         "Payment Information ID",
// 			Rules:         "This is the identifier for the payment information.",
// 			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
// 			Documentation: `Unique identifier for the payment information.`,
// 		},
// 		PaymentMethod: common.ElementHelper{
// 			Title:         "Payment Method",
// 			Rules:         "This is the method used for the payment.",
// 			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
// 			Documentation: `Method used for the payment.`,
// 		},
// 		PaymentCondition: BuildPaymentConditionHelper(),
// 		RequestedExecutionDate: common.ElementHelper{
// 			Title:         "Requested Execution Date",
// 			Rules:         "This is the date when the payment is requested to be executed.",
// 			Type:          `ISODate (based on date)`,
// 			Documentation: `Date when the payment is requested to be executed.`,
// 		},
// 		ExpiryDate: common.ElementHelper{
// 			Title:         "Expiry Date",
// 			Rules:         "This is the date when the payment request expires.",
// 			Type:          `ISODate (based on date)`,
// 			Documentation: `Date when the payment request expires.`,
// 		},
// 		Debtor:           common.TransactionPartyHelper{},
// 		CreditTransferId: common.TransactionDetailReferenceHelper{},
// 		CreditTransferPaymentType: common.ElementHelper{
// 			Title:         "Credit Transfer Payment Type",
// 			Rules:         "This is the type of payment for the credit transfer.",
// 			Type:          `PaymentTypeInfo`,
// 			Documentation: `Type of payment for the credit transfer.`,
// 		},
// 		CreditTransferAmount: common.CurrencyAndAmountHelper{},
// 		CreditTransferChargeBearer: common.ElementHelper{
// 			Title:         "Credit Transfer Charge Bearer",
// 			Rules:         "This indicates who bears the charges for the credit transfer.",
// 			Type:          `ChargeBearerType`,
// 			Documentation: `Specifies which party/parties will bear the charges associated with the processing of the payment transaction.`,
// 		},
// 		CreditTransferCreditor: common.TransactionPartyHelper{},
// 		CreditTransferPurposeCode: common.ElementHelper{
// 			Title:         "Credit Transfer Purpose Code",
// 			Rules:         "This is the purpose code for the credit transfer.",
// 			Type:          `PurposeCodeType`,
// 			Documentation: `Specifies the purpose of the credit transfer.`,
// 		},
// 		CreditTransferRemitInfo:             common.RemittanceInformationHelper{},
// 		CreditTransferRelatedRemittanceInfo: common.RelatedRemittanceInformationHelper{},
// 	}
// }
