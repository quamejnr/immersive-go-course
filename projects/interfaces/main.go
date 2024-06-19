package main

type OurByteBuffer struct {
	buf      []byte
	lastRead int
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
	s := b.buf[b.lastRead:]
	for i := 0; i < len(p) && i < len(s); i++ {
		p[i] = s[i]
		count++
	}
	if count < len(s) {
		b.lastRead += count
	} else {
		b.lastRead = 0
	}
	return count, nil
}
