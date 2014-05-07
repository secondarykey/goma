package dic

type Unknown struct {
	category *CharCategory
	spaceId int64
}

func NewUnknown(dir string) (*Unknown, error) {
	ukn := new(Unknown)
	var err error
	ukn.category, err = NewCharCategory(dir)
	if err != nil {
		return nil, err
	}
	ukn.spaceId = ukn.category.category(" "[0]).id
	return ukn, nil
}

func (this *Unknown) search() {
}
