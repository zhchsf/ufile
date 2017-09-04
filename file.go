package ufile

import (
  "os/exec"
  // "fmt"
  "strings"
)

// // 获取文件大小
// func fileSize(file string) (int64, error) {
//   f, e := os.Stat(file)
//   if e != nil {
//     return 0, e
//   }
//   return f.Size(), nil
// }

// 获取文件mime-type
// 返回的mime-type包含换行符，末尾去掉
func fileMimeType(filepath string) string {
  mime, err := exec.Command("file", "--brief", "--mime-type", filepath).Output()
  if err != nil {
    return ""
  }
  return strings.Replace(string(mime), "\n", "", -1) 
}
