package csv

import (
	"strings"
	"testing"
)

/*
name,high score
Aya,10
Prisha,30
Charlie,-1
Margot,25
*/
func TestParser(t *testing.T) {
	t.Run("Test successful parsing", func(t *testing.T) {
		tests := []struct {
			content  string
			expected map[string]int
		}{
			{
				content:  "name,high score\nAya,10\nPrisha,30",
				expected: map[string]int{"Aya": 10, "Prisha": 30},
			},
			{
				content:  "name,high score\nCharlie,-1\nMargot,25",
				expected: map[string]int{"Charlie": -1, "Margot": 25},
			},
		}
		var p Parser

		for _, tt := range tests {
			s := strings.NewReader(tt.content)
			records, err := p.Parse(s)
			if err != nil {
				t.Error(err)
			}
			if !compareMaps(records, tt.expected) {
				t.Fatalf("wanted %v, got %v\n", tt.expected, records)
			}
		}
	})
}

func compareMaps(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}
	for k, v1 := range map1 {
		v2, ok := map2[k]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}
