package sdk

import (
	"errors"
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/captcha"
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/cloudc"
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk/types"
	"github.com/iotdreamcatcher/cloud-sdk-go/third_party/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Captcha struct {
	captcha.CaptchaRpcServiceClient
	Num         int64
	Status      int
	Retry       int
	Configs     map[string]*cloudc.ModelCaptchaConfig
	UsingConfig *cloudc.ModelCaptchaConfig
}

func NewCaptcha(RpcClientConf *zrpc.RpcClientConf) *Captcha {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	client := captcha.NewCaptchaRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &Captcha{
		CaptchaRpcServiceClient: client,
		Num:                     1,
	}
}

func (s *Captcha) GetAllConfigs() map[string]*cloudc.ModelCaptchaConfig {
	return s.Configs
}

func (s *Captcha) GetConfig() (*cloudc.ModelCaptchaConfig, error) {
	if s.UsingConfig != nil {
		return s.UsingConfig, nil
	}
	return nil, errors.New("not set config")
}

func (s *Captcha) SetConfig(key string) error {
	if _, ok := s.Configs[key]; !ok {
		return errors.New("not exist")
	}
	s.UsingConfig = s.Configs[key]
	return nil
}

func (c *Sdk) InitCaptcha() *Sdk {
	c.Captcha = NewCaptcha(c.Config.RpcClientConf)
	c.Captcha.Status = types.STATUS_READY
	c.Captcha.Retry += 1
	_, err := c.CaptchaLoadCloudConfig()
	if err != nil {
		logx.Errorf("oss cloud config sync err: %v+", err)
		return c
	}
	return c
}

func (c *Sdk) CaptchaCheckStatus() *Sdk {
	if c.Captcha.Status != types.STATUS_READY {
		if c.Captcha.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitCaptcha()
		}
	}
	return c
}

func (c *Sdk) CaptchaLoadCloudConfig() (*Sdk, error) {
	// note: 加载云端配置
	res, err := c.CloudCCheckStatus().CloudCCaptchaConfigGetAll()
	if err != nil {
		logx.Errorf("captcha cloud config sync err: %v+", err)
		return c, err
	}
	if res.Code != response.SUCCESS {
		logx.Errorf("captcha cloud config sync err: %v+", res.Msg)
		return c, errors.New(res.Msg)
	}
	temp := make(map[string]*cloudc.ModelCaptchaConfig)
	temp = res.Data.Configs
	// note: 判断是否存在default的配置文件，如果不存在，给予用户警告提示
	if _, ok := temp["default"]; !ok {
		logx.Alert("请注意，您的账号下没有default的配置文件，将会导致验证码发送失败，发送前请SetConfig(key string)")
	}

	c.Captcha.Configs = temp
	c.Captcha.UsingConfig = temp["default"]
	return c, nil
}

func (c *Sdk) CaptchaGenerate() (*captcha.CaptchaResp, error) {
	// note: 读取当前使用配置
	if c.Captcha.UsingConfig == nil {
		c.Captcha.UsingConfig = c.Captcha.Configs["default"]
	}
	res, err := c.Captcha.CaptchaGenerate(c.SonyCtx(), &captcha.CaptchaReq{
		Key:       c.Captcha.UsingConfig.Key,
		DotCount:  c.Captcha.UsingConfig.DotCount,
		MaxSkew:   c.Captcha.UsingConfig.MaxSkew,
		KeyLong:   c.Captcha.UsingConfig.KeyLong,
		ImgWidth:  c.Captcha.UsingConfig.ImgWidth,
		ImgHeight: c.Captcha.UsingConfig.ImgHeight,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
