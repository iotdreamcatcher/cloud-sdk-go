package sdk

import (
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/iot"
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk/types"
	"github.com/zeromicro/go-zero/zrpc"
)

type IotHome struct {
	iot.IotHomeServiceClient
	Num    int64
	Status int
	Retry  int
}

func NewIot(RpcClientConf *zrpc.RpcClientConf) *IotHome {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	iotHomeServiceClient := iot.NewIotHomeServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &IotHome{
		IotHomeServiceClient: iotHomeServiceClient,
		Num:                  1,
	}
}

func (c *Sdk) InitIotHome() *Sdk {
	c.IotHome = NewIot(c.Config.RpcClientConf)
	c.IotHome.Status = types.STATUS_READY
	c.IotHome.Retry += 1
	return c
}

func (c *Sdk) IotHomeCheckStatus() *Sdk {
	if c.IotHome.Status != types.STATUS_READY {
		if c.IotHome.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitIotHome()
		}
	}
	return c
}

func (c *Sdk) IotHomeCreate(in *iot.CreateHomeReq) (*iot.CreateResp, error) {
	res, err := c.IotHome.Create(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotHomeDelete(in *iot.DeleteReq) (*iot.Response, error) {
	res, err := c.IotHome.Delete(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
