package fednowtest

import (
	"os"
	"path/filepath"
	"testing"

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
	}

	for _, tc := range cases {
		t.Run(filepath.Base(tc.messageFilepath), func(t *testing.T) {
			bs, err := os.ReadFile(tc.messageFilepath)
			require.NoError(t, err)

			for _, needle := range tc.containsBefore {
				require.Contains(t, string(bs), needle)
			}

			flipped, err := FlipMessageDirection(bs)
			require.NoError(t, err)

			for _, needle := range tc.containsAfter {
				require.Contains(t, string(flipped), needle)
			}
		})
	}
}
