package csv

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

/*
name,high score
Aya,10
Prisha,30
Charlie,-1
Margot,25
*/

type Parser struct{}

func (p Parser) Parse(r io.Reader) (map[string]int, error) {
	records := map[string]int{}
	s := bufio.NewScanner(r)

	skipHeader := true
	var count int
	for s.Scan() {
		count++
		if skipHeader {
			skipHeader = false
			continue
		}
		line := s.Text()
		val := strings.Split(line, ",")
		score, err := strconv.Atoi(val[1])
		if err != nil {
			err := fmt.Errorf("error converting score on line %d, %w\n", count, err)
			return nil, err
		}
		records[val[0]] = score
	}

	return records, nil
}
