package json

import (
	"encoding/json"
	"fmt"
	"io"
)

type Record struct {
	Name      string `json:"name"`
	HighScore int    `json:"high_score"`
}

type Records []Record

type Parser struct {
}

func (p Parser) Parse(r io.Reader) (map[string]int, error) {
	results := map[string]int{}
	var records Records
	if err := json.NewDecoder(r).Decode(&records); err != nil {
		error := fmt.Errorf("error parsing file %w\n", err)
		return nil, error
	}
	for _, r := range records {
		results[r.Name] = r.HighScore
	}

	return results, nil

}
