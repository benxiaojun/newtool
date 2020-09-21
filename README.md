
## 运行容器

```
docker run -it --rm -v ${pwd}:/go/src -w /go/src golang
```

## 设置GOPROXY代理

```
go env -w GOPROXY=https://goproxy.cn,direct
```

## 编译

cd 进入到需要编译的文件目录，比如 cd /app/satool/hangjia/

```
go build .
```

可能要用到下面命令编译  否则会出现容器中无法执行程序 报not found
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
```