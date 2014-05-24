package trie

import (
)

type Searcher struct {
	keySetSize int32
	base []int32
	chck []uint16
	begs []int32
	lens []int16
	tail string
}

func NewSearcher(_ string) (*Searcher,error) {
	this := new(Searcher)
	return this,nil
}

func (this *Searcher) Size() int32 {
	return this.keySetSize
}

/*
func (this* Searcher) Search(key []byte) int32 {
	node := this.base[0]
	in := NewKeyStream(key)
	for code := in.Read();;code=in.Read() {
		idx := node + int32(code)
		node = this.base[idx]
		if this.chck[idx] == code {
			if node >= 0 {
				//TODO ここなに？
			} else if in.Eos() || this.keyExists(in,node) {
				return -1
			}
		}
	}
}
*/

/*
func (this *Searcher)EachCommonPrefix(key []byte,start int32,fn Callback) {
	node := this.base[0]
	offset := 0
	in := NewKeyStreamStart(key,start)

	code:=in.Read()
	for {
		terminalIdx := node + int32(Node.Chck.TERMINATE_CODE)
		if this.chck[terminalIdx] == Node.Chck.TERMINATE_CODE {
			fn.call(start,offset,Node.Base.ID(this.base[terminalIdx]))
			if code == Node.Chck.TERMINATE_CODE {
				return
			}
		}

		idx := node + int32(code)
		node = this.base[idx]

		if this.chck[idx] == code {
			if node >= 0 {
				code=in.Read()
				offset++
				continue
			} else {
				this.call_if_keyIncluding(in,node,start,offset,fn)
			}
		}
		code=in.Read()
		offset++
	}
}


func (this *Searcher)call_if_keyIncluding(in *KeyStream, node int32, start int32, offset int, fn Callback) {
	id := Node.Base.ID(node)

	if strings.Contains(this.tail,this.begs[id],this.lens[id]) {
		fn.call(start,offset+this.lens[id]+1,id)
	}
}
*/


