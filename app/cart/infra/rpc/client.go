package rpc

import (
	"sync"

	"github.com/YiD11/gomall/app/cart/conf"
	cartUtils "github.com/YiD11/gomall/app/cart/utils"
	"github.com/YiD11/gomall/common/clientsuite"
	productservice "github.com/YiD11/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	ProductClient productservice.Client

	CurrentServiceName = conf.GetConf().Kitex.Service
	RegistryAddr       = conf.GetConf().Registry.RegistryAddress
	once          sync.Once
	err           error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	var opts []client.Option
	opts = append(opts, client.WithSuite(clientsuite.CommonClientSuite{CurrentServiceName: CurrentServiceName, RegistryAddr: RegistryAddr}))
	// r, err := etcd.NewEtcdResolver(conf.GetConf().Registry.RegistryAddress)
	// cartUtils.MustHandleErr(err)
	// opts = append(opts, client.WithResolver(r))

	ProductClient, err = productservice.NewClient("product", opts...)
	cartUtils.MustHandleErr(err)
}
