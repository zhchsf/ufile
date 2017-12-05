# ufile

## Installation

```sh
go get github.com/zhchsf/ufile
```
```go
import "github.com/zhchsf/ufile"
```

## Getting Started

config ufile settings:
```go
ufile.SetConfig(
    ufile.Config{
        PublicKey:    GlobolConfig.Ufile.PublicKey,
        PrivateKey:   GlobolConfig.Ufile.PrivateKey,
        UserAgent:    GlobolConfig.Ufile.UserAgent,
        ExpireSecond: GlobolConfig.Ufile.ExpireSecond,
        UrlSuffix:    GlobolConfig.Ufile.UrlSuffix,
    },
)
```

upload to ufile:
```go
ufile.Upload(filepath string, key string, bucketName string) (bool, error)
```

get 私密地址：
```go
ufile.PrivateCloudUrl(key string, bucketName string) string
```
