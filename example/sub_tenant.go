/*
*

	@author: taco
	@Date: 2023/10/27
	@Time: 11:12

*
*/
package main

import (
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/tenant"
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
	s = s.AutoAuth().InitSubTenant()
	s.SubTenantCheckStatus()

	result1, err := s.AddTenant(&tenant.EditTenantParams{
		Account:     "xxx",
		Password:    "xxx",
		Phone:       "xxx",
		Email:       "xxxx",
		Name:        "xxx",
		Avatar:      "",
		ParentId:    0,
		AccountType: 1,
	})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result1)

	//note: test tenantList
	result2, err := s.TenantList(&tenant.TenantListParams{})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result2)

	//note: test tenantList
	result3, err := s.TenantInfo(&tenant.TenantInfoParams{
		Id: 1,
	})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result3)

	//note: test tenantList
	result4, err := s.PwdTenant(&tenant.EditTenantParams{
		ID:       9,
		Password: "huanghua",
	})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result4)

	result5, err := s.EditTenant(&tenant.EditTenantParams{
		ID:          9,
		Email:       "edit@qq.com",
		Name:        "edit_test",
		Avatar:      "edit_test",
		AccountType: 1,
	})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result5)

	result6, err := s.DeleteTenant(&tenant.DeleteTenantParams{
		ID: 2,
	})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result6)

}
