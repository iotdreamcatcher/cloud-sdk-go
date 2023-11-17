package main

import (
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/push"
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

func main() {

	Endpoint := os.Getenv("ENDPOINT")

	// 本示例从环境变量中获取AccessKey ID和AccessKey Secret。
	AccessKeyId := os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	s := sdk.NewSdk().WithConfig(sdk.DefaultConfig(AccessKeyId, AccessKeySecret, []string{Endpoint})).AutoAuth().InitCloudC().InitPush()

	// note: 以下函数可以检查cloudc是否初始化成功，不调用也可以
	s = s.PushCheckStatus()

	iosPush, err := s.PushIosPush(&push.IosPushReq{
		Key:     "jiachao_ios",
		Msg:     "xxxx",
		IsDebug: false,
		Token: []string{
			"654b61c9eec16a5fd731ce223a60095d3da58e0570c37885102d8c5ed2ff2967",
			"46298ebf8a0079e20ec89bc834c7b51a16217ff9b375d6b6136402db5c14dcc4",
			"add718349b45dc230ac8c45be925788f6e4016b27656c9449f62f4775c1aad3c",
			"db5b4b3ffbb874d7c98c4a001db627d413a62fd00db7c28254e9a5d108c914dc",
			"59177e38f0dbcd68d6b04f9cd67b508931e78cbab96d0e37972d196492b75794",
		},
	})
	if err != nil {
		logx.Errorf("打印一下请求的错误:%+v", err)
		return
	}
	logx.Infof("打印一下请求的结果:%+v", iosPush)
}
