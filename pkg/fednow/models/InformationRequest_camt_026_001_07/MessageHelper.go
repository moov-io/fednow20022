package InformationRequest_camt_026_001_07

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

type JustificationHelper struct {
	MissingInformation           common.MissingCodeAndInfoHelper
	IncorrectInformation         common.MissingCodeAndInfoHelper
	PossibleDuplicateInstruction common.ElementHelper
}

func BuildJustificationHelper() JustificationHelper {
	return JustificationHelper{
		MissingInformation:   common.BuildMissingCodeAndInfoHelper(),
		IncorrectInformation: common.BuildMissingCodeAndInfoHelper(),
		PossibleDuplicateInstruction: common.ElementHelper{
			Title:         "Possible Duplicate Instruction",
			Rules:         "",
			Type:          `boolean`,
			Documentation: `Indicates whether or not the referred item is a possible duplicate of a previous instruction or entry.`,
		},
	}
}

type MessageHelper struct {
	AssignmentId    common.ElementHelper
	CreatedDateTime common.ElementHelper
	Assignments     common.AssignmentsHelper
	CaseId          common.ElementHelper
	Creator         common.CreatorHelper
	Initiation      common.PaymentInfomationHelper
	Interbank       common.PaymentInfomationHelper
	Justification   JustificationHelper
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
		Creator:       common.BuildCreatorHelper(),
		Initiation:    common.BuildPaymentInfomationHelper(),
		Interbank:     common.BuildPaymentInfomationHelper(),
		Justification: BuildJustificationHelper(),
	}
}
