# cloud-sdk-go

## 项目简介

1. golang sdk
2. 即开即用
3. 接入自动签名
4. 接入云端配置函数
5. 接入图片验证码的生成
6. 接入邮件模块
7. 接入短信模块
8. 接入子账号模块
9. 接入微信模块
10. 接入软文模块

### 安装

使用 `go get` 下载安装 SDK:

```sh
go get -u github.com/iotdreamcatcher/cloud-sdk-go
```


### 快速开始

```go
package main

import (
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/cloudc"
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

func main() {

	Endpoint := os.Getenv("ENDPOINT")

	// 本示例从环境变量中获取AccessKey ID和AccessKey Secret。
	AccessKeyId := os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	s := sdk.NewSdk().WithConfig(sdk.DefaultConfig(AccessKeyId, AccessKeySecret, []string{Endpoint})).AutoAuth().InitCloudC()

	res, err := s.CloudCCheckStatus().CloudCWechatConfigGet(&cloudc.ConfigGetParams{
		Key: "default",
	})
	if err != nil {
		return
	}
	logx.Infof("打印一下请求的结果:%+v", res)
}

```
