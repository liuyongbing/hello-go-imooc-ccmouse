package main

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/liuyongbing/hello-go-imooc-ccmouse/hello_viper/structure"
)

func GetEnvInfo(env string) string {
	viper.AutomaticEnv()

	return viper.GetString(env)
}

func main() {
	// 根据环境变量加载配置文件
	debug := GetEnvInfo("PATH")
	pathStr := ""
	configFileMode := "prd"
	if pathStr != debug {
		configFileMode = "dev"
	}
	configFileName := fmt.Sprintf("../config-%s.yaml", configFileMode)

	v := viper.New()
	v.SetConfigFile(configFileName)

	// 读取配置文件内容
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	fmt.Println(v.Get("name"))

	// 将配置文件内容映射到 配置 struct
	serverConfig := structure.ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)
	fmt.Printf("%V", v.Get("name"))
	fmt.Println("")

	// 动态监控配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file channed:", e.Name)
		v.ReadInConfig()
		v.Unmarshal(&serverConfig)
		fmt.Println(serverConfig)
	})

	time.Sleep(time.Second * 100)
}
