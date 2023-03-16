[toc]
# 文档

```go
https://go-zero.dev/cn/docs/introduction //新
https://legacy.go-zero.dev/cn/goctl.html //老
```



# 准备工作

```go
//下载go-zero go mo tidy


//下载goctl
go install github.com/zeromicro/go-zero/tools/goctl@latest

//下载protoc
https://github.com/protocolbuffers/protobuf/releases

//下载protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

//下载protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

//工具
https://github.com/Mikaelemmmm/go-zero-looklook
```

# goctl命令

```go
//goctl命令，api文件生成代码
goctl api go -api *.api -dir ../  --style=goZero

//goctl命令，proto:文件生成代码
goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=goZero


//goctl命令，生成Dockerfile
goctl  docker -go user.go

//生成k8s部署文件
goctl kube deploy --name user --namespace sg-bs -image user:latest  -o user.yml -port 8080 -nodePort 31001
```

# goctl model

通过表反向自动生成模型，包含表的增删改查

**userModel.sh chmod 777 userModel.sh**

```sh
#!/usr/bin/env bash

# 使用方法：
# ./genModel.sh bsmaster users
# ./genModel.sh bsmaster bike
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package


#生成的表名
tables=$2
#表生成的genmodel目录
modeldir=./${tables}Model

# 数据库配置
host=192.168.0.214
port=32432
dbname=$1 # 库名
username=root
passwd=123456


echo "开始创建库：$dbname 的表：$tables"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=true --style=goZero

#-cache=false 如果需要redis则打开
```

# goctl 生成api文档 goctl api doc

```go
goctl api doc --dir ./
```

# grpcui独立调试rpc工具

```go
grpcui工具独立调试grpc接口

//下载安装工具
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest

//启动工具
grpcui -plaintext localhost:8080
gRPC Web UI available at http://127.0.0.1:60551/...
```

![](https://img2023.cnblogs.com/blog/1736414/202211/1736414-20221127222544843-1555555890.png)
