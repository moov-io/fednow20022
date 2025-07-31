package fednow

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMessageID(t *testing.T) {
	cases := []struct {
		name                string
		when                time.Time
		partyIdentification string
		reference           string
		expected            string
	}{
		{
			name:                "short reference",
			when:                time.Date(2025, time.July, 31, 10, 30, 0, 0, time.UTC),
			partyIdentification: "987654321",
			reference:           "AABBCCDDEEFF",
			expected:            "20250731987654321AABBCCDDEEFF",
		},
		{
			name:                "long everything",
			when:                time.Date(2025, time.July, 31, 10, 30, 0, 0, time.UTC),
			partyIdentification: "9876543210000",
			reference:           "AABBCCDDEEFF000000000",
			expected:            "20250731987654321AABBCCDDEEFF000000",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := MessageID(tc.when, tc.partyIdentification, tc.reference)

			require.Equal(t, tc.expected, got)
		})
	}
}
