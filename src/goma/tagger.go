package goma

import (
	"goma/dic"
)

type Tagger struct {
	wdc *dic.WordDic
	unk *dic.Unknown
	mtx *dic.Matrix
}

func NewTagger(dir string) (*Tagger,error) {
	tg := new(Tagger)
	var err error
	tg.wdc, err = dic.NewWordDic(dir)
	if err != nil {
		return nil,err
	}
	tg.unk,err = dic.NewUnknown(dir)
	if err != nil {
		return nil,err
	}
	tg.mtx,err = dic.NewMatrix(dir)
	if err != nil {
		return nil,err
	}
	return tg,nil
}

func (this *Tagger) Wakati(word string) ([]string,error) {
}

func (this *Tagger) Parse(word string) ([]Morpheme,error) {
}
