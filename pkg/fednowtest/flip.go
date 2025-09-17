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
		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowPaymentStatus != nil:
			xfer := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowPaymentStatus
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowPaymentStatus: &fednow_incoming_external.FedNowPaymentStatus{
						AppHdr:   xfer.AppHdr,
						Document: xfer.Document,
					},
				},
			})

		case fednowOutgoingDoc.FedNowOutgoingMessage.FedNowCustomerCreditTransfer != nil:
			xfer := fednowOutgoingDoc.FedNowOutgoingMessage.FedNowCustomerCreditTransfer
			return xml.Marshal(fednow_incoming_external.FedNowIncoming{
				FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
					FedNowCustomerCreditTransfer: &fednow_incoming_external.FedNowCustomerCreditTransfer{
						AppHdr:   xfer.AppHdr,
						Document: xfer.Document,
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
		case fednowIncomingDoc.FedNowIncomingMessage.FedNowPaymentStatus != nil:
			xfer := fednowIncomingDoc.FedNowIncomingMessage.FedNowPaymentStatus
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowPaymentStatus: &fednow_outgoing_external.FedNowPaymentStatus{
						AppHdr:   xfer.AppHdr,
						Document: xfer.Document,
					},
				},
			})

		case fednowIncomingDoc.FedNowIncomingMessage.FedNowCustomerCreditTransfer != nil:
			xfer := fednowIncomingDoc.FedNowIncomingMessage.FedNowCustomerCreditTransfer
			return xml.Marshal(fednow_outgoing_external.FedNowOutgoing{
				FedNowOutgoingMessage: fednow_outgoing_external.FedNowOutgoingMessage{
					FedNowCustomerCreditTransfer: &fednow_outgoing_external.FedNowCustomerCreditTransfer{
						AppHdr:   xfer.AppHdr,
						Document: xfer.Document,
					},
				},
			})
		}
	}

	return input, nil
}
