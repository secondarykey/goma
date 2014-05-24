package util

import (
	"strings"
	"io/ioutil"
	"errors"
	"os"
)

func GetFileList(path,ext string) ([]string,error) {

	info,err := os.Stat(path)
	if err != nil {
		return nil,err
	}
	if !info.IsDir() {
		return nil,errors.New("not dir")
	}
	infos,err := ioutil.ReadDir(path)
	if err != nil {
		return nil,err
	}

	var files []string
	for _,info := range infos {
		name := info.Name()
		if strings.LastIndex(name,ext) != -1 {
			files = append(files,name)
		}
	}
	return files,nil
}

