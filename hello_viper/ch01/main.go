package main

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/liuyongbing/hello-go-imooc-ccmouse/hello_viper/structure"
)

func main() {
	v := viper.New()
	v.SetConfigFile("../config.yaml")

	// 读取配置文件内容
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	fmt.Println(v.Get("name"))

	// 将配置文件内容映射到 配置 struct
	sererConfig := structure.ServerConfig{}
	if err := v.Unmarshal(&sererConfig); err != nil {
		panic(err)
	}
	fmt.Println(sererConfig)
	fmt.Printf("%V", v.Get("name"))
}
