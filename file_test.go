package ufile

import (
	"testing"
	// "fmt"
)

func Test_fileMimeType(t *testing.T) {
	// fmt.Println("--------------")
	// fmt.Println(mimeType("./ufile.go"))
	// fmt.Println("--------------")
	if fileMimeType("./ufile.go") != "text/plain" {
		t.Error("test mimeType failed")
	} else {
		t.Log("test mimeType passed")
	}
}
