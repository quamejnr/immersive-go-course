package main

import (
	"slices"
	"testing"
	"testing/fstest"
)

func TestLs(t *testing.T) {

	fs := fstest.MapFS{
		"hello.txt": {Data: []byte("")},
		"hi.txt":    {Data: []byte("")},
		"data.txt":  {Data: []byte("")},
	}

	dir, err := fs.ReadDir(".")
	if err != nil {
		t.Fatalf("error reading file %v", err)
	}
	got := ls(dir)
	want := []string{"data.txt", "hello.txt", "hi.txt"}

	if slices.Compare(got, want) != 0 {
		t.Errorf("got %v wanted %v\n", got, want)
	}

}
