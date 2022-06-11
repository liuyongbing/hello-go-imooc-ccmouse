package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

/*
Register 服务注册
*/
func Register(addr string, port int, name string, tags []string, id string) {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 服务健康检查对象
	check := api.AgentServiceCheck{
		HTTP:                           "http://192.168.31.141:8021/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	// 服务注册对象
	registration := api.AgentServiceRegistration{
		Name:    name,
		ID:      id,
		Tags:    tags,
		Port:    port,
		Address: addr,
		Check:   &check,
	}

	// 注册服务
	err = client.Agent().ServiceRegister(&registration)
	if err != nil {
		panic(err)
	}
}

/*
AllServices 发现服务
*/
func AllServices() {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}

	for key, _ := range data {
		fmt.Println(key)
	}
}

/*
FilterServices 发现服务
*/
func FilterServices() {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	filter := `Service == "user-web"`
	data, err := client.Agent().ServicesWithFilter(filter)
	if err != nil {
		panic(err)
	}

	for key, _ := range data {
		fmt.Println(key)
	}
}

func main() {
	// 服务注册
	addr := "127.0.0.1"
	port := 8021
	name := "user-web"
	id := "user-web"
	tags := []string{
		"consul-demo",
		"gosrv-register",
	}
	Register(addr, port, name, tags, id)

	// 服务发现
	AllServices()

	// 服务发现: 筛选
	FilterServices()
}
