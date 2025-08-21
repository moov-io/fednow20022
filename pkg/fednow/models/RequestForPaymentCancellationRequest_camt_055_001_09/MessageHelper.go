package RequestForPaymentCancellationRequest_camt_055_001_09

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type MessageHelper struct {
	MessageId                      common.ElementHelper
	CreatedDateTime                common.ElementHelper
	Assigners                      common.AssignmentsHelper
	CaseId                         common.ElementHelper
	Creator                        common.PartyContactHelper
	PaymentCancellationId          common.ElementHelper
	GroupInfo                      common.GroupInformationHelper
	TransactionInfo                common.TransactionDetailReferenceHelper
	OriginalInstructedAmount       common.CurrencyAndAmountHelper
	OriginalRequestedExecutionDate common.ElementHelper
	Reason                         common.ReturnReasonHelper
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
		Assigners: common.BuildAssignmentsHelper(),
		CaseId: common.ElementHelper{
			Title:         "Case ID",
			Rules:         "This is the unique identifier for the case.",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identifier for the case.`,
		},
		Creator: common.BuildPartyContactHelper(),
		PaymentCancellationId: common.ElementHelper{
			Title:         "Payment Cancellation ID",
			Rules:         "This is the unique identifier for the payment cancellation.",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identifier for the payment cancellation.`,
		},
		GroupInfo:                common.BuildGroupInformationHelper(),
		TransactionInfo:          common.BuildTransactionDetailReferenceHelper(),
		OriginalInstructedAmount: common.BuildCurrencyAndAmountHelper(),
		OriginalRequestedExecutionDate: common.ElementHelper{
			Title:         "Original Requested Execution Date",
			Rules:         "This is the original requested execution date.",
			Type:          `ISODate (based on date)`,
			Documentation: `Original requested execution date.`,
		},
		Reason: common.BuildReturnReasonHelper(),
	}
}
