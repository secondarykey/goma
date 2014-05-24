package build

import (
	"io"
	"strings"
	"goma/util"
	"goma/trie"
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

	keyList := make([]string,0,50000)
	files,err := util.GetFileList(input,".csv")
	if err != nil {
		return err
	}
	for _,file := range files {
		wordList,err := collectKey(input + "/" + file,"")
		if err != nil {
			return err
		}
		keyList = append(keyList,wordList...)
	}

	unkList,err := collectKey(input + "/unk.def",KEY_PREFIX)
	if err != nil {
		return err
	}
	keyList = append(keyList,unkList...)

	builder := trie.NewBuilder(keyList)
	builder.Build()

	builder.save(output + "/word2id")
	return nil
}

func collectKey(path string,prefix string) ([]string,error){
	rd,err := util.NewReadLine(path)
	defer rd.Close()
	if err !=  nil {
		return nil,err
	}

	list := make([]string,10000)
	for {
		line,err := rd.Read()
		if err != nil {
			if err != io.EOF {
				return nil,err
			} else {
				break
			}
		}
		idx := strings.Index(line,",")
		data := line[:idx]
		list = append(list,prefix + data)
	}
	return list,nil
}

func buildWordInfo(input,output string) error {

	wid := NewSearcher()






	return nil
}

