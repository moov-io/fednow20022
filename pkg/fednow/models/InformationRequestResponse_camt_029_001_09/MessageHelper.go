package InformationRequestResponse_camt_029_001_09

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type MessageHelper struct {
	AssignmentId                 common.ElementHelper
	CreatedDateTime              common.ElementHelper
	Assignments                  common.AssignmentsHelper
	CaseId                       common.ElementHelper
	Creator                      common.CreatorHelper
	Status                       common.ElementHelper
	Interbank                    common.PaymentInfomationHelper
	ResolutionRelatedInformation common.TransactionDetailReferenceHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		AssignmentId: common.ElementHelper{
			Title:         "Assignment ID",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Uniquely identifies the case assignment.`,
		},
		CreatedDateTime: common.ElementHelper{
			Title:         "Created Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `The date and time when the message was created.`,
		},
		Assignments: common.BuildAssignmentsHelper(),
		CaseId: common.ElementHelper{
			Title:         "Case ID",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Uniquely identifies the case.`,
		},
		Creator: common.BuildCreatorHelper(),
		Status: common.ElementHelper{
			Title:         "Status",
			Rules:         "",
			Type:          `InvestigationExecutionConfirmationCode (based on string)`,
			Documentation: `Status of the investigation execution confirmation.`,
		},
		Interbank:                    common.BuildPaymentInfomationHelper(),
		ResolutionRelatedInformation: common.BuildTransactionDetailReferenceHelper(),
	}
}
