package build

import (
	"testing"

)

func TestBuildWordIdMap(t *testing.T) {

	input := "/home/secondarykey/IdeaProjects/igo/dictionary"
	output := "/home/secondarykey/IdeaProjects/igo/ipadic"

	err := buildWordIdMap(input,output)
	if err != nil {
		t.Error(err)
	}

	err := buildWordIdMap(input,output)

}
