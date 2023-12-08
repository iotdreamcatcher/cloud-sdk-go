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

	s := sdk.NewSdk().WithConfig(sdk.DefaultConfig(AccessKeyId, AccessKeySecret, []string{Endpoint})).AutoAuth().InitCloudC().InitIot()

	// note: 以下函数可以检查cloudc是否初始化成功，不调用也可以
	s = s.IotCheckStatus()

	//res, err := s.IotCheckStatus().IotHomeCreate(&iot.CreateHomeReq{
	//	Name:    "测试家庭",
	//	Address: "测试地址",
	//	UserId:  6,
	//})
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//logx.Infof("打印一下请求的结果:%+v", res)

	//res, err := s.IotCheckStatus().IotRoomCreate(&iot.CreateRoomReq{
	//	Name:   "卧室",
	//	UserId: 6,
	//	HomeId: 1,
	//})
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//logx.Infof("打印一下请求的结果:%+v", res)

	//res, err := s.IotCheckStatus().IotHomeQuery(&iot.QueryReq{
	//	Id: 1,
	//})
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//
	//databt, _ := json.Marshal(res.Data)
	//
	//logx.Infof("打印一下请求的结果:%s", string(databt))

	//res, err := s.IotCheckStatus().IotHomeQueryList(&iot.QueryHomeListReq{
	//	UserId: 6,
	//})
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//
	//databt, _ := json.Marshal(res.Data)
	//
	//logx.Infof("打印一下请求的结果:%s", string(databt))
}
