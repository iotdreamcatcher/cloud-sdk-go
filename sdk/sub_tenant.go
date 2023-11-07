/*
*

	@author: taco
	@Date: 2023/10/27
	@Time: 11:07

*
*/
package sdk

import (
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/tenant"
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk/types"
	"github.com/zeromicro/go-zero/zrpc"
)

type SubTenant struct {
	tenant.TenantRpcServiceClient
	Num    int64
	Status int
	Retry  int
}

func NewSubTenant(RpcClientConf *zrpc.RpcClientConf) *SubTenant {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	client := tenant.NewTenantRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &SubTenant{
		TenantRpcServiceClient: client,
		Num:                    1,
	}
}

func (c *Sdk) InitSubTenant() *Sdk {
	c.SubTenant = NewSubTenant(c.Config.RpcClientConf)
	c.SubTenant.Status = types.STATUS_READY
	c.SubTenant.Retry += 1
	return c
}

func (c *Sdk) SubTenantCheckStatus() *Sdk {
	if c.SubTenant.Status != types.STATUS_READY {
		if c.SubTenant.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitSubTenant()
		}
	}
	return c
}

func (c *Sdk) PwdTenant(in *tenant.EditTenantParams) (*tenant.TenantResp, error) {
	res, err := c.SubTenant.PwdTenant(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) EditTenant(in *tenant.EditTenantParams) (*tenant.TenantResp, error) {
	res, err := c.SubTenant.EditTenant(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) DeleteTenant(in *tenant.DeleteTenantParams) (*tenant.TenantResp, error) {
	res, err := c.SubTenant.DeleteTenant(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) TenantInfo(in *tenant.TenantInfoParams) (*tenant.TenantInfoResp, error) {
	res, err := c.SubTenant.TenantInfo(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) TenantList(in *tenant.TenantListParams) (*tenant.TenantListResp, error) {
	res, err := c.SubTenant.TenantList(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) AddTenant(in *tenant.EditTenantParams) (*tenant.AddTenantResp, error) {
	res, err := c.SubTenant.AddTenant(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
