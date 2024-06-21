package main

import (
	"unicode"
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
	var i int
	for i = 0; i < len(p) && i < len(b.buf); i++ {
		p[i] = b.buf[i]
	}
	b.buf = b.buf[i:]
	return i, nil
}

type FilteringPipe struct {
	OurByteBuffer
}

func (f *FilteringPipe) Write(p []byte) (int, error) {
	for _, v := range p {
		if !unicode.IsDigit(rune(v)) {
			f.buf = append(f.buf, v)
		}
	}
	return len(p), nil
}

func NewFilteringPipe(w OurByteBuffer) *FilteringPipe {
	return &FilteringPipe{w}
}
