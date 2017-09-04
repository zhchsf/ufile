package ufile

import (
  "bytes"
  "strconv"
  "time"
  "os"
  "net/http"
  // "fmt"
  "bufio"
  "errors"
)

// filepath：上传的文件绝对路径
// key：ufile的key，要自己保证唯一性
// bucketName: bucket name
func Upload(filepath string, key string, bucketName string) (bool, error) {
  fi, err := os.Stat(filepath)
  if err != nil {
    return false, err
  }
  fileSize := string(fi.Size())

  fileType := fileMimeType(filepath)
  signedStr := uploadSignedStr(key, bucketName, fileType)
  authorization := "UCloud " + config.PublicKey + ":" + signedStr
  url := uploadUrl(key, bucketName)

  file, err := os.Open(filepath)
  defer file.Close()
  if err != nil {
    return false, err
  }

  client := &http.Client{}
  req, err := http.NewRequest("PUT", url, bufio.NewReader(file))
  req.Header.Set("User-Agent", config.UserAgent)
  req.Header.Set("Content-Length", fileSize)
  req.Header.Set("Content-Type", fileType)
  req.Header.Set("Authorization", authorization)
  response, err := client.Do(req)
  if err != nil {
    return false, err
  }
  if response.StatusCode == 200 {
    return true, nil
  } else {
    return false, errors.New("上传到ufile失败，请检查")
  }
  return true, nil
}

func uploadUrl(key string, bucketName string) string {
  var buf bytes.Buffer
  buf.WriteString("http://")
  buf.WriteString(bucketName)
  buf.WriteString(config.UrlSuffix)
  buf.WriteString("/")
  buf.WriteString(key)
  return buf.String()
}

// 私有空间文件访问url
func PrivateCloudUrl(key string, bucketName string) string {
  expire := config.ExpireSecond + time.Now().Unix()
  expireStr := strconv.FormatInt(expire, 10)
  signedStr := downloadSignedStr(key, bucketName, expireStr)

  var buf bytes.Buffer
  buf.WriteString("http://")
  buf.WriteString(bucketName)
  buf.WriteString(config.UrlSuffix)
  buf.WriteString("/")
  buf.WriteString(key)
  buf.WriteString("?UCloudPublicKey=")
  buf.WriteString(config.PublicKey)
  buf.WriteString("&Expires=")
  buf.WriteString(expireStr)
  buf.WriteString("&Signature=")
  buf.WriteString(signedStr)
  return buf.String()
}

