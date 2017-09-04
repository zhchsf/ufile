package ufile

type Config struct {
  PublicKey    string
  PrivateKey   string
  UserAgent    string
  ExpireSecond int64
  UrlSuffix    string
}

var config Config

func SetConfig(conf Config) Config {
  config = conf
  return config
}
