package dic

import (
	"goma"
)

type CharCategory struct {
	categorys []*Category
	char2id   []int
	eqlMasks  []int
}

type Category struct {
	id     int32
	length int32
	invoke bool
	group  bool
}

func NewCharCategory(dir string) (*CharCategory, error) {
	cc := new(CharCategory)
	var err error
	cc.categorys, err = cc.readCategorys(dir)
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func (this *CharCategory) readCategorys(dir string) ([]*Category, error) {
	fm, err := goma.NewFileMapped(dir + "/char.category")
	if err != nil {
		return nil, err
	}
	data := fm.getInt32Array(fm.size())
	size := len(data) / 4
	cgs := make([]Category, size)
	for i := 0; i < size; i++ {
		cgs[i] = NewCategory(data[i*4], data[i*4+1], data[i*4+2] == 1, data[i*4+3] == 1)
	}
	return cgs, nil
}

func (this *CharCategory) category(code byte) *Category {
	return this.categorys[this.char2id[code]]
}

func (this *CharCategory) isCompatible(code1, code2 byte) bool {
	return this.eqlMasks[code1]&this.eqlMasks[code2] != 0
}

func NewCategory(i int32, l int32, iv bool, g bool) *Category {
	cg := new(Category)
	cg.id = i
	cg.length = l
	cg.invoke = iv
	cg.group = g
	return cg
}
