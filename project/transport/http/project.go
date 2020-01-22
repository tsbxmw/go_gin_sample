package http

import (
	"fmt"
	"go_gin_sample/project/config"
	"go_gin_sample/project/routers"
	"go_gin_sample/project/worker"
	"strconv"

	"github.com/gin-gonic/gin"
	common "github.com/tsbxmw/gin_common"
	"github.com/tsbxmw/gin_common/handler"
	"github.com/tsbxmw/gin_common/middleware"
)

type (
	HttpServer struct {
		common.HttpServerImpl
	}
)

func (httpServer HttpServer) ServeWorker(mode string) {
	if mode == "" || mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	common.InitDB(httpServer.DbUri)
	common.InitRedisPool("tcp", httpServer.RedisHost+":"+string(httpServer.RedisPort), httpServer.RedisPassword, httpServer.RedisDB)
	common.InitLogger()
	config.LoadConfig(gin.Mode())
	worker.CornWork()
}

func (httpServer HttpServer) Serve(mode string) {
	fmt.Println("httpserver", httpServer.SvcName)
	if mode == "" || mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	engin := gin.New()
	// init logger
	middleware.LoggerInit(engin, "./log/carphone.log")
	// init tracer
	//middleware.TracerInit(engin, httpServer.JaegerAddr, httpServer.SvcName)
	//defer (*middleware.TracerCloser).Close()
	config.LoadConfig(gin.Mode())

	// init redis
	common.InitRedisPool("tcp", httpServer.RedisHost+":"+string(httpServer.RedisPort), httpServer.RedisPassword, httpServer.RedisDB)
	// init exception
	middleware.ExceptionInit(engin)
	handler.HandlerInit(engin)
	// init router
	routers.InitRouter(engin)
	// init consul
	//consulRegister := consul.ConsulRegister{
	//	Address:                        httpServer.Address,
	//	Port:                           httpServer.Port,
	//	ConsulAddress:                  httpServer.ConsulAddr,
	//	ConsulPort:                     httpServer.ConsulPort,
	//	Service:                        httpServer.SvcName,
	//	Tag:                            []string{httpServer.SvcName},
	//	DeregisterCriticalServiceAfter: time.Second * 60,
	//	Interval:                       time.Second * 60,
	//}
	//
	//consulRegister.RegisterHTTP()
	common.InitDB(httpServer.DbUri)
	common.LogrusLogger.Info("serve on " + strconv.Itoa(httpServer.Port))
	if err := engin.Run("0.0.0.0:" + strconv.Itoa(httpServer.Port)); err != nil {
		panic(err)
	}
}

func (httpServer HttpServer) Shutdown() {
	//consulName := httpServer.SvcName + "-" + common.LocalIP()
	//common.LogrusLogger.Info("Consul Deregister Now ", consulName)
	//// init consul
	//consulRegister := consul.ConsulRegister{
	//	Address:                        httpServer.Address,
	//	Port:                           httpServer.Port,
	//	ConsulAddress:                  httpServer.ConsulAddr,
	//	ConsulPort:                     httpServer.ConsulPort,
	//	Service:                        httpServer.SvcName,
	//	Tag:                            []string{httpServer.SvcName},
	//	DeregisterCriticalServiceAfter: time.Second * 60,
	//	Interval:                       time.Second * 60,
	//}
	//consulClient := consulRegister.NewConsulClient()
	//if err := consulClient.Agent().ServiceDeregister(consulName); err != nil {
	//	common.LogrusLogger.Error(err)
	//	panic(err)
	//}
}

func (httpServer HttpServer) Init(config common.ServiceConfig, configPath string) common.HttpServer {
	configReal := config.ConfigFromFileName(configPath).(common.ServiceConfigImpl)
	httpServer.SvcName = configReal.ServiceName
	httpServer.Address = configReal.HttpAddr
	httpServer.Port = configReal.Port
	httpServer.DbUri = configReal.DbUri
	httpServer.ConsulAddr = configReal.ConsulAddr
	httpServer.JaegerAddr = configReal.JaegerAddr
	httpServer.ConsulPort = configReal.ConsulPort
	httpServer.RedisDB = configReal.RedisDB
	httpServer.RedisHost = configReal.RedisHost
	httpServer.RedisPassword = configReal.RedisPassword
	httpServer.RedisPort = configReal.RedisPort
	return httpServer
}
