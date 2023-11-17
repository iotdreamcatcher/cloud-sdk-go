package sdk

import (
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/push"
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk/types"
	"github.com/zeromicro/go-zero/zrpc"
)

type Push struct {
	push.PushRpcServiceClient
	Num    int64
	Status int
	Retry  int
}

func NewPush(RpcClientConf *zrpc.RpcClientConf) *Push {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	client := push.NewPushRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &Push{
		PushRpcServiceClient: client,
		Num:                  1,
	}
}

func (c *Sdk) InitPush() *Sdk {
	c.Push = NewPush(c.Config.RpcClientConf)
	c.Push.Status = types.STATUS_READY
	c.Push.Retry += 1
	return c
}

func (c *Sdk) PushCheckStatus() *Sdk {
	if c.Push.Status != types.STATUS_READY {
		if c.Push.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitPush()
		}
	}
	return c
}

func (c *Sdk) PushIosPush(in *push.IosPushReq) (*push.PushResp, error) {
	res, err := c.Push.IosPush(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) PushHuaweiPush(in *push.HuaweiPushReq) (*push.PushResp, error) {
	res, err := c.Push.HuaweiPush(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) PushXiaomiPush(in *push.XiaomiPushReq) (*push.PushResp, error) {
	res, err := c.Push.XiaomiPush(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
