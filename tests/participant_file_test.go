package tests

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/fednow20022/gen/FedNow_Service_Release_1__Edition_08_06_2024_LIV_FedNowParticipantFile_admi_998_001_02_20241122_1718"
	"github.com/moov-io/fednow20022/gen/sup_FedNowParticipantFile_admi_998_001_02_20250730_1500"

	"cloud.google.com/go/civil"
	"github.com/stretchr/testify/require"
)

func TestParticipantFile(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("..", "testdata", "FedNowParticipantFile_Step1_admi.998.xml"))
	require.NoError(t, err)

	// When reading a full fednow message it would be using the following with their FedIncomgingMessage wrapper
	// gen/Message_FedNowParticipantFile_admi_998_001_02/model.go

	// For this test we are just verifying that extraction of the Proprietary Data (xml inside of xml) can be fully read
	var outer FedNow_Service_Release_1__Edition_08_06_2024_LIV_FedNowParticipantFile_admi_998_001_02_20241122_1718.Document
	err = xml.Unmarshal(bs, &outer)
	require.NoError(t, err)

	require.Equal(t, "FPFL", string(outer.AdmstnPrtryMsg.PrtryData.Tp))

	innerXML := strings.TrimSpace(outer.AdmstnPrtryMsg.PrtryData.Data.Item)
	require.Contains(t, innerXML, "<Document xmlns=")

	// Read the inner XML
	var inner sup_FedNowParticipantFile_admi_998_001_02_20250730_1500.Document
	err = xml.Unmarshal([]byte(innerXML), &inner)
	require.NoError(t, err)

	bizDay := inner.Admi998SuplDataV01.PtcptFile.BizDay
	require.Equal(t, "2023-08-18", civil.Date(bizDay).String())

	profiles := inner.Admi998SuplDataV01.PtcptFile.PtcptPrfl
	require.Len(t, profiles, 8)

	require.Equal(t, "011104238", string(profiles[0].Id))
	require.Equal(t, "Bank A", string(profiles[0].Nm))
	require.Equal(t, []string{"CTSR", "RFPR"}, toString(profiles[0].Svcs))

	require.Equal(t, "021040078", string(profiles[1].Id))
	require.Equal(t, "Bank B", string(profiles[1].Nm))
	require.Equal(t, []string{"CTSR", "RFPR"}, toString(profiles[1].Svcs))

	require.Equal(t, "021307481", string(profiles[2].Id))
	require.Equal(t, "Bank C", string(profiles[2].Nm))
	require.Equal(t, []string{"CTSR", "RFPR"}, toString(profiles[2].Svcs))

	require.Equal(t, "231981435", string(profiles[3].Id))
	require.Equal(t, "Bank D", string(profiles[3].Nm))
	require.Equal(t, []string{"CTSR", "RFPR"}, toString(profiles[3].Svcs))

	require.Equal(t, "211590752", string(profiles[4].Id))
	require.Equal(t, "Faster Payment System FPS", string(profiles[4].Nm))
	require.Equal(t, []string{"CTSR"}, toString(profiles[4].Svcs))

	require.Equal(t, "122240120", string(profiles[5].Id))
	require.Equal(t, "Bank E", string(profiles[5].Nm))
	require.Equal(t, []string{"CTSR", "RFPR"}, toString(profiles[5].Svcs))

	require.Equal(t, "113185029", string(profiles[6].Id))
	require.Equal(t, "Bank F", string(profiles[6].Nm))
	require.Equal(t, []string{"CTSR", "RFPR"}, toString(profiles[6].Svcs))

	require.Equal(t, "113194159", string(profiles[7].Id))
	require.Equal(t, "Bank G", string(profiles[7].Nm))
	require.Equal(t, []string{"CTRO"}, toString(profiles[7].Svcs))
}

func toString[T ~string](input []T) []string {
	out := make([]string, len(input))
	for idx := range input {
		out[idx] = string(input[idx])
	}
	return out
}
