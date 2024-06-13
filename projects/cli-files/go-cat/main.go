package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	for _, fname := range os.Args[1:] {
		f, err := os.Open(fname)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
		out := cat(f)
		fmt.Print(out)
		f.Close()
	}
}

func cat(r io.Reader) string {
	out := new(bytes.Buffer)
	if _, err := io.Copy(out, r); err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	return out.String()
}
