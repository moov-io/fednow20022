package fednowtest

import (
	"encoding/xml"

	"github.com/moov-io/fednow20022/gen/fednow_incoming_external"
	"github.com/moov-io/fednow20022/gen/fednow_outgoing_external"
)

// FlipMessageDirection will attempt to switch XML messages from incoming to outgoing and vice-versa.
// This is needed so we can receive out own originated messages (e.g. pacs.008)
//
// Right now only a subset of possible FedNow messages are supported
func FlipMessageDirection(input []byte) ([]byte, error) {
	var fednowOutgoingDoc fednow_outgoing_external.FedNowOutgoing
	fednowError := xml.Unmarshal(input, &fednowOutgoingDoc)
	if fednowError == nil {
		// Convert the message into a FedNowIncomingMessage
		switch {
		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowMessageReject != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowMessageReject
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowMessageReject: &fednow_incoming_external.FedNowMessageReject{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowBroadcast != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowBroadcast
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowParticipantBroadcast: &fednow_incoming_external.FedNowParticipantBroadcast{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowReceiptAcknowledgement != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowReceiptAcknowledgement
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowReceiptAcknowledgement: &fednow_incoming_external.FedNowReceiptAcknowledgement{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowInformationRequest != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowInformationRequest
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowInformationRequest: &fednow_incoming_external.FedNowInformationRequest{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowAdditionalPaymentInformation != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowAdditionalPaymentInformation
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowAdditionalPaymentInformation: &fednow_incoming_external.FedNowAdditionalPaymentInformation{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowInformationRequestResponse != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowInformationRequestResponse
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowInformationRequestResponse: &fednow_incoming_external.FedNowInformationRequestResponse{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowRequestForPaymentCancellationRequestResponse != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowRequestForPaymentCancellationRequestResponse
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowRequestForPaymentCancellationRequestResponse: &fednow_incoming_external.FedNowRequestForPaymentCancellationRequestResponse{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowReturnRequestResponse != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowReturnRequestResponse
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowReturnRequestResponse: &fednow_incoming_external.FedNowReturnRequestResponse{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowRequestForPaymentCancellationRequest != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowRequestForPaymentCancellationRequest
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowRequestForPaymentCancellationRequest: &fednow_incoming_external.FedNowRequestForPaymentCancellationRequest{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowReturnRequest != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowReturnRequest
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowReturnRequest: &fednow_incoming_external.FedNowReturnRequest{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowPaymentStatus != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowPaymentStatus
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowPaymentStatus: &fednow_incoming_external.FedNowPaymentStatus{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowPaymentReturn != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowPaymentReturn
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowPaymentReturn: &fednow_incoming_external.FedNowPaymentReturn{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowCustomerCreditTransfer != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowCustomerCreditTransfer
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowCustomerCreditTransfer: &fednow_incoming_external.FedNowCustomerCreditTransfer{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowInstitutionCreditTransfer != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowInstitutionCreditTransfer
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowInstitutionCreditTransfer: &fednow_incoming_external.FedNowInstitutionCreditTransfer{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowPaymentStatusRequest != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowPaymentStatusRequest
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowPaymentStatusRequest: &fednow_incoming_external.FedNowPaymentStatusRequest{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowRequestForPayment != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowRequestForPayment
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowRequestForPayment: &fednow_incoming_external.FedNowRequestForPayment{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowRequestForPaymentResponse != nil:
			msg := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowRequestForPaymentResponse
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowRequestForPaymentResponse: &fednow_incoming_external.FedNowRequestForPaymentResponse{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		}
	}

	var fednowIncomingDoc fednow_incoming_external.FedNowIncoming
	fednowError2 := xml.Unmarshal(input, &fednowIncomingDoc)
	if fednowError2 == nil {
		// Convert the message into a FedNowOutgoing
		switch {
		case fednowIncomingDoc.FedNowIncomingMessage.FedNowMessageReject != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowMessageReject
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowMessageReject: &fednow_outgoing_external.FedNowMessageReject{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowParticipantBroadcast != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowParticipantBroadcast
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowBroadcast: &fednow_outgoing_external.FedNowBroadcast{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowReceiptAcknowledgement != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowReceiptAcknowledgement
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowReceiptAcknowledgement: &fednow_outgoing_external.FedNowReceiptAcknowledgement{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowInformationRequest != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowInformationRequest
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowInformationRequest: &fednow_outgoing_external.FedNowInformationRequest{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowAdditionalPaymentInformation != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowAdditionalPaymentInformation
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowAdditionalPaymentInformation: &fednow_outgoing_external.FedNowAdditionalPaymentInformation{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowInformationRequestResponse != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowInformationRequestResponse
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowInformationRequestResponse: &fednow_outgoing_external.FedNowInformationRequestResponse{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowRequestForPaymentCancellationRequestResponse != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowRequestForPaymentCancellationRequestResponse
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowRequestForPaymentCancellationRequestResponse: &fednow_outgoing_external.FedNowRequestForPaymentCancellationRequestResponse{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowReturnRequestResponse != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowReturnRequestResponse
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowReturnRequestResponse: &fednow_outgoing_external.FedNowReturnRequestResponse{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowRequestForPaymentCancellationRequest != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowRequestForPaymentCancellationRequest
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowRequestForPaymentCancellationRequest: &fednow_outgoing_external.FedNowRequestForPaymentCancellationRequest{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowReturnRequest != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowReturnRequest
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowReturnRequest: &fednow_outgoing_external.FedNowReturnRequest{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowPaymentStatus != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowPaymentStatus
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowPaymentStatus: &fednow_outgoing_external.FedNowPaymentStatus{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowPaymentReturn != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowPaymentReturn
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowPaymentReturn: &fednow_outgoing_external.FedNowPaymentReturn{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowCustomerCreditTransfer != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowCustomerCreditTransfer
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowCustomerCreditTransfer: &fednow_outgoing_external.FedNowCustomerCreditTransfer{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowInstitutionCreditTransfer != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowInstitutionCreditTransfer
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowInstitutionCreditTransfer: &fednow_outgoing_external.FedNowInstitutionCreditTransfer{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowPaymentStatusRequest != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowPaymentStatusRequest
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowPaymentStatusRequest: &fednow_outgoing_external.FedNowPaymentStatusRequest{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowRequestForPayment != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowRequestForPayment
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowRequestForPayment: &fednow_outgoing_external.FedNowRequestForPayment{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowRequestForPaymentResponse != nil:
			msg := fednowIncomingDoc.FedNowIncomingMessage.FedNowRequestForPaymentResponse
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowRequestForPaymentResponse: &fednow_outgoing_external.FedNowRequestForPaymentResponse{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		}
	}

	return input, nil
}
