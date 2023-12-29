package sdk

import (
	"errors"
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/cloudc"
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/oss"
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk/types"
	"github.com/iotdreamcatcher/third_party/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Oss struct {
	oss.OssRpcServiceClient
	Num         int64
	Status      int
	Retry       int
	Configs     map[string]*cloudc.ModelOssConfig
	UsingConfig *cloudc.ModelOssConfig
}

func NewOss(RpcClientConf *zrpc.RpcClientConf) *Oss {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	client := oss.NewOssRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &Oss{
		OssRpcServiceClient: client,
		Num:                 1,
	}
}

func (s *Oss) GetAllConfigs() map[string]*cloudc.ModelOssConfig {
	return s.Configs
}

func (s *Oss) GetConfig() (*cloudc.ModelOssConfig, error) {
	if s.UsingConfig != nil {
		return s.UsingConfig, nil
	}
	return nil, errors.New("not set config")
}

func (s *Oss) SetConfig(key string) error {
	if _, ok := s.Configs[key]; !ok {
		return errors.New("not exist")
	}
	s.UsingConfig = s.Configs[key]
	return nil
}

func (c *Sdk) InitOss() *Sdk {
	c.Oss = NewOss(c.Config.RpcClientConf)
	c.Oss.Status = types.STATUS_READY
	c.Oss.Retry += 1
	_, err := c.OssLoadCloudConfig()
	if err != nil {
		logx.Errorf("oss cloud config sync err: %v+", err)
		return c
	}
	return c
}

func (c *Sdk) OssCheckStatus() *Sdk {
	if c.Oss.Status != types.STATUS_READY {
		if c.Oss.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitOss()
		}
	}
	return c
}

func (c *Sdk) OssLoadCloudConfig() (*Sdk, error) {
	// note: 加载云端配置
	res, err := c.CloudCCheckStatus().CloudCOssConfigGetAll()
	if err != nil {
		logx.Errorf("oss cloud config sync err: %v+", err)
		return c, err
	}
	if res.Code != response.SUCCESS {
		logx.Errorf("oss cloud config sync err: %v+", res.Msg)
		return c, errors.New(res.Msg)
	}
	temp := make(map[string]*cloudc.ModelOssConfig)
	temp = res.Data.Configs
	// note: 判断是否存在default的配置文件，如果不存在，给予用户警告提示
	if _, ok := temp["default"]; !ok {
		logx.Alert("请注意，您的账号下没有default的配置文件，将会导致oss生成失败，发送前请SetConfig(key string)")
	}

	c.Oss.Configs = temp
	c.Oss.UsingConfig = temp["default"]
	return c, nil
}

func (c *Sdk) OssGenerateUploadSign(Filename, UploadDir, CallBack string, IsCallBack bool) (*oss.GenerateUploadSignParamsResp, error) {
	// note: 读取当前使用配置
	if c.Oss.UsingConfig == nil {
		c.Oss.UsingConfig = c.Oss.Configs["default"]
	}
	res, err := c.Oss.GenerateUploadSign(c.SonyCtx(), &oss.GenerateUploadSignParams{
		AccessKeyId:     c.Oss.UsingConfig.AppID,
		AccessKeySecret: c.Oss.UsingConfig.AppSecret,
		Endpoint:        c.Oss.UsingConfig.Endpoint,
		Domain:          c.Oss.UsingConfig.Domain,
		Prefix:          c.Oss.UsingConfig.Prefix,
		Filename:        Filename,
		UploadDir:       UploadDir,
		IsCallBack:      IsCallBack,
		CallBack:        CallBack,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
