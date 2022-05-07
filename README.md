# hello-go-imooc-ccmouse
慕课网 Go开发工程师 学习之路 https://class.imooc.com/sale/go

## 设置代理加速
```
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=on
```

## 生成存根文件: hello_grpc/proto
```
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```