package sdk

import (
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/cloudc"
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/syncer"
	"github.com/zeromicro/go-zero/zrpc"
)

type Syncer struct {
	syncer.SyncerRpcServiceClient
	Num         int64
	Status      int
	Retry       int
	Configs     map[string]*cloudc.ModelOssConfig
	UsingConfig *cloudc.ModelOssConfig
}

func NewScyner(RpcClientConf *zrpc.RpcClientConf) *Syncer {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	client := syncer.NewSyncerRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &Syncer{
		SyncerRpcServiceClient: client,
		Num:                    1,
	}
}
