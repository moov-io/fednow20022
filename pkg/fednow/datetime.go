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

func (i ISODate) MarshalText() (string, error) {
	return time.Time(i).Truncate(time.Second).Format(defaultDateFormat), nil
}

func (i ISODate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	content, err := i.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement([]byte(content), start)
}

func (i *ISODate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}

	tt, err := time.Parse(defaultDateFormat, content)
	if err != nil {
		return fmt.Errorf("invalid ISODate format: %v", err)
	}

	*i = ISODate(tt)
	return nil
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

func (i ISODateTime) MarshalText() (string, error) {
	return time.Time(i).Truncate(time.Second).Format(defaultDateTimeFormat), nil
}

func (i ISODateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	content, err := i.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement([]byte(content), start)
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
