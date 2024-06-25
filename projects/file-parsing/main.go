package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/quamejnr/parser/parsers"
	"github.com/quamejnr/parser/parsers/binary"
	"github.com/quamejnr/parser/parsers/csv"
	"github.com/quamejnr/parser/parsers/json"
	"github.com/quamejnr/parser/parsers/repeated_json"
)

type scoreRecord struct {
	name  string
	score int
}

func getParser(fname string) (parsers.Parser, error) {
	parsers := map[string]parsers.Parser{
		"data.csv":             csv.Parser{},
		"json.txt":             json.Parser{},
		"repeated-json.txt":    repeated_json.Parser{},
		"custom-binary-be.bin": binary.Parser{},
		"custom-binary-le.bin": binary.Parser{},
	}

	p, ok := parsers[fname]
	if !ok {
		err := fmt.Errorf("file extension %s not supported", fname)
		return nil, err
	}
	return p, nil
}

func getMaxMinScore(records map[string]int) (high scoreRecord, low scoreRecord) {
	var max, min int

	for k, v := range records {
		if v > max {
			max = v
			high.name = k
			high.score = v
		}
		if v < min {
			min = v
			low.name = k
			low.score = v
		}
	}
	return high, low
}

func main() {
	paths := os.Args[1:]

	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening file: %s\n", err)
			return
		}

		_, fname := filepath.Split(path)
		p, err := getParser(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error getting parser: %v\n", err)
			return
		}
		records, err := p.Parse(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing data: %s", err)
			return
		}

		high, low := getMaxMinScore(records)

    fmt.Println(path)
		fmt.Printf("%s had the highest score %d\n", high.name, high.score)
		fmt.Printf("%s had the lowest score %d\n", low.name, low.score)
    fmt.Println()

	}
}
