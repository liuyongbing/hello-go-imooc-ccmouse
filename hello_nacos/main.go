package main

import (
	"encoding/json"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"github.com/liuyongbing/hello-go-imooc-ccmouse/hello_nacos/config"
)

func main() {
	// Nacos 服务器配置
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}

	// Nacos 客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         "229535b5-2256-4339-a1f3-66d06605344d",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建 Nacos 客户端连接
	nacosClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}

	// 读取 Nacos 配置信息
	configInfo, err := nacosClient.GetConfig(vo.ConfigParam{
		DataId: "user-web",
		Group:  "dev",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Nacos 原始配置内容：")
	fmt.Println(configInfo)

	serverConfig := config.ServerConfig{}
	json.Unmarshal([]byte(configInfo), &serverConfig)

	fmt.Println("本地转换后内容：")
	fmt.Println(serverConfig)

	// 监听 Nacos 配置信息变化
	// err = nacosClient.ListenConfig(vo.ConfigParam{
	// 	DataId: "user-web",
	// 	Group:  "dev",
	// 	OnChange: func(namespace, group, dataId, data string) {
	// 		fmt.Println("配置文件变化")
	// 	},
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// time.Sleep(time.Second * 20)
}
