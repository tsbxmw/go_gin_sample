# go gin sample

依托于 `gin_common` 构建的脚手架。

https://github.com/tsbxmw/gin_common

# 依赖

- github.com/tsbxmw/gin_common


# 使用方法

## 下载代码

```shell
git clone https://github.com/tsbxmw/go_gin_sample
```

### 放置到 GOPATH/src 下

```shell
cp -rf go_gin_sample $GOPATH/src/
```

### 编译

```shell
go build app/main.go
```


### 直接运行

```go
go run app/main.go --config=./project/config/dev.json --mode=debug httpserver
```