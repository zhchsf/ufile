package ufile

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	base64 "encoding/base64"
)

// 签名 & base64
func signature(signStr string) string {
	hmac_str := hmac.New(sha1.New, []byte(config.PrivateKey))
	hmac_str.Write([]byte(signStr))
	return base64.StdEncoding.EncodeToString(hmac_str.Sum(nil))
}

// 上传时的signature
func uploadSignedStr(key string, bucketName string, fileType string) string {
	var buf bytes.Buffer
	buf.WriteString("PUT\n\n")
	buf.WriteString(fileType)
	buf.WriteString("\n\n/")
	buf.WriteString(bucketName)
	buf.WriteString("/")
	buf.WriteString(key)
	return signature(buf.String())
}

// 私有空间访问的signature
func downloadSignedStr(key string, bucketName string, expireStr string) string {
	var buf bytes.Buffer
	buf.WriteString("GET\n\n\n")
	buf.WriteString(expireStr)
	buf.WriteString("\n/")
	buf.WriteString(bucketName)
	buf.WriteString("/")
	buf.WriteString(key)

	return signature(buf.String())
}
