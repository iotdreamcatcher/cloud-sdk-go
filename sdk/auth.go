package sdk

import (
	"context"
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/shield"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"time"
)

type Auth struct {
	shield.ShieldServiceClient
	Num     int64
	Timeout time.Duration
}

func NewAuth(RpcClientConf *zrpc.RpcClientConf) *Auth {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	clent := shield.NewShieldServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &Auth{
		clent,
		1,
		time.Duration(RpcClientConf.Timeout) * time.Millisecond,
	}
}

func (c *Auth) RequestPing() (*shield.PingResp, error) {
	ctx := context.Background()
	newCtx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()
	c.Num = c.Num + 1

	res, err := c.Ping(newCtx, &shield.PingParams{
		Num: c.Num,
	})

	logx.Infof("打印结果 %v", res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Auth) RequestLogin(AccessKeyId, AccessKeySecret string) (*shield.AuthApiResp, error) {
	ctx := context.Background()
	newCtx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()
	loginResult, err := c.AuthLoginWithApi(newCtx, &shield.AuthApiParams{
		AccessKey:    AccessKeyId,
		AccessSecret: AccessKeySecret,
	})
	if err != nil {
		return nil, err
	}
	return loginResult, nil
}

func (c *Auth) RequestRefresh(AccessKeyId, RefreshToken string) (*shield.RefreshTokenResp, error) {
	ctx := context.Background()
	newCtx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()
	result, err := c.RefreshToken(newCtx, &shield.RefreshTokenParams{
		AccessKey: AccessKeyId,
		Token:     RefreshToken,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
