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

	s := sdk.NewSdk().WithConfig(sdk.DefaultConfig(AccessKeyId, AccessKeySecret, []string{Endpoint})).AutoAuth().InitCloudC().InitEms().InitOss()
	login, err := s.InitAuth().AuthPing().AuthLogin()
	if err != nil {
		return
	}
	logx.Infof("打印login结果 %v", login)

	refresh, err := s.AuthRefresh()
	if err != nil {
		return
	}
	logx.Infof("打印refresh结果 %v", refresh)
	s.AutoAuth()

	select {}
}
