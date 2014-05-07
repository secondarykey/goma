package util

import (
	"os"
	"bufio"
)

type ReadLine struct {
	rd *bufio.Reader
}

func NewReadLine(name string) (*ReadLine, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rdr := bufio.NewReaderSize(f, 4096)
	rl := &ReadLine {
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

