package binary

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

type Parser struct {
}

func (p Parser) Parse(r io.Reader) (map[string]int, error) {
	results := map[string]int{}

	buf := bufio.NewReader(r)
	byteOrder, err := getEndianness(r)
	if err != nil {
		return nil, fmt.Errorf("failed to get endianness: %w", err)
	}

	for {
		if _, err := buf.Peek(1); errors.Is(err, io.EOF) {
			break
		}
		var highScore int32
		if err := binary.Read(buf, byteOrder, &highScore); err != nil {
			return nil, fmt.Errorf("error parsing high score: %w", err)
		}
		nameWithTrailingNull, err := buf.ReadString('\x00')
		if err != nil {
			return nil, fmt.Errorf("error parsing name: %w", err)
		}
		name := nameWithTrailingNull[:len(nameWithTrailingNull)-1]
		results[name] = int(highScore)
	}

	return results, nil
}

func getEndianness(r io.Reader) (binary.ByteOrder, error) {
	buf := make([]byte, 2)
	_, err := r.Read(buf)
	if err != nil {
		return nil, err
	}
	if buf[0] == '\xFE' && buf[1] == '\xFF' {
		return binary.BigEndian, nil
	} else if buf[0] == '\xFF' && buf[1] == '\xFE' {
		return binary.LittleEndian, nil
	} else {
		return nil, fmt.Errorf("byte order not recognized")
	}
}
