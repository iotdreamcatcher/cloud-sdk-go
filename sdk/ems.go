package sdk

import (
	"errors"
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/cloudc"
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/ems"
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk/types"
	"github.com/iotdreamcatcher/cloud-sdk-go/third_party/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Ems struct {
	ems.EmsRpcServiceClient
	Num         int64
	Status      int
	Retry       int
	Configs     map[string]*cloudc.ModelEmsConfig
	UsingConfig *cloudc.ModelEmsConfig
}

func NewEms(RpcClientConf *zrpc.RpcClientConf) *Ems {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	client := ems.NewEmsRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &Ems{
		EmsRpcServiceClient: client,
		Num:                 1,
	}
}

func (s *Ems) GetAllConfigs() map[string]*cloudc.ModelEmsConfig {
	return s.Configs
}

func (s *Ems) GetConfig() (*cloudc.ModelEmsConfig, error) {
	if s.UsingConfig != nil {
		return s.UsingConfig, nil
	}
	return nil, errors.New("not set config")
}

func (s *Ems) SetConfig(key string) error {
	if _, ok := s.Configs[key]; !ok {
		return errors.New("not exist")
	}
	s.UsingConfig = s.Configs[key]
	return nil
}

func (c *Sdk) InitEms() *Sdk {
	c.Ems = NewEms(c.Config.RpcClientConf)
	c.Ems.Status = types.STATUS_READY
	c.Ems.Retry += 1
	_, err := c.EmsLoadCloudConfig()
	if err != nil {
		logx.Errorf("ems cloud config sync err: %v+", err)
		return c
	}
	return c
}

func (c *Sdk) EmsCheckStatus() *Sdk {
	if c.Ems.Status != types.STATUS_READY {
		if c.Ems.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitEms()
		}
	}
	return c
}

func (c *Sdk) EmsLoadCloudConfig() (*Sdk, error) {
	// note: 加载云端配置
	res, err := c.CloudCCheckStatus().CloudCEmsConfigGetAll()
	if err != nil {
		logx.Errorf("ems cloud config sync err: %v+", err)
		return c, err
	}
	if res.Code != response.SUCCESS {
		logx.Errorf("ems cloud config sync err: %v+", res.Msg)
		return c, errors.New(res.Msg)
	}
	temp := make(map[string]*cloudc.ModelEmsConfig)
	temp = res.Data.Configs
	// note: 判断是否存在default的配置文件，如果不存在，给予用户警告提示
	if _, ok := temp["default"]; !ok {
		logx.Alert("请注意，您的账号下没有default的配置文件，将会导致邮件发送失败，发送前请SetConfig(key string)")
	}

	c.Ems.Configs = temp
	c.Ems.UsingConfig = temp["default"]
	return c, nil
}

func (c *Sdk) EmsSendEms(RecipientEmail []string, cc []*ems.Cc, Subject, SendType, SendBody string) (*ems.EmsResp, error) {
	// note: 读取当前使用配置
	if c.Ems.UsingConfig == nil {
		c.Ems.UsingConfig = c.Ems.Configs["default"]
	}
	res, err := c.Ems.SendEmsRpc(c.SonyCtx(), &ems.EmsReq{
		RecipientEmail: RecipientEmail,
		Cc:             cc,
		Subject:        Subject,
		SendType:       SendType,
		SendBody:       SendBody,
		SenderName:     c.Ems.UsingConfig.SenderName,
		SenderMail:     c.Ems.UsingConfig.SenderMail,
		SenderPwd:      c.Ems.UsingConfig.SenderPwd,
		Host:           c.Ems.UsingConfig.Host,
		Port:           c.Ems.UsingConfig.Port,
		Protocol:       c.Ems.UsingConfig.Protocol,
		Key:            c.Ems.UsingConfig.Key,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
