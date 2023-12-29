package sdk

import (
	"errors"
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/cloudc"
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/wechat"
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk/types"
	"github.com/iotdreamcatcher/third_party/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Wechat struct {
	wechat.WechatRpcServiceClient
	Num         int64
	Status      int
	Retry       int
	Configs     map[string]*cloudc.ModelWechatConfig
	UsingConfig *cloudc.ModelWechatConfig
}

func NewWechat(RpcClientConf *zrpc.RpcClientConf) *Wechat {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	client := wechat.NewWechatRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &Wechat{
		WechatRpcServiceClient: client,
		Num:                    1,
	}
}

func (s *Wechat) GetAllConfigs() map[string]*cloudc.ModelWechatConfig {
	return s.Configs
}

func (s *Wechat) GetConfig() (*cloudc.ModelWechatConfig, error) {
	if s.UsingConfig != nil {
		return s.UsingConfig, nil
	}
	return nil, errors.New("not set config")
}

func (s *Wechat) SetConfig(key string) error {
	if _, ok := s.Configs[key]; !ok {
		return errors.New("not exist")
	}
	s.UsingConfig = s.Configs[key]
	return nil
}

func (c *Sdk) InitWechat() *Sdk {
	c.Wechat = NewWechat(c.Config.RpcClientConf)
	c.Wechat.Status = types.STATUS_READY
	c.Wechat.Retry += 1
	_, err := c.WechatLoadCloudConfig()
	if err != nil {
		logx.Errorf("oss cloud config sync err: %v+", err)
		return c
	}
	return c
}

func (c *Sdk) WechatCheckStatus() *Sdk {
	if c.Wechat.Status != types.STATUS_READY {
		if c.Wechat.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitWechat()
		}
	}
	return c
}

func (c *Sdk) WechatLoadCloudConfig() (*Sdk, error) {
	// note: 加载云端配置
	res, err := c.CloudCCheckStatus().CloudCWechatConfigGetAll()
	if err != nil {
		logx.Errorf("wechat cloud config sync err: %v+", err)
		return c, err
	}
	if res.Code != response.SUCCESS {
		logx.Errorf("wechat cloud config sync err: %v+", res.Msg)
		return c, errors.New(res.Msg)
	}
	temp := make(map[string]*cloudc.ModelWechatConfig)
	temp = res.Data.Configs
	// note: 判断是否存在default的配置文件，如果不存在，给予用户警告提示
	if _, ok := temp["default"]; !ok {
		logx.Alert("请注意，您的账号下没有default的配置文件，将会导致微信模块初始化失败，发送前请SetConfig(key string)")
	}

	c.Wechat.Configs = temp
	c.Wechat.UsingConfig = temp["default"]
	return c, nil
}

func (c *Sdk) WechatWebRedirectWechat(in *wechat.WebRedirectReq) (*wechat.WebRedirectResp, error) {
	// note: 读取当前使用配置
	if c.Wechat.UsingConfig == nil {
		c.Wechat.UsingConfig = c.Wechat.Configs["default"]
	}
	res, err := c.Wechat.WebRedirectWechat(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) WechatCode2Token(in *wechat.CodeReq) (*wechat.CodeResp, error) {
	// note: 读取当前使用配置
	if c.Wechat.UsingConfig == nil {
		c.Wechat.UsingConfig = c.Wechat.Configs["default"]
	}
	res, err := c.Wechat.Code2Token(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) WechatRefreshUserToken(in *wechat.RefreshReq) (*wechat.RefreshResp, error) {
	// note: 读取当前使用配置
	if c.Wechat.UsingConfig == nil {
		c.Wechat.UsingConfig = c.Wechat.Configs["default"]
	}
	res, err := c.Wechat.RefreshUserToken(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) WechatUserToken2UserInfo(in *wechat.TokenReq) (*wechat.UserInfoResp, error) {
	// note: 读取当前使用配置
	if c.Wechat.UsingConfig == nil {
		c.Wechat.UsingConfig = c.Wechat.Configs["default"]
	}
	res, err := c.Wechat.UserToken2UserInfo(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) WechatOfficialAccountAccessToken(in *wechat.OaKeyReq) (*wechat.OaAccessTokenResp, error) {
	// note: 读取当前使用配置
	if c.Wechat.UsingConfig == nil {
		c.Wechat.UsingConfig = c.Wechat.Configs["default"]
	}
	res, err := c.Wechat.OfficialAccountAccessToken(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) WechatWebAutoRedirectWechat(in *wechat.WebAutoRedirectReq) (*wechat.WebAutoRedirectResp, error) {
	// note: 读取当前使用配置
	if c.Wechat.UsingConfig == nil {
		c.Wechat.UsingConfig = c.Wechat.Configs["default"]
	}
	res, err := c.Wechat.WebAutoRedirectWechat(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
