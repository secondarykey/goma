package goma

type Morpheme struct {
	Surface string
	Feature string
	Start int32
}

func NewMorpheme(surface,feature string,start int) *Morpheme {
	mp := new(Morpheme)
	mp.Surface = surface
	mp.Feature = feature
	mp.Start = start
}
