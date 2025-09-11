//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Define the replacements
	replacements := []struct {
		old string
		new string
	}{
		// imports
		{`"time"`, `"github.com/moov-io/fednow20022/pkg/fednow"`},
		// Date
		{`type ISODate time.Time`, ``},
		{` ISODate `, ` fednow.ISODate `},
		{` *ISODate `, ` *fednow.ISODate `},
		// DateTime
		{`type ISODateTime time.Time`, ``},
		{` ISODateTime `, ` fednow.ISODateTime `},
		{` *ISODateTime `, ` *fednow.ISODateTime `},
	}

	// Walk through the directory
	err := filepath.WalkDir("gen", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Only process regular files with .go extension
		if !d.IsDir() && strings.HasSuffix(strings.ToLower(path), ".go") {
			if err := replaceInFile(path, replacements); err != nil {
				fmt.Printf("Error processing %s: %v\n", path, err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Text replacement completed.")
}

func replaceInFile(path string, replacements []struct{ old, new string }) error {
	// Read the file
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Perform replacements
	modified := data
	changed := false
	for _, r := range replacements {
		if bytes.Contains(modified, []byte(r.old)) {
			modified = bytes.ReplaceAll(modified, []byte(r.old), []byte(r.new))
			changed = true
		}
	}

	// Write back to the file only if changes were made
	if changed {
		err = os.WriteFile(path, modified, 0644)
		if err != nil {
			return err
		}
		fmt.Printf("Updated %s\n", path)
	}
	return nil
}
