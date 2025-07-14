package ReturnRequest_camt_056_001_08

// type CancellationReasonInfoHelper struct {
// 	Contact common.PartyContactHelper
// 	Reason  common.ReturnReasonHelper
// }

// func BuildCancellationReasonInfoHelper() CancellationReasonInfoHelper {
// 	return CancellationReasonInfoHelper{
// 		Contact: common.BuildPartyContactHelper(),
// 		Reason:  common.BuildReturnReasonHelper(),
// 	}
// }

// type CreatorHelper struct {
// 	Agent common.AgentHelper
// 	Party common.PartyContactHelper
// }

// func BuildCreatorHelper() CreatorHelper {
// 	return CreatorHelper{
// 		Agent: common.BuildAgentHelper(),
// 		Party: common.BuildPartyContactHelper(),
// 	}
// }

// type MessageHelper struct {
// 	MessageId                         common.ElementHelper
// 	CreatedDateTime                   common.ElementHelper
// 	Assignment                        common.AssignmentsHelper
// 	CaseId                            common.ElementHelper
// 	CaseCreator                       CreatorHelper
// 	OriginalGroup                     common.GroupInformationHelper
// 	OriginalTransaction               common.TransactionDetailReferenceHelper
// 	OriginalInterbankSettlementAmount common.CurrencyAndAmountHelper
// 	OriginalInterbankSettlementDate   common.ElementHelper
// 	CancelReason                      CancellationReasonInfoHelper
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
// 		Assignment: common.BuildAssignmentsHelper(),
// 		CaseId: common.ElementHelper{
// 			Title:         "Case ID",
// 			Rules:         "This is the unique identifier for the case.",
// 			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
// 			Documentation: `Unique identifier for the case.`,
// 		},
// 		CaseCreator:                       BuildCreatorHelper(),
// 		OriginalGroup:                     common.BuildGroupInformationHelper(),
// 		OriginalTransaction:               common.BuildTransactionDetailReferenceHelper(),
// 		OriginalInterbankSettlementAmount: common.BuildCurrencyAndAmountHelper(),
// 		OriginalInterbankSettlementDate: common.ElementHelper{
// 			Title:         "Original Interbank Settlement Date",
// 			Rules:         "This is the date when the original interbank settlement occurred.",
// 			Type:          `ISODate (based on date)`,
// 			Documentation: `Date when the original interbank settlement occurred.`,
// 		},
// 		CancelReason: BuildCancellationReasonInfoHelper(),
// 	}
// }
