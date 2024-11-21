package rpc

import (
	"sync"

	"github.com/YiD11/gomall/app/checkout/conf"
	"github.com/YiD11/gomall/common/clientsuite"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	productservice "github.com/YiD11/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
)

var (
	CartClient    cartservice.Client
	ProductClient productservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once

	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress

	err error
)

func InitClient() {
	once.Do(func() {
		initCartClient()
		initPaymentClient()
		initProductClient()
		initOrderClient()
	})
}

func initCartClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		klog.Fatal(err)
	}
}

func initProductClient() {
	var opts []client.Option

	opts = append(opts, client.WithSuite(clientsuite.CommonClientSuite{CurrentServiceName: ServiceName, RegistryAddr: RegistryAddr}))

	// r, err := etcd.NewEtcdResolver(conf.GetConf().Registry.RegistryAddress)
	// if err != nil {
	// 	klog.Fatal(err)
	// }

	// opts = append(opts,
	// client.WithResolver(r),
	// client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service,}),
	// client.WithTransportProtocol(transport.GRPC),
	// client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	// )

	ProductClient, err = productservice.NewClient("product", opts...)
	if err != nil {
		klog.Fatal(err)
	}
}

func initPaymentClient() {
	var opts []client.Option

	opts = append(opts, client.WithSuite(clientsuite.CommonClientSuite{CurrentServiceName: ServiceName, RegistryAddr: RegistryAddr}))

	// r, err := etcd.NewEtcdResolver(conf.GetConf().Registry.RegistryAddress)
	// if err != nil {
	// 	klog.Fatal(err)
	// }

	// opts = append(opts,
	// 	client.WithResolver(r),
	// client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service,}),
	// 	client.WithTransportProtocol(transport.GRPC),
	// 	client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	// )

	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		klog.Fatal(err)
	}
}

func initOrderClient() {
	var opts []client.Option

	opts = append(opts, client.WithSuite(clientsuite.CommonClientSuite{CurrentServiceName: ServiceName, RegistryAddr: RegistryAddr}))

	// r, err := etcd.NewEtcdResolver(conf.GetConf().Registry.RegistryAddress)
	// if err != nil {
	// 	klog.Fatal(err)
	// }
	// opts = append(opts,
	// 	client.WithResolver(r),
	// 	client.WithTransportProtocol(transport.GRPC),
	// 	client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	// )

	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		klog.Fatal(err)
	}
}
