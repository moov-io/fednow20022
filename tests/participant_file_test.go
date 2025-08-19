package tests

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/fednow20022/gen/fednow_outgoing_external"
	"github.com/moov-io/fednow20022/gen/sup_FedNowParticipantFile_admi_998_001_02"

	"github.com/stretchr/testify/require"
)

func TestParticipantFile(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("..", "testdata", "fednow-sample-message-admi998.xml"))
	require.NoError(t, err)

	var outer fednow_outgoing_external.FedNowOutgoing
	err = xml.Unmarshal(bs, &outer)
	require.NoError(t, err)

	require.Nil(t, outer.FedNowTechnicalHeader)
	require.NotNil(t, outer.FedNowOutgoingMessage.FedNowParticipantFile)

	file := outer.FedNowOutgoingMessage.FedNowParticipantFile

	// App Header
	require.Equal(t, "021150706", string(file.AppHdr.Fr.FIId.FinInstnId.ClrSysMmbId.MmbId))
	require.Equal(t, "111111111", string(file.AppHdr.To.FIId.FinInstnId.ClrSysMmbId.MmbId))
	require.Equal(t, "20230818FedNowReferencePartFile", string(file.AppHdr.BizMsgIdr))

	// Document
	require.Equal(t, "FPFL", string(file.Document.AdmstnPrtryMsg.PrtryData.Tp))

	// Find the inner XML to read out
	input := string(bs)

	start := "<Data>"
	startIdx := strings.Index(input, start) + len(start)

	end := "</Data>"
	endIdx := strings.Index(input, end) - len(end)

	innerXML := strings.TrimSpace(input[startIdx:endIdx])
	require.Contains(t, innerXML, "<Document xmlns=")

	// Read the inner XML
	var inner sup_FedNowParticipantFile_admi_998_001_02.Document
	err = xml.Unmarshal([]byte(innerXML), &inner)
	require.NoError(t, err)

	bizDay := inner.Admi998SuplDataV01.PtcptFile.BizDay
	require.Equal(t, "2023-08-18", time.Time(bizDay).Format("2006-01-02"))

	profiles := inner.Admi998SuplDataV01.PtcptFile.PtcptPrfl
	require.Len(t, profiles, 8)

	require.Equal(t, "222222222", string(profiles[0].Id))
	require.Equal(t, "Bank A", string(profiles[0].Nm))
	require.Equal(t, []string{"CTRO", "CTSR", "RFPR"}, toString(profiles[0].Svcs))

	require.Equal(t, "333333333", string(profiles[1].Id))
	require.Equal(t, "Bank B", string(profiles[1].Nm))
	require.Equal(t, []string{"CTRO", "CTSR", "RFPR"}, toString(profiles[1].Svcs))

	require.Equal(t, "444444444", string(profiles[2].Id))
	require.Equal(t, "Bank C", string(profiles[2].Nm))
	require.Equal(t, []string{"CTRO", "CTSR", "RFPR"}, toString(profiles[2].Svcs))

	require.Equal(t, "555555555", string(profiles[3].Id))
	require.Equal(t, "Bank D", string(profiles[3].Nm))
	require.Equal(t, []string{"CTRO", "CTSR", "RFPR"}, toString(profiles[3].Svcs))

	require.Equal(t, "666666666", string(profiles[4].Id))
	require.Equal(t, "Faster Payment System FPS", string(profiles[4].Nm))
	require.Equal(t, []string{"CTRO", "CTSR"}, toString(profiles[4].Svcs))

	require.Equal(t, "777777777", string(profiles[5].Id))
	require.Equal(t, "Bank E", string(profiles[5].Nm))
	require.Equal(t, []string{"CTRO", "CTSR", "RFPR"}, toString(profiles[5].Svcs))

	require.Equal(t, "888888888", string(profiles[6].Id))
	require.Equal(t, "Bank F", string(profiles[6].Nm))
	require.Equal(t, []string{"CTRO", "CTSR", "RFPR"}, toString(profiles[6].Svcs))

	require.Equal(t, "999999999", string(profiles[7].Id))
	require.Equal(t, "Bank G", string(profiles[7].Nm))
	require.Equal(t, []string{"CTRO", "CTSR", "RFPR"}, toString(profiles[7].Svcs))
}

func toString[T ~string](input []T) []string {
	out := make([]string, len(input))
	for idx := range input {
		out[idx] = string(input[idx])
	}
	return out
}
