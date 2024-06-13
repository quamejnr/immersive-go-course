package main

import (
	"testing"
	"testing/fstest"
)

func TestCat(t *testing.T) {

	tests := []struct {
		filename string
		content  string
	}{
		{"hello.txt", "Hello World"},
		{"hi.txt", "Hi World"},
	}

  // create temporary filesystem
	fs := make(fstest.MapFS)
	for _, tt := range tests {
		fs[tt.filename] = &fstest.MapFile{Data: []byte(tt.content)}
	}

	t.Run("Testing concatenation", func(t *testing.T) {
		for _, tt := range tests {
			f, err := fs.Open(tt.filename)
			if err != nil {
				t.Fatalf("error opening file %v", err)
			}

			got := cat(f)
			expected := tt.content

			if got != expected {
				t.Errorf("got %q wanted %q", got, expected)
			}

		}
	})
}
