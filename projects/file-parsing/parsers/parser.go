package parsers

import "io"

type Parser interface {
	Parse(io.Reader) (map[string]int, error)
}
