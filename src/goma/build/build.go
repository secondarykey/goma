package build

import "os"

func Build(input,output string) error {

	os.Mkdir(output,os.FileMode(666))
	err := buildWordDic(input,output)
	if err != nil {
		return err
	}
	err = buildMatrix(input,output)
	if err != nil {
		return err
	}
	err = buildCategory(input,output)
	if err != nil {
		return err
	}
	return nil
}


