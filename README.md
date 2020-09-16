
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