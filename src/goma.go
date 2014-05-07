package main

import (
	"fmt"
	"goma"
	"goma/build"
)

func Usage() {
	fmt.Println("Usage:")
	fmt.Print("go run goma.go")
	fmt.Print("  --mode <wakati|parse|build>")
	fmt.Println("  --input <dictionary directory>")
	fmt.Println("")
	fmt.Println("In order to get your consumerkey and consumersecret, you must register an 'app' at twitter.com:")
	fmt.Println("https://dev.twitter.com/apps/new")
}

func main() {
	fmt.Printf("Hello world!")
	mode := "build"
	word := "すもももももももものうち"
	input := "dictionary"
	binaryDir := "ipadic"

	if mode == "build" {
		err := build.Build(input,binaryDir)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success!")
	} else {
		tagger,err := goma.NewTagger(binaryDir)
		if err != nil {
			panic(err)
		}

		if mode == "wakati" {
			wordArray,err := tagger.Wakati(word)
			if err != nil {
				panic(err)
			}
			for word := range wordArray {
				fmt.Print(word + " ")
			}
		} else if mode == "parse" {
			morpArray,err := tagger.Parse(word)
			if err != nil {
				panic(err)
			}
			for morp := range morpArray {
				fmt.Println(morp.Suface + "\t" + morp.Feature)
			}

		}
	}
}
