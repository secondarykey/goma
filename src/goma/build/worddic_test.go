package build

import "testing"


func TestBuildWordIdMap(t *testing.T) {

	input := "/home/secondarykey/IdeaProjects/igo/dictionary"
	output := "/home/secondarykey/IdeaProjects/igo/ipadic"

	buildWordIdMap(input,output)

}
