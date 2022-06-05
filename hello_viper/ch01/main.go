package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println(v.Get("name"))
}
