package main

import (
	"encoding/json"
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

	// note: 以下函数可以检查cloudc是否初始化成功，不调用也可以
	s = s.CloudCCheckStatus()

	//res1, err := s.CloudCCheckStatus().CloudCWechatConfigSet(&cloudc.WechatConfigSetParams{
	//	Key:       "xxxxxx",
	//	GhId:      "xxxxxx",
	//	AppID:     "xxxxxx",
	//	AppSecret: "xxxxxx",
	//	AppName:   "xxxxxx",
	//})
	//if err != nil {
	//	return
	//}
	//logx.Infof("打印一下请求的结果:%+v", res1)
	//
	//res2, err := s.CloudCCheckStatus().CloudCWechatConfigGet(&cloudc.ConfigGetParams{
	//	Key: "default",
	//})
	//if err != nil {
	//	return
	//}
	//logx.Infof("打印一下请求的结果:%+v", res2)

	//res3, err := s.CloudCCheckStatus().CloudCSystemConfigSet(&cloudc.SystemConfigSetParams{
	//	Key:    "key_xx",
	//	Name:   "xx",
	//	Desc:   "xx管理服务",
	//	Domain: "xxxxx",
	//	Ip:     "xxxx",
	//	Port:   0,
	//})
	//if err != nil {
	//	return
	//}
	//logx.Infof("打印一下请求的结果:%+v", res3)
	//
	//res4, err := s.CloudCCheckStatus().CloudCSystemConfigGet(&cloudc.ConfigGetParams{
	//	Key: "key_am",
	//})
	//if err != nil {
	//	return
	//}
	//logx.Infof("打印一下请求的结果:%+v", res4)

	//type iosConfig struct {
	//	P12Path string
	//	P12Pwd  string
	//	Topic   string
	//}
	//
	//c := iosConfig{
	//	P12Path: "etc/jiachao.p12",
	//	P12Pwd:  "123456",
	//	Topic:   "com.dreamcatcher.jiachao",
	//}

	type XiaomiConfig struct {
		APPID       string `json:"appid"`
		SECRET      string `json:"secret"`
		PackageName string `json:"package_name"`
		Title       string `json:"title"`
	}

	type HuaweiConfig struct {
		AppId     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
		AuthUrl   string `json:"auth_url"`
		PushUrl   string `json:"push_url"`
	}

	//c := XiaomiConfig{
	//	APPID:       "2882303761520177745",
	//	SECRET:      "UFudmkLyKMYaWB5hCOIYuQ==",
	//	PackageName: "com.dc.jiachao",
	//	Title:       "家超",
	//}

	c := HuaweiConfig{
		AppId:     "106867201",
		AppSecret: "3d72f9e04bf4142e56fd63d06b3d364bb54e4246abb3e67c3d6c11561e6606d8",
		AuthUrl:   "https://login.cloud.huawei.com/oauth2/v3/token",
		PushUrl:   "https://api.push.hicloud.com",
	}
	cbt, _ := json.Marshal(c)

	res5, err := s.CloudCCheckStatus().CloudCPushConfigSet(&cloudc.PushConfigSetParams{
		Key:      "jiachao_huawei",
		Title:    "家超",
		Body:     string(cbt),
		Platform: 2,
	})
	if err != nil {
		return
	}
	logx.Infof("打印一下请求的结果:%+v", res5)

	res6, err := s.CloudCCheckStatus().CloudCPushConfigGet(&cloudc.ConfigGetParams{
		Key: "jiachao_xiaomi",
	})
	if err != nil {
		return
	}
	logx.Infof("打印一下请求的结果:%+v", res6)
}
