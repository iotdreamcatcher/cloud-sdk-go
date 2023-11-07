package sdk

import (
	"github.com/zeromicro/go-zero/zrpc"
	"time"
)

const (
	defaultTimeout = 5000
)

type Config struct {
	AutoRetry        bool          `default:"false"`
	MaxRetryTimes    int           `default:"3"`
	Debug            bool          `default:"false"`
	Timeout          time.Duration `default:"5000"`
	AutoRefreshToken bool          `default:"true"`
	AccessKeyId      string
	AccessKeySecret  string
	RpcClientConf    *zrpc.RpcClientConf
}

func DefaultConfig(AccessKeyId, AccessKeySecret string, Endpoints []string) (config *Config) {
	config = &Config{
		AutoRetry:        false,
		MaxRetryTimes:    3,
		Debug:            false,
		Timeout:          defaultTimeout,
		AutoRefreshToken: true,
		AccessKeyId:      AccessKeyId,
		AccessKeySecret:  AccessKeySecret,
		RpcClientConf:    newRpcClientConf(Endpoints),
	}
	return
}

func NewConfig(AccessKeyId, AccessKeySecret string, Endpoints []string) (config *Config) {
	config = &Config{
		AutoRetry:        false,
		MaxRetryTimes:    3,
		Debug:            false,
		Timeout:          defaultTimeout,
		AutoRefreshToken: true,
		AccessKeyId:      AccessKeyId,
		AccessKeySecret:  AccessKeySecret,
		RpcClientConf:    newRpcClientConf(Endpoints),
	}
	return
}

func newRpcClientConf(Endpoints []string) *zrpc.RpcClientConf {
	return &zrpc.RpcClientConf{
		Endpoints: Endpoints,
		Timeout:   defaultTimeout,
		NonBlock:  true,
	}
}

func (c *Config) WithAutoRetry(isAutoRetry bool) *Config {
	c.AutoRetry = isAutoRetry
	return c
}

func (c *Config) WithMaxRetryTimes(MaxRetryTimes int) *Config {
	c.MaxRetryTimes = MaxRetryTimes
	return c
}

func (c *Config) WithDebug(Debug bool) *Config {
	c.Debug = Debug
	return c
}

func (c *Config) WithTimeout(Timeout time.Duration) *Config {
	c.Timeout = Timeout
	c.RpcClientConf.Timeout = int64(Timeout)
	return c
}

func (c *Config) WithAutoRefreshToken(AutoRefreshToken bool) *Config {
	c.AutoRefreshToken = AutoRefreshToken
	return c
}

func (c *Config) WithAccessKeyId(AccessKeyId string) *Config {
	c.AccessKeyId = AccessKeyId
	return c
}

func (c *Config) WithAccessKeySecret(AccessKeySecret string) *Config {
	c.AccessKeySecret = AccessKeySecret
	return c
}
