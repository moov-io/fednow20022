package fednow

import (
	"encoding/xml"
	"fmt"
	"time"
)

// ISODate is a custom type wrapping time.Time
type ISODate time.Time

func (i ISODate) MarshalText() (string, error) {
	return time.Time(i).Truncate(time.Second).Format("2006-01-02"), nil
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

	// Parse the date in YYYY-MM-DD format
	tt, err := time.Parse("2006-01-02", content)
	if err != nil {
		return fmt.Errorf("invalid ISODate format: %v", err)
	}

	*i = ISODate(tt)
	return nil
}

type ISODateTime time.Time

func (i ISODateTime) MarshalText() (string, error) {
	return time.Time(i).Truncate(time.Second).Format("2006-01-02T15:04:05-0700"), nil
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

	// Parse the date in YYYY-MM-DD format
	tt, err := time.Parse("2006-01-02T15:04:05-0700", content)
	if err != nil {
		return fmt.Errorf("invalid ISODate format: %v", err)
	}

	*i = ISODateTime(tt)
	return nil
}
