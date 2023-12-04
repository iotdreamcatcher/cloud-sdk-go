package main

import (
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk"
	"os"
)

func main() {

	Endpoint := os.Getenv("ENDPOINT")

	// 本示例从环境变量中获取AccessKey ID和AccessKey Secret。
	AccessKeyId := os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	s := sdk.NewSdk().WithConfig(sdk.DefaultConfig(AccessKeyId, AccessKeySecret, []string{Endpoint})).AutoAuth().InitCloudC().InitIotHome()

	// note: 以下函数可以检查cloudc是否初始化成功，不调用也可以
	s = s.IotHomeCheckStatus()

	//res, err := s.IotHomeCheckStatus().IotHomeCreate(&iot.CreateHomeReq{
	//	Name:    "测试家庭",
	//	Address: "测试地址",
	//	UserId:  6,
	//})
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//logx.Infof("打印一下请求的结果:%+v", res)

	//res, err := s.IotHomeCheckStatus().IotHomeDelete(&iot.DeleteReq{
	//	UserId: 6,
	//	Id:     4,
	//})
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//logx.Infof("打印一下请求的结果:%+v", res)
}
