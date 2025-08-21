package fednow_test

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/fednow20022/pkg/fednow"

	"github.com/stretchr/testify/require"
)

func TestISODateFormat(t *testing.T) {
	when := time.Date(2019, time.March, 21, 0, 0, 0, 0, time.UTC)

	out, err := fednow.ISODate(when).MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2019-03-21", string(out))

	bs, err := xml.Marshal(fednow.ISODate(when))
	require.NoError(t, err)
	require.Equal(t, "<ISODate>2019-03-21</ISODate>", string(bs))

	var read fednow.ISODate
	err = xml.Unmarshal([]byte("<ISODate>2019-03-21</ISODate>"), &read)
	require.NoError(t, err)
}

// func TestISODateTimeFormat(t *testing.T) {
// 	loc, _ := time.LoadLocation("America/New_York")
// 	when := time.Date(2019, time.March, 21, 10, 36, 19, 0, loc)
//
// 	require.Equal(t, fednow.ISODateTime(when), fednow.UnmarshalISODateTime("2019-03-21T10:36:19"))
//
// 	out, err := fednow.ISODateTime(when).MarshalText()
// 	require.NoError(t, err)
// 	require.Equal(t, "2019-03-21T10:36:19", string(out))
//
// 	out, err = xml.Marshal(fednow.ISODateTime(when))
// 	require.NoError(t, err)
// 	require.Equal(t, "<ISODateTime>2019-03-21T10:36:19</ISODateTime>", string(out))
//
// 	var read fednow.ISODateTime
// 	err = xml.Unmarshal([]byte("<ISODateTime>2019-03-21T10:36:19</ISODateTime>"), &read)
// 	require.NoError(t, err)
// 	require.True(t, when.Equal(time.Time(read)))
// }
