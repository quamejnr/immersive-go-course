package main

import (
	"slices"
	"testing"
)

func TestBuffer(t *testing.T) {
	t.Run("Test Bytes()", func(t *testing.T) {
		tString := "Hello"
		b := NewBufferString(tString)
		got := b.Bytes()

		if slices.Compare(got, []byte(tString)) != 0 {
			t.Errorf("wanted %s, got %s\n", tString, string(got))
		}
	})

	t.Run("Test Write()", func(t *testing.T) {
		b := NewBufferString("Hello")
		_, err := b.Write([]byte("World"))
		if err != nil {
			t.Fatalf("error writing to file %v", err)
		}
		got := b.Bytes()
		expected := []byte("HelloWorld")
		if slices.Compare(got, expected) != 0 {
			t.Errorf("wanted %s, got %s\n", string(expected), string(got))
		}
	})

	t.Run("Test Read() big", func(t *testing.T) {
		tString := "Hello"
		b := NewBufferString(tString)
		buf := make([]byte, 6)
		got, err := b.Read(buf)
		if err != nil {
			t.Fatalf("error writing to file %v", err)
		}
		expected := 5
		if got != expected {
			t.Errorf("wanted %d, got %d\n", expected, got)
		}
	})
	t.Run("Test Read() small", func(t *testing.T) {
		tString := "Hello"
		b := NewBufferString(tString)
		buf := make([]byte, 2)
		got, err := b.Read(buf)
		if err != nil {
			t.Fatalf("error writing to file %v", err)
		}
		expected := len(buf)
		if got != expected {
			t.Errorf("wanted %d, got %d\n", expected, got)
		}
	})
}
