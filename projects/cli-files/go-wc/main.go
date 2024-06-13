package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]

	var tlc, twc, tbc int
	for _, fname := range files {
		f, err := os.Open(fname)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		c := wc(f)
		fmt.Printf("%6d %6d %6d %6s\n", c.Lines, c.Words, c.Bytes, fname)
		tlc += c.Lines
		twc += c.Words
		tbc += c.Bytes

		f.Close()
	}
	if len(files) > 2 {
		fmt.Printf("%6d %6d %6d %6s\n", tlc, twc, tbc, "total")
	}
}

type Count struct {
	Lines int
	Words int
	Bytes int
}

func wc(r io.Reader) Count {
	count := Count{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()
		count.Lines++
		count.Words += len(strings.Fields(s))
		count.Bytes += len(s)
	}

	return count

}
