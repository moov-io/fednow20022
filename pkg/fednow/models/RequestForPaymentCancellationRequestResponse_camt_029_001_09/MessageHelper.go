package RequestForPaymentCancellationRequestResponse_camt_029_001_09

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type CancellationStatusReasonHelper struct {
	OriginatorName common.ElementHelper
	ReasonCode     common.ElementHelper
}

func BuildCancellationStatusReasonHelper() CancellationStatusReasonHelper {
	return CancellationStatusReasonHelper{
		OriginatorName: common.ElementHelper{
			Title:         "Originator Name",
			Rules:         "This is the name of the originator of the cancellation request.",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Name of the originator of the cancellation request.`,
		},
		ReasonCode: common.ElementHelper{
			Title:         "Reason Code",
			Rules:         "This is the reason code for the cancellation request.",
			Type:          `ReasonCodeType (based on string)`,
			Documentation: `Reason code for the cancellation request.`,
		},
	}
}

type MessageHelper struct {
	MessageId                      common.ElementHelper
	CreatedDateTime                common.ElementHelper
	Assigners                      common.AssignmentsHelper
	ResolvedCaseId                 common.ElementHelper
	Creator                        common.PartyContactHelper
	Status                         common.ElementHelper
	PaymentCancellationId          common.ElementHelper
	GroupInfo                      common.GroupInformationHelper
	TransactionInfo                common.TransactionDetailReferenceHelper
	OriginalInstructedAmount       common.CurrencyAndAmountHelper
	OriginalRequestedExecutionDate common.ElementHelper
	Reason                         CancellationStatusReasonHelper
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
		ResolvedCaseId: common.ElementHelper{
			Title:         "Resolved Case ID",
			Rules:         "This is the unique identifier for the resolved case.",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identifier for the resolved case.`,
		},
		Creator: common.BuildPartyContactHelper(),
		Status: common.ElementHelper{
			Title:         "Status",
			Rules:         "This indicates the status of the payment cancellation request.",
			Type:          `CancellationRequestResponse (based on string)`,
			Documentation: `Status of the payment cancellation request, such as cancelled, rejected, or pending.`,
		},
		PaymentCancellationId: common.ElementHelper{
			Title:         "Payment Cancellation ID",
			Rules:         "This is the unique identifier for the payment cancellation.",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identifier for the payment cancellation.`,
		},
		GroupInfo:       common.BuildGroupInformationHelper(),
		TransactionInfo: common.BuildTransactionDetailReferenceHelper(),

		Reason: BuildCancellationStatusReasonHelper(),
	}
}
