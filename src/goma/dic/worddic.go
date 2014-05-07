package dic

import (
	"goma/trie"
)

type Callback interface {
	call(*ViterbiNode)
	isEmpty() bool
}

type WordDic struct {
	trie        trie.Searcher
	data        string
	indices     []int32
	costs       []int16
	leftIds     []int16
	tightIds    []int16
	dataOffsets []int32
}

func NewWordDic(dir string) (*WordDic, error) {
	this := new(WordDic)
	var err error
	this.trie,err = trie.NewSearcher(dir+"/word2id")
	if err != nil {
		return nil,err
	}
	return this, nil
}

