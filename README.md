# qweather-sdk-go

[pkg.go.dev](https://pkg.go.dev/github.com/taadis/qweather-sdk-go)

## 开发文档

本SDK以和风天气官方的开发文档为主，字段和请求可以直接对照官方的开发文档

[开发文档|和风天气开发服务](https://dev.qweather.com/docs/api/)

## 单元测试

```
go test -v .
```

## 安装依赖

在你的Go项目中执行以下命令来下载并安装依赖

```
go get github.com/taadis/qweather-sdk-go
```

执行成功后,可以在你的Go项目go.mod依赖文件中看到此变更内容.

## 使用示例

```go
package main

import (
	"log"

	"github.com/taadis/qweather-sdk-go"
)

func main() {
	// new client
	client := qweather.NewClient()

	// top city request
	topCityReq := qweather.NewV2TopCityRequest()
	topCityReq.Key = "your_qweather_key"
	topCityRsp, err := client.V2TopCity(topCityReq)
	if err != nil {
		log.Fatalf("client.V2TopCity error:%+v", err)
	}

	log.Printf("client.V2TopCity rsp:%s", topCityRsp.String())

	// others request
}

```

启用日志

内置了 `log/slog` 默认日志示例，可配置选项参数 `WithLogEnabled(true)` 来启用日志输出，以便排查定位问题

```go
client := NewClient(WithLogEnabled(true))

```

使用自定义 `log/slog`

也可以通过选项参数 `WithLogger(logger)` 来传入自定义的 `log/slog`

```go
logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelInfo,
}))

client := NewClient(WithLogger(logger), WithLogEnabled(true))

```

> your_qweather_key 为你在和风天气服务注册并申请的应用key,不要填错了哦.

还没有注册的小伙伴可以从这里开始 [和风天气开发服务](https://dev.qweather.com/)

更多接口示例可以参考此代码库中的各类_test.go结尾的测试示例文件.