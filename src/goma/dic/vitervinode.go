package dic

type ViterbiNode struct {
	Cost    int32
	Prev    interface{}
	WordId  int32
	LeftId  int16
	RightId int16
	Start   int32
	Length  int16
	IsSpace bool
}

func NewViterbiNode(wid int32, beg int32, len int16, wordCost int16, l int16, r int16, space bool) *ViterbiNode {
	this := new(ViterbiNode)

	this.WordId = wid
	this.LeftId = l
	this.RightId = r
	this.Length = len
	this.Cost = wordCost
	this.IsSpace = space
	this.Start = beg
	return this
}

func MakeBOSEOS() *ViterbiNode {
	return NewViterbiNode(0, 0, 0, 0, 0, 0, false)
}
