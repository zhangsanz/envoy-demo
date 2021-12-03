Envoy demo
==========

来源：

https://github.com/linux-china/envoy-demo

https://github.com/linux-china/grpc-go-demo

解决了一些问题后，将整个项目更新如当前。

后续会继续补充README

# 使用

需要先在grpc-go-demo目录下编译，并使用Dockerfile构建镜像

```shell
# 进入目录
cd grpc-go-demo
# 编译
# 需要有go环境
go mod init github.com/linux_china/grpc-go-demo
go mod tidy
go build
# 重命名
mv grpc-go-demo app

# 构建Dockerfile
docker build -t linuxchina/grpc-go-demo .
```

在envoy-demo目录下使用docker-compose启动所有的服务

```shell
docker-compose up
```

打开http://localhost:3000/ 访问Grafana dashboard,初始账户与密码为admin/admin



# Envoy testing

* MySQL：password is 123456

```
mysql -h 127.0.0.1 -P 3307 -u root -p demo
```

* httpbin

```
curl http://localhost:10800/ip
```

* grpc testing：需要先从https://github.com/ktr0731/evans安装evans

```shell
# evans
# 进入proto目录
cd envoy-demo/proto
# 下载evans源码安装
wget https://github.com/ktr0731/evans/releases/download/0.10.0/evans_linux_amd64.tar.gz
# 解压
tar -zvxf evans_linux_amd64.tar.gz
# 复制
cp evans /usr/local/bin
# evans的简单使用
# 参考链接：https://github.com/ktr0731/evans
# 在protp目录下
evans repl greeter.proto
# 使用
show package
package greeter_api
show service
service Greeter
call SayHello
# 示例
greeter_api.Greeter@127.0.0.1:50051> call SayHello
name (TYPE_STRING) => a
{
  "message": "Hello a!"
}
```

* baidu service test

```
curl http://localhost:10000/
```



