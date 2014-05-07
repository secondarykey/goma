package dic

import (
	"goma"
)

type Matrix struct {
	leftSize  int32
	rightSize int32
	matrix    []int16
}

func NewMatrix(dir string) (*Matrix, error) {
	mtx := new(Matrix)
	fm, err := goma.NewFileMapped(dir + "/matrix.bin")
	if err != nil {
		return nil, err
	}
	mtx.leftSize = fm.getInt()
	mtx.rightSize = fm.getInt()
	mtx.matrix = fm.getInt16Array(mtx.leftSize * mtx.rightSize)
	fm.Close()
	return mtx, nil
}

func (this *Matrix) LinkCost(leftId, rightId int32) {
	return this.matrix[rightId*this.leftSize+leftId]
}
