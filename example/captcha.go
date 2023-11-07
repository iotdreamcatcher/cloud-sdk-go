package main

import (
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

func main() {

	Endpoint := os.Getenv("ENDPOINT")

	// 本示例从环境变量中获取AccessKey ID和AccessKey Secret。
	AccessKeyId := os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	s := sdk.NewSdk().WithConfig(sdk.DefaultConfig(AccessKeyId, AccessKeySecret, []string{Endpoint})).AutoAuth().InitCloudC().InitCaptcha()

	// note: 以下函数可以检查cloudc是否初始化成功，不调用也可以
	s = s.CaptchaCheckStatus()

	result, err := s.CaptchaCheckStatus().CaptchaGenerate()
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下结果:%+v", result)
}
