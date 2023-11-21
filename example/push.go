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

	/*	iosPush, err := s.PushIosPush(&push.IosPushReq{
			Key:     "jiachao_ios",
			Msg:     "测试推送",
			IsDebug: false,
			Token: []string{
				//"654b61c9eec16a5fd731ce223a60095d3da58e0570c37885102d8c5ed2ff2967",
				//"46298ebf8a0079e20ec89bc834c7b51a16217ff9b375d6b6136402db5c14dcc4",
				//"add718349b45dc230ac8c45be925788f6e4016b27656c9449f62f4775c1aad3c",
				//"db5b4b3ffbb874d7c98c4a001db627d413a62fd00db7c28254e9a5d108c914dc",
				"579fe9714be714158c17db0857cd4a88cae84515ccd44bb7e26e5e185fba140f",
				"c56493a04da69314dc9fafaafca280e1ab3fa00b8b0adcb85c5b7acd0b8d83cd",

				//"9fcabcf2a1adddac83f843ea3b5dee0d2a3ef62cb6824a2cf63ce8cc0dde7a8e",
				//"480a45a1ffebad87ce3aa3f225d81a31410f5ebe70c6f729611ddd0d365272ad",
				//"f8c2ff862583604dd00f037de38f3e0dee65d6a7a753dd522980764e326a99ff",
			},
			Type:   "",
			Action: "",
		})
		if err != nil {
			logx.Errorf("打印一下请求的错误:%+v", err)
			return
		}
		logx.Infof("打印一下请求的结果:%+v", iosPush)*/

	xiaomiPush, err := s.PushXiaomiPush(&push.AndroidPushReq{
		Key:     "jiachao_xiaomi",
		Msg:     "test",
		IsDebug: false,
		Token: []string{
			"QLGLprJh/icXPhAQ64JdpeB2z7c2upK7cDmgWm21UIq30ZtypRV8KquiVtNnblKk",
			//"IQAAAACy0vPXAAAogoHarTDcuHDSTKO8WhXDDQ9yP8A7EyvwzXkWNVep3OrBKTYjUsy8u0XC9ulPSQoWHyrnD-pcJgVpyWs9s5oMiCLFdUINUbXI1A",
		},
		Type:   "intent",
		Action: "intent:#Intent;component=com.dc.jiachao/com.login.activity.CreateActivity;end",
		Scheme: "https://www.bejson.com",
		Extra:  "",
	})
	if err != nil {
		logx.Errorf("打印一下请求的错误:%+v", err)
		return
	}
	logx.Infof("打印一下请求的结果:%+v", xiaomiPush)

	/*	huaweiPush, err := s.PushHuaweiPush(&push.AndroidPushReq{
			Key:     "jiachao_huawei",
			Msg:     "test",
			IsDebug: false,
			Token: []string{
				"IQAAAACy0vPXAABEPk8L96m5BsoMbPgyh1mbf297exzaiJhp4K3kB6YtbDxhgmZE-1KV6mGQlBRUtqV8BoQHqnW2Kxaw9jYsxI3IefAOVTzgD_CKtw",
			},
			MsgType: "DEVICE_REMINDER",
			Type:    "intent",
			Action:  "com.dc.jiachao.intent.action.test",
			Scheme:  "https://www.bejson.com",
			Extra:   "",
		})
		if err != nil {
			logx.Errorf("打印一下请求的错误:%+v", err)
			return
		}
		logx.Infof("打印一下请求的结果:%+v", huaweiPush)*/

}
