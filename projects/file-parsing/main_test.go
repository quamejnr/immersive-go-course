package main

import (
	"reflect"
	"testing"

	"github.com/quamejnr/parser/parsers"
	"github.com/quamejnr/parser/parsers/binary"
	"github.com/quamejnr/parser/parsers/csv"
	"github.com/quamejnr/parser/parsers/json"
	"github.com/quamejnr/parser/parsers/repeated_json"
)

func TestGetParser(t *testing.T) {
	subtests := []struct {
		filename string
		expected parsers.Parser
	}{
		{
			"data.csv",
			csv.Parser{},
		},
		{
			"data.csv",
			csv.Parser{},
		},
		{
			"json.txt",
			json.Parser{},
		},
		{
			"custom-binary-be.bin",
			binary.Parser{},
		},
		{
			"repeated-json.txt",
			repeated_json.Parser{},
		},
	}

	for _, tt := range subtests {
		got, err := getParser(tt.filename)
		if err != nil {
			t.Error("error parsing file", err)
		}
		if reflect.TypeOf(got) != reflect.TypeOf(tt.expected) {
			t.Fatalf("got %T expected %T", got, tt.expected)
		}

	}
}
