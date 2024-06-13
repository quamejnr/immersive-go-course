package main

import (
	"testing"
	"testing/fstest"
)

func TestWC(t *testing.T) {

	tests := []struct {
		filename string
		content  string
		expected Count
	}{
		{"1.txt", "Hello World", Count{Lines: 1, Words: 2, Bytes: 11}},
		{"2.txt", "Hello World\nI am good", Count{Lines: 2, Words: 5, Bytes: 20}},
		{"3.txt", "Hello World\nGo is fun\n", Count{Lines: 2, Words: 5, Bytes: 20}},
		{"4.txt", "One line\n", Count{Lines: 1, Words: 2, Bytes: 8}},
		{"5.txt", "", Count{Lines: 0, Words: 0, Bytes: 0}},
		{"6.txt", "Multiple words on a single line", Count{Lines: 1, Words: 6, Bytes: 31}},
	}

	fs := make(fstest.MapFS)
	for _, tt := range tests {
		fs[tt.filename] = &fstest.MapFile{Data: []byte(tt.content)}
	}

	t.Run("Testing wc", func(t *testing.T) {
		for _, tt := range tests {
			f, err := fs.Open(tt.filename)
			if err != nil {
				t.Fatalf("error opening a file %v", err)
			}
			got := wc(f)
			if got != tt.expected {
				t.Errorf("wanted %v got %v", tt.expected, got)
			}
		}
	})

}
