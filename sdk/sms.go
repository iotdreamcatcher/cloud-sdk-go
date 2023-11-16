package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/cloudc"
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/sms"
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk/types"
	"github.com/iotdreamcatcher/cloud-sdk-go/third_party/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Sms struct {
	sms.SmsRpcServiceClient
	Num         int64
	Status      int
	Retry       int
	Configs     map[string]*cloudc.ModelSmsConfig
	UsingConfig *cloudc.ModelSmsConfig
}

func NewSms(RpcClientConf *zrpc.RpcClientConf) *Sms {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	client := sms.NewSmsRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &Sms{
		SmsRpcServiceClient: client,
		Num:                 1,
	}
}

func (s *Sms) GetAllConfigs() map[string]*cloudc.ModelSmsConfig {
	return s.Configs
}

func (s *Sms) GetConfig() (*cloudc.ModelSmsConfig, error) {
	if s.UsingConfig != nil {
		return s.UsingConfig, nil
	}
	return nil, errors.New("not set config")
}

func (s *Sms) SetConfig(key string) error {
	if _, ok := s.Configs[key]; !ok {
		return errors.New("not exist")
	}
	s.UsingConfig = s.Configs[key]
	return nil
}

func (c *Sdk) InitSms() *Sdk {
	c.Sms = NewSms(c.Config.RpcClientConf)
	c.Sms.Status = types.STATUS_READY
	c.Sms.Retry += 1
	_, err := c.SmsLoadCloudConfig()
	if err != nil {
		logx.Errorf("oss cloud config sync err: %v+", err)
		return c
	}
	return c
}

func (c *Sdk) SmsCheckStatus() *Sdk {
	if c.Sms.Status != types.STATUS_READY {
		if c.Sms.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitSms()
		}
	}
	return c
}

func (c *Sdk) SmsLoadCloudConfig() (*Sdk, error) {
	// note: 加载云端配置
	res, err := c.CloudCCheckStatus().CloudCSmsConfigGetAll()
	if err != nil {
		logx.Errorf("ems cloud config sync err: %v+", err)
		return c, err
	}
	if res.Code != response.SUCCESS {
		logx.Errorf("ems cloud config sync err: %v+", res.Msg)
		return c, errors.New(res.Msg)
	}
	temp := make(map[string]*cloudc.ModelSmsConfig)
	temp = res.Data.Configs
	// note: 判断是否存在default的配置文件，如果不存在，给予用户警告提示
	if _, ok := temp["default"]; !ok {
		logx.Alert("请注意，您的账号下没有default的配置文件，将会导致短信发送失败，发送前请SetConfig(key string)")
	}

	c.Sms.Configs = temp
	c.Sms.UsingConfig = temp["default"]
	return c, nil
}

func (c *Sdk) SmsSendSms(phone string, args ...any) (*sms.SmsResp, error) {
	// note: 读取当前使用配置
	if c.Sms.UsingConfig == nil {
		c.Sms.UsingConfig = c.Sms.Configs["default"]
	}
	param := make(map[string]string)
	tempStr := fmt.Sprintf(c.Sms.UsingConfig.ParamsStr, args...)
	if err := json.Unmarshal([]byte(tempStr), &param); err != nil {
		return nil, err
	}
	res, err := c.Sms.SendSms(c.SonyCtx(), &sms.SmsParams{
		Mobile:          phone,
		Template:        c.Sms.UsingConfig.Template,
		Params:          param,
		Type:            c.Sms.UsingConfig.Type,
		SignName:        c.Sms.UsingConfig.SignName,
		AccessKeyId:     c.Sms.UsingConfig.AppID,
		AccessKeySecret: c.Sms.UsingConfig.AppSecret,
		Endpoint:        c.Sms.UsingConfig.Endpoint,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
