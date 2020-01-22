package main

import (
	"fmt"
	common "github.com/tsbxmw/gin_common"
	"go_gin_sample/project/transport/http"
	"os"
)

func main() {
	httpServer := http.HttpServer{}
	config := common.ServiceConfigImpl{}
	app, err := common.App("car_phone_v1", "wechat notification for car user", httpServer, config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
