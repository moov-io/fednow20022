package fednow_test

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/moov-io/fednow20022/gen/fednow_incoming_external"
	"github.com/moov-io/fednow20022/gen/head_001_001_02"
	"github.com/moov-io/fednow20022/pkg/fednow"

	"github.com/stretchr/testify/require"
)

func TestISODateFormat(t *testing.T) {
	when := time.Date(2019, time.March, 21, 0, 0, 0, 123, time.UTC)
	fmt.Println(when.Format(time.RFC3339))

	out, err := fednow.ISODate(when).MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2019-03-21", string(out))

	fmt.Println(out)

	bs, err := xml.Marshal(fednow.ISODate(when))
	require.NoError(t, err)
	require.Equal(t, "<ISODate>2019-03-21</ISODate>", string(bs))

	var read fednow.ISODate
	err = xml.Unmarshal([]byte("<ISODate>2019-03-21</ISODate>"), &read)
	require.NoError(t, err)
}

func TestISODateTimeFormat(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	when := time.Date(2019, time.March, 21, 10, 36, 19, 0, loc)

	out, err := fednow.ISODateTime(when).MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2019-03-21T10:36:19-04:00", string(out))

	bs, err := xml.Marshal(fednow.ISODateTime(when))
	require.NoError(t, err)
	require.Equal(t, "<ISODateTime>2019-03-21T10:36:19-04:00</ISODateTime>", string(bs))

	var read fednow.ISODateTime
	err = xml.Unmarshal([]byte("<ISODateTime>2019-03-21T10:36:19-04:00</ISODateTime>"), &read)
	require.NoError(t, err)

	fmt.Println(when.Format(time.RFC3339))
	fmt.Println(time.Time(read).Format(time.RFC3339))

	require.True(t, when.Equal(time.Time(read)))
}

func TestDateTimeXML(t *testing.T) {
	mountain, err := time.LoadLocation("America/Denver")
	require.NoError(t, err)

	when := time.Date(2025, time.September, 10, 12, 20, 0, 0, mountain)

	t.Run("head_001_001_02.AppHdr", func(t *testing.T) {
		write := fednow_incoming_external.FedNowIncoming{
			FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
				FedNowCustomerCreditTransfer: &fednow_incoming_external.FedNowCustomerCreditTransfer{
					AppHdr: head_001_001_02.BusinessApplicationHeaderV02{
						CreDt: fednow.ISODateTime(when),
					},
				},
			},
		}

		bs, err := xml.Marshal(write)
		require.NoError(t, err)

		require.Contains(t, string(bs), `<CreDt>2025-09-10T12:20:00-06:00</CreDt>`)

		var read fednow_incoming_external.FedNowIncoming
		err = xml.Unmarshal(bs, &read)
		require.NoError(t, err)

		xfer := read.FedNowIncomingMessage.FedNowCustomerCreditTransfer
		require.NotNil(t, xfer)

		require.Equal(t, "2025-09-10T12:20:00-06:00", time.Time(xfer.AppHdr.CreDt).Format(time.RFC3339))
		require.Nil(t, xfer.AppHdr.BizPrcgDt)
	})

	t.Run("empty", func(t *testing.T) {
		bs := []byte(`<FedNowIncoming><FedNowIncomingMessage><FedNowCustomerCreditTransfer><AppHdr><Fr></Fr><To></To><BizMsgIdr></BizMsgIdr><MsgDefIdr></MsgDefIdr><CreDt></CreDt></AppHdr><Document><FIToFICstmrCdtTrf><GrpHdr><MsgId></MsgId><CreDtTm>0001-01-01T00:00:00+00:00</CreDtTm><NbOfTxs></NbOfTxs><SttlmInf><SttlmMtd></SttlmMtd></SttlmInf></GrpHdr></FIToFICstmrCdtTrf></Document></FedNowCustomerCreditTransfer></FedNowIncomingMessage></FedNowIncoming>`)

		var read fednow_incoming_external.FedNowIncoming
		err = xml.Unmarshal(bs, &read)
		require.ErrorContains(t, err, `CreDt (namespace=) has no date time format found for ""`)
	})
}
