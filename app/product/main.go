package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/YiD11/gomall/app/product/biz/dal"
	"github.com/YiD11/gomall/app/product/conf"
	"github.com/YiD11/gomall/common/mtl"
	"github.com/YiD11/gomall/common/serversuite"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"

	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	MetricsPort  = conf.GetConf().Kitex.MetricsPort
	RegistryAddr = conf.GetConf().Registry.RegistryAddress

	err error
)

func main() {
	if os.Getenv("GO_ENV") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Panicln(err)
		}
	}

	dal.Init()

	mtl.InitMetric(ServiceName, MetricsPort, RegistryAddr)
	p := mtl.InitTracing(ServiceName)
	defer p.Shutdown(context.Background())
	opts := kitexInit()

	opts = append(opts, server.WithSuite(serversuite.CommonServerSuite{CurrentServiceName: ServiceName, RegistryAddr: RegistryAddr}))

	svr := productcatalogservice.NewServer(new(ProductCatalogServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		log.Panicln(err)
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	opts = append(opts, server.WithSuite(serversuite.CommonServerSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))

	// service info
	// opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
	// 	ServiceName: conf.GetConf().Kitex.Service,
	// }))

	// registry
	// r, err := etcd.NewEtcdRegistry(conf.GetConf().Registry.RegistryAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// opts = append(opts, server.WithRegistry(r))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
