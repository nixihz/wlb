package clients

import (
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"wlb/internal/conf"
)

// LarkCfg 飞书sdk client 配置文件
type LarkCfg struct {
	Name       string
	Id         string
	Token      string
	AppSecret  string
	EncryptKey string
}

func NewLarkClient(data *conf.Data) *lark.Client {
	larkCfg := LarkCfg{
		Name:       data.Lark.FeishuAppName,
		Token:      data.Lark.FeishuAppVerificationToken,
		AppSecret:  data.Lark.FeishuAppSecret,
		EncryptKey: data.Lark.FeishuAppEventEncryptKey,
		Id:         data.Lark.FeishuAppId,
	}
	return lark.NewClient(larkCfg.Id, larkCfg.AppSecret, lark.WithLogLevel(larkcore.LogLevelDebug))
}
