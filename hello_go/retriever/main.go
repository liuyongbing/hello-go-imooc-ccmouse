package main

import (
	"fmt"

	"github.com/liuyongbing/hello-go-imooc-ccmouse/hello_go/retriever/testing"
)

// 1. 封装成 func retrieve()
func retrieve(url string) string {
	//response, err := http.Get(url)
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer response.Body.Close()
	//
	//bytes, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	panic("ioutil.ReadAll err" + err.Error())
	//}
	//
	//return string(bytes)

	// 2. 不能自己封装， 需调用架构组提供的方法
	//retriever := infra.Retriever{}
	//contents := retriever.Get(url)
	//return contents

	// 3. 封装成 func getRetriever()
	retriever := getRetriever()
	contents := retriever.Get(url)
	return contents
}

// 3. 封装成 func getRetriever()
//func getRetriever() infra.Retriever {
//	return infra.Retriever{}
//}

// 4. 测试团队介入
//func getRetriever() testing.Retriever {
//	return testing.Retriever{}
//}

// 5. 兼容架构组 & 测试组， 申明 retriever interface, 并重写 func getRetriever()
type retriever interface {
	Get(string) string
}

func getRetriever() retriever {
	return testing.Retriever{}
}

func main() {

	// 0. 原生写法
	//response, err := http.Get("https://www.imooc.com")
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer response.Body.Close()
	//
	//bytes, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	panic("ioutil.ReadAll err" + err.Error())
	//}
	//fmt.Printf("%s\n", bytes)

	url := "https://www.imooc.com"
	contents := retrieve(url)
	fmt.Println(contents)
}
