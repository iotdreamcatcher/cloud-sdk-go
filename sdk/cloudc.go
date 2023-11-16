package sdk

import (
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/cloudc"
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk/types"
	"github.com/zeromicro/go-zero/zrpc"
)

type CloudC struct {
	cloudc.CloudCServiceClient
	Num    int64
	Status int
	Retry  int
}

func NewCloudC(RpcClientConf *zrpc.RpcClientConf) *CloudC {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	client := cloudc.NewCloudCServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &CloudC{
		CloudCServiceClient: client,
		Num:                 1,
	}
}

func (c *Sdk) InitCloudC() *Sdk {
	c.CloudC = NewCloudC(c.Config.RpcClientConf)
	c.CloudC.Status = types.STATUS_READY
	c.CloudC.Retry += 1
	return c
}

func (c *Sdk) CloudCCheckStatus() *Sdk {
	if c.CloudC.Status != types.STATUS_READY {
		if c.CloudC.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitCloudC()
		}
	}
	return c
}

func (c *Sdk) CloudCWechatConfigSet(in *cloudc.WechatConfigSetParams) (*cloudc.ConfigResp, error) {
	res, err := c.CloudC.WechatConfigSet(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCWechatConfigGet(in *cloudc.ConfigGetParams) (*cloudc.WechatConfigGetResp, error) {
	res, err := c.CloudC.WechatConfigGet(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCWechatConfigGetAll() (*cloudc.WechatConfigGetAllResp, error) {
	res, err := c.CloudC.WechatConfigGetAll(c.SonyCtx(), &cloudc.ConfigGetAllParams{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCSmsConfigSet(in *cloudc.SmsConfigSetParams) (*cloudc.ConfigResp, error) {
	res, err := c.CloudC.SmsConfigSet(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCSmsConfigGet(in *cloudc.ConfigGetParams) (*cloudc.SmsConfigGetResp, error) {
	res, err := c.CloudC.SmsConfigGet(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCSmsConfigGetAll() (*cloudc.SmsConfigGetAllResp, error) {
	res, err := c.CloudC.SmsConfigGetAll(c.SonyCtx(), &cloudc.ConfigGetAllParams{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCEmsConfigSet(in *cloudc.EmsConfigSetParams) (*cloudc.ConfigResp, error) {
	res, err := c.CloudC.EmsConfigSet(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCEmsConfigGet(in *cloudc.ConfigGetParams) (*cloudc.EmsConfigGetResp, error) {
	res, err := c.CloudC.EmsConfigGet(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCEmsConfigGetAll() (*cloudc.EmsConfigGetAllResp, error) {
	res, err := c.CloudC.EmsConfigGetAll(c.SonyCtx(), &cloudc.ConfigGetAllParams{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCOssConfigSet(in *cloudc.OssConfigSetParams) (*cloudc.ConfigResp, error) {
	res, err := c.CloudC.OssConfigSet(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCOssConfigGet(in *cloudc.ConfigGetParams) (*cloudc.OssConfigGetResp, error) {
	res, err := c.CloudC.OssConfigGet(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCOssConfigGetAll() (*cloudc.OssConfigGetAllResp, error) {
	res, err := c.CloudC.OssConfigGetAll(c.SonyCtx(), &cloudc.ConfigGetAllParams{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCCaptchaConfigSet(in *cloudc.CaptchaConfigSetParams) (*cloudc.ConfigResp, error) {
	res, err := c.CloudC.CaptchaConfigSet(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCCaptchaConfigGet(in *cloudc.ConfigGetParams) (*cloudc.CaptchaConfigGetResp, error) {
	res, err := c.CloudC.CaptchaConfigGet(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCCaptchaConfigGetAll() (*cloudc.CaptchaConfigGetAllResp, error) {
	res, err := c.CloudC.CaptchaConfigGetAll(c.SonyCtx(), &cloudc.ConfigGetAllParams{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCSystemConfigSet(in *cloudc.SystemConfigSetParams) (*cloudc.ConfigResp, error) {
	res, err := c.CloudC.SystemConfigSet(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCSystemConfigGet(in *cloudc.ConfigGetParams) (*cloudc.SystemConfigGetResp, error) {
	res, err := c.CloudC.SystemConfigGet(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) CloudCSystemConfigGetAll() (*cloudc.SystemConfigGetAllResp, error) {
	res, err := c.CloudC.SystemConfigGetAll(c.SonyCtx(), &cloudc.ConfigGetAllParams{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
