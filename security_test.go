package ufile

import (
	"testing"
	// "fmt"
)

func Test_uploadSignedStr(t *testing.T) {
	SetConfig(testNewConfig())
	str := uploadSignedStr("1111", "ymj", "image/png")
	// expectedStr := signature("PUT\n\nimage/png\n\n/ymj/1111")
	expectedStr := "sCTyUSEJ+O1VAiOIg0NVLUeATIk="
	if str != expectedStr {
		t.Error("test uploadSignedStr failed")
	} else {
		t.Log("test uploadSignedStr passed")
	}
}

func Test_downloadSignedStr(t *testing.T) {
	SetConfig(testNewConfig())
	str := downloadSignedStr("1111", "ymj", "1504513724")
	expectedStr := "y6eecjZwALQoPPC0obMzR1NYtsc="
	if str != expectedStr {
		t.Error("test downloadSignedStr failed")
	} else {
		t.Log("test downloadSignedStr passed")
	}
}
