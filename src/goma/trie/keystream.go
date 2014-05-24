package trie

type KeyStream struct {
	s   string
	cur int32
}

func NewKeyStream(key string) *KeyStream {
	this := new(KeyStream)
	this.s = key
	this.cur = 0
	return this
}

func NewKeyStreamStart(key string, start int32) *KeyStream {
	this := new(KeyStream)
	this.s = key
	this.cur = start
	return this
}

func (this *KeyStream) Rest() string {
	src := this.s[this.cur:]
	dst := make([]byte,len(src))
	copy(dst,src)
	return string(dst)
}

func (this *KeyStream) Read() uint16 {
	if this.Eos() {
		return Node.Chck.TERMINATE_CODE
	} else {
		rt := this.s[this.cur]
		this.cur++
		return uint16(rt)
	}
}

func (this *KeyStream) Eos() bool {
	return (this.cur == int32(len(this.s)))
}
