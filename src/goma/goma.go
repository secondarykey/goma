package goma

import (
	"io/ioutil"
	"os"
	"bufio"
)

type FileMapped struct {
	content []byte
	cur     int32
}

func NewFileMapped(name string) (*FileMapped, error) {
	fm := new(FileMapped)
	var err error
	fm.content, err = ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return fm, nil
}

func (this *FileMapped) GetInt32() int32 {
	return int32(this.mapping(4))
}

func (this *FileMapped) GetInt32Array(size int32) []int32 {
	return []int32(this.mapping(size))
}

//ShortArray
func (this *FileMapped) GetInt16Array(size int32) []int16 {
	return []int16(this.mapping(size * 2))
}

func (this *FileMapped) GetUint16Array(size int32) []uint16 {
	return this.mapping(size * 2)
}

func (this *FileMapped) Size() int32 {
	return len(this.content)
}

func (this *FileMapped) mapping(size int32) []byte {
	this.cur += size
	return this.content[this.cur-size : size]
}



