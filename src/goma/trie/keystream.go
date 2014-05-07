package trie

type KeyStream struct {
	s   []byte
	cur int32
}

func NewKeyStream(key []byte) *KeyStream {
	this := new(KeyStream)
	this.s = key
	this.cur = 0
	return this
}

func NewKeyStreamStart(key []byte, start int32) *KeyStream {
	this := new(KeyStream)
	this.s = key
	this.cur = start
	return this
}

func (this *KeyStream) Rest() string {
	return this.s[this.cur:]
}

func (this *KeyStream) Read() uint16 {
	if this.Eos() {
		return Node.Chck.TERMINATE_CODE
	} else {
		this.cur++
		return this.s[this.cur:]
	}
}

func (this *KeyStream) Eos() bool {
	return (this.cur == len(this.s))
}
