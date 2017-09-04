package ufile

import (
  "testing"
  // "fmt"
)

func testNewConfig() Config {
  return Config {
    PublicKey:    "public_key",
    PrivateKey:   "private_key",
    UserAgent:    "ymj go client",
    ExpireSecond: 300,
    UrlSuffix:    ".cn-bj.ufileos.com",
  }
}

func TestSetConfig(t *testing.T) {
  conf := testNewConfig()
  // fmt.Println(SetConfig(conf))
  if SetConfig(conf) == conf {
    t.Log("test SetConfig passed")
  } else {
    t.Error("test SetConfig failed")
  }
}
