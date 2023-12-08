package sdk

import (
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/iot"
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk/types"
	"github.com/zeromicro/go-zero/zrpc"
)

type Iot struct {
	iot.IotServiceClient
	Num    int64
	Status int
	Retry  int
}

func NewIot(RpcClientConf *zrpc.RpcClientConf) *Iot {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	iotHomeServiceClient := iot.NewIotServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &Iot{
		IotServiceClient: iotHomeServiceClient,
		Num:              1,
	}
}

func (c *Sdk) InitIot() *Sdk {
	c.Iot = NewIot(c.Config.RpcClientConf)
	c.Iot.Status = types.STATUS_READY
	c.Iot.Retry += 1
	return c
}

func (c *Sdk) IotCheckStatus() *Sdk {
	if c.Iot.Status != types.STATUS_READY {
		if c.Iot.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitIot()
		}
	}
	return c
}

func (c *Sdk) IotHomeCreate(in *iot.CreateHomeReq) (*iot.Response, error) {
	res, err := c.Iot.HomeCreate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotHomeUpdate(in *iot.UpdateHomeReq) (*iot.Response, error) {
	res, err := c.Iot.HomeUpdate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotHomeDelete(in *iot.DeleteReq) (*iot.Response, error) {
	res, err := c.Iot.HomeDelete(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotHomeDeleteIds(in *iot.DeleteIdsReq) (*iot.Response, error) {
	res, err := c.Iot.HomeDeleteIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotHomeQuery(in *iot.QueryReq) (*iot.QueryHomeResp, error) {
	res, err := c.Iot.HomeQuery(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotHomeQueryList(in *iot.QueryHomeListReq) (*iot.QueryHomeListResp, error) {
	res, err := c.Iot.HomeQueryList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotHomeUpdateStatus(in *iot.UpdateStatusReq) (*iot.Response, error) {
	res, err := c.Iot.HomeUpdateStatus(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotMemberCreate(in *iot.CreateMemberReq) (*iot.Response, error) {
	res, err := c.Iot.MemberCreate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotMemberUpdate(in *iot.UpdateMemberReq) (*iot.Response, error) {
	res, err := c.Iot.MemberUpdate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotMemberDelete(in *iot.DeleteReq) (*iot.Response, error) {
	res, err := c.Iot.MemberDelete(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotMemberDeleteIds(in *iot.DeleteIdsReq) (*iot.Response, error) {
	res, err := c.Iot.MemberDeleteIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotMemberQuery(in *iot.QueryReq) (*iot.QueryMemberResp, error) {
	res, err := c.Iot.MemberQuery(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotMemberQueryList(in *iot.QueryMemberListReq) (*iot.QueryMemberListResp, error) {
	res, err := c.Iot.MemberQueryList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotMemberUpdateStatus(in *iot.UpdateStatusReq) (*iot.Response, error) {
	res, err := c.Iot.HomeUpdateStatus(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotRoomCreate(in *iot.CreateRoomReq) (*iot.Response, error) {
	res, err := c.Iot.RoomCreate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotRoomUpdate(in *iot.UpdateRoomReq) (*iot.Response, error) {
	res, err := c.Iot.RoomUpdate(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotRoomDelete(in *iot.DeleteReq) (*iot.Response, error) {
	res, err := c.Iot.RoomDelete(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotRoomDeleteIds(in *iot.DeleteIdsReq) (*iot.Response, error) {
	res, err := c.Iot.RoomDeleteIds(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotRoomQuery(in *iot.QueryReq) (*iot.QueryRoomResp, error) {
	res, err := c.Iot.RoomQuery(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotRoomQueryList(in *iot.QueryRoomListReq) (*iot.QueryRoomListResp, error) {
	res, err := c.Iot.RoomQueryList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) IotRoomUpdateStatus(in *iot.UpdateStatusReq) (*iot.Response, error) {
	res, err := c.Iot.HomeUpdateStatus(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
