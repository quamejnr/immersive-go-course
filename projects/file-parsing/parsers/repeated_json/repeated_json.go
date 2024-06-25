package repeated_json

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Record struct {
	Name      string `json:"name"`
	HighScore int    `json:"high_score"`
}

type Parser struct {
}

func (p Parser) Parse(r io.Reader) (map[string]int, error) {
	results := map[string]int{}
	var record Record

	s := bufio.NewScanner(r)

	for s.Scan() {
		t := s.Text()
    fmt.Println(t)
		if strings.HasPrefix(t, "#") {
			continue
		}
    if len(t) <= 0 {
      continue
    }
		if err := json.Unmarshal([]byte(t), &record); err != nil {
			error := fmt.Errorf("error parsing file: %w", err)
			return nil, error
		}
		results[record.Name] = record.HighScore
	}
	return results, nil

}
