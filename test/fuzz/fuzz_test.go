// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fuzz

import (
	"encoding/xml"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/fednow20022/gen/fednow_incoming_external"
	"github.com/moov-io/fednow20022/gen/fednow_outgoing_external"
	"github.com/moov-io/fednow20022/pkg/fednow"
	"github.com/moov-io/fednow20022/pkg/fednowtest"
)

func FuzzFedNowXML(f *testing.F) {
	populateCorpus(f)

	f.Fuzz(func(t *testing.T, contents string) {
		if len(contents) > 1<<20 {
			t.Skip()
		}
		data := []byte(contents)

		var incoming fednow_incoming_external.FedNowIncoming
		if err := xml.Unmarshal(data, &incoming); err == nil {
			_, _ = xml.Marshal(&incoming)
		}

		var outgoing fednow_outgoing_external.FedNowOutgoing
		if err := xml.Unmarshal(data, &outgoing); err == nil {
			_, _ = xml.Marshal(&outgoing)
		}

		// Direction flip helper
		if out, err := fednowtest.FlipMessageDirection(data); err == nil && len(out) > 0 {
			_, _ = fednowtest.FlipMessageDirection(out)
		}
	})
}

func FuzzFedNowTypes(f *testing.F) {
	// Seed datetime / amount edge cases used throughout generated types
	f.Add("2019-03-21")
	f.Add("2019-03-21T10:36:19-04:00")
	f.Add("0")
	f.Add("1234.56")
	f.Add("")
	f.Add("not-a-date")

	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > 256 {
			t.Skip()
		}

		var d fednow.ISODate
		_ = xml.Unmarshal([]byte("<ISODate>"+xmlEscape(s)+"</ISODate>"), &d)
		_, _ = d.MarshalText()

		var dt fednow.ISODateTime
		_ = xml.Unmarshal([]byte("<ISODateTime>"+xmlEscape(s)+"</ISODateTime>"), &dt)
		_, _ = dt.MarshalText()

		var a fednow.Amount
		_ = xml.Unmarshal([]byte("<Amt>"+xmlEscape(s)+"</Amt>"), &a)
	})
}

func xmlEscape(s string) string {
	var b strings.Builder
	if err := xml.EscapeText(&b, []byte(s)); err != nil {
		return ""
	}
	return b.String()
}

func populateCorpus(f *testing.F) {
	f.Helper()

	f.Add("")
	f.Add("<FedNowIncoming></FedNowIncoming>")
	f.Add("<FedNowOutgoing></FedNowOutgoing>")

	roots := []string{
		filepath.Join("..", "..", "testdata"),
		filepath.Join("..", "..", "tests", "testdata"),
	}
	for _, root := range roots {
		_ = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil || info == nil || info.IsDir() {
				return nil
			}
			if strings.HasSuffix(strings.ToLower(path), ".xml") {
				bs, err := os.ReadFile(path)
				if err != nil || len(bs) > 512*1024 {
					return nil
				}
				f.Add(string(bs))
			}
			return nil
		})
	}
}
