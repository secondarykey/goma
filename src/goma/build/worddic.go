package build

import (
	"io"
	"strings"
	"fmt"
	"goma/util"
)

func buildWordDic(input,output string)(error) {

	//input = unk.def + ***.csv

	//単語のID word2id
	//単語データ word.dat
	//単語配列インデックス word.ary.idx

	//ファイルを開く
	err := buildWordIdMap(input,output)
	if err != nil {
		return err
	}
	return buildWordInfo(input,output)
}

func buildWordIdMap(input,output string) error {
	keyList,err := collectKey(input + "/unk.def",KEY_PREFIX)
	if err != nil {
		return err
	}

	for _,data := range keyList {
		fmt.Println(data)
	}
	return nil
}

func collectKey(path string,prefix string) ([]string,error){
	rd,err := util.NewReadLine(path)
	if err !=  nil {
		return nil,err
	}

	var list []string
	for {
		line,err := rd.Read()
		if err != io.EOF {
			return nil,err
		} else {
			break
		}
		idx := strings.Index(line,",")
		data := line[:idx]
		list = append(list,prefix + data)
	}
	return list,nil
}

func buildWordInfo(input,output string) error {
	return nil
}

