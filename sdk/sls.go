package sdk

import (
	"errors"
	"fmt"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/cloudc"
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk/types"
	"github.com/iotdreamcatcher/cloud-sdk-go/third_party/response"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
	"os"
	"os/signal"
	"time"
)

type Sls struct {
	Producer    *producer.Producer
	Num         int64
	Status      int
	Retry       int
	Configs     map[string]*cloudc.ModelSlsConfig
	UsingConfig *cloudc.ModelSlsConfig
}

func NewSls() *Sls {
	return &Sls{
		Num: 1,
	}
}

func (s *Sls) GetAllConfigs() map[string]*cloudc.ModelSlsConfig {
	return s.Configs
}

func (s *Sls) GetConfig() (*cloudc.ModelSlsConfig, error) {
	if s.UsingConfig != nil {
		return s.UsingConfig, nil
	}
	return nil, errors.New("not set config")
}

func (s *Sls) SetConfig(key string) error {
	if _, ok := s.Configs[key]; !ok {
		return errors.New("not exist")
	}
	s.UsingConfig = s.Configs[key]
	return nil
}

func (c *Sdk) InitSls() *Sdk {
	c.Sls = NewSls()
	_, err := c.SlsLoadCloudConfig()
	if err != nil {
		logx.Errorf("Sls cloud config sync err: %v+", err)
		return c
	}
	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = c.Sls.UsingConfig.Endpoint
	credentialsProvider := sls.NewStaticCredentialsProvider(c.Sls.UsingConfig.AppID, c.Sls.UsingConfig.AppSecret, "")
	producerConfig.CredentialsProvider = credentialsProvider
	producerInstance := producer.InitProducer(producerConfig)
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Kill, os.Interrupt)
	producerInstance.Start() // 启动producer实例
	c.Sls.Producer = producerInstance
	c.Sls.Status = types.STATUS_READY
	c.Sls.Retry += 1
	return c
}

func (c *Sdk) SlsCheckStatus() *Sdk {
	if c.Sls.Status != types.STATUS_READY {
		if c.Sls.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitSls()
		}
	}
	return c
}

func (c *Sdk) SlsLoadCloudConfig() (*Sdk, error) {
	// note: 加载云端配置
	res, err := c.CloudCCheckStatus().CloudCSlsConfigGetAll()
	if err != nil {
		logx.Errorf("sls cloud config sync err: %v+", err)
		return c, err
	}
	if res.Code != response.SUCCESS {
		logx.Errorf("sls cloud config sync err: %v+", res.Msg)
		return c, errors.New(res.Msg)
	}
	temp := make(map[string]*cloudc.ModelSlsConfig)
	temp = res.Data.Configs
	// note: 判断是否存在default的配置文件，如果不存在，给予用户警告提示
	if _, ok := temp["default"]; !ok {
		logx.Alert("请注意，您的账号下没有default的配置文件，将会导致日志服务发送失败，发送前请SetConfig(key string)")
	}

	c.Sls.Configs = temp
	c.Sls.UsingConfig = temp["default"]
	return c, nil
}

func (c *Sdk) SlsSendLog(log *sls.Log, sourceIp string) error {
	// note: 读取当前使用配置
	if c.Sls.UsingConfig == nil {
		c.Sls.UsingConfig = c.Sls.Configs["default"]
	}
	//logx.Infof("打印当前日志push的目的地: 配置key: %s", c.Sls.UsingConfig.Key)
	//logx.Infof("打印当前日志push的目的地: 配置名字: %s", c.Sls.UsingConfig.Name)
	//logx.Infof("打印当前日志push的目的地: project: %s", c.Sls.UsingConfig.ProjectName)
	//logx.Infof("打印当前日志push的目的地: logStoreName: %s", c.Sls.UsingConfig.LogStoreName)
	//logx.Infof("打印当前日志push的目的地: topic: %s", c.Sls.UsingConfig.Topic)
	err := c.Sls.Producer.SendLog(c.Sls.UsingConfig.ProjectName, c.Sls.UsingConfig.LogStoreName, c.Sls.UsingConfig.Topic, sourceIp, log)
	if err != nil {
		return err
	}
	return nil
}

type TemplateLog struct {
	AppName     string `json:"app_name"`
	AppPlatform string `json:"app_platform"`
	RequestId   string `json:"requestId"`
	Time        string `json:"time"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	Body        string `json:"body"`
	Message     string `json:"message"`
	UserId      string `json:"user_id"`
	UniqueId    string `json:"unique_id"`
}

func (c *Sdk) NewSlsLogMsg(s *TemplateLog) *sls.Log {
	content := []*sls.LogContent{}
	nowTime := time.Now().Unix()
	content = append(content, &sls.LogContent{
		Key:   proto.String("app_name"),
		Value: proto.String(s.AppName),
	})
	content = append(content, &sls.LogContent{
		Key:   proto.String("app_platform"),
		Value: proto.String(s.AppPlatform),
	})
	content = append(content, &sls.LogContent{
		Key:   proto.String("requestId"),
		Value: proto.String(s.RequestId),
	})

	tempTime, _ := time.Parse(types.TimeLayout, s.Time)
	content = append(content, &sls.LogContent{
		Key:   proto.String("timestamp"),
		Value: proto.String(fmt.Sprintf("%d", tempTime.Unix())),
	})
	content = append(content, &sls.LogContent{
		Key:   proto.String("time"),
		Value: proto.String(s.Time),
	})
	content = append(content, &sls.LogContent{
		Key:   proto.String("path"),
		Value: proto.String(s.Path),
	})
	content = append(content, &sls.LogContent{
		Key:   proto.String("method"),
		Value: proto.String(s.Method),
	})
	content = append(content, &sls.LogContent{
		Key:   proto.String("body"),
		Value: proto.String(s.Body),
	})
	content = append(content, &sls.LogContent{
		Key:   proto.String("message"),
		Value: proto.String(s.Message),
	})
	content = append(content, &sls.LogContent{
		Key:   proto.String("user_id"),
		Value: proto.String(s.UserId),
	})
	content = append(content, &sls.LogContent{
		Key:   proto.String("unique_id"),
		Value: proto.String(s.UniqueId),
	})
	return &sls.Log{
		Time:     proto.Uint32(uint32(nowTime)),
		Contents: content,
	}
}
