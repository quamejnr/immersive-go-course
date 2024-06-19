package main

import (
	"bytes"
	"fmt"
)

type OurByteBuffer struct {
	buf []byte
}

func NewBufferString(s string) *OurByteBuffer {
	var b OurByteBuffer
	b.buf = []byte(s)
	return &b
}

func (b *OurByteBuffer) Write(p []byte) (int, error) {
	b.buf = append(b.buf, p...)
	return len(p), nil
}

func (b *OurByteBuffer) Bytes() []byte {
	return b.buf
}

func (b *OurByteBuffer) Read(p []byte) (int, error) {
	var count int
	for i := 0; i < len(p) && i < len(b.buf); i++ {
		p[i] = b.buf[i]
		count++
	}
	b.buf = b.buf[count:]
	return count, nil
}

