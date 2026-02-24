package fednowtest

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/moov-io/fednow20022/gen/fednow_incoming_external"
	"github.com/moov-io/fednow20022/gen/fednow_outgoing_external"
	"github.com/stretchr/testify/require"
)

func TestFlipMessageContents(t *testing.T) {
	cases := []struct {
		messageFilepath string
		containsBefore  []string
		containsAfter   []string
	}{
		{
			messageFilepath: filepath.Join("..", "..", "testdata", "sample-participant-to-fednow-message.xml"),
			containsBefore:  []string{"FedNowIncoming", "FedNowCustomerCreditTransfer"},
			containsAfter:   []string{"FedNowOutgoing", "FedNowCustomerCreditTransfer"},
		},
		{
			messageFilepath: filepath.Join("..", "..", "testdata", "fednow-sample-message-outgoing-pacs002.xml"),
			containsBefore:  []string{"FedNowOutgoing", "FedNowPaymentStatus"},
			containsAfter:   []string{"FedNowIncoming", "FedNowPaymentStatus"},
		},
		{
			messageFilepath: filepath.Join("..", "..", "testdata", "fednow-sample-message-pacs002.xml"),
			containsBefore:  []string{"FedNowIncoming", "FedNowPaymentStatus"},
			containsAfter:   []string{"FedNowOutgoing", "FedNowPaymentStatus"},
		},
		{
			messageFilepath: filepath.Join("..", "..", "testdata", "admi004-ping.xml"),
			containsBefore:  []string{"FedNowIncoming", "FedNowParticipantBroadcast"},
			containsAfter:   []string{"FedNowOutgoing", "FedNowBroadcast"},
		},
	}

	for _, tc := range cases {
		t.Run(filepath.Base(tc.messageFilepath), func(t *testing.T) {
			bs, err := os.ReadFile(tc.messageFilepath)
			require.NoError(t, err)

			// Verify before
			for _, needle := range tc.containsBefore {
				require.Contains(t, string(bs), needle)
			}

			// flip direction and verify
			flipped, err := FlipMessageDirection(bs)
			require.NoError(t, err)

			for _, needle := range tc.containsAfter {
				require.Contains(t, string(flipped), needle)
			}

			// flip back and verify before
			flipped, err = FlipMessageDirection(flipped)
			require.NoError(t, err)
			require.NotEmpty(t, flipped)

			// read the final as a message
			var fednowOutgoingDoc fednow_outgoing_external.FedNowOutgoing
			outgoingErr := xml.Unmarshal(flipped, &fednowOutgoingDoc)

			var fednowIncomingDoc fednow_incoming_external.FedNowIncoming
			incomingErr := xml.Unmarshal(flipped, &fednowIncomingDoc)

			if outgoingErr != nil && incomingErr != nil {
				t.Fatal(errors.Join(
					fmt.Errorf("outgoing: %w", outgoingErr),
					fmt.Errorf("incoming: %w", incomingErr),
				))
			}

			if outgoingErr != nil {
				require.True(t, verifyUnmarshaled(fednowIncomingDoc))
			}
			if incomingErr != nil {
				require.True(t, verifyUnmarshaled(fednowOutgoingDoc))
			}
		})
	}
}

// verifyUnmarshaled checks that the unmarshaled document has at least one non-nil member
func verifyUnmarshaled(doc interface{}) bool {
	v := reflect.ValueOf(doc)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return false
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.IsValid() && !field.IsZero() {
			return true
		}
	}
	return false
}
