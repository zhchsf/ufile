package ufile

import (
	"testing"
	// "fmt"
)

// 测试上传失败的情况，配置错误
// 使用正确配置可以上传成功，暂时不知道如何mock http请求，后续继续完善
func TestUploadFailed(t *testing.T) {
	SetConfig(testNewConfig())
	success, _ := Upload("./ufile_test.go", "ufile_go_testing_file", "ymjimages")
	// fmt.Println(err)
	if success {
		t.Error("test Upload failed")
	} else {
		t.Log("test Upload passed")
	}
}

func Test_uploadUrl(t *testing.T) {
	SetConfig(testNewConfig())
	if uploadUrl("1111", "ymj") == "http://ymj.cn-bj.ufileos.com/1111" {
		t.Log("test uploadUrl passed")
	} else {
		t.Error("test uploadUrl failed")
	}
}

// 只简单调用了一次，测试待完善: mock time
func TestPrivateCloudUrl(t *testing.T) {
	SetConfig(testNewConfig())
	if PrivateCloudUrl("1111", "ymj") != "" {
		t.Log("test PrivateCloudUrl passed")
	} else {
		t.Error("test PrivateCloudUrl failed")
	}
}
