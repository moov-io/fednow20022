package fednow

import (
	"encoding/xml"
	"fmt"
	"time"
)

// ISODate is a custom type wrapping time.Time
type ISODate time.Time

var (
	defaultDateFormat = "2006-01-02"
)

// MarshalText implements encoding.TextMarshaler. It writes the date as
// YYYY-MM-DD.
func (i ISODate) MarshalText() ([]byte, error) {
	return []byte(time.Time(i).Truncate(time.Second).Format(defaultDateFormat)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler. It reads a date as
// YYYY-MM-DD.
func (i *ISODate) UnmarshalText(text []byte) error {
	tt, err := time.Parse(defaultDateFormat, string(text))
	if err != nil {
		return fmt.Errorf("invalid ISODate format: %v", err)
	}

	*i = ISODate(tt)
	return nil
}

func (i ISODate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	content, err := i.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(content, start)
}

func (i *ISODate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}
	return i.UnmarshalText([]byte(content))
}

var (
	defaultDateTimeFormat = "2006-01-02T15:04:05-07:00"

	acceptedDateTimeFormats = []string{
		defaultDateTimeFormat,
		"2006-01-02T15:04:05-0700",
		"2006-01-02T15:04:05.000",
		"2006-01-02T15:04:05.00000",
		"2006-01-02T15:04:05.000Z",
		"2006-01-02T15:04:05.00000Z",
		"2006-01-02T15:04:05.000-07:00",
		"2006-01-02T15:04:05.00000-07:00",
	}
)

type ISODateTime time.Time

// MarshalText implements encoding.TextMarshaler. It writes the date time
// as YYYY-MM-DDThh:mm:ss with the UTC offset, for example
// 2019-03-21T10:36:19-04:00.
func (i ISODateTime) MarshalText() ([]byte, error) {
	return []byte(time.Time(i).Truncate(time.Second).Format(defaultDateTimeFormat)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler. It accepts the same
// formats as UnmarshalXML.
func (i *ISODateTime) UnmarshalText(text []byte) error {
	content := string(text)
	for idx := range acceptedDateTimeFormats {
		tt, err := time.Parse(acceptedDateTimeFormats[idx], content)
		if err == nil {
			*i = ISODateTime(tt)
			return nil
		}
	}
	return fmt.Errorf("ISODateTime has no date time format found for %q", content)
}

func (i ISODateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	content, err := i.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(content, start)
}

func (i *ISODateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}

	for idx := range acceptedDateTimeFormats {
		tt, err := time.Parse(acceptedDateTimeFormats[idx], content)
		if err == nil {
			*i = ISODateTime(tt)
			return nil
		}
	}

	return fmt.Errorf("%s (namespace=%s) has no date time format found for %q", start.Name.Local, start.Name.Space, content)
}
