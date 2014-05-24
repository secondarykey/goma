package util

import (
	"os"
	"bufio"
)

type ReadLine struct {
	f *os.File
	rd *bufio.Reader
}

func NewReadLine(name string) (*ReadLine, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	rdr := bufio.NewReaderSize(f, 4096)
	rl := &ReadLine {
		f : f,
		rd : rdr,
	}
	return rl,nil
}

func (self *ReadLine) Read() (string, error) {
	l, _, err := self.rd.ReadLine()
	if err != nil {
		return "", err
	}
	return string(l), nil
}

func (self *ReadLine) Close() {
	self.f.Close()
}

