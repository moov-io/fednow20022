package tests

import (
	"bytes"
	"encoding/xml"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func xmlMatches[T any](tb testing.TB, doc T, expectedPath string) {
	tb.Helper()

	marshaled, err := xml.MarshalIndent(doc, "", "  ")
	require.NoError(tb, err)

	read, err := os.ReadFile(expectedPath)
	require.NoError(tb, err)

	require.Equal(tb, string(bytes.TrimSpace(read)), string(bytes.TrimSpace(marshaled)))
}
